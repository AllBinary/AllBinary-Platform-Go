/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

//FileJ2SE should be the same as J2SEForJ2ME
type AbFileSystem struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var AbFileSystem_Instance *AbFileSystem = NewAbFileSystem()

func (this *AbFileSystem) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewAbFileSystem() *AbFileSystem {
    this := &AbFileSystem{}
    this.initFields()
    return this
}


func AbFileSystem_GetInstance() *AbFileSystem {
        return AbFileSystem_Instance
}


func (this *AbFileSystem) IsDirectoryOrFile(path string) bool {
        var file *File = NewFile(path); _ = file
        return file.IsDirectory() || file.IsFile()
}


func (this *AbFileSystem) IsDirectory(path string) bool {
        return NewFile(path).IsDirectory()
}


func (this *AbFileSystem) GetFilesAsStringArrayForPath(currentDirPath string) []string {
        var file *File = NewFile(currentDirPath); _ = file
        if file.Exists() {
            return file.List()
        } else {
            return StringUtil_GetInstance().GetArrayInstance()
        }
}


func (this *AbFileSystem) ReadAsString(fileName string) string {
        var bytes []int8 = make([]int8, 1000000); _ = bytes
        return this.ReadAsString(fileName, bytes)
}


// Java overloaded method: Go needs a different name
func (this *AbFileSystem) ReadAsString(fileName string, bytes []int8) string {
        var closeable any = NullUtil_GetInstance().NULL_OBJECT; _ = closeable
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "readAsString", e)
                }
            this.Close(closeable)
            }()
            var idFile *InputStream = NewFileInputStream(fileName); _ = idFile
            closeable = idFile
            var size int = idFile.Read(bytes); _ = size
            if size > 0 {
                tryControl1, tryResult1 = 1, string(bytes[0:0 + size]); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *AbFileSystem) Close(closeable any) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CLOSE, e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if closeable != nil {
                var inputStream *InputStream = (any(closeable).(*InputStream)); _ = inputStream
                inputStream.Close()
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

