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

type NullRotationAnimationFactory struct {
    NULL_ROTATION_ANIMATION_ARRAY []IRotationAnimation
    NULL_ANIMATION IAnimation
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface

var NullRotationAnimationFactory_NULL_ROTATION_ANIMATION_FACTORY *NullRotationAnimationFactory = NewNullRotationAnimationFactory()

func (this *NullRotationAnimationFactory) initFields() {
    this.NULL_ROTATION_ANIMATION_ARRAY = make([]IRotationAnimation, 0)
    this.NULL_ANIMATION = NewNullRotationAnimation2()
}


func NewNullRotationAnimationFactory() *NullRotationAnimationFactory {
    this := &NullRotationAnimationFactory{}
    this.initFields()
    return this
}


func NullRotationAnimationFactory_GetFactoryInstance() *NullRotationAnimationFactory {
        return NullRotationAnimationFactory_NULL_ROTATION_ANIMATION_FACTORY
}


func (this *NullRotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        return this.NULL_ANIMATION
}


// throws Exception
func (this *NullRotationAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.NULL_ANIMATION
}


func (this *NullRotationAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

type NullRotationAnimation2 struct {
    NullRotationAnimation
}

    

    func NewNullRotationAnimationFactory() *NullRotationAnimationFactory {
        this := &NullRotationAnimationFactory{}
            
        return this
    }
    

    func (this *NullRotationAnimationFactory) PaintXY(graphics IGraphics, x int, y int) {
    }

