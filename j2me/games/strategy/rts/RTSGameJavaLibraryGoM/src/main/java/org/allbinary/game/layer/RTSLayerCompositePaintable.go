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

type RTSLayerCompositePaintable struct {
    Paintable
    UpgradeCost string
    DownGradeCost string
    UPGRADE string
    DOWNGRADE string
    UpgradableRTSLayerHudPaintable *UpgradableRTSLayerHudPaintable
}

// The Java class RTSLayerCompositePaintable has subclasses so its variables are this interface
type IRTSLayerCompositePaintable interface {
    IPaintable
    RTSLayerCompositePaintable_() *RTSLayerCompositePaintable
    Update(rtsLayer IRTSLayer)
    Paint(graphics IGraphics)
    SetUpgradeCost(upgradeCost string)
    GetUpgradeCost() string
    SetDownGradeCost(downGradeCost string)
    GetDownGradeCost() string
}

func (this *RTSLayerCompositePaintable) RTSLayerCompositePaintable_() *RTSLayerCompositePaintable {
    return this
}


func (this *RTSLayerCompositePaintable) initFields() {
    this.UpgradeCost = StringUtil_GetInstance().EMPTY_STRING
    this.DownGradeCost = StringUtil_GetInstance().EMPTY_STRING
    this.UPGRADE = "Up: $"
    this.DOWNGRADE = "Down: $-"
    
}


func NewRTSLayerCompositePaintable(upgradableRTSLayerHudPaintable *UpgradableRTSLayerHudPaintable) *RTSLayerCompositePaintable {
    this := &RTSLayerCompositePaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.UpgradableRTSLayerHudPaintable = upgradableRTSLayerHudPaintable
    return this
}


func (this *RTSLayerCompositePaintable) Update(rtsLayer IRTSLayer) {
        if rtsLayer.IsUpgradeable() {
            this.SetUpgradeCost(fmt.Sprint(NewStringMaker().Append(this.UPGRADE).Appendint(rtsLayer.GetUpgradeCost())))
        } else {
            this.SetUpgradeCost(StringUtil_GetInstance().EMPTY_STRING)
        }
        if rtsLayer.IsDowngradeable() {
            this.SetDownGradeCost(fmt.Sprint(NewStringMaker().Append(this.DOWNGRADE).Appendint(rtsLayer.GetDowngradeCost())))
        } else {
            this.SetDownGradeCost(StringUtil_GetInstance().EMPTY_STRING)
        }
}


func (this *RTSLayerCompositePaintable) Paint(graphics IGraphics) {
        graphics.DrawString(this.GetDownGradeCost(), this.UpgradableRTSLayerHudPaintable.TextX, this.UpgradableRTSLayerHudPaintable.CostY1, 0)
        graphics.DrawString(this.GetUpgradeCost(), this.UpgradableRTSLayerHudPaintable.TextX, this.UpgradableRTSLayerHudPaintable.CostY, 0)
}


func (this *RTSLayerCompositePaintable) SetUpgradeCost(upgradeCost string) {
        this.UpgradeCost = upgradeCost
}


func (this *RTSLayerCompositePaintable) GetUpgradeCost() string {
        return this.UpgradeCost
}


func (this *RTSLayerCompositePaintable) SetDownGradeCost(downGradeCost string) {
        this.DownGradeCost = downGradeCost
}


func (this *RTSLayerCompositePaintable) GetDownGradeCost() string {
        return this.DownGradeCost
}

