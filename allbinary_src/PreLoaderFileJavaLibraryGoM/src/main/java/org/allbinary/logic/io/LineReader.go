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

type LineReader struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    FileName string
    String string
    Index int
    IdFile IAbFileInputStream
    IdData IDataInputStream
}


func (this *LineReader) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.Index = 0
    
    
}


func NewLineReader(fileName string) *LineReader {
    this := &LineReader{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("File: " + fileName, this, "LineReader", e)
            {
            }
                }
            }()
            var bytes []int8 = make([]int8, 100000); _ = bytes
            this.FileName = fileName
            this.IdFile = NewAbFileInputStream(fileName)
            this.IdFile.Read(bytes)
            this.String = string(bytes)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *LineReader) HasNext() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "hasNext", e)
            {
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var nextIndex int = this.String.IndexOf(CommonSeps_GetInstance().NEW_LINE, this.Index); _ = nextIndex
            if nextIndex == -1 {
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *LineReader) Next() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "next", e)
            {
            }
            panic(e)
                }
            }()
            var nextIndex int = this.String.IndexOf(CommonSeps_GetInstance().NEW_LINE, this.Index); _ = nextIndex
            if nextIndex == -1 {
                panic(errors.New(fmt.Sprint("next() should have been called first")))
            }
            var temp string = this.String[this.Index:nextIndex]; _ = temp
            this.Index = nextIndex + 1
            tryControl1, tryResult1 = 1, temp; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

