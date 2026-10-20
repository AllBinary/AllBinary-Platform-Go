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

type CollidableWeaponBehavior struct {
    CollidableDestroyableDamageableBehavior
    Collided bool
    CollisionHelper ICollisionHelper
}

var CollidableWeaponBehavior_NULL_COLLIDABLE_WEAPON_BEHAVIOR any = NullUtil_GetInstance().NULL_OBJECT

func (this *CollidableWeaponBehavior) initFields() {
    
    
}


func CollidableWeaponBehavior_GetNullInstance() *CollidableWeaponBehavior {
        if CollidableWeaponBehavior_NULL_COLLIDABLE_WEAPON_BEHAVIOR == NullUtil_GetInstance().NULL_OBJECT {
            CollidableWeaponBehavior_NULL_COLLIDABLE_WEAPON_BEHAVIOR = NewCollidableWeaponBehavior(false)
        }
        return any(CollidableWeaponBehavior_NULL_COLLIDABLE_WEAPON_BEHAVIOR).(*CollidableWeaponBehavior)
}


func NewCollidableWeaponBehavior(collidable bool) *CollidableWeaponBehavior {
    this := &CollidableWeaponBehavior{}
        this.CollidableDestroyableDamageableBehavior = *NewCollidableDestroyableDamageableBehavior(collidable)
        this.initFields()
        this.SetCollided(false)
        this.SetCollidable(true)
        this.CollisionHelper = CollidableHelperFactory_GetInstance()
    return this
}


func (this *CollidableWeaponBehavior) Init(sourceLayerInterface IAllBinaryLayer) {
        this.SetCollided(false)
        this.CollisionHelper.SetOwnerLayerInterface(sourceLayerInterface)
}


func (this *CollidableWeaponBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool {
        if this.CollisionHelper.IsCollidable(collisionLayer) {
            if ownerLayer.GetGroupInterface()[0] != collisionLayer.GetGroupInterface()[0] {
                return this.CollidableDestroyableDamageableBehavior.IsCollision(ownerLayer, collisionLayer)
            }
        }
        return false
}


// throws Exception
func (this *CollidableWeaponBehavior) Collide(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) {
        this.CollidableDestroyableDamageableBehavior.Collide(ownerLayer, collisionLayer)
        this.Collided = true
}


func (this *CollidableWeaponBehavior) IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool {
        if this.CollisionHelper.IsCollidable(any(collidableInterfaceCompositeInterface).(ICollidableCompositeLayer)) {
            var layerInterface IAllBinaryLayer = any(collidableInterfaceCompositeInterface).(IAllBinaryLayer); _ = layerInterface
            if ownerLayer.GetGroupInterface()[0] != layerInterface.GetGroupInterface()[0] {
                if this.LayerCollisionUtil.IsCollision(ownerLayer, layerInterface) {
                    return true
                }
            }
        }
        return false
}


// throws Exception
func (this *CollidableWeaponBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        var ownerDamageableInterface DamageableInterface = any(ownerLayer).(DamageableInterface); _ = ownerDamageableInterface
        var damageableInterface DamageableInterface = any(collidableInterfaceCompositeInterface).(DamageableInterface); _ = damageableInterface
        this.DamageUtil.Process(ownerDamageableInterface, damageableInterface)
        this.Collided = true
}


func (this *CollidableWeaponBehavior) GetCollisionTypeWith(layerInterface IAllBinaryLayer) *CollisionType {
        return CollisionTypeFactory_GetInstance().COLLISION
}


func (this *CollidableWeaponBehavior) IsCollided() bool {
        return this.Collided
}


func (this *CollidableWeaponBehavior) SetCollided(collided bool) {
        this.Collided = collided
}

