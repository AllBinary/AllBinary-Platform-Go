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
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageCopyUtil *ImageCopyUtil
    GameFeatureFactory *GameFeatureFactory
    Features *Features
    NO_COPY string
}

var ImageScaleUtil_Instance *ImageScaleUtil = NewImageScaleUtil()

func (this *ImageScaleUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageCopyUtil = ImageCopyUtil_GetInstance()
    this.GameFeatureFactory = GameFeatureFactory_GetInstance()
    this.Features = Features_GetInstance()
    this.NO_COPY = "SWT should not copy images after initial loading as the alpha is not honored"
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
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return originalImage
        }
        var width int = originalImage.GetWidth(); _ = width
        var height int = originalImage.GetHeight(); _ = height
        var scaleX float32 = scaleNominatorX / scaleDenominatorX; _ = scaleX
        var scaleY float32 = scaleNominatorY / scaleDenominatorY; _ = scaleY
        var scaledImage IImage = this.ImageCopyUtil.CreateImageWH(originalImage, int((scaleX * float32(width))), int((scaleY * float32(height))), mutable); _ = scaledImage
        return scaledImage
}


// throws Exception
func (this *ImageScaleUtil) Scale2(originalImage IImage, originalImageArray []IImage, ximageToShowArray []IImage, unused int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32) {
        this.Scale(originalImage, originalImageArray, ximageToShowArray, unused, scaleX, scaleY, maxScaleX, maxScaleY, true)
}


// throws Exception
func (this *ImageScaleUtil) Scale(originalImage IImage, originalImageArray []IImage, ximageToShowArray []IImage, unused int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32, mutable bool) {
        var width int = originalImage.GetWidth(); _ = width
        var height int = originalImage.GetHeight(); _ = height
        var scaledImage IImage = this.ImageCopyUtil.CreateImageWH(originalImage, int((scaleX * float32(width))), int((scaleY * float32(height))), mutable); _ = scaledImage
        originalImageArray[0] = scaledImage
}

