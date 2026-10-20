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

type TweenedImageAnimation struct {
    IndexedAnimation
    Image IImage
    Dx int
    Dy int
    TotalFrames int
    CircularIndexUtil *CircularIndexUtil
    Anchor int
}


func (this *TweenedImageAnimation) initFields() {
    
    
    
    
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewTweenedImageAnimation(image IImage, dx int, dy int, totalFrames int, animationBehavior IAnimationBehavior) *TweenedImageAnimation {
    this := &TweenedImageAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.Image = image
        this.Dx = dx
        this.Dy = dy
        this.TotalFrames = totalFrames
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(this.TotalFrames)
    return this
}


// throws Exception
func (this *TweenedImageAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *TweenedImageAnimation) NextFrame() {
        this.CircularIndexUtil.Next()
}


func (this *TweenedImageAnimation) PreviousFrame() {
        this.CircularIndexUtil.Previous()
}


func (this *TweenedImageAnimation) SetFrame(index int) {
        this.CircularIndexUtil.SetIndex(index)
}


func (this *TweenedImageAnimation) GetFrame() int {
        return this.CircularIndexUtil.GetIndex()
}


func (this *TweenedImageAnimation) GetSize() int {
        return this.TotalFrames
}


func (this *TweenedImageAnimation) SetSequence(sequence []int) {
}


func (this *TweenedImageAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *TweenedImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        var frame int = this.GetFrame(); _ = frame
        var currentX int = this.Dx * frame; _ = currentX
        var currentY int = this.Dy * frame; _ = currentY
        graphics.DrawImage(this.Image, x + currentX, y + currentY, this.Anchor)
}


// throws Exception
func (this *TweenedImageAnimation) Close() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}


// throws Throwable
func (this *TweenedImageAnimation) Finalize() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}

