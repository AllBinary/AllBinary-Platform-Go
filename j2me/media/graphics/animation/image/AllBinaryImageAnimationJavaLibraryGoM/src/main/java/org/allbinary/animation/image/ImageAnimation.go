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

type ImageAnimation struct {
    IndexedAnimation
    Image IImage
    Anchor int
}

// The Java class ImageAnimation has subclasses so its variables are this interface
type IImageAnimation interface {
    IIndexedAnimation
    ImageAnimation_() *ImageAnimation
    GetAnimationSize() int
    NextFrame()
    PreviousFrame()
    SetFrame(index int)
    GetFrame() int
    GetSize() int
    SetSequence(sequence []int)
    GetSequence() []int
    PaintXY(graphics IGraphics, x int, y int)
    GetImage() IImage
    Close()
    Finalize()
}

func (this *ImageAnimation) ImageAnimation_() *ImageAnimation {
    return this
}


func (this *ImageAnimation) initFields() {
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewImageAnimation(image IImage, animationBehavior IAnimationBehavior) *ImageAnimation {
    this := &ImageAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.Image = image
    return this
}


// throws Exception
func (this *ImageAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *ImageAnimation) NextFrame() {
}


func (this *ImageAnimation) PreviousFrame() {
}


func (this *ImageAnimation) SetFrame(index int) {
}


func (this *ImageAnimation) GetFrame() int {
        return 0
}


func (this *ImageAnimation) GetSize() int {
        return 1
}


func (this *ImageAnimation) SetSequence(sequence []int) {
}


func (this *ImageAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *ImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.Image, x, y, this.Anchor)
}


func (this *ImageAnimation) GetImage() IImage {
        return this.Image
}


// throws Exception
func (this *ImageAnimation) Close() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}


// throws Throwable
func (this *ImageAnimation) Finalize() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}

