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

type LayerPartialCellPositionsUtil struct {
}

var LayerPartialCellPositionsUtil_Instance *LayerPartialCellPositionsUtil = NewLayerPartialCellPositionsUtil()

func NewLayerPartialCellPositionsUtil() *LayerPartialCellPositionsUtil {
    this := &LayerPartialCellPositionsUtil{}
    return this
}


// throws Exception
func (this *LayerPartialCellPositionsUtil) GetAll(geographicMapInterface IBasicGeographicMap, layerInterface IAllBinaryLayer, reusableList IBasicArrayList) IBasicArrayList {
        reusableList.Clear()
        var topLeftGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP(), layerInterface.GetYP()); _ = topLeftGeographicMapCellPosition
        var topRightGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + layerInterface.GetWidth(), layerInterface.GetYP()); _ = topRightGeographicMapCellPosition
        var bottomLeftGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP(), layerInterface.GetYP() + layerInterface.GetHeight()); _ = bottomLeftGeographicMapCellPosition
        var bottomRightGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + layerInterface.GetWidth(), layerInterface.GetYP() + layerInterface.GetHeight()); _ = bottomRightGeographicMapCellPosition
        if topLeftGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            reusableList.Add(topLeftGeographicMapCellPosition)
        }
        if topRightGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && topLeftGeographicMapCellPosition != topRightGeographicMapCellPosition {
            reusableList.Add(topRightGeographicMapCellPosition)
        }
        if bottomLeftGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && !reusableList.Contains(bottomLeftGeographicMapCellPosition) {
            reusableList.Add(bottomLeftGeographicMapCellPosition)
        }
        if bottomRightGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && !reusableList.Contains(bottomRightGeographicMapCellPosition) {
            reusableList.Add(bottomRightGeographicMapCellPosition)
        }
        return reusableList
}


// throws Exception
func (this *LayerPartialCellPositionsUtil) GetAllDXY(geographicMapInterface IBasicGeographicMap, layerInterface IAllBinaryLayer, dx int, dy int, reusableList IBasicArrayList) IBasicArrayList {
        reusableList.Clear()
        var topLeftGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + dx, layerInterface.GetYP() + dy); _ = topLeftGeographicMapCellPosition
        var topRightGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + layerInterface.GetWidth() + dx, layerInterface.GetYP() + dy); _ = topRightGeographicMapCellPosition
        var bottomLeftGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + dx, layerInterface.GetYP() + layerInterface.GetHeight() + dy); _ = bottomLeftGeographicMapCellPosition
        var bottomRightGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(layerInterface.GetXP() + layerInterface.GetWidth() + dx, layerInterface.GetYP() + layerInterface.GetHeight() + dy); _ = bottomRightGeographicMapCellPosition
        if topLeftGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            reusableList.Add(topLeftGeographicMapCellPosition)
        }
        if topRightGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && topLeftGeographicMapCellPosition != topRightGeographicMapCellPosition {
            reusableList.Add(topRightGeographicMapCellPosition)
        }
        if bottomLeftGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && !reusableList.Contains(bottomLeftGeographicMapCellPosition) {
            reusableList.Add(bottomLeftGeographicMapCellPosition)
        }
        if bottomRightGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && !reusableList.Contains(bottomRightGeographicMapCellPosition) {
            reusableList.Add(bottomRightGeographicMapCellPosition)
        }
        return reusableList
}


func LayerPartialCellPositionsUtil_GetInstance() *LayerPartialCellPositionsUtil {
        return LayerPartialCellPositionsUtil_Instance
}

