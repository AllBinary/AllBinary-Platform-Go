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

type ImageAnimationFactory struct {
    BaseImageAnimationFactory
}



// throws Exception
func ImageAnimationFactory_CreateFactory(image IImage) *ImageAnimationFactory {
        return NewImageAnimationFactory(image, image.GetWidth(), image.GetHeight(), 0, 0, AnimationBehaviorFactory_GetInstance())
}


func NewImageAnimationFactory(image IImage, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *ImageAnimationFactory {
    this := &ImageAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, dx, dy, animationBehaviorFactory)
    return this
}


// throws Exception
func (this *ImageAnimationFactory) GetInstance(instanceId int) IAnimation {
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        var openGLUtil *OpenGLUtil = OpenGLUtil_GetInstance(); _ = openGLUtil
        scaledImage = openGLUtil.AddImage(scaledImage)
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAdjustedImageAnimation(scaledImage, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewImageAnimation(scaledImage, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}

