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

type WaypointInfoHudPaintable struct {
    SelectionHudPaintable
    KeyvalueDrawString *KeyValueDrawString
    RtsLayerP PathFindingLayerInterface
    TextLine2Y int
}

// The Java class WaypointInfoHudPaintable has subclasses so its variables are this interface
type IWaypointInfoHudPaintable interface {
    ISelectionHudPaintable
    WaypointInfoHudPaintable_() *WaypointInfoHudPaintable
    UpdateMeasurement(graphics IGraphics)
    UpdateSelectionInfo()
    SetRtsLayer(rtsLayer IRTSLayer)
    GetRtsLayer() PathFindingLayerInterface
    Paint(graphics IGraphics)
}

func (this *WaypointInfoHudPaintable) WaypointInfoHudPaintable_() *WaypointInfoHudPaintable {
    return this
}


func (this *WaypointInfoHudPaintable) initFields() {
    
    this.RtsLayerP = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
    
}


func NewWaypointInfoHudPaintable() *WaypointInfoHudPaintable {
    this := &WaypointInfoHudPaintable{}
    this.SelectionHudPaintable = *NewSelectionHudPaintable()
    this.initFields()
        this.KeyvalueDrawString = NewKeyValueDrawString("Owner: ", this.TextX)
    return this
}


func (this *WaypointInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
        this.SelectionHudPaintable.UpdateMeasurement(graphics)
        var font Font = graphics.GetFont(); _ = font
        this.TextLine2Y = (this.Y + font.GetHeight())
}


func (this *WaypointInfoHudPaintable) UpdateSelectionInfo() {
        var rtsLayer IRTSLayer = any(this.RtsLayerP).(IRTSLayer); _ = rtsLayer
        this.SetName(rtsLayer.GetName())
        this.SetAnimationInterface(rtsLayer.GetVerticleBuildAnimationInterface())
        this.KeyvalueDrawString.Update(rtsLayer.GetParentLayer().GetName())
}


func (this *WaypointInfoHudPaintable) SetRtsLayer(rtsLayer IRTSLayer) {
        this.RtsLayerP = rtsLayer
}


func (this *WaypointInfoHudPaintable) GetRtsLayer() PathFindingLayerInterface {
        return this.RtsLayerP
}


func (this *WaypointInfoHudPaintable) Paint(graphics IGraphics) {
        this.SelectionHudPaintable.Paint(graphics)
        this.KeyvalueDrawString.Paint(graphics, this.TextLine2Y)
        this.GetAnimationInterface().PaintXY(graphics, this.ImageX, this.Y)
}

