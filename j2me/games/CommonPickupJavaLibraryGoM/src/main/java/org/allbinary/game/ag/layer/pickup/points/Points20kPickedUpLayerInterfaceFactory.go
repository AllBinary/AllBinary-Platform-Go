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

type Points20kPickedUpLayerInterfaceFactory struct {
    PickedUpLayerInterfaceFactory
    Pool *PointsLayerCircularStaticPool
}

// Implements PickupProcessorInterface

var Points20kPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface = CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY

func (this *Points20kPickedUpLayerInterfaceFactory) initFields() {
    this.Pool = NewPointsLayerCircularStaticPool(NewPointsLayerFactory(GameConfigurationUtil_GetInstance().GetCompetitionValue() * 20000), 1)
}


// throws Exception
func Points20kPickedUpLayerInterfaceFactory_Init() {
        Points20kPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface = NewPoints20kPickedUpLayerInterfaceFactory()
}


func NewPoints20kPickedUpLayerInterfaceFactory() *Points20kPickedUpLayerInterfaceFactory {
    this := &Points20kPickedUpLayerInterfaceFactory{}
        this.PickedUpLayerInterfaceFactory = *NewPickedUpLayerInterfaceFactory(PickedUpLayerTypeFactory_GetInstance().POINTS, IconLayerFactory_GetInstance(FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(PointsResources_GetInstance().RESOURCE_2).GetInstance(0), 10, 10), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(PointsResources_GetInstance().RESOURCE_2).GetInstance(0))
        this.initFields()
    return this
}


func (this *Points20kPickedUpLayerInterfaceFactory) GetTotal() int {
        return 1
}


func Points20kPickedUpLayerInterfaceFactory_GetInstance() PickedUpLayerInterfaceFactoryInterface {
        return Points20kPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface
}


// throws Exception
func (this *Points20kPickedUpLayerInterfaceFactory) Process(sourceLayerInterface IAllBinaryLayer) {
        this.Pool.Visit(sourceLayerInterface)
}

