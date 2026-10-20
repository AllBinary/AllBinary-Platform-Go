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

type ImageBaseRotationAnimation struct {
    RotationAnimation
    Image IImage
    Anchor int
}

// The Java class ImageBaseRotationAnimation has subclasses so its variables are this interface
type IImageBaseRotationAnimation interface {
    IRotationAnimation
    ImageBaseRotationAnimation_() *ImageBaseRotationAnimation
    GetAnimationSize() int
    SetSequence(sequence []int)
    GetSequence() []int
    PaintXY(graphics IGraphics, x int, y int)
    GetImage() IImage
    GetWidth() int
    GetHeight() int
    Close()
    Finalize()
}

func (this *ImageBaseRotationAnimation) ImageBaseRotationAnimation_() *ImageBaseRotationAnimation {
    return this
}


func (this *ImageBaseRotationAnimation) initFields() {
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewImageBaseRotationAnimation(image IImage, angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) *ImageBaseRotationAnimation {
    this := &ImageBaseRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(angleInfo, CircularIndexUtil_CreateInstance(int(totalAngle) / int(angleInfo.GetAngleIncrementInfo().GetAngleIncrement())), animationBehavior)
        this.initFields()
        this.Image = image
    return this
}


// throws Exception
func (this *ImageBaseRotationAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *ImageBaseRotationAnimation) SetSequence(sequence []int) {
}


func (this *ImageBaseRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *ImageBaseRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.Image, x, y, this.Anchor)
}


func (this *ImageBaseRotationAnimation) GetImage() IImage {
        return this.Image
}


func (this *ImageBaseRotationAnimation) GetWidth() int {
        return this.Image.GetWidth()
}


func (this *ImageBaseRotationAnimation) GetHeight() int {
        return this.Image.GetHeight()
}


// throws Exception
func (this *ImageBaseRotationAnimation) Close() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}


// throws Throwable
func (this *ImageBaseRotationAnimation) Finalize() {
        DisposalUtil_GetInstance().DisposeImage(this.Image)
}

