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

type WaypointPathRunnableBase struct {
    LayerDistanceUtil *LayerDistanceUtil
    PathFindingLayer PathFindingLayerInterface
    TargetPathFindingLayer PathFindingLayerInterface
    PriorityP int
    RunningP bool
}

// Implements RunnableInterface, PriorityRunnable

// The Java class WaypointPathRunnableBase has subclasses so its variables are this interface
type IWaypointPathRunnableBase interface {
    WaypointPathRunnableBase_() *WaypointPathRunnableBase
    IsRunning() bool
    SetRunning(isRunning bool)
    SetThread(thread IThread)
    SetTargetLayer(waypointLayer PathFindingLayerInterface)
    SetPriority()
    GetTargetLayer() PathFindingLayerInterface
    SetUnitLayer(unitLayer PathFindingLayerInterface)
    GetPriority() int
    IsDone() bool
    Run()
    Reset()
}

func (this *WaypointPathRunnableBase) WaypointPathRunnableBase_() *WaypointPathRunnableBase {
    return this
}


func (this *WaypointPathRunnableBase) initFields() {
    this.LayerDistanceUtil = LayerDistanceUtil_GetInstance()
    this.PathFindingLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
    this.TargetPathFindingLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
    this.PriorityP = Integer_MAX_VALUE
    
}

func NewWaypointPathRunnableBase() *WaypointPathRunnableBase {
    this := &WaypointPathRunnableBase{}
    this.initFields()
    return this
}


func (this *WaypointPathRunnableBase) IsRunning() bool {
        return this.RunningP
}


func (this *WaypointPathRunnableBase) SetRunning(isRunning bool) {
        this.RunningP = isRunning
}


// throws Exception
func (this *WaypointPathRunnableBase) SetThread(thread IThread) {
}


func (this *WaypointPathRunnableBase) SetTargetLayer(waypointLayer PathFindingLayerInterface) {
        this.TargetPathFindingLayer = waypointLayer
        this.SetPriority()
}


func (this *WaypointPathRunnableBase) SetPriority() {
        var pathFindingLayer IAllBinaryLayer = any(this.PathFindingLayer).(IAllBinaryLayer); _ = pathFindingLayer
        var targetLayer IAllBinaryLayer = any(this.TargetPathFindingLayer).(IAllBinaryLayer); _ = targetLayer
        if targetLayer != nil {
            var distance int = this.LayerDistanceUtil.GetDistance(targetLayer, pathFindingLayer); _ = distance
            var distanceCategory int = distance / 70; _ = distanceCategory
            this.PriorityP = distanceCategory
        }
}


func (this *WaypointPathRunnableBase) GetTargetLayer() PathFindingLayerInterface {
        return this.TargetPathFindingLayer
}


func (this *WaypointPathRunnableBase) SetUnitLayer(unitLayer PathFindingLayerInterface) {
        this.PathFindingLayer = unitLayer
}


func (this *WaypointPathRunnableBase) GetPriority() int {
        return this.PriorityP
}


func (this *WaypointPathRunnableBase) IsDone() bool {
        return true
}


func (this *WaypointPathRunnableBase) Run() {
}


func (this *WaypointPathRunnableBase) Reset() {
}

