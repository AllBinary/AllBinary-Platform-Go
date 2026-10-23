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

type DeleteCloud struct {
    LogUtil *LogUtil
}


func (this *DeleteCloud) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewDeleteCloud() *DeleteCloud {
    this := &DeleteCloud{}
    this.initFields()
    return this
}


func (this *DeleteCloud) Delete(prePath string, current int, total int) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to copy installer files into cloud", this, "initialize()", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if AbIOSystem_GetInstance().IsType("com.vobject.appengine.java.io") {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append(URLGLOBALS.GetWebappPath())
                stringBuffer.Append(prePath)
                var path IAbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = path
                var file *AbFile = AbFile_CreateAbFileFromAbPath(path); _ = file
                var fileBasicArrayList IBasicArrayList = Directory_GetInstance().Search(file, true); _ = fileBasicArrayList
                var size int = fileBasicArrayList.Size(); _ = size
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Searched: ")
                stringBuffer.Append(path.ToFileSystemString())
                stringBuffer.Append(" BasicArrayList: ")
                stringBuffer.Appendint(size)
                var portion int = size / total + 1; _ = portion
                var start int = portion * current; _ = start
                var end int = start + portion; _ = end
                if end > size {
                    end = size
                }
                stringBuffer.Append(" Section: ")
                stringBuffer.Appendint(start)
                stringBuffer.Append(" - ")
                stringBuffer.Appendint(end)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "initialize()")
                var nextFile *AbFile; _ = nextFile
                for index := int(start); index < end; index++ {
                    nextFile = any(fileBasicArrayList.Get(index)).(*AbFile)
                    if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 bool) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                            }
                        }()
                        if strings.Index(nextFile.GetPath(), string(FileUploadData_GetInstance().FILE)) < 0 {
                            nextFile.Delete()
                        }
                        return
                    }(); tryControl2 == 1 {
                        tryControl1, tryResult1 = 1, tryResult2; return
                    }
                }
                this.LogUtil.PutF("Deleted Files From Cloud", this, "initialize()")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

