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

type MineCountedPickedUpLayerInterfaceFactory struct {
    CountedPickedUpLayerInterfaceFactory
    WeaponLayerCircularStaticPool *WeaponLayerCircularPool
    WeaponProperties *WeaponProperties
    SmallIntegerSingletonFactory *SmallIntegerSingletonFactory
}

var MineCountedPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface = CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY

func (this *MineCountedPickedUpLayerInterfaceFactory) initFields() {
    this.WeaponLayerCircularStaticPool = MineLayerCircularStaticPool_GetInstance()
    this.WeaponProperties = NewWeaponProperties(-1, -1, 0, 8000, int16(10))
    this.SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance()
}


// throws Exception
func MineCountedPickedUpLayerInterfaceFactory_Init() {
        MineCountedPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface = NewMineCountedPickedUpLayerInterfaceFactory()
}


func NewMineCountedPickedUpLayerInterfaceFactory() *MineCountedPickedUpLayerInterfaceFactory {
    this := &MineCountedPickedUpLayerInterfaceFactory{}
        this.CountedPickedUpLayerInterfaceFactory = *NewCountedPickedUpLayerInterfaceFactory(PickedUpLayerTypeFactory_GetInstance().MINE, IconLayerFactory_GetInstance(FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(MineWeaponResources_GetInstance().DROP_ICON_RESOURCE).GetInstance(0), 10, 10), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(MineWeaponResources_GetInstance().DROP_RESOURCE).GetInstance(0))
        this.initFields()
    return this
}


func MineCountedPickedUpLayerInterfaceFactory_GetInstance() PickedUpLayerInterfaceFactoryInterface {
        return MineCountedPickedUpLayerInterfaceFactory_PickedUpLayerInterfaceFactoryInterface
}


func (this *MineCountedPickedUpLayerInterfaceFactory) GetTotal() int {
        return 3
}


// throws Exception
func (this *MineCountedPickedUpLayerInterfaceFactory) GetNextInstance(hashtable *ABHashtable[any, any], x int, y int, z int) IAllBinaryLayer {
        var sourceLayerInterface IAllBinaryLayer = any(hashtable.Get(any(this.SmallIntegerSingletonFactory.GetAt(0)).(any))).(IAllBinaryLayer); _ = sourceLayerInterface
        var angle int = any(hashtable.Get(any(this.SmallIntegerSingletonFactory.GetAt(1)).(any))).(int); _ = angle
        var scoreableInterface ScoreableInterface = any(sourceLayerInterface).(ScoreableInterface); _ = scoreableInterface
        return this.WeaponLayerCircularStaticPool.GetInstanceWeaponLayer(sourceLayerInterface, x, y, z, int(angle), 0, this.WeaponProperties, scoreableInterface)
}

