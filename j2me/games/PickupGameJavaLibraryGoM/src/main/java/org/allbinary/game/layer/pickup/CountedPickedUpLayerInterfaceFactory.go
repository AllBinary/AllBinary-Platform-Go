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

type CountedPickedUpLayerInterfaceFactory struct {
    PickedUpLayerInterfaceFactory
    Id int
}

// Implements CountedPickedUpLayerInterfaceFactoryInterface

// The Java class CountedPickedUpLayerInterfaceFactory has subclasses so its variables are this interface
type ICountedPickedUpLayerInterfaceFactory interface {
    IPickedUpLayerInterfaceFactory
    CountedPickedUpLayerInterfaceFactory_() *CountedPickedUpLayerInterfaceFactory
    GetNextInstance(hashtable *ABHashtable[any, any], x int, y int, z int) IAllBinaryLayer
    GetTotal() int
    GetId() int
}

func (this *CountedPickedUpLayerInterfaceFactory) CountedPickedUpLayerInterfaceFactory_() *CountedPickedUpLayerInterfaceFactory {
    return this
}

var CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY ICountedPickedUpLayerInterfaceFactory = NewCountedPickedUpLayerInterfaceFactory(PickedUpLayerTypeFactory_GetInstance().NONE, IconLayer_NULL_ICON_LAYER, NullAnimationFactory_GetFactoryInstance().GetInstance(0))

func (this *CountedPickedUpLayerInterfaceFactory) initFields() {
    
}


func NewCountedPickedUpLayerInterfaceFactory(pickeUpLayerType *PickedUpLayerType, iconLayer *IconLayer, animationInterface IAnimation) *CountedPickedUpLayerInterfaceFactory {
    this := &CountedPickedUpLayerInterfaceFactory{}
        this.PickedUpLayerInterfaceFactory = *NewPickedUpLayerInterfaceFactory(pickeUpLayerType, iconLayer, animationInterface)
        this.initFields()
        var list IBasicArrayList = CountedPickedUpLayerInterfaceFactoryPool_GetInstance().GetList(); _ = list
        this.Id = list.Size()
        list.Add(this)
    return this
}


// throws Exception
func (this *CountedPickedUpLayerInterfaceFactory) GetNextInstance(hashtable *ABHashtable[any, any], x int, y int, z int) IAllBinaryLayer {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}


func (this *CountedPickedUpLayerInterfaceFactory) GetTotal() int {
        return -1
}


func (this *CountedPickedUpLayerInterfaceFactory) GetId() int {
        return this.Id
}

