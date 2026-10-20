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

type SimultaneousCompoundIndexAnimationInterfaceFactory struct {
    BasicAnimationInterfaceFactoryInterfaceArray []AnimationInterfaceFactoryInterface
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface


func (this *SimultaneousCompoundIndexAnimationInterfaceFactory) initFields() {
    
    
}


func NewSimultaneousCompoundIndexAnimationInterfaceFactory(basicAnimationInterfaceFactoryInterfaceArray []AnimationInterfaceFactoryInterface, animationBehaviorFactory IAnimationBehaviorFactory) *SimultaneousCompoundIndexAnimationInterfaceFactory {
    this := &SimultaneousCompoundIndexAnimationInterfaceFactory{}
    this.initFields()
        this.BasicAnimationInterfaceFactoryInterfaceArray = basicAnimationInterfaceFactoryInterfaceArray
        this.AnimationBehaviorFactory = animationBehaviorFactory
    return this
}


// throws Exception
func (this *SimultaneousCompoundIndexAnimationInterfaceFactory) GetInstance(instanceId int) IAnimation {
        var size int = len(this.BasicAnimationInterfaceFactoryInterfaceArray); _ = size
        var animationInterfaceArray []IIndexedAnimation = make([]IIndexedAnimation, size); _ = animationInterfaceArray
        for index := int(0); index < size; index++ {
            animationInterfaceArray[index] = any(this.BasicAnimationInterfaceFactoryInterfaceArray[index].GetInstance(instanceId)).(IIndexedAnimation)
        }
        return this.GetOrCreate(animationInterfaceArray)
}


func (this *SimultaneousCompoundIndexAnimationInterfaceFactory) GetOrCreate(animationInterfaceArray []IIndexedAnimation) IAnimation {
        return NewSimultaneousCompoundIndexAnimation(animationInterfaceArray, this.AnimationBehaviorFactory.GetOrCreateInstance())
}


func (this *SimultaneousCompoundIndexAnimationInterfaceFactory) GetBasicAnimationInterfaceFactoryInterfaceArray() []AnimationInterfaceFactoryInterface {
        return this.BasicAnimationInterfaceFactoryInterfaceArray
}


func (this *SimultaneousCompoundIndexAnimationInterfaceFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

