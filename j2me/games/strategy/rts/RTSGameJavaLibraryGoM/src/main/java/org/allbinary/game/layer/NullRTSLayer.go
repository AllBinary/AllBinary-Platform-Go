/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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
    . "org/allbinary/game/ag"
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

type NullRTSLayer struct {
    NullPathFindingLayer
}

// Implements RTSPlayerLayerInterface

var NullRTSLayer_NULL_RTS_LAYER *NullRTSLayer = NewNullRTSLayer()

func NewNullRTSLayer() *NullRTSLayer {
    this := &NullRTSLayer{}
    this.NullPathFindingLayer = *NewNullPathFindingLayer()
    return this
}


func (this *NullRTSLayer) GetName() string {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRTSLayer) ProcessInput(layerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) InitInputProcessors() {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetGameNotificationHud() IGameNotificationHud {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetCurrentScrollSelectionForm() IScrollSelectionForm {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) SetCurrentScrollSelectionForm(currentScrollSelectionForm IScrollSelectionForm) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetBuildingScrollSelectionForm() IScrollSelectionForm {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetPlayerGameInput() IPlayerGameInput {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) OnEvent(eventObject IAllBinaryEventObject) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRTSLayer) OnDestroyed(destroyedEvent *DestroyedEvent) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRTSLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetCapital() ICapital {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetRTSFormInput() IRTSFormInput {
        return NullRTSFormInputFactory_GetInstance()
}


func (this *NullRTSLayer) Add(sound ISound) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) GetPlayerOwnedRTSLayers() IPlayerOwnedRTSLayers {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRTSLayer) ImplmentsArtificialIntelligenceCompositeInterface() bool {
        panic(errors.New(fmt.Sprint()))
}

