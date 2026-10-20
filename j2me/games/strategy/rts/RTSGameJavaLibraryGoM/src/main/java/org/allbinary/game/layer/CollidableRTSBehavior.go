/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type CollidableRTSBehavior struct {
    CollidableDestroyableDamageableBehavior
}

// The Java class CollidableRTSBehavior has subclasses so its variables are this interface
type ICollidableRTSBehavior interface {
    ICollidableDestroyableDamageableBehavior
    CollidableRTSBehavior_() *CollidableRTSBehavior
    IsCollision(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) bool
    Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer)
    CollideNone(ownerLayer ICollidableCompositeLayer, collidableInterface ICollidableCompositeLayer)
}

func (this *CollidableRTSBehavior) CollidableRTSBehavior_() *CollidableRTSBehavior {
    return this
}



func NewCollidableRTSBehavior(collidable bool) *CollidableRTSBehavior {
    this := &CollidableRTSBehavior{}
        this.CollidableDestroyableDamageableBehavior = *NewCollidableDestroyableDamageableBehavior(collidable)
    return this
}


func (this *CollidableRTSBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) bool {
        var collisionTypeFactory *CollisionTypeFactory = CollisionTypeFactory_GetInstance(); _ = collisionTypeFactory
        var collisionType *CollisionType = collidableInterfaceCompositeInterface.GetCollidableInferface().GetCollisionTypeWith(ownerLayer); _ = collisionType
        if collisionType == collisionTypeFactory.PICKUP {
        } else if collisionType == collisionTypeFactory.COLLISION {
            return this.CollidableDestroyableDamageableBehavior.IsCollision(ownerLayer, collidableInterfaceCompositeInterface)
        } else {
            if collidableInterfaceCompositeInterface.GetX2() <= ownerLayer.GetXP() || collidableInterfaceCompositeInterface.GetY2() <= ownerLayer.GetYP() || collidableInterfaceCompositeInterface.GetYP() >= ownerLayer.GetY2() || collidableInterfaceCompositeInterface.GetXP() >= ownerLayer.GetX2() {
                return false
            } else {
                return true
            }
        }
        return false
}


// throws Exception
func (this *CollidableRTSBehavior) Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var collisionTypeFactory *CollisionTypeFactory = CollisionTypeFactory_GetInstance(); _ = collisionTypeFactory
        var collisionType *CollisionType = collidableInterfaceCompositeInterface.GetCollidableInferface().GetCollisionTypeWith(ownerLayer); _ = collisionType
        if collisionType == collisionTypeFactory.PICKUP {
        } else if collisionType == collisionTypeFactory.COLLISION {
            this.CollidableDestroyableDamageableBehavior.Collide(ownerLayer, collidableInterfaceCompositeInterface)
        } else {
            this.CollideNone(ownerLayer, collidableInterfaceCompositeInterface)
        }
}


// throws Exception
func (this *CollidableRTSBehavior) CollideNone(ownerLayer ICollidableCompositeLayer, collidableInterface ICollidableCompositeLayer) {
}

