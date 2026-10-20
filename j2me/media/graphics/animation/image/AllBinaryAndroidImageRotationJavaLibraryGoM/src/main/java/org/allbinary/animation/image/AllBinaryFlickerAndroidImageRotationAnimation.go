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

type AllBinaryFlickerAndroidImageRotationAnimation struct {
    ImageBaseRotationAnimation
    ImageModifierUtil *ImageModifierUtil
    AndroidImageUtil *AndroidImageUtil
    Matrix *Matrix
    HalfWidth int
    HalfHeight int
    Increment int16
    OriginalImage IImage
}

// The Java class AllBinaryFlickerAndroidImageRotationAnimation has subclasses so its variables are this interface
type IAllBinaryFlickerAndroidImageRotationAnimation interface {
    IImageBaseRotationAnimation
    AllBinaryFlickerAndroidImageRotationAnimation_() *AllBinaryFlickerAndroidImageRotationAnimation
    SetBasicColorP(basicColor *BasicColor)
    ChangeBasicColor(basicColor *BasicColor)
    SetAlpha(alpha int)
    NextRotation()
    PreviousRotation()
    UpdateImage()
    SetFrame(index int)
}

func (this *AllBinaryFlickerAndroidImageRotationAnimation) AllBinaryFlickerAndroidImageRotationAnimation_() *AllBinaryFlickerAndroidImageRotationAnimation {
    return this
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) initFields() {
    this.ImageModifierUtil = ImageModifierUtil_GetInstanceOrCreate()
    this.AndroidImageUtil = AndroidImageUtil_GetInstance()
    this.Matrix = NewMatrix()
    
    
    
    
}


func NewAllBinaryFlickerAndroidImageRotationAnimation(originalImage IImage, image IImage, angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) *AllBinaryFlickerAndroidImageRotationAnimation {
    this := &AllBinaryFlickerAndroidImageRotationAnimation{}
        this.ImageBaseRotationAnimation = *NewImageBaseRotationAnimation(image, angleInfo, totalAngle, animationBehavior)
        this.initFields()
        this.OriginalImage = originalImage
        this.HalfWidth = (image.GetWidth() >> 1)
        this.HalfHeight = (image.GetHeight() >> 1)
        this.Increment = (this.AngleInfo.GetAngleIncrementInfo().GetAngleIncrement())
    return this
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            this.ImageModifierUtil.SetColor(this.OriginalImage, this.GetImage(), 0, basicColor)
            changed = true
        }
        this.ImageBaseRotationAnimation.SetBasicColorP(basicColor)
        if changed {
            this.Matrix.SetRotate(0.0, float32(this.HalfWidth), float32(this.HalfHeight))
            this.UpdateImage()
        }
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) ChangeBasicColor(basicColor *BasicColor) {
        this.SetBasicColorP(basicColor)
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) SetAlpha(alpha int) {
        var changed bool = false; _ = changed
        if this.AlphaP != alpha {
            changed = true
        }
        this.ImageBaseRotationAnimation.SetAlpha(alpha)
        this.ImageModifierUtil.SetAlpha(this.OriginalImage, this.GetImage(), 0, this.AlphaP)
        if changed {
            this.Matrix.SetRotate(0.0, float32(this.HalfWidth), float32(this.HalfHeight))
            this.UpdateImage()
        }
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) NextRotation() {
        this.ImageBaseRotationAnimation.NextRotation()
        this.Matrix.SetRotate(float32(this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) PreviousRotation() {
        this.ImageBaseRotationAnimation.PreviousRotation()
        this.Matrix.SetRotate(float32(-this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) UpdateImage() {
        this.AndroidImageUtil.Rotate(this.GetImage(), this.OriginalImage, this.Matrix, this.ImageModifierUtil.Paint)
}


func (this *AllBinaryFlickerAndroidImageRotationAnimation) SetFrame(index int) {
        var currentFrame int = this.CircularIndexUtil.GetIndex(); _ = currentFrame
        this.CircularIndexUtil.SetIndex(index)
        var newFrame int = this.CircularIndexUtil.GetIndex(); _ = newFrame
        this.AngleInfo.AdjustAngle(newFrame)
        this.Matrix.SetRotate(float32((newFrame - currentFrame)) * float32(this.Increment), float32(this.HalfWidth), float32(this.HalfHeight))
        this.UpdateImage()
}

