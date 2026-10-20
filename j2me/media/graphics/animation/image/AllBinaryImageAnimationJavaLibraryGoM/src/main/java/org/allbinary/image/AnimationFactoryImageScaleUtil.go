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

type AnimationFactoryImageScaleUtil struct {
    ImageScaleUtil *ImageJ2MEScaleUtil
}

var AnimationFactoryImageScaleUtil_Instance *AnimationFactoryImageScaleUtil = NewAnimationFactoryImageScaleUtil()

func (this *AnimationFactoryImageScaleUtil) initFields() {
    this.ImageScaleUtil = ImageJ2MEScaleUtil_GetInstance()
}

func NewAnimationFactoryImageScaleUtil() *AnimationFactoryImageScaleUtil {
    this := &AnimationFactoryImageScaleUtil{}
    this.initFields()
    return this
}


func AnimationFactoryImageScaleUtil_GetInstance() *AnimationFactoryImageScaleUtil {
        return AnimationFactoryImageScaleUtil_Instance
}


// throws Exception
func (this *AnimationFactoryImageScaleUtil) CreateImage(image IImage, width int, height int, scaleWidth int, scaleHeight int) IImage {
        var scaledImage IImage; _ = scaledImage
        if scaleWidth != 0 && scaleHeight != 0 {
            var scaleX float32 = (float32(scaleWidth)) / (float32(width)); _ = scaleX
            var scaleY float32 = (float32(scaleHeight)) / (float32(height)); _ = scaleY
            if (scaleX == 1.0 && scaleY == 1.0) || (scaleX == 0.0 || scaleY == 0.0) {
                scaledImage = image
            } else {
                var scaledWidth int = int((scaleX * float32(width))); _ = scaledWidth
                var scaledHeight int = int((scaleY * float32(height))); _ = scaledHeight
                scaledImage = this.ImageScaleUtil.Scale(image, scaledWidth, scaledHeight)
            }
        } else {
            scaledImage = image
        }
        return scaledImage
}


// throws Exception
func (this *AnimationFactoryImageScaleUtil) ProcessAdjust(baseImageAnimationFactory IBaseImageAnimationFactory) {
        var scaleProperties *ScaleProperties = baseImageAnimationFactory.GetScalePropertiesP(); _ = scaleProperties
        if scaleProperties.ScaleWidth != 0 && scaleProperties.ScaleHeight != 0 {
            var animationFactoryInitializationVisitor *AnimationFactoryInitializationVisitor = baseImageAnimationFactory.GetAnimationFactoryInitializationVisitorP(); _ = animationFactoryInitializationVisitor
            var scaleX float32 = (float32(scaleProperties.ScaleWidth)) / (float32(animationFactoryInitializationVisitor.Width)); _ = scaleX
            var scaleY float32 = (float32(scaleProperties.ScaleHeight)) / (float32(animationFactoryInitializationVisitor.Height)); _ = scaleY
            if (scaleX == 1.0 && scaleY == 1.0) || (scaleX == 0.0 || scaleY == 0.0) {
            } else {
                animationFactoryInitializationVisitor.Dx = int((float32(animationFactoryInitializationVisitor.OriginalDx) * scaleX))
                animationFactoryInitializationVisitor.Dy = int((float32(animationFactoryInitializationVisitor.OriginalDy) * scaleY))
            }
        } else {
        }
}

