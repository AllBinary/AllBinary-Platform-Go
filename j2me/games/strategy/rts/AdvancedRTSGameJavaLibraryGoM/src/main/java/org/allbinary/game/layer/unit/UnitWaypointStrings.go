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

type UnitWaypointStrings struct {
    WANDERING string
    THINKING string
    THINKING_ABOUT_TARGET string
    TARGET string
    KILL string
    STOP string
    WAYPOINT_DESTROYED_SHORT string
    WAYPOINT_DESTROYED string
    ALL_VISITED_SHORT string
    ALL_VISITED string
    ALREADY_THERE_SHORT string
    ALREADY_THERE string
    NEXT_PATH_NODE string
}

var UnitWaypointStrings_Instance *UnitWaypointStrings = NewUnitWaypointStrings()

func (this *UnitWaypointStrings) initFields() {
    this.WANDERING = "Order?"
    this.THINKING = "Thinking"
    this.THINKING_ABOUT_TARGET = "Hmmm"
    this.TARGET = "Target"
    this.KILL = "Kill!"
    this.STOP = "Stop"
    this.WAYPOINT_DESTROYED_SHORT = "Uh Oh"
    this.WAYPOINT_DESTROYED = "Waypoint Destroyed"
    this.ALL_VISITED_SHORT = "Arrived"
    this.ALL_VISITED = "All Visited"
    this.ALREADY_THERE_SHORT = "Again?"
    this.ALREADY_THERE = "Already There"
    this.NEXT_PATH_NODE = fmt.Sprint(CommonLabels_GetInstance().NEXT) + " Path Node"
}

func NewUnitWaypointStrings() *UnitWaypointStrings {
    this := &UnitWaypointStrings{}
    this.initFields()
    return this
}


func UnitWaypointStrings_GetInstance() *UnitWaypointStrings {
        return UnitWaypointStrings_Instance
}

