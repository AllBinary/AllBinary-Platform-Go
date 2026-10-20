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

type PickupBehavior struct {
    LogUtil *LogUtil
    CountedIndex int
    OwnerLayerInterface ICollidableCompositeLayer
}

// Implements PickupBehaviorInterface

var PickupBehavior_NULL_PICKUP_BEHAVIOR any = NullUtil_GetInstance().NULL_OBJECT

func (this *PickupBehavior) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func PickupBehavior_GetNullInstance() *PickupBehavior {
        if PickupBehavior_NULL_PICKUP_BEHAVIOR == NullUtil_GetInstance().NULL_OBJECT {
            PickupBehavior_NULL_PICKUP_BEHAVIOR = NewPickupBehavior(CollidableCompositeLayer_GetNullInstance(), 0)
        }
        return any(PickupBehavior_NULL_PICKUP_BEHAVIOR).(*PickupBehavior)
}


func NewPickupBehavior(ownerLayerInterface ICollidableCompositeLayer, countedIndex int) *PickupBehavior {
    this := &PickupBehavior{}
    this.initFields()
        this.OwnerLayerInterface = ownerLayerInterface
        this.CountedIndex = countedIndex
    return this
}


func (this *PickupBehavior) DoPickup(pickupProcessorInterface PickupProcessorInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "doPickup", e)
                }
            }()
            var healthInterfaceCompositeInterface HealthInterfaceCompositeInterface = any(this.OwnerLayerInterface).(HealthInterfaceCompositeInterface); _ = healthInterfaceCompositeInterface
            if healthInterfaceCompositeInterface.GetHealthInterface().IsAlive() {
                pickupProcessorInterface.Process(this.OwnerLayerInterface)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PickupBehavior) DoPickupLayer(pickupLayerInterface PickedUpLayerInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_PutOE(commonStrings.EXCEPTION, this, "doPickup", e)
                }
            }()
            var pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface = pickupLayerInterface.GetPickedUpLayerInterfaceFactoryInterface(); _ = pickedUpLayerInterfaceFactoryInterface
            var pickedUpLayerType *PickedUpLayerType = pickedUpLayerInterfaceFactoryInterface.GetPickedUpLayerType(); _ = pickedUpLayerType
            var pickedUpLayerTypeFactory *PickedUpLayerTypeFactory = PickedUpLayerTypeFactory_GetInstance(); _ = pickedUpLayerTypeFactory
            if pickedUpLayerType == pickedUpLayerTypeFactory.BEAM || pickedUpLayerType == pickedUpLayerTypeFactory.MINE || pickedUpLayerType == pickedUpLayerTypeFactory.PROJECTILE || pickedUpLayerType == pickedUpLayerTypeFactory.ENHANCEMENT {
                this.Add(any(pickedUpLayerInterfaceFactoryInterface).(CountedPickedUpLayerInterfaceFactoryInterface))
            } else if pickedUpLayerInterfaceFactoryInterface.GetPickedUpLayerType() == pickedUpLayerTypeFactory.PART {
                var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(this.OwnerLayerInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
                collidableDestroyableDamageableLayer.AddPart(pickedUpLayerInterfaceFactoryInterface)
            } else if pickedUpLayerType == pickedUpLayerTypeFactory.NONE {
            } else {
                var pickupProcessorInterface PickupProcessorInterface = any(pickedUpLayerInterfaceFactoryInterface).(PickupProcessorInterface); _ = pickupProcessorInterface
                this.DoPickup(pickupProcessorInterface)
            }
            pickupLayerInterface.SetPickedUp()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PickupBehavior) Add(countedPickedUpLayerInterfaceFactoryInterface CountedPickedUpLayerInterfaceFactoryInterface) {
        var partIndex int = countedPickedUpLayerInterfaceFactoryInterface.GetId() + this.CountedIndex; _ = partIndex
        this.AddAt(countedPickedUpLayerInterfaceFactoryInterface, partIndex)
}


func (this *PickupBehavior) AddAt(countedPickedUpLayerInterfaceFactoryInterface CountedPickedUpLayerInterfaceFactoryInterface, partIndex int) {
        var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(this.OwnerLayerInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
        var countedLayerInterfaceFactory *CountedLayerInterfaceFactoryPart = any(collidableDestroyableDamageableLayer.GetPartInterfaceArray()[partIndex]).(*CountedLayerInterfaceFactoryPart); _ = countedLayerInterfaceFactory
        countedLayerInterfaceFactory.SetTotal(countedLayerInterfaceFactory.GetTotal() + countedPickedUpLayerInterfaceFactoryInterface.GetTotal())
}


func (this *PickupBehavior) GetFirstAvailableCountedLayerInterfaceFactory(pickedUpLayerTypeArray []*PickedUpLayerType) *CountedLayerInterfaceFactoryPart {
        var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(this.OwnerLayerInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
        var partInterfaceArray []PartInterface = collidableDestroyableDamageableLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var nextCountedLayerInterfaceFactory *CountedLayerInterfaceFactoryPart; _ = nextCountedLayerInterfaceFactory
        var countedPickedUpLayerInterfaceFactory ICountedPickedUpLayerInterfaceFactory; _ = countedPickedUpLayerInterfaceFactory
        for index := int(this.CountedIndex); index < size; index++ {
            nextCountedLayerInterfaceFactory = any(collidableDestroyableDamageableLayer.GetPartInterfaceArray()[index]).(*CountedLayerInterfaceFactoryPart)
            countedPickedUpLayerInterfaceFactory = nextCountedLayerInterfaceFactory.GetCountedPickedUpLayerInterfaceFactory()
            for index2 := int(len(pickedUpLayerTypeArray)); Java_PreDecrement(&index2) >= 0;  {
                if countedPickedUpLayerInterfaceFactory.GetPickedUpLayerType() == pickedUpLayerTypeArray[index2] {
                    if nextCountedLayerInterfaceFactory.GetTotal() > 0 {
                        return nextCountedLayerInterfaceFactory
                    }
                }
            }
        }
        return CountedLayerInterfaceFactoryPart_NULL_COUNTED_LAYER_INTERFACE_FACTORY
}


// throws Exception
func (this *PickupBehavior) GetCountedLayerInterfaceFactoryForSlot(slotIndex int) *CountedLayerInterfaceFactoryPart {
        var currentSlot int = 0; _ = currentSlot
        var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(this.OwnerLayerInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
        var partInterfaceArray []PartInterface = collidableDestroyableDamageableLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var nextCountedLayerInterfaceFactory *CountedLayerInterfaceFactoryPart; _ = nextCountedLayerInterfaceFactory
        for index := int(this.CountedIndex); index < size; index++ {
            nextCountedLayerInterfaceFactory = any(collidableDestroyableDamageableLayer.GetPartInterfaceArray()[index]).(*CountedLayerInterfaceFactoryPart)
            if nextCountedLayerInterfaceFactory.GetTotal() > 0 {
                if currentSlot == slotIndex {
                    return nextCountedLayerInterfaceFactory
                } else {
                    currentSlot++
                }
            }
        }
        return CountedLayerInterfaceFactoryPart_NULL_COUNTED_LAYER_INTERFACE_FACTORY
}


func (this *PickupBehavior) GetCountedIndex() int {
        return this.CountedIndex
}

