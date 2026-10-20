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

type ImageRotationUtil struct {
    ImageJ2SERotationUtil *ImageJ2SERotationUtil
}

var ImageRotationUtil_Instance *ImageRotationUtil = NewImageRotationUtil()

func (this *ImageRotationUtil) initFields() {
    this.ImageJ2SERotationUtil = ImageJ2SERotationUtil_GetInstance()
}


func ImageRotationUtil_GetInstance() *ImageRotationUtil {
        return ImageRotationUtil_Instance
}


func NewImageRotationUtil() *ImageRotationUtil {
    this := &ImageRotationUtil{}
    this.initFields()
    return this
}


func (this *ImageRotationUtil) RotateImage(originalImage IImage, image IImage, totalAngle int) {
        if image.IsMutable() {
            var j2seImage *J2SEMutableImage = any(image).(*J2SEMutableImage); _ = j2seImage
            var newBufferedImage *BufferedImage = any(j2seImage.GetImage()).(*BufferedImage); _ = newBufferedImage
            var originalAwtImage IImage = nil; _ = originalAwtImage
            if originalImage.IsMutable() {
                var originalJ2SEImage *J2SEMutableImage = any(originalImage).(*J2SEMutableImage); _ = originalJ2SEImage
                originalAwtImage = any(originalJ2SEImage.GetImage()).(IImage)
            } else {
                var originalJ2SEImage *J2SEImmutableImage = any(originalImage).(*J2SEImmutableImage); _ = originalJ2SEImage
                originalAwtImage = any(originalJ2SEImage.GetImage()).(IImage)
            }
            this.ImageJ2SERotationUtil.RotateImage(originalAwtImage, newBufferedImage, totalAngle)
        } else {
            PreLogUtil_PutOE("Not Mutable", this, "getRotatedImage", errors.New(fmt.Sprint()))
        }
}


// throws Exception
func (this *ImageRotationUtil) CreateRotatedImage(originalImage IImage, rotationInDegrees int) IImage {
        var image IImage = ImageCreationUtil_GetInstance().CreateImageWH(originalImage.GetWidth(), originalImage.GetHeight()); _ = image
        if image.IsMutable() {
            var originalAwtImage IImage = nil; _ = originalAwtImage
            if originalImage.IsMutable() {
                var originalJ2SEImage *J2SEMutableImage = any(originalImage).(*J2SEMutableImage); _ = originalJ2SEImage
                originalAwtImage = any(originalJ2SEImage.GetImage()).(IImage)
            } else {
                var originalJ2SEImage *J2SEImmutableImage = any(originalImage).(*J2SEImmutableImage); _ = originalJ2SEImage
                originalAwtImage = any(originalJ2SEImage.GetImage()).(IImage)
            }
            var j2seImage *J2SEMutableImage = any(image).(*J2SEMutableImage); _ = j2seImage
            var newBufferedImage *BufferedImage = any(j2seImage.GetImage()).(*BufferedImage); _ = newBufferedImage
            var bufferedImage *BufferedImage = this.ImageJ2SERotationUtil.CreateRotatedImage(originalAwtImage, newBufferedImage, rotationInDegrees); _ = bufferedImage
            return image
        } else {
            panic(errors.New(fmt.Sprint("Not Mutable")))
        }
}

