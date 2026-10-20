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

type PickupLayerCircularStaticPool struct {
    AllBinaryLayerCircularPool
}

var PickupLayerCircularStaticPool_SINGLETON *PickupLayerCircularStaticPool = NewPickupLayerCircularStaticPool()


func NewPickupLayerCircularStaticPool() *PickupLayerCircularStaticPool {
    this := &PickupLayerCircularStaticPool{}
    this.AllBinaryLayerCircularPool = *NewAllBinaryLayerCircularPool()
    return this
}


func PickupLayerCircularStaticPool_GetInstance() *PickupLayerCircularStaticPool {
        return PickupLayerCircularStaticPool_SINGLETON
}


// throws Exception
func (this *PickupLayerCircularStaticPool) GetInstanceXYZ(pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface, x int, y int, z int) PickedUpLayerInterface {
        var pickupLayer IPickupLayer = any(this.GetNextInstance()).(IPickupLayer); _ = pickupLayer
        pickupLayer.InitXYZ(x, y, z)
        pickupLayer.Init(pickedUpLayerInterfaceFactoryInterface, pickedUpLayerInterfaceFactoryInterface.GetAnimationInterface())
        return pickupLayer
}

