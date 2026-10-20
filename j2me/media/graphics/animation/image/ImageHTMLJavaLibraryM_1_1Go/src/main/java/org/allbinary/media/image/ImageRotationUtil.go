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

type ImageRotationUtil struct {
}

var ImageRotationUtil_Instance *ImageRotationUtil = NewImageRotationUtil()


func ImageRotationUtil_GetInstance() *ImageRotationUtil {
        return ImageRotationUtil_Instance
}


func NewImageRotationUtil() *ImageRotationUtil {
    this := &ImageRotationUtil{}
    return this
}


func (this *ImageRotationUtil) RotateImage(originalImage IImage, image IImage, totalAngle int) IImage {
        if image.IsMutable() {
            var htmlImage *PlaynMutableImage = any(image).(*PlaynMutableImage); _ = htmlImage
            var canvasImage *CanvasImage = any(htmlImage.GetImage()).(*CanvasImage); _ = canvasImage
            var canvasSurface *CanvasSurface = htmlImage.GetCanvasSurface(canvasImage); _ = canvasSurface
            canvasSurface.Translate(originalImage.GetWidth() / 2, originalImage.GetHeight() / 2)
            var image2 IImage = this.RotateImageCanvasSurface(originalImage, image, canvasSurface, totalAngle); _ = image2
            this.DrawImage(originalImage, image, canvasSurface)
            return image2
        } else {
            return nil
        }
}


func (this *ImageRotationUtil) RotateImageCanvasSurfaceClear(originalImage IImage, image IImage, canvasSurface *CanvasSurface, totalAngle int) IImage {
        if image.IsMutable() {
            canvasSurface.Translate(-originalImage.GetWidth() / 2, -originalImage.GetHeight() / 2)
            canvasSurface.Clear()
            canvasSurface.Translate(originalImage.GetWidth() / 2, originalImage.GetHeight() / 2)
            return this.RotateImageCanvasSurface(originalImage, image, canvasSurface, totalAngle)
        } else {
            return nil
        }
}


func (this *ImageRotationUtil) RotateImageCanvasSurface(originalImage IImage, image IImage, canvasSurface *CanvasSurface, totalAngle int) IImage {
        if image.IsMutable() {
            canvasSurface.Rotate(float32((float64(totalAngle) * math.Pi / 180)))
            return image
        } else {
            return nil
        }
}


func (this *ImageRotationUtil) DrawImage(originalImage IImage, image IImage, canvasSurface *CanvasSurface) {
        var originalPlayNImage IImage = nil; _ = originalPlayNImage
        if originalImage.IsMutable() {
            var originalHTMLImage *PlaynMutableImage = any(originalImage).(*PlaynMutableImage); _ = originalHTMLImage
            originalPlayNImage = any(originalHTMLImage.GetImage()).(IImage)
        } else {
            var originalHTMLImage *PlaynImmutableImage = any(originalImage).(*PlaynImmutableImage); _ = originalHTMLImage
            originalPlayNImage = any(originalHTMLImage.GetImage()).(IImage)
        }
        canvasSurface.DrawImage(originalPlayNImage, -originalImage.GetWidth() / 2, -originalImage.GetHeight() / 2)
}


// throws Exception
func (this *ImageRotationUtil) CreateRotatedImage(originalImage IImage, rotationInDegrees int) IImage {
        var image IImage = ImageCreationUtil_GetInstance().CreateImageWH(originalImage.GetWidth(), originalImage.GetHeight()); _ = image
        var rotatedImage IImage = this.RotateImage(originalImage, image, rotationInDegrees); _ = rotatedImage
        if rotatedImage == nil {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
        return rotatedImage
}

