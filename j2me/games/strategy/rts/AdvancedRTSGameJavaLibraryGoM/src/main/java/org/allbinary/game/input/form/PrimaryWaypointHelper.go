/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type PrimaryWaypointHelper struct {
    WaypointLayer PathFindingLayerInterface
}

var PrimaryWaypointHelper_Instance *PrimaryWaypointHelper = NewPrimaryWaypointHelper()

func (this *PrimaryWaypointHelper) initFields() {
    this.WaypointLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
}


func PrimaryWaypointHelper_GetInstance() *PrimaryWaypointHelper {
        return PrimaryWaypointHelper_Instance
}


func NewPrimaryWaypointHelper() *PrimaryWaypointHelper {
    this := &PrimaryWaypointHelper{}
    this.initFields()
    return this
}


func (this *PrimaryWaypointHelper) GetWaypointLayer() PathFindingLayerInterface {
        return this.WaypointLayer
}


func (this *PrimaryWaypointHelper) SetWaypointLayer(waypointLayer IWaypointLayer) {
        this.WaypointLayer = waypointLayer
}

