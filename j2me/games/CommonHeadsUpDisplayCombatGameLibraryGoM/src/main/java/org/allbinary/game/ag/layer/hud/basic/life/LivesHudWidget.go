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

type LivesHudWidget struct {
    BasicHud
    LifeInterface *Life
    XArray []int
    AnimationInterface IAnimation
}

// Implements PaintableInterface


func (this *LivesHudWidget) initFields() {
    
    
    
}


func NewLivesHudWidget(animationInterface IAnimation, lifeInterface *Life, location int, direction int) *LivesHudWidget {
    this := &LivesHudWidget{}
        this.BasicHud = *NewBasicHud(location, direction, 2, BasicColorFactory_GetInstance().WHITE)
        this.initFields()
        this.LifeInterface = lifeInterface
        var size int = int(this.GetLifeInterface().GetMaxlives()); _ = size
        this.XArray = make([]int, size)
        this.AnimationInterface = animationInterface
        this.UpdateMaxWidth = int(this.LifeInterface.GetMaxlives()) * 16
        this.UpdateMaxHeight = 16
    return this
}


func (this *LivesHudWidget) UpdateMeasurement(graphics IGraphics) {
        this.BasicHud.UpdateMeasurement(graphics)
        var maxLives int = int(this.GetLifeInterface().GetMaxlives()); _ = maxLives
        for index := int(0); index < maxLives; index++ {
            this.XArray[index] = this.GetX() + (index * 16)
        }
}


func (this *LivesHudWidget) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        this.BasicHud.OnDisplayChangeEvent(displayChangeEvent)
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *LivesHudWidget) SetX(x int) {
        this.BasicHud.SetX(x)
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *LivesHudWidget) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        for index := int(int(this.GetLifeInterface().Get())); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterface.PaintXY(graphics, this.XArray[index], this.GetY())
        }
}


func (this *LivesHudWidget) PaintThreed(graphics IGraphics) {
}


func (this *LivesHudWidget) GetLifeInterface() *Life {
        return this.LifeInterface
}

