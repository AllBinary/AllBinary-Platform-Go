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

type FallingPickupLayer struct {
    PickupLayer
}

// Implements TickableInterface

var FallingPickupLayer_NAME string = "FallingPickupLayer"


func NewFallingPickupLayer(remoteInfo *RemoteInfo) *FallingPickupLayer {
    this := &FallingPickupLayer{}
        this.PickupLayer = *NewPickupLayer(FallingPickupLayer_NAME, remoteInfo, 0, CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY, NullAnimationFactory_GetFactoryInstance().GetInstance(0), NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 0, 0), ViewPosition_GetInstanceD())
    return this
}


func (this *FallingPickupLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        this.SetPosition(this.X, this.Y + 1, this.Z)
}


func (this *FallingPickupLayer) ImplmentsTickableInterface() bool {
        return true
}

