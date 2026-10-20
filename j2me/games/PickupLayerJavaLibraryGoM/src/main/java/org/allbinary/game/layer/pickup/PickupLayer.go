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

type PickupLayer struct {
    MultiPlayerGameLayer
    PickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface
    Destroyed bool
    AnimationInterface IAnimation
}

// Implements PickedUpLayerInterface, PickupableInterface

// The Java class PickupLayer has subclasses so its variables are this interface
type IPickupLayer interface {
    IMultiPlayerGameLayer
    PickupLayer_() *PickupLayer
    Init(pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface, animationInterface IAnimation)
    InitXYZ(x int, y int, z int)
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    GetPickedUpLayerInterfaceFactoryInterface() PickedUpLayerInterfaceFactoryInterface
    SetPickedUp()
    IsDestroyed() bool
    SetDestroyed(destroyed bool)
    Damage(damage int, damageType int)
    GetDamage(damageType int) int
    Set(gl GL)
}

func (this *PickupLayer) PickupLayer_() *PickupLayer {
    return this
}


func (this *PickupLayer) initFields() {
    this.PickedUpLayerInterfaceFactoryInterface = CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY
    
    this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
}


func NewPickupLayer(name string, remoteInfo *RemoteInfo, total int, pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface, animationInterface IAnimation, rectangle *Rectangle, viewPosition IViewPositionBase) *PickupLayer {
    this := &PickupLayer{}
        this.MultiPlayerGameLayer = *NewMultiPlayerGameLayer(remoteInfo, BasicGroupFactory_GetInstance().NONE_ARRAY, name, rectangle, viewPosition)
        this.initFields()
        this.SetCollidableInferface(CollidableAlwaysPickupNeverCollideBehaviorFactory_GetInstance().CreateBehavior())
        this.SetLayerWidth(10)
        this.SetLayerHeight(10)
        this.Init(pickedUpLayerInterfaceFactoryInterface, animationInterface)
    return this
}


func (this *PickupLayer) Init(pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface, animationInterface IAnimation) {
        this.PickedUpLayerInterfaceFactoryInterface = pickedUpLayerInterfaceFactoryInterface
        this.AnimationInterface = animationInterface
        this.SetDestroyed(false)
}


func (this *PickupLayer) InitXYZ(x int, y int, z int) {
        this.SetPosition(x, y, z)
}


func (this *PickupLayer) Paint(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        var viewX int = viewPosition.GetX(); _ = viewX
        var viewY int = viewPosition.GetY(); _ = viewY
        this.AnimationInterface.PaintXY(graphics, viewX, viewY)
}


func (this *PickupLayer) PaintThreed(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        var viewX int = viewPosition.GetX(); _ = viewX
        var viewY int = viewPosition.GetY(); _ = viewY
        this.AnimationInterface.PaintThreedXYZ(graphics, viewX, viewY, 3)
}


func (this *PickupLayer) GetPickedUpLayerInterfaceFactoryInterface() PickedUpLayerInterfaceFactoryInterface {
        return this.PickedUpLayerInterfaceFactoryInterface
}


func (this *PickupLayer) SetPickedUp() {
        this.SetDestroyed(true)
}


func (this *PickupLayer) IsDestroyed() bool {
        return this.Destroyed
}


func (this *PickupLayer) SetDestroyed(destroyed bool) {
        this.Destroyed = destroyed
        if this.IsDestroyed() {
            DestroyedLayerProcessor_GetInstance().Add(this)
        }
}


func (this *PickupLayer) Damage(damage int, damageType int) {
}


func (this *PickupLayer) GetDamage(damageType int) int {
        return 0
}


// throws Exception
func (this *PickupLayer) Set(gl GL) {
        var openGLSurfaceChangedInterface OpenGLSurfaceChangedInterface = any(this.AnimationInterface).(OpenGLSurfaceChangedInterface); _ = openGLSurfaceChangedInterface
        openGLSurfaceChangedInterface.Set(gl)
}

