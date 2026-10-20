/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type ImageModifierUtil struct {
    LogUtil *LogUtil
    ImageCopyUtil *ImageCopyUtil
    Features *Features
}

var ImageModifierUtil_Instance *ImageModifierUtil = NewImageModifierUtil()

func (this *ImageModifierUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageCopyUtil = ImageCopyUtil_GetInstance()
    this.Features = Features_GetInstance()
}

func NewImageModifierUtil() *ImageModifierUtil {
    this := &ImageModifierUtil{}
    this.initFields()
    return this
}


func ImageModifierUtil_GetInstanceOrCreate() *ImageModifierUtil {
        return ImageModifierUtil_Instance
}


func (this *ImageModifierUtil) SetBasicColorP(basicColor *BasicColor) {
}


func (this *ImageModifierUtil) SetColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
        var newBufferedImage IImage; _ = newBufferedImage
        if image.IsMutable() {
            var swtImage *SwtMutableImage = any(image).(*SwtMutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        } else {
            var swtImage *SwtImmutableImage = any(image).(*SwtImmutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        }
        var r int16 = basicColor.Red; _ = r
        var g int16 = basicColor.Green; _ = g
        var b int16 = basicColor.Blue; _ = b
        var width int = image.GetWidth(); _ = width
        var height int = image.GetHeight(); _ = height
        var imageData *ImageData = newBufferedImage.GetImageData(); _ = imageData
        if imageData.Depth == 24 {
            var rx int; _ = rx
            var gx int; _ = gx
            var bx int; _ = bx
            var index3 int; _ = index3
            for index := int(0); index < width; index++ {
                for index2 := int(0); index2 < height; index2++ {
                    index3 = (index2 * imageData.BytesPerLine) + (index * 4)
                    rx = (imageData.Data[index3] & 0xFF) << 16
                    gx = (imageData.Data[index3 + 1] & 0xFF) << 8
                    bx = (imageData.Data[index3 + 2] & 0xFF)
                    if rx != 0 || gx != 0 || bx != 0 {
                        imageData.Data[index3] = int8(r)
                        imageData.Data[index3 + 1] = int8(g)
                        imageData.Data[index3 + 2] = int8(b)
                    }
                }
            }
        }
}


func (this *ImageModifierUtil) ChangeColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
        if this.Features.IsFeature(OpenGLFeatureFactory_GetInstance().OPENGL) {
            if image.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(image).(IOpenGLESImage)); _ = openGLESImage
                openGLESImage.OpenGLESImage_().ImageProcessor = ModifierOpenGLESImageProcessor_GetInstance()
                openGLESImage.OpenGLESImage_().OpenGLESImageProperties.Redf = basicColor.GetRedComponent()
                openGLESImage.OpenGLESImage_().OpenGLESImageProperties.Greenf = basicColor.GetGreenComponent()
                openGLESImage.OpenGLESImage_().OpenGLESImageProperties.Bluef = basicColor.GetBlueComponent()
                return
            }
        }
        var newBufferedImage IImage; _ = newBufferedImage
        if image.IsMutable() {
            var swtImage *SwtMutableImage = any(image).(*SwtMutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        } else {
            var swtImage *SwtImmutableImage = any(image).(*SwtImmutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        }
        var MAX float32 = 255.0; _ = MAX
        var r float32 = (float32(basicColor.Red)) / MAX; _ = r
        var g float32 = (float32(basicColor.Green)) / MAX; _ = g
        var b float32 = (float32(basicColor.Blue)) / MAX; _ = b
        var rx int; _ = rx
        var gx int; _ = gx
        var bx int; _ = bx
        var imageData *ImageData = newBufferedImage.GetImageData(); _ = imageData
        if imageData.Depth == 24 {
            var size int = len(imageData.Data); _ = size
            for index := int(0); index < size; index += 4 {
                rx = (imageData.Data[index] & 0xFF) << 16
                gx = (imageData.Data[index + 1] & 0xFF) << 8
                bx = (imageData.Data[index + 2] & 0xFF)
                rx = int(rx * r)
                gx = int(gx * g)
                bx = int(bx * b)
                imageData.Data[index] = int8(rx)
                imageData.Data[index + 1] = int8(gx)
                imageData.Data[index + 2] = int8(bx)
            }
        }
}


func (this *ImageModifierUtil) SetAlpha(unusedOriginalImage IImage, image IImage, imageIndex int, alphaInt int) {
        if this.Features.IsFeature(OpenGLFeatureFactory_GetInstance().OPENGL) {
            if image.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(image).(IOpenGLESImage)); _ = openGLESImage
                openGLESImage.OpenGLESImage_().ImageProcessor = ModifierOpenGLESImageProcessor_GetInstance()
                openGLESImage.OpenGLESImage_().OpenGLESImageProperties.Alphaf = (float32(alphaInt)) / 255
                return
            }
        }
        this.SetAlphaSWT(image, alphaInt)
}


func (this *ImageModifierUtil) SetAlphaSWT(image IImage, alphaInt int) {
        var alpha int8 = int8(alphaInt); _ = alpha
        var newBufferedImage IImage; _ = newBufferedImage
        if image.IsMutable() {
            var swtImage *SwtMutableImage = any(image).(*SwtMutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        } else {
            var swtImage *SwtImmutableImage = any(image).(*SwtImmutableImage); _ = swtImage
            newBufferedImage = any(swtImage.GetImage()).(IImage)
        }
        newBufferedImage.GetImageData().Alpha = alpha
}


func (this *ImageModifierUtil) GetImageArray(originalImageArray []IImage) []IImage {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []IImage) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, "getImageArray", e)
            tryControl1, tryResult1 = 1, originalImageArray; return
                }
            }()
            var newImageArray []IImage = make([]IImage, len(originalImageArray)); _ = newImageArray
            var size int = len(newImageArray); _ = size
            var originalImage IImage; _ = originalImage
            for index := int(0); index < size; index++ {
                originalImage = originalImageArray[index]
                var image IImage = this.ImageCopyUtil.CreateImage(originalImage); _ = image
                newImageArray[index] = image
            }
            tryControl1, tryResult1 = 1, newImageArray; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ImageModifierUtil) HandleImage(imageArray []IImage, index int, image IImage) {
}


func (this *ImageModifierUtil) Reset() {
}

