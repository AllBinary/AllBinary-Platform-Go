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

type CellPosition struct {
    V
    Id int
    Column int
    Row int
}

// The Java class CellPosition has subclasses so its variables are this interface
type ICellPosition interface {
    IV
    CellPosition_() *CellPosition
    GetColumn() int
    GetRow() int
    ToString() string
    GetId() int
}

func (this *CellPosition) CellPosition_() *CellPosition {
    return this
}


func (this *CellPosition) initFields() {
    
    
    
}


func NewCellPosition(column int, row int, columns int, rows int) *CellPosition {
    this := &CellPosition{}
    this.V = *NewV()
    this.initFields()
        this.Column = column
        this.Row = row
        this.Id = (row * columns) + column
    return this
}


func (this *CellPosition) GetColumn() int {
        return this.Column
}


func (this *CellPosition) GetRow() int {
        return this.Row
}


func (this *CellPosition) ToString() string {
        return CellPosition_ToStringColRow(this.GetColumn(), this.GetRow())
}

func (this *CellPosition) String() string {
    return this.ToString()
}


func CellPosition_ToStringCellPosition(basicGeographicMapCellPosition ICellPosition) string {
        return CellPosition_ToStringColRow(basicGeographicMapCellPosition.GetColumn(), basicGeographicMapCellPosition.GetRow())
}


func CellPosition_ToStringColRow(i_column int, i_row int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        stringBuffer.Append(commonSeps.PARENTHESIS_OPEN)
        stringBuffer.Appendint(i_column)
        stringBuffer.Append(commonSeps.COMMA)
        stringBuffer.Appendint(i_row)
        stringBuffer.Append(commonSeps.PARENTHESIS_CLOSE)
        return fmt.Sprint(stringBuffer)
}


func (this *CellPosition) GetId() int {
        return this.Id
}

