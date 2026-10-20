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

type CollidableBaseBehavior struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    LayerCollisionUtil *LayerCollisionUtil
    Collidable bool
}

// Implements CollidableInterface

// The Java class CollidableBaseBehavior has subclasses so its variables are this interface
type ICollidableBaseBehavior interface {
    CollidableBaseBehavior_() *CollidableBaseBehavior
    Update()
    GetName() string
    SetCollidable(collidable bool)
    IsCollidable(ownerLayer ICollidableCompositeLayer) bool
    Collide(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer)
    IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool
    IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool
    CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface)
    Paint(ownerLayer ICollidableCompositeLayer, graphics IGraphics)
    GetCollisionTypeWith(layerInterface IAllBinaryLayer) *CollisionType
}

func (this *CollidableBaseBehavior) CollidableBaseBehavior_() *CollidableBaseBehavior {
    return this
}


func (this *CollidableBaseBehavior) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.LayerCollisionUtil = LayerCollisionUtil_GetInstance()
    this.Collidable = true
}


func NewCollidableBaseBehavior(collidable bool) *CollidableBaseBehavior {
    this := &CollidableBaseBehavior{}
    this.initFields()
        this.Collidable = collidable
    return this
}


func (this *CollidableBaseBehavior) Update() {
}


func (this *CollidableBaseBehavior) GetName() string {
        return reflect.TypeOf(this).String()
}


func (this *CollidableBaseBehavior) SetCollidable(collidable bool) {
        this.Collidable = collidable
}


func (this *CollidableBaseBehavior) IsCollidable(ownerLayer ICollidableCompositeLayer) bool {
        return this.Collidable
}


// throws Exception
func (this *CollidableBaseBehavior) Collide(ownerLayer ICollidableCompositeLayer, allbinaryCollidableLayer ICollidableCompositeLayer) {
        panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.NOT_IMPLEMENTED).Append(CommonLabels_GetInstance().COLON_SEP).Append(reflect.TypeOf(this).String())))))
}


func (this *CollidableBaseBehavior) IsCollision(ownerLayer ICollidableCompositeLayer, collisionLayer ICollidableCompositeLayer) bool {
        return this.LayerCollisionUtil.IsCollision(ownerLayer, collisionLayer)
}


func (this *CollidableBaseBehavior) IsCollisionInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) bool {
        return this.LayerCollisionUtil.IsCollision(ownerLayer, any(collidableInterfaceCompositeInterface).(IAllBinaryLayer))
}


// throws Exception
func (this *CollidableBaseBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *CollidableBaseBehavior) Paint(ownerLayer ICollidableCompositeLayer, graphics IGraphics) {
}


func (this *CollidableBaseBehavior) GetCollisionTypeWith(layerInterface IAllBinaryLayer) *CollisionType {
        return CollisionTypeFactory_GetInstance().NONE
}

