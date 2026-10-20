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

type AutoCompoundRotationAnimation struct {
    RotationAnimation
    AnimationInterfaceArray []IRotationAnimation
}

// Implements CompoundAnimationInterface


func (this *AutoCompoundRotationAnimation) initFields() {
    
}


func NewAutoCompoundRotationAnimation(animationInterfaceArray []IRotationAnimation, animationBehavior IAnimationBehavior) *AutoCompoundRotationAnimation {
    this := &AutoCompoundRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(this.AnimationInterfaceArray))
    return this
}


func (this *AutoCompoundRotationAnimation) SetFrame(index int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrame(index)
}


func (this *AutoCompoundRotationAnimation) GetFrame() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetFrame()
}


// throws Exception
func (this *AutoCompoundRotationAnimation) GetAnimationSize() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetAnimationSize()
}


func (this *AutoCompoundRotationAnimation) GetSize() int {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetSize()
}


func (this *AutoCompoundRotationAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PreviousFrame()
}


// throws Exception
func (this *AutoCompoundRotationAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].NextFrame()
}


func (this *AutoCompoundRotationAnimation) PreviousRotation() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PreviousRotation()
}


func (this *AutoCompoundRotationAnimation) NextRotation() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].NextRotation()
}


func (this *AutoCompoundRotationAnimation) SetSequence(sequence []int) {
}


func (this *AutoCompoundRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *AutoCompoundRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintXY(graphics, x, y)
}


func (this *AutoCompoundRotationAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintThreedXYZ(graphics, x, y, z)
}


func (this *AutoCompoundRotationAnimation) GetCurrentAnimation() IIndexedAnimation {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]
}


func (this *AutoCompoundRotationAnimation) NextAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Next()
        this.SetFrame(frame)
}


func (this *AutoCompoundRotationAnimation) PreviousAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Previous()
        this.SetFrame(frame)
}


func (this *AutoCompoundRotationAnimation) SetAnimation(index int) {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.SetIndex(index)
        this.SetFrame(frame)
}


func (this *AutoCompoundRotationAnimation) SetFrameByDirection(direction *Direction) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrameByDirection(direction)
}


func (this *AutoCompoundRotationAnimation) SetFrameToAngle(angle IAngle) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].SetFrameToAngle(angle)
}


func (this *AutoCompoundRotationAnimation) AdjustFrameToAngle(newAngle IAngle) {
        this.AdjustFrame(newAngle.GetValue())
}


func (this *AutoCompoundRotationAnimation) AdjustFrame(angle int16) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].AdjustFrame(angle)
}


func (this *AutoCompoundRotationAnimation) GetAngleInfoP() *AngleInfo {
        return this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].GetAngleInfoP()
}


func (this *AutoCompoundRotationAnimation) GetAnimationInterfaceArray() []IRotationAnimation {
        return this.AnimationInterfaceArray
}


func (this *AutoCompoundRotationAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IRotationAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}


// throws Exception
func (this *AutoCompoundRotationAnimation) Set(gl GL) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].Set(gl)
        }
}

