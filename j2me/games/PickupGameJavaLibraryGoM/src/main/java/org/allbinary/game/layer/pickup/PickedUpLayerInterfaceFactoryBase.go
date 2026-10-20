/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type PickedUpLayerInterfaceFactoryBase struct {
}

// Implements PickedUpLayerInterfaceFactoryInterface

// The Java class PickedUpLayerInterfaceFactoryBase has subclasses so its variables are this interface
type IPickedUpLayerInterfaceFactoryBase interface {
    PickedUpLayerInterfaceFactoryBase_() *PickedUpLayerInterfaceFactoryBase
    GetIconLayer() IAllBinaryLayer
    GetAnimationInterface() IAnimation
    GetPickedUpLayerType() *PickedUpLayerType
}

func (this *PickedUpLayerInterfaceFactoryBase) PickedUpLayerInterfaceFactoryBase_() *PickedUpLayerInterfaceFactoryBase {
    return this
}

var PickedUpLayerInterfaceFactoryBase_NULL_PICKED_UP_LAYER_FACTORY_BASE IPickedUpLayerInterfaceFactoryBase = NewPickedUpLayerInterfaceFactoryBase()

func NewPickedUpLayerInterfaceFactoryBase() *PickedUpLayerInterfaceFactoryBase {
    this := &PickedUpLayerInterfaceFactoryBase{}
    return this
}


func (this *PickedUpLayerInterfaceFactoryBase) GetIconLayer() IAllBinaryLayer {
        panic(errors.New(fmt.Sprint()))
}


func (this *PickedUpLayerInterfaceFactoryBase) GetAnimationInterface() IAnimation {
        panic(errors.New(fmt.Sprint()))
}


func (this *PickedUpLayerInterfaceFactoryBase) GetPickedUpLayerType() *PickedUpLayerType {
        panic(errors.New(fmt.Sprint()))
}

