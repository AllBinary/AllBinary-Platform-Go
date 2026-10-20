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

type OneRowSpriteIndexedAnimationFactory struct {
    BaseImageAnimationFactory
    AnimationFactorySpriteScaleUtil *AnimationFactorySpriteScaleUtil
}


func (this *OneRowSpriteIndexedAnimationFactory) initFields() {
    this.AnimationFactorySpriteScaleUtil = AnimationFactorySpriteScaleUtil_GetInstance()
}


// throws Exception
func OneRowSpriteIndexedAnimationFactory_CreateFactoryDX(image IImage, dx int, animationBehaviorFactory IAnimationBehaviorFactory) *OneRowSpriteIndexedAnimationFactory {
        var oneRowSpriteIndexedAnimationFactory *OneRowSpriteIndexedAnimationFactory = NewOneRowSpriteIndexedAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), image.GetHeight(), image.GetHeight(), 0, 0, animationBehaviorFactory); _ = oneRowSpriteIndexedAnimationFactory
        oneRowSpriteIndexedAnimationFactory.InitW(dx)
        return oneRowSpriteIndexedAnimationFactory
}


// throws Exception
func OneRowSpriteIndexedAnimationFactory_CreateFactoryDXY(image IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *OneRowSpriteIndexedAnimationFactory {
        var oneRowSpriteIndexedAnimationFactory *OneRowSpriteIndexedAnimationFactory = NewOneRowSpriteIndexedAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), image.GetHeight(), image.GetHeight(), 0, 0, animationBehaviorFactory); _ = oneRowSpriteIndexedAnimationFactory
        oneRowSpriteIndexedAnimationFactory.Init(dx, dy)
        return oneRowSpriteIndexedAnimationFactory
}


// throws Exception
func OneRowSpriteIndexedAnimationFactory_CreateFactoryWH(image IImage, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) *OneRowSpriteIndexedAnimationFactory {
        return NewOneRowSpriteIndexedAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
}


// throws Exception
func OneRowSpriteIndexedAnimationFactory_CreateFactory(image IImage, animationBehaviorFactory IAnimationBehaviorFactory) *OneRowSpriteIndexedAnimationFactory {
        return NewOneRowSpriteIndexedAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), image.GetHeight(), image.GetHeight(), 0, 0, animationBehaviorFactory)
}


func NewOneRowSpriteIndexedAnimationFactory(image IImage, sequenceArray []int, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *OneRowSpriteIndexedAnimationFactory {
    this := &OneRowSpriteIndexedAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, sequenceArray, width, height, dx, dy, animationBehaviorFactory)
        this.initFields()
    return this
}


func (this *OneRowSpriteIndexedAnimationFactory) Init(dx int, dy int) {
        this.AnimationFactoryInitializationVisitor.Dx = dx
        this.AnimationFactoryInitializationVisitor.Dy = dy
}


func (this *OneRowSpriteIndexedAnimationFactory) InitW(dx int) {
        this.Init(dx, 0)
        this.AnimationFactoryInitializationVisitor.Dx += -(this.AnimationFactoryInitializationVisitor.Width >> 2)
}


// throws Exception
func (this *OneRowSpriteIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        var sprite *Sprite = this.AnimationFactorySpriteScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = sprite
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            return NewAdjustedSpriteIndexedAnimation(sprite, this.GetImage(), BasicColorUtil_GetInstance().ZERO_ARRAY, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewSpriteIndexedAnimation(sprite, this.GetImage(), BasicColorUtil_GetInstance().ZERO_ARRAY, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}

