/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type Waypoint2SelectedLogHelper struct {
    Waypoint2LogHelper
    LogUtil *LogUtil
    CommonSeps *CommonSeps
    PositionStrings *PositionStrings
    StringUtil *StringUtil
    ENEMY_POSSIBLE_TARGET string
    CLOSE_RANGE string
    SENSOR_RANGE string
    ONLY_TARGETS string
    SHORTER_THAN_CURRENT_TARGET_DISTANCE string
    CURRENT_TARGET_DESTROYED string
    CURRENT_TARGET_NOT_AVAILABLE string
    PROCESS_POSSIBLE_TARGET string
}

var Waypoint2SelectedLogHelper_InstanceC *Waypoint2SelectedLogHelper = NewWaypoint2SelectedLogHelper()

func (this *Waypoint2SelectedLogHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.PositionStrings = PositionStrings_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.ENEMY_POSSIBLE_TARGET = " Enemy Possible Target: "
    this.CLOSE_RANGE = "closeRange: "
    this.SENSOR_RANGE = " sensorRange: "
    this.ONLY_TARGETS = " Only Targets: "
    this.SHORTER_THAN_CURRENT_TARGET_DISTANCE = " isShorterThanCurrentTargetDistance: "
    this.CURRENT_TARGET_DESTROYED = " isCurrentTargetDestroyed: "
    this.CURRENT_TARGET_NOT_AVAILABLE = " isCurrentTargetNotAvailable: "
    this.PROCESS_POSSIBLE_TARGET = "selected: processPossibleTarget1"
}

func NewWaypoint2SelectedLogHelper() *Waypoint2SelectedLogHelper {
    this := &Waypoint2SelectedLogHelper{}
    this.Waypoint2LogHelper = *NewWaypoint2LogHelper()
    this.initFields()
    return this
}


func Waypoint2SelectedLogHelper_GetInstance() IWaypoint2LogHelper {
        return Waypoint2SelectedLogHelper_InstanceC
}


func (this *Waypoint2SelectedLogHelper) InitRange(associatedAdvancedRTSGameLayer PathFindingLayerInterface, closeRange int, sensorRange int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CLOSE_RANGE)
        stringBuffer.Appendint(closeRange)
        stringBuffer.Append(this.SENSOR_RANGE)
        stringBuffer.Appendint(sensorRange)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: initRange")
}


func (this *Waypoint2SelectedLogHelper) ProcessPossibleTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int, isShorterThanCurrentTargetDistance bool, isCurrentTargetDestroyed bool) {
        if waypointBehaviorBase.IsWaypointListEmptyOrOnlyTargets() && waypointBehaviorBase.IsInSensorRange(any(layerInterface).(ICollidableDestroyableDamageableLayer), anotherTargetDistance) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
            stringBuffer.Append(this.ONLY_TARGETS)
            stringBuffer.Appendboolean(waypointBehaviorBase.IsWaypointListEmptyOrOnlyTargets())
            stringBuffer.Append(this.SHORTER_THAN_CURRENT_TARGET_DISTANCE)
            stringBuffer.Appendboolean(isShorterThanCurrentTargetDistance)
            stringBuffer.Append(this.CURRENT_TARGET_DESTROYED)
            stringBuffer.Appendboolean(isCurrentTargetDestroyed)
            stringBuffer.Append(this.CURRENT_TARGET_NOT_AVAILABLE)
            stringBuffer.Appendboolean((waypointBehaviorBase.GetCurrentTargetLayerInterface() == nil))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.PROCESS_POSSIBLE_TARGET)
        }
}


func (this *Waypoint2SelectedLogHelper) ProcessSetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(waypointBehaviorBase.GetCurrentTargetingStateString())
        stringBuffer.Append(this.ENEMY_POSSIBLE_TARGET)
        stringBuffer.Append(layerInterface.GetName())
        stringBuffer.Append(" anotherTargetDistance: ")
        stringBuffer.Appendint(anotherTargetDistance)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processPossibleTarget2")
}


func (this *Waypoint2SelectedLogHelper) ProcessPossibleTargetCloser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(waypointBehaviorBase.GetCurrentTargetingStateString())
        stringBuffer.Append(this.ENEMY_POSSIBLE_TARGET)
        stringBuffer.Append(layerInterface.GetName())
        stringBuffer.Append(" anotherTargetDistance: ")
        stringBuffer.Appendint(anotherTargetDistance)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processPossibleTarget3")
}


func (this *Waypoint2SelectedLogHelper) SetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(waypointBehaviorBase.GetName())
        stringBuffer.Append(" targeting: ")
        stringBuffer.Append(layerInterface.GetName())
        stringBuffer.Append(" at: ")
        stringBuffer.Appendint(anotherTargetDistance)
        stringBuffer.Append(" instead of: ")
        stringBuffer.Append(waypointBehaviorBase.GetCurrentTargetingStateString())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: setTarget")
}


func (this *Waypoint2SelectedLogHelper) SetTargetPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF("Target Destroyed", this, "selected: setTargetPath")
}


func (this *Waypoint2SelectedLogHelper) SetTargetPathIgnoreNewPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Set Path To Target: ")
        stringBuffer.Append(waypointBehaviorBase.GetCurrentTargetLayerInterface().GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.StringUtil.ToString(waypointBehaviorBase.GetCurrentTargetLayerInterface()))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: setTargetPath")
}


func (this *Waypoint2SelectedLogHelper) ProcessWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase, targetLayer PathFindingLayerInterface, size int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Size: ")
        stringBuffer.Appendint(size)
        stringBuffer.Append(" Waypoint: ")
        stringBuffer.Append(targetLayer.GetName())
        stringBuffer.Append(" Target: ")
        stringBuffer.Append(this.StringUtil.ToString(waypointBehaviorBase.GetCurrentTargetLayerInterface()))
        stringBuffer.Append(" isTrackingWaypoint: ")
        stringBuffer.Appendboolean(waypointBehaviorBase.IsTrackingWaypoint())
        stringBuffer.Append(" WaypointOverridesAttacking: ")
        stringBuffer.Appendboolean(waypointBehaviorBase.IsWaypointOverridesAttacking())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processWaypoint")
}


func (this *Waypoint2SelectedLogHelper) ProcessWaypointTracked(associatedAdvancedRTSGameLayer PathFindingLayerInterface, waypointBehaviorBase IWaypointBehaviorBase) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" currentPathGeographicMapCellPosition: ")
        stringBuffer.Append(this.StringUtil.ToString(waypointBehaviorBase.GetCurrentPathGeographicMapCellPosition()))
        stringBuffer.Append(" nextUnvisitedPathGeographicMapCellPosition: ")
        stringBuffer.Append(this.StringUtil.ToString(waypointBehaviorBase.GetNextUnvisitedPathGeographicMapCellPosition()))
        stringBuffer.Append(" Tracked: ")
        stringBuffer.Append(this.StringUtil.ToString(waypointBehaviorBase.GetCurrentGeographicMapCellHistory().GetTracked()))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processWaypoint")
}


func (this *Waypoint2SelectedLogHelper) ProcessWaypointTrackedVisit(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellPosition IGeographicMapCellPosition) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Visited: ")
        stringBuffer.Append(this.StringUtil.ToString(geographicMapCellPosition))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processWaypoint")
}


func (this *Waypoint2SelectedLogHelper) ProcessWaypointTrackedWithoutProgress(associatedAdvancedRTSGameLayer PathFindingLayerInterface, reason string) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(reason)), this, "turnTo")
}


func (this *Waypoint2SelectedLogHelper) Wander(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(associatedAdvancedRTSGameLayer.GetName(), this, "wander")
}


func (this *Waypoint2SelectedLogHelper) TargetDestroyed(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" Target Destroyed")), this, "selected: processTargeting")
}


func (this *Waypoint2SelectedLogHelper) ProcessTargeting(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.PositionStrings.DX_LABEL)
        stringBuffer.Appendint(dx)
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.PositionStrings.DY_LABEL)
        stringBuffer.Appendint(dy)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processTargeting")
}


func (this *Waypoint2SelectedLogHelper) ProcessTargetingNonWayPoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Targeting Non Waypoint ")
        stringBuffer.Append(this.PositionStrings.DX_LABEL)
        stringBuffer.Appendint(dx)
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.PositionStrings.DY_LABEL)
        stringBuffer.Appendint(dy)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: processTargeting")
}


func (this *Waypoint2SelectedLogHelper) RemoveWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, waypointLayer PathFindingLayerInterface, reason string) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(unitWaypointBehavior2.GetName())
        stringBuffer.Append(" remove: ")
        stringBuffer.Append(waypointLayer.GetName())
        stringBuffer.Append(" because: ")
        stringBuffer.Append(reason)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: removeWaypoint")
}


func (this *Waypoint2SelectedLogHelper) RemoveWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, list IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(unitWaypointBehavior2.GetName())
        stringBuffer.Append(" Waypoints: ")
        stringBuffer.Append(this.StringUtil.ToString(list))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: removeWaypoint")
}


func (this *Waypoint2SelectedLogHelper) RemoveWaypointClear(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" Clearing removed waypoint")), this, "selected: removeWaypoint")
}


func (this *Waypoint2SelectedLogHelper) ClearTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" Cleared Target")), this, "selected: clearTarget")
}


func (this *Waypoint2SelectedLogHelper) AddWaypointFromUser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Adding Waypoint to Selected: ")
        stringBuffer.Append(advancedRTSGameLayer.GetName())
        if advancedRTSGameLayer.GetParentLayer() != nil {
            stringBuffer.Append(" With Parent: ")
            stringBuffer.Append(advancedRTSGameLayer.GetParentLayer().GetName())
        }
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: onWaypointEvent")
}


func (this *Waypoint2SelectedLogHelper) TargetMovedSoRetarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" - target moved so retargeting")), this, "turnTo")
}

