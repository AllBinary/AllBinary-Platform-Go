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

type ImageArrayAnimationFactory struct {
    Dx int
    Dy int
    ImageArray []IImage
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface


func (this *ImageArrayAnimationFactory) initFields() {
    this.Dx = 0
    this.Dy = 0
    
    
}


func NewImageArrayAnimationFactory(imageArray []IImage, dx int, dy int, animationBehaviorFactory IAnimationBehaviorFactory) *ImageArrayAnimationFactory {
    this := &ImageArrayAnimationFactory{}
    this.initFields()
        this.ImageArray = imageArray
        this.AnimationBehaviorFactory = animationBehaviorFactory
        this.Dx = dx
        this.Dy = dy
    return this
}


// throws Exception
func (this *ImageArrayAnimationFactory) GetInstance(instanceId int) IAnimation {
        if this.Dx != 0 || this.Dy != 0 {
            return NewAdjustedImageArrayAnimation(this.ImageArray, this.Dx, this.Dy, this.AnimationBehaviorFactory.GetOrCreateInstance())
        } else {
            return NewImageArrayAnimation(this.ImageArray, AnimationBehavior_GetInstance())
        }
}


func (this *ImageArrayAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

