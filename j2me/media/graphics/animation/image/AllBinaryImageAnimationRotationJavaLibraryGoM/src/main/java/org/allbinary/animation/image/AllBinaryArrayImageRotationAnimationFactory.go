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

type AllBinaryArrayImageRotationAnimationFactory struct {
    BaseImageAnimationFactory
    ImageArray []IImage
    AngleIncrement int
}


func (this *AllBinaryArrayImageRotationAnimationFactory) initFields() {
    this.ImageArray = NullImage_NULL_IMAGE_ARRAY
    
}


// throws Exception
func AllBinaryArrayImageRotationAnimationFactory_CreateFactory(image IImage, dx int, dy int) *AllBinaryArrayImageRotationAnimationFactory {
        var arrayImageRotationAnimationFactory *AllBinaryArrayImageRotationAnimationFactory = NewAllBinaryArrayImageRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), int(AngleFactory_GetInstance().TOTAL_ANGLE) / GameConfigurationCentral_GetInstance().GetGameControlFidelity(), AnimationBehaviorFactory_GetInstance()); _ = arrayImageRotationAnimationFactory
        arrayImageRotationAnimationFactory.InitDXY(dx, dy)
        return arrayImageRotationAnimationFactory
}


// throws Exception
func AllBinaryArrayImageRotationAnimationFactory_CreateFactoryA(image IImage, dx int, dy int, angleIncrement int) *AllBinaryArrayImageRotationAnimationFactory {
        var arrayImageRotationAnimationFactory *AllBinaryArrayImageRotationAnimationFactory = NewAllBinaryArrayImageRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), angleIncrement, AnimationBehaviorFactory_GetInstance()); _ = arrayImageRotationAnimationFactory
        arrayImageRotationAnimationFactory.InitDXY(dx, dy)
        return arrayImageRotationAnimationFactory
}


func NewAllBinaryArrayImageRotationAnimationFactory(image IImage, width int, height int, angleIncrement int, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinaryArrayImageRotationAnimationFactory {
    this := &AllBinaryArrayImageRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.AngleIncrement = angleIncrement
        this.Init()
    return this
}


// throws Exception
func (this *AllBinaryArrayImageRotationAnimationFactory) Init() {
        this.SetImageArray(ImageToRotationImageArrayUtil_GetInstance().Generate(this.GetImage(), this.GetAngleIncrement(), int(AngleFactory_GetInstance().TOTAL_ANGLE)))
}


func (this *AllBinaryArrayImageRotationAnimationFactory) InitDXY(dx int, dy int) {
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
        this.AnimationFactoryInitializationVisitor.OriginalDx = dx
        this.AnimationFactoryInitializationVisitor.OriginalDy = dy
}


// throws Exception
func (this *AllBinaryArrayImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        var scaledImageArray []IImage = make([]IImage, len(this.ImageArray)); _ = scaledImageArray
        var size int = len(scaledImageArray); _ = size
        var image IImage = this.GetImage(); _ = image
        for index := int(0); index < size; index++ {
            scaledImageArray[index] = this.AnimationFactoryImageScaleUtil.CreateImage(image, image.GetWidth(), image.GetHeight(), this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight)
        }
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            this.AnimationFactoryImageScaleUtil.ProcessAdjust(this)
            return NewAdjustedImageArrayRotationAnimation(scaledImageArray, AngleInfo_GetInstance(int16(this.GetAngleIncrement())), int(AngleFactory_GetInstance().TOTAL_ANGLE), this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewImageArrayRotationAnimation(scaledImageArray, AngleInfo_GetInstance(int16(this.GetAngleIncrement())), int(AngleFactory_GetInstance().TOTAL_ANGLE), this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *AllBinaryArrayImageRotationAnimationFactory) SetImageArray(imageArray []IImage) {
        this.ImageArray = imageArray
}


func (this *AllBinaryArrayImageRotationAnimationFactory) GetAngleIncrement() int {
        return this.AngleIncrement
}

