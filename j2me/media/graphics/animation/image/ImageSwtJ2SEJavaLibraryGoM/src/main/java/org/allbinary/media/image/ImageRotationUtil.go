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
    ImageSwtRotationUtil *ImageSwtRotationUtil
    Features *Features
}

var ImageRotationUtil_Instance *ImageRotationUtil = NewImageRotationUtil()

func (this *ImageRotationUtil) initFields() {
    this.ImageSwtRotationUtil = ImageSwtRotationUtil_GetInstance()
    this.Features = Features_GetInstance()
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
        var originalImage2 IImage = originalImage; _ = originalImage2
        if this.Features.IsFeature(OpenGLFeatureFactory_GetInstance().OPENGL) {
            if image.GetType() >= OpenGLESImage_TYPE {
                var openGLESImage IOpenGLESImage = (any(image).(IOpenGLESImage)); _ = openGLESImage
                openGLESImage.OpenGLESImage_().OpenGLESImageProperties.Angle = totalAngle
                return
            } else {
                PreLogUtil_Put("OpenGL but image is: " + fmt.Sprint(image) + " type: " + fmt.Sprint(image.GetType()), this, "rotateImage")
            }
        }
        var originalSwtImage IImage = (any(originalImage2.GetImage()).(IImage)); _ = originalSwtImage
        var gc *GC = nil; _ = gc
        var copySwtImage IImage = nil; _ = copySwtImage
        if image.IsMutable() {
            var swtImage *SwtMutableImage = any(image).(*SwtMutableImage); _ = swtImage
            copySwtImage = any(swtImage.GetImage()).(IImage)
            gc = swtImage.GetGc()
            this.ImageSwtRotationUtil.RotateImage(originalSwtImage, copySwtImage, gc, totalAngle)
        } else {
            PreLogUtil_Put("No rotation for: " + fmt.Sprint(image), this, "rotateImage")
        }
}


// throws Exception
func (this *ImageRotationUtil) CreateRotatedImage(originalImage IImage, rotationInDegrees int) IImage {
        return originalImage
}

