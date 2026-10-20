/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
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

type TrackingEvent struct {
    AllBinaryEventObject
    LayerInterface IAllBinaryLayer
}


func (this *TrackingEvent) initFields() {
    this.LayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
}


func NewTrackingEvent(allBinaryLayerInterface any) *TrackingEvent {
    this := &TrackingEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(allBinaryLayerInterface)
        this.initFields()
        if allBinaryLayerInterface != TrackingEventHandler_GetInstance() {
            this.SetLayerInterface(any(allBinaryLayerInterface).(IAllBinaryLayer))
        }
    return this
}


func (this *TrackingEvent) Init(anyType any) {
        this.SetSource(anyType)
}


func (this *TrackingEvent) GetLayerInterface() IAllBinaryLayer {
        return this.LayerInterface
}


func (this *TrackingEvent) SetLayerInterface(layerInterface IAllBinaryLayer) {
        this.LayerInterface = layerInterface
}


func (this *TrackingEvent) SetLayerInterfaceForCircularStaticPool(layerInterface IAllBinaryLayer) {
        this.LayerInterface = layerInterface
}


func (this *TrackingEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("TrackingEvent: \n")
        stringBuffer.Append("LayerInterface: ")
        stringBuffer.Append(fmt.Sprint(this.LayerInterface))
        return fmt.Sprint(stringBuffer)
}

func (this *TrackingEvent) String() string {
    return this.ToString()
}

