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

type RotationAnimation struct {
    IndexedAnimation
    FrameUtil *FrameUtil
    DirectionUtil *DirectionUtil
    AngleInfo *AngleInfo
    CircularIndexUtil *CircularIndexUtil
}

// Implements RotationAnimationInterface

// The Java class RotationAnimation has subclasses so its variables are this interface
type IRotationAnimation interface {
    IIndexedAnimation
    RotationAnimation_() *RotationAnimation
    NextRotationX()
    PreviousRotationX()
    NextRotation()
    PreviousRotation()
    NextRotationZ()
    PreviousRotationZ()
    SetFrame(index int)
    SetFrameByDirection(direction *Direction)
    SetFrameToAngle(angle IAngle)
    AdjustFrameToAngle(angle IAngle)
    AdjustFrame(angle int16)
    GetFrame() int
    GetSize() int
    GetAngleInfoP() *AngleInfo
    ToString() string
}

func (this *RotationAnimation) RotationAnimation_() *RotationAnimation {
    return this
}


func (this *RotationAnimation) initFields() {
    this.FrameUtil = FrameUtil_GetInstance()
    this.DirectionUtil = DirectionUtil_GetInstance()
    
    
}


func NewRotationAnimation(angleInfo *AngleInfo, circularIndexUtil *CircularIndexUtil, animationBehavior IAnimationBehavior) *RotationAnimation {
    this := &RotationAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.AngleInfo = angleInfo
        this.CircularIndexUtil = circularIndexUtil
    return this
}


func (this *RotationAnimation) NextRotationX() {
}


func (this *RotationAnimation) PreviousRotationX() {
}


func (this *RotationAnimation) NextRotation() {
        this.AngleInfo.AdjustAngle(this.CircularIndexUtil.Next())
}


func (this *RotationAnimation) PreviousRotation() {
        this.AngleInfo.AdjustAngle(this.CircularIndexUtil.Previous())
}


func (this *RotationAnimation) NextRotationZ() {
}


func (this *RotationAnimation) PreviousRotationZ() {
}


func (this *RotationAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
        var newFrame int = this.CircularIndexUtil.GetIndex(); _ = newFrame
        this.AngleInfo.AdjustAngle(newFrame)
}


func (this *RotationAnimation) SetFrameByDirection(direction *Direction) {
        var angle IAngle = this.DirectionUtil.GetFrameAngle(direction); _ = angle
        this.AdjustFrameToAngle(angle)
}


func (this *RotationAnimation) SetFrameToAngle(angle IAngle) {
        this.AdjustFrameToAngle(angle)
}


func (this *RotationAnimation) AdjustFrameToAngle(angle IAngle) {
        this.AdjustFrame(angle.GetValue())
}


func (this *RotationAnimation) AdjustFrame(angle int16) {
        this.SetFrame(this.FrameUtil.GetFrameForAngle(angle, int(this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())))
}


func (this *RotationAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *RotationAnimation) GetSize() int {
        return this.CircularIndexUtil.GetSize()
}


func (this *RotationAnimation) GetAngleInfoP() *AngleInfo {
        return this.AngleInfo
}


func (this *RotationAnimation) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.IndexedAnimation.ToString()).Append("circularIndexUtil: ").Append(fmt.Sprint(this.CircularIndexUtil)).Append("angleInfo: ").Append(fmt.Sprint(this.AngleInfo)))
}

func (this *RotationAnimation) String() string {
    return this.ToString()
}

