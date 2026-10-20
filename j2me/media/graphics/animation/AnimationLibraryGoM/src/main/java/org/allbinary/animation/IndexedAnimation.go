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

type IndexedAnimation struct {
    Animation
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AnimationBehaviorP IAnimationBehavior
}

// Implements IndexedAnimationInterface

// The Java class IndexedAnimation has subclasses so its variables are this interface
type IIndexedAnimation interface {
    IAnimation
    IndexedAnimation_() *IndexedAnimation
    GetAnimationBehavior() IAnimationBehavior
    Reset()
    SetFrame(index int)
    GetFrame() int
    GetAnimationSize() int
    GetSize() int
    PreviousFrame()
    IsLastFrame() bool
    SetSequence(sequence []int)
    GetSequence() []int
    SetState(indexedAnimation IIndexedAnimation)
    GetWidth() int
    GetHeight() int
}

func (this *IndexedAnimation) IndexedAnimation_() *IndexedAnimation {
    return this
}


func (this *IndexedAnimation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewIndexedAnimation(animationBehavior IAnimationBehavior) *IndexedAnimation {
    this := &IndexedAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.AnimationBehaviorP = animationBehavior
    return this
}


func (this *IndexedAnimation) GetAnimationBehavior() IAnimationBehavior {
        return this.AnimationBehaviorP
}


func (this *IndexedAnimation) Reset() {
        this.AnimationBehaviorP.Reset()
        this.SetFrame(0)
}


func (this *IndexedAnimation) SetFrame(index int) {
}


func (this *IndexedAnimation) GetFrame() int {
        return 0
}


// throws Exception
func (this *IndexedAnimation) GetAnimationSize() int {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *IndexedAnimation) GetSize() int {
        return 0
}


func (this *IndexedAnimation) PreviousFrame() {
}


func (this *IndexedAnimation) IsLastFrame() bool {
        if this.GetFrame() == this.GetSize() - 1 {
            return true
        } else {
            return false
        }
}


func (this *IndexedAnimation) SetSequence(sequence []int) {
}


func (this *IndexedAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *IndexedAnimation) SetState(indexedAnimation IIndexedAnimation) {
        this.SetFrame(indexedAnimation.GetFrame())
        this.SetDx(indexedAnimation.GetDx())
        this.SetDy(indexedAnimation.GetDy())
}


func (this *IndexedAnimation) GetWidth() int {
        return 0
}


func (this *IndexedAnimation) GetHeight() int {
        return 0
}

