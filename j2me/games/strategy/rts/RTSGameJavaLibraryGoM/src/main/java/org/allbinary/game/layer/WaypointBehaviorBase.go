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

type WaypointBehaviorBase struct {
    OwnedWaypointList IBasicArrayList
    WaypointOverridesAttacking bool
    Waypoint IWaypointBase
    SteeringVisitorList IBasicArrayList
}

// Implements TickableInterface

// The Java class WaypointBehaviorBase has subclasses so its variables are this interface
type IWaypointBehaviorBase interface {
    WaypointBehaviorBase_() *WaypointBehaviorBase
    GetName() string
    IsRunning() bool
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    GetOwnedWaypointList() IBasicArrayList
    GetWaypoint() IWaypointBase
    SetWaypoint(waypoint IWaypointBase)
    GetSteeringVisitorList() IBasicArrayList
    GetMovementLogicAsString() string
    IsWaypointListEmptyOrOnlyTargets() bool
    IsInSensorRange(layerInterface ICollidableDestroyableDamageableLayer, targetDistance int) bool
    GetCurrentTargetLayerInterface() ICollidableDestroyableDamageableLayer
    GetCurrentTargetingStateString() string
    IsTrackingWaypoint() bool
    GetCurrentPathGeographicMapCellPosition() IGeographicMapCellPosition
    GetCurrentGeographicMapCellHistory() *GeographicMapCellHistory
    IsWaypointOverridesAttacking() bool
    GetNextUnvisitedPathGeographicMapCellPosition() IGeographicMapCellPosition
    NeedToMove() bool
    SetWaypointPathsList(waypointPathsList IBasicArrayList)
    GetWaypointPathsList() IBasicArrayList
    SetTarget(layerInterface PathFindingLayerInterface)
    SetTargetWithDistance(layerInterface PathFindingLayerInterface, anotherTargetDistance int)
    ClearTarget()
    UpdatePathOnTargetMove(reason string)
}

func (this *WaypointBehaviorBase) WaypointBehaviorBase_() *WaypointBehaviorBase {
    return this
}


func (this *WaypointBehaviorBase) initFields() {
    
    this.WaypointOverridesAttacking = true
    this.Waypoint = WaypointBase_NULL_WAYPOINT_BASE
    this.SteeringVisitorList = NewBasicArrayListD()
}


func NewWaypointBehaviorBase() *WaypointBehaviorBase {
    this := &WaypointBehaviorBase{}
    this.initFields()
        this.OwnedWaypointList = NewBasicArrayListS(1)
    return this
}


func (this *WaypointBehaviorBase) GetName() string {
        return reflect.TypeOf(this).String()
}


func (this *WaypointBehaviorBase) IsRunning() bool {
        return false
}


// throws Exception
func (this *WaypointBehaviorBase) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
}


func (this *WaypointBehaviorBase) GetOwnedWaypointList() IBasicArrayList {
        return this.OwnedWaypointList
}


func (this *WaypointBehaviorBase) GetWaypoint() IWaypointBase {
        return this.Waypoint
}


func (this *WaypointBehaviorBase) SetWaypoint(waypoint IWaypointBase) {
        this.Waypoint = waypoint
}


func (this *WaypointBehaviorBase) GetSteeringVisitorList() IBasicArrayList {
        return this.SteeringVisitorList
}


func (this *WaypointBehaviorBase) GetMovementLogicAsString() string {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) IsWaypointListEmptyOrOnlyTargets() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) IsInSensorRange(layerInterface ICollidableDestroyableDamageableLayer, targetDistance int) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) GetCurrentTargetLayerInterface() ICollidableDestroyableDamageableLayer {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) GetCurrentTargetingStateString() string {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) IsTrackingWaypoint() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) GetCurrentPathGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) GetCurrentGeographicMapCellHistory() *GeographicMapCellHistory {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) IsWaypointOverridesAttacking() bool {
        return this.WaypointOverridesAttacking
}


func (this *WaypointBehaviorBase) GetNextUnvisitedPathGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) NeedToMove() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) SetWaypointPathsList(waypointPathsList IBasicArrayList) {
        panic(errors.New(fmt.Sprint()))
}


func (this *WaypointBehaviorBase) GetWaypointPathsList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBehaviorBase) SetTarget(layerInterface PathFindingLayerInterface) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBehaviorBase) SetTargetWithDistance(layerInterface PathFindingLayerInterface, anotherTargetDistance int) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBehaviorBase) ClearTarget() {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBehaviorBase) UpdatePathOnTargetMove(reason string) {
}

