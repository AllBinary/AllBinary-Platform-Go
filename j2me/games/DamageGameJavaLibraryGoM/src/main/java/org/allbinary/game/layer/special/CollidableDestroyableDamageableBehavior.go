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

type CollidableDestroyableDamageableBehavior struct {
    CollidableBaseBehavior
    DamageUtil *DamageUtil
}

// The Java class CollidableDestroyableDamageableBehavior has subclasses so its variables are this interface
type ICollidableDestroyableDamageableBehavior interface {
    ICollidableBaseBehavior
    CollidableDestroyableDamageableBehavior_() *CollidableDestroyableDamageableBehavior
    IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool
    Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer)
    IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool
    CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface)
}

func (this *CollidableDestroyableDamageableBehavior) CollidableDestroyableDamageableBehavior_() *CollidableDestroyableDamageableBehavior {
    return this
}


func (this *CollidableDestroyableDamageableBehavior) initFields() {
    this.DamageUtil = DamageUtil_GetInstance()
}


func NewCollidableDestroyableDamageableBehavior(collidable bool) *CollidableDestroyableDamageableBehavior {
    this := &CollidableDestroyableDamageableBehavior{}
        this.CollidableBaseBehavior = *NewCollidableBaseBehavior(collidable)
        this.initFields()
    return this
}


func (this *CollidableDestroyableDamageableBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool {
        if ownerLayer.GetGroupInterface()[0] != collisionLayer.GetGroupInterface()[0] {
            return this.CollidableBaseBehavior.IsCollision(ownerLayer, collisionLayer)
        }
        return false
}


// throws Exception
func (this *CollidableDestroyableDamageableBehavior) Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var ownerCollidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(ownerLayer).(ICollidableDestroyableDamageableLayer); _ = ownerCollidableDestroyableDamageableLayer
        var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(collidableInterfaceCompositeInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
        ownerCollidableDestroyableDamageableLayer.Damage(collidableDestroyableDamageableLayer.GetDamage(0), 0)
}


func (this *CollidableDestroyableDamageableBehavior) IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool {
        ForcedLogUtil_Log("No Longer Used", this)
        return false
}


// throws Exception
func (this *CollidableDestroyableDamageableBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        ForcedLogUtil_Log("No Longer Used", this)
}

