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

type WaypointLayerInterface interface {
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
}


