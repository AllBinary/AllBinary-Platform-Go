/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSGameStrings struct {
    DRAG_TO_SPOT string
    NOT_YOURS string
    NEW_UNIT string
    NEW_WAYPOINT string
    BUILDING string
    UPGRADE string
    DOWNGRADE string
    NO_MONEY string
    SPOT_TAKEN string
    BUILDING_COLLISION string
    STRUCTURE_TO_CLOSE string
    MAP_EDGE string
    ROAD_COLLISION string
    SELECT_BUILD_SPOT string
    BUILD_ON_PATH string
    DRAGGABLE string
}

var RTSGameStrings_Instance *RTSGameStrings = NewRTSGameStrings()

func (this *RTSGameStrings) initFields() {
    this.DRAG_TO_SPOT = "Please Drag"
    this.NOT_YOURS = "Not Yours"
    this.NEW_UNIT = "New Unit"
    this.NEW_WAYPOINT = "New Waypoint"
    this.BUILDING = "Building"
    this.UPGRADE = "Upgrading"
    this.DOWNGRADE = "Downgrading"
    this.NO_MONEY = "Not Enough Money"
    this.SPOT_TAKEN = "Spot Taken"
    this.BUILDING_COLLISION = "Building Collision"
    this.STRUCTURE_TO_CLOSE = "Building To Close"
    this.MAP_EDGE = "Map Edge"
    this.ROAD_COLLISION = "Road Collision"
    this.SELECT_BUILD_SPOT = "Select Build Spot"
    this.BUILD_ON_PATH = "Not On Path"
    this.DRAGGABLE = "Draggable"
}

func NewRTSGameStrings() *RTSGameStrings {
    this := &RTSGameStrings{}
    this.initFields()
    return this
}


func RTSGameStrings_GetInstance() *RTSGameStrings {
        return RTSGameStrings_Instance
}

