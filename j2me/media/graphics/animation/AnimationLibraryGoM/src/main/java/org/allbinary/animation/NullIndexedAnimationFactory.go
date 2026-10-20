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

type NullIndexedAnimationFactory struct {
    NULL_ANIMATION IAnimation
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface

var NullIndexedAnimationFactory_NULL_INDEXED_ANIMATION_ARRAY []IIndexedAnimation = make([]IIndexedAnimation, 0)
var NullIndexedAnimationFactory_NULL_INDEXED_ANIMATION_FACTORY *NullIndexedAnimationFactory = NewNullIndexedAnimationFactory()

func (this *NullIndexedAnimationFactory) initFields() {
    this.NULL_ANIMATION = NewNullIndexedAnimationInner()
}


func NewNullIndexedAnimationFactory() *NullIndexedAnimationFactory {
    this := &NullIndexedAnimationFactory{}
    this.initFields()
    return this
}


func NullIndexedAnimationFactory_GetFactoryInstance() *NullIndexedAnimationFactory {
        return NullIndexedAnimationFactory_NULL_INDEXED_ANIMATION_FACTORY
}


// throws Exception
func (this *NullIndexedAnimationFactory) GetInstance(instanceId int) IAnimation {
        return this.NULL_ANIMATION
}


// throws Exception
func (this *NullIndexedAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.NULL_ANIMATION
}


func (this *NullIndexedAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

type NullIndexedAnimationInner struct {
    NullIndexedAnimation
}

    

    func NewNullIndexedAnimationFactory() *NullIndexedAnimationFactory {
        this := &NullIndexedAnimationFactory{}
            
        return this
    }
    

    func (this *NullIndexedAnimationFactory) PaintXY(graphics IGraphics, x int, y int) {
    }

