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

type WaypointSelectedLogHelper struct {
    WaypointLogHelper
    LogUtil *LogUtil
    CommonSeps *CommonSeps
    StringUtil *StringUtil
}

var WaypointSelectedLogHelper_InstanceC *WaypointSelectedLogHelper = NewWaypointSelectedLogHelper()

func (this *WaypointSelectedLogHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
}

func NewWaypointSelectedLogHelper() *WaypointSelectedLogHelper {
    this := &WaypointSelectedLogHelper{}
    this.WaypointLogHelper = *NewWaypointLogHelper()
    this.initFields()
    return this
}


func WaypointSelectedLogHelper_GetInstance() *WaypointSelectedLogHelper {
        return WaypointSelectedLogHelper_InstanceC
}


func (this *WaypointSelectedLogHelper) OnWaypointEvent(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" onWaypointEvent: ")
        stringBuffer.Append(advancedRTSGameLayer.GetName())
        if advancedRTSGameLayer.GetParentLayer() != nil {
            stringBuffer.Append(" With Parent: ")
            stringBuffer.Append(advancedRTSGameLayer.GetParentLayer().GetName())
        }
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "onWaypointEvent")
}


func (this *WaypointSelectedLogHelper) AddWaypointFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Adding Waypoint: ")
        stringBuffer.Append(advancedRTSGameLayer.GetName())
        if advancedRTSGameLayer.GetParentLayer() != nil {
            stringBuffer.Append(" With Parent: ")
            stringBuffer.Append(advancedRTSGameLayer.GetParentLayer().GetName())
        }
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: onWaypointEvent")
}


func (this *WaypointSelectedLogHelper) AddWaypointFromBuildingList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, advancedRTSGameLayer PathFindingLayerInterface, list IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" Waypoints: ")
        stringBuffer.Append(this.StringUtil.ToString(list))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: onWaypointEvent")
}


func (this *WaypointSelectedLogHelper) InsertWaypoint(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(waypointName)
        stringBuffer.Append(" Insert: ")
        stringBuffer.Append(rtsLayer.GetName())
        stringBuffer.Append(" for Waypoint: ")
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: insertWaypoint")
}


func (this *WaypointSelectedLogHelper) InsertWaypointList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, index int, rtsLayer ICollidableDestroyableDamageableLayer, waypointName string, list IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(waypointName)
        stringBuffer.Append(" Insert: ")
        stringBuffer.Append(rtsLayer.GetName())
        stringBuffer.Append(" Waypoints: ")
        stringBuffer.Append(this.StringUtil.ToString(list))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: insertWaypoint")
}


func (this *WaypointSelectedLogHelper) SetRandomGeographicMapCellHistoryList(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList) {
        var size int = pathsList.Size(); _ = size
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(CommonLabels_GetInstance().START)
        stringBuffer.Appendint(size)
        stringBuffer.Append(" -> ")
        stringBuffer.Append(this.StringUtil.ToString(pathsList))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: setRandomGeographicMapCellHistory")
}


func (this *WaypointSelectedLogHelper) MoveAwayFromBuilding(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" Trying to move away from building")), this, "selected: moveAwayFromBuilding")
}


func (this *WaypointSelectedLogHelper) NeedToMove(associatedAdvancedRTSGameLayer PathFindingLayerInterface, unitWaypointBehavior IWaypointBehaviorBase) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(unitWaypointBehavior.GetMovementLogicAsString())), this, "needToMove")
}


func (this *WaypointSelectedLogHelper) SetRandomGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" pathsList is null")), this, "setRandomGeographicMapCellHistory")
}


func (this *WaypointSelectedLogHelper) SetCurrentPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingCurrentPathGeographicMapCellPosition IGeographicMapCellPosition, currentPathGeographicMapCellPosition IGeographicMapCellPosition) {
        if existingCurrentPathGeographicMapCellPosition == nil {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" currentPathGeographicMapCellPosition: ").Append(this.StringUtil.ToString(currentPathGeographicMapCellPosition))), this, "setCurrentPathGeographicMapCellPosition")
        }
}


func (this *WaypointSelectedLogHelper) SetNextUnvisitedPathGeographicMapCellPosition(associatedAdvancedRTSGameLayer PathFindingLayerInterface, existingNextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition) {
        if existingNextUnvisitedPathGeographicMapCellPosition != nextUnvisitedPathGeographicMapCellPosition {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" nextUnvisitedPathGeographicMapCellPosition: ").Append(this.StringUtil.ToString(nextUnvisitedPathGeographicMapCellPosition))), this, "setNextUnvisitedPathGeographicMapCellPosition")
        }
}

