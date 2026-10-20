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

type BasicProjectileWeaponPart struct {
    BasicWeaponPart
    WeaponLayerCircularStaticPool *WeaponLayerCircularPool
}


func (this *BasicProjectileWeaponPart) initFields() {
    
}


func BasicProjectileWeaponPart_CreateBasicProjectileWeaponPart(animationInterface IAnimation, weaponLayerCircularStaticPool *WeaponLayerCircularPool) *BasicProjectileWeaponPart {
        return NewBasicProjectileWeaponPart(animationInterface, AllBinaryLayer_NULL_ALLBINARY_LAYER, weaponLayerCircularStaticPool, WeaponProperties_NULL_WEAPON_PROPERTIES, NoScoreable_GetInstance(), RelativeRelationship_NULL_RELATIVE_RELATIONSHIP)
}


func NewBasicProjectileWeaponPart(animationInterface IAnimation, sourceLayerInterface IAllBinaryLayer, weaponLayerCircularStaticPool *WeaponLayerCircularPool, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship) *BasicProjectileWeaponPart {
    this := &BasicProjectileWeaponPart{}
        this.BasicWeaponPart = *NewBasicWeaponPart(animationInterface, sourceLayerInterface, weaponProperties, scoreableInterface, relativeRelationship)
        this.initFields()
        this.WeaponLayerCircularStaticPool = weaponLayerCircularStaticPool
    return this
}


// throws Exception
func (this *BasicProjectileWeaponPart) ProcessScore(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface) {
        var weaponLayer IWeaponLayer = this.WeaponLayerCircularStaticPool.GetInstanceWeaponLayer(this.GetOwnerLayerInterface(), this.RelativeRelationship.GetX(), this.RelativeRelationship.GetY(), this.RelativeRelationship.GetZ(), int(angle), int(otherAngle), weaponProperties, scoreableInterface); _ = weaponLayer
        allbinaryLayerManager.Append(weaponLayer)
}

