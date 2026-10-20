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
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    IIOIMAGE_POOL_NAME string
    CREATE_BUFFERED_IMAGE string
}

var ImageUtil_Instance *ImageUtil = NewImageUtil()

func (this *ImageUtil) initFields() {
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
    return this
}


func (this *ImageUtil) GetDefaultConfiguration() *GraphicsConfiguration {
        var ge *GraphicsEnvironment = GraphicsEnvironment_GetLocalGraphicsEnvironment(); _ = ge
        var gd *GraphicsDevice = ge.GetDefaultScreenDevice(); _ = gd
        return gd.GetDefaultConfiguration()
}


func (this *ImageUtil) Create(width int, height int) *BufferedImage {
        var graphicsConfiguration *GraphicsConfiguration = this.GetDefaultConfiguration(); _ = graphicsConfiguration
        return graphicsConfiguration.CreateCompatibleImage(width, height, Transparency_TRANSLUCENT)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImageForResize(bufferedImageArray []*BufferedImage, percent int, scale bool) []*BufferedImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []*BufferedImage = make([]*BufferedImage, size); _ = scaledBufferedImageArray
        var bufferedImage *BufferedImage; _ = bufferedImage
        for index := int(0); index < size; index++ {
            bufferedImage = bufferedImageArray[index]
            var newWidth int = int((bufferedImage.GetWidth() * percent / 100)); _ = newWidth
            var newHeight int = int((bufferedImage.GetHeight() * percent / 100)); _ = newHeight
            scaledBufferedImageArray[index] = this.CreateBufferedImage(bufferedImage, newWidth, newHeight, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage2(bufferedImageArray []*BufferedImage, percent float32, scale bool) []*BufferedImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []*BufferedImage = make([]*BufferedImage, size); _ = scaledBufferedImageArray
        var bufferedImage *BufferedImage; _ = bufferedImage
        for index := int(0); index < size; index++ {
            bufferedImage = bufferedImageArray[index]
            var newWidth int = int((bufferedImage.GetWidth() * percent)); _ = newWidth
            var newHeight int = int((bufferedImage.GetHeight() * percent)); _ = newHeight
            scaledBufferedImageArray[index] = this.CreateBufferedImage(bufferedImage, newWidth, newHeight, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage3(bufferedImageArray []*BufferedImage, width int, height int, scale bool) []*BufferedImage {
        var size int = len(bufferedImageArray); _ = size
        var scaledBufferedImageArray []*BufferedImage = make([]*BufferedImage, size); _ = scaledBufferedImageArray
        for index := int(0); index < size; index++ {
            scaledBufferedImageArray[index] = this.CreateBufferedImage(bufferedImageArray[index], width, height, scale)
        }
        return scaledBufferedImageArray
}


// throws Exception
func (this *ImageUtil) CreateBufferedImageForSave(bufferedImage *BufferedImage, newWidth int, newHeight int) *BufferedImage {
        return this.CreateBufferedImage(bufferedImage, newWidth, newHeight, true)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImage(bufferedImage *BufferedImage, newWidth int, newHeight int, scale bool) *BufferedImage {
        return this.CreateBufferedImageAllowTranslate(bufferedImage, newWidth, newHeight, scale, false)
}


// throws Exception
func (this *ImageUtil) CreateBufferedImageAllowTranslate(bufferedImage *BufferedImage, newWidth int, newHeight int, scale bool, allowTranslate bool) *BufferedImage {
        var width float64 = bufferedImage.GetWidth(); _ = width
        var height float64 = bufferedImage.GetHeight(); _ = height
        var d_newWidth float64 = float64(newWidth); _ = d_newWidth
        var d_newHeight float64 = float64(newHeight); _ = d_newHeight
        var widthRatio float64 = d_newWidth / width; _ = widthRatio
        var heightRatio float64 = d_newHeight / height; _ = heightRatio
        var ratioX float64 = 1.0; _ = ratioX
        var ratioY float64 = 1.0; _ = ratioY
        if scale {
            ratioX = widthRatio
            ratioY = heightRatio
        }
        var affineTransform *AffineTransform = AffineTransform_GetScaleInstance(ratioX, ratioY); _ = affineTransform
        if !scale && allowTranslate {
            var dx float64 = (float64(newWidth) - width) / 2.0; _ = dx
            var dy float64 = (float64(newHeight) - height) / 2.0; _ = dy
            affineTransform.Translate(dx, dy)
        }
        var newBufferedImage *BufferedImage = NewBufferedImage(newWidth, newHeight, BufferedImage_TYPE_INT_ARGB_PRE); _ = newBufferedImage
        var graphics *Graphics2D = newBufferedImage.CreateGraphics(); _ = graphics
        graphics.DrawRenderedImage(bufferedImage, affineTransform)
        return newBufferedImage
}


// throws Exception
func (this *ImageUtil) CreateBufferedImageWithLargerCanvas(bufferedImage *BufferedImage, newWidth int, newHeight int) *BufferedImage {
        var width float64 = bufferedImage.GetWidth(); _ = width
        var height float64 = bufferedImage.GetHeight(); _ = height
        var d_newWidth float64 = float64(newWidth); _ = d_newWidth
        var d_newHeight float64 = float64(newHeight); _ = d_newHeight
        var widthRatio float64 = d_newWidth / width; _ = widthRatio
        var heightRatio float64 = d_newHeight / height; _ = heightRatio
        var dx2 int = int((float64(newWidth) - width)); _ = dx2
        var dy2 int = int((float64(newHeight) - height)); _ = dy2
        var dx int = dx2 / 2; _ = dx
        var dy int = dy2 / 2; _ = dy
        var dx4 int = dx2 / 4; _ = dx4
        var dy4 int = dy2 / 4; _ = dy4
        var newBufferedImage *BufferedImage = NewBufferedImage(newWidth, newHeight, BufferedImage_TYPE_INT_ARGB_PRE); _ = newBufferedImage
        var graphics *Graphics2D = newBufferedImage.CreateGraphics(); _ = graphics
        graphics.DrawImage(bufferedImage, dx4, dy4, nil)
        return newBufferedImage
}


func (this *ImageUtil) ConvertToBufferedImage(toolkitImage IImage) *BufferedImage {
        if toolkitImage == nil {
            return nil
        }
        var bufferedImage *BufferedImage = NewBufferedImage(toolkitImage.GetWidth(nil), toolkitImage.GetHeight(nil), BufferedImage_TYPE_INT_ARGB); _ = bufferedImage
        var g2d *Graphics2D = bufferedImage.CreateGraphics(); _ = g2d
        g2d.DrawImage(toolkitImage, 0, 0, nil)
        g2d.Dispose()
        return bufferedImage
}


func (this *ImageUtil) GetBufferedImage(image IImage) *BufferedImage {
        var bufferedImage *BufferedImage; _ = bufferedImage
        if image.IsMutable() {
            var j2seImage *J2SEMutableImage = any(image).(*J2SEMutableImage); _ = j2seImage
            bufferedImage = any(j2seImage.GetImage()).(*BufferedImage)
        } else {
            var j2seImage *J2SEImmutableImage = any(image).(*J2SEImmutableImage); _ = j2seImage
            bufferedImage = any(j2seImage.GetImage()).(*BufferedImage)
        }
        return bufferedImage
}


func (this *ImageUtil) ToString(bufferedImage *BufferedImage) string {
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(" BufferedImage -")
        stringBuffer.Append(commonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(int(bufferedImage.GetWidth()))
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(int(bufferedImage.GetHeight()))
        stringBuffer.Append(" Type: ")
        stringBuffer.Appendint(int(bufferedImage.GetType()))
        return fmt.Sprint(stringBuffer)
}

