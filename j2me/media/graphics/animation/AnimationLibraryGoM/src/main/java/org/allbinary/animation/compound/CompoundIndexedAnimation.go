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

type CompoundIndexedAnimation struct {
    IndexedAnimation
    CircularIndexUtil *CircularIndexUtil
    AnimationInterfaceArray []IAnimation
}

// Implements CompoundAnimationInterface


func (this *CompoundIndexedAnimation) initFields() {
    
    
}


func NewCompoundIndexedAnimation(animationInterfaceArray []IAnimation, animationBehavior IAnimationBehavior) *CompoundIndexedAnimation {
    this := &CompoundIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(this.AnimationInterfaceArray))
    return this
}


func (this *CompoundIndexedAnimation) SetFrame(index int) {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]).(IIndexedAnimation); _ = indexedAnimation
        indexedAnimation.SetFrame(index)
}


func (this *CompoundIndexedAnimation) GetFrame() int {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]).(IIndexedAnimation); _ = indexedAnimation
        return indexedAnimation.GetFrame()
}


func (this *CompoundIndexedAnimation) GetSize() int {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]).(IIndexedAnimation); _ = indexedAnimation
        return indexedAnimation.GetSize()
}


func (this *CompoundIndexedAnimation) PreviousFrame() {
        var indexedAnimation IIndexedAnimation = any(this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]).(IIndexedAnimation); _ = indexedAnimation
        indexedAnimation.PreviousFrame()
}


func (this *CompoundIndexedAnimation) SetSequence(sequence []int) {
}


func (this *CompoundIndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


// throws Exception
func (this *CompoundIndexedAnimation) NextFrame() {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].NextFrame()
}


func (this *CompoundIndexedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintXY(graphics, x, y)
}


func (this *CompoundIndexedAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
        this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()].PaintThreedXYZ(graphics, x, y, z)
}


func (this *CompoundIndexedAnimation) GetCurrentAnimation() IIndexedAnimation {
        return any(this.AnimationInterfaceArray[this.CircularIndexUtil.GetIndex()]).(IIndexedAnimation)
}


func (this *CompoundIndexedAnimation) NextAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Next()
        this.SetFrame(frame)
}


func (this *CompoundIndexedAnimation) PreviousAnimation() {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.Previous()
        this.SetFrame(frame)
}


func (this *CompoundIndexedAnimation) SetAnimation(index int) {
        var frame int = this.GetFrame(); _ = frame
        this.CircularIndexUtil.SetIndex(index)
        this.SetFrame(frame)
}


func (this *CompoundIndexedAnimation) GetAnimationInterfaceArray() []IAnimation {
        return this.AnimationInterfaceArray
}


func (this *CompoundIndexedAnimation) SetAnimationInterfaceArray(animationInterfaceArray []IAnimation) {
        this.AnimationInterfaceArray = animationInterfaceArray
}


// throws Exception
func (this *CompoundIndexedAnimation) Set(gl GL) {
        for index := int(len(this.AnimationInterfaceArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationInterfaceArray[index].Set(gl)
        }
}

