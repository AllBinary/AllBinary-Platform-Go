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

type MineLayerCircularStaticPool struct {
}

var MineLayerCircularStaticPool_WEAPON_LAYER_CIRCULAR_STATIC_POOL *WeaponLayerCircularPool = NewWeaponLayerCircularPool()

func NewMineLayerCircularStaticPool() *MineLayerCircularStaticPool {
    this := &MineLayerCircularStaticPool{}
    return this
}


// throws Exception
func MineLayerCircularStaticPool_Init(viewPositionFactoryInterface ViewPositionFactoryInterface) {
        MineCountedPickedUpLayerInterfaceFactory_Init()
        MineLayerCircularStaticPool_WEAPON_LAYER_CIRCULAR_STATIC_POOL.InitAllBinaryLayerCircularPool(NewMineLayerFactory(viewPositionFactoryInterface), 20)
}


func MineLayerCircularStaticPool_GetInstance() *WeaponLayerCircularPool {
        return MineLayerCircularStaticPool_WEAPON_LAYER_CIRCULAR_STATIC_POOL
}

