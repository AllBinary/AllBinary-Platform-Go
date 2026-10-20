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

type SimpleWeaponLayer struct {
    WeaponLayer
}

// The Java class SimpleWeaponLayer has subclasses so its variables are this interface
type ISimpleWeaponLayer interface {
    IWeaponLayer
    SimpleWeaponLayer_() *SimpleWeaponLayer
}

func (this *SimpleWeaponLayer) SimpleWeaponLayer_() *SimpleWeaponLayer {
    return this
}



// throws Exception
func SimpleWeaponLayer_CreateDestroyed() IAnimation {
        return FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().GetProcedural(ExplosionResources_GetInstance().THIRD_EXPLOSION_RESOURCE).GetInstanceAnimation(NullIndexedAnimationFactory_GetFactoryInstance().GetInstance(0))
}


func NewSimpleWeaponLayer(name string, remoteInfo *RemoteInfo, multiPlayerType int, movement IMovement, animationInterface IAnimation, destroyedAnimationInterface IAnimation, rectangle *Rectangle, viewPosition IViewPositionBase) *SimpleWeaponLayer {
    this := &SimpleWeaponLayer{}
        this.WeaponLayer = *NewWeaponLayer(name, remoteInfo, multiPlayerType, movement, animationInterface, destroyedAnimationInterface, rectangle, viewPosition)
        this.SetCollidableInferface(NewCollidableWeaponBehavior(true))
    return this
}

