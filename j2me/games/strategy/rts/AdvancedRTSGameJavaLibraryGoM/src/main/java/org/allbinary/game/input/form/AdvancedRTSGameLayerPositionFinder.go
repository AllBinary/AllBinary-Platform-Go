/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type AdvancedRTSGameLayerPositionFinder struct {
}

// Implements LayerPositionFinderInterface

var AdvancedRTSGameLayerPositionFinder_Instance *AdvancedRTSGameLayerPositionFinder = NewAdvancedRTSGameLayerPositionFinder()


func AdvancedRTSGameLayerPositionFinder_GetInstance() *AdvancedRTSGameLayerPositionFinder {
        return AdvancedRTSGameLayerPositionFinder_Instance
}


func NewAdvancedRTSGameLayerPositionFinder() *AdvancedRTSGameLayerPositionFinder {
    this := &AdvancedRTSGameLayerPositionFinder{}
    return this
}


func (this *AdvancedRTSGameLayerPositionFinder) GetLayerInterface(geographicMapCellPosition IGeographicMapCellPosition) IAllBinaryLayer {
        var layerInterface IAllBinaryLayer = DropCellPositionHistory_GetInstance().GetLayerInterface(geographicMapCellPosition); _ = layerInterface
        if layerInterface == AllBinaryLayer_NULL_ALLBINARY_LAYER {
            layerInterface = WaypointCellPositionHistory_GetInstance().GetLayerInterface(geographicMapCellPosition)
        }
        return layerInterface
}

