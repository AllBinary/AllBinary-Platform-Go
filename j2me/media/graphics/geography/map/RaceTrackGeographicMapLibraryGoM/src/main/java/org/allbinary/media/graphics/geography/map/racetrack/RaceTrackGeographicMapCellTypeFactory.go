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

type RaceTrackGeographicMapCellTypeFactory struct {
    GeographicMapCellTypeFactory
    EMPTY_CELL_TYPE IGeographicMapCellType
    EASY_CELL_TYPE IGeographicMapCellType
    BOTTOM_LEFT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    BOTTOM_RIGHT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    TOP_LEFT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    TOP_RIGHT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    VERTICAL_STRAIGHT_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    START_LINE_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    DEFAULT_FINISH_LINE_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    FINISH_LINE_ROAD_CELL_TYPE IGeographicMapCellType
}

// The Java class RaceTrackGeographicMapCellTypeFactory has subclasses so its variables are this interface
type IRaceTrackGeographicMapCellTypeFactory interface {
    IGeographicMapCellTypeFactory
    RaceTrackGeographicMapCellTypeFactory_() *RaceTrackGeographicMapCellTypeFactory
    GetStartType() int
    GetEndType() int
    GetEmptyType() int
    IsPath(cellType IGeographicMapCellType) bool
}

func (this *RaceTrackGeographicMapCellTypeFactory) RaceTrackGeographicMapCellTypeFactory_() *RaceTrackGeographicMapCellTypeFactory {
    return this
}

var RaceTrackGeographicMapCellTypeFactory_InstanceC IRaceTrackGeographicMapCellTypeFactory = NewRaceTrackGeographicMapCellTypeFactory()

func (this *RaceTrackGeographicMapCellTypeFactory) initFields() {
    this.EMPTY_CELL_TYPE = NewGeographicMapCellType(0, -1)
    this.EASY_CELL_TYPE = this.EMPTY_CELL_TYPE
    
    
    
    
    
    
    
    
    this.FINISH_LINE_ROAD_CELL_TYPE = GeographicMapCellType_NULL_GEOGRAPHIC_MAP_CELL_TYPE
}


func RaceTrackGeographicMapCellTypeFactory_GetInstance() IRaceTrackGeographicMapCellTypeFactory {
        return RaceTrackGeographicMapCellTypeFactory_InstanceC
}


func NewRaceTrackGeographicMapCellTypeFactory() *RaceTrackGeographicMapCellTypeFactory {
    this := &RaceTrackGeographicMapCellTypeFactory{}
    this.GeographicMapCellTypeFactory = *NewGeographicMapCellTypeFactory()
    this.initFields()
        var raceTrackGeographicMapCellTypes *RaceTrackGeographicMapCellTypes = RaceTrackGeographicMapCellTypes_GetInstance(); _ = raceTrackGeographicMapCellTypes
        this.BOTTOM_LEFT_TURN_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.BOTTOM_LEFT_TURN_ROAD_CELL_TYPE
        this.BOTTOM_RIGHT_TURN_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.BOTTOM_RIGHT_TURN_ROAD_CELL_TYPE
        this.TOP_LEFT_TURN_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.TOP_LEFT_TURN_ROAD_CELL_TYPE
        this.TOP_RIGHT_TURN_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.TOP_RIGHT_TURN_ROAD_CELL_TYPE
        this.HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE
        this.VERTICAL_STRAIGHT_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.VERTICAL_STRAIGHT_ROAD_CELL_TYPE
        this.START_LINE_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.START_LINE_ROAD_CELL_TYPE
        this.DEFAULT_FINISH_LINE_ROAD_CELL_TYPE = raceTrackGeographicMapCellTypes.DEFAULT_FINISH_LINE_ROAD_CELL_TYPE
    return this
}


func (this *RaceTrackGeographicMapCellTypeFactory) GetStartType() int {
        return this.START_LINE_ROAD_CELL_TYPE.GetType()
}


func (this *RaceTrackGeographicMapCellTypeFactory) GetEndType() int {
        return this.FINISH_LINE_ROAD_CELL_TYPE.GetType()
}


func (this *RaceTrackGeographicMapCellTypeFactory) GetEmptyType() int {
        return this.EMPTY_CELL_TYPE.GetType()
}


func (this *RaceTrackGeographicMapCellTypeFactory) IsPath(cellType IGeographicMapCellType) bool {
        if cellType != this.EMPTY_CELL_TYPE && cellType != this.EASY_CELL_TYPE {
            return true
        }
        return false
}

