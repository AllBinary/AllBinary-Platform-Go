/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type WaypointPathRunnable struct {
    WaypointPathRunnableBase
    LogUtil *LogUtil
}


func (this *WaypointPathRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewWaypointPathRunnable() *WaypointPathRunnable {
    this := &WaypointPathRunnable{}
    this.WaypointPathRunnableBase = *NewWaypointPathRunnableBase()
    this.initFields()
    return this
}


func (this *WaypointPathRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
            this.SetRunning(false)
                }
            }()
            this.PathFindingLayer.GetWaypointRunnableLogHelper().Start(this.PathFindingLayer)
            var waypointBehavior IWaypointBehaviorBase = this.PathFindingLayer.GetWaypointBehavior(); _ = waypointBehavior
            var geographicMapCellPosition IGeographicMapCellPosition = this.PathFindingLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
            if geographicMapCellPosition == nil {
                panic(errors.New(fmt.Sprint("Should never be running here")))
            }
            waypointBehavior.SetWaypointPathsList(this.TargetPathFindingLayer.GetWaypointBehavior().GetWaypoint().GetPathsListRunnable(geographicMapCellPosition))
            this.PathFindingLayer.GetWaypointRunnableLogHelper().End(this.PathFindingLayer)
            time.Sleep(time.Duration(15 + int64((this.PriorityP * 2))) * time.Millisecond)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WaypointPathRunnable) IsDone() bool {
        return true
}

