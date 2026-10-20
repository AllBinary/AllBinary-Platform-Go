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

type RTSLayerHudPaintable struct {
    SelectionHudPaintable
    WeaponProperties []string
    RtsLayer ICollidableDestroyableDamageableLayer
    CostY int
    CostY1 int
    FontHeight int
}

// Implements UpdateMyFontInterface

var RTSLayerHudPaintable_Instance *RTSLayerHudPaintable = NewRTSLayerHudPaintable()

func (this *RTSLayerHudPaintable) initFields() {
    this.WeaponProperties = StringUtil_GetInstance().GetArrayInstance()
    this.RtsLayer = CollidableDestroyableDamageableLayer_GetNullInstance()
    
    
    
}


func RTSLayerHudPaintable_GetInstance() *RTSLayerHudPaintable {
        return RTSLayerHudPaintable_Instance
}


func NewRTSLayerHudPaintable() *RTSLayerHudPaintable {
    this := &RTSLayerHudPaintable{}
    this.SelectionHudPaintable = *NewSelectionHudPaintable()
    this.initFields()
    return this
}


func (this *RTSLayerHudPaintable) UpdateMeasurement(graphics IGraphics) {
        this.SelectionHudPaintable.UpdateMeasurement(graphics)
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.CostY1 = (this.Y + ((len(this.WeaponProperties) + 1) * this.FontHeight))
        if !AndroidUtil_IsAndroid() {
            this.CostY = this.CostY1
        } else {
            this.CostY = (this.Y + ((len(this.WeaponProperties) + 2) * this.FontHeight))
        }
}


func (this *RTSLayerHudPaintable) UpdateSelectionInfo() {
        this.SetName(this.GetRtsLayer().GetName())
        var partInterface IBasicWeaponPart = any(this.GetRtsLayer().GetPartInterfaceArray()[0]).(IBasicWeaponPart); _ = partInterface
        this.WeaponProperties = partInterface.GetWeaponProperties().ToStringArray()
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *RTSLayerHudPaintable) Paint(graphics IGraphics) {
        this.SelectionHudPaintable.Paint(graphics)
        var size int = len(this.WeaponProperties); _ = size
        for index := int(0); index < size; index++ {
            graphics.DrawString(this.WeaponProperties[index], this.TextX, this.Y + ((index + 1) * this.FontHeight), 0)
        }
}


func (this *RTSLayerHudPaintable) SetRtsLayer(rtsLayer IRTSLayer) {
        this.RtsLayer = rtsLayer
}


func (this *RTSLayerHudPaintable) GetRtsLayer() ICollidableDestroyableDamageableLayer {
        return this.RtsLayer
}

