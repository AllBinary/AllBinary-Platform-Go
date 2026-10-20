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

type ImageSwtRotationUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageUtil *ImageUtil
    GameFeatureFactory *GameFeatureFactory
    Features *Features
    NO_COPY string
}

var ImageSwtRotationUtil_Instance *ImageSwtRotationUtil = NewImageSwtRotationUtil()

func (this *ImageSwtRotationUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageUtil = ImageUtil_GetInstance()
    this.GameFeatureFactory = GameFeatureFactory_GetInstance()
    this.Features = Features_GetInstance()
    this.NO_COPY = "SWT should not copy images after initial loading as the alpha is not honored"
}


func ImageSwtRotationUtil_GetInstance() *ImageSwtRotationUtil {
        return ImageSwtRotationUtil_Instance
}


func NewImageSwtRotationUtil() *ImageSwtRotationUtil {
    this := &ImageSwtRotationUtil{}
    this.initFields()
    return this
}


func (this *ImageSwtRotationUtil) GetRotatedImage(bufferedImage IImage, totalAngle int) IImage {
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return bufferedImage
        }
        return bufferedImage
}


func (this *ImageSwtRotationUtil) GetRotatedImage2(bufferedImage IImage, newBufferedImage IImage, gc *GC, angle float32) IImage {
        if !this.Features.IsFeature(this.GameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION) {
            this.LogUtil.PutF(this.NO_COPY, this, this.CommonStrings.CONSTRUCTOR)
            return bufferedImage
        }
        return bufferedImage
}


func (this *ImageSwtRotationUtil) RotateImage(originalSwtImage IImage, newSwtImage IImage, gc *GC, totalAngle int) IImage {
        return this.GetRotatedImage2(originalSwtImage, newSwtImage, gc, totalAngle)
}


func (this *ImageSwtRotationUtil) GetRotatedImageSWT(bufferedImage IImage, newSwtImage IImage, gc *GC, totalAngle int) IImage {
        return this.GetRotatedImage2(bufferedImage, newSwtImage, gc, totalAngle)
}


func (this *ImageSwtRotationUtil) GetRotatedImages(bufferedImage IImage, numberOfFrames int, totalAngle int) []IImage {
        var bufferedImageArray []IImage = make([]IImage, numberOfFrames); _ = bufferedImageArray
        var angle int; _ = angle
        var size int = len(bufferedImageArray); _ = size
        for index := int(0); index < size; index++ {
            angle = (totalAngle / size) * index
            bufferedImageArray[index] = this.GetRotatedImage(bufferedImage, angle)
        }
        return bufferedImageArray
}


func (this *ImageSwtRotationUtil) CreateSpriteImage(bufferedImageArray []IImage) IImage {
        var columns int = 9; _ = columns
        var rows int = 0; _ = rows
        var size int = len(bufferedImageArray); _ = size
        if size < columns {
            columns = size
        }
        rows = (size / columns)
        if size % columns != 0 {
            rows++
        }
        var firstBufferedImage IImage = bufferedImageArray[0]; _ = firstBufferedImage
        var imageData *ImageData = firstBufferedImage.GetImageData(); _ = imageData
        var bufferedImage IImage = this.ImageUtil.Create(imageData.Width * columns, imageData.Height * rows); _ = bufferedImage
        var gc *GC = NewGC(bufferedImage); _ = gc
        gc.SetAntialias(SWT_ON)
        gc.SetInterpolation(SWT_HIGH)
        var columnIndex int = 0; _ = columnIndex
        var rowIndex int = 0; _ = rowIndex
        var nextBufferedImage IImage; _ = nextBufferedImage
        for index := int(0); index < size; index++ {
            if index / 9 != 0 && index % 9 == 0 {
                rowIndex++
                columnIndex = 0
            }
            nextBufferedImage = bufferedImageArray[index]
            imageData = nextBufferedImage.GetImageData()
            gc.DrawImage(bufferedImage, 0, 0, bufferedImage.GetBounds().Width, bufferedImage.GetBounds().Height, imageData.Width * columnIndex, imageData.Height * rowIndex, imageData.Width, imageData.Height)
            columnIndex++
        }
        gc.Dispose()
        return bufferedImage
}

