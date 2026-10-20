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

type WeaponLayer struct {
    MultiPlayerGameLayer
    AnimationInterface IAnimation
    InitAnimationInterface IAnimation
    DestroyedAnimationInterface IAnimation
    SourceLayerInterface IAllBinaryLayer
    TotalDamage int
    InitDamage int
    Destroyed bool
    ScoreableInterface ScoreableInterface
    Movement IMovement
    WeaponProperties *WeaponProperties
    MultiPlayerType int
    CollidableWeaponBehavior *CollidableWeaponBehavior
}

// Implements TickableInterface

// The Java class WeaponLayer has subclasses so its variables are this interface
type IWeaponLayer interface {
    IMultiPlayerGameLayer
    WeaponLayer_() *WeaponLayer
    SetCollidableInferface(collidableInferface ICollidableBaseBehavior)
    InitXYZ(x int, y int, z int)
    Init(sourceLayerInterface IAllBinaryLayer, angle int, otherAngle int, weaponProperties *WeaponProperties, scoreable ScoreableInterface)
    SetGroupInterfaces()
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    IsDestroyed() bool
    Damage(damage int, damageType int)
    GivePoints(total int)
    IsExhausted() bool
    GetDamage(damageType int) int
    SetDestroyed(destroyed bool)
    GetInitDamage() int
    SetAnimationInterface(animationInterface IAnimation)
    GetAnimationInterface() IAnimation
    GetInitAnimationInterfaceP() IAnimation
    SetInitAnimationInterfaceP(initAnimationInterface IAnimation)
    GetDestroyedAnimationInterface() IAnimation
    GetOwnerLayerInterface() IAllBinaryLayer
    SetOwnerLayerInterface(sourceLayerInterface IAllBinaryLayer)
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    GetMovement() IMovement
    SetMovement(movement IMovement)
    GetType() int
    GetMultiPlayerType() int
    SetWeaponPropertiesP(weaponProperties *WeaponProperties)
    GetWeaponPropertiesP() *WeaponProperties
    ImplmentsTickableInterface() bool
    Set(gl GL)
}

func (this *WeaponLayer) WeaponLayer_() *WeaponLayer {
    return this
}


func (this *WeaponLayer) initFields() {
    this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
    
    
    this.SourceLayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
    
    
    this.Destroyed = true
    this.ScoreableInterface = NoScoreable_GetInstance()
    
    this.WeaponProperties = WeaponProperties_NULL_WEAPON_PROPERTIES
    
    this.CollidableWeaponBehavior = CollidableWeaponBehavior_GetNullInstance()
}


func NewWeaponLayer(name string, remoteInfo *RemoteInfo, multiPlayerType int, movement IMovement, animationInterface IAnimation, destroyedAnimationInterface IAnimation, rectangle *Rectangle, viewPosition IViewPositionBase) *WeaponLayer {
    this := &WeaponLayer{}
        this.MultiPlayerGameLayer = *NewMultiPlayerGameLayer(remoteInfo, BasicGroupFactory_GetInstance().NONE_ARRAY, name, rectangle, viewPosition)
        this.initFields()
        this.InitAnimationInterface = animationInterface
        this.SetAnimationInterface(animationInterface)
        this.DestroyedAnimationInterface = destroyedAnimationInterface
        this.Movement = movement
        this.MultiPlayerType = multiPlayerType
    return this
}


func (this *WeaponLayer) SetCollidableInferface(collidableInferface ICollidableBaseBehavior) {
        this.MultiPlayerGameLayer.SetCollidableInferface(collidableInferface)
        this.CollidableWeaponBehavior = any(this.GetCollidableInferface()).(*CollidableWeaponBehavior)
}


// throws Exception
func (this *WeaponLayer) InitXYZ(x int, y int, z int) {
        this.X = x
        this.Y = y
        this.Z = z
}


// throws Exception
func (this *WeaponLayer) Init(sourceLayerInterface IAllBinaryLayer, angle int, otherAngle int, weaponProperties *WeaponProperties, scoreable ScoreableInterface) {
        this.SetWeaponPropertiesP(weaponProperties)
        this.SetReadyForExplosion(false)
        this.SetAnimationInterface(this.GetInitAnimationInterfaceP())
        this.SetOwnerLayerInterface(sourceLayerInterface)
        this.CollidableWeaponBehavior.Init(sourceLayerInterface)
        this.SetGroupInterfaces()
        if scoreable != nil {
            this.ScoreableInterface = scoreable
        }
        this.TotalDamage = 0
        this.InitDamage = weaponProperties.GetDamage()
        this.SetDestroyed(false)
        this.Movement.Init(weaponProperties.GetSpeed(), angle, otherAngle)
}


func (this *WeaponLayer) SetGroupInterfaces() {
        var basicGroupFactory *BasicGroupFactory = BasicGroupFactory_GetInstance(); _ = basicGroupFactory
        if this.SourceLayerInterface.GetGroupInterface()[0] == basicGroupFactory.ENEMY {
            this.SetGroupInterface(basicGroupFactory.ENEMY_ARRAY)
        } else {
            this.SetGroupInterface(this.SourceLayerInterface.GetGroupInterface())
        }
}


// throws Exception
func (this *WeaponLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.IsExhausted() && !this.IsDestroyed() {
            if this.IsReadyForExplosion() {
                var indexedAnimationInterface IIndexedAnimation = any(this.GetAnimationInterface()).(IIndexedAnimation); _ = indexedAnimationInterface
                var currentFrame int = indexedAnimationInterface.GetFrame(); _ = currentFrame
                if currentFrame < indexedAnimationInterface.GetAnimationSize() - 1 {
                    indexedAnimationInterface.NextFrame()
                } else {
                    this.SetDestroyed(true)
                }
            } else {
                var collidableWeaponBehavior *CollidableWeaponBehavior = any(this.GetCollidableInferface()).(*CollidableWeaponBehavior); _ = collidableWeaponBehavior
                if collidableWeaponBehavior.IsCollided() {
                    this.SetAnimationInterface(this.DestroyedAnimationInterface)
                    this.GetMovement().Stop()
                    this.SetReadyForExplosion(true)
                } else {
                    this.SetDestroyed(true)
                }
            }
        }
        this.Movement.Process(this)
        this.TotalDamage += int(this.WeaponProperties.GetDissipation())
}


func (this *WeaponLayer) IsDestroyed() bool {
        return this.Destroyed
}


func (this *WeaponLayer) Damage(damage int, damageType int) {
        this.TotalDamage += damage * damage
}


func (this *WeaponLayer) GivePoints(total int) {
        this.ScoreableInterface.AddPoints(GameConfigurationUtil_GetInstance().GetCompetitionValue() * total)
}


func (this *WeaponLayer) IsExhausted() bool {
        if this.TotalDamage > this.GetInitDamage() || this.TotalDamage < 0 {
            return true
        } else {
            return false
        }
}


func (this *WeaponLayer) GetDamage(damageType int) int {
        if !this.IsExhausted() {
            var total int = this.GetInitDamage() - this.TotalDamage; _ = total
            this.GivePoints(total)
            return total
        }
        return 0
}


func (this *WeaponLayer) SetDestroyed(destroyed bool) {
        this.Destroyed = destroyed
        if this.IsDestroyed() {
            this.ScoreableInterface = NoScoreable_GetInstance()
            DestroyedLayerProcessor_GetInstance().Add(this)
        }
}


func (this *WeaponLayer) GetInitDamage() int {
        return this.InitDamage
}


func (this *WeaponLayer) SetAnimationInterface(animationInterface IAnimation) {
        this.AnimationInterface = animationInterface
}


func (this *WeaponLayer) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *WeaponLayer) GetInitAnimationInterfaceP() IAnimation {
        return this.InitAnimationInterface
}


func (this *WeaponLayer) SetInitAnimationInterfaceP(initAnimationInterface IAnimation) {
        this.InitAnimationInterface = initAnimationInterface
}


func (this *WeaponLayer) GetDestroyedAnimationInterface() IAnimation {
        return this.DestroyedAnimationInterface
}


func (this *WeaponLayer) GetOwnerLayerInterface() IAllBinaryLayer {
        return this.SourceLayerInterface
}


func (this *WeaponLayer) SetOwnerLayerInterface(sourceLayerInterface IAllBinaryLayer) {
        this.SourceLayerInterface = sourceLayerInterface
}


func (this *WeaponLayer) Paint(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        this.AnimationInterface.PaintXY(graphics, viewPosition.GetX() - this.GetHalfWidth(), viewPosition.GetY() - this.GetHalfHeight())
}


func (this *WeaponLayer) PaintThreed(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        this.AnimationInterface.PaintThreedXYZ(graphics, viewPosition.GetX() - this.GetHalfWidth(), viewPosition.GetY() - this.GetHalfHeight(), viewPosition.GetZ() - this.GetHalfDepth())
}


func (this *WeaponLayer) GetMovement() IMovement {
        return this.Movement
}


func (this *WeaponLayer) SetMovement(movement IMovement) {
        this.Movement = movement
}


func (this *WeaponLayer) GetType() int {
        return WeaponLayer_GetStaticType()
}


func WeaponLayer_GetStaticType() int {
        return 0
}


func (this *WeaponLayer) GetMultiPlayerType() int {
        return this.MultiPlayerType
}


func (this *WeaponLayer) SetWeaponPropertiesP(weaponProperties *WeaponProperties) {
        this.WeaponProperties = weaponProperties
}


func (this *WeaponLayer) GetWeaponPropertiesP() *WeaponProperties {
        return this.WeaponProperties
}


func (this *WeaponLayer) ImplmentsTickableInterface() bool {
        return true
}


// throws Exception
func (this *WeaponLayer) Set(gl GL) {
        this.InitAnimationInterface.Set(gl)
}

