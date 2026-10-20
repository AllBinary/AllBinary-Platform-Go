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
    Anchor int
    ImageBasicArrayList IBasicArrayList
    ScaledImageBasicArrayList IBasicArrayList
    ScaleXBasicArrayList IBasicArrayList
    ScaleYBasicArrayList IBasicArrayList
}

var ImageScaleUtil_Instance *ImageScaleUtil = NewImageScaleUtil()

func (this *ImageScaleUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Anchor = Anchor_TOP_LEFT
    this.ImageBasicArrayList = NewBasicArrayListD()
    this.ScaledImageBasicArrayList = NewBasicArrayListD()
    this.ScaleXBasicArrayList = NewBasicArrayListD()
    this.ScaleYBasicArrayList = NewBasicArrayListD()
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
        var scaleX float32 = scaleNominatorX / scaleDenominatorX; _ = scaleX
        var scaleY float32 = scaleNominatorY / scaleDenominatorY; _ = scaleY
        return this.CreateImage(imageCache, originalImage, scaleX, scaleY, cached)
}


// throws Exception
func (this *ImageScaleUtil) CreateImage(imageCache IImageCache, originalImage IImage, scaleX float32, scaleY float32, cached bool) IImage {
        var originalAndroidImage AndroidImageInterface = any(originalImage).(AndroidImageInterface); _ = originalAndroidImage
        var originalBitmap *Bitmap = originalAndroidImage.GetBitmap(); _ = originalBitmap
        var width int = int((originalBitmap.GetWidth() * scaleX)); _ = width
        var height int = int((originalBitmap.GetHeight() * scaleY)); _ = height
        var index int = this.ImageBasicArrayList.IndexOf(originalImage); _ = index
        var alreadyAvailable bool = false; _ = alreadyAvailable
        if index >= 0 {
            var scaleX2Float float32 = any(this.ScaleXBasicArrayList.Get(index)).(float32); _ = scaleX2Float
            var scaleX2 float32 = float32(scaleX2Float); _ = scaleX2
            var scaleY2Float float32 = any(this.ScaleYBasicArrayList.Get(index)).(float32); _ = scaleY2Float
            var scaleY2 float32 = float32(scaleY2Float); _ = scaleY2
            if scaleX2 == scaleX && scaleY2 == scaleY {
                alreadyAvailable = true
            }
        }
        if alreadyAvailable {
            return any(this.ScaledImageBasicArrayList.Get(index)).(IImage)
        } else {
            var scaledImage IImage = this.GetScaledImage(imageCache, originalImage, scaleX, scaleY, width, height, cached); _ = scaledImage
            this.ImageBasicArrayList.Add(originalImage)
            this.ScaledImageBasicArrayList.Add(scaledImage)
            this.ScaleXBasicArrayList.Add(scaleX)
            this.ScaleYBasicArrayList.Add(scaleY)
            return scaledImage
        }
}


// throws Exception
func (this *ImageScaleUtil) GetScaledImage(imageCache IImageCache, originalImage IImage, scaleX float32, scaleY float32, width int, height int, cached bool) IImage {
        var image IImage = NullImage_NULL_IMAGE; _ = image
        if cached {
            image = imageCache.Get(reflect.TypeOf(this).String(), width, height)
        } else {
            image = imageCache.Get(this.CommonStrings.CREATE_IMAGE, width, height)
        }
        if image.IsMutable() {
            var matrix *Matrix = NewMatrix(); _ = matrix
            this.Scale(image, matrix, scaleX, scaleY)
            image.GetGraphics().DrawImage(originalImage, 0, 0, this.Anchor)
            return image
        } else {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
}


func (this *ImageScaleUtil) Scale(image IImage, matrix *Matrix, scaleX float32, scaleY float32) {
        matrix.SetScale(scaleX, scaleY)
        var canvas ICanvas = (any(image).(AndroidImageInterface)).GetCanvas(); _ = canvas
        canvas.Concat(matrix)
}

