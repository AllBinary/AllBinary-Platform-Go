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

type CustomMapGenerator struct {
    CustomMapGeneratorBase
    CustomMapArray [][]int
    RaceTrackGeographicMap *RaceTrackGeographicMap
    GeographicMapCellTypeFactory IGeographicMapCellTypeFactory
    DropCellPositionHistory *DropCellPositionHistory
}


func (this *CustomMapGenerator) initFields() {
    
    
    
    this.DropCellPositionHistory = DropCellPositionHistory_GetInstance()
}


func NewCustomMapGenerator(raceTrackGeographicMap *RaceTrackGeographicMap) *CustomMapGenerator {
    this := &CustomMapGenerator{}
    this.CustomMapGeneratorBase = *NewCustomMapGeneratorBase()
    this.initFields()
        this.RaceTrackGeographicMap = raceTrackGeographicMap
        this.GeographicMapCellTypeFactory = this.RaceTrackGeographicMap.GetGeographicMapCellTypeFactory()
        var mapArray [][]int = raceTrackGeographicMap.GetRaceTrackData().GetMapArray(); _ = mapArray
        var size2 int = len(mapArray[0]); _ = size2
        this.CustomMapArray = func() [][]int { array := make([][]int, len(mapArray)); for index := range array { array[index] = make([]int, size2) }; return array }()
        if len(mapArray) != len(this.CustomMapArray) || len(mapArray[0]) != len(this.CustomMapArray[0]) {
            panic(errors.New(fmt.Sprint("Array Incorrect")))
        }
    return this
}


// throws Exception
func (this *CustomMapGenerator) CopyMapIntoCustomMap() {
        var mapArray [][]int = this.RaceTrackGeographicMap.GetRaceTrackData().GetMapArray(); _ = mapArray
        var startIndex2 int = len(mapArray[0]) - 1; _ = startIndex2
        for index := int(len(mapArray) - 1); index >= 0; index-- {
            for index2 := int(startIndex2); index2 >= 0; index2-- {
                this.CustomMapArray[index][index2] = this.GetCustomType(index2, index, mapArray[index][index2])
            }
        }
}


// throws Exception
func (this *CustomMapGenerator) GetCustomType(column int, row int, currentType int) int {
        var emptyType int = this.GeographicMapCellTypeFactory.GetEmptyType(); _ = emptyType
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.RaceTrackGeographicMap.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var cellPosition ICellPosition = geographicMapCellPositionFactory.GetAt(column, row); _ = cellPosition
        if this.DropCellPositionHistory.IsCellPositionWithDrop(cellPosition) {
            var layer IAllBinaryLayer = this.DropCellPositionHistory.GetLayerInterface(cellPosition); _ = layer
            if layer == AllBinaryLayer_NULL_ALLBINARY_LAYER {
                return emptyType
            }
            var rtsLayer IAdvancedRTSGameLayer = any(layer).(IAdvancedRTSGameLayer); _ = rtsLayer
            if !(rtsLayer.GetType() == WaypointLayer_GetStaticType()) {
                return emptyType
            }
        }
        return currentType
}


func (this *CustomMapGenerator) GetCustomMapArray() [][]int {
        return this.CustomMapArray
}

