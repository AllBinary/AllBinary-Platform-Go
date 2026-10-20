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

type AllBinaryAndroidImageRotationAnimationFactory struct {
    BaseImageAnimationFactory
    AngleIncrement int16
    ResizeCanvasForRotation bool
}

// The Java class AllBinaryAndroidImageRotationAnimationFactory has subclasses so its variables are this interface
type IAllBinaryAndroidImageRotationAnimationFactory interface {
    IBaseImageAnimationFactory
    AllBinaryAndroidImageRotationAnimationFactory_() *AllBinaryAndroidImageRotationAnimationFactory
    Init(dx int, dy int)
    GetCanvasImage() IImage
    GetInstance(instanceId int) IAnimation
}

func (this *AllBinaryAndroidImageRotationAnimationFactory) AllBinaryAndroidImageRotationAnimationFactory_() *AllBinaryAndroidImageRotationAnimationFactory {
    return this
}


func (this *AllBinaryAndroidImageRotationAnimationFactory) initFields() {
    
    
}


// throws Exception
func AllBinaryAndroidImageRotationAnimationFactory_CreateFactory(image IImage, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) IAllBinaryAndroidImageRotationAnimationFactory {
        return NewAllBinaryAndroidImageRotationAnimationFactory(image, width, height, int16((int(AngleFactory_GetInstance().TOTAL_ANGLE) / GameConfigurationCentral_GetInstance().GetGameControlFidelity())), animationBehaviorFactory, false)
}


// throws Exception
func AllBinaryAndroidImageRotationAnimationFactory_CreateFactoryA(image IImage, width int, height int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory) IAllBinaryAndroidImageRotationAnimationFactory {
        return NewAllBinaryAndroidImageRotationAnimationFactory(image, width, height, angleIncrement, animationBehaviorFactory, false)
}


// throws Exception
func AllBinaryAndroidImageRotationAnimationFactory_CreateFactoryU(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) IAllBinaryAndroidImageRotationAnimationFactory {
        var androidImageRotationAnimationFactory IAllBinaryAndroidImageRotationAnimationFactory = NewAllBinaryAndroidImageRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), int16((int(AngleFactory_GetInstance().TOTAL_ANGLE) / GameConfigurationCentral_GetInstance().GetGameControlFidelity())), animationBehaviorFactory, false); _ = androidImageRotationAnimationFactory
        androidImageRotationAnimationFactory.Init(-(image.GetWidth() >> 2), -(image.GetHeight() >> 2))
        return androidImageRotationAnimationFactory
}


// throws Exception
func AllBinaryAndroidImageRotationAnimationFactory_CreateFactoryDXYU(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) IAllBinaryAndroidImageRotationAnimationFactory {
        var androidImageRotationAnimationFactory IAllBinaryAndroidImageRotationAnimationFactory = NewAllBinaryAndroidImageRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), int16((int(AngleFactory_GetInstance().TOTAL_ANGLE) / GameConfigurationCentral_GetInstance().GetGameControlFidelity())), animationBehaviorFactory, false); _ = androidImageRotationAnimationFactory
        androidImageRotationAnimationFactory.Init(dx, dy)
        return androidImageRotationAnimationFactory
}


// throws Exception
func AllBinaryAndroidImageRotationAnimationFactory_CreateFactoryADXY(image IImage, width int, height int, dx int, dy int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory, resizeCanvasForRotation bool) IAllBinaryAndroidImageRotationAnimationFactory {
        var androidImageRotationAnimationFactory IAllBinaryAndroidImageRotationAnimationFactory = NewAllBinaryAndroidImageRotationAnimationFactory(image, width, height, angleIncrement, animationBehaviorFactory, resizeCanvasForRotation); _ = androidImageRotationAnimationFactory
        androidImageRotationAnimationFactory.Init(dx, dy)
        return androidImageRotationAnimationFactory
}


func (this *AllBinaryAndroidImageRotationAnimationFactory) Init(dx int, dy int) {
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
        this.AnimationFactoryInitializationVisitor.OriginalDx = dx
        this.AnimationFactoryInitializationVisitor.OriginalDy = dy
}


func NewAllBinaryAndroidImageRotationAnimationFactory(image IImage, width int, height int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory, resizeCanvasForRotation bool) *AllBinaryAndroidImageRotationAnimationFactory {
    this := &AllBinaryAndroidImageRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.AngleIncrement = angleIncrement
        this.ResizeCanvasForRotation = resizeCanvasForRotation
    return this
}


// throws Exception
func (this *AllBinaryAndroidImageRotationAnimationFactory) GetCanvasImage() IImage {
        var features *Features = Features_GetInstance(); _ = features
        if this.ResizeCanvasForRotation && !features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            return ImageCopyUtil_GetInstance().CreateImageScale(this.GetImage(), 1.44, false)
        } else {
            return this.GetImage()
        }
}


// throws Exception
func (this *AllBinaryAndroidImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        var canvasImage IImage = this.GetCanvasImage(); _ = canvasImage
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(canvasImage, this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        var copyOfScaledImage IImage = ImageCopyUtil_GetInstance().CreateImage(scaledImage); _ = copyOfScaledImage
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAllBinaryAdjustedAndroidImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.AngleIncrement), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewAllBinaryFlickerAndroidImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.AngleIncrement), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}

