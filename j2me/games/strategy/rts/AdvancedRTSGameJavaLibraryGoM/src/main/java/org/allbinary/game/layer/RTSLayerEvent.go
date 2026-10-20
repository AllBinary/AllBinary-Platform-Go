/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
 *  
 * By agreeing to this license you and any business entity you represent are 
 * legally bound to the AllBinary Open License Version 1 legal agreement. 
 *  
 * You may obtain the AllBinary Open License Version 1 legal agreement from 
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository. 
 *  
 * Created By: Travis Berthelot  
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

type RTSLayerEvent struct {
    AllBinaryEventObject
    RtsLayer PathFindingLayerInterface
}


func (this *RTSLayerEvent) initFields() {
    this.RtsLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
}


func NewRTSLayerEvent(rtsLayer PathFindingLayerInterface) *RTSLayerEvent {
    this := &RTSLayerEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(rtsLayer)
        this.initFields()
    return this
}


func (this *RTSLayerEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("RTSLayerEvent: \n")
        stringBuffer.Append("RTSLayer: ")
        stringBuffer.Append(fmt.Sprint(this.GetRtsLayer()))
        return fmt.Sprint(stringBuffer)
}

func (this *RTSLayerEvent) String() string {
    return this.ToString()
}


func (this *RTSLayerEvent) GetRtsLayer() PathFindingLayerInterface {
        return this.RtsLayer
}


func (this *RTSLayerEvent) SetRtsLayer(rtsLayer PathFindingLayerInterface) {
        this.RtsLayer = rtsLayer
}

