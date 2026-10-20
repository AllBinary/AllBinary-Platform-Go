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

type EmptyRaceRaceTrackGeographicMapCellTypeFactory struct {
    RaceTrackGeographicMapCellTypeFactory
}

var EmptyRaceRaceTrackGeographicMapCellTypeFactory_InstanceC *EmptyRaceRaceTrackGeographicMapCellTypeFactory = NewEmptyRaceRaceTrackGeographicMapCellTypeFactory()


func EmptyRaceRaceTrackGeographicMapCellTypeFactory_GetInstance() *EmptyRaceRaceTrackGeographicMapCellTypeFactory {
        return EmptyRaceRaceTrackGeographicMapCellTypeFactory_InstanceC
}


func NewEmptyRaceRaceTrackGeographicMapCellTypeFactory() *EmptyRaceRaceTrackGeographicMapCellTypeFactory {
    this := &EmptyRaceRaceTrackGeographicMapCellTypeFactory{}
    this.RaceTrackGeographicMapCellTypeFactory = *NewRaceTrackGeographicMapCellTypeFactory()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.EMPTY_CELL_TYPE = NewRaceTrackGeographicMapCellType(commonStrings.EMPTY, 0, PathFindingNodeCostInfoData_GetInstance().MAX_NODE_COST, 1)
        this.EASY_CELL_TYPE = NewRaceTrackGeographicMapCellType("Easy", 19, 1, 0)
        this.FINISH_LINE_ROAD_CELL_TYPE = this.DEFAULT_FINISH_LINE_ROAD_CELL_TYPE
    return this
}

