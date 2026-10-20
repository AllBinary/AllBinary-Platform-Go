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

type UpgradableWideRTSLayerPaintable struct {
    RTSLayerCompositePaintable
    Cost string
}


func (this *UpgradableWideRTSLayerPaintable) initFields() {
    this.Cost = StringUtil_GetInstance().EMPTY_STRING
}


func NewUpgradableWideRTSLayerPaintable(upgradableRTSLayerHudPaintable *UpgradableRTSLayerHudPaintable) *UpgradableWideRTSLayerPaintable {
    this := &UpgradableWideRTSLayerPaintable{}
        this.RTSLayerCompositePaintable = *NewRTSLayerCompositePaintable(upgradableRTSLayerHudPaintable)
        this.initFields()
    return this
}


func (this *UpgradableWideRTSLayerPaintable) Update(rtsLayer IRTSLayer) {
        this.RTSLayerCompositePaintable.Update(rtsLayer)
        this.SetCost(fmt.Sprint(NewStringMaker().Append(this.GetUpgradeCost()).Append(CommonSeps_GetInstance().SPACE).Append(this.GetDownGradeCost())))
}


func (this *UpgradableWideRTSLayerPaintable) Paint(graphics IGraphics) {
        graphics.DrawString(this.GetCost(), this.UpgradableRTSLayerHudPaintable.TextX, this.UpgradableRTSLayerHudPaintable.CostY, 0)
}


func (this *UpgradableWideRTSLayerPaintable) SetCost(cost string) {
        this.Cost = cost
}


func (this *UpgradableWideRTSLayerPaintable) GetCost() string {
        return this.Cost
}

