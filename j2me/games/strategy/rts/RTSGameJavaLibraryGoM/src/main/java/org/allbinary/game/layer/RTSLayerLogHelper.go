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

type RTSLayerLogHelper struct {
}

// The Java class RTSLayerLogHelper has subclasses so its variables are this interface
type IRTSLayerLogHelper interface {
    RTSLayerLogHelper_() *RTSLayerLogHelper
    SetClosestGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList)
    TrackTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, dx int, dy int, reason string)
    TurnTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int, angleInfo *AngleInfo, angle int, movementAngle *NamedAngle, evading bool, targetAngle int)
    DoneMoving(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    CloseEnough(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    MovingLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    MovingRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    MovingUp(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    MovingDown(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    CurrentMoveEnded(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    Evade(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    RotateLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    RotateRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    Handle(associatedAdvancedRTSGameLayer PathFindingLayerInterface, movementAngle *NamedAngle)
    NoRotation(associatedAdvancedRTSGameLayer PathFindingLayerInterface)
    NotOnPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellHistory *GeographicMapCellHistory, currentGeographicMapCellPosition IGeographicMapCellPosition, pathList IBasicArrayList)
}

func (this *RTSLayerLogHelper) RTSLayerLogHelper_() *RTSLayerLogHelper {
    return this
}

var RTSLayerLogHelper_Instance IRTSLayerLogHelper = NewRTSLayerLogHelper()

func NewRTSLayerLogHelper() *RTSLayerLogHelper {
    this := &RTSLayerLogHelper{}
    return this
}


func RTSLayerLogHelper_GetInstance() IRTSLayerLogHelper {
        return RTSLayerLogHelper_Instance
}


func (this *RTSLayerLogHelper) SetClosestGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList) {
}


func (this *RTSLayerLogHelper) TrackTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, dx int, dy int, reason string) {
}


func (this *RTSLayerLogHelper) TurnTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int, angleInfo *AngleInfo, angle int, movementAngle *NamedAngle, evading bool, targetAngle int) {
}


func (this *RTSLayerLogHelper) DoneMoving(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) CloseEnough(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) MovingLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) MovingRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) MovingUp(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) MovingDown(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) CurrentMoveEnded(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) Evade(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) RotateLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) RotateRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) Handle(associatedAdvancedRTSGameLayer PathFindingLayerInterface, movementAngle *NamedAngle) {
}


func (this *RTSLayerLogHelper) NoRotation(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
}


func (this *RTSLayerLogHelper) NotOnPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellHistory *GeographicMapCellHistory, currentGeographicMapCellPosition IGeographicMapCellPosition, pathList IBasicArrayList) {
}

