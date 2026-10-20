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

type ImageCopyUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameFeatureFactory *GameFeatureFactory
    Features *Features
    OpenGLFeatureFactory *OpenGLFeatureFactory
    OpenGLUtil *OpenGLUtil
    OpenGLESImageExclusionUtil *OpenGLESImageExclusionUtil
    NO_COPY string
    UNABLE string
    INFORMATION string
    LEADERBOARD string
}

var ImageCopyUtil_Instance *ImageCopyUtil = NewImageCopyUtil()

func (this *ImageCopyUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GameFeatureFactory = GameFeatureFactory_GetInstance()
    this.Features = Features_GetInstance()
    this.OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance()
    this.OpenGLUtil = OpenGLUtil_GetInstance()
    this.OpenGLESImageExclusionUtil = OpenGLESImageExclusionUtil_GetInstance()
    this.NO_COPY = "SWT should not copy images after initial loading as the alpha is not honored"
    this.UNABLE = "Unable to scale image without dimensions"
    this.INFORMATION = "about_"
    this.LEADERBOARD = "leaderboard"
}


func ImageCopyUtil_GetInstance() *ImageCopyUtil {
        return ImageCopyUtil_Instance
}


func NewImageCopyUtil() *ImageCopyUtil {
    this := &ImageCopyUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *ImageCopyUtil) CreateImageForRotation(originalImage IImage) IImage {
        if this.Features.IsFeature(this.OpenGLFeatureFactory.OPENGL) {
            var image IImage = originalImage; _ = image
            image = this.OpenGLUtil.AddImage(image)
            return image
        } else {
            return this.CreateImage(originalImage)
        }
}


// throws Exception
func (this *ImageCopyUtil) CreateImage(originalImage IImage) IImage {
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return originalImage
        }
        var originalImage2 IImage = originalImage; _ = originalImage2
        if this.Features.IsFeature(this.OpenGLFeatureFactory.OPENGL) {
            if originalImage.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(originalImage).(IOpenGLESImage)); _ = openGLESImage
                originalImage2 = openGLESImage.OpenGLESImage_().OpenGLBitmap.GetImage()
            }
        }
        var image IImage = nil; _ = image
        if originalImage2.IsMutable() {
            var originalMutableImage *SwtMutableImage = any(originalImage2).(*SwtMutableImage); _ = originalMutableImage
            image = NewSwtMutableImage(SwtDeviceComponent_CreateImage(originalMutableImage.Image.GetImageData()))
        } else {
            var originalImmutableImage *SwtImmutableImage = any(originalImage2).(*SwtImmutableImage); _ = originalImmutableImage
            image = NewSwtMutableImage(SwtDeviceComponent_CreateImage(originalImmutableImage.Image.GetImageData()))
        }
        image = this.OpenGLUtil.AddImage(image)
        return image
}


// throws Exception
func (this *ImageCopyUtil) CreateImageWH(originalImage IImage, width int, height int, mutable bool) IImage {
        var originalImage2 IImage = originalImage; _ = originalImage2
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return originalImage
        }
        if this.Features.IsFeature(this.OpenGLFeatureFactory.OPENGL) {
            if this.OpenGLESImageExclusionUtil.IsCustomScaling(originalImage) {
                return this.CreateImage2(originalImage, width, height, mutable)
            }
            if originalImage.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(originalImage).(IOpenGLESImage)); _ = openGLESImage
                originalImage2 = openGLESImage.OpenGLESImage_().OpenGLBitmap.GetImage()
            }
        }
        var image IImage = nil; _ = image
        if originalImage2.IsMutable() && mutable {
            var originalMutableImage *SwtMutableImage = any(originalImage2).(*SwtMutableImage); _ = originalMutableImage
            if originalMutableImage.GetImage() != nil {
                var imageData2 *ImageData = originalMutableImage.Image.GetImageData(); _ = imageData2
                var imageData *ImageData = imageData2.ScaledTo(width, height); _ = imageData
                image = NewSwtMutableImage(SwtDeviceComponent_CreateImage(imageData))
            } else {
                image = NewSwtMutableImage(originalMutableImage.GetName(), NewPostLoadSwtMutableImageProcessor(originalMutableImage))
            }
        } else {
            var originalImmutableImage *SwtImmutableImage = any(originalImage2).(*SwtImmutableImage); _ = originalImmutableImage
            if originalImmutableImage.GetImage() != nil {
                var imageData2 *ImageData = originalImmutableImage.Image.GetImageData(); _ = imageData2
                if width > 0 && height > 0 {
                    var imageData *ImageData = imageData2.ScaledTo(width, height); _ = imageData
                    image = NewSwtImmutableImage(originalImage.GetName(), SwtDeviceComponent_CreateImage(imageData))
                } else {
                    this.LogUtil.PutF(this.UNABLE, this, this.CommonStrings.CONSTRUCTOR)
                    image = NewSwtImmutableImage(originalImmutableImage.GetName(), NewPostLoadSwtImmutableImageProcessor(originalImmutableImage))
                }
            } else {
                image = NewSwtImmutableImage(originalImmutableImage.GetName(), NewPostLoadSwtImmutableImageProcessor(originalImmutableImage))
            }
        }
        image = this.OpenGLUtil.AddImage(image)
        return image
}


// throws Exception
func (this *ImageCopyUtil) CreateImage2(originalImage IImage, width int, height int, mutable bool) IImage {
        var originalImage2 IImage = originalImage; _ = originalImage2
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return originalImage
        }
        var image IImage = nil; _ = image
        if this.Features.IsFeature(this.OpenGLFeatureFactory.OPENGL) {
            var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
            if originalImage.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(originalImage).(IOpenGLESImage)); _ = openGLESImage
                var openGLESImageProperties *OpenGLESImageProperties = openGLESImage.OpenGLESImage_().OpenGLESImageProperties; _ = openGLESImageProperties
                var openGLWidth int = openGLESImage.OpenGLESImage_().OpenGLBitmap.GetWidth(); _ = openGLWidth
                var openGLHeight int = openGLESImage.OpenGLESImage_().OpenGLBitmap.GetHeight(); _ = openGLHeight
                openGLESImageProperties.ScaleX2 = ((float32(width)) / float32(openGLWidth))
                openGLESImageProperties.ScaleX = openGLESImageProperties.ScaleX2
                openGLESImageProperties.ScaleY2 = ((float32(height)) / float32(openGLHeight))
                openGLESImageProperties.ScaleY = openGLESImageProperties.ScaleY2
                openGLESImageProperties.ScaleX = openGLESImageProperties.ScaleX * 0.75
                openGLESImageProperties.ScaleY = openGLESImageProperties.ScaleY * 0.75
                if originalImage.GetHeight() % 2 != 0 {
                    openGLESImageProperties.ScaleX2 += 1.0 / float32(openGLWidth)
                    openGLESImageProperties.ScaleY2 += 1.0 / float32(openGLHeight)
                }
                image = openGLESImage
            } else {
                image = originalImage
                var openGLESImage IOpenGLESImage = any(this.OpenGLUtil.AddImage(image)).(IOpenGLESImage); _ = openGLESImage
                var openGLESImageProperties *OpenGLESImageProperties = openGLESImage.OpenGLESImage_().OpenGLESImageProperties; _ = openGLESImageProperties
                openGLESImageProperties.ScaleX2 = float32(((float32(width)) / openGLESImage.GetWidth()))
                openGLESImageProperties.ScaleX = openGLESImageProperties.ScaleX2
                openGLESImageProperties.ScaleY2 = float32(((float32(height)) / openGLESImage.GetHeight()))
                openGLESImageProperties.ScaleY = openGLESImageProperties.ScaleY2
                if strings.HasPrefix(image.GetName(), this.INFORMATION) {
                    openGLESImage.OpenGLESImage_().OpenGLESImageTranslate = NewOpenGLESDeviceImageTranslate()
                    var openGLESDeviceImageTranslate *OpenGLESDeviceImageTranslate = any(openGLESImage.OpenGLESImage_().OpenGLESImageTranslate).(*OpenGLESDeviceImageTranslate); _ = openGLESDeviceImageTranslate
                    openGLESDeviceImageTranslate.TranslateX = float32(displayInfoSingleton.GetLastWidth()) / 1.4 / openGLESImageProperties.ScaleX
                    openGLESDeviceImageTranslate.TranslateY = float32(displayInfoSingleton.GetLastHeight() / 4) / openGLESImageProperties.ScaleY
                    openGLESImageProperties.ScaleX = openGLESImageProperties.ScaleX * 0.25
                    openGLESImageProperties.ScaleY = openGLESImageProperties.ScaleY * 0.25
                } else {
                    openGLESImageProperties.ScaleX = openGLESImageProperties.ScaleX * 0.75
                    openGLESImageProperties.ScaleY = openGLESImageProperties.ScaleY * 0.75
                    if strings.HasPrefix(image.GetName(), this.LEADERBOARD) {
                        openGLESImage.OpenGLESImage_().OpenGLESImageTranslate = NewOpenGLESDeviceImageTranslate()
                        var openGLESDeviceImageTranslate *OpenGLESDeviceImageTranslate = any(openGLESImage.OpenGLESImage_().OpenGLESImageTranslate).(*OpenGLESDeviceImageTranslate); _ = openGLESDeviceImageTranslate
                        openGLESDeviceImageTranslate.TranslateX = float32(-displayInfoSingleton.GetLastWidth() / 40) / openGLESImageProperties.ScaleX
                    }
                    if strings.HasPrefix(image.GetName(), this.LEADERBOARD) {
                    } else {
                        openGLESImageProperties.ScaleX = 1.0
                        openGLESImageProperties.ScaleY = 1.0
                    }
                }
                image = openGLESImage
            }
        } else {
            if originalImage2.IsMutable() && mutable {
                var originalMutableImage *SwtMutableImage = any(originalImage2).(*SwtMutableImage); _ = originalMutableImage
                if originalMutableImage.GetImage() != nil {
                    var imageData2 *ImageData = originalMutableImage.Image.GetImageData(); _ = imageData2
                    var imageData *ImageData = imageData2.ScaledTo(width, height); _ = imageData
                    image = NewSwtMutableImage(SwtDeviceComponent_CreateImage(imageData))
                } else {
                    image = NewSwtMutableImage(originalMutableImage.GetName(), NewPostLoadSwtMutableImageProcessor(originalMutableImage))
                }
            } else {
                var originalImmutableImage *SwtImmutableImage = any(originalImage2).(*SwtImmutableImage); _ = originalImmutableImage
                if originalImmutableImage.GetImage() != nil {
                    var imageData2 *ImageData = originalImmutableImage.Image.GetImageData(); _ = imageData2
                    var imageData *ImageData = imageData2.ScaledTo(width, height); _ = imageData
                    image = NewSwtImmutableImage(originalImage.GetName(), SwtDeviceComponent_CreateImage(imageData))
                } else {
                    image = NewSwtImmutableImage(originalImmutableImage.GetName(), NewPostLoadSwtImmutableImageProcessor(originalImmutableImage))
                }
            }
            image = this.OpenGLUtil.AddImage(image)
        }
        return image
}


// throws Exception
func (this *ImageCopyUtil) CreateImageScale(originalImage IImage, canvasScale float32, resize bool) IImage {
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return originalImage
        }
        var newWidth int = int((originalImage.GetWidth() * canvasScale)); _ = newWidth
        var newHeight int = int((originalImage.GetHeight() * canvasScale)); _ = newHeight
        if resize {
            if newWidth < newHeight {
                newWidth = newHeight
            }
            if newHeight < newWidth {
                newHeight = newWidth
            }
        }
        var image IImage = this.CreateImageWH(originalImage, newWidth, newHeight, true); _ = image
        if image.IsMutable() {
            var originalImageData *ImageData = (any(originalImage.GetImage()).(IImage)).GetImageData(); _ = originalImageData
            var mutableImage2 *SwtMutableImage = any(image).(*SwtMutableImage); _ = mutableImage2
            var halfWidthDelta int = (newWidth - originalImage.GetWidth()) / 2; _ = halfWidthDelta
            var halfHeightDelta int = (newHeight - originalImage.GetHeight()) / 2; _ = halfHeightDelta
            var originalPixelArray []int = make([]int, originalImage.GetWidth() * originalImage.GetHeight()); _ = originalPixelArray
            var newPixelArray []int = make([]int, image.GetWidth() * image.GetHeight()); _ = newPixelArray
            var width int = originalImage.GetWidth(); _ = width
            var height int = originalImage.GetHeight(); _ = height
            for i := int(0); i < height; i++ {
                originalImageData.GetPixels(0, i, width, originalPixelArray, (i * width))
            }
            for index := int(halfWidthDelta); index < width; index++ {
                for index2 := int(halfHeightDelta); index2 < height; index2++ {
                    newPixelArray[index + (index2 * width)] = originalPixelArray[(index - halfWidthDelta) + ((index2 - halfHeightDelta) * width)]
                }
            }
            (any(mutableImage2.GetImage()).(IImage)).GetImageData().SetPixels(0, 0, image.GetWidth(), newPixelArray, 0)
            return image
        } else {
            var originalImmutableImage *SwtImmutableImage = any(image).(*SwtImmutableImage); _ = originalImmutableImage
            var originalImageData *ImageData = (any(originalImage.GetImage()).(IImage)).GetImageData(); _ = originalImageData
            var halfWidthDelta int = (newWidth - originalImage.GetWidth()) / 2; _ = halfWidthDelta
            var halfHeightDelta int = (newHeight - originalImage.GetHeight()) / 2; _ = halfHeightDelta
            var originalPixelArray []int = make([]int, originalImage.GetWidth() * originalImage.GetHeight()); _ = originalPixelArray
            var newPixelArray []int = make([]int, image.GetWidth() * image.GetHeight()); _ = newPixelArray
            var width int = originalImage.GetWidth(); _ = width
            var height int = originalImage.GetHeight(); _ = height
            for i := int(0); i < height; i++ {
                originalImageData.GetPixels(0, i, width, originalPixelArray, (i * width))
            }
            for index := int(halfWidthDelta); index < width; index++ {
                for index2 := int(halfHeightDelta); index2 < height; index2++ {
                    newPixelArray[index + (index2 * width)] = originalPixelArray[(index - halfWidthDelta) + ((index2 - halfHeightDelta) * width)]
                }
            }
            (any(originalImmutableImage.GetImage()).(IImage)).GetImageData().SetPixels(0, 0, image.GetWidth(), newPixelArray, 0)
            return image
        }
}

