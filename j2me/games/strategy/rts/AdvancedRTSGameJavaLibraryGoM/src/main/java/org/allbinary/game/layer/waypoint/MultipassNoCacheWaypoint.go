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

type MultipassNoCacheWaypoint struct {
    WaypointBase
    LogUtil *LogUtil
    OwnerLayer PathFindingLayerInterface
}

// Implements BuildingEventListenerInterface


func (this *MultipassNoCacheWaypoint) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewMultipassNoCacheWaypoint(ownerLayer PathFindingLayerInterface, sound ISound) *MultipassNoCacheWaypoint {
    this := &MultipassNoCacheWaypoint{}
        this.WaypointBase = *NewWaypointBase(sound)
        this.initFields()
        this.OwnerLayer = ownerLayer
    return this
}


// throws Exception
func (this *MultipassNoCacheWaypoint) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.WaypointBase.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
}


func (this *MultipassNoCacheWaypoint) ReleaseCachedPaths() {
}


// throws Exception
func (this *MultipassNoCacheWaypoint) GetPathFindingInfo(geographicMapCellPosition IGeographicMapCellPosition) *PathFindingInfo {
        var startGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPosition; _ = startGeographicMapCellPosition
        var endGeographicMapCellPosition IGeographicMapCellPosition = this.GetEndGeographicMapCellPosition(startGeographicMapCellPosition); _ = endGeographicMapCellPosition
        if endGeographicMapCellPosition == nil {
            return PathFindingInfo_NULL_PATH_FINDING_INFO
        }
        if startGeographicMapCellPosition == endGeographicMapCellPosition {
            if this.OwnerLayer.ShouldHandleStartSameAsEnd() {
                return PathFindingInfo_NULL_PATH_FINDING_INFO
            } else {
                panic(errors.New(fmt.Sprint("Start should not be End: " + fmt.Sprint(endGeographicMapCellPosition))))
            }
        }
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var raceTrackGeographicMap *RaceTrackGeographicMap = any(geographicMapInterface).(*RaceTrackGeographicMap); _ = raceTrackGeographicMap
        var raceTrackGeographicMapCellTypeFactory IGeographicMapCellTypeFactory = raceTrackGeographicMap.GetGeographicMapCellTypeFactory(); _ = raceTrackGeographicMapCellTypeFactory
        var customMapGenerator *CustomMapGenerator = any(raceTrackGeographicMap.GetCustomMapGenerator()).(*CustomMapGenerator); _ = customMapGenerator
        customMapGenerator.CopyMapIntoCustomMap()
        var customMapArray [][]int = customMapGenerator.GetCustomMapArray(); _ = customMapArray
        customMapArray[startGeographicMapCellPosition.GetRow()][startGeographicMapCellPosition.GetColumn()] = raceTrackGeographicMapCellTypeFactory.GetStartType()
        customMapArray[endGeographicMapCellPosition.GetRow()][endGeographicMapCellPosition.GetColumn()] = raceTrackGeographicMapCellTypeFactory.GetEndType()
        var pathFindingInfo *PathFindingInfo = PathFindingInfoFactory_GetInstance().GetInstancePathFindingInfo(raceTrackGeographicMap, customMapArray); _ = pathFindingInfo
        return pathFindingInfo
}


// throws Exception
func (this *MultipassNoCacheWaypoint) GetPathsList(geographicMapCellPosition IGeographicMapCellPosition, pathFindingInfo *PathFindingInfo, multipassState *MultipassState) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "getPathsList", e)
            tryControl1, tryResult1 = 1, BasicArrayListUtil_GetInstance().GetImmutableInstance(); return
                }
            }()
            if pathFindingInfo == nil {
                this.LogUtil.PutF("pathFindingInfo: " + fmt.Sprint(pathFindingInfo), this, "getPathsList")
                tryControl1, tryResult1 = 1, BasicArrayListUtil_GetInstance().GetImmutableInstance(); return
            }
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            var list IBasicArrayList = PathGenerator_GetInstance().CreateN(geographicMapInterface, pathFindingInfo, 2, multipassState); _ = list
            if list != nil && list != BasicArrayListUtil_GetInstance().GetImmutableInstance() {
                if list.Size() < 1 {
                    panic(errors.New(fmt.Sprint("Path may have been cleared by clearing caching")))
                }
                tryControl1, tryResult1 = 1, list; return
            }
            tryControl1, tryResult1 = 1, BasicArrayListUtil_GetInstance().GetImmutableInstance(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *MultipassNoCacheWaypoint) GetEndGeographicMapCellPosition(startGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        var endList IBasicArrayList = this.OwnerLayer.GetEndGeographicMapCellPositionList(); _ = endList
        var endGeographicMapCellPosition IGeographicMapCellPosition = any(BasicArrayListUtil_GetInstance().GetRandom(endList)).(IGeographicMapCellPosition); _ = endGeographicMapCellPosition
        if startGeographicMapCellPosition == endGeographicMapCellPosition {
            for index := int(endList.Size() - 1); index >= 0; index-- {
                endGeographicMapCellPosition = any(endList.Get(index)).(IGeographicMapCellPosition)
                if startGeographicMapCellPosition != endGeographicMapCellPosition {
                    break
                }
            }
        }
        return endGeographicMapCellPosition
}


func (this *MultipassNoCacheWaypoint) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *MultipassNoCacheWaypoint) OnBuildingEvent(event *RTSLayerEvent) {
}


func (this *MultipassNoCacheWaypoint) Reset() {
        this.GetConnectedWaypointList().Clear()
}


// throws Exception
func (this *MultipassNoCacheWaypoint) Visit(unitLayer PathFindingLayerInterface) {
        if unitLayer.GetGroupInterface()[0] != this.OwnerLayer.GetGroupInterface()[0] {
            return
        }
        var size int = this.GetConnectedWaypointList().Size(); _ = size
        if size > 0 {
            var unitWaypointBehavior IUnitWaypointBehavior = any(unitLayer.GetWaypointBehavior()).(IUnitWaypointBehavior); _ = unitWaypointBehavior
            for this.GetConnectedWaypointList().Size() > 0 {
                var rtsLayer ICollidableDestroyableDamageableLayer = any(this.GetConnectedWaypointList().Get(0)).(ICollidableDestroyableDamageableLayer); _ = rtsLayer
                if rtsLayer.IsDestroyed() {
                    this.GetConnectedWaypointList().Remove(rtsLayer)
                } else {
                    unitLayer.HandleCost(this.OwnerLayer)
                    unitWaypointBehavior.InsertWaypoint(0, rtsLayer)
                    break
                }
            }
        }
}

