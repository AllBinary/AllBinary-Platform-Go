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

type AllBinaryJ2SEImageRotationAnimation struct {
    ImageBaseRotationAnimation
    ImageRotationUtil *ImageRotationUtil
    ImageModifierUtil *ImageModifierUtil
    RealOriginalImage IImage
    OriginalImageArray []IImage
    TwoImages []IImage
    ImageToShow IImage
    BufferedImageIndex int
    LastScaleX float32
    LastScaleY float32
    ScaleX float32
    ScaleY float32
    MaxScaleX float32
    MaxScaleY float32
    AlphaProcessor IModifierBaseProcessor
    SetColorProcessor IModifierBaseProcessor
    ChangeColorProcessor IModifierBaseProcessor
    ScaleProcessor IScaleBaseProcessor
}

// The Java class AllBinaryJ2SEImageRotationAnimation has subclasses so its variables are this interface
type IAllBinaryJ2SEImageRotationAnimation interface {
    IImageBaseRotationAnimation
    AllBinaryJ2SEImageRotationAnimation_() *AllBinaryJ2SEImageRotationAnimation
    SetBasicColorP(basicColor *BasicColor)
    ChangeBasicColor(basicColor *BasicColor)
    SetAlpha(alpha int)
    SetScale(scaleX float32, scaleY float32)
    SetMaxScale(maxScaleX float32, maxScaleY float32)
    NextRotation()
    PreviousRotation()
    UpdateImage()
    SetFrame(index int)
    Swap()
    PaintXY(graphics IGraphics, x int, y int)
    Close()
    Finalize()
}

func (this *AllBinaryJ2SEImageRotationAnimation) AllBinaryJ2SEImageRotationAnimation_() *AllBinaryJ2SEImageRotationAnimation {
    return this
}


func (this *AllBinaryJ2SEImageRotationAnimation) initFields() {
    this.ImageRotationUtil = ImageRotationUtil_GetInstance()
    this.ImageModifierUtil = ImageModifierUtil_GetInstanceOrCreate()
    
    this.OriginalImageArray = make([]IImage, 1)
    this.TwoImages = make([]IImage, 2)
    
    
    this.LastScaleX = 1.0
    this.LastScaleY = 1.0
    this.ScaleX = 1.0
    this.ScaleY = 1.0
    
    
    this.AlphaProcessor = ModifierBaseProcessor_GetInstance()
    this.SetColorProcessor = ModifierBaseProcessor_GetInstance()
    this.ChangeColorProcessor = ModifierBaseProcessor_GetInstance()
    this.ScaleProcessor = ScaleProcessor_GetInstance()
}


func NewAllBinaryJ2SEImageRotationAnimation(originalImage IImage, image IImage, angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) *AllBinaryJ2SEImageRotationAnimation {
    this := &AllBinaryJ2SEImageRotationAnimation{}
        this.ImageBaseRotationAnimation = *NewImageBaseRotationAnimation(image, angleInfo, totalAngle, animationBehavior)
        this.initFields()
        this.RealOriginalImage = originalImage
        this.OriginalImageArray[0] = originalImage
        this.ImageToShow = image
        this.TwoImages[0] = image
        this.TwoImages[1] = ImageCopyUtil_GetInstance().CreateImageForRotation(image)
    return this
}


func (this *AllBinaryJ2SEImageRotationAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.ImageBaseRotationAnimation.SetBasicColorP(basicColor)
        if changed {
            this.SetColorProcessor = SetColorProcessor_GetInstance()
            this.UpdateImage()
        }
}


func (this *AllBinaryJ2SEImageRotationAnimation) ChangeBasicColor(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetChangeBasicColor() == nil || this.GetChangeBasicColor().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.ImageBaseRotationAnimation.ChangeBasicColor(basicColor)
        if changed {
            this.ChangeColorProcessor = ChangeColorProcessor_GetInstance()
            this.UpdateImage()
        }
}


func (this *AllBinaryJ2SEImageRotationAnimation) SetAlpha(alpha int) {
        var changed bool = false; _ = changed
        if this.AlphaP != alpha {
            changed = true
        }
        this.ImageBaseRotationAnimation.SetAlpha(alpha)
        if changed {
            this.AlphaProcessor = AlphaProcessor_GetInstance()
            this.UpdateImage()
        }
}


func (this *AllBinaryJ2SEImageRotationAnimation) SetScale(scaleX float32, scaleY float32) {
        this.ScaleX = scaleX
        this.ScaleY = scaleY
        if this.ScaleX != this.LastScaleX || this.ScaleY != this.LastScaleY {
            this.ScaleProcessor.Update(this.RealOriginalImage, this.OriginalImageArray, this.TwoImages, this.BufferedImageIndex, this.ScaleX, this.ScaleY, this.MaxScaleX, this.MaxScaleY)
            this.UpdateImage()
        }
        this.LastScaleX = this.ScaleX
        this.LastScaleY = this.ScaleY
}


func (this *AllBinaryJ2SEImageRotationAnimation) SetMaxScale(maxScaleX float32, maxScaleY float32) {
        this.MaxScaleX = maxScaleX
        this.MaxScaleY = maxScaleY
        this.ScaleProcessor.Update(this.RealOriginalImage, this.OriginalImageArray, this.TwoImages, this.BufferedImageIndex, this.ScaleX, this.ScaleY, this.MaxScaleX, this.MaxScaleY)
        this.UpdateImage()
}


func (this *AllBinaryJ2SEImageRotationAnimation) NextRotation() {
        this.ImageBaseRotationAnimation.NextRotation()
        this.UpdateImage()
}


func (this *AllBinaryJ2SEImageRotationAnimation) PreviousRotation() {
        this.ImageBaseRotationAnimation.PreviousRotation()
        this.UpdateImage()
}


func (this *AllBinaryJ2SEImageRotationAnimation) UpdateImage() {
        this.ImageRotationUtil.RotateImage(this.OriginalImageArray[0], this.TwoImages[this.BufferedImageIndex], int(this.AngleInfo.GetAngle()) + 90)
        this.AlphaProcessor.UpdateAlpha(this.ImageModifierUtil, NullImage_NULL_IMAGE, this.TwoImages[this.BufferedImageIndex], 0, this.AlphaP)
        this.SetColorProcessor.UpdateColor(this.ImageModifierUtil, NullImage_NULL_IMAGE, this.TwoImages[this.BufferedImageIndex], 0, this.BasicColor)
        this.ChangeColorProcessor.UpdateColor(this.ImageModifierUtil, NullImage_NULL_IMAGE, this.TwoImages[this.BufferedImageIndex], 0, this.ChangeBasicColorP)
        this.Swap()
}


func (this *AllBinaryJ2SEImageRotationAnimation) SetFrame(index int) {
        this.ImageBaseRotationAnimation.SetFrame(index)
        this.UpdateImage()
}


func (this *AllBinaryJ2SEImageRotationAnimation) Swap() {
        this.ImageToShow = this.TwoImages[this.BufferedImageIndex]
        if this.BufferedImageIndex == 0 {
            this.BufferedImageIndex = 1
        } else {
            this.BufferedImageIndex = 0
        }
}


func (this *AllBinaryJ2SEImageRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.ImageToShow, x, y, this.Anchor)
}


// throws Exception
func (this *AllBinaryJ2SEImageRotationAnimation) Close() {
        this.ImageBaseRotationAnimation.Close()
        var disposalUtil *DisposalUtil = DisposalUtil_GetInstance(); _ = disposalUtil
        var size2 int = len(this.TwoImages); _ = size2
        for index := int(0); index < size2; index++ {
            disposalUtil.DisposeImage(this.TwoImages[index])
        }
        var size int = len(this.OriginalImageArray); _ = size
        for index := int(0); index < size; index++ {
            disposalUtil.DisposeImage(this.OriginalImageArray[index])
        }
        disposalUtil.DisposeImage(this.RealOriginalImage)
        disposalUtil.DisposeImage(this.ImageToShow)
}


// throws Throwable
func (this *AllBinaryJ2SEImageRotationAnimation) Finalize() {
        this.ImageBaseRotationAnimation.Finalize()
        var disposalUtil *DisposalUtil = DisposalUtil_GetInstance(); _ = disposalUtil
        var size2 int = len(this.TwoImages); _ = size2
        for index := int(0); index < size2; index++ {
            disposalUtil.DisposeImage(this.TwoImages[index])
        }
        var size int = len(this.OriginalImageArray); _ = size
        for index := int(0); index < size; index++ {
            disposalUtil.DisposeImage(this.OriginalImageArray[index])
        }
        disposalUtil.DisposeImage(this.RealOriginalImage)
        disposalUtil.DisposeImage(this.ImageToShow)
}

