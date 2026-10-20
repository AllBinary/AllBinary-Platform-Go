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

type WaypointRunnableSelectedLogHelper struct {
    WaypointRunnableLogHelper
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var WaypointRunnableSelectedLogHelper_InstanceC *WaypointRunnableSelectedLogHelper = NewWaypointRunnableSelectedLogHelper()

func (this *WaypointRunnableSelectedLogHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewWaypointRunnableSelectedLogHelper() *WaypointRunnableSelectedLogHelper {
    this := &WaypointRunnableSelectedLogHelper{}
    this.WaypointRunnableLogHelper = *NewWaypointRunnableLogHelper()
    this.initFields()
    return this
}


func WaypointRunnableSelectedLogHelper_GetInstance() *WaypointRunnableSelectedLogHelper {
        return WaypointRunnableSelectedLogHelper_InstanceC
}


func (this *WaypointRunnableSelectedLogHelper) Start(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(this.CommonStrings.START_RUNNABLE)), this, this.CommonStrings.RUN)
}


func (this *WaypointRunnableSelectedLogHelper) End(pathFindingLayerInterface PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(pathFindingLayerInterface.GetName()).Append(this.CommonStrings.END_RUNNABLE)), this, this.CommonStrings.RUN)
}

