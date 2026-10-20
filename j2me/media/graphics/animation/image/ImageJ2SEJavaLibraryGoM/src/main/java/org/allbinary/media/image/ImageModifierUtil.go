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
    ImageUtil *ImageUtil
    ColorArray []int
}

var ImageModifierUtil_Instance *ImageModifierUtil = NewImageModifierUtil()

func (this *ImageModifierUtil) initFields() {
    this.ImageUtil = ImageUtil_GetInstance()
    this.ColorArray = make([]int, 1)
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
        var newBufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(image); _ = newBufferedImage
        var r int16 = basicColor.Red; _ = r
        var g int16 = basicColor.Green; _ = g
        var b int16 = basicColor.Blue; _ = b
        var width int = image.GetWidth(); _ = width
        var height int = image.GetHeight(); _ = height
        var colorModel *ColorModel; _ = colorModel
        var dataElements any; _ = dataElements
        var ax int; _ = ax
        var rx int; _ = rx
        var gx int; _ = gx
        var bx int; _ = bx
        for index := int(0); index < width; index++ {
            for index2 := int(0); index2 < height; index2++ {
                colorModel = newBufferedImage.GetColorModel()
                dataElements = newBufferedImage.GetRaster().GetDataElements(index, index2, nil)
                ax = colorModel.GetAlpha(dataElements)
                rx = colorModel.GetRed(dataElements)
                gx = colorModel.GetGreen(dataElements)
                bx = colorModel.GetBlue(dataElements)
                if rx != 0 || gx != 0 || bx != 0 {
                    newBufferedImage.SetRGB(index, index2, (ax << 24) | (r << 16) | (g << 8) | int((b)))
                }
            }
        }
}


func (this *ImageModifierUtil) ChangeColor(unusedOriginalImage IImage, image IImage, imageIndex int, basicColor *BasicColor) {
        var newBufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(image); _ = newBufferedImage
        var MAX float32 = 255.0; _ = MAX
        var r float32 = (float32(basicColor.Red)) / MAX; _ = r
        var g float32 = (float32(basicColor.Green)) / MAX; _ = g
        var b float32 = (float32(basicColor.Blue)) / MAX; _ = b
        var width int = image.GetWidth(); _ = width
        var height int = image.GetHeight(); _ = height
        var colorModel *ColorModel; _ = colorModel
        var dataElements any; _ = dataElements
        var ax int; _ = ax
        var rx int; _ = rx
        var gx int; _ = gx
        var bx int; _ = bx
        for index := int(0); index < width; index++ {
            for index2 := int(0); index2 < height; index2++ {
                colorModel = newBufferedImage.GetColorModel()
                dataElements = newBufferedImage.GetRaster().GetDataElements(index, index2, nil)
                ax = colorModel.GetAlpha(dataElements)
                rx = colorModel.GetRed(dataElements)
                gx = colorModel.GetGreen(dataElements)
                bx = colorModel.GetBlue(dataElements)
                rx = int(rx * r)
                gx = int(gx * g)
                bx = int(bx * b)
                newBufferedImage.SetRGB(index, index2, (ax << 24) | (rx << 16) | (gx << 8) | (bx))
            }
        }
}


func (this *ImageModifierUtil) SetAlpha(unusedOriginalImage IImage, image IImage, imageIndex int, alphaInt int) {
        var alpha int8 = int8(alphaInt); _ = alpha
        var newBufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(image); _ = newBufferedImage
        var width int = image.GetWidth(); _ = width
        var height int = image.GetHeight(); _ = height
        var raster *WritableRaster = newBufferedImage.GetAlphaRaster(); _ = raster
        for index := int(0); index < width; index++ {
            for index2 := int(0); index2 < height; index2++ {
                raster.GetPixel(index, index2, this.ColorArray)
                this.ColorArray[0] = int(alpha) & this.ColorArray[0]
                raster.SetPixel(index, index2, this.ColorArray)
            }
        }
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
                var image IImage = ImageCreationUtil_GetInstance().CreateImageWH(originalImage.GetWidth(), originalImage.GetHeight()); _ = image
                image.GetGraphics().DrawImage(originalImage, 0, 0, Anchor_TOP_LEFT)
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

