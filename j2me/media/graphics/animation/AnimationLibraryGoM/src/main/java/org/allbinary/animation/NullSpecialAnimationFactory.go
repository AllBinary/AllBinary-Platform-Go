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

type NullSpecialAnimationFactory struct {
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface

var NullSpecialAnimationFactory_NULL_SPECIAL_ANIMATION_FACTORY *NullSpecialAnimationFactory = NewNullSpecialAnimationFactory()


func NewNullSpecialAnimationFactory() *NullSpecialAnimationFactory {
    this := &NullSpecialAnimationFactory{}
    return this
}


func NullSpecialAnimationFactory_GetFactoryInstance() *NullSpecialAnimationFactory {
        return NullSpecialAnimationFactory_NULL_SPECIAL_ANIMATION_FACTORY
}


// throws Exception
func (this *NullSpecialAnimationFactory) GetInstance(instanceId int) IAnimation {
        return SpecialAnimation_GetInstance()
}


// throws Exception
func (this *NullSpecialAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return SpecialAnimation_GetInstance()
}


func (this *NullSpecialAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

