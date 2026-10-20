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

type ImageSegmentAnimation struct {
    ImageAnimation
    StartX int
    StartY int
    DrawWidth int
    DrawHeight int
    CircularIndexUtil *CircularIndexUtil
    SequenceArray []int
}

// The Java class ImageSegmentAnimation has subclasses so its variables are this interface
type IImageSegmentAnimation interface {
    IImageAnimation
    ImageSegmentAnimation_() *ImageSegmentAnimation
    NextFrame()
    PreviousFrame()
    SetFrame(index int)
    GetFrame() int
    GetSize() int
    SetSequence(sequenceArray []int)
    GetSequence() []int
    PaintXY(graphics IGraphics, x int, y int)
    GetStartX() int
    SetStartX(startX int)
    GetStartY() int
    SetStartY(startY int)
    GetDrawWidth() int
    SetDrawWidth(drawWidth int)
    GetDrawHeight() int
    SetDrawHeight(drawHeight int)
}

func (this *ImageSegmentAnimation) ImageSegmentAnimation_() *ImageSegmentAnimation {
    return this
}


func (this *ImageSegmentAnimation) initFields() {
    
    
    
    
    
    
}


func NewImageSegmentAnimation(image IImage, sequenceArray []int, animationBehavior IAnimationBehavior) *ImageSegmentAnimation {
    this := &ImageSegmentAnimation{}
        this.ImageAnimation = *NewImageAnimation(image, animationBehavior)
        this.initFields()
        this.StartX = 0
        this.StartY = 0
        this.DrawWidth = this.GetImage().GetWidth()
        this.DrawHeight = this.GetImage().GetHeight()
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(this.DrawHeight)
        this.SequenceArray = sequenceArray
    return this
}


func (this *ImageSegmentAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
}


func (this *ImageSegmentAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
}


func (this *ImageSegmentAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
}


func (this *ImageSegmentAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *ImageSegmentAnimation) GetSize() int {
        return this.CircularIndexUtil.GetSize()
}


func (this *ImageSegmentAnimation) SetSequence(sequenceArray []int) {
        this.SequenceArray = sequenceArray
}


func (this *ImageSegmentAnimation) GetSequence() []int {
        return this.SequenceArray
}


func (this *ImageSegmentAnimation) PaintXY(graphics IGraphics, x int, y int) {
}


func (this *ImageSegmentAnimation) GetStartX() int {
        return this.StartX
}


func (this *ImageSegmentAnimation) SetStartX(startX int) {
        this.StartX = startX
}


func (this *ImageSegmentAnimation) GetStartY() int {
        return this.StartY
}


func (this *ImageSegmentAnimation) SetStartY(startY int) {
        this.StartY = startY
}


func (this *ImageSegmentAnimation) GetDrawWidth() int {
        return this.DrawWidth
}


func (this *ImageSegmentAnimation) SetDrawWidth(drawWidth int) {
        this.DrawWidth = drawWidth
}


func (this *ImageSegmentAnimation) GetDrawHeight() int {
        return this.DrawHeight
}


func (this *ImageSegmentAnimation) SetDrawHeight(drawHeight int) {
        this.DrawHeight = drawHeight
}

