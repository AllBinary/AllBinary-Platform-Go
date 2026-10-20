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

type CompoundAnimationInterfaceFactory struct {
    BasicAnimationInterfaceFactoryInterfaceArrayP []AnimationInterfaceFactoryInterface
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface

// The Java class CompoundAnimationInterfaceFactory has subclasses so its variables are this interface
type ICompoundAnimationInterfaceFactory interface {
    CompoundAnimationInterfaceFactory_() *CompoundAnimationInterfaceFactory
    GetInstance(instanceId int) IAnimation
    CreateArray(size int) []IAnimation
    CreateAnimation(animationInterfaceArray []IAnimation) IAnimation
    GetBasicAnimationInterfaceFactoryInterfaceArray() []AnimationInterfaceFactoryInterface
    SetInitialScale(scaleProperties *ScaleProperties)
}

func (this *CompoundAnimationInterfaceFactory) CompoundAnimationInterfaceFactory_() *CompoundAnimationInterfaceFactory {
    return this
}


func (this *CompoundAnimationInterfaceFactory) initFields() {
    
    
}


func NewCompoundAnimationInterfaceFactory(basicAnimationInterfaceFactoryInterfaceArray []AnimationInterfaceFactoryInterface, animationBehaviorFactory IAnimationBehaviorFactory) *CompoundAnimationInterfaceFactory {
    this := &CompoundAnimationInterfaceFactory{}
    this.initFields()
        this.BasicAnimationInterfaceFactoryInterfaceArrayP = basicAnimationInterfaceFactoryInterfaceArray
        this.AnimationBehaviorFactory = animationBehaviorFactory
    return this
}


// throws Exception
func (this *CompoundAnimationInterfaceFactory) GetInstance(instanceId int) IAnimation {
        var size int = len(this.BasicAnimationInterfaceFactoryInterfaceArrayP); _ = size
        var animationInterfaceArray []IAnimation = this.CreateArray(size); _ = animationInterfaceArray
        for index := int(0); index < size; index++ {
            animationInterfaceArray[index] = this.BasicAnimationInterfaceFactoryInterfaceArrayP[index].GetInstance(instanceId)
        }
        return this.CreateAnimation(animationInterfaceArray)
}


func (this *CompoundAnimationInterfaceFactory) CreateArray(size int) []IAnimation {
        return NullAnimationFactory_GetFactoryInstance().EMPTY_ARRAY
}


func (this *CompoundAnimationInterfaceFactory) CreateAnimation(animationInterfaceArray []IAnimation) IAnimation {
        return NullAnimationFactory_GetFactoryInstance().GetInstance(0)
}


func (this *CompoundAnimationInterfaceFactory) GetBasicAnimationInterfaceFactoryInterfaceArray() []AnimationInterfaceFactoryInterface {
        return this.BasicAnimationInterfaceFactoryInterfaceArrayP
}


func (this *CompoundAnimationInterfaceFactory) SetInitialScale(scaleProperties *ScaleProperties) {
        var size int = len(this.BasicAnimationInterfaceFactoryInterfaceArrayP); _ = size
        for index := int(0); index < size; index++ {
            this.BasicAnimationInterfaceFactoryInterfaceArrayP[index].SetInitialScale(scaleProperties)
        }
}

