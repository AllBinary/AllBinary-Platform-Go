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

type CollidableNeverCollideBehavior struct {
    CollidableBaseBehavior
}

// The Java class CollidableNeverCollideBehavior has subclasses so its variables are this interface
type ICollidableNeverCollideBehavior interface {
    ICollidableBaseBehavior
    CollidableNeverCollideBehavior_() *CollidableNeverCollideBehavior
    IsCollision(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer) bool
    Collide(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer)
    CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface)
    IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool
}

func (this *CollidableNeverCollideBehavior) CollidableNeverCollideBehavior_() *CollidableNeverCollideBehavior {
    return this
}



func NewCollidableNeverCollideBehavior(isCollidable bool) *CollidableNeverCollideBehavior {
    this := &CollidableNeverCollideBehavior{}
        this.CollidableBaseBehavior = *NewCollidableBaseBehavior(isCollidable)
    return this
}


func (this *CollidableNeverCollideBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer) bool {
        return false
}


// throws Exception
func (this *CollidableNeverCollideBehavior) Collide(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer) {
}


func (this *CollidableNeverCollideBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
}


func (this *CollidableNeverCollideBehavior) IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool {
        return false
}

