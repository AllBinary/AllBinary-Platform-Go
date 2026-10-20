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

type ImageScaleUtil struct {
}

var ImageScaleUtil_Instance *ImageScaleUtil = NewImageScaleUtil()


func ImageScaleUtil_GetInstance() *ImageScaleUtil {
        return ImageScaleUtil_Instance
}


func NewImageScaleUtil() *ImageScaleUtil {
    this := &ImageScaleUtil{}
    return this
}


// throws Exception
func (this *ImageScaleUtil) CreateImage2(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool) IImage {
        return this.CreateImage3(imageCache, originalImage, scaleNominatorX, scaleDenominatorX, scaleNominatorY, scaleDenominatorY, cached, true)
}


// throws Exception
func (this *ImageScaleUtil) CreateImage3(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool, mutable bool) IImage {
        var scaleX float32 = scaleNominatorX / scaleDenominatorX; _ = scaleX
        var scaleY float32 = scaleNominatorY / scaleDenominatorY; _ = scaleY
        return this.CreateImage(imageCache, originalImage, scaleX, scaleY, cached)
}


// throws Exception
func (this *ImageScaleUtil) CreateImage(imageCache IImageCache, originalImage IImage, scaleX float32, scaleY float32, cached bool) IImage {
        return this.CreateImageInternal(originalImage, scaleX, scaleY)
}


// throws Exception
func (this *ImageScaleUtil) CreateImageInternal(originalImage IImage, scaleX float32, scaleY float32) IImage {
        var image IImage = ImageCreationUtil_GetInstance(int((originalImage.GetWidth() * scaleX)), int((originalImage.GetHeight() * scaleY))); _ = image
        if image.IsMutable() {
            var originalPlayNImage IImage = nil; _ = originalPlayNImage
            if originalImage.IsMutable() {
                var originalHTMLImage *PlaynMutableImage = any(originalImage).(*PlaynMutableImage); _ = originalHTMLImage
                originalPlayNImage = any(originalHTMLImage.GetImage()).(IImage)
            } else {
                var originalHTMLImage *PlaynImmutableImage = any(originalImage).(*PlaynImmutableImage); _ = originalHTMLImage
                originalPlayNImage = any(originalHTMLImage.GetImage()).(IImage)
            }
            var htmlImage *PlaynMutableImage = any(image).(*PlaynMutableImage); _ = htmlImage
            var canvasImage *CanvasImage = any(htmlImage.GetImage()).(*CanvasImage); _ = canvasImage
            var canvasSurface *CanvasSurface = htmlImage.GetCanvasSurface(canvasImage); _ = canvasSurface
            canvasSurface.DrawImage(originalPlayNImage, 0, 0, image.GetWidth(), image.GetHeight(), 0, 0, originalPlayNImage.Width(), originalPlayNImage.Height())
            return image
        } else {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
}

