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

type CompoundRotationAnimation struct {
    RotationAnimation
    AnimationInterfaceArray []IRotationAnimation
}

// Implements CompoundAnimationInterface


func (this *CompoundRotationAnimation) initFields() {
    
}


func NewCompoundRotationAnimation(animationInterfaceArray []IRotationAnimation, animationBehavior IAnimationBehavior) *CompoundRotationAnimation {
    this := &CompoundRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(this.AnimationInterfaceArray))
    return this
}


func (this *CompoundRotationAnimation) SetFrame(index int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrame(index)
}


func (this *CompoundRotationAnimation) GetFrame() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetFrame()
}


// throws Exception
func (this *CompoundRotationAnimation) GetAnimationSize() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetAnimationSize()
}


func (this *CompoundRotationAnimation) GetSize() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetSize()
}


func (this *CompoundRotationAnimation) PreviousFrame() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PreviousFrame()
}


// throws Exception
func (this *CompoundRotationAnimation) NextFrame() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].NextFrame()
}


func (this *CompoundRotationAnimation) PreviousRotation() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PreviousRotation()
}


func (this *CompoundRotationAnimation) NextRotation() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].NextRotation()
}


func (this *CompoundRotationAnimation) SetSequence(sequence []int) {
}


func (this *CompoundRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *CompoundRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintXY(graphics, x, y)
}


func (this *CompoundRotationAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintThreedXYZ(graphics, x, y, z)
}


func (this *CompoundRotationAnimation) GetCurrentAnimation() IIndexedAnimation {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]
}


func (this *CompoundRotationAnimation) NextAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Next()
        this.SetFrame(frame)
}


func (this *CompoundRotationAnimation) PreviousAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Previous()
        this.SetFrame(frame)
}


func (this *CompoundRotationAnimation) SetAnimation(index int) {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.SetIndex(index)
        this.SetFrame(frame)
}


func (this *CompoundRotationAnimation) SetFrameByDirection(direction *Direction) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrameByDirection(direction)
}


func (this *CompoundRotationAnimation) SetFrameToAngle(angle IAngle) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrameToAngle(angle)
}


func (this *CompoundRotationAnimation) AdjustFrameToAngle(newAngle IAngle) {
        this.AdjustFrame(newAngle.GetValue())
}


func (this *CompoundRotationAnimation) AdjustFrame(angle int16) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].AdjustFrame(angle)
}


func (this *CompoundRotationAnimation) GetAngleInfoP() *AngleInfo {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetAngleInfoP()
}


func (this *CompoundRotationAnimation) GetAnimationInterfaceArray() []IRotationAnimation {
        return this.AnimationInterfaceArray
}


func (this *CompoundRotationAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IRotationAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}


// throws Exception
func (this *CompoundRotationAnimation) Set(gl GL) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].Set(gl)
        }
}

