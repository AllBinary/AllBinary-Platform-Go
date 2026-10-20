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

type SpriteIndexedAnimationFactory struct {
    BaseImageAnimationFactory
    AnimationFactorySpriteScaleUtil *AnimationFactorySpriteScaleUtil
    BasicColorArray []*BasicColor
}

// Implements ProceduralAnimationInterfaceFactoryInterface


func (this *SpriteIndexedAnimationFactory) initFields() {
    this.AnimationFactorySpriteScaleUtil = AnimationFactorySpriteScaleUtil_GetInstance()
    this.BasicColorArray = BasicColorUtil_GetInstance().ZERO_ARRAY
}


// throws Exception
func SpriteIndexedAnimationFactory_CreateCWHDXY(image IImage, basicColorArray []*BasicColor, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *SpriteIndexedAnimationFactory {
        return NewSpriteIndexedAnimationFactory(image, basicColorArray, PrimitiveIntUtil_GetArrayInstance(), width, height, dx, dy, animationBehaviorFactory)
}


// throws Exception
func SpriteIndexedAnimationFactory_CreateCWH(image IImage, basicColorArray []*BasicColor, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) *SpriteIndexedAnimationFactory {
        var spriteIndexedAnimationFactory *SpriteIndexedAnimationFactory = NewSpriteIndexedAnimationFactory(image, basicColorArray, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory); _ = spriteIndexedAnimationFactory
        spriteIndexedAnimationFactory.InitHWH()
        return spriteIndexedAnimationFactory
}


// throws Exception
func SpriteIndexedAnimationFactory_CreateWH(image IImage, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) *SpriteIndexedAnimationFactory {
        return NewSpriteIndexedAnimationFactory(image, BasicColorUtil_GetInstance().ZERO_ARRAY, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
}


func NewSpriteIndexedAnimationFactory(image IImage, basicColorArray []*BasicColor, sequenceArray []int, width int, height int, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *SpriteIndexedAnimationFactory {
    this := &SpriteIndexedAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, sequenceArray, width, height, dx, dy, animationBehaviorFactory)
        this.initFields()
        this.BasicColorArray = basicColorArray
    return this
}


func (this *SpriteIndexedAnimationFactory) InitHWH() {
        this.AnimationFactoryInitializationVisitor.Dx = -(this.AnimationFactoryInitializationVisitor.Width >> 2)
        this.AnimationFactoryInitializationVisitor.Dy = -(this.AnimationFactoryInitializationVisitor.Height >> 2)
        if AndroidUtil_IsAndroid() {
            this.AnimationFactoryInitializationVisitor.Dx += 3
        }
}


// throws Exception
func (this *SpriteIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        var sprite *Sprite = this.AnimationFactorySpriteScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = sprite
        if this.AnimationFactoryInitializationVisitor.Dx != 0 || this.AnimationFactoryInitializationVisitor.Dy != 0 {
            return NewAdjustedSpriteIndexedAnimation(sprite, this.GetImage(), this.BasicColorArray, this.AnimationFactoryInitializationVisitor.Dx, this.AnimationFactoryInitializationVisitor.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewSpriteIndexedAnimation(sprite, this.GetImage(), this.BasicColorArray, this.AnimationBehaviorFactory.GetOrCreateInstance())
        }
}


// throws Exception
func (this *SpriteIndexedAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.GetInstance(0)
}

