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

type WaypointLogHelper struct {
}

// The Java class WaypointLogHelper has subclasses so its variables are this interface
type IWaypointLogHelper interface {
    WaypointLogHelper_() *WaypointLogHelper
    OnWaypointEvent(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface)
    AddWaypointFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface)
    AddWaypointFromBuildingList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface, list IBasicArrayList)
    InsertWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string)
    InsertWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string, list IBasicArrayList)
    SetRandomGeographicMapCellHistoryList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList)
    MoveAwayFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    NeedToMove(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior IWaypointBehaviorBase)
    SetRandomGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    SetCurrentPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingCurrentPathGeographicMapCellPosition IGeographicMapCellPosition, currentPathGeographicMapCellPosition IGeographicMapCellPosition)
    SetNextUnvisitedPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingNextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition)
}

func (this *WaypointLogHelper) WaypointLogHelper_() *WaypointLogHelper {
    return this
}

var WaypointLogHelper_Instance IWaypointLogHelper = NewWaypointLogHelper()

func NewWaypointLogHelper() *WaypointLogHelper {
    this := &WaypointLogHelper{}
    return this
}


func WaypointLogHelper_GetInstance() IWaypointLogHelper {
        return WaypointLogHelper_Instance
}


func (this *WaypointLogHelper) OnWaypointEvent(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *WaypointLogHelper) AddWaypointFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *WaypointLogHelper) AddWaypointFromBuildingList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface, list IBasicArrayList) {
}


func (this *WaypointLogHelper) InsertWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string) {
}


func (this *WaypointLogHelper) InsertWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string, list IBasicArrayList) {
}


func (this *WaypointLogHelper) SetRandomGeographicMapCellHistoryList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList) {
}


func (this *WaypointLogHelper) MoveAwayFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *WaypointLogHelper) NeedToMove(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior IWaypointBehaviorBase) {
}


func (this *WaypointLogHelper) SetRandomGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *WaypointLogHelper) SetCurrentPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingCurrentPathGeographicMapCellPosition IGeographicMapCellPosition, currentPathGeographicMapCellPosition IGeographicMapCellPosition) {
}


func (this *WaypointLogHelper) SetNextUnvisitedPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingNextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition) {
}

