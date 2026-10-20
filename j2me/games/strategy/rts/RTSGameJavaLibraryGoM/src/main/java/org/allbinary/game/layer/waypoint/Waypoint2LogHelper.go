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

type Waypoint2LogHelper struct {
}

// The Java class Waypoint2LogHelper has subclasses so its variables are this interface
type IWaypoint2LogHelper interface {
    Waypoint2LogHelper_() *Waypoint2LogHelper
    InitRange(associatedAdvancedRTSGameLayer PathFindingLayerInterface, closeRange int, sensorRange int)
    ProcessPossibleTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int, isShorterThanCurrentTargetDistance bool, isCurrentTargetDestroyed bool)
    ProcessSetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int)
    ProcessPossibleTargetCloser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int)
    SetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int)
    SetTargetPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    SetTargetPathIgnoreNewPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase)
    ProcessWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, targetLayer PathFindingLayerInterface, size int)
    ProcessWaypointTracked(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase)
    ProcessWaypointTrackedVisit(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellPosition IGeographicMapCellPosition)
    ProcessWaypointTrackedWithoutProgress(associatedAdvancedRTSGameLayer PathFindingLayerInterface, reason string)
    Wander(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    TargetDestroyed(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    ProcessTargeting(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int)
    ProcessTargetingNonWayPoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int)
    RemoveWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, waypointLayer PathFindingLayerInterface, reason string)
    RemoveWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, list IBasicArrayList)
    RemoveWaypointClear(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    ClearTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    AddWaypointFromUser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface)
    TargetMovedSoRetarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
}

func (this *Waypoint2LogHelper) Waypoint2LogHelper_() *Waypoint2LogHelper {
    return this
}

var Waypoint2LogHelper_Instance IWaypoint2LogHelper = NewWaypoint2LogHelper()

func NewWaypoint2LogHelper() *Waypoint2LogHelper {
    this := &Waypoint2LogHelper{}
    return this
}


func Waypoint2LogHelper_GetInstance() IWaypoint2LogHelper {
        return Waypoint2LogHelper_Instance
}


func (this *Waypoint2LogHelper) InitRange(associatedAdvancedRTSGameLayer PathFindingLayerInterface, closeRange int, sensorRange int) {
}


func (this *Waypoint2LogHelper) ProcessPossibleTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int, isShorterThanCurrentTargetDistance bool, isCurrentTargetDestroyed bool) {
}


func (this *Waypoint2LogHelper) ProcessSetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
}


func (this *Waypoint2LogHelper) ProcessPossibleTargetCloser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
}


func (this *Waypoint2LogHelper) SetTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
}


func (this *Waypoint2LogHelper) SetTargetPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) SetTargetPathIgnoreNewPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase) {
}


func (this *Waypoint2LogHelper) ProcessWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, targetLayer PathFindingLayerInterface, size int) {
}


func (this *Waypoint2LogHelper) ProcessWaypointTracked(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase) {
}


func (this *Waypoint2LogHelper) ProcessWaypointTrackedVisit(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellPosition IGeographicMapCellPosition) {
}


func (this *Waypoint2LogHelper) ProcessWaypointTrackedWithoutProgress(associatedAdvancedRTSGameLayer PathFindingLayerInterface, reason string) {
}


func (this *Waypoint2LogHelper) Wander(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) TargetDestroyed(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) ProcessTargeting(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int) {
}


func (this *Waypoint2LogHelper) ProcessTargetingNonWayPoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int) {
}


func (this *Waypoint2LogHelper) RemoveWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, waypointLayer PathFindingLayerInterface, reason string) {
}


func (this *Waypoint2LogHelper) RemoveWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior2 IWaypointBehaviorBase, list IBasicArrayList) {
}


func (this *Waypoint2LogHelper) RemoveWaypointClear(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) ClearTarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) AddWaypointFromUser(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *Waypoint2LogHelper) TargetMovedSoRetarget(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}

