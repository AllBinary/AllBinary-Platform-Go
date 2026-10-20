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

type GeographicMapCellPositionEvent struct {
    AllBinaryEventObject
    GeographicMapCellPosition ICellPosition
}


func (this *GeographicMapCellPositionEvent) initFields() {
    this.GeographicMapCellPosition = CellPositionFactory_GetInstance().NONE
}


func NewGeographicMapCellPositionEvent(anyType any, geographicMapCellPosition ICellPosition) *GeographicMapCellPositionEvent {
    this := &GeographicMapCellPositionEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
        this.GeographicMapCellPosition = geographicMapCellPosition
    return this
}


func (this *GeographicMapCellPositionEvent) Init(anyType any) {
        this.SetSource(anyType)
}


func (this *GeographicMapCellPositionEvent) SetGeographicMapCellPosition(geographicMapCellPosition IGeographicMapCellPosition) {
        this.GeographicMapCellPosition = geographicMapCellPosition
}


func (this *GeographicMapCellPositionEvent) GetGeographicMapCellPosition() IGeographicMapCellPosition {
        return any(this.GeographicMapCellPosition).(IGeographicMapCellPosition)
}


func (this *GeographicMapCellPositionEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("CellPositionEvent: ")
        stringBuffer.Append("\nGeographicMapCellPosition: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetGeographicMapCellPosition()))
        return fmt.Sprint(stringBuffer)
}

func (this *GeographicMapCellPositionEvent) String() string {
    return this.ToString()
}

