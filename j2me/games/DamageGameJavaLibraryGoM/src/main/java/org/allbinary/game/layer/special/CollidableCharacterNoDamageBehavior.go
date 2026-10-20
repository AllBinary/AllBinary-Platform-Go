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

type CollidableCharacterNoDamageBehavior struct {
    CollidableDestroyableDamageableBehavior
}



func NewCollidableCharacterNoDamageBehavior(collidable bool) *CollidableCharacterNoDamageBehavior {
    this := &CollidableCharacterNoDamageBehavior{}
        this.CollidableDestroyableDamageableBehavior = *NewCollidableDestroyableDamageableBehavior(collidable)
    return this
}


// throws Exception
func (this *CollidableCharacterNoDamageBehavior) Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var collisionTypeFactory *CollisionTypeFactory = CollisionTypeFactory_GetInstance(); _ = collisionTypeFactory
        var collisionType *CollisionType = collidableInterfaceCompositeInterface.GetCollidableInferface().GetCollisionTypeWith(ownerLayer); _ = collisionType
        if collisionType == collisionTypeFactory.PICKUP {
            var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(ownerLayer).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
            collidableDestroyableDamageableLayer.GetPickupBehavior().DoPickupLayer(any(collidableInterfaceCompositeInterface).(PickedUpLayerInterface))
        } else {
        }
}


// throws Exception
func (this *CollidableCharacterNoDamageBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}

