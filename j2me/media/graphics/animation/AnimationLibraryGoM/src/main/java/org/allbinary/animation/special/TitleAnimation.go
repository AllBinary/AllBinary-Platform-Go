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

type TitleAnimation struct {
    SpecialAnimation
    SystemWrapper *ABSystemWrapper
    DeltaX float32
    DeltaY float32
    DeltaZ float32
    BasicColorArray []*BasicColor
    DxArray []int
    DyArray []int
    WidthP int
    AnimationInterfaceArray []IIndexedAnimation
    SizeP int
    Y int
    LastFrameStartTime int64
    DisplayInfoSingleton *DisplayInfoSingleton
    CLEAR_COLOR *BasicColor
}


func (this *TitleAnimation) initFields() {
    this.SystemWrapper = ABSystemWrapper_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.CLEAR_COLOR = BasicColorFactory_GetInstance().CLEAR_COLOR
}


func TitleAnimation_CreateAnimation(animationInterfaceArray []IIndexedAnimation, basicColorArray []*BasicColor, dxArray []int, dyArray []int, y int, width int) *TitleAnimation {
        return NewTitleAnimation(animationInterfaceArray, basicColorArray, dxArray, dyArray, y, width, NewIndexedAnimationBehavior(1, 250))
}


func NewTitleAnimation(animationInterfaceArray []IIndexedAnimation, basicColorArray []*BasicColor, dxArray []int, dyArray []int, y int, width int, animationBehavior IAnimationBehavior) *TitleAnimation {
    this := &TitleAnimation{}
        this.SpecialAnimation = *NewSpecialAnimation(animationBehavior)
        this.initFields()
        this.LastFrameStartTime = this.SystemWrapper.CurrentTimeMillis()
        this.AnimationInterfaceArray = animationInterfaceArray
        this.SizeP = len(this.AnimationInterfaceArray)
        this.BasicColorArray = basicColorArray
        this.DxArray = dxArray
        this.DyArray = dyArray
        this.Y = y
        this.WidthP = width
        this.Reset()
    return this
}


func (this *TitleAnimation) NextFrame() {
        var currentTime int64 = this.SystemWrapper.CurrentTimeMillis(); _ = currentTime
        var totalTimeElapsed int64 = currentTime - this.LastFrameStartTime; _ = totalTimeElapsed
        var indexedAnimationBehavior *IndexedAnimationBehavior = any(this.GetAnimationBehavior()).(*IndexedAnimationBehavior); _ = indexedAnimationBehavior
        if totalTimeElapsed > indexedAnimationBehavior.FrameDelayTime {
            this.PreviousFrame()
            this.LastFrameStartTime = currentTime
        }
        if this.AnimationInterfaceArray[0].GetFrame() == 0 {
            indexedAnimationBehavior.LoopIndex++
        }
}


func (this *TitleAnimation) IsComplete() bool {
        var indexedAnimationBehavior *IndexedAnimationBehavior = any(this.GetAnimationBehavior()).(*IndexedAnimationBehavior); _ = indexedAnimationBehavior
        if indexedAnimationBehavior.LoopTotal == -1 || indexedAnimationBehavior.LoopIndex < indexedAnimationBehavior.LoopTotal || this.GetFrame() != 0 {
            return false
        } else {
            return true
        }
}


func (this *TitleAnimation) SetSequence(sequence []int) {
}


func (this *TitleAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *TitleAnimation) GetSize() int {
        return this.AnimationInterfaceArray[0].GetSize()
}


func (this *TitleAnimation) GetFrame() int {
        return this.AnimationInterfaceArray[0].GetFrame()
}


func (this *TitleAnimation) SetFrame(frame int) {
        for index := int(0); index < this.SizeP; index++ {
            this.AnimationInterfaceArray[index].SetFrame(frame)
        }
}


func (this *TitleAnimation) SetLastFrame() {
        this.SetFrame(this.GetSize() - 1)
}


func (this *TitleAnimation) Reset() {
        this.SetLastFrame()
        var indexedAnimationBehavior *IndexedAnimationBehavior = (any(this.GetAnimationBehavior()).(*IndexedAnimationBehavior)); _ = indexedAnimationBehavior
        indexedAnimationBehavior.Reset()
}


func (this *TitleAnimation) PreviousFrame() {
        for index := int(0); index < this.SizeP; index++ {
            this.AnimationInterfaceArray[index].PreviousFrame()
        }
}


func (this *TitleAnimation) PaintFrameXY(graphics IGraphics, frame int, x int, y int) {
        this.SetFrame(frame)
        this.PaintXY(graphics, x, y)
}


func (this *TitleAnimation) PaintXY(graphics IGraphics, ax int, ay int) {
        var x int = 0; _ = x
        if this.WidthP != Integer_MIN_VALUE {
            x = ((this.DisplayInfoSingleton.GetLastWidth() - this.WidthP) / 2)
        }
        var deltaX int; _ = deltaX
        var deltaY int; _ = deltaY
        for index := int(0); index < this.SizeP; index++ {
            deltaX = this.DxArray[index] + x
            deltaY = this.DyArray[index] + this.Y
            if this.BasicColorArray[index] != this.CLEAR_COLOR {
                this.BasicSetColorUtil.SetBasicColorP(graphics, this.BasicColorArray[index])
            }
            this.AnimationInterfaceArray[index].PaintXY(graphics, deltaX, deltaY)
        }
}

