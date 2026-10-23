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

type ZipFileUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var ZipFileUtil_Instance *ZipFileUtil = NewZipFileUtil()

func (this *ZipFileUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewZipFileUtil() *ZipFileUtil {
    this := &ZipFileUtil{}
    this.initFields()
    return this
}


func ZipFileUtil_GetInstance() *ZipFileUtil {
        return ZipFileUtil_Instance
}


func (this *ZipFileUtil) Create(outFilename string, fileBasicArrayList IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "create", e)
                }
            }()
            var zipFile *AbFile = AbFile_CreateAbFile(outFilename); _ = zipFile
            if zipFile.Exists() {
                zipFile.Delete()
            }
            zipFile.CreateNewFile()
            var outputStream *ZipOutputStream = NewZipOutputStream(AbFileOutputStream_CreateFromAbFile(zipFile)); _ = outputStream
            var streamUtil *StreamUtil = StreamUtil_GetInstance(); _ = streamUtil
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var byteArray []int8 = make([]int8, 16384); _ = byteArray
            var fileInputStream IAbFileInputStream; _ = fileInputStream
            var size int = fileBasicArrayList.Size(); _ = size
            var current int = 0; _ = current
            for i := int(0); i < size; i++ {
                var file *AbFile = any(fileBasicArrayList.Get(i)).(*AbFile); _ = file
                if file.IsDirectory() {
                } else if file.IsFile() {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Appendint(current)
                    stringBuffer.Append(" Creating Zip File Entry: ")
                    stringBuffer.Append(file.GetPath())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "create()")
                    if tryControl2 := func() (tryControl2 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        this.LogUtil.PutF("Skipping File (Probably Local): " + fmt.Sprint(file.GetPath()), this, "create()")
                            }
                        }()
                        fileInputStream = NewAbFileInputStream(file)
                        outputStream.PutNextEntry(NewZipEntry(file.GetPath()))
                        streamUtil.Get(fileInputStream, outputStream, byteArray)
                        outputStream.CloseEntry()
                        streamUtil.Close(fileInputStream)
                        return
                    }(); tryControl2 == 1 {
                        tryControl1 = 1; return
                    }
                }
                current++
            }
            streamUtil.Close(outputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *ZipFileUtil) Unzip(path string, zipFile *AbFile, fileName string) {
        var buffer []int8 = make([]int8, 16384); _ = buffer
        var inputStream *ZipInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(inputStream)
            }()
            inputStream = NewZipInputStream(NewAbFileInputStream(zipFile))
            var zipEntry *ZipEntry; _ = zipEntry
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var fileUtil *FileUtil = FileUtil_GetInstance(); _ = fileUtil
            for (func() any { zipEntry = inputStream.GetNextEntry(); return zipEntry }()) != nil {
                var entryName string = zipEntry.GetName(); _ = entryName
                if zipEntry.IsDirectory() {
                } else {
                    var entryFile *AbFile = AbFile_CreateAbFile(path + entryName); _ = entryFile
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("ZipEntry Name: ")
                    stringBuffer.Append(entryName)
                    stringBuffer.Append(" isFile: ")
                    stringBuffer.Appendboolean(entryFile.IsFile())
                    stringBuffer.Append(" isDirectory: ")
                    stringBuffer.Appendboolean(entryFile.IsDirectory())
                    stringBuffer.Append(" getParent: ")
                    stringBuffer.Append(entryFile.GetParent())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "unzip")
                    var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(entryFile); _ = dataOutputStream
                    fileUtil.Write(inputStream, dataOutputStream, buffer)
                }
                inputStream.CloseEntry()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

