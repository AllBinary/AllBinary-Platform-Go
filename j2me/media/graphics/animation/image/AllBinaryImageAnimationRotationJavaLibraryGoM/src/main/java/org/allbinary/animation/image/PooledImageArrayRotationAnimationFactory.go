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

//TWB - Adjustments should be done in the resource creation and not at the animation level
type PooledImageArrayRotationAnimationFactory struct {
    AllBinaryImageRotationAnimationInfo any
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface


func (this *PooledImageArrayRotationAnimationFactory) initFields() {
    this.AllBinaryImageRotationAnimationInfo = NullUtil_GetInstance().NULL_OBJECT
    
}


// throws Exception
func PooledImageArrayRotationAnimationFactory_CreateFactory(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *PooledImageArrayRotationAnimationFactory {
        var totalAngle int = int(AngleFactory_GetInstance().TOTAL_ANGLE); _ = totalAngle
        var angleIncrement int = (totalAngle / GameConfigurationCentral_GetInstance().GetGameControlFidelity()); _ = angleIncrement
        return NewPooledImageArrayRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), -(image.GetWidth() >> 2), -(image.GetHeight() >> 2), angleIncrement, animationBehaviorFactory)
}


// throws Exception
func PooledImageArrayRotationAnimationFactory_CreateDXY(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *PooledImageArrayRotationAnimationFactory {
        var totalAngle int = int(AngleFactory_GetInstance().TOTAL_ANGLE); _ = totalAngle
        var angleIncrement int = (totalAngle / GameConfigurationCentral_GetInstance().GetGameControlFidelity()); _ = angleIncrement
        return NewPooledImageArrayRotationAnimationFactory(image, image.GetWidth(), image.GetHeight(), dx, dy, angleIncrement, animationBehaviorFactory)
}


func NewPooledImageArrayRotationAnimationFactory(image IImage, width int, height int, dx int, dy int, angleIncrement int, animationBehaviorFactory IAnimationBehaviorFactory) *PooledImageArrayRotationAnimationFactory {
    this := &PooledImageArrayRotationAnimationFactory{}
    this.initFields()
        this.AnimationBehaviorFactory = animationBehaviorFactory
        var totalAngle int = int(AngleFactory_GetInstance().TOTAL_ANGLE); _ = totalAngle
        var imageArray []IImage = ImageToRotationImageArrayUtil_GetInstance().Generate(image, angleIncrement, totalAngle); _ = imageArray
        this.AllBinaryImageRotationAnimationInfo = NewImageArrayRotationAnimationInfo(imageArray, angleIncrement, totalAngle, dx, dy)
    return this
}


// throws Exception
func (this *PooledImageArrayRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        return AdjustedImageArrayRotationAnimation_CreateAdjustedImageArrayRotationAnimationInfo(this.AllBinaryImageRotationAnimationInfo, this.AnimationBehaviorFactory.GetOrCreateInstance())
}


func (this *PooledImageArrayRotationAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

