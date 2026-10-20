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

type LazyImageRotationAnimationFactory struct {
    LayoutIndex int
    AnimationInterfaceFactoryInterface IBaseImageAnimationFactory
    ScaleProperties *ScaleProperties
}

// Implements AnimationInterfaceFactoryInterface


func (this *LazyImageRotationAnimationFactory) initFields() {
    
    
    this.ScaleProperties = ScaleProperties_Instance
}


func NewLazyImageRotationAnimationFactory(layoutIndex int, associatedLazyAnimationId int, animationInterfaceFactoryInterface IBaseImageAnimationFactory) *LazyImageRotationAnimationFactory {
    this := &LazyImageRotationAnimationFactory{}
    this.initFields()
        this.LayoutIndex = layoutIndex
        this.AnimationInterfaceFactoryInterface = animationInterfaceFactoryInterface
        ImageCacheFactory_GetInstance().ImageCache_().HasAnyLazyAnimationFactories = true
        if this.AnimationInterfaceFactoryInterface.BaseImageAnimationFactory_().AnimationBehaviorFactory == AnimationBehaviorFactory_GetInstance() {
            ForcedLogUtil_Log("Using default AnimationBehaviorFactory with IndexedAnimationFactory", this)
        }
    return this
}


// throws Exception
func (this *LazyImageRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        if this.AnimationInterfaceFactoryInterface.GetImage().IsReady() {
            this.AnimationInterfaceFactoryInterface.SetInitialScale(this.ScaleProperties)
            return this.AnimationInterfaceFactoryInterface.GetInstance(instanceId)
        } else {
            return NewLazyImageRotationAnimation(this.LayoutIndex, instanceId, this.ScaleProperties, this.AnimationInterfaceFactoryInterface, this.AnimationInterfaceFactoryInterface.BaseImageAnimationFactory_().AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


func (this *LazyImageRotationAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
        this.ScaleProperties = scaleProperties
}

