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

type BaseAnimationInterfaceFactoryInterfaceComposite struct {
    BasicAnimationInterfaceFactoryInterfaceArray []BasicAnimationInterfaceFactoryInterface
}

// Implements AnimationInterfaceFactoryInterface


func (this *BaseAnimationInterfaceFactoryInterfaceComposite) initFields() {
    
}


func NewBaseAnimationInterfaceFactoryInterfaceComposite(basicAnimationInterfaceFactoryInterfaceArray []BasicAnimationInterfaceFactoryInterface) *BaseAnimationInterfaceFactoryInterfaceComposite {
    this := &BaseAnimationInterfaceFactoryInterfaceComposite{}
    this.initFields()
        this.BasicAnimationInterfaceFactoryInterfaceArray = basicAnimationInterfaceFactoryInterfaceArray
    return this
}


// throws Exception
func (this *BaseAnimationInterfaceFactoryInterfaceComposite) GetInstance(instanceId int) IAnimation {
        return NullAnimationFactory_GetFactoryInstance().GetInstance(0)
}


func (this *BaseAnimationInterfaceFactoryInterfaceComposite) GetBasicAnimationInterfaceFactoryInterfaceArray() []BasicAnimationInterfaceFactoryInterface {
        return this.BasicAnimationInterfaceFactoryInterfaceArray
}


func (this *BaseAnimationInterfaceFactoryInterfaceComposite) SetBasicAnimationInterfaceFactoryInterfaceArray(basicAnimationInterfaceFactoryInterfaceArray []BasicAnimationInterfaceFactoryInterface) {
        this.BasicAnimationInterfaceFactoryInterfaceArray = basicAnimationInterfaceFactoryInterfaceArray
}


func (this *BaseAnimationInterfaceFactoryInterfaceComposite) SetInitialScale(scaleProperties *ScaleProperties) {
}

