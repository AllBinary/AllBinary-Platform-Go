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

type ColorCompositeSpriteIndexedAnimationFactory struct {
    BaseImageAnimationFactory
    AnimationFactorySpriteScaleUtil *AnimationFactorySpriteScaleUtil
    BasicColorArray []*BasicColor
}

// Implements ProceduralAnimationInterfaceFactoryInterface


func (this *ColorCompositeSpriteIndexedAnimationFactory) initFields() {
    this.AnimationFactorySpriteScaleUtil = AnimationFactorySpriteScaleUtil_GetInstance()
    
}


func NewColorCompositeSpriteIndexedAnimationFactory(image IImage, basicColorArray []*BasicColor, width int, height int, animationBehaviorFactory IAnimationBehaviorFactory) *ColorCompositeSpriteIndexedAnimationFactory {
    this := &ColorCompositeSpriteIndexedAnimationFactory{}
        this.BaseImageAnimationFactory = *NewBaseImageAnimationFactory(image, PrimitiveIntUtil_GetArrayInstance(), width, height, 0, 0, animationBehaviorFactory)
        this.initFields()
        this.BasicColorArray = basicColorArray
    return this
}


// throws Exception
func (this *ColorCompositeSpriteIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        var sprite *Sprite = this.AnimationFactorySpriteScaleUtil.CreateImage(this.GetImage(), this.AnimationFactoryInitializationVisitor.Width, this.AnimationFactoryInitializationVisitor.Height, this.ScaleProperties.ScaleWidth, this.ScaleProperties.ScaleHeight); _ = sprite
        return NewSpriteIndexedAnimation(sprite, this.GetImage(), this.BasicColorArray, this.AnimationBehaviorFactory.GetOrCreateInstance())
}


// throws Exception
func (this *ColorCompositeSpriteIndexedAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.GetInstance(0)
}

