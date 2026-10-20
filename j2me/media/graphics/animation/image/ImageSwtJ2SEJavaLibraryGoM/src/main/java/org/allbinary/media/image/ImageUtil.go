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

type ImageUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    IIOIMAGE_POOL_NAME string
    CREATE_BUFFERED_IMAGE string
}

var ImageUtil_Instance *ImageUtil = NewImageUtil()

func (this *ImageUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.IIOIMAGE_POOL_NAME = "IIOIMAGE_POOL_NAME"
    this.CREATE_BUFFERED_IMAGE = "createBufferedImage"
}


func ImageUtil_GetInstance() *ImageUtil {
        return ImageUtil_Instance
}


func NewImageUtil() *ImageUtil {
    this := &ImageUtil{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *ImageUtil) Create(width int, height int) IImage {
        var display *Display = BareMain_Shell.GetDisplay(); _ = display
        return NewImage(display, width, height)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage(bufferedImageArray []IImage, percent int, scale bool) []IImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []IImage = make([]IImage, size); _ = scaledBufferedImageArray
        var bufferedImage IImage; _ = bufferedImage
        var imageData *ImageData; _ = imageData
        for index := int(0); index < size; index++ {
            bufferedImage = bufferedImageArray[index]
            imageData = bufferedImage.GetImageData()
            var newWidth int = (imageData.Width * percent / 100); _ = newWidth
            var newHeight int = (imageData.Height * percent / 100); _ = newHeight
            scaledBufferedImageArray[index] = this.CreateBufferedImage5(bufferedImage, newWidth, newHeight, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage2(bufferedImageArray []IImage, percent float32, scale bool) []IImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []IImage = make([]IImage, size); _ = scaledBufferedImageArray
        var bufferedImage IImage; _ = bufferedImage
        var imageData *ImageData; _ = imageData
        for index := int(0); index < size; index++ {
            bufferedImage = bufferedImageArray[index]
            imageData = bufferedImage.GetImageData()
            var newWidth int = int((float32(imageData.Width) * percent)); _ = newWidth
            var newHeight int = int((float32(imageData.Height) * percent)); _ = newHeight
            scaledBufferedImageArray[index] = this.CreateBufferedImage5(bufferedImage, newWidth, newHeight, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage3(bufferedImageArray []IImage, width int, height int, scale bool) []IImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []IImage = make([]IImage, size); _ = scaledBufferedImageArray
        for index := int(0); index < size; index++ {
            scaledBufferedImageArray[index] = this.CreateBufferedImage5(bufferedImageArray[index], width, height, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage4(bufferedImage IImage, newWidth int, newHeight int) IImage {
        return this.CreateBufferedImage5(bufferedImage, newWidth, newHeight, true)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage5(bufferedImage IImage, newWidth int, newHeight int, scale bool) IImage {
        return this.CreateBufferedImage6(bufferedImage, newWidth, newHeight, scale, false)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage6(bufferedImage IImage, newWidth int, newHeight int, scale bool, allowTranslate bool) IImage {
        var imageData *ImageData = bufferedImage.GetImageData(); _ = imageData
        var width float32 = float32(imageData.Width); _ = width
        var height float32 = float32(imageData.Height); _ = height
        var d_newWidth float32 = float32(newWidth); _ = d_newWidth
        var d_newHeight float32 = float32(newHeight); _ = d_newHeight
        var widthRatio float32 = d_newWidth / width; _ = widthRatio
        var heightRatio float32 = d_newHeight / height; _ = heightRatio
        var ratioX float32 = 1.0; _ = ratioX
        var ratioY float32 = 1.0; _ = ratioY
        if scale {
            ratioX = widthRatio
            ratioY = heightRatio
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Appendfloat(width).Append(this.CommonSeps.FORWARD_SLASH).Appendfloat(height).Append(this.CommonSeps.COLON).Appendint(newWidth).Append(this.CommonSeps.FORWARD_SLASH).Appendint(newHeight).Append(this.CommonSeps.COLON).Appendfloat(widthRatio).Append(this.CommonSeps.FORWARD_SLASH).Appendfloat(heightRatio)), this, this.CREATE_BUFFERED_IMAGE)
        var dx float32 = 0.0; _ = dx
        var dy float32 = 0.0; _ = dy
        if !scale && allowTranslate {
            dx = (float32(newWidth) - width) / 2.0
            dy = (float32(newHeight) - height) / 2.0
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Translate dx: ").Appendfloat(dx).Append(" dy: ").Appendfloat(dy)), this, this.CREATE_BUFFERED_IMAGE)
        }
        var newBufferedImage IImage = NewImage(BareMain_Shell.GetDisplay(), newWidth, newHeight); _ = newBufferedImage
        var gc *GC = NewGC(newBufferedImage); _ = gc
        var affineTransform *Transform = NewTransform(gc.GetDevice()); _ = affineTransform
        affineTransform.Scale(ratioX, ratioY)
        affineTransform.Translate(dx, dy)
        gc.SetTransform(affineTransform)
        gc.SetAntialias(SWT_ON)
        gc.SetInterpolation(SWT_HIGH)
        gc.DrawImage(bufferedImage, 0, 0, bufferedImage.GetBounds().Width, bufferedImage.GetBounds().Height, 0, 0, newWidth, newHeight)
        gc.Dispose()
        return newBufferedImage
}


func (this *ImageUtil) DrawTransformedImage(image IImage, parent IComposite, dx int, dy int, width int, height int) ICanvas {
        var imageCanvas ICanvas = NewCanvas(parent, SWT_NONE); _ = imageCanvas
        imageCanvas.AddPaintListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageUtil) PaintControl(event *PaintEvent) {
                var gc *GC = event.Gc; _ = gc
                var affineTransform *Transform = NewTransform(gc.GetDevice()); _ = affineTransform
                affineTransform.Translate(dx, dy)
                gc.SetAntialias(SWT_ON)
                gc.SetInterpolation(SWT_HIGH)
                gc.DrawImage(image, 0, 0, image.GetBounds().Width, image.GetBounds().Height, 0, 0, width, height)
        }
*/ NewPaintListener())
        return imageCanvas
}


func (this *ImageUtil) ToString(bufferedImage IImage) string {
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var imageData *ImageData = bufferedImage.GetImageData(); _ = imageData
        stringBuffer.Append(" org.eclipse.swt.graphics.Image: ")
        stringBuffer.Append(commonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(imageData.Width)
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(imageData.Height)
        stringBuffer.Append(" Type: ")
        stringBuffer.Appendint(int(imageData.Type))
        return fmt.Sprint(stringBuffer)
}

