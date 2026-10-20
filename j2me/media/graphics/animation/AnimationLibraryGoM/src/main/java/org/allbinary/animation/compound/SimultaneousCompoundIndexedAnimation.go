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

type SimultaneousCompoundIndexedAnimation struct {
    IndexedAnimation
    AnimationInterfaceArray []IAnimation
}


func (this *SimultaneousCompoundIndexedAnimation) initFields() {
    
}


func NewSimultaneousCompoundIndexedAnimation(animationInterfaceArray []IAnimation, animationBehavior IAnimationBehavior) *SimultaneousCompoundIndexedAnimation {
    this := &SimultaneousCompoundIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
    return this
}


func (this *SimultaneousCompoundIndexedAnimation) SetFrame(frameIndex int) {
        var indexedAnimation IIndexedAnimation; _ = indexedAnimation
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            indexedAnimation = any(this.AnimationInterfaceArray[index]).(IIndexedAnimation)
            indexedAnimation.SetFrame(frameIndex)
        }
}


func (this *SimultaneousCompoundIndexedAnimation) GetFrame() int {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[0]).(IIndexedAnimation); _ = indexedAnimation
        return indexedAnimation.GetFrame()
}


func (this *SimultaneousCompoundIndexedAnimation) GetSize() int {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[0]).(IIndexedAnimation); _ = indexedAnimation
        return indexedAnimation.GetSize()
}


func (this *SimultaneousCompoundIndexedAnimation) PreviousFrame() {
        var indexedAnimation IIndexedAnimation; _ = indexedAnimation
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            indexedAnimation = any(this.AnimationInterfaceArray[index]).(IIndexedAnimation)
            indexedAnimation.PreviousFrame()
        }
}


func (this *SimultaneousCompoundIndexedAnimation) SetSequence(sequence []int) {
}


func (this *SimultaneousCompoundIndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


// throws Exception
func (this *SimultaneousCompoundIndexedAnimation) NextFrame() {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].NextFrame()
        }
}


func (this *SimultaneousCompoundIndexedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        var size int = len(this.AnimationInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            this.AnimationInterfaceArray[index].PaintXY(graphics, x, y)
        }
}


func (this *SimultaneousCompoundIndexedAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        var size int = len(this.AnimationInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            this.AnimationInterfaceArray[index].PaintThreedXYZ(graphics, x, y, z)
        }
}


func (this *SimultaneousCompoundIndexedAnimation) GetAnimationInterfaceArray() []IAnimation {
        return this.AnimationInterfaceArray
}


func (this *SimultaneousCompoundIndexedAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}

