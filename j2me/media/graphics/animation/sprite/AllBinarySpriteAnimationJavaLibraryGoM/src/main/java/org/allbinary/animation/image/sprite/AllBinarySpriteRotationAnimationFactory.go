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

type AllBinarySpriteRotationAnimationFactory struct {
    BaseImageAnimationFactory
    AnimationFactorySpriteScaleUtil *AnimationFactorySpriteScaleUtil
}

// Implements ProceduralAnimationInterfaceFactoryInterface


func (this *AllBinarySpriteRotationAnimationFactory) initFields() {
    this.AnimationFactorySpriteScaleUtil = AnimationFactorySpriteScaleUtil_GetInstance()
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateWHDY(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        var spriteRotationAnimationFactory *AllBinarySpriteRotationAnimationFactory = NewAllBinarySpriteRotationAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), (image.GetHeight() >> 2), (image.GetHeight() >> 2), 0, 0, animationBehaviorFactory); _ = spriteRotationAnimationFactory
        spriteRotationAnimationFactory.InitWH(dx, dy)
        return spriteRotationAnimationFactory
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateWH(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        return AllBinarySpriteRotationAnimationFactory_CreateWHDY(image, 0, 0, animationBehaviorFactory)
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateDXY(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        var spriteRotationAnimationFactory *AllBinarySpriteRotationAnimationFactory = NewAllBinarySpriteRotationAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), image.GetHeight(), image.GetHeight(), 0, 0, animationBehaviorFactory); _ = spriteRotationAnimationFactory
        spriteRotationAnimationFactory.Init(dx, dy)
        return spriteRotationAnimationFactory
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateDXYQ(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        var spriteRotationAnimationFactory *AllBinarySpriteRotationAnimationFactory = NewAllBinarySpriteRotationAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), (image.GetHeight() >> 2), (image.GetHeight() >> 2), 0, 0, animationBehaviorFactory); _ = spriteRotationAnimationFactory
        spriteRotationAnimationFactory.Init(dx, dy)
        return spriteRotationAnimationFactory
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateQ(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        return NewAllBinarySpriteRotationAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), (image.GetHeight() >> 2), (image.GetHeight() >> 2), 0, 0, animationBehaviorFactory)
}


// throws Exception
func AllBinarySpriteRotationAnimationFactory_CreateWHF(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
        return NewAllBinarySpriteRotationAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), image.GetHeight(), image.GetHeight(), 0, 0, animationBehaviorFactory)
}


func NewAllBinarySpriteRotationAnimationFactory(image IImage, sequenceArray []int, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *AllBinarySpriteRotationAnimationFactory {
    this := &AllBinarySpriteRotationAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, sequenceArray, width, height, dx, dy, animationBehaviorFactory)
        this.initFields()
        if animationBehaviorFactory == AnimationBehaviorFactory_GetInstance() {
            ForcedLogUtil_Log("Using default AnimationBehaviorFactory with IndexedAnimationFactory", this)
        }
    return this
}


func (this *AllBinarySpriteRotationAnimationFactory) Init(dx int, dy int) {
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
}


func (this *AllBinarySpriteRotationAnimationFactory) InitWH(dx int, dy int) {
        this.Init(dx, dy)
        this.AnimationFactoryInitializationVisitor.Dx += -this.AnimationFactoryInitializationVisitor.Width / 5
        this.AnimationFactoryInitializationVisitor.Dy += -this.AnimationFactoryInitializationVisitor.Height / 5
}


// throws Exception
func (this *AllBinarySpriteRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        var sprite *Sprite = this.AnimationFactorySpriteScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = sprite
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            return NewAllBinaryAdjustedSpriteRotationAnimation(sprite, this.GetImage(), this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewAllBinarySpriteRotationAnimation(sprite, this.GetImage(), this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


// throws Exception
func (this *AllBinarySpriteRotationAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.GetInstance(0)
}

