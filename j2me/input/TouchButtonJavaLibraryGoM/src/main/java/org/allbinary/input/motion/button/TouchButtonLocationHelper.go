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

type TouchButtonLocationHelper struct {
    TotalColumns int
    TotalRows int
    ColumnsRemainder int
    RowsRemainder int
    ColumnsRemainderHalf int
    RowsRemainderHalf int
}


func (this *TouchButtonLocationHelper) initFields() {
    
    
    
    
    
    
}


func NewTouchButtonLocationHelper() *TouchButtonLocationHelper {
    this := &TouchButtonLocationHelper{}
    this.initFields()
        var buttonSize int = CommonButtons_GetInstance().STANDARD_BUTTON_SIZE; _ = buttonSize
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.TotalColumns = int64(math.Floor(float64(float32(displayInfo.GetLastWidth()) / float32(buttonSize)) + 0.5))
        this.TotalRows = int64(math.Floor(float64(float32(displayInfo.GetLastHeight()) / float32(buttonSize)) + 0.5))
        this.ColumnsRemainder = displayInfo.GetLastWidth() % buttonSize
        this.ColumnsRemainderHalf = (this.ColumnsRemainder >> 1)
        this.RowsRemainder = displayInfo.GetLastHeight() % buttonSize
        this.RowsRemainderHalf = (this.RowsRemainder >> 1)
    return this
}


func (this *TouchButtonLocationHelper) GetTotalColumns() int {
        return this.TotalColumns
}


func (this *TouchButtonLocationHelper) GetTotalRows() int {
        return this.TotalRows
}


func (this *TouchButtonLocationHelper) GetColumnsRemainder() int {
        return this.ColumnsRemainder
}


func (this *TouchButtonLocationHelper) GetRowsRemainder() int {
        return this.RowsRemainder
}


func (this *TouchButtonLocationHelper) GetColumnsRemainderHalf() int {
        return this.ColumnsRemainderHalf
}


func (this *TouchButtonLocationHelper) GetRowsRemainderHalf() int {
        return this.RowsRemainderHalf
}


func (this *TouchButtonLocationHelper) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("totalColumns: ")
        stringBuffer.Appendint(this.TotalColumns)
        stringBuffer.Append(" totalRows: ")
        stringBuffer.Appendint(this.TotalRows)
        stringBuffer.Append(" columnsRemainder: ")
        stringBuffer.Appendint(this.ColumnsRemainder)
        stringBuffer.Append(" rowsRemainder: ")
        stringBuffer.Appendint(this.RowsRemainder)
        stringBuffer.Append(" columnsRemainderHalf: ")
        stringBuffer.Appendint(this.ColumnsRemainderHalf)
        stringBuffer.Append(" rowsRemainderHalf: ")
        stringBuffer.Appendint(this.RowsRemainderHalf)
        return fmt.Sprint(stringBuffer)
}

func (this *TouchButtonLocationHelper) String() string {
    return this.ToString()
}

