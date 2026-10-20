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

type PathFindingNodeCostInfoFactory struct {
    PathFindingNodeCostInfoFactoryBase
}



func NewPathFindingNodeCostInfoFactory(max_ int) *PathFindingNodeCostInfoFactory {
    this := &PathFindingNodeCostInfoFactory{}
    this.PathFindingNodeCostInfoFactoryBase = *NewPathFindingNodeCostInfoFactoryBase()
    return this
}


// throws Exception
func (this *PathFindingNodeCostInfoFactory) Create(geographicMapInterface IBasicGeographicMap, goingToGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition, costFromStart int64, costToEnd int64) {
}


// throws Exception
func (this *PathFindingNodeCostInfoFactory) GetTotalCost(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition) int64 {
        var geographicMapCellType IGeographicMapCellType = geographicMapInterface.GetCellTypeAt(comingFromGeographicMapCellPosition); _ = geographicMapCellType
        var geographicMapCellType2 IGeographicMapCellType = geographicMapInterface.GetCellTypeAt(geographicMapCellPosition); _ = geographicMapCellType2
        var raceTrackGeographicMapCellType *RaceTrackGeographicMapCellType = any(geographicMapCellType).(*RaceTrackGeographicMapCellType); _ = raceTrackGeographicMapCellType
        var raceTrackGeographicMapCellType2 *RaceTrackGeographicMapCellType = any(geographicMapCellType2).(*RaceTrackGeographicMapCellType); _ = raceTrackGeographicMapCellType2
        return int64(raceTrackGeographicMapCellType.GetTravelCost()) + int64(raceTrackGeographicMapCellType2.GetTravelCost())
}

