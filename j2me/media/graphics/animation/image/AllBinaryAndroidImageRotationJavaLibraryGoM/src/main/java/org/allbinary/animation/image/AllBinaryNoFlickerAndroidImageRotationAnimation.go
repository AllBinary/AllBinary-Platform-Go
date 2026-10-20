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

type AllBinaryNoFlickerAndroidImageRotationAnimation struct {
    ImageBaseRotationAnimation
    ImageModifierUtil *ImageModifierUtil
    AndroidImageUtil *AndroidImageUtil
    Matrix *Matrix
    HalfWidth int
    HalfHeight int
    Increment int16
    OriginalImage IImage
    TwoImages []IImage
    ImageToShow IImage
    BufferedImageIndex int
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) initFields() {
    this.ImageModifierUtil = ImageModifierUtil_GetInstanceOrCreate()
    this.AndroidImageUtil = AndroidImageUtil_GetInstance()
    this.Matrix = NewMatrix()
    
    
    
    
    this.TwoImages = make([]IImage, 2)
    
    
}


func NewAllBinaryNoFlickerAndroidImageRotationAnimation(originalImage IImage, image IImage, angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) *AllBinaryNoFlickerAndroidImageRotationAnimation {
    this := &AllBinaryNoFlickerAndroidImageRotationAnimation{}
        this.ImageBaseRotationAnimation = *NewImageBaseRotationAnimation(image, angleInfo, totalAngle, animationBehavior)
        this.initFields()
        this.OriginalImage = originalImage
        this.HalfWidth = (image.GetWidth() >> 1)
        this.HalfHeight = (image.GetHeight() >> 1)
        this.Increment = (this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())
        this.ImageToShow = image
        this.TwoImages[0] = image
        this.TwoImages[1] = ImageCopyUtil_GetInstance().CreateImage(image)
    return this
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            this.ImageModifierUtil.SetColor(this.OriginalImage, this.ImageToShow, 0, basicColor)
            changed = true
        }
        this.ImageBaseRotationAnimation.SetBasicColorP(basicColor)
        if changed {
            this.Matrix.SetRotate(0.0, float32(this.HalfWidth), float32(this.HalfHeight))
            this.UpdateImage()
        }
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) ChangeBasicColor(basicColor *BasicColor) {
        this.SetBasicColorP(basicColor)
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) SetAlpha(alpha int) {
        var changed bool = false; _ = changed
        if this.AlphaP != alpha {
            changed = true
        }
        this.ImageBaseRotationAnimation.SetAlpha(alpha)
        this.ImageModifierUtil.SetAlpha(this.OriginalImage, this.ImageToShow, 0, this.AlphaP)
        if changed {
            this.Matrix.SetRotate(0.0, float32(this.HalfWidth), float32(this.HalfHeight))
            this.UpdateImage()
        }
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) NextRotation() {
        this.ImageBaseRotationAnimation.NextRotation()
        this.Matrix.SetRotate(float32(this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) PreviousRotation() {
        this.ImageBaseRotationAnimation.PreviousRotation()
        this.Matrix.SetRotate(float32(-this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) UpdateImage() {
        this.AndroidImageUtil.Rotate(this.TwoImages[this.BufferedImageIndex], this.OriginalImage, this.Matrix, this.ImageModifierUtil.Paint)
        this.Swap()
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) SetFrame(index int) {
        var currentFrame int = this.CircularIndexUtil.GetIndex(); _ = currentFrame
        this.CircularIndexUtil.SetIndex(index)
        var newFrame int = this.CircularIndexUtil.GetIndex(); _ = newFrame
        this.AngleInfo.AdjustAngle(newFrame)
        this.Matrix.SetRotate(float32((newFrame - currentFrame)) * float32(this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) Swap() {
        this.ImageToShow = this.TwoImages[this.BufferedImageIndex]
        if this.BufferedImageIndex == 0 {
            this.BufferedImageIndex = 1
        } else {
            this.BufferedImageIndex = 0
        }
}


func (this *AllBinaryNoFlickerAndroidImageRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.ImageToShow, x, y, this.Anchor)
}

