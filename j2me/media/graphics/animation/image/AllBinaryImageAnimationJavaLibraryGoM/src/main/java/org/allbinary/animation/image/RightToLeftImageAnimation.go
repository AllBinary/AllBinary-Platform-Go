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

type RightToLeftImageAnimation struct {
    ImageSegmentAnimation
    StartWidth int
    CurrentWidth int
    Anchor int
}

// The Java class RightToLeftImageAnimation has subclasses so its variables are this interface
type IRightToLeftImageAnimation interface {
    IImageSegmentAnimation
    RightToLeftImageAnimation_() *RightToLeftImageAnimation
    NextFrame()
    PreviousFrame()
    SetFrame(index int)
    Update()
    PaintXY(graphics IGraphics, x int, y int)
}

func (this *RightToLeftImageAnimation) RightToLeftImageAnimation_() *RightToLeftImageAnimation {
    return this
}


func (this *RightToLeftImageAnimation) initFields() {
    
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewRightToLeftImageAnimation(image IImage, animationBehavior IAnimationBehavior) *RightToLeftImageAnimation {
    this := &RightToLeftImageAnimation{}
        this.ImageSegmentAnimation = *NewImageSegmentAnimation(image, PrimitiveIntUtil_GetArrayInstance(), animationBehavior)
        this.initFields()
        this.SetFrame(this.GetSize() - 1)
    return this
}


func (this *RightToLeftImageAnimation) NextFrame() {
        this.ImageSegmentAnimation.PreviousFrame()
        this.Update()
}


func (this *RightToLeftImageAnimation) PreviousFrame() {
        this.ImageSegmentAnimation.NextFrame()
        this.Update()
}


func (this *RightToLeftImageAnimation) SetFrame(index int) {
        this.ImageSegmentAnimation.SetFrame(index)
        this.Update()
}


func (this *RightToLeftImageAnimation) Update() {
        this.StartWidth = this.GetDrawWidth() * this.GetFrame() / this.GetSize()
        this.CurrentWidth = this.GetDrawWidth() - this.StartWidth
}


func (this *RightToLeftImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawRegion(this.GetImage(), this.StartWidth, this.GetStartY(), this.CurrentWidth, this.GetDrawHeight(), Sprite_TRANS_NONE, x + this.StartWidth, y, this.Anchor)
}

