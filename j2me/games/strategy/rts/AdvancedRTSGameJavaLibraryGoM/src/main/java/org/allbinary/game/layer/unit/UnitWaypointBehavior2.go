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

type UnitWaypointBehavior2 struct {
    UnitWaypointBehavior
    LogUtil *LogUtil
    UnitWaypointStrings *UnitWaypointStrings
    BasicColorFactory *BasicColorFactory
    LayerDistanceUtil *LayerDistanceUtil
    PathFindingThreadPool IThreadPool
    ProgressTimeDelayHelper ITimeDelayHelper
    WanderPathsList IBasicArrayList
    WaypointPathRunnable IWaypointPathRunnableBase
    SensorRange int
    CloseRange int
    NextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition
    AfterNextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition
    WaitingOnTargetPath bool
    WaitingOnWaypointPath bool
    TargetWithoutCachedPathLayerInterface ICollidableDestroyableDamageableLayer
    BuildingSteeringVisitor *BuildingSteeringVisitor
}

var UnitWaypointBehavior2_RunningWaypointPathList IBasicArrayList = NewBasicArrayListD()
var UnitWaypointBehavior2_TARGET_DISTANCE string = "Target Distance"
var UnitWaypointBehavior2_TARGET_LAYER string = "Target Layer"

func (this *UnitWaypointBehavior2) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.UnitWaypointStrings = UnitWaypointStrings_GetInstance()
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.LayerDistanceUtil = LayerDistanceUtil_GetInstance()
    this.PathFindingThreadPool = PathFindingThreadPool_GetInstance()
    
    
    
    this.SensorRange = 0
    this.CloseRange = 0
    this.NextUnvisitedPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.AfterNextUnvisitedPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    
    
    this.TargetWithoutCachedPathLayerInterface = CollidableDestroyableDamageableLayer_GetNullInstance()
    this.BuildingSteeringVisitor = NewBuildingSteeringVisitor()
}


func NewUnitWaypointBehavior2(ownerAdvancedRTSGameLayer *UnitLayer, fakeWaypoint IAdvancedRTSGameLayer) *UnitWaypointBehavior2 {
    this := &UnitWaypointBehavior2{}
        this.UnitWaypointBehavior = *NewUnitWaypointBehavior(ownerAdvancedRTSGameLayer, fakeWaypoint)
        this.initFields()
        this.ProgressTimeDelayHelper = NewTimeDelayHelper(5000)
        this.WanderPathsList = NewBasicArrayListD()
        var features *Features = Features_GetInstance(); _ = features
        this.WaypointPathRunnable = func() any { if J2MEUtil_IsHTML() { return NewMultipassWaypointPathRunnable() }; return NewWaypointPathRunnable() }()
    return this
}


func (this *UnitWaypointBehavior2) InitRange(weaponRange int) {
        this.UnitWaypointBehavior.InitRange(weaponRange)
        this.CloseRange = weaponRange
        this.SensorRange = weaponRange * 4
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.InitRange(this.AssociatedAdvancedRTSGameLayer, this.CloseRange, this.SensorRange)
}


func (this *UnitWaypointBehavior2) GetNextUnvisitedPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.NextUnvisitedPathGeographicMapCellPosition
}


func (this *UnitWaypointBehavior2) IsRunning() bool {
        if this.WaypointPathRunnable.IsRunning() {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.WaypointPathRunnable.IsRunning() {
            if this.WaypointPathsListP != UnitWaypointBehavior2_RunningWaypointPathList {
                this.WaypointPathRunnable.SetRunning(false)
                if this.WaitingOnTargetPath {
                    this.SetTargetPath()
                } else if this.WaitingOnWaypointPath {
                    this.SetWaypointPath(any(this.WaypointPathRunnable.GetTargetLayer()).(IAdvancedRTSGameLayer))
                } else {
                    panic(errors.New(fmt.Sprint("Should not happen")))
                }
            } else {
                return
            }
        }
        this.ProcessTargetList()
        if !this.WaypointPathRunnable.IsRunning() {
            this.ProcessWaypoint()
        } else {
            return
        }
        if !this.WaypointPathRunnable.IsRunning() {
            this.ProcessTargeting()
        } else {
            return
        }
        if !this.WaypointPathRunnable.IsRunning() {
            this.TeleportIfNoProgress()
        } else {
            return
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) OnEnemyMovement(layerInterface IAdvancedRTSGameLayer) {
        var anotherTargetDistance int = this.LayerDistanceUtil.GetDistance(this.AssociatedAdvancedRTSGameLayer, layerInterface); _ = anotherTargetDistance
        if layerInterface == this.CurrentTargetLayerInterfaceP {
            this.SetCurrentTargetDistance(anotherTargetDistance)
        } else {
            this.ProcessPossibleTarget(layerInterface, anotherTargetDistance)
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) ProcessPossibleTarget(layerInterface IAdvancedRTSGameLayer, anotherTargetDistance int) {
        var isShorterThanCurrentTargetDistance bool = this.GetCurrentTargetDistance() > anotherTargetDistance; _ = isShorterThanCurrentTargetDistance
        var isCurrentTargetDestroyed bool = this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() && this.CurrentTargetLayerInterfaceP.IsDestroyed(); _ = isCurrentTargetDestroyed
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessPossibleTarget(this.AssociatedAdvancedRTSGameLayer, this, layerInterface, anotherTargetDistance, isShorterThanCurrentTargetDistance, isCurrentTargetDestroyed)
        if this.IsWaypointListEmptyOrOnlyTargets() && this.IsInSensorRange(layerInterface, anotherTargetDistance) && (isShorterThanCurrentTargetDistance || isCurrentTargetDestroyed) {
            this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessSetTarget(this.AssociatedAdvancedRTSGameLayer, this, layerInterface, anotherTargetDistance)
            this.SetTarget(layerInterface, anotherTargetDistance)
        } else if this.IsCloseRange(layerInterface, anotherTargetDistance) && (isShorterThanCurrentTargetDistance || isCurrentTargetDestroyed) {
            this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessPossibleTargetCloser(this.AssociatedAdvancedRTSGameLayer, this, layerInterface, anotherTargetDistance)
            this.SetTarget(layerInterface, anotherTargetDistance)
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) TeleportIfNoProgress() {
        if this.IsTrackingWaypoint() || this.AssociatedAdvancedRTSGameLayer.GetParentLayer().IsDestroyed() {
            if this.ProgressTimeDelayHelper.IsTimeTNT() && this.NextUnvisitedPathGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
                this.AssociatedAdvancedRTSGameLayer.TeleportTo(this.NextUnvisitedPathGeographicMapCellPosition)
            }
            if this.GetCompleteTimeDelayHelper().IsTimeTNT() {
                var geographicMapCellPosition IGeographicMapCellPosition = any(this.CurrentGeographicMapCellHistoryP.GetTracked().Get(this.CurrentGeographicMapCellHistoryP.GetSize() - 1)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
                this.AssociatedAdvancedRTSGameLayer.TeleportTo(geographicMapCellPosition)
            }
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) SetTarget(layerInterface IAdvancedRTSGameLayer, anotherTargetDistance int) {
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.SetTarget(this.AssociatedAdvancedRTSGameLayer, this, layerInterface, anotherTargetDistance)
        this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.TARGET, BasicColorFactory_GetInstance().GREEN)
        this.AssociatedAdvancedRTSGameLayer.SetLoad(0)
        this.SetCurrentTargetDistance(anotherTargetDistance)
        this.SetCurrentTargetLayerInterface(layerInterface)
        this.SetTrackingWaypoint(false)
        this.TargetList.Clear()
        if !this.IsCloseRange(layerInterface, anotherTargetDistance) && this.CanInsertWaypoint(0, this.CurrentTargetLayerInterfaceP) {
            var geographicMapCellPosition IGeographicMapCellPosition = this.AssociatedAdvancedRTSGameLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
            var pathFindingLayerInterface PathFindingLayerInterface = (any(this.CurrentTargetLayerInterfaceP).(PathFindingLayerInterface)); _ = pathFindingLayerInterface
            var waypoint IWaypointBase = pathFindingLayerInterface.GetWaypointBehavior().GetWaypoint(); _ = waypoint
            var list IBasicArrayList = waypoint.GetPathsListFromCacheOnly(geographicMapCellPosition); _ = list
            this.SetWaypointPathsList(list)
            if this.WaypointPathsListP == nil {
                this.TargetWithoutCachedPathLayerInterface = this.CurrentTargetLayerInterfaceP
            } else if this.WaypointPathsListP.Size() != 0 {
                this.SetTargetPath()
            }
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) SetTargetPath() {
        if this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() {
            if this.CurrentTargetLayerInterfaceP.IsDestroyed() {
                this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.SetTargetPath(this.AssociatedAdvancedRTSGameLayer)
                this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.KILL, this.BasicColorFactory.ORANGE)
                this.ClearTarget()
                return
            }
            if this.CurrentTargetLayerInterfaceP == this.WaypointPathRunnable.GetTargetLayer() {
                this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.SetTargetPathIgnoreNewPath(this.AssociatedAdvancedRTSGameLayer, this)
                this.InsertWaypoint(0, this.CurrentTargetLayerInterfaceP)
                this.SetRandomGeographicMapCellHistory(this.WaypointPathsListP)
            }
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) SetGeographicMapCellHistoryPath(geographicMapCellPositionBasicArrayList IBasicArrayList) {
        this.SetCurrentPathGeographicMapCellPosition(SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION)
        this.NextUnvisitedPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
        this.UnitWaypointBehavior.SetGeographicMapCellHistoryPath(geographicMapCellPositionBasicArrayList)
}


// throws Exception
func (this *UnitWaypointBehavior2) ProcessWaypoint() {
        var size int = this.TargetList.Size(); _ = size
        if size > 0 {
            var targetLayer IAdvancedRTSGameLayer = any(this.TargetList.Get(0)).(IAdvancedRTSGameLayer); _ = targetLayer
            this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessWaypoint(this.AssociatedAdvancedRTSGameLayer, this, targetLayer, size)
            if targetLayer.IsDestroyed() {
                this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.WAYPOINT_DESTROYED_SHORT, this.BasicColorFactory.YELLOW)
                this.RemoveWaypoint(targetLayer, this.UnitWaypointStrings.WAYPOINT_DESTROYED)
            } else {
                var geographicMapCellPosition IGeographicMapCellPosition = this.AssociatedAdvancedRTSGameLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
                if this.IsTrackingWaypoint() {
                    this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessWaypointTracked(this.AssociatedAdvancedRTSGameLayer, this)
                    if this.VisitIfAtMidPoint(geographicMapCellPosition) {
                        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessWaypointTrackedVisit(this.AssociatedAdvancedRTSGameLayer, geographicMapCellPosition)
                    }
                    if this.CurrentGeographicMapCellHistoryP.IsAllVisited2() && this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() {
                        var oldWaypointLayer IAdvancedRTSGameLayer = any(this.CurrentTargetLayerInterfaceP).(IAdvancedRTSGameLayer); _ = oldWaypointLayer
                        oldWaypointLayer.GetWaypointBehavior().GetWaypoint().Visit(this.AssociatedAdvancedRTSGameLayer)
                        this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.ALL_VISITED_SHORT, this.BasicColorFactory.GREEN)
                        this.RemoveWaypoint(any(this.CurrentTargetLayerInterfaceP).(IAdvancedRTSGameLayer), this.UnitWaypointStrings.ALL_VISITED)
                    }
                } else if this.CurrentTargetLayerInterfaceP == nil || this.WaypointOverridesAttacking {
                    var list IBasicArrayList = targetLayer.GetWaypointBehavior().GetWaypoint().GetPathsListFromCacheOnly(geographicMapCellPosition); _ = list
                    this.SetWaypointPathsList(list)
                    if this.WaypointPathsListP == nil {
                        this.WaitingOnWaypointPath = true
                        this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.THINKING, this.BasicColorFactory.GREEN)
                        this.RunWaypointPathTask(targetLayer)
                        return
                    }
                    this.SetWaypointPath(targetLayer)
                }
            }
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) Wander() {
        if this.CurrentGeographicMapCellHistoryP.IsAllVisited2() {
            this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.Wander(this.AssociatedAdvancedRTSGameLayer)
            this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.WANDERING, this.BasicColorFactory.RED)
            this.WanderPathsList.Clear()
            this.WanderPathsList.Add(this.AssociatedAdvancedRTSGameLayer.GetSurroundingGeographicMapCellPositionList())
            this.SetRandomGeographicMapCellHistory(this.WanderPathsList)
        }
        this.VisitIfAtMidPoint(this.GetCurrentPathGeographicMapCellPosition())
        this.UpdateCurrentPathGeographicMapCellPosition()
        this.AssociatedAdvancedRTSGameLayer.TrackTo(this.UnitWaypointStrings.NEXT_PATH_NODE)
}


func (this *UnitWaypointBehavior2) VisitIfAtMidPoint(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var unitLayer *UnitLayer = this.AssociatedAdvancedRTSGameLayer; _ = unitLayer
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION && this.NextUnvisitedPathGeographicMapCellPosition == geographicMapCellPosition {
            var point IGPoint = geographicMapCellPosition.GetMidPoint(); _ = point
            var afterNextPoint IGPoint = this.AfterNextUnvisitedPathGeographicMapCellPosition.GetMidPoint(); _ = afterNextPoint
            var beyondMidPoint bool = true; _ = beyondMidPoint
            if geographicMapCellPosition.GetColumn() == this.AfterNextUnvisitedPathGeographicMapCellPosition.GetColumn() {
            } else if point.GetX() < afterNextPoint.GetX() {
                if unitLayer.GetXP() + unitLayer.GetHalfWidth() < point.GetX() {
                    beyondMidPoint = false
                }
            } else {
                if unitLayer.GetXP() + unitLayer.GetHalfWidth() > point.GetX() {
                    beyondMidPoint = false
                }
            }
            if geographicMapCellPosition.GetRow() == this.AfterNextUnvisitedPathGeographicMapCellPosition.GetRow() {
            } else if point.GetY() < afterNextPoint.GetY() {
                if unitLayer.GetYP() + unitLayer.GetHalfHeight() < point.GetY() {
                    beyondMidPoint = false
                }
            } else {
                if unitLayer.GetYP() + unitLayer.GetHalfHeight() > point.GetY() {
                    beyondMidPoint = false
                }
            }
            if beyondMidPoint {
                return this.CurrentGeographicMapCellHistoryP.Visit(geographicMapCellPosition)
            }
        }
        return false
}


// throws Exception
func (this *UnitWaypointBehavior2) ProcessTargetList() {
        this.TargetWithoutCachedPathLayerInterface = CollidableDestroyableDamageableLayer_GetNullInstance()
        for index := int(this.GetPossibleTargetList().Size() - 1); index >= 0; index-- {
            var layerInterface IAdvancedRTSGameLayer = any(this.GetPossibleTargetList().Get(index)).(IAdvancedRTSGameLayer); _ = layerInterface
            if layerInterface.IsDestroyed() {
                this.GetPossibleTargetList().Remove(layerInterface)
            } else {
                this.OnEnemyMovement(layerInterface)
            }
        }
        if this.TargetWithoutCachedPathLayerInterface != CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.WaitingOnTargetPath = true
            this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.THINKING_ABOUT_TARGET, this.BasicColorFactory.GREEN)
            this.RunWaypointPathTask(any(this.CurrentTargetLayerInterfaceP).(IAdvancedRTSGameLayer))
        }
        this.GetPossibleTargetList().Clear()
}


// throws Exception
func (this *UnitWaypointBehavior2) ProcessTargeting() {
        if this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() && (this.IsInSensorRange(this.CurrentTargetLayerInterfaceP, this.GetCurrentTargetDistance()) || this.IsTrackingWaypoint()) {
            if this.CurrentTargetLayerInterfaceP.IsDestroyed() {
                this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.TargetDestroyed(this.AssociatedAdvancedRTSGameLayer)
                this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.KILL, this.BasicColorFactory.ORANGE)
                this.ClearTarget()
                return
            }
            var dx int = 0; _ = dx
            var dy int = 0; _ = dy
            if this.IsTrackingWaypoint() {
                this.UpdateCurrentPathGeographicMapCellPosition()
                var point IGPoint = this.NextUnvisitedPathGeographicMapCellPosition.GetMidPoint(); _ = point
                dx = this.AssociatedAdvancedRTSGameLayer.GetXP() + this.AssociatedAdvancedRTSGameLayer.GetHalfWidth() - point.GetX()
                dy = this.AssociatedAdvancedRTSGameLayer.GetYP() + this.AssociatedAdvancedRTSGameLayer.GetHalfHeight() - point.GetY()
                this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessTargeting(this.AssociatedAdvancedRTSGameLayer, dx, dy)
            } else {
                this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ProcessTargetingNonWayPoint(this.AssociatedAdvancedRTSGameLayer, dx, dy)
                dx = (this.AssociatedAdvancedRTSGameLayer.GetXP() + this.AssociatedAdvancedRTSGameLayer.GetHalfWidth()) - (this.CurrentTargetLayerInterfaceP.GetXP() + this.CurrentTargetLayerInterfaceP.GetHalfWidth())
                dy = (this.AssociatedAdvancedRTSGameLayer.GetYP() + this.AssociatedAdvancedRTSGameLayer.GetHalfHeight()) - (this.CurrentTargetLayerInterfaceP.GetYP() + this.CurrentTargetLayerInterfaceP.GetHalfHeight())
            }
            this.AssociatedAdvancedRTSGameLayer.TrackToDXY(dx, dy)
        } else {
            if this.AssociatedAdvancedRTSGameLayer.GetParentLayer().IsDestroyed() {
                this.Wander()
            } else {
                if this.AssociatedAdvancedRTSGameLayer.ShowMoreCaptionStates {
                    this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.STOP, this.BasicColorFactory.YELLOW)
                }
                this.AssociatedAdvancedRTSGameLayer.AllStop()
            }
        }
}


func (this *UnitWaypointBehavior2) UpdateCurrentPathGeographicMapCellPosition() {
        this.SetLastPathGeographicMapCellPosition(this.GetCurrentPathGeographicMapCellPosition())
        this.SetCurrentPathGeographicMapCellPosition(this.NextUnvisitedPathGeographicMapCellPosition)
        this.NextUnvisitedPathGeographicMapCellPosition = this.CurrentGeographicMapCellHistoryP.GetFirstUnvisited()
        this.AfterNextUnvisitedPathGeographicMapCellPosition = this.CurrentGeographicMapCellHistoryP.GetAfterIfNotLast(this.NextUnvisitedPathGeographicMapCellPosition)
        if this.GetCurrentPathGeographicMapCellPosition() != this.NextUnvisitedPathGeographicMapCellPosition {
            this.ProgressTimeDelayHelper.SetStartTimeTNT()
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) SetWaypointPath(waypointLayer IAdvancedRTSGameLayer) {
        if this.WaypointPathsListP.Size() != 0 {
            this.SetCurrentTargetLayerInterface(waypointLayer)
            var MAX int = Integer_MAX_VALUE; _ = MAX
            this.SetCurrentTargetDistance(MAX)
            this.SetRandomGeographicMapCellHistory(this.WaypointPathsListP)
        } else {
            waypointLayer.GetWaypointBehavior().GetWaypoint().Visit(this.AssociatedAdvancedRTSGameLayer)
            this.AssociatedAdvancedRTSGameLayer.GetCaptionAnimationHelper().Update(this.UnitWaypointStrings.ALREADY_THERE_SHORT, this.BasicColorFactory.YELLOW)
            this.RemoveWaypoint(waypointLayer, this.UnitWaypointStrings.ALREADY_THERE)
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) RunWaypointPathTask(waypointLayer IAdvancedRTSGameLayer) {
        this.SetWaypointPathsList(UnitWaypointBehavior2_RunningWaypointPathList)
        if this.WaypointPathRunnable.IsRunning() {
            panic(errors.New(fmt.Sprint("Should never be running here")))
        }
        this.WaypointPathRunnable.SetRunning(true)
        this.WaypointPathRunnable.SetUnitLayer(this.AssociatedAdvancedRTSGameLayer)
        this.WaypointPathRunnable.SetTargetLayer(waypointLayer)
        this.PathFindingThreadPool.RunTask(this.WaypointPathRunnable)
}


// throws Exception
func (this *UnitWaypointBehavior2) RemoveWaypoint(waypointLayer IRTSLayer, reason string) {
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.RemoveWaypoint(this.AssociatedAdvancedRTSGameLayer, this, waypointLayer, reason)
        this.TargetList.Remove(waypointLayer)
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.RemoveWaypointList(this.AssociatedAdvancedRTSGameLayer, this, this.TargetList)
        if this.CurrentTargetLayerInterfaceP == waypointLayer {
            this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.RemoveWaypointClear(this.AssociatedAdvancedRTSGameLayer)
            this.ClearTarget()
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) ClearTarget() {
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.ClearTarget(this.AssociatedAdvancedRTSGameLayer)
        this.SetCurrentTargetLayerInterface(CollidableDestroyableDamageableLayer_GetNullInstance())
        this.SetTrackingWaypoint(false)
        var MAX int = Integer_MAX_VALUE; _ = MAX
        this.SetCurrentTargetDistance(MAX)
        TrackingEventHandler_GetInstance().FireEvent(this.AssociatedAdvancedRTSGameLayer.GetTrackingEvent())
}


func (this *UnitWaypointBehavior2) IsWaypointListEmptyOrOnlyTargets() bool {
        var list IBasicArrayList = this.TargetList; _ = list
        if list.Size() == 0 {
            return true
        }
        for index := int(list.Size() - 1); index >= 0; index-- {
            var layerInterface IAdvancedRTSGameLayer = any(list.Get(index)).(IAdvancedRTSGameLayer); _ = layerInterface
            if layerInterface.IsWaypointListEmptyOrOnlyTargets() {
                return false
            }
        }
        return true
}


func (this *UnitWaypointBehavior2) IsCloseRange(layerInterface IAdvancedRTSGameLayer, targetDistance int) bool {
        return targetDistance < this.CloseRange + layerInterface.GetHalfHeight()
}


func (this *UnitWaypointBehavior2) IsInSensorRange(layerInterface ICollidableDestroyableDamageableLayer, targetDistance int) bool {
        return targetDistance < this.SensorRange + layerInterface.GetHalfHeight()
}


func (this *UnitWaypointBehavior2) GetCurrentTargetingStateString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if this.CurrentTargetLayerInterfaceP != CollidableDestroyableDamageableLayer_GetNullInstance() {
            stringBuffer.Append(UnitWaypointBehavior2_TARGET_LAYER)
            stringBuffer.Append(this.CommonSeps.SPACE)
            stringBuffer.Append(this.CurrentTargetLayerInterfaceP.GetName())
            stringBuffer.Append(" with ")
            stringBuffer.Append(UnitWaypointBehavior2_TARGET_DISTANCE)
            stringBuffer.Append(this.CommonSeps.SPACE)
            stringBuffer.Appendint(this.GetCurrentTargetDistance())
        }
        return fmt.Sprint(stringBuffer)
}


// throws Exception
func (this *UnitWaypointBehavior2) AddWaypointFromUser(advancedRTSGameLayer IAdvancedRTSGameLayer) {
        if advancedRTSGameLayer.IsDestroyed() {
            panic(errors.New(fmt.Sprint("Trying to add a dead: " + fmt.Sprint(advancedRTSGameLayer))))
        }
        this.AssociatedAdvancedRTSGameLayer.Waypoint2LogHelperP.AddWaypointFromUser(this.AssociatedAdvancedRTSGameLayer, advancedRTSGameLayer)
        this.TargetList.Clear()
        this.TargetList.Add(advancedRTSGameLayer)
        this.ClearTarget()
}


// throws Exception
func (this *UnitWaypointBehavior2) AddBuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) {
        if !this.BuildingSteeringVisitor.GetList().Contains(allbinaryLayer) {
            this.BuildingSteeringVisitor.GetList().Add(allbinaryLayer)
            this.BuildingSteeringVisitor.GetPositionList().Add(cellPosition)
        }
        if !this.GetSteeringVisitorList().Contains(this.BuildingSteeringVisitor) {
            this.GetSteeringVisitorList().Add(this.BuildingSteeringVisitor)
        }
}


// throws Exception
func (this *UnitWaypointBehavior2) BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool {
        return this.AssociatedAdvancedRTSGameLayer.BuildingChase(allbinaryLayer, cellPosition)
}

type BuildingSteeringVisitor struct {
    SteeringVisitor
    PositionList IBasicArrayList
}


func (this *BuildingSteeringVisitor) initFields() {
        this.PositionList = NewBasicArrayListD()
}

func NewBuildingSteeringVisitor() *BuildingSteeringVisitor {
    this := &BuildingSteeringVisitor{}
    this.SteeringVisitor = *NewSteeringVisitor()
    this.initFields()
    return this
}
    

    func (this *UnitWaypointBehavior2) Visit(anyType any) any {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.Put(commonStrings.EXCEPTION, this, "visit", e)
                tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                    }
                }()
                if this.GetList().Size() > 0 {
                    var allbinaryLayer IAllBinaryLayer = any(this.GetList().Get(0)).(IAllBinaryLayer); _ = allbinaryLayer
                    var cellPosition IGeographicMapCellPosition = any(this.GetPositionList().Get(0)).(IGeographicMapCellPosition); _ = cellPosition
                    var clear_ bool = this.BuildingChase(allbinaryLayer, cellPosition); _ = clear_
                    if clear_ {
                        this.GetList().Clear()
                        this.PositionList.Clear()
                        tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                    }
                    tryControl1, tryResult1 = 1, BooleanFactory_GetInstance().FALSE; return
                }
                tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            panic("unreachable")
    }
    

    func (this *UnitWaypointBehavior2) GetPositionList() IBasicArrayList {
            return this.PositionList
    }

