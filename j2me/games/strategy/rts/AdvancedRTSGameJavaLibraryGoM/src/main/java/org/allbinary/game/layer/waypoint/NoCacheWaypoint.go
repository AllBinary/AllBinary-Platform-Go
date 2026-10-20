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

type NoCacheWaypoint struct {
    WaypointBase
    OwnerLayer PathFindingLayerInterface
}

// Implements BuildingEventListenerInterface


func (this *NoCacheWaypoint) initFields() {
    
}


func NewNoCacheWaypoint(ownerLayer PathFindingLayerInterface, sound ISound) *NoCacheWaypoint {
    this := &NoCacheWaypoint{}
        this.WaypointBase = *NewWaypointBase(sound)
        this.initFields()
        this.OwnerLayer = ownerLayer
    return this
}


// throws Exception
func (this *NoCacheWaypoint) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.WaypointBase.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
}


func (this *NoCacheWaypoint) ReleaseCachedPaths() {
}


// throws Exception
func (this *NoCacheWaypoint) GetPathsListFromCacheOnly(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        return BasicArrayListUtil_GetInstance().GetImmutableInstance()
}


// throws Exception
func (this *NoCacheWaypoint) GetPathsListRunnable(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        return this.CreatePaths(geographicMapCellPosition)
}


// throws Exception
func (this *NoCacheWaypoint) GetEndGeographicMapCellPosition(startGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
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


// throws Exception
func (this *NoCacheWaypoint) CreatePaths(startGeographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        var endGeographicMapCellPosition IGeographicMapCellPosition = this.GetEndGeographicMapCellPosition(startGeographicMapCellPosition); _ = endGeographicMapCellPosition
        if endGeographicMapCellPosition == nil {
            return BasicArrayListUtil_GetInstance().GetImmutableInstance()
        }
        if startGeographicMapCellPosition == endGeographicMapCellPosition {
            if this.OwnerLayer.ShouldHandleStartSameAsEnd() {
                return BasicArrayListUtil_GetInstance().GetImmutableInstance()
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
        var list IBasicArrayList = PathGenerator_GetInstance().GetInstanceNoCache(geographicMapInterface, pathFindingInfo, 2); _ = list
        if list.Size() < 1 {
            panic(errors.New(fmt.Sprint("Path may have been cleared by clearing caching")))
        }
        return list
}


func (this *NoCacheWaypoint) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *NoCacheWaypoint) OnBuildingEvent(event *RTSLayerEvent) {
}


func (this *NoCacheWaypoint) Reset() {
        this.GetConnectedWaypointList().Clear()
}


// throws Exception
func (this *NoCacheWaypoint) Visit(unitLayer PathFindingLayerInterface) {
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

