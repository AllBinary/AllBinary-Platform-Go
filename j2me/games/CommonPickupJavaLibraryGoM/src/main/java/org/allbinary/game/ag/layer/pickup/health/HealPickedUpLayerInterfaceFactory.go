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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type HealPickedUpLayerInterfaceFactory struct {
    PickedUpLayerInterfaceFactory
    Pool *HealLayerCircularStaticPool
}

// Implements PickupProcessorInterface

var HealPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface = CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY

func (this *HealPickedUpLayerInterfaceFactory) initFields() {
    this.Pool = NewHealLayerCircularStaticPool(NewHealLayerFactory(), 1)
}


// throws Exception
func HealPickedUpLayerInterfaceFactory_Init() {
        HealPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface = NewHealPickedUpLayerInterfaceFactory()
}


func NewHealPickedUpLayerInterfaceFactory() *HealPickedUpLayerInterfaceFactory {
    this := &HealPickedUpLayerInterfaceFactory{}
        this.PickedUpLayerInterfaceFactory = *NewPickedUpLayerInterfaceFactory(PickedUpLayerTypeFactory_GetInstance().HEAL, IconLayerFactory_GetInstance(FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(HealthResources_GetInstance().RESOURCE).GetInstance(0), 10, 10), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(HealthResources_GetInstance().RESOURCE).GetInstance(0))
        this.initFields()
    return this
}


func (this *HealPickedUpLayerInterfaceFactory) GetTotal() int {
        return 1
}


func HealPickedUpLayerInterfaceFactory_GetInstance() PickedUpLayerInterfaceFactoryInterface {
        return HealPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface
}


// throws Exception
func (this *HealPickedUpLayerInterfaceFactory) Process(sourceLayerInterface IAllBinaryLayer) {
        this.Pool.Visit(sourceLayerInterface)
}

