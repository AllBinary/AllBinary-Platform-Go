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

type Waypoint struct {
    WaypointBase
    LogUtil *LogUtil
    BasicArrayListUtil *BasicArrayListUtil
    OwnerLayer PathFindingLayerInterface
    EndList IBasicArrayList
    Paths [][]IBasicArrayList
}

// Implements BuildingEventListenerInterface

// The Java class Waypoint has subclasses so its variables are this interface
type IWaypoint interface {
    IWaypointBase
    Waypoint_() *Waypoint
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    ReleaseCachedPaths()
    GetPathsListFromCacheOnly(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    GetPathsListRunnable(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    GetEndGeographicMapCellPosition(startGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition
    CreatePaths(startGeographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    OnEvent(eventObject IAllBinaryEventObject)
    OnBuildingEvent(event *RTSLayerEvent)
    ClearPaths()
    Reset()
    Visit(unitLayer PathFindingLayerInterface)
}

func (this *Waypoint) Waypoint_() *Waypoint {
    return this
}


func (this *Waypoint) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    
    this.EndList = this.BasicArrayListUtil.GetImmutableInstance()
    this.Paths = this.BasicArrayListUtil.NULL_ARRAY_OF_ARRAY
}


func NewWaypoint(ownerLayer PathFindingLayerInterface, sound ISound) *Waypoint {
    this := &Waypoint{}
        this.WaypointBase = *NewWaypointBase(sound)
        this.initFields()
        this.OwnerLayer = ownerLayer
    return this
}


// throws Exception
func (this *Waypoint) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.WaypointBase.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        if this.Paths != this.BasicArrayListUtil.NULL_ARRAY_OF_ARRAY {
            panic(errors.New(fmt.Sprint()))
        }
        this.EndList = this.OwnerLayer.GetEndGeographicMapCellPositionList()
        this.Paths = func() [][]IBasicArrayList { array := make([][]IBasicArrayList, tiledLayer.GetColumns()); for index := range array { array[index] = make([]IBasicArrayList, tiledLayer.GetRows()) }; return array }()
        this.ClearPaths()
}


func (this *Waypoint) ReleaseCachedPaths() {
        for columnIndex := int(len(this.Paths)); Java_PreDecrement(&columnIndex) >= 0;  {
            for rowIndex := int(len(this.Paths)); Java_PreDecrement(&rowIndex) >= 0;  {
                var pathsList IBasicArrayList = this.Paths[columnIndex][rowIndex]; _ = pathsList
                if pathsList != this.BasicArrayListUtil.GetImmutableInstance() {
                    BasicGeographicMapExtractedPathsCacheFactory_GetInstance().Release(pathsList)
                }
            }
        }
}


// throws Exception
func (this *Waypoint) GetPathsListFromCacheOnly(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        return this.Paths[geographicMapCellPosition.GetColumn()][geographicMapCellPosition.GetRow()]
}


// throws Exception
func (this *Waypoint) GetPathsListRunnable(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        var pathsList IBasicArrayList = this.Paths[geographicMapCellPosition.GetColumn()][geographicMapCellPosition.GetRow()]; _ = pathsList
        if pathsList == this.BasicArrayListUtil.GetImmutableInstance() {
            pathsList = this.CreatePaths(geographicMapCellPosition)
            this.Paths[geographicMapCellPosition.GetColumn()][geographicMapCellPosition.GetRow()] = pathsList
        }
        return pathsList
}


func (this *Waypoint) GetEndGeographicMapCellPosition(startGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        var endGeographicMapCellPosition IGeographicMapCellPosition = any(this.BasicArrayListUtil.GetRandom(this.EndList)).(IGeographicMapCellPosition); _ = endGeographicMapCellPosition
        if startGeographicMapCellPosition == endGeographicMapCellPosition {
            for index := int(this.EndList.Size() - 1); index >= 0; index-- {
                endGeographicMapCellPosition = any(this.EndList.Get(index)).(IGeographicMapCellPosition)
                if startGeographicMapCellPosition != endGeographicMapCellPosition {
                    break
                }
            }
        }
        return endGeographicMapCellPosition
}


// throws Exception
func (this *Waypoint) CreatePaths(startGeographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        if this.EndList.Size() == 0 {
            return this.BasicArrayListUtil.GetImmutableInstance()
        }
        var endGeographicMapCellPosition IGeographicMapCellPosition = this.GetEndGeographicMapCellPosition(startGeographicMapCellPosition); _ = endGeographicMapCellPosition
        if endGeographicMapCellPosition == nil {
            return this.BasicArrayListUtil.GetImmutableInstance()
        }
        if startGeographicMapCellPosition == endGeographicMapCellPosition {
            if this.OwnerLayer.ShouldHandleStartSameAsEnd() {
                return this.BasicArrayListUtil.GetImmutableInstance()
            } else {
                panic(errors.New(fmt.Sprint("Start should not be End: " + fmt.Sprint(this.EndList.Size()))))
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


func (this *Waypoint) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *Waypoint) OnBuildingEvent(event *RTSLayerEvent) {
        this.ClearPaths()
}


func (this *Waypoint) ClearPaths() {
        for columnIndex := int(len(this.Paths)); Java_PreDecrement(&columnIndex) >= 0;  {
            for rowIndex := int(len(this.Paths)); Java_PreDecrement(&rowIndex) >= 0;  {
                this.Paths[columnIndex][rowIndex] = this.BasicArrayListUtil.GetImmutableInstance()
            }
        }
}


func (this *Waypoint) Reset() {
        this.GetConnectedWaypointList().Clear()
        this.ReleaseCachedPaths()
}


// throws Exception
func (this *Waypoint) Visit(unitLayer PathFindingLayerInterface) {
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

