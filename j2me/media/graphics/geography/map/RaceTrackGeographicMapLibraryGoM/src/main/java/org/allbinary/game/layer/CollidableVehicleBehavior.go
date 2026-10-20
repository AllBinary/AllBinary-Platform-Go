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

type CollidableVehicleBehavior struct {
    CollidableDestroyableDamageableBehavior
    TotalImpactVelocity int64
    HalfImpactVelocityX int64
    HalfImpactVelocityY int64
}


func (this *CollidableVehicleBehavior) initFields() {
    this.TotalImpactVelocity = 0
    
    
}


func NewCollidableVehicleBehavior(collidable bool) *CollidableVehicleBehavior {
    this := &CollidableVehicleBehavior{}
        this.CollidableDestroyableDamageableBehavior = *NewCollidableDestroyableDamageableBehavior(collidable)
        this.initFields()
    return this
}


// throws Exception
func (this *CollidableVehicleBehavior) Collide(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var collisionTypeFactory *CollisionTypeFactory = CollisionTypeFactory_GetInstance(); _ = collisionTypeFactory
        var collisionType *CollisionType = collidableInterfaceCompositeInterface.GetCollidableInferface().GetCollisionTypeWith(ownerLayer); _ = collisionType
        if collisionType == collisionTypeFactory.PICKUP {
            var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = (any(ownerLayer).(ICollidableDestroyableDamageableLayer)); _ = collidableDestroyableDamageableLayer
            collidableDestroyableDamageableLayer.GetPickupBehavior().DoPickupLayer(any(collidableInterfaceCompositeInterface).(PickedUpLayerInterface))
        } else if collisionType == collisionTypeFactory.COLLISION {
            this.CollidableDestroyableDamageableBehavior.Collide(ownerLayer, collidableInterfaceCompositeInterface)
        } else {
            this.CollideVehicle(ownerLayer, any(collidableInterfaceCompositeInterface).(VehiclePropertiesCompositeInterface))
        }
}


func (this *CollidableVehicleBehavior) CollideInterface(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface) {
        ForcedLogUtil_Log("Don't Use Interface Version It Is Slower", this)
}


func (this *CollidableVehicleBehavior) CollideVehicle(ownerLayer ICollidableCompositeLayer, vehiclePropertiesCompositeInterface VehiclePropertiesCompositeInterface) {
        var ownerVehicleLayerInterface VehiclePropertiesCompositeInterface = (any(ownerLayer).(VehiclePropertiesCompositeInterface)); _ = ownerVehicleLayerInterface
        var ownerVehicleProperties *VehicleProperties = ownerVehicleLayerInterface.GetVehicleProperties(); _ = ownerVehicleProperties
        var vehicleProperties *VehicleProperties = vehiclePropertiesCompositeInterface.GetVehicleProperties(); _ = vehicleProperties
        var impactVelocityX int64 = ownerVehicleProperties.GetVelocityProperties().GetVelocityXBasicDecimalP().GetUnscaled() - vehicleProperties.GetVelocityProperties().GetVelocityXBasicDecimalP().GetUnscaled(); _ = impactVelocityX
        var impactVelocityY int64 = ownerVehicleProperties.GetVelocityProperties().GetVelocityYBasicDecimalP().GetUnscaled() - vehicleProperties.GetVelocityProperties().GetVelocityYBasicDecimalP().GetUnscaled(); _ = impactVelocityY
        if impactVelocityX != 0 {
            this.HalfImpactVelocityX = (impactVelocityX >> 1)
            ownerVehicleProperties.GetVelocityProperties().GetVelocityXBasicDecimalP().Addlong(this.HalfImpactVelocityX)
            vehicleProperties.GetVelocityProperties().GetVelocityXBasicDecimalP().Addlong(this.HalfImpactVelocityX)
        } else {
        }
        if impactVelocityY != 0 {
            this.HalfImpactVelocityY = (impactVelocityY >> 1)
            ownerVehicleProperties.GetVelocityProperties().GetVelocityYBasicDecimalP().Addlong(this.HalfImpactVelocityY)
            vehicleProperties.GetVelocityProperties().GetVelocityXBasicDecimalP().Addlong(this.HalfImpactVelocityY)
        } else {
        }
        this.CollideFriction(ownerVehicleLayerInterface)
        this.CollideFriction(vehiclePropertiesCompositeInterface)
        this.TotalImpactVelocity = impactVelocityX + impactVelocityY
}


func (this *CollidableVehicleBehavior) CollideFriction(vehiclePropertiesCompositeInterface VehiclePropertiesCompositeInterface) {
        var vehicleFrictionProperties *VehicleFrictionProperties = vehiclePropertiesCompositeInterface.GetVehicleProperties().GetVehicleFrictionProperties(); _ = vehicleFrictionProperties
        vehiclePropertiesCompositeInterface.GetVehicleProperties().GetVehicleFrictionProperties().Friction(vehiclePropertiesCompositeInterface.GetVehicleProperties().GetVelocityProperties(), vehicleFrictionProperties.GetCrashFrictionNominator())
}

