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

type GeographicMapCellPositionTracking struct {
    CurrentGeographicMapCellPosition IGeographicMapCellPosition
    PreviousGeographicMapCellPosition IGeographicMapCellPosition
    NewGeographicMapCellPosition IGeographicMapCellPosition
    CurrentCellType IGeographicMapCellType
}


func (this *GeographicMapCellPositionTracking) initFields() {
    this.CurrentGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.PreviousGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.NewGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.CurrentCellType = GeographicMapCellType_NULL_GEOGRAPHIC_MAP_CELL_TYPE
}


func NewGeographicMapCellPositionTracking() *GeographicMapCellPositionTracking {
    this := &GeographicMapCellPositionTracking{}
    this.initFields()
    return this
}


// throws Exception
func (this *GeographicMapCellPositionTracking) GetCurrentOverCellType(geographicMap IBasicGeographicMap) IGeographicMapCellType {
        var geographicMapCellPosition IGeographicMapCellPosition = this.CurrentGeographicMapCellPosition; _ = geographicMapCellPosition
        if geographicMapCellPosition != this.PreviousGeographicMapCellPosition {
            this.CurrentCellType = geographicMap.GetCellTypeAt(geographicMapCellPosition)
        }
        return this.CurrentCellType
}


// throws Exception
func (this *GeographicMapCellPositionTracking) UpdateStart(gameCanvasInterface AllBinaryGameCanvasInterface, layer IAllBinaryLayer) bool {
        if gameCanvasInterface != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(gameCanvasInterface.GetLayerManager()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            this.NewGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(layer.GetXP() + layer.GetHalfWidth(), layer.GetYP() + layer.GetHalfHeight())
            if this.NewGeographicMapCellPosition != this.CurrentGeographicMapCellPosition {
                this.CurrentGeographicMapCellPosition = this.NewGeographicMapCellPosition
                return true
            }
        }
        return false
}


func (this *GeographicMapCellPositionTracking) UpdateEnd() bool {
        if this.CurrentGeographicMapCellPosition != this.PreviousGeographicMapCellPosition {
            this.PreviousGeographicMapCellPosition = this.CurrentGeographicMapCellPosition
            return true
        }
        return false
}


func (this *GeographicMapCellPositionTracking) GetPreviousGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.PreviousGeographicMapCellPosition
}


func (this *GeographicMapCellPositionTracking) GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.CurrentGeographicMapCellPosition
}

