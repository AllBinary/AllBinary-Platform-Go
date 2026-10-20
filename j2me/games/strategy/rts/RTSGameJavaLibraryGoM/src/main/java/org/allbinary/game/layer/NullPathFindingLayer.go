/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type NullPathFindingLayer struct {
}

// Implements PathFindingLayerInterface

// The Java class NullPathFindingLayer has subclasses so its variables are this interface
type INullPathFindingLayer interface {
    NullPathFindingLayer_() *NullPathFindingLayer
    GetName() string
    GetEndGeographicMapCellPositionList() IBasicArrayList
    ShouldHandleStartSameAsEnd() bool
    GetGeographicMapCellPositionArea() IGeographicMapCellPositionAreaBase
    GetHudPaintable() ISelectionHudPaintable
    HandleCost(ownerLayer PathFindingLayerInterface)
    GetParentLayer() PathFindingLayerInterface
    GetRTSLayer2LogHelper() IRTSLayer2LogHelper
    SetTarget(targetGameLayer PathFindingLayerInterface)
    Init(geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPositionBasicArrayList IBasicArrayList)
    GetMoveOutOfBuildAreaPath(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    SetClosestGeographicMapCellHistory(pathsList IBasicArrayList)
    TeleportTo(geographicMapCellPosition IGeographicMapCellPosition)
    SetLoad(resource int)
    GetSurroundingGeographicMapCellPositionList() IBasicArrayList
    IsSelected() bool
    ShouldAddWaypointFromBuilding() bool
    IsShowMoreCaptionStates() bool
    TrackTo(reason string)
    TrackToDXY(dx int, dy int)
    BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool
    GetTrackingEvent() *TrackingEvent
    AllStop()
    IsWaypointListEmptyOrOnlyTargets() bool
    GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition
    GetTopLeftGeographicMapCellPosition() IGeographicMapCellPosition
    GetCaptionAnimationHelper() ICaptionAnimationHelperBase
    GetWaypointBehavior() IWaypointBehaviorBase
    GetWaypointLogHelper() IWaypointLogHelper
    GetWaypoint2LogHelper() IWaypoint2LogHelper
    GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper
    GetGroupInterface() []*Group
    IsDestroyed() bool
    GetSourceId() int
}

func (this *NullPathFindingLayer) NullPathFindingLayer_() *NullPathFindingLayer {
    return this
}

var NullPathFindingLayer_NULL_PATH_FINDING_LAYER INullPathFindingLayer = NewNullPathFindingLayer()

func NewNullPathFindingLayer() *NullPathFindingLayer {
    this := &NullPathFindingLayer{}
    return this
}


func (this *NullPathFindingLayer) GetName() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) GetEndGeographicMapCellPositionList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) ShouldHandleStartSameAsEnd() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetGeographicMapCellPositionArea() IGeographicMapCellPositionAreaBase {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetHudPaintable() ISelectionHudPaintable {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) HandleCost(ownerLayer PathFindingLayerInterface) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetParentLayer() PathFindingLayerInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetRTSLayer2LogHelper() IRTSLayer2LogHelper {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) SetTarget(targetGameLayer PathFindingLayerInterface) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) Init(geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPositionBasicArrayList IBasicArrayList) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetMoveOutOfBuildAreaPath(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) SetClosestGeographicMapCellHistory(pathsList IBasicArrayList) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) TeleportTo(geographicMapCellPosition IGeographicMapCellPosition) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) SetLoad(resource int) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) GetSurroundingGeographicMapCellPositionList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) IsSelected() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) ShouldAddWaypointFromBuilding() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) IsShowMoreCaptionStates() bool {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) TrackTo(reason string) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) TrackToDXY(dx int, dy int) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetTrackingEvent() *TrackingEvent {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) AllStop() {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) IsWaypointListEmptyOrOnlyTargets() bool {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) GetTopLeftGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetCaptionAnimationHelper() ICaptionAnimationHelperBase {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetWaypointBehavior() IWaypointBehaviorBase {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetWaypointLogHelper() IWaypointLogHelper {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetWaypoint2LogHelper() IWaypoint2LogHelper {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetGroupInterface() []*Group {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullPathFindingLayer) IsDestroyed() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullPathFindingLayer) GetSourceId() int {
        panic(errors.New(fmt.Sprint()))
}

