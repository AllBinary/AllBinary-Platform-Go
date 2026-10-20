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

type SpecialAnimation struct {
    IndexedAnimation
}

// Implements SpecialAnimationInterface

// The Java class SpecialAnimation has subclasses so its variables are this interface
type ISpecialAnimation interface {
    IIndexedAnimation
    SpecialAnimation_() *SpecialAnimation
    NextFrame()
    SetLastFrame()
    IsComplete() bool
    Open()
    Close()
    Process()
}

func (this *SpecialAnimation) SpecialAnimation_() *SpecialAnimation {
    return this
}

var SpecialAnimation_NULL_SPECIAL_ANIMATION ISpecialAnimation = NewSpecialAnimation(NewIndexedAnimationBehavior(1, 250))


func SpecialAnimation_GetInstance() ISpecialAnimation {
        return SpecialAnimation_NULL_SPECIAL_ANIMATION
}


func NewSpecialAnimation(animationBehavior IAnimationBehavior) *SpecialAnimation {
    this := &SpecialAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
    return this
}


func (this *SpecialAnimation) NextFrame() {
        var indexedAnimationBehavior *IndexedAnimationBehavior = (any(this.AnimationBehaviorP).(*IndexedAnimationBehavior)); _ = indexedAnimationBehavior
        indexedAnimationBehavior.LoopIndex++
}


func (this *SpecialAnimation) SetLastFrame() {
}


func (this *SpecialAnimation) IsComplete() bool {
        return false
}


func (this *SpecialAnimation) Open() {
}


func (this *SpecialAnimation) Close() {
}


func (this *SpecialAnimation) Process() {
}

