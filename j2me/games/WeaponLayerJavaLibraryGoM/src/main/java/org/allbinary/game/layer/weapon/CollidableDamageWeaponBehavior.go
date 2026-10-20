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

type CollidableDamageWeaponBehavior struct {
    CollidableDestroyableDamageableBehavior
    Collided bool
    CollisionHelper ICollisionHelper
}


func (this *CollidableDamageWeaponBehavior) initFields() {
    
    
}


func NewCollidableDamageWeaponBehavior(collidable bool) *CollidableDamageWeaponBehavior {
    this := &CollidableDamageWeaponBehavior{}
        this.CollidableDestroyableDamageableBehavior = *NewCollidableDestroyableDamageableBehavior(collidable)
        this.initFields()
        this.SetCollided(false)
        this.SetCollidable(true)
        this.CollisionHelper = CollidableHelperFactory_GetInstance()
    return this
}


func (this *CollidableDamageWeaponBehavior) Init(sourceLayerInterface IAllBinaryLayer) {
        this.SetCollided(false)
        this.CollisionHelper.SetOwnerLayerInterface(sourceLayerInterface)
}


func (this *CollidableDamageWeaponBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool {
        if this.CollisionHelper.IsCollidable(collisionLayer) {
            if ownerLayer.GetGroupInterface()[0] != collisionLayer.GetGroupInterface()[0] {
                return this.CollidableDestroyableDamageableBehavior.IsCollision(ownerLayer, collisionLayer)
            }
        }
        return false
}


// throws Exception
func (this *CollidableDamageWeaponBehavior) Collide(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) {
        this.CollidableDestroyableDamageableBehavior.Collide(ownerLayer, collisionLayer)
        this.Collided = true
}


func (this *CollidableDamageWeaponBehavior) IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool {
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
func (this *CollidableDamageWeaponBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        this.DamageUtil.Process(any(ownerLayer).(DamageableInterface), any(collidableInterfaceCompositeInterface).(DamageableInterface))
        this.Collided = true
}


func (this *CollidableDamageWeaponBehavior) GetCollisionTypeWith(layerInterface IAllBinaryLayer) *CollisionType {
        return CollisionTypeFactory_GetInstance().COLLISION
}


func (this *CollidableDamageWeaponBehavior) IsCollided() bool {
        return this.Collided
}


func (this *CollidableDamageWeaponBehavior) SetCollided(collided bool) {
        this.Collided = collided
}

