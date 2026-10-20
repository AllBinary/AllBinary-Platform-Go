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

type HealthBar struct {
    Paintable
    AnimationInterface IHealthBarAnimation
    HealthInterface HealthInterface
    AllbinaryLayer IAllBinaryLayer
}

// Implements HealthListenerInterface


func (this *HealthBar) initFields() {
    
    
    
}


func NewHealthBar(layerInterface IAllBinaryLayer, healthInterface *Health, animationInterface IHealthBarAnimation, direction int) *HealthBar {
    this := &HealthBar{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.AllbinaryLayer = layerInterface
        this.AnimationInterface = animationInterface
        this.HealthInterface = healthInterface
        this.HealthInterface.AddListener(this)
        this.OnHealthChange()
    return this
}


func (this *HealthBar) OnHealthChange() {
        this.AnimationInterface.OnHealthChange(this.HealthInterface.GetHealth() * this.AllbinaryLayer.GetWidth() / this.HealthInterface.GetMaxHealth())
}


func (this *HealthBar) Paint(graphics IGraphics) {
        this.AnimationInterface.PaintXY(graphics, 0, 0)
}


func (this *HealthBar) PaintThreed(graphics IGraphics) {
        this.AnimationInterface.PaintThreedXYZ(graphics, 0, 0, 0)
}

