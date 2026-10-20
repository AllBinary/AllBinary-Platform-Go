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

type RaceTrackGeographicMapCellTypes struct {
    BOTTOM_LEFT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    BOTTOM_RIGHT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    TOP_LEFT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    TOP_RIGHT_TURN_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    VERTICAL_STRAIGHT_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    START_LINE_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
    DEFAULT_FINISH_LINE_ROAD_CELL_TYPE *RaceTrackGeographicMapCellType
}

var RaceTrackGeographicMapCellTypes_Instance *RaceTrackGeographicMapCellTypes = NewRaceTrackGeographicMapCellTypes()

func (this *RaceTrackGeographicMapCellTypes) initFields() {
    
    
    
    
    
    
    
    
}


func RaceTrackGeographicMapCellTypes_GetInstance() *RaceTrackGeographicMapCellTypes {
        return RaceTrackGeographicMapCellTypes_Instance
}


func NewRaceTrackGeographicMapCellTypes() *RaceTrackGeographicMapCellTypes {
    this := &RaceTrackGeographicMapCellTypes{}
    this.initFields()
        this.BOTTOM_LEFT_TURN_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Bottom Left Turn", 1, 1, 0)
        this.BOTTOM_RIGHT_TURN_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Bottom Right Turn", 2, 1, 0)
        this.TOP_LEFT_TURN_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Top Left Turn", 3, 1, 0)
        this.TOP_RIGHT_TURN_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Top Right Turne", 4, 1, 0)
        this.HORIZONTAL_STRAIGHT_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Horizontal Straight", 5, 1, 0)
        this.VERTICAL_STRAIGHT_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Vertical Straight", 6, 1, 0)
        this.START_LINE_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Start Line", 7, 1, 0)
        this.DEFAULT_FINISH_LINE_ROAD_CELL_TYPE = NewRaceTrackGeographicMapCellType("Finish Line", 8, 1, 0)
    return this
}

