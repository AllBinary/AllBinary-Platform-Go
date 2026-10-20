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

type NullAnimationFactory struct {
    NULL_ANIMATION IAnimation
    EMPTY_ARRAY []IAnimation
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface

var NullAnimationFactory_NULL_NOT_FOR_USE_ANIMATION_FACTORY *NullAnimationFactory = NewNullAnimationFactory()
var NullAnimationFactory_NULL_ANIMATION_FACTORY *NullAnimationFactory = NewNullAnimationFactory()

func (this *NullAnimationFactory) initFields() {
    this.NULL_ANIMATION = /* Java anonymous class body: implement it as a named type
        

    func (this *NullAnimationFactory) PaintXY(graphics IGraphics, x int, y int) {
    }
*/ NewAnimation()
    this.EMPTY_ARRAY = make([]IAnimation, 0)
}


func NewNullAnimationFactory() *NullAnimationFactory {
    this := &NullAnimationFactory{}
    this.initFields()
    return this
}


func NullAnimationFactory_GetFactoryInstance() *NullAnimationFactory {
        return NullAnimationFactory_NULL_ANIMATION_FACTORY
}


func (this *NullAnimationFactory) GetInstance(instanceId int) IAnimation {
        return this.NULL_ANIMATION
}


// throws Exception
func (this *NullAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return this.NULL_ANIMATION
}


func (this *NullAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

