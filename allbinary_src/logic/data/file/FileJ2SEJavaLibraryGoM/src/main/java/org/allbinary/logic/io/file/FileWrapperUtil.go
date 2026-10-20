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

type FileWrapperUtil struct {
}


func NewFileWrapperUtil() *FileWrapperUtil {
    this := &FileWrapperUtil{}
    return this
}


func FileWrapperUtil_WrapFiles(files []any) []*AbFile {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []*AbFile) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutSE("Exception Wrapping Files", "FileWrapperUtil", "wrapFiles", e)
            tryControl1, tryResult1 = 1, make([]*AbFile, 0); return
                }
            }()
            var abFileArray []*AbFile = make([]*AbFile, len(files)); _ = abFileArray
            for index := int(len(files) - 1); index >= 0; index-- {
                abFileArray[index] = NewAbFile(any(files[index]).(*File))
            }
            tryControl1, tryResult1 = 1, abFileArray; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func FileWrapperUtil_WrapFile(file *File) *AbFile {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *AbFile) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutSE("Exception Wrapping File", "FileWrapperUtil", "wrapFile", e)
            tryControl1, tryResult1 = 1, AbFile_NULL_FILE; return
                }
            }()
            tryControl1, tryResult1 = 1, NewAbFile(file); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

