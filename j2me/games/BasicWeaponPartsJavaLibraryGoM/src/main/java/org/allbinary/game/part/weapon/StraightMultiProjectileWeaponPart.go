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

type StraightMultiProjectileWeaponPart struct {
    BasicWeaponPart
    WeaponLayerCircularStaticPool *WeaponLayerCircularPool
    Total int
    NoDecimalTrigTable *NoDecimalTrigTable
}


func (this *StraightMultiProjectileWeaponPart) initFields() {
    
    
    this.NoDecimalTrigTable = NoDecimalTrigTable_GetInstance()
}


func StraightMultiProjectileWeaponPart_CreateStraightMultiProjectileWeaponPart(animationInterface IAnimation, weaponLayerCircularStaticPool *WeaponLayerCircularPool) *StraightMultiProjectileWeaponPart {
        return NewStraightMultiProjectileWeaponPart(animationInterface, AllBinaryLayer_NULL_ALLBINARY_LAYER, weaponLayerCircularStaticPool, 2, WeaponProperties_NULL_WEAPON_PROPERTIES, NoScoreable_GetInstance(), RelativeRelationship_NULL_RELATIVE_RELATIONSHIP)
}


func NewStraightMultiProjectileWeaponPart(animationInterface IAnimation, sourceLayerInterface IAllBinaryLayer, weaponLayerCircularStaticPool *WeaponLayerCircularPool, total int, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship) *StraightMultiProjectileWeaponPart {
    this := &StraightMultiProjectileWeaponPart{}
        this.BasicWeaponPart = *NewBasicWeaponPart(animationInterface, sourceLayerInterface, weaponProperties, scoreableInterface, relativeRelationship)
        this.initFields()
        this.Total = total
        this.WeaponLayerCircularStaticPool = weaponLayerCircularStaticPool
    return this
}


// throws Exception
func (this *StraightMultiProjectileWeaponPart) ProcessScore(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface) {
        var sourceLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = sourceLayerInterface
        var reducedWidth int = (sourceLayerInterface.GetWidth() * 8 / 10); _ = reducedWidth
        var halfWidth int = (reducedWidth >> 1); _ = halfWidth
        var x int = this.RelativeRelationship.GetX(); _ = x
        var y int = this.RelativeRelationship.GetY(); _ = y
        var sine int64; _ = sine
        var beamX int; _ = beamX
        var cosine int64; _ = cosine
        var beamY int; _ = beamY
        var weaponLayer IWeaponLayer; _ = weaponLayer
        var increment int = reducedWidth / (this.Total - 1); _ = increment
        var next int = -halfWidth; _ = next
        for index := int(0); index < this.Total; index++ {
            sine = (int64(next) * this.NoDecimalTrigTable.Sin(int(angle)))
            beamX = int((sine / int64(this.NoDecimalTrigTable.SCALE)))
            cosine = (int64(next) * this.NoDecimalTrigTable.Cos(int(angle)))
            beamY = -int((cosine / int64(this.NoDecimalTrigTable.SCALE)))
            weaponLayer = this.WeaponLayerCircularStaticPool.GetInstanceWeaponLayer(sourceLayerInterface, x + beamX, y + beamY, 0, int(angle), int(otherAngle), weaponProperties, scoreableInterface)
            allbinaryLayerManager.Append(weaponLayer)
            next += increment
        }
}

