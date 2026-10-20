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

type UpgradableRTSLayerHudPaintable struct {
    SelectionHudPaintable
    PERCENT string
    RtsLayer ICollidableDestroyableDamageableLayer
    CostY int
    CostY1 int
    PercentCompleteX2 int
    RtsLayerCompositePaintableLateInit IRTSLayerCompositePaintable
    CharHeight int
    CharWidth int
    PercentComplete int
    PercentCompleteX int
    PercentCompleteArray []rune
    CurrentTotalDigits int
}

var UpgradableRTSLayerHudPaintable_Instance *UpgradableRTSLayerHudPaintable = NewUpgradableRTSLayerHudPaintable()

func (this *UpgradableRTSLayerHudPaintable) initFields() {
    this.PERCENT = "%"
    this.RtsLayer = CollidableDestroyableDamageableLayer_GetNullInstance()
    
    
    
    
    
    
    
    
    this.PercentCompleteArray = CharArrayFactory_GetInstance().GetZeroCharArray()
    
}


func UpgradableRTSLayerHudPaintable_GetInstance() *UpgradableRTSLayerHudPaintable {
        return UpgradableRTSLayerHudPaintable_Instance
}


func NewUpgradableRTSLayerHudPaintable() *UpgradableRTSLayerHudPaintable {
    this := &UpgradableRTSLayerHudPaintable{}
    this.SelectionHudPaintable = *NewSelectionHudPaintable()
    this.initFields()
    return this
}


func (this *UpgradableRTSLayerHudPaintable) UpdateMeasurement(graphics IGraphics) {
        this.SelectionHudPaintable.UpdateMeasurement(graphics)
        var font Font = graphics.GetFont(); _ = font
        this.CharHeight = font.GetHeight()
        this.CharWidth = MyFontProcessor_DefaultCharWidth(font)
}


func (this *UpgradableRTSLayerHudPaintable) Update() {
        this.SelectionHudPaintable.Update()
        var commonButtons *CommonButtons = CommonButtons_GetInstance(); _ = commonButtons
        this.CostY = (this.Y + commonButtons.STANDARD_BUTTON_SIZE)
        this.CostY1 = (this.Y + commonButtons.STANDARD_BUTTON_SIZE - this.CharHeight)
        this.PercentCompleteX2 = this.ImageX + commonButtons.STANDARD_BUTTON_SIZE - this.CharWidth
        var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
        if displayInfoSingleton.GetLastWidth() > 320 {
            this.RtsLayerCompositePaintableLateInit = NewUpgradableWideRTSLayerPaintable(this)
        } else {
            this.RtsLayerCompositePaintableLateInit = NewRTSLayerCompositePaintable(this)
        }
}


func (this *UpgradableRTSLayerHudPaintable) UpdateSelectionInfo() {
        var rtsLayer IRTSLayer = any(this.GetRtsLayer()).(IRTSLayer); _ = rtsLayer
        this.RtsLayerCompositePaintableLateInit.Update(rtsLayer)
        this.SetAnimationInterface(rtsLayer.GetVerticleBuildAnimationInterface())
        this.SetName(rtsLayer.GetName())
}


func (this *UpgradableRTSLayerHudPaintable) UpdateInfo() {
        var rtsLayer IRTSLayer = any(this.GetRtsLayer()).(IRTSLayer); _ = rtsLayer
        this.PercentComplete = rtsLayer.GetPercentComplete()
        if this.PercentComplete < 10 {
            this.PercentCompleteX = 32
        } else if this.PercentComplete < 100 {
            this.PercentCompleteX = 24
        } else {
            this.PercentCompleteX = 16
        }
        this.PercentCompleteArray = this.GetPrimitiveLongUtil().GetCharArray(this.PercentComplete)
        this.CurrentTotalDigits = this.GetPrimitiveLongUtil().GetCurrentTotalDigits()
}


func (this *UpgradableRTSLayerHudPaintable) Paint(graphics IGraphics) {
        this.SelectionHudPaintable.Paint(graphics)
        this.RtsLayerCompositePaintableLateInit.Paint(graphics)
        graphics.DrawChars(this.PercentCompleteArray, 0, this.CurrentTotalDigits, this.ImageX + this.PercentCompleteX, this.CostY, 0)
        graphics.DrawString(this.PERCENT, this.PercentCompleteX2, this.CostY, 0)
        this.GetAnimationInterface().PaintXY(graphics, this.ImageX, this.Y)
}


func (this *UpgradableRTSLayerHudPaintable) SetRtsLayer(rtsLayer IRTSLayer) {
        this.RtsLayer = rtsLayer
}


func (this *UpgradableRTSLayerHudPaintable) GetRtsLayer() ICollidableDestroyableDamageableLayer {
        return this.RtsLayer
}

