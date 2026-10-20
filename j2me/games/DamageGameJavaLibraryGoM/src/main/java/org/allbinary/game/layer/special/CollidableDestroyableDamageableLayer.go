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

type CollidableDestroyableDamageableLayer struct {
    CollidableCompositeLayer
    GroupInterface []*Group
    ReadyForExplosion bool
    InitWidth int
    InitHeight int
    InitX int
    InitY int
    InitZ int
    PartInterfaceArrayP []PartInterface
    PickupBehavior *PickupBehavior
    AllBinaryGameLayerManagerP IAllBinaryGameLayerManager
    IsDraggable bool
    IsDragged bool
    IsDropped bool
}

// Implements DestroyableInterface, DamageableInterface, PickupCompositeInterface, SpecialGameInputInterface

// The Java class CollidableDestroyableDamageableLayer has subclasses so its variables are this interface
type ICollidableDestroyableDamageableLayer interface {
    ICollidableCompositeLayer
    CollidableDestroyableDamageableLayer_() *CollidableDestroyableDamageableLayer
    InitPosition()
    InitPositionXYZ(x int, y int, z int)
    Paint(graphics IGraphics)
    Damage(damage int, damageType int)
    GetDamage(damageType int) int
    IsDestroyed() bool
    GetGroupInterface() []*Group
    SetGroupInterface(teamInterface []*Group)
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    IsReadyForExplosion() bool
    SetReadyForExplosion(isReadyForExplosion bool)
    GetInitWidth() int
    GetInitHeight() int
    GetInitX() int
    GetInitY() int
    Up()
    Down()
    Right()
    Left()
    StrafeLeft()
    StrafeRight()
    Fire(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special1(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special2(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special3(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special4(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special5(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    UpAnalog(yAnalogValue int)
    DownAnalog(yAnalogValue int)
    RightAnalog(xAnalogValue int)
    LeftAnalog(xAnalogValue int)
    RightTrigger(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, xAnalogValue int)
    LeftTrigger(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, xAnalogValue int)
    OnDestroyed(destroyedEvent *DestroyedEvent)
    GetPickupBehavior() *PickupBehavior
    SetPickupBehavior(pickupBehavior *PickupBehavior)
    AddPart(pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface)
    SetPartInterfaceArray(partInterfaceArray []PartInterface)
    GetPartInterfaceArray() []PartInterface
    ToStringAppend(stringBuffer *StringMaker)
    ToString() string
}

func (this *CollidableDestroyableDamageableLayer) CollidableDestroyableDamageableLayer_() *CollidableDestroyableDamageableLayer {
    return this
}

var CollidableDestroyableDamageableLayer_NULL_COLLIDABLE_DESTROYABLE_DAMAGE_LAYER any = NullUtil_GetInstance().NULL_OBJECT
var CollidableDestroyableDamageableLayer_READYFOREXPLOSION string = "ReadyForExplosion: "

func (this *CollidableDestroyableDamageableLayer) initFields() {
    
    
    
    
    
    
    
    this.PartInterfaceArrayP = PartInterfaceUtil_GetZeroArray()
    
    this.AllBinaryGameLayerManagerP = AllBinaryGameLayerManager_GetNullInstance()
    this.IsDraggable = false
    this.IsDragged = false
    this.IsDropped = false
}


func CollidableDestroyableDamageableLayer_GetNullInstance() ICollidableDestroyableDamageableLayer {
        if CollidableDestroyableDamageableLayer_NULL_COLLIDABLE_DESTROYABLE_DAMAGE_LAYER == NullUtil_GetInstance().NULL_OBJECT {
            CollidableDestroyableDamageableLayer_NULL_COLLIDABLE_DESTROYABLE_DAMAGE_LAYER = NewCollidableDestroyableDamageableLayer(BasicGroupFactory_GetInstance().NONE_ARRAY, StringUtil_GetInstance().EMPTY_STRING, RectangleFactory_SINGLETON, ViewPositionBase_NULL_VIEW_POSITION)
        }
        return any(CollidableDestroyableDamageableLayer_NULL_COLLIDABLE_DESTROYABLE_DAMAGE_LAYER).(ICollidableDestroyableDamageableLayer)
}


func NewCollidableDestroyableDamageableLayer(groupInterface []*Group, name string, layerInfo *Rectangle, viewPosition IViewPositionBase) *CollidableDestroyableDamageableLayer {
    this := &CollidableDestroyableDamageableLayer{}
        this.CollidableCompositeLayer = *NewCollidableCompositeLayer(name, layerInfo, viewPosition, CollidableNeverCollideBehaviorFactory_GetInstance())
        this.initFields()
        this.InitWidth = layerInfo.GetWidth()
        this.InitHeight = layerInfo.GetHeight()
        this.GroupInterface = groupInterface
        this.PickupBehavior = PickupBehavior_GetNullInstance()
    return this
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) InitPosition() {
        this.SetPosition(this.InitX, this.InitY, this.InitZ)
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) InitPositionXYZ(x int, y int, z int) {
        this.InitX = x
        this.InitY = y
        this.InitZ = z
}


func (this *CollidableDestroyableDamageableLayer) Paint(graphics IGraphics) {
        this.CollidableCompositeLayer.Paint(graphics)
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Damage(damage int, damageType int) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) GetDamage(damageType int) int {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) IsDestroyed() bool {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *CollidableDestroyableDamageableLayer) GetGroupInterface() []*Group {
        return this.GroupInterface
}


func (this *CollidableDestroyableDamageableLayer) SetGroupInterface(teamInterface []*Group) {
        this.GroupInterface = teamInterface
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.AllBinaryGameLayerManagerP = allBinaryGameLayerManager
        if this.AllBinaryGameLayerManagerP == nil {
            panic(errors.New(fmt.Sprint()))
        }
}


func (this *CollidableDestroyableDamageableLayer) IsReadyForExplosion() bool {
        return this.ReadyForExplosion
}


func (this *CollidableDestroyableDamageableLayer) SetReadyForExplosion(isReadyForExplosion bool) {
        this.ReadyForExplosion = isReadyForExplosion
}


func (this *CollidableDestroyableDamageableLayer) GetInitWidth() int {
        return this.InitWidth
}


func (this *CollidableDestroyableDamageableLayer) GetInitHeight() int {
        return this.InitHeight
}


func (this *CollidableDestroyableDamageableLayer) GetInitX() int {
        return this.InitX
}


func (this *CollidableDestroyableDamageableLayer) GetInitY() int {
        return this.InitY
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Up() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Down() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Right() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Left() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) StrafeLeft() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) StrafeRight() {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Fire(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Special1(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Special2(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Special3(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Special4(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) Special5(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) UpAnalog(yAnalogValue int) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) DownAnalog(yAnalogValue int) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) RightAnalog(xAnalogValue int) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) LeftAnalog(xAnalogValue int) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) RightTrigger(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, xAnalogValue int) {
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) LeftTrigger(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, xAnalogValue int) {
}


func (this *CollidableDestroyableDamageableLayer) OnDestroyed(destroyedEvent *DestroyedEvent) {
}


func (this *CollidableDestroyableDamageableLayer) GetPickupBehavior() *PickupBehavior {
        return this.PickupBehavior
}


func (this *CollidableDestroyableDamageableLayer) SetPickupBehavior(pickupBehavior *PickupBehavior) {
        this.PickupBehavior = pickupBehavior
}


// throws Exception
func (this *CollidableDestroyableDamageableLayer) AddPart(pickedUpLayerInterfaceFactoryInterface PickedUpLayerInterfaceFactoryInterface) {
}


func (this *CollidableDestroyableDamageableLayer) SetPartInterfaceArray(partInterfaceArray []PartInterface) {
        this.PartInterfaceArrayP = partInterfaceArray
}


func (this *CollidableDestroyableDamageableLayer) GetPartInterfaceArray() []PartInterface {
        return this.PartInterfaceArrayP
}


func (this *CollidableDestroyableDamageableLayer) ToStringAppend(stringBuffer *StringMaker) {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        this.CollidableCompositeLayer.ToStringAppend(stringBuffer)
        stringBuffer.Append(commonSeps.NEW_LINE)
        var groupInterfaceArray []*Group = this.GetGroupInterface(); _ = groupInterfaceArray
        var size int = len(groupInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            stringBuffer.Append(fmt.Sprint(groupInterfaceArray[index]))
            stringBuffer.Append(commonSeps.COMMA)
        }
        var pickupBehavior *PickupBehavior = this.GetPickupBehavior(); _ = pickupBehavior
        if pickupBehavior != nil {
            stringBuffer.Append(commonSeps.NEW_LINE)
            stringBuffer.Append(fmt.Sprint(pickupBehavior))
        }
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(CollidableDestroyableDamageableLayer_READYFOREXPLOSION)
        stringBuffer.Appendboolean(this.ReadyForExplosion)
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.CommonStrings.INIT)
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Appendint(this.InitX)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Appendint(this.InitX)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Appendint(this.InitZ)
}


func (this *CollidableDestroyableDamageableLayer) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        this.ToStringAppend(stringBuffer)
        return fmt.Sprint(stringBuffer)
}

func (this *CollidableDestroyableDamageableLayer) String() string {
    return this.ToString()
}

