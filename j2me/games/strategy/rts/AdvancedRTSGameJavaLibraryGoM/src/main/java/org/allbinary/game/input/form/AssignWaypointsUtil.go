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

type AssignWaypointsUtil struct {
    LogUtil *LogUtil
    WAYPOINT_EVENT *RTSLayerEvent
}

var AssignWaypointsUtil_Instance *AssignWaypointsUtil = NewAssignWaypointsUtil()

func (this *AssignWaypointsUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.WAYPOINT_EVENT = NewRTSLayerEvent(NullPathFindingLayer_NULL_PATH_FINDING_LAYER)
}

func NewAssignWaypointsUtil() *AssignWaypointsUtil {
    this := &AssignWaypointsUtil{}
    this.initFields()
    return this
}


func AssignWaypointsUtil_GetInstance() *AssignWaypointsUtil {
        return AssignWaypointsUtil_Instance
}


func (this *AssignWaypointsUtil) Set(unitLayer *UnitLayer, ownerAdvancedRTSGameLayer IAdvancedRTSGameLayer) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "set", e)
                }
            }()
            var list IBasicArrayList = ownerAdvancedRTSGameLayer.GetWaypointBehavior().GetOwnedWaypointList(); _ = list
            var unitWaypointBehavior IUnitWaypointBehavior = any(unitLayer.GetWaypointBehavior()).(IUnitWaypointBehavior); _ = unitWaypointBehavior
            for index := int(list.Size() - 1); index >= 0; index-- {
                var waypointLayer IRTSLayer = any(list.Get(index)).(IRTSLayer); _ = waypointLayer
                this.WAYPOINT_EVENT.SetRtsLayer(waypointLayer)
                unitWaypointBehavior.OnWaypointEvent(this.WAYPOINT_EVENT)
            }
            var waypointLayer PathFindingLayerInterface = PrimaryWaypointHelper_GetInstance().GetWaypointLayer(); _ = waypointLayer
            if waypointLayer != NullPathFindingLayer_NULL_PATH_FINDING_LAYER {
                this.WAYPOINT_EVENT.SetRtsLayer(waypointLayer)
                unitWaypointBehavior.OnWaypointEvent(this.WAYPOINT_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

