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

type BottomToTopImageAnimation struct {
    ImageSegmentAnimation
    StartHeight int
    CurrentHeight int
    Anchor int
}


func (this *BottomToTopImageAnimation) initFields() {
    
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewBottomToTopImageAnimation(image IImage, animationBehavior IAnimationBehavior) *BottomToTopImageAnimation {
    this := &BottomToTopImageAnimation{}
        this.ImageSegmentAnimation = *NewImageSegmentAnimation(image, PrimitiveIntUtil_GetArrayInstance(), animationBehavior)
        this.initFields()
        this.SetFrame(this.GetSize() - 1)
    return this
}


func (this *BottomToTopImageAnimation) NextFrame() {
        this.ImageSegmentAnimation.PreviousFrame()
        this.Update()
}


func (this *BottomToTopImageAnimation) PreviousFrame() {
        this.ImageSegmentAnimation.NextFrame()
        this.Update()
}


func (this *BottomToTopImageAnimation) SetFrame(index int) {
        this.ImageSegmentAnimation.SetFrame(index)
        this.Update()
}


func (this *BottomToTopImageAnimation) Update() {
        this.StartHeight = this.GetDrawHeight() * this.GetFrame() / this.GetSize()
        this.CurrentHeight = this.GetDrawHeight() - this.StartHeight
}


func (this *BottomToTopImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawRegion(this.GetImage(), this.GetStartX(), this.StartHeight, this.GetDrawWidth(), this.CurrentHeight, Sprite_TRANS_NONE, x, y + this.StartHeight, this.Anchor)
}

