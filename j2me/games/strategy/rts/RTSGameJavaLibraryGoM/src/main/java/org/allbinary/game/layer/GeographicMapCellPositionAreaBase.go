/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type GeographicMapCellPositionAreaBase struct {
}

// Implements GeographicMapCellPositionAreaInterface

// The Java class GeographicMapCellPositionAreaBase has subclasses so its variables are this interface
type IGeographicMapCellPositionAreaBase interface {
    GeographicMapCellPositionAreaBase_() *GeographicMapCellPositionAreaBase
    GetNextSurroundingGeographicMapCellPosition() IGeographicMapCellPosition
    GetOccupyingGeographicMapCellPositionList() IBasicArrayList
    GetSurroundingGeographicMapCellPositionList() IBasicArrayList
    Update(geographicMapInterface IBasicGeographicMap)
}

func (this *GeographicMapCellPositionAreaBase) GeographicMapCellPositionAreaBase_() *GeographicMapCellPositionAreaBase {
    return this
}

var GeographicMapCellPositionAreaBase_NULL_GEOGRPAHIC_MAP_POSITION_AREA_BASE IGeographicMapCellPositionAreaBase = NewGeographicMapCellPositionAreaBase()

func NewGeographicMapCellPositionAreaBase() *GeographicMapCellPositionAreaBase {
    this := &GeographicMapCellPositionAreaBase{}
    return this
}


func (this *GeographicMapCellPositionAreaBase) GetNextSurroundingGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *GeographicMapCellPositionAreaBase) GetOccupyingGeographicMapCellPositionList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


func (this *GeographicMapCellPositionAreaBase) GetSurroundingGeographicMapCellPositionList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *GeographicMapCellPositionAreaBase) Update(geographicMapInterface IBasicGeographicMap) {
        panic(errors.New(fmt.Sprint()))
}

