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

type ImageArraySpecialAnimation struct {
    ImageArrayAnimation
    TimeDelayHelper ITimeDelayHelper
    LoopCount int
    LoopCountTotal int
}

// Implements SpecialAnimationInterface


func (this *ImageArraySpecialAnimation) initFields() {
    this.TimeDelayHelper = NewTimeDelayHelper(270)
    this.LoopCount = 0
    this.LoopCountTotal = 1
}


func NewImageArraySpecialAnimation(imageArray []IImage, animationBehavior IAnimationBehavior) *ImageArraySpecialAnimation {
    this := &ImageArraySpecialAnimation{}
        this.ImageArrayAnimation = *NewImageArrayAnimation(imageArray, animationBehavior)
        this.initFields()
        this.Reset()
    return this
}


func (this *ImageArraySpecialAnimation) GetLoopCount() int {
        return this.LoopCount
}


func (this *ImageArraySpecialAnimation) SetLastFrame() {
        this.CircularIndexUtil.SetIndex(this.CircularIndexUtil.GetSize() - 1)
}


func (this *ImageArraySpecialAnimation) Reset() {
        this.SetFrame(this.GetSize() - 1)
        this.LoopCount = 0
}


func (this *ImageArraySpecialAnimation) NextFrame() {
        if this.GetFrame() > 0 {
            if this.TimeDelayHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
                this.PreviousFrame()
                if this.GetFrame() == 0 {
                    this.LoopCount++
                }
            }
        }
}


func (this *ImageArraySpecialAnimation) IsComplete() bool {
        if this.LoopCount < this.LoopCountTotal || this.GetFrame() != 0 {
            return false
        } else {
            return true
        }
}


func (this *ImageArraySpecialAnimation) PaintXY(graphics IGraphics, x int, y int) {
        x = DisplayInfoSingleton_GetInstance().GetLastHalfWidth() - (this.GetImage(this.CircularIndexUtil.GetIndex()).GetWidth() >> 1)
        y = 5
        this.ImageArrayAnimation.PaintXY(graphics, x, y)
}

