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

type RTSLayer2LogHelper struct {
}

// The Java class RTSLayer2LogHelper has subclasses so its variables are this interface
type IRTSLayer2LogHelper interface {
    RTSLayer2LogHelper_() *RTSLayer2LogHelper
    SteeringUp(pathFindingLayerInterface PathFindingLayerInterface)
    SteeringFireOrStop(pathFindingLayerInterface PathFindingLayerInterface)
    BuildingAbove(pathFindingLayerInterface PathFindingLayerInterface)
    BuildingChaseLeft(pathFindingLayerInterface PathFindingLayerInterface, angle int)
    BuildingChaseRight(pathFindingLayerInterface PathFindingLayerInterface, angle int)
    BuildingDown(pathFindingLayerInterface PathFindingLayerInterface)
    BuildingChaseDown(pathFindingLayerInterface PathFindingLayerInterface, angle int)
    BuildingChaseUp(pathFindingLayerInterface PathFindingLayerInterface, angle int)
    BuildingLeft()
    BuildingRight()
}

func (this *RTSLayer2LogHelper) RTSLayer2LogHelper_() *RTSLayer2LogHelper {
    return this
}

var RTSLayer2LogHelper_Instance IRTSLayer2LogHelper = NewRTSLayer2LogHelper()

func NewRTSLayer2LogHelper() *RTSLayer2LogHelper {
    this := &RTSLayer2LogHelper{}
    return this
}


func RTSLayer2LogHelper_GetInstance() IRTSLayer2LogHelper {
        return RTSLayer2LogHelper_Instance
}


func (this *RTSLayer2LogHelper) SteeringUp(pathFindingLayerInterface PathFindingLayerInterface) {
}


func (this *RTSLayer2LogHelper) SteeringFireOrStop(pathFindingLayerInterface PathFindingLayerInterface) {
}


func (this *RTSLayer2LogHelper) BuildingAbove(pathFindingLayerInterface PathFindingLayerInterface) {
}


func (this *RTSLayer2LogHelper) BuildingChaseLeft(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
}


func (this *RTSLayer2LogHelper) BuildingChaseRight(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
}


func (this *RTSLayer2LogHelper) BuildingDown(pathFindingLayerInterface PathFindingLayerInterface) {
}


func (this *RTSLayer2LogHelper) BuildingChaseDown(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
}


func (this *RTSLayer2LogHelper) BuildingChaseUp(pathFindingLayerInterface PathFindingLayerInterface, angle int) {
}


func (this *RTSLayer2LogHelper) BuildingLeft() {
}


func (this *RTSLayer2LogHelper) BuildingRight() {
}

