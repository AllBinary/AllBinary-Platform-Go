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

type AllBinaryHTMLImageRotationAnimation struct {
    ImageBaseRotationAnimation
    ImageRotationUtil *ImageRotationUtil
    ImageModifierUtil *ImageModifierUtil
    OriginalImage IImage
    TwoImages []IImage
    CanvasSurfaceArray []*CanvasSurface
    ImageToShow IImage
    BufferedImageIndex int
    AlphaProcessor IAlphaBaseProcessor
}

// The Java class AllBinaryHTMLImageRotationAnimation has subclasses so its variables are this interface
type IAllBinaryHTMLImageRotationAnimation interface {
    IImageBaseRotationAnimation
    AllBinaryHTMLImageRotationAnimation_() *AllBinaryHTMLImageRotationAnimation
    GetCanvasSurface(image IImage) *CanvasSurface
    SetBasicColorP(basicColor *BasicColor)
    SetAlpha(alpha int)
    NextRotation()
    PreviousRotation()
    UpdateImage()
    SetFrame(index int)
    Swap()
    PaintXY(graphics IGraphics, x int, y int)
}

func (this *AllBinaryHTMLImageRotationAnimation) AllBinaryHTMLImageRotationAnimation_() *AllBinaryHTMLImageRotationAnimation {
    return this
}


func (this *AllBinaryHTMLImageRotationAnimation) initFields() {
    this.ImageRotationUtil = ImageRotationUtil_GetInstance()
    this.ImageModifierUtil = ImageModifierUtil_GetInstanceOrCreate()
    
    this.TwoImages = make([]IImage, 2)
    this.CanvasSurfaceArray = make([]*CanvasSurface, 2)
    
    
    this.AlphaProcessor = AlphaBaseProcessor_GetInstance()
}


func NewAllBinaryHTMLImageRotationAnimation(originalImage IImage, image IImage, angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) *AllBinaryHTMLImageRotationAnimation {
    this := &AllBinaryHTMLImageRotationAnimation{}
        this.ImageBaseRotationAnimation = *NewImageBaseRotationAnimation(image, angleInfo, totalAngle, animationBehavior)
        this.initFields()
        this.OriginalImage = originalImage
        this.ImageToShow = image
        this.TwoImages[0] = image
        this.TwoImages[1] = ImageCopyUtil_GetInstance().CreateImage(image)
        this.CanvasSurfaceArray[0] = this.GetCanvasSurface(this.TwoImages[0])
        this.CanvasSurfaceArray[1] = this.GetCanvasSurface(this.TwoImages[1])
    return this
}


func (this *AllBinaryHTMLImageRotationAnimation) GetCanvasSurface(image IImage) *CanvasSurface {
        var htmlImage *PlaynMutableImage = any(image).(*PlaynMutableImage); _ = htmlImage
        var canvasImage *CanvasImage = any(htmlImage.GetImage()).(*CanvasImage); _ = canvasImage
        var canvasSurface *CanvasSurface = htmlImage.GetCanvasSurface(canvasImage); _ = canvasSurface
        canvasSurface.Translate(this.OriginalImage.GetWidth() / 2, this.OriginalImage.GetHeight() / 2)
        return canvasSurface
}


func (this *AllBinaryHTMLImageRotationAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.ImageBaseRotationAnimation.SetBasicColorP(basicColor)
        if changed {
            this.UpdateImage()
        }
}


func (this *AllBinaryHTMLImageRotationAnimation) SetAlpha(alpha int) {
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


func (this *AllBinaryHTMLImageRotationAnimation) NextRotation() {
        this.ImageBaseRotationAnimation.NextRotation()
        this.UpdateImage()
}


func (this *AllBinaryHTMLImageRotationAnimation) PreviousRotation() {
        this.ImageBaseRotationAnimation.PreviousRotation()
        this.UpdateImage()
}


func (this *AllBinaryHTMLImageRotationAnimation) UpdateImage() {
        var canvasSurface *CanvasSurface = this.CanvasSurfaceArray[this.BufferedImageIndex]; _ = canvasSurface
        canvasSurface.Save()
        this.ImageRotationUtil.RotateImageCanvasSurfaceClear(this.OriginalImage, this.TwoImages[this.BufferedImageIndex], canvasSurface, int(this.AngleInfo.GetAngle()) + 90)
        this.AlphaProcessor.SetAlpha(this.ImageModifierUtil, this.OriginalImage, this.TwoImages[this.BufferedImageIndex], this.AlphaP)
        this.ImageRotationUtil.DrawImage(this.OriginalImage, this.ImageToShow, canvasSurface)
        canvasSurface.Restore()
        this.Swap()
}


func (this *AllBinaryHTMLImageRotationAnimation) SetFrame(index int) {
        this.ImageBaseRotationAnimation.SetFrame(index)
        this.UpdateImage()
}


func (this *AllBinaryHTMLImageRotationAnimation) Swap() {
        this.ImageToShow = this.TwoImages[this.BufferedImageIndex]
        if this.BufferedImageIndex == 0 {
            this.BufferedImageIndex = 1
        } else {
            this.BufferedImageIndex = 0
        }
}


func (this *AllBinaryHTMLImageRotationAnimation) PaintXY(graphics IGraphics, x int, y int) {
        graphics.DrawImage(this.ImageToShow, x, y, this.Anchor)
}

