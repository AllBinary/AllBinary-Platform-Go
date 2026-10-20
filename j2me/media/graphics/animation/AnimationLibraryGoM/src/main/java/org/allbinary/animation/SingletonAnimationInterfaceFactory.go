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

type SingletonAnimationInterfaceFactory struct {
    AnimationInterface IAnimation
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface

// The Java class SingletonAnimationInterfaceFactory has subclasses so its variables are this interface
type ISingletonAnimationInterfaceFactory interface {
    SingletonAnimationInterfaceFactory_() *SingletonAnimationInterfaceFactory
    GetInstance(instanceId int) IAnimation
    GetInstanceAnimation(animationInterface IAnimation) IAnimation
    SetInitialScale(scaleProperties *ScaleProperties)
}

func (this *SingletonAnimationInterfaceFactory) SingletonAnimationInterfaceFactory_() *SingletonAnimationInterfaceFactory {
    return this
}


func (this *SingletonAnimationInterfaceFactory) initFields() {
    
}


func NewSingletonAnimationInterfaceFactory(animationInterface IAnimation) *SingletonAnimationInterfaceFactory {
    this := &SingletonAnimationInterfaceFactory{}
    this.initFields()
        this.AnimationInterface = animationInterface
    return this
}


// throws Exception
func (this *SingletonAnimationInterfaceFactory) GetInstance(instanceId int) IAnimation {
        return this.AnimationInterface
}


// throws Exception
func (this *SingletonAnimationInterfaceFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.AnimationInterface
}


func (this *SingletonAnimationInterfaceFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

