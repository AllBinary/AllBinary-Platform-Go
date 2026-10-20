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

type VectorRotationAnimation struct {
    VectorBaseRotationAnimation
}



func NewVectorRotationAnimation(currentPoints [][][]int, basicColor *BasicColor, animationBehavior IAnimationBehavior) *VectorRotationAnimation {
    this := &VectorRotationAnimation{}
        this.VectorBaseRotationAnimation = *NewVectorBaseRotationAnimation(AngleInfo_GetInstance(int16((int(AngleFactory_GetInstance().TOTAL_ANGLE) / len(currentPoints)))), currentPoints, basicColor, animationBehavior)
        this.AngleInfo.AdjustAngle(this.GetFrame())
    return this
}


func (this *VectorRotationAnimation) SetFrameByDirection(direction *Direction) {
        var angle IAngle = this.DirectionUtil.GetFrameAngle(direction); _ = angle
        this.AdjustFrameToAngle(angle)
}


func (this *VectorRotationAnimation) SetFrameToAngle(angle IAngle) {
        this.AdjustFrameToAngle(angle)
}


func (this *VectorRotationAnimation) SetFrame(index int) {
        this.VectorBaseRotationAnimation.SetFrame(index)
        this.AngleInfo.AdjustAngle(this.GetFrame())
}


func (this *VectorRotationAnimation) AdjustFrame(newAngle int16) {
        this.SetFrame(this.FrameUtil.GetFrameForAngle(newAngle, int(this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())))
}


func (this *VectorRotationAnimation) AdjustFrameToAngle(newAngle IAngle) {
        this.AdjustFrame(newAngle.GetValue())
}

