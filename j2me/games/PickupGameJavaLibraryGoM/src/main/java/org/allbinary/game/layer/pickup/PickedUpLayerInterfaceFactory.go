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

type PickedUpLayerInterfaceFactory struct {
    PickedUpLayerInterfaceFactoryBase
    PickedUpLayerType *PickedUpLayerType
    IconLayer *IconLayer
    AnimationInterface IAnimation
}

// The Java class PickedUpLayerInterfaceFactory has subclasses so its variables are this interface
type IPickedUpLayerInterfaceFactory interface {
    IPickedUpLayerInterfaceFactoryBase
    PickedUpLayerInterfaceFactory_() *PickedUpLayerInterfaceFactory
    GetIconLayer() IAllBinaryLayer
    GetAnimationInterface() IAnimation
    GetPickedUpLayerType() *PickedUpLayerType
}

func (this *PickedUpLayerInterfaceFactory) PickedUpLayerInterfaceFactory_() *PickedUpLayerInterfaceFactory {
    return this
}


func (this *PickedUpLayerInterfaceFactory) initFields() {
    
    
    
}


func NewPickedUpLayerInterfaceFactory(pickeUpLayerType *PickedUpLayerType, iconLayer *IconLayer, animationInterface IAnimation) *PickedUpLayerInterfaceFactory {
    this := &PickedUpLayerInterfaceFactory{}
    this.PickedUpLayerInterfaceFactoryBase = *NewPickedUpLayerInterfaceFactoryBase()
    this.initFields()
        this.PickedUpLayerType = pickeUpLayerType
        this.AnimationInterface = animationInterface
        this.IconLayer = iconLayer
    return this
}


func (this *PickedUpLayerInterfaceFactory) GetIconLayer() IAllBinaryLayer {
        return this.IconLayer
}


func (this *PickedUpLayerInterfaceFactory) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *PickedUpLayerInterfaceFactory) GetPickedUpLayerType() *PickedUpLayerType {
        return this.PickedUpLayerType
}

