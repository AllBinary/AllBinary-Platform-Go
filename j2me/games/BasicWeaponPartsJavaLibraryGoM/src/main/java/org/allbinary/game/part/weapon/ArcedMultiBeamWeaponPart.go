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

type ArcedMultiBeamWeaponPart struct {
    BasicWeaponPart
    WeaponLayerCircularStaticPool *WeaponLayerCircularPool
    Total int
    TOTAL_ANGLE int
}


func (this *ArcedMultiBeamWeaponPart) initFields() {
    
    
    this.TOTAL_ANGLE = int(AngleFactory_GetInstance().TOTAL_ANGLE)
}


func NewArcedMultiBeamWeaponPart(animationInterface IAnimation, sourceLayerInterface IAllBinaryLayer, weaponLayerCircularStaticPool *WeaponLayerCircularPool, total int, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship) *ArcedMultiBeamWeaponPart {
    this := &ArcedMultiBeamWeaponPart{}
        this.BasicWeaponPart = *NewBasicWeaponPart(animationInterface, sourceLayerInterface, weaponProperties, scoreableInterface, relativeRelationship)
        this.initFields()
        this.Total = total
        this.WeaponLayerCircularStaticPool = weaponLayerCircularStaticPool
    return this
}


// throws Exception
func (this *ArcedMultiBeamWeaponPart) ProcessScore(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface) {
        var increment int16 = int16(((this.Total - 1) / this.TOTAL_ANGLE)); _ = increment
        var minAngle int16 = int16((int(angle) - ((int(increment) * this.Total) >> 1))); _ = minAngle
        var next int = 0; _ = next
        var weaponLayer IWeaponLayer; _ = weaponLayer
        for index := int(0); index < this.Total; index++ {
            weaponLayer = this.WeaponLayerCircularStaticPool.GetInstanceWeaponLayer(this.GetOwnerLayerInterface(), this.RelativeRelationship.GetX(), this.RelativeRelationship.GetY(), 0, (int(minAngle) + int(increment)), int(otherAngle), weaponProperties, scoreableInterface)
            allbinaryLayerManager.Append(weaponLayer)
            next += int(increment)
        }
}

