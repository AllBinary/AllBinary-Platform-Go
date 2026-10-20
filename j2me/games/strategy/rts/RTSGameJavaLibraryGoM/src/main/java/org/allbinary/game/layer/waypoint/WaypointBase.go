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

type WaypointBase struct {
    ConnectedWaypointList IBasicArrayList
    Sound ISound
    AllBinaryGameLayerManagerP IAllBinaryGameLayerManager
}

// Implements EventListenerInterface

// The Java class WaypointBase has subclasses so its variables are this interface
type IWaypointBase interface {
    WaypointBase_() *WaypointBase
    GetSound() ISound
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    GetConnectedWaypointList() IBasicArrayList
    OnEvent(eventObject IAllBinaryEventObject)
    GetPathFindingInfo(geographicMapCellPosition IGeographicMapCellPosition) *PathFindingInfo
    GetPathsList(geographicMapCellPosition IGeographicMapCellPosition, pathFindingInfo *PathFindingInfo, multipassState *MultipassState) IBasicArrayList
    GetPathsListRunnable(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    GetPathsListFromCacheOnly(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    Visit(unitLayer PathFindingLayerInterface)
    Reset()
    GetType() int
}

func (this *WaypointBase) WaypointBase_() *WaypointBase {
    return this
}

var WaypointBase_NULL_WAYPOINT_BASE IWaypointBase = NewWaypointBase(NoSound_GetInstance())

func (this *WaypointBase) initFields() {
    this.ConnectedWaypointList = NewBasicArrayListD()
    
    this.AllBinaryGameLayerManagerP = AllBinaryGameLayerManager_GetNullInstance()
}


func NewWaypointBase(sound ISound) *WaypointBase {
    this := &WaypointBase{}
    this.initFields()
        this.Sound = sound
    return this
}


func (this *WaypointBase) GetSound() ISound {
        return this.Sound
}


// throws Exception
func (this *WaypointBase) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.AllBinaryGameLayerManagerP = allBinaryGameLayerManager
}


func (this *WaypointBase) GetConnectedWaypointList() IBasicArrayList {
        return this.ConnectedWaypointList
}


func (this *WaypointBase) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *WaypointBase) GetPathFindingInfo(geographicMapCellPosition IGeographicMapCellPosition) *PathFindingInfo {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBase) GetPathsList(geographicMapCellPosition IGeographicMapCellPosition, pathFindingInfo *PathFindingInfo, multipassState *MultipassState) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBase) GetPathsListRunnable(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBase) GetPathsListFromCacheOnly(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *WaypointBase) Visit(unitLayer PathFindingLayerInterface) {
}


func (this *WaypointBase) Reset() {
}


func (this *WaypointBase) GetType() int {
        return 0
}

