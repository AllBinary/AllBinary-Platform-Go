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
    ImageUtil *ImageUtil
    ImageJ2SEUtil *ImageJ2SEUtil
    ImageCreationUtil *ImageCreationUtil
}

var ImageScaleUtil_Instance *ImageScaleUtil = NewImageScaleUtil()

func (this *ImageScaleUtil) initFields() {
    this.ImageUtil = ImageUtil_GetInstance()
    this.ImageJ2SEUtil = ImageJ2SEUtil_GetInstance()
    this.ImageCreationUtil = ImageCreationUtil_GetInstance()
}


func ImageScaleUtil_GetInstance() *ImageScaleUtil {
        return ImageScaleUtil_Instance
}


func NewImageScaleUtil() *ImageScaleUtil {
    this := &ImageScaleUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *ImageScaleUtil) CreateImage2(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool) IImage {
        return this.CreateImage3(imageCache, originalImage, scaleNominatorX, scaleDenominatorX, scaleNominatorY, scaleDenominatorY, cached, true)
}


// throws Exception
func (this *ImageScaleUtil) CreateImage3(imageCache IImageCache, originalImage IImage, scaleNominatorX float32, scaleDenominatorX float32, scaleNominatorY float32, scaleDenominatorY float32, cached bool, mutable bool) IImage {
        var width int = originalImage.GetWidth(); _ = width
        var height int = originalImage.GetHeight(); _ = height
        var scaleX float32 = scaleNominatorX / scaleDenominatorX; _ = scaleX
        var scaleY float32 = scaleNominatorY / scaleDenominatorY; _ = scaleY
        var image IImage; _ = image
        if cached {
            image = imageCache.Get(reflect.TypeOf(this).String(), int((float32(width) * scaleX)), int((float32(height) * scaleY)))
        } else {
            image = imageCache.Get("createImage", int((float32(width) * scaleX)), int((float32(height) * scaleY)))
        }
        this.Scale(originalImage, image, scaleX, scaleY, true)
        return image
}


// throws Exception
func (this *ImageScaleUtil) Scale2(originalImage IImage, originalImageArray []IImage, ximageToShowArray []IImage, unused int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32) {
        if maxScaleX * originalImage.GetWidth() > originalImageArray[0].GetWidth() || maxScaleY * originalImage.GetHeight() > originalImageArray[0].GetHeight() {
            originalImageArray[0] = this.ImageCreationUtil.CreateImage(originalImage.GetWidth(), originalImage.GetHeight(), maxScaleX, maxScaleY)
        }
        this.Scale(originalImage, originalImageArray[0], scaleX, scaleY, false)
}


func (this *ImageScaleUtil) Scale(originalImage IImage, newMaxSizeImage IImage, scaleX float32, scaleY float32, clear_ bool) {
        var bufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(originalImage); _ = bufferedImage
        var newBufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(newMaxSizeImage); _ = newBufferedImage
        var at *AffineTransform = AffineTransform_GetScaleInstance(scaleX, scaleY); _ = at
        var g *Graphics2D = newBufferedImage.CreateGraphics(); _ = g
        if clear_ {
            g.SetBackground(this.ImageJ2SEUtil.TRANSPARENT_COLOR)
            g.ClearRect(0, 0, newBufferedImage.GetWidth(), newBufferedImage.GetHeight())
        }
        g.DrawRenderedImage(bufferedImage, at)
        g.Dispose()
}

