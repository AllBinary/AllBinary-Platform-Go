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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type HealthIncludingLifeBar struct {
    Paintable
    AnimationInterface IHealthBarAnimation
    Life *Life
    HealthInterface HealthInterface
    AllbinaryLayer IAllBinaryLayer
}

// Implements HealthListenerInterface


func (this *HealthIncludingLifeBar) initFields() {
    
    
    
    
}


// throws Exception
func HealthIncludingLifeBar_CreateHealthIncludingLifeBar(layerInterface IAllBinaryLayer, life *Life, healthInterface *Health, location int, direction int) *HealthIncludingLifeBar {
        return NewHealthIncludingLifeBar(layerInterface, life, healthInterface, NewHealthBarTwodAnimation(layerInterface, location), direction)
}


func NewHealthIncludingLifeBar(layerInterface IAllBinaryLayer, life *Life, healthInterface *Health, animationInterface IHealthBarAnimation, direction int) *HealthIncludingLifeBar {
    this := &HealthIncludingLifeBar{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.AllbinaryLayer = layerInterface
        this.AnimationInterface = animationInterface
        this.Life = life
        this.HealthInterface = healthInterface
        this.HealthInterface.AddListener(this)
        this.OnHealthChange()
    return this
}


func (this *HealthIncludingLifeBar) OnHealthChange() {
        this.AnimationInterface.OnHealthChange(((this.HealthInterface.GetHealth() + (this.HealthInterface.GetMaxHealth() * int(this.Life.Get()))) * this.AllbinaryLayer.GetWidth()) / (this.HealthInterface.GetMaxHealth() + (this.HealthInterface.GetMaxHealth() * int(this.Life.GetStartLives()))))
}


func (this *HealthIncludingLifeBar) Paint(graphics IGraphics) {
        this.AnimationInterface.PaintXY(graphics, 0, 0)
}


func (this *HealthIncludingLifeBar) PaintThreed(graphics IGraphics) {
        this.AnimationInterface.PaintThreedXYZ(graphics, 0, 0, 0)
}

