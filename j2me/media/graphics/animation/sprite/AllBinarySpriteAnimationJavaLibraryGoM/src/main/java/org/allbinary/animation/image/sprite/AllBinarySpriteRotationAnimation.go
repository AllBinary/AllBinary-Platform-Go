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

type AllBinarySpriteRotationAnimation struct {
    HackRotationSpriteIndexedAnimation
}

// The Java class AllBinarySpriteRotationAnimation has subclasses so its variables are this interface
type IAllBinarySpriteRotationAnimation interface {
    IHackRotationSpriteIndexedAnimation
    AllBinarySpriteRotationAnimation_() *AllBinarySpriteRotationAnimation
    NextRotation()
    PreviousRotation()
    SetFrameByDirection(direction *Direction)
    SetFrameToAngle(angle IAngle)
    SetFrame(index int)
    AdjustFrameToAngle(angle IAngle)
    AdjustFrame(angle int16)
}

func (this *AllBinarySpriteRotationAnimation) AllBinarySpriteRotationAnimation_() *AllBinarySpriteRotationAnimation {
    return this
}



func NewAllBinarySpriteRotationAnimation(sprite *Sprite, image IImage, animationBehavior IAnimationBehavior) *AllBinarySpriteRotationAnimation {
    this := &AllBinarySpriteRotationAnimation{}
        this.HackRotationSpriteIndexedAnimation = *NewHackRotationSpriteIndexedAnimation(sprite, image, AngleInfo_GetInstance(int16((int(AngleFactory_GetInstance().TOTAL_ANGLE) / sprite.GetRawFrameCount()))), animationBehavior)
        this.AngleInfo.AdjustAngle(this.Sprite.GetFrame())
    return this
}


func (this *AllBinarySpriteRotationAnimation) NextRotation() {
        this.Sprite.NextFrame()
        this.AngleInfo.AdjustAngle(this.Sprite.GetFrame())
}


func (this *AllBinarySpriteRotationAnimation) PreviousRotation() {
        this.Sprite.PrevFrame()
        this.AngleInfo.AdjustAngle(this.Sprite.GetFrame())
}


func (this *AllBinarySpriteRotationAnimation) SetFrameByDirection(direction *Direction) {
        var angle IAngle = this.DirectionUtil.GetFrameAngle(direction); _ = angle
        this.AdjustFrameToAngle(angle)
}


func (this *AllBinarySpriteRotationAnimation) SetFrameToAngle(angle IAngle) {
        this.AdjustFrameToAngle(angle)
}


func (this *AllBinarySpriteRotationAnimation) SetFrame(index int) {
        this.Sprite.SetFrame(index)
        this.AngleInfo.AdjustAngle(this.GetFrame())
}


func (this *AllBinarySpriteRotationAnimation) AdjustFrameToAngle(angle IAngle) {
        this.AdjustFrame(angle.GetValue())
}


func (this *AllBinarySpriteRotationAnimation) AdjustFrame(angle int16) {
        this.SetFrame(this.FrameUtil.GetFrameForAngle(angle, int(this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())))
}

