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

type AllBinaryJ2SEImageRotationAnimationFactory struct {
    BaseImageAnimationFactory
    AngleIncrementP int16
    ResizeCanvasForRotation bool
}

// The Java class AllBinaryJ2SEImageRotationAnimationFactory has subclasses so its variables are this interface
type IAllBinaryJ2SEImageRotationAnimationFactory interface {
    IBaseImageAnimationFactory
    AllBinaryJ2SEImageRotationAnimationFactory_() *AllBinaryJ2SEImageRotationAnimationFactory
    GetCanvasImage() IImage
    GetInstance(instanceId int) IAnimation
    GetAngleIncrement() int16
}

func (this *AllBinaryJ2SEImageRotationAnimationFactory) AllBinaryJ2SEImageRotationAnimationFactory_() *AllBinaryJ2SEImageRotationAnimationFactory {
    return this
}


func (this *AllBinaryJ2SEImageRotationAnimationFactory) initFields() {
    
    
}


// throws Exception
func AllBinaryJ2SEImageRotationAnimationFactory_CreateDXY(image IImage, width int, height int, dx int, dy int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory, resizeCanvasForRotation bool) IAllBinaryJ2SEImageRotationAnimationFactory {
        var imageRotationAnimationFactory IAllBinaryJ2SEImageRotationAnimationFactory = NewAllBinaryJ2SEImageRotationAnimationFactory(image, width, height, angleIncrement, animationBehaviorFactory, resizeCanvasForRotation); _ = imageRotationAnimationFactory
        imageRotationAnimationFactory.AllBinaryJ2SEImageRotationAnimationFactory_().AnimationFactoryInitializationVisitor.Dx = dx
        imageRotationAnimationFactory.AllBinaryJ2SEImageRotationAnimationFactory_().AnimationFactoryInitializationVisitor.Dy = dy
        imageRotationAnimationFactory.AllBinaryJ2SEImageRotationAnimationFactory_().AnimationFactoryInitializationVisitor.OriginalDx = dx
        imageRotationAnimationFactory.AllBinaryJ2SEImageRotationAnimationFactory_().AnimationFactoryInitializationVisitor.OriginalDy = dy
        return imageRotationAnimationFactory
}


func NewAllBinaryJ2SEImageRotationAnimationFactory(image IImage, width int, height int, angleIncrement int16, animationBehaviorFactory IAnimationBehaviorFactory, resizeCanvasForRotation bool) *AllBinaryJ2SEImageRotationAnimationFactory {
    this := &AllBinaryJ2SEImageRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.AngleIncrementP = angleIncrement
        this.ResizeCanvasForRotation = resizeCanvasForRotation
    return this
}


// throws Exception
func (this *AllBinaryJ2SEImageRotationAnimationFactory) GetCanvasImage() IImage {
        var features *Features = Features_GetInstance(); _ = features
        if this.ResizeCanvasForRotation && !features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            return ImageCopyUtil_GetInstance().CreateImageScale(this.GetImage(), 1.44, false)
        } else {
            return this.GetImage()
        }
}


// throws Exception
func (this *AllBinaryJ2SEImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        var imageCopyUtil *ImageCopyUtil = ImageCopyUtil_GetInstance(); _ = imageCopyUtil
        var canvasImage IImage = this.GetCanvasImage(); _ = canvasImage
        var scaledImage IImage = this.AnimationFactoryImageScaleUtil.CreateImage(canvasImage, this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = scaledImage
        var openGLUtil *OpenGLUtil = OpenGLUtil_GetInstance(); _ = openGLUtil
        scaledImage = openGLUtil.AddImage(scaledImage)
        var copyOfScaledImage IImage = imageCopyUtil.CreateImageForRotation(scaledImage); _ = copyOfScaledImage
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAllBinaryAdjustedJ2SEImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.AngleIncrementP), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewAllBinaryJ2SEImageRotationAnimation(scaledImage, copyOfScaledImage, AngleInfo_GetInstance(this.AngleIncrementP), AngleFactory_GetInstance().TOTAL_ANGLE, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *AllBinaryJ2SEImageRotationAnimationFactory) GetAngleIncrement() int16 {
        return this.AngleIncrementP
}

