/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2023 AllBinary
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

type ImageJ2MEScaleUtil struct {
}

var ImageJ2MEScaleUtil_Instance *ImageJ2MEScaleUtil = NewImageJ2MEScaleUtil()


func ImageJ2MEScaleUtil_GetInstance() *ImageJ2MEScaleUtil {
        return ImageJ2MEScaleUtil_Instance
}


func NewImageJ2MEScaleUtil() *ImageJ2MEScaleUtil {
    this := &ImageJ2MEScaleUtil{}
    return this
}


// throws Exception
func (this *ImageJ2MEScaleUtil) ScaleArray(images []IImage, width int, height int) []IImage {
        var scaledImages []IImage = make([]IImage, len(images)); _ = scaledImages
        for index := int(len(images)); Java_PreDecrement(&index) >= 0;  {
            scaledImages[index] = this.Scale(images[index], width, height)
        }
        return scaledImages
}


// throws Exception
func (this *ImageJ2MEScaleUtil) Scale(image IImage, width int, height int) IImage {
        var sourceWidth int = image.GetWidth(); _ = sourceWidth
        var sourceHeight int = image.GetHeight(); _ = sourceHeight
        var originalData []int = make([]int, image.GetWidth() * image.GetHeight()); _ = originalData
        image.GetRGB(originalData, 0, image.GetWidth(), 0, 0, image.GetWidth(), image.GetHeight())
        var scaledData []int = make([]int, width * height); _ = scaledData
        var heightRatioFactor int = (8 * sourceHeight) / height; _ = heightRatioFactor
        var widthRatioFactor int = (8 * sourceWidth) / width; _ = widthRatioFactor
        var scaledIndex int = len(scaledData) - 1; _ = scaledIndex
        var dx int; _ = dx
        var dy int; _ = dy
        for index := int(height); Java_PreDecrement(&index) >= 0;  {
            dy = (index * heightRatioFactor) >> 3
            for index2 := int(width); Java_PreDecrement(&index2) >= 0;  {
                dx = (index2 * widthRatioFactor) >> 3
                scaledData[Java_PostDecrement(&scaledIndex)] = originalData[(sourceWidth * dy) + dx]
            }
        }
        var scaledImage IImage = Image_CreateRGBImage(scaledData, width, height, true); _ = scaledImage
        return scaledImage
}

