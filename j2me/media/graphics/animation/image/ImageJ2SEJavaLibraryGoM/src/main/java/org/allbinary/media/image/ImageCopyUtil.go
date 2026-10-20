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
    ImageUtil *ImageUtil
    ImageCreationUtil *ImageCreationUtil
    Anchor int
}

var ImageCopyUtil_Instance *ImageCopyUtil = NewImageCopyUtil()

func (this *ImageCopyUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ImageUtil = ImageUtil_GetInstance()
    this.ImageCreationUtil = ImageCreationUtil_GetInstance()
    this.Anchor = Anchor_TOP_LEFT
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
        return this.CreateImage(originalImage)
}


// throws Exception
func (this *ImageCopyUtil) CreateImage(originalImage IImage) IImage {
        var image IImage = this.ImageCreationUtil.CreateImageWH(originalImage.GetWidth(), originalImage.GetHeight()); _ = image
        if image.IsMutable() {
            image.GetGraphics().DrawImage(originalImage, 0, 0, this.Anchor)
            return image
        } else {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
}


// throws Exception
func (this *ImageCopyUtil) CreateImageScale(originalImage IImage, canvasScale float32, resize bool) IImage {
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
        var originalBufferedImage *BufferedImage = this.ImageUtil.GetBufferedImage(originalImage); _ = originalBufferedImage
        var bufferedImage *BufferedImage = this.ImageUtil.CreateBufferedImageWithLargerCanvas(originalBufferedImage, newWidth, newHeight); _ = bufferedImage
        var image *J2SEImmutableImage = NewJ2SEImmutableImage(bufferedImage); _ = image
        return image
}

