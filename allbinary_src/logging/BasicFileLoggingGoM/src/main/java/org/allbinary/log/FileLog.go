/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type FileLog struct {
}

var FileLog_StringUtil *StringUtil = StringUtil_GetInstance()
var FileLog_LogLength int64 = 33000000
var FileLog_LogPath string = "g:\\log\\"
var FileLog_Extension string = "dat"
var FileLog_FileName string = "log." + fmt.Sprint(FileLog_Extension)
var FileLog_BackupFileName string = FileLog_FileName.Concat(".bak")
var FileLog_ORG_ALLBINARY string = "org.allbinary: "
var FileLog_FirstTime bool = true
var FileLog_LogFile *File = NewFile(FileLog_LogPath, FileLog_FileName)
var FileLog_LogFileBak *File
var FileLog_FileOut *BufferedWriter
var FileLog_BackupIndex int


func NewFileLog() *FileLog {
    this := &FileLog{}
    return this
}


// synchronized
func FileLog_CreateLogFile() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println("Error Creating Log: " + fmt.Sprint(e))
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            FileLog_LogFile = NewFile(FileLog_LogPath, FileLog_FileName)
            if !FileLog_FirstTime {
                FileLog_FileOut = NewBufferedWriter(NewFileWriter(FileLog_LogFile))
            } else {
                FileLog_FirstTime = false
                var raFile *RandomAccessFile = NewRandomAccessFile(FileLog_LogFile, "rw"); _ = raFile
                raFile.Seek(raFile.Length())
                FileLog_FileOut = NewBufferedWriter(NewFileWriter(raFile.GetFD()))
            }
            var canWrite bool = FileLog_LogFile.CanWrite(); _ = canWrite
            tryControl1, tryResult1 = 1, canWrite; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func FileLog_CreateLogFileBackup() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println("Error Creating Backup: " + fmt.Sprint(e))
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
            FileLog_LogFileBak = NewFile(FileLog_LogPath, fmt.Sprint(stringMaker.Append(FileLog_BackupFileName).Append(CommonSeps_GetInstance().PERIOD).Appendint(FileLog_BackupIndex)))
            for FileLog_LogFileBak.IsFile() {
                FileLog_BackupIndex++
                stringMaker.Delete(0, stringMaker.Length())
                FileLog_LogFileBak = NewFile(FileLog_LogPath, fmt.Sprint(stringMaker.Append(FileLog_BackupFileName).Append(CommonSeps_GetInstance().PERIOD).Appendint(FileLog_BackupIndex)))
            }
            var line string = ""; _ = line
            var tmpOut *BufferedWriter = NewBufferedWriter(NewFileWriter(FileLog_LogFileBak)); _ = tmpOut
            var tmpIn *BufferedReader = NewBufferedReader(NewFileReader(FileLog_LogFile)); _ = tmpIn
            for (func() any { line = tmpIn.ReadLine(); return line }()) != nil {
                tmpOut.Write(line, 0, len(line))
                tmpOut.NewLine()
            }
            tmpOut.Flush()
            tmpOut.Close()
            tmpIn.Close()
            FileLog_LogFile.Delete()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func FileLog_Put(specialMessage string, anyType any, functionName string) string {
        return FileLog_Put(specialMessage, anyType, functionName, nil)
}


// Java overloaded method: Go needs a different name
// synchronized
func FileLog_Put(specialMessage string, anyType any, functionName string, exception error) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "Error: " + fmt.Sprint(e); return
                }
            }()
            if FileLog_FirstTime == true {
                FileLog_CreateLogFileBackup()
                if FileLog_CreateLogFile() == false {
                    tryControl1, tryResult1 = 1, "Log File Creation Error"; return
                }
            }
            var length int64 = int64(FileLog_LogFile.Length()); _ = length
            if length > FileLog_LogLength {
                if !(FileLog_CreateLogFileBackup()) || FileLog_CreateLogFile() == false {
                    tryControl1, tryResult1 = 1, "Couldn't Create a Backup file"; return
                }
            }
            var className string = FileLog_StringUtil.NULL_STRING; _ = className
            if functionName == "" {
                functionName = FileLog_StringUtil.NULL_STRING
            }
            if specialMessage == "" {
                specialMessage = FileLog_StringUtil.NULL_STRING
            }
            if reflect.TypeOf(anyType).String() != "" {
                className = reflect.TypeOf(anyType).String()
            }
            var message string = LogFormatUtil_GetInstance().Get(className, functionName, specialMessage, exception); _ = message
            FileLog_FileOut.Write(message, 0, len(message))
            FileLog_FileOut.NewLine()
            FileLog_FileOut.Flush()
            tryControl1, tryResult1 = 1, fmt.Sprint(FileLog_ORG_ALLBINARY) + message; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// Java overloaded method: Go needs a different name
// synchronized
func FileLog_Put(specialMessage string, className string, functionName string) string {
        return FileLog_Put(specialMessage, className, functionName, nil)
}


// Java overloaded method: Go needs a different name
// synchronized
func FileLog_Put(specialMessage string, className string, functionName string, exception error) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "Error: " + fmt.Sprint(e); return
                }
            }()
            if FileLog_FirstTime == true {
                FileLog_CreateLogFileBackup()
                if FileLog_CreateLogFile() == false {
                    tryControl1, tryResult1 = 1, "Log File Creation Error"; return
                }
            }
            var length int64 = int64(FileLog_LogFile.Length()); _ = length
            if length > FileLog_LogLength {
                if !(FileLog_CreateLogFileBackup()) || FileLog_CreateLogFile() == false {
                    tryControl1, tryResult1 = 1, "Couldn't Create a Backup file"; return
                }
            }
            if functionName == "" {
                functionName = FileLog_StringUtil.NULL_STRING
            }
            if specialMessage == "" {
                specialMessage = FileLog_StringUtil.NULL_STRING
            }
            var message string = LogFormatUtil_GetInstance().Get(className, functionName, specialMessage, exception); _ = message
            FileLog_FileOut.Write(message, 0, len(message))
            FileLog_FileOut.NewLine()
            FileLog_FileOut.Flush()
            tryControl1, tryResult1 = 1, fmt.Sprint(FileLog_ORG_ALLBINARY) + message; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func FileLog_GetFilePath() string {
        return FileLog_LogPath + FileLog_FileName
}

