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

type SimultaneousCompoundRotationAnimation struct {
    RotationAnimation
    AnimationInterfaceArray []IRotationAnimation
}


func (this *SimultaneousCompoundRotationAnimation) initFields() {
    
}


func NewSimultaneousCompoundRotationAnimation(animationInterfaceArray []IRotationAnimation, animationBehavior IAnimationBehavior) *SimultaneousCompoundRotationAnimation {
    this := &SimultaneousCompoundRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
    return this
}


func (this *SimultaneousCompoundRotationAnimation) SetFrame(frameIndex int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].SetFrame(frameIndex)
        }
}


func (this *SimultaneousCompoundRotationAnimation) GetFrame() int {
        return this.AnimationInterfaceArray[0].GetFrame()
}


// throws Exception
func (this *SimultaneousCompoundRotationAnimation) GetAnimationSize() int {
        return this.AnimationInterfaceArray[0].GetAnimationSize()
}


func (this *SimultaneousCompoundRotationAnimation) GetSize() int {
        return this.AnimationInterfaceArray[0].GetSize()
}


func (this *SimultaneousCompoundRotationAnimation) PreviousFrame() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PreviousFrame()
        }
}


// throws Exception
func (this *SimultaneousCompoundRotationAnimation) NextFrame() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].NextFrame()
        }
}


func (this *SimultaneousCompoundRotationAnimation) PreviousRotation() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PreviousRotation()
        }
}


func (this *SimultaneousCompoundRotationAnimation) NextRotation() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].NextRotation()
        }
}


func (this *SimultaneousCompoundRotationAnimation) SetSequence(sequence []int) {
}


func (this *SimultaneousCompoundRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *SimultaneousCompoundRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PaintXY(graphics, x, y)
        }
}


func (this *SimultaneousCompoundRotationAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PaintThreedXYZ(graphics, x, y, z)
        }
}


func (this *SimultaneousCompoundRotationAnimation) SetFrameByDirection(direction *Direction) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].SetFrameByDirection(direction)
        }
}


func (this *SimultaneousCompoundRotationAnimation) SetFrameToAngle(angle IAngle) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].SetFrameToAngle(angle)
        }
}


func (this *SimultaneousCompoundRotationAnimation) AdjustFrameToAngle(newAngle IAngle) {
        this.AdjustFrame(newAngle.GetValue())
}


func (this *SimultaneousCompoundRotationAnimation) AdjustFrame(angle int16) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].AdjustFrame(angle)
        }
}


func (this *SimultaneousCompoundRotationAnimation) GetAngleInfoP() *AngleInfo {
        return this.AnimationInterfaceArray[0].GetAngleInfoP()
}


func (this *SimultaneousCompoundRotationAnimation) GetAnimationInterfaceArray() []IRotationAnimation {
        return this.AnimationInterfaceArray
}


func (this *SimultaneousCompoundRotationAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IRotationAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}


// throws Exception
func (this *SimultaneousCompoundRotationAnimation) Set(gl GL) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].Set(gl)
        }
}

