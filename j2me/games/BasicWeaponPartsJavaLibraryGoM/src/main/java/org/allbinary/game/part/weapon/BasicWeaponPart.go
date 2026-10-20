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

type BasicWeaponPart struct {
    AnimationInterface IAnimation
    OwnerLayerInterface IAllBinaryLayer
    WeaponProperties *WeaponProperties
    ScoreableInterface ScoreableInterface
    RelativeRelationship IRelativeRelationship
}

// Implements PartInterface, SalvoInterface

// The Java class BasicWeaponPart has subclasses so its variables are this interface
type IBasicWeaponPart interface {
    BasicWeaponPart_() *BasicWeaponPart
    Init(sourceLayerInterface IAllBinaryLayer, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship)
    Process(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16)
    ProcessScore(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface)
    GetOwnerLayerInterface() IAllBinaryLayer
    SetOwnerLayerInterface(ownerLayerInterface IAllBinaryLayer)
    GetAnimationInterface() IAnimation
    SetAnimationInterface(animationInterface IAnimation)
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    GetWeaponProperties() *WeaponProperties
    SetWeaponProperties(weaponProperties *WeaponProperties)
}

func (this *BasicWeaponPart) BasicWeaponPart_() *BasicWeaponPart {
    return this
}

var BasicWeaponPart_NULL_BASIC_WEAPON_PART IBasicWeaponPart = BasicWeaponPart_CreateBasicWeaponPart(NullAnimationFactory_GetFactoryInstance().GetInstance(0))

func (this *BasicWeaponPart) initFields() {
    this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
    this.OwnerLayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
    this.WeaponProperties = WeaponProperties_NULL_WEAPON_PROPERTIES
    this.ScoreableInterface = NoScoreable_GetInstance()
    this.RelativeRelationship = RelativeRelationship_NULL_RELATIVE_RELATIONSHIP
}


func BasicWeaponPart_CreateBasicWeaponPart(animationInterface IAnimation) IBasicWeaponPart {
        return NewBasicWeaponPart(animationInterface, AllBinaryLayer_NULL_ALLBINARY_LAYER, WeaponProperties_NULL_WEAPON_PROPERTIES, NoScoreable_GetInstance(), RelativeRelationship_NULL_RELATIVE_RELATIONSHIP)
}


func NewBasicWeaponPart(animationInterface IAnimation, sourceLayerInterface IAllBinaryLayer, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship) *BasicWeaponPart {
    this := &BasicWeaponPart{}
    this.initFields()
        this.Init(sourceLayerInterface, weaponProperties, scoreableInterface, relativeRelationship)
        this.SetAnimationInterface(animationInterface)
    return this
}


func (this *BasicWeaponPart) Init(sourceLayerInterface IAllBinaryLayer, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface, relativeRelationship IRelativeRelationship) {
        this.SetOwnerLayerInterface(sourceLayerInterface)
        this.SetWeaponProperties(weaponProperties)
        this.ScoreableInterface = scoreableInterface
        this.RelativeRelationship = relativeRelationship
}


// throws Exception
func (this *BasicWeaponPart) Process(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16) {
        this.ProcessScore(allbinaryLayerManager, angle, otherAngle, this.GetWeaponProperties(), this.ScoreableInterface)
}


// throws Exception
func (this *BasicWeaponPart) ProcessScore(allbinaryLayerManager IAllBinaryLayerManager, angle int16, otherAngle int16, weaponProperties *WeaponProperties, scoreableInterface ScoreableInterface) {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}


func (this *BasicWeaponPart) GetOwnerLayerInterface() IAllBinaryLayer {
        return this.OwnerLayerInterface
}


func (this *BasicWeaponPart) SetOwnerLayerInterface(ownerLayerInterface IAllBinaryLayer) {
        this.OwnerLayerInterface = ownerLayerInterface
}


func (this *BasicWeaponPart) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *BasicWeaponPart) SetAnimationInterface(animationInterface IAnimation) {
        this.AnimationInterface = animationInterface
}


func (this *BasicWeaponPart) Paint(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetOwnerLayerInterface().GetViewPosition(); _ = viewPosition
        var viewX int = viewPosition.GetX(); _ = viewX
        var viewY int = viewPosition.GetY(); _ = viewY
        this.AnimationInterface.PaintXY(graphics, viewX, viewY)
}


func (this *BasicWeaponPart) PaintThreed(graphics IGraphics) {
}


func (this *BasicWeaponPart) GetWeaponProperties() *WeaponProperties {
        return this.WeaponProperties
}


func (this *BasicWeaponPart) SetWeaponProperties(weaponProperties *WeaponProperties) {
        this.WeaponProperties = weaponProperties
}

