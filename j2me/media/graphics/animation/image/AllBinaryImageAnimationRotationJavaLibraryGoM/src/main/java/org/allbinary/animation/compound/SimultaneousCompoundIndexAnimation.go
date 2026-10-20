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

type SimultaneousCompoundIndexAnimation struct {
    IndexedAnimation
    AnimationInterfaceArray []IIndexedAnimation
}


func (this *SimultaneousCompoundIndexAnimation) initFields() {
    
}


func NewSimultaneousCompoundIndexAnimation(animationInterfaceArray []IIndexedAnimation, animationBehavior IAnimationBehavior) *SimultaneousCompoundIndexAnimation {
    this := &SimultaneousCompoundIndexAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
    return this
}


func (this *SimultaneousCompoundIndexAnimation) SetFrame(frameIndex int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].SetFrame(frameIndex)
        }
}


func (this *SimultaneousCompoundIndexAnimation) GetFrame() int {
        return this.AnimationInterfaceArray[0].GetFrame()
}


// throws Exception
func (this *SimultaneousCompoundIndexAnimation) GetAnimationSize() int {
        return this.AnimationInterfaceArray[0].GetAnimationSize()
}


func (this *SimultaneousCompoundIndexAnimation) GetSize() int {
        return this.AnimationInterfaceArray[0].GetSize()
}


func (this *SimultaneousCompoundIndexAnimation) PreviousFrame() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PreviousFrame()
        }
}


// throws Exception
func (this *SimultaneousCompoundIndexAnimation) NextFrame() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].NextFrame()
        }
}


func (this *SimultaneousCompoundIndexAnimation) SetSequence(sequence []int) {
}


func (this *SimultaneousCompoundIndexAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *SimultaneousCompoundIndexAnimation) PaintXY(graphics IGraphics, x int, y int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PaintXY(graphics, x, y)
        }
}


func (this *SimultaneousCompoundIndexAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].PaintThreedXYZ(graphics, x, y, z)
        }
}


// throws Exception
func (this *SimultaneousCompoundIndexAnimation) Set(gl GL) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].Set(gl)
        }
}


func (this *SimultaneousCompoundIndexAnimation) GetAnimationInterfaceArray() []IIndexedAnimation {
        return this.AnimationInterfaceArray
}


func (this *SimultaneousCompoundIndexAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IIndexedAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}

