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

type RTSLayer2SelectedLogHelper struct {
    RTSLayer2LogHelper
    LogUtil *LogUtil
    TRACK_TO_FIRE_OR_MOVE string
    BUILDING_CHASE string
}

var RTSLayer2SelectedLogHelper_InstanceC *RTSLayer2SelectedLogHelper = NewRTSLayer2SelectedLogHelper()

func (this *RTSLayer2SelectedLogHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TRACK_TO_FIRE_OR_MOVE = "trackTo:fireOrMove"
    this.BUILDING_CHASE = "buildingChase"
}

func NewRTSLayer2SelectedLogHelper() *RTSLayer2SelectedLogHelper {
    this := &RTSLayer2SelectedLogHelper{}
    this.RTSLayer2LogHelper = *NewRTSLayer2LogHelper()
    this.initFields()
    return this
}


func RTSLayer2SelectedLogHelper_GetInstance() *RTSLayer2SelectedLogHelper {
        return RTSLayer2SelectedLogHelper_InstanceC
}


func (this *RTSLayer2SelectedLogHelper) SteeringUp(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Up")), this, this.TRACK_TO_FIRE_OR_MOVE)
}


func (this *RTSLayer2SelectedLogHelper) SteeringFireOrStop(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Fire/All Stop")), this, this.TRACK_TO_FIRE_OR_MOVE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingAbove(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" Building Above")), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingChaseLeft(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Going Left: ").Appendint(angle)), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingChaseRight(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Going Right: ").Appendint(angle)), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingDown(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" Building Below")), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingChaseDown(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Going Down: ").Appendint(angle)), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingChaseUp(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" steering - Going Up: ").Appendint(angle)), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingLeft(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" Building Left")), this, this.BUILDING_CHASE)
}


func (this *RTSLayer2SelectedLogHelper) BuildingRight(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(" Building Right")), this, this.BUILDING_CHASE)
}

