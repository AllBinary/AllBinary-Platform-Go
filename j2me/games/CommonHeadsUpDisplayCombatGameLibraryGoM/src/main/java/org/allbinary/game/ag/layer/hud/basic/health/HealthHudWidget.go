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

type HealthHudWidget struct {
    BasicHud
    HealthScale int
    AnimationInterface IAnimation
    HealthInterface *Health
    Max int
    XArray []int
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    TimeDelayHelper ITimeDelayHelper
    SlowBeatTimeDelayHelper ITimeDelayHelper
    MediumBeatTimeDelayHelper ITimeDelayHelper
    FastBeatTimeDelayHelper ITimeDelayHelper
}

// Implements PaintableInterface, HealthListenerInterface


func (this *HealthHudWidget) initFields() {
    
    
    
    
    
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.TimeDelayHelper = NoTimeDelayHelper_SINGLETON
    this.SlowBeatTimeDelayHelper = NewTimeDelayHelper(1280)
    this.MediumBeatTimeDelayHelper = NewTimeDelayHelper(640)
    this.FastBeatTimeDelayHelper = NewTimeDelayHelper(320)
}


func NewHealthHudWidget(animationInterface IAnimation, healthInterface *Health, location int, direction int) *HealthHudWidget {
    this := &HealthHudWidget{}
        this.BasicHud = *NewBasicHud(location, direction, 2, BasicColorFactory_GetInstance().WHITE)
        this.initFields()
        this.AnimationInterface = animationInterface
        this.HealthInterface = healthInterface
        this.HealthInterface.AddListener(this)
        this.HealthScale = (this.HealthInterface.GetMaxHealth() / 6) + 1
        this.OnHealthChange()
        this.UpdateMaxWidth = healthInterface.GetMaxHealth() * 16
        this.UpdateMaxHeight = 16
        this.XArray = make([]int, 30)
        this.Update()
    return this
}


func (this *HealthHudWidget) Update() {
        if this.XArray != nil {
            for index := int(0); index < len(this.XArray); index++ {
                this.XArray[index] = this.GetX() + (index * 16)
            }
        }
}


func (this *HealthHudWidget) SetX(x int) {
        this.BasicHud.SetX(x)
        this.Update()
}


func (this *HealthHudWidget) OnHealthChange() {
        this.Max = (this.HealthInterface.GetHealth() / this.HealthScale)
        this.TimeDelayHelper = NoTimeDelayHelper_SINGLETON
        if this.Max <= 1 && this.HealthInterface.IsAlive() {
            this.Max = 1
            if this.HealthScale - this.HealthInterface.GetHealth() > (this.HealthScale * 2) / 3 {
                this.TimeDelayHelper = this.SlowBeatTimeDelayHelper
            } else if this.HealthScale - this.HealthInterface.GetHealth() > this.HealthScale / 3 {
                this.TimeDelayHelper = this.MediumBeatTimeDelayHelper
            } else {
                this.TimeDelayHelper = this.FastBeatTimeDelayHelper
            }
        }
}


func (this *HealthHudWidget) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        for index := int(0); index < this.Max; index++ {
            if this.TimeDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
                this.AnimationInterface.PaintXY(graphics, this.XArray[index], this.GetY())
            }
        }
}


func (this *HealthHudWidget) PaintThreed(graphics IGraphics) {
}

