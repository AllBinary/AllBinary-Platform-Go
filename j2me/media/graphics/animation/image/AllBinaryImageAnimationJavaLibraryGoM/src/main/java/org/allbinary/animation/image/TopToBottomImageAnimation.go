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

type TopToBottomImageAnimation struct {
    ImageSegmentAnimation
    CurrentHeight int
    Anchor int
}


func (this *TopToBottomImageAnimation) initFields() {
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewTopToBottomImageAnimation(image IImage, sequenceArray []int, animationBehavior IAnimationBehavior) *TopToBottomImageAnimation {
    this := &TopToBottomImageAnimation{}
        this.ImageSegmentAnimation = *NewImageSegmentAnimation(image, sequenceArray, animationBehavior)
        this.initFields()
        this.CurrentHeight = 0
        if this.GetSequence()[0] == -1 {
            this.SetFrame(this.GetSize() - 1)
        }
    return this
}


func (this *TopToBottomImageAnimation) NextFrame() {
        if this.GetSequence()[0] == -1 {
            this.ImageSegmentAnimation.PreviousFrame()
        } else {
            this.ImageSegmentAnimation.NextFrame()
        }
        this.Update()
}


func (this *TopToBottomImageAnimation) PreviousFrame() {
        if this.GetSequence()[0] == -1 {
            this.ImageSegmentAnimation.NextFrame()
        } else {
            this.ImageSegmentAnimation.PreviousFrame()
        }
        this.Update()
}


func (this *TopToBottomImageAnimation) SetFrame(index int) {
        this.ImageSegmentAnimation.SetFrame(index)
        this.Update()
}


func (this *TopToBottomImageAnimation) Update() {
        this.CurrentHeight = this.GetDrawHeight() * this.GetFrame() / this.GetSize()
}


func (this *TopToBottomImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawRegion(this.GetImage(), this.GetStartX(), this.GetStartY(), this.GetDrawWidth(), this.CurrentHeight, Sprite_TRANS_NONE, x, y, this.Anchor)
}

