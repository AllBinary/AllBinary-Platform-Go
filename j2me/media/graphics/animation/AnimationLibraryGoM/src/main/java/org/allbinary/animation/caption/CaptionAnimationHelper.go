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

type CaptionAnimationHelper struct {
    CaptionAnimationHelperBase
    TimeHelper ITimeDelayHelper
    TextAnimation ITextAnimation
    CaptionAnimation IAnimation
    Animation IAnimation
}


func (this *CaptionAnimationHelper) initFields() {
    this.TimeHelper = NewTimeDelayHelper(620)
    this.TextAnimation = NewTextAnimation(StringUtil_GetInstance().EMPTY_STRING, AnimationBehavior_GetInstance())
    
    this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
}


func NewCaptionAnimationHelper(captionAnimation IAnimation, captionDx int, captionDy int, dx int, dy int) *CaptionAnimationHelper {
    this := &CaptionAnimationHelper{}
    this.CaptionAnimationHelperBase = *NewCaptionAnimationHelperBase()
    this.initFields()
        this.CaptionAnimation = NewCaptionAnimation(captionAnimation, this.TextAnimation, captionDx, captionDy, dx, dy)
    return this
}


func (this *CaptionAnimationHelper) IsShowing() bool {
        if this.Animation == this.CaptionAnimation {
            return true
        } else {
            return false
        }
}


func (this *CaptionAnimationHelper) Tick() {
        if this.TimeHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
        }
}


func (this *CaptionAnimationHelper) Update(message string, basicColor *BasicColor) {
        this.TextAnimation.SetBasicColorP(basicColor)
        this.TextAnimation.SetText(message)
        this.Animation = this.CaptionAnimation
        this.TimeHelper.SetStartTimeTNT()
}


func (this *CaptionAnimationHelper) PaintXY(graphics IGraphics, x int, y int) {
        this.Animation.PaintXY(graphics, x, y)
}

