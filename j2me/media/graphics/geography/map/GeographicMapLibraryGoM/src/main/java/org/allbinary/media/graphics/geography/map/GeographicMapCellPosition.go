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

type GeographicMapCellPosition struct {
    CellPosition
    Point IGPoint
    MidPoint IGPoint
}

// The Java class GeographicMapCellPosition has subclasses so its variables are this interface
type IGeographicMapCellPosition interface {
    ICellPosition
    GeographicMapCellPosition_() *GeographicMapCellPosition
    GetPoint() IGPoint
    GetMidPoint() IGPoint
    SetMidPoint(midPoint IGPoint)
}

func (this *GeographicMapCellPosition) GeographicMapCellPosition_() *GeographicMapCellPosition {
    return this
}


func (this *GeographicMapCellPosition) initFields() {
    
    
}


func NewGeographicMapCellPosition(column int, row int, columns int, rows int, width int, height int) *GeographicMapCellPosition {
    this := &GeographicMapCellPosition{}
        this.CellPosition = *NewCellPosition(column, row, columns, rows)
        this.initFields()
        var x int = this.GetColumn() * width; _ = x
        var y int = this.GetRow() * height; _ = y
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        this.Point = pointFactory.CreateXY(x, y)
        this.MidPoint = pointFactory.CreateXY(x + (width >> 1), y + (height >> 1))
    return this
}


func (this *GeographicMapCellPosition) GetPoint() IGPoint {
        return this.Point
}


func (this *GeographicMapCellPosition) GetMidPoint() IGPoint {
        return this.MidPoint
}


func (this *GeographicMapCellPosition) SetMidPoint(midPoint IGPoint) {
        this.MidPoint = midPoint
}

