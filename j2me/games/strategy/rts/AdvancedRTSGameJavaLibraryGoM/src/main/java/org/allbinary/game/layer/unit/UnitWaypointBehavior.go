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

type UnitWaypointBehavior struct {
    WaypointBehaviorBase
    CommonSeps *CommonSeps
    LongWeaponRange int
    SensorAction *SensorAction
    CompleteTimeDelayHelper ITimeDelayHelper
    CurrentGeographicMapCellHistoryP *GeographicMapCellHistory
    LastPathGeographicMapCellPosition IGeographicMapCellPosition
    CurrentPathGeographicMapCellPosition IGeographicMapCellPosition
    FAKE_WAYPOINT_LAYER ICollidableDestroyableDamageableLayer
    TargetList IBasicArrayList
    PossibleTargetList IBasicArrayList
    AssociatedAdvancedRTSGameLayer *UnitLayer
    Moving bool
    MovingFromStopped bool
    WaypointPathsListP IBasicArrayList
    CurrentTargetDistance int
    CurrentTargetLayerInterfaceP ICollidableDestroyableDamageableLayer
    TrackingWaypoint bool
    RepeatedToLong ITimeDelayHelper
}

// Implements WaypointEventListenerInterface

// The Java class UnitWaypointBehavior has subclasses so its variables are this interface
type IUnitWaypointBehavior interface {
    IWaypointBehaviorBase
    UnitWaypointBehavior_() *UnitWaypointBehavior
    InitRange(weaponRange int)
    OnEvent(eventObject IAllBinaryEventObject)
    OnWaypointEvent(event *RTSLayerEvent)
    AddWaypointFromUser(advancedRTSGameLayer IAdvancedRTSGameLayer)
    AddWaypointFromBuilding(advancedRTSGameLayer IAdvancedRTSGameLayer)
    IsWaypointListEmptyOrOnlyTargets() bool
    InsertWaypoint(index int, rtsLayer ICollidableDestroyableDamageableLayer) bool
    Move()
    SetRandomGeographicMapCellHistory(pathsList IBasicArrayList)
    SetGeographicMapCellHistoryPath(geographicMapCellPositionBasicArrayList IBasicArrayList)
    CanInsertWaypoint(index int, rtsLayer ICollidableDestroyableDamageableLayer) bool
    AddBuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition)
    MoveAwayFromBuilding(buildingLayer *BuildingLayer)
    NeedToMove() bool
    GetMovementLogicAsString() string
    IsMovingFromStopped() bool
    SetMovingFromStopped(movingFromStopped bool)
    SetWaypointPathsList(waypointPathsList IBasicArrayList)
    GetWaypointPathsList() IBasicArrayList
    IsMoving() bool
    SetMoving(moving bool)
    IsTrackingWaypoint() bool
    GetPossibleTargetList() IBasicArrayList
    SetLastPathGeographicMapCellPosition(lastPathGeographicMapCellPosition IGeographicMapCellPosition)
    GetLastPathGeographicMapCellPosition() IGeographicMapCellPosition
    SetCurrentPathGeographicMapCellPosition(currentPathGeographicMapCellPosition IGeographicMapCellPosition)
    GetCurrentPathGeographicMapCellPosition() IGeographicMapCellPosition
    SetSensorAction(sensorAction *SensorAction)
    GetSensorAction() *SensorAction
    SetCurrentTargetLayerInterface(currentTargetLayerInterface ICollidableDestroyableDamageableLayer)
    GetCurrentTargetLayerInterface() ICollidableDestroyableDamageableLayer
    SetCurrentTargetDistance(currentTargetDistance int)
    GetCurrentTargetDistance() int
    GetCompleteTimeDelayHelper() ITimeDelayHelper
    GetCurrentGeographicMapCellHistory() *GeographicMapCellHistory
    SetTrackingWaypoint(trackingWaypoint bool)
}

func (this *UnitWaypointBehavior) UnitWaypointBehavior_() *UnitWaypointBehavior {
    return this
}

var UnitWaypointBehavior_PATHING string = "Pathing"

func (this *UnitWaypointBehavior) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    this.LongWeaponRange = 0
    this.SensorAction = SensorActionFactory_GetInstance().ATTACK
    
    
    this.LastPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.CurrentPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    
    
    
    
    this.Moving = false
    this.MovingFromStopped = false
    this.WaypointPathsListP = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    this.CurrentTargetDistance = Integer_MAX_VALUE
    this.CurrentTargetLayerInterfaceP = CollidableDestroyableDamageableLayer_GetNullInstance()
    
    this.RepeatedToLong = NewTimeDelayHelper(22000)
}


func NewUnitWaypointBehavior(associatedAdvancedRTSGameLayer *UnitLayer, fakeWaypoint ICollidableDestroyableDamageableLayer) *UnitWaypointBehavior {
    this := &UnitWaypointBehavior{}
    this.WaypointBehaviorBase = *NewWaypointBehaviorBase()
    this.initFields()
        this.AssociatedAdvancedRTSGameLayer = associatedAdvancedRTSGameLayer
        this.CompleteTimeDelayHelper = NewTimeDelayHelper(30000)
        this.TargetList = NewBasicArrayListD()
        this.PossibleTargetList = NewBasicArrayListD()
        this.SetWaypointPathsList(BasicArrayListUtil_GetInstance().GetImmutableInstance())
        this.CurrentGeographicMapCellHistoryP = NewGeographicMapCellHistory()
        this.FAKE_WAYPOINT_LAYER = fakeWaypoint
    return this
}


func (this *UnitWaypointBehavior) InitRange(weaponRange int) {
        this.LongWeaponRange = weaponRange / 2
}


func (this *UnitWaypointBehavior) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *UnitWaypointBehavior) OnWaypointEvent(event *RTSLayerEvent) {
        var advancedRTSGameLayer IAdvancedRTSGameLayer = any(event.GetRtsLayer()).(IAdvancedRTSGameLayer); _ = advancedRTSGameLayer
        this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.OnWaypointEvent(this.AssociatedAdvancedRTSGameLayer, advancedRTSGameLayer)
        if this.AssociatedAdvancedRTSGameLayer.IsSelected() {
            this.AddWaypointFromUser(advancedRTSGameLayer)
        } else if advancedRTSGameLayer.ShouldAddWaypointFromBuilding() {
            this.AddWaypointFromBuilding(advancedRTSGameLayer)
        }
}


// throws Exception
func (this *UnitWaypointBehavior) AddWaypointFromUser(advancedRTSGameLayer IAdvancedRTSGameLayer) {
}


// throws Exception
func (this *UnitWaypointBehavior) AddWaypointFromBuilding(advancedRTSGameLayer IAdvancedRTSGameLayer) {
        if advancedRTSGameLayer == PrimaryWaypointHelper_GetInstance().GetWaypointLayer() || advancedRTSGameLayer.GetParentLayer() == this.AssociatedAdvancedRTSGameLayer.GetParentLayer() {
            if !this.TargetList.Contains(advancedRTSGameLayer) {
                this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.AddWaypointFromBuilding(this.AssociatedAdvancedRTSGameLayer, advancedRTSGameLayer)
                if advancedRTSGameLayer.IsDestroyed() {
                    panic(errors.New(fmt.Sprint("Trying to add a dead: " + fmt.Sprint(advancedRTSGameLayer))))
                }
                this.TargetList.Add(advancedRTSGameLayer)
                this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.AddWaypointFromBuildingList(this.AssociatedAdvancedRTSGameLayer, advancedRTSGameLayer, this.TargetList)
            }
        }
}


func (this *UnitWaypointBehavior) IsWaypointListEmptyOrOnlyTargets() bool {
        return false
}


// throws Exception
func (this *UnitWaypointBehavior) InsertWaypoint(index int, rtsLayer ICollidableDestroyableDamageableLayer) bool {
        if this.CanInsertWaypoint(index, rtsLayer) {
            this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.InsertWaypoint(this.AssociatedAdvancedRTSGameLayer, index, rtsLayer, this.GetName())
            if rtsLayer.IsDestroyed() {
                panic(errors.New(fmt.Sprint("Trying to add a dead: " + fmt.Sprint(rtsLayer))))
            }
            this.TargetList.AddAt(index, rtsLayer)
            this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.InsertWaypointList(this.AssociatedAdvancedRTSGameLayer, index, rtsLayer, this.GetName(), this.TargetList)
            return true
        }
        return false
}


func (this *UnitWaypointBehavior) Move() {
        if this.IsMoving() {
            this.SetMovingFromStopped(false)
        } else {
            this.SetMovingFromStopped(true)
        }
        this.SetMoving(false)
}


// throws Exception
func (this *UnitWaypointBehavior) SetRandomGeographicMapCellHistory(pathsList IBasicArrayList) {
        var size int = pathsList.Size(); _ = size
        this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.SetRandomGeographicMapCellHistoryList(this.AssociatedAdvancedRTSGameLayer, pathsList)
        if size > 0 {
            var geographicMapCellPositionBasicArrayList IBasicArrayList = any(BasicArrayListUtil_GetInstance().GetRandom(pathsList)).(IBasicArrayList); _ = geographicMapCellPositionBasicArrayList
            this.SetGeographicMapCellHistoryPath(geographicMapCellPositionBasicArrayList)
        }
}


// throws Exception
func (this *UnitWaypointBehavior) SetGeographicMapCellHistoryPath(geographicMapCellPositionBasicArrayList IBasicArrayList) {
        this.LastPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
        if this.AssociatedAdvancedRTSGameLayer.ShowMoreCaptionStates {
            this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(UnitWaypointBehavior_PATHING, BasicColorFactory_GetInstance().GREEN)
        }
        this.CurrentGeographicMapCellHistoryP.Init()
        this.AssociatedAdvancedRTSGameLayer.Init(this.CurrentGeographicMapCellHistoryP, geographicMapCellPositionBasicArrayList)
        this.SetTrackingWaypoint(true)
        this.GetCompleteTimeDelayHelper().SetStartTimeTNT()
}


func (this *UnitWaypointBehavior) CanInsertWaypoint(index int, rtsLayer ICollidableDestroyableDamageableLayer) bool {
        if this.TargetList.Size() > 4 {
        } else if this.TargetList.Contains(rtsLayer) {
        } else {
            return true
        }
        return false
}


// throws Exception
func (this *UnitWaypointBehavior) AddBuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) {
}


// throws Exception
func (this *UnitWaypointBehavior) MoveAwayFromBuilding(buildingLayer *BuildingLayer) {
        var geographicMapCellPosition IGeographicMapCellPosition = this.AssociatedAdvancedRTSGameLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
        var list IBasicArrayList = buildingLayer.GeographicMapCellPositionAreaBase.GetOccupyingGeographicMapCellPositionList(); _ = list
        if list.Contains(geographicMapCellPosition) {
            if this.InsertWaypoint(0, this.FAKE_WAYPOINT_LAYER) {
                this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.MoveAwayFromBuilding(this.AssociatedAdvancedRTSGameLayer)
                this.SetCurrentTargetLayerInterface(any(this.FAKE_WAYPOINT_LAYER).(ICollidableDestroyableDamageableLayer))
                var pathsList IBasicArrayList = buildingLayer.GetMoveOutOfBuildAreaPath(geographicMapCellPosition); _ = pathsList
                this.AssociatedAdvancedRTSGameLayer.SetClosestGeographicMapCellHistory(pathsList)
            }
        }
}


func (this *UnitWaypointBehavior) NeedToMove() bool {
        this.AssociatedAdvancedRTSGameLayer.WaypointLogHelperP.NeedToMove(this.AssociatedAdvancedRTSGameLayer, this)
        if this.IsTrackingWaypoint() || this.SensorAction == SensorActionFactory_GetInstance().EVADE || (this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() && this.GetCurrentTargetDistance() >= this.LongWeaponRange + this.CurrentTargetLayerInterfaceP.GetHalfHeight()) {
            this.RepeatedToLong.SetStartTimeTNT()
            return true
        }
        if this.RepeatedToLong.IsTimeTNT() {
            var message string = "Repeating too long: " + fmt.Sprint(this.GetMovementLogicAsString()); _ = message
            ForcedLogUtil_Log(message, this.AssociatedAdvancedRTSGameLayer)
        }
        return false
}


func (this *UnitWaypointBehavior) GetMovementLogicAsString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("isTrackingWaypoint: ")
        stringBuffer.Appendboolean(this.IsTrackingWaypoint())
        stringBuffer.Append(" sensorAction: ")
        stringBuffer.Append(this.SensorAction.Name)
        stringBuffer.Append(" getCurrentTargetLayerInterface: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.CurrentTargetLayerInterfaceP))
        if this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() {
            stringBuffer.Append(" Target Range: ")
            stringBuffer.Appendint(this.GetCurrentTargetDistance())
            stringBuffer.Append(" >= ")
            stringBuffer.Appendint(this.LongWeaponRange + this.CurrentTargetLayerInterfaceP.GetHalfHeight())
        }
        return fmt.Sprint(stringBuffer)
}


func (this *UnitWaypointBehavior) IsMovingFromStopped() bool {
        return this.MovingFromStopped
}


func (this *UnitWaypointBehavior) SetMovingFromStopped(movingFromStopped bool) {
        this.MovingFromStopped = movingFromStopped
}


func (this *UnitWaypointBehavior) SetWaypointPathsList(waypointPathsList IBasicArrayList) {
        this.WaypointPathsListP = waypointPathsList
}


func (this *UnitWaypointBehavior) GetWaypointPathsList() IBasicArrayList {
        return this.WaypointPathsListP
}


func (this *UnitWaypointBehavior) IsMoving() bool {
        return this.Moving
}


func (this *UnitWaypointBehavior) SetMoving(moving bool) {
        this.Moving = moving
}


func (this *UnitWaypointBehavior) IsTrackingWaypoint() bool {
        return this.TrackingWaypoint
}


func (this *UnitWaypointBehavior) GetPossibleTargetList() IBasicArrayList {
        return this.PossibleTargetList
}


func (this *UnitWaypointBehavior) SetLastPathGeographicMapCellPosition(lastPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.LastPathGeographicMapCellPosition = lastPathGeographicMapCellPosition
}


func (this *UnitWaypointBehavior) GetLastPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.LastPathGeographicMapCellPosition
}


func (this *UnitWaypointBehavior) SetCurrentPathGeographicMapCellPosition(currentPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.CurrentPathGeographicMapCellPosition = currentPathGeographicMapCellPosition
}


func (this *UnitWaypointBehavior) GetCurrentPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.CurrentPathGeographicMapCellPosition
}


func (this *UnitWaypointBehavior) SetSensorAction(sensorAction *SensorAction) {
        this.SensorAction = sensorAction
}


func (this *UnitWaypointBehavior) GetSensorAction() *SensorAction {
        return this.SensorAction
}


func (this *UnitWaypointBehavior) SetCurrentTargetLayerInterface(currentTargetLayerInterface ICollidableDestroyableDamageableLayer) {
        this.CurrentTargetLayerInterfaceP = currentTargetLayerInterface
}


func (this *UnitWaypointBehavior) GetCurrentTargetLayerInterface() ICollidableDestroyableDamageableLayer {
        return this.CurrentTargetLayerInterfaceP
}


func (this *UnitWaypointBehavior) SetCurrentTargetDistance(currentTargetDistance int) {
        this.CurrentTargetDistance = currentTargetDistance
}


func (this *UnitWaypointBehavior) GetCurrentTargetDistance() int {
        return this.CurrentTargetDistance
}


func (this *UnitWaypointBehavior) GetCompleteTimeDelayHelper() ITimeDelayHelper {
        return this.CompleteTimeDelayHelper
}


func (this *UnitWaypointBehavior) GetCurrentGeographicMapCellHistory() *GeographicMapCellHistory {
        return this.CurrentGeographicMapCellHistoryP
}


func (this *UnitWaypointBehavior) SetTrackingWaypoint(trackingWaypoint bool) {
        this.TrackingWaypoint = trackingWaypoint
}

