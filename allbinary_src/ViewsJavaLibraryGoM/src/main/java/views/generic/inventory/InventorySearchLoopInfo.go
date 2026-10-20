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

type InventorySearchLoopInfo struct {
    LastPage int
    CurrentPage int
    NumberOfResultsOnCurrentPage int
    NumberOfResults int
    CURRENT string
    NUM string
    LAST string
}


func (this *InventorySearchLoopInfo) initFields() {
    this.LastPage = -1
    this.CurrentPage = -1
    this.NumberOfResultsOnCurrentPage = 0
    this.NumberOfResults = 0
    this.CURRENT = "current: "
    this.NUM = " num: "
    this.LAST = " last: "
}

func NewInventorySearchLoopInfo() *InventorySearchLoopInfo {
    this := &InventorySearchLoopInfo{}
    this.initFields()
    return this
}


func (this *InventorySearchLoopInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.CURRENT)
        stringBuffer.Appendint(this.CurrentPage)
        stringBuffer.Append(this.NUM)
        stringBuffer.Appendint(this.NumberOfResults)
        stringBuffer.Append(this.LAST)
        stringBuffer.Appendint(this.LastPage)
        return fmt.Sprint(stringBuffer)
}

func (this *InventorySearchLoopInfo) String() string {
    return this.ToString()
}

