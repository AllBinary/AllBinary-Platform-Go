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

type ImageArrayBaseRotationAnimation struct {
    RotationAnimation
    ImageModifierUtil *ImageModifierUtil
    OriginalImageArray []IImage
    ImageArray []IImage
    CurrentImage IImage
    TotalFrames int
    Anchor int
}

// The Java class ImageArrayBaseRotationAnimation has subclasses so its variables are this interface
type IImageArrayBaseRotationAnimation interface {
    IRotationAnimation
    ImageArrayBaseRotationAnimation_() *ImageArrayBaseRotationAnimation
    GetAnimationSize() int
    SetAlpha(alpha int)
    NextRotation()
    PreviousRotation()
    SetFrame(index2 int)
    SetSequence(sequence []int)
    GetSequence() []int
    SetImageArray(imageArray []IImage)
    PaintXY(graphics IGraphics, x int, y int)
    Close()
    Finalize()
}

func (this *ImageArrayBaseRotationAnimation) ImageArrayBaseRotationAnimation_() *ImageArrayBaseRotationAnimation {
    return this
}


func (this *ImageArrayBaseRotationAnimation) initFields() {
    this.ImageModifierUtil = ImageModifierUtil_GetInstanceOrCreate()
    
    this.ImageArray = NullImage_NULL_IMAGE_ARRAY
    
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewImageArrayBaseRotationAnimation(originalImageArray []IImage, angleInfo *AngleInfo, animationBehavior IAnimationBehavior) *ImageArrayBaseRotationAnimation {
    this := &ImageArrayBaseRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(angleInfo, CircularIndexUtil_CreateInstance(360 / int(angleInfo.GetAngleIncrementInfo().GetAngleIncrement())), animationBehavior)
        this.initFields()
        this.OriginalImageArray = originalImageArray
        this.SetImageArray(this.ImageModifierUtil.GetImageArray(originalImageArray))
        this.CurrentImage = this.ImageArray[this.CircularIndexUtil.GetIndex()]
    return this
}


// throws Exception
func (this *ImageArrayBaseRotationAnimation) GetAnimationSize() int {
        return this.GetSize()
}


func (this *ImageArrayBaseRotationAnimation) SetAlpha(alpha int) {
        if this.AlphaP != alpha {
            this.RotationAnimation.SetAlpha(alpha)
            this.ImageModifierUtil.Reset()
            var index int = this.CircularIndexUtil.GetIndex(); _ = index
            this.ImageModifierUtil.SetAlpha(this.OriginalImageArray[index], this.ImageArray[index], index, this.AlphaP)
        }
}


func (this *ImageArrayBaseRotationAnimation) NextRotation() {
        this.RotationAnimation.NextRotation()
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.ImageModifierUtil.SetAlpha(this.OriginalImageArray[index], this.ImageArray[index], index, this.AlphaP)
        this.CurrentImage = this.ImageArray[index]
}


func (this *ImageArrayBaseRotationAnimation) PreviousRotation() {
        this.RotationAnimation.PreviousRotation()
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.ImageModifierUtil.SetAlpha(this.OriginalImageArray[index], this.ImageArray[index], index, this.AlphaP)
        this.CurrentImage = this.ImageArray[index]
}


func (this *ImageArrayBaseRotationAnimation) SetFrame(index2 int) {
        this.RotationAnimation.SetFrame(index2)
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.ImageModifierUtil.SetAlpha(this.OriginalImageArray[index], this.ImageArray[index], index, this.AlphaP)
        this.CurrentImage = this.ImageArray[index]
}


func (this *ImageArrayBaseRotationAnimation) SetSequence(sequence []int) {
}


func (this *ImageArrayBaseRotationAnimation) GetSequence() []int {
        return PrimitiveIntUtil_GetArrayInstance()
}


func (this *ImageArrayBaseRotationAnimation) SetImageArray(imageArray []IImage) {
        this.ImageArray = imageArray
        this.TotalFrames = len(imageArray)
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(this.TotalFrames)
}


func (this *ImageArrayBaseRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.CurrentImage, x, y, this.Anchor)
}


// throws Exception
func (this *ImageArrayBaseRotationAnimation) Close() {
        var disposalUtil *DisposalUtil = DisposalUtil_GetInstance(); _ = disposalUtil
        var size2 int = len(this.ImageArray); _ = size2
        for index := int(0); index < size2; index++ {
            disposalUtil.DisposeImage(this.ImageArray[index])
        }
        var size int = len(this.OriginalImageArray); _ = size
        for index := int(0); index < size; index++ {
            disposalUtil.DisposeImage(this.OriginalImageArray[index])
        }
        disposalUtil.DisposeImage(this.CurrentImage)
}


// throws Throwable
func (this *ImageArrayBaseRotationAnimation) Finalize() {
        var disposalUtil *DisposalUtil = DisposalUtil_GetInstance(); _ = disposalUtil
        var size2 int = len(this.ImageArray); _ = size2
        for index := int(0); index < size2; index++ {
            disposalUtil.DisposeImage(this.ImageArray[index])
        }
        var size int = len(this.OriginalImageArray); _ = size
        for index := int(0); index < size; index++ {
            disposalUtil.DisposeImage(this.OriginalImageArray[index])
        }
        disposalUtil.DisposeImage(this.CurrentImage)
}

