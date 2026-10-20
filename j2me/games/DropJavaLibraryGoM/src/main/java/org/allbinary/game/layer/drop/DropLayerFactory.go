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

type DropLayerFactory struct {
    List IBasicArrayList
    BasicArrayListUtil *BasicArrayListUtil
}

var DropLayerFactory_DROP_LAYER_FACTORY *DropLayerFactory = NewDropLayerFactory()

func (this *DropLayerFactory) initFields() {
    this.List = NewBasicArrayListD()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewDropLayerFactory() *DropLayerFactory {
    this := &DropLayerFactory{}
    this.initFields()
    return this
}


func (this *DropLayerFactory) Clear() {
        this.List.Clear()
}


func DropLayerFactory_GetInstance() *DropLayerFactory {
        return DropLayerFactory_DROP_LAYER_FACTORY
}


func (this *DropLayerFactory) GetSize() int {
        return this.List.Size()
}


func (this *DropLayerFactory) GetRandomInstance() PickedUpLayerInterfaceFactoryInterface {
        return any(this.BasicArrayListUtil.GetRandom(this.List)).(PickedUpLayerInterfaceFactoryInterface)
}


func (this *DropLayerFactory) Add(layerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface) {
        this.List.Add(layerInterfaceFactoryInterface)
}


func (this *DropLayerFactory) GetList() IBasicArrayList {
        return this.List
}

