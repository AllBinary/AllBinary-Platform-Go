/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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
    . "org/allbinary/game/ag"
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

type UnitLayer struct {
    AdvancedRTSGameLayer
    BasicColorFactory *BasicColorFactory
    GroupCommonFactory *GroupCommonFactory
    LayerPartialCellPositionsUtil *LayerPartialCellPositionsUtil
    LayerDistanceUtil *LayerDistanceUtil
    AngleFactory *AngleFactory
    CellPositionFactory *CellPositionFactory
    MaxResourceLoad int16
    VehicleProperties *VehicleProperties
    AccelerationBasicDecimal *BasicDecimal
    DecelerationBasicDecimal *BasicDecimal
    TrackingEvent *TrackingEvent
    DamageFloaters IDamageFloaters
    DamageFloatersPaintableInterface IPaintable
    HealthBar IPaintable
    MoveSoundInterface ISound
    InitResourceAnimation IRotationAnimation
    ResourceAnimation IIndexedAnimation
    InitPathAnimation *PathAnimation
    PathAnimation IAnimation
    CaptionAnimationHelper ICaptionAnimationHelperBase
    DecalAnimation IRotationAnimation
    WaypointLayerInterfaceFactoryInterface LayerInterfaceFactoryInterface
    ResourceLoad int
    WeaponRange int
    RtsLogHelper IRTSLayerLogHelper
    RotationAnimationInterfaceP IRotationAnimation
    MovementAngle *NamedAngle
    SteeringInsideGeographicMapCellPosition IGeographicMapCellPosition
    SensorGeographicMapCellPositionList IBasicArrayList
    CAPITAL_EVENT *CapitalEvent
}

// Implements BuildingEventListenerInterface, TrackingEventListenerInterface

var UnitLayer_PartialPositionList IBasicArrayList = NewBasicArrayListS(4)
var UnitLayer_SENSOR_RANGE_MULTIPLIER int = 6
var UnitLayer_MOVE string = "Moving"

func (this *UnitLayer) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.GroupCommonFactory = GroupCommonFactory_GetInstance()
    this.LayerPartialCellPositionsUtil = LayerPartialCellPositionsUtil_GetInstance()
    this.LayerDistanceUtil = LayerDistanceUtil_GetInstance()
    this.AngleFactory = AngleFactory_GetInstance()
    this.CellPositionFactory = CellPositionFactory_GetInstance()
    
    
    this.AccelerationBasicDecimal = NewBasicDecimal(1600)
    this.DecelerationBasicDecimal = NewBasicDecimal(-1000)
    
    
    
    
    
    
    this.ResourceAnimation = any(NullIndexedAnimationFactory_GetFactoryInstance().GetInstance(0)).(IIndexedAnimation)
    
    this.PathAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
    this.CaptionAnimationHelper = NewCaptionAnimationHelper(FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(CaptionResources_GetInstance().RESOURCE).GetInstance(0), -23, -25, 6, 0)
    
    
    this.ResourceLoad = 0
    
    this.RtsLogHelper = RTSLayerLogHelper_GetInstance()
    
    this.MovementAngle = this.AngleFactory.NOT_ANGLE
    this.SteeringInsideGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.SensorGeographicMapCellPositionList = NewBasicArrayListD()
    this.CAPITAL_EVENT = NewCapitalEvent(this)
}


func NewUnitLayer(remoteInfo *RemoteInfo, parentLayer PathFindingLayerInterface, advancedRTSProperties *AdvancedRTSProperties, groupInterface []*Group, rootName string, name string, vehicleProperties *VehicleProperties, healthInterface *Health, maxResourceLoad int, moveSoundInterface ISound, waypointLayerInterfaceFactoryInterface LayerInterfaceFactoryInterface, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, decalAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, resourceAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, direction *Direction, x int, y int, z int, viewPosition IViewPositionBase) *UnitLayer {
    this := &UnitLayer{}
        this.AdvancedRTSGameLayer = *NewAdvancedRTSGameLayer(remoteInfo, parentLayer, advancedRTSProperties, groupInterface, rootName, name, healthInterface, NewWaypointRTSFormInput(groupInterface, true), animationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface, rectangle, x, y, viewPosition)
        this.initFields()
        this.SetCollidableInferface(NewCollidableUnitBehavior(true))
        this.WaypointLayerInterfaceFactoryInterface = waypointLayerInterfaceFactoryInterface
        this.MaxResourceLoad = maxResourceLoad.ShortValue()
        this.MoveSoundInterface = moveSoundInterface
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().DAMAGE_FLOATERS) {
            this.DamageFloaters = NewPtsDamageFloaters(this)
            this.DamageFloatersPaintableInterface = this.DamageFloaters
        } else {
            this.DamageFloatersPaintableInterface = NullPaintable_GetInstance()
            this.DamageFloaters = NewDamageFloaters()
        }
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().HEALTH_BARS) {
            this.HealthBar = NewHealthBar(this, this.GetHealthInterface(), NewHealthBarTwodAnimation(any(this).(IAllBinaryLayer), BasicHudFactory_GetInstance().BOTTOMLEFT), -1)
        } else {
            this.HealthBar = NullPaintable_GetInstance()
        }
        this.DecalAnimation = any(decalAnimationInterfaceFactoryInterface.GetInstance(0)).(IRotationAnimation)
        this.InitResourceAnimation = any(resourceAnimationInterfaceFactoryInterface.GetInstance(0)).(IRotationAnimation)
        this.InitResourceAnimation.SetFrameByDirection(direction)
        this.DecalAnimation.SetFrameByDirection(direction)
        this.RotationAnimationInterfaceP = any(this.IndexedButShouldBeRotationAnimationInterface).(IRotationAnimation)
        this.RotationAnimationInterfaceP.SetFrameByDirection(direction)
        this.SetMaxLevel(12)
        this.VehicleProperties = vehicleProperties
        this.TrackingEvent = NewTrackingEvent(this)
        this.InitPathAnimation = NewPathAnimation(this, LinePathRelativeAnimation_GetInstance())
    return this
}


// throws Exception
func (this *UnitLayer) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.AdvancedRTSGameLayer.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        this.InitPathAnimation.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
}


// throws Exception
func (this *UnitLayer) UpdateWaypointBehavior(geographicMapInterface IBasicGeographicMap) {
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        hashtable.Put(this.GroupCommonFactory.ID, this.GetGroupInterface())
        hashtable.Put(Layer_ID, this)
        hashtable.Put(AllBinaryGameLayerManager_ID, this.AllBinaryGameLayerManagerP)
        this.SetWaypointBehavior(NewUnitWaypointBehavior2(this, any(this.WaypointLayerInterfaceFactoryInterface.GetNextInstance(hashtable, this.X, this.Y, this.Z)).(IAdvancedRTSGameLayer)))
        var features *Features = Features_GetInstance(); _ = features
        var waypoint IWaypointBase = func() any { if J2MEUtil_IsHTML() { return NewMultipassNoCacheWaypoint(this, AttackSound_GetInstance()) }; return NewNoCacheWaypoint(this, AttackSound_GetInstance()) }(); _ = waypoint
        this.GetWaypointBehavior().SetWaypoint(waypoint)
        this.AdvancedRTSGameLayer.UpdateWaypointBehavior(geographicMapInterface)
        this.InitRangeHack()
}


// throws Exception
func (this *UnitLayer) Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        this.AdvancedRTSGameLayer.Construct(rtsPlayerLayerInterface)
        TrackingEventHandler_GetInstance().AddListenerInterface(this)
        WaypointEventHandlerFactory_GetInstance(this.GetGroupInterface()[0]).AddListenerInterface(this.GetUnitWaypointBehavior())
        BuildingEventHandler_GetInstance().AddListenerInterface(this)
        this.GetUnitWaypointBehavior().SetCurrentPathGeographicMapCellPosition(this.GetCurrentGeographicMapCellPosition())
        this.UpdateSensorGeographicMapCellPositionList()
        this.GetUnitWaypointBehavior().SetLastPathGeographicMapCellPosition(this.GetUnitWaypointBehavior().GetCurrentPathGeographicMapCellPosition())
}


// throws Exception
func (this *UnitLayer) UpdateSensorGeographicMapCellPositionList() {
        if VisibleCellPositionsSingleton_GetInstance().ShouldProcess() {
            var currentGeographicMapCellPosition IGeographicMapCellPosition = this.GetCurrentGeographicMapCellPosition(); _ = currentGeographicMapCellPosition
            this.SensorGeographicMapCellPositionList.Clear()
            this.SensorGeographicMapCellPositionList.Add(currentGeographicMapCellPosition)
            var sensorRange int = this.WeaponRange * UnitLayer_SENSOR_RANGE_MULTIPLIER; _ = sensorRange
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
            var totalCells int = (sensorRange / tiledLayer.GetCellHeight()) / 2; _ = totalCells
            var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
            var column int = currentGeographicMapCellPosition.GetColumn(); _ = column
            var row int = currentGeographicMapCellPosition.GetRow(); _ = row
            var lastColumn int = column + totalCells; _ = lastColumn
            var lastRow int = row + totalCells; _ = lastRow
            if lastColumn > tiledLayer.GetColumns() {
                lastColumn = tiledLayer.GetColumns()
            }
            if lastRow > tiledLayer.GetRows() {
                lastRow = tiledLayer.GetRows()
            }
            var firstColumn int = column - totalCells; _ = firstColumn
            var firstRow int = row - totalCells; _ = firstRow
            if firstColumn < 0 {
                firstColumn = 0
            }
            if firstRow < 0 {
                firstRow = 0
            }
            for index := int(lastColumn - 1); index >= firstColumn; index-- {
                for index2 := int(lastRow - 1); index2 >= firstRow; index2-- {
                    var geographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(index, index2); _ = geographicMapCellPosition
                    if !this.SensorGeographicMapCellPositionList.Contains(geographicMapCellPosition) {
                        this.SensorGeographicMapCellPositionList.Add(geographicMapCellPosition)
                    }
                }
            }
        }
}


func (this *UnitLayer) GetSensorGeographicMapCellPositionList() IBasicArrayList {
        return this.SensorGeographicMapCellPositionList
}


func (this *UnitLayer) Select() {
        this.PathAnimation = this.InitPathAnimation
        this.AdvancedRTSGameLayer.Select()
}


func (this *UnitLayer) Deselect() {
        this.PathAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
        this.AdvancedRTSGameLayer.Deselect()
}


func (this *UnitLayer) SetSelected(selected bool) {
        this.AdvancedRTSGameLayer.SetSelected(selected)
        if selected {
            if this.Debug {
                this.RtsLogHelper = RTSLayerSelectedLogHelper_GetInstance()
            }
            this.RtsLayer2LogHelper = RTSLayer2SelectedLogHelper_GetInstance()
        } else {
            if this.Debug {
                this.RtsLogHelper = RTSLayerLogHelper_GetInstance()
            }
            this.RtsLayer2LogHelper = RTSLayer2LogHelper_GetInstance()
        }
}


// throws Exception
func (this *UnitLayer) SetClosestGeographicMapCellHistory(pathsList IBasicArrayList) {
        this.RtsLogHelper.SetClosestGeographicMapCellHistory(this, pathsList)
        var closestIndex int = -1; _ = closestIndex
        var shortestDistance int = Integer_MAX_VALUE; _ = shortestDistance
        var currentDistance int = Integer_MAX_VALUE; _ = currentDistance
        for index := int(pathsList.Size() - 1); index >= 0; index-- {
            var geographicMapCellPositionBasicArrayList IBasicArrayList = any(pathsList.Get(index)).(IBasicArrayList); _ = geographicMapCellPositionBasicArrayList
            var geographicMapCellPosition IGeographicMapCellPosition = any(geographicMapCellPositionBasicArrayList.Get(geographicMapCellPositionBasicArrayList.Size() - 1)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
            currentDistance = this.LayerDistanceUtil.GetDistanceAt(this, geographicMapCellPosition.GetMidPoint())
            if currentDistance < shortestDistance {
                shortestDistance = currentDistance
                closestIndex = index
            }
        }
        if closestIndex >= 0 {
            var geographicMapCellPositionBasicArrayList IBasicArrayList = any(pathsList.Get(closestIndex)).(IBasicArrayList); _ = geographicMapCellPositionBasicArrayList
            var geographicMapCellPosition IGeographicMapCellPosition = any(geographicMapCellPositionBasicArrayList.Get(0)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
            this.TeleportTo(geographicMapCellPosition)
        }
}


// throws Exception
func (this *UnitLayer) Init(geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPositionBasicArrayList IBasicArrayList) {
        geographicMapCellHistory.TrackAll(geographicMapCellPositionBasicArrayList)
}


func (this *UnitLayer) OnMovement(trackingEvent *TrackingEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onMovement", e)
                }
            }()
            var layerInterface IAdvancedRTSGameLayer = any(trackingEvent.GetLayerInterface()).(IAdvancedRTSGameLayer); _ = layerInterface
            if layerInterface.GetGroupInterface()[0] != this.GetGroupInterface()[0] {
                this.GetUnitWaypointBehavior().GetPossibleTargetList().Add(layerInterface)
                layerInterface.OnMovementFound(this.GetTrackingEvent())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *UnitLayer) OnMovementFound(trackingEvent *TrackingEvent) {
        var layerInterface IAdvancedRTSGameLayer = any(trackingEvent.GetLayerInterface()).(IAdvancedRTSGameLayer); _ = layerInterface
        if layerInterface.GetGroupInterface()[0] != this.GetGroupInterface()[0] {
            this.GetUnitWaypointBehavior().GetPossibleTargetList().Add(layerInterface)
        }
}


func (this *UnitLayer) InitRangeHack() {
        var basicWeaponPart IBasicWeaponPart = any(this.GetPartInterfaceArray()[0]).(IBasicWeaponPart); _ = basicWeaponPart
        var weaponProperties *WeaponProperties = basicWeaponPart.GetWeaponProperties(); _ = weaponProperties
        this.WeaponRange = weaponProperties.GetRange()
        this.InitRangeAnimation = AdjustedCircleAnimation_CreateW(this.WeaponRange, this.WeaponRange, this.GetWidth(), this.BasicColorFactory.GREEN)
        var sensorRange int = this.WeaponRange * UnitLayer_SENSOR_RANGE_MULTIPLIER; _ = sensorRange
        this.InitSensorRangeAnimation = AdjustedCircleAnimation_CreateW(sensorRange, sensorRange, this.GetWidth(), this.BasicColorFactory.RED)
        this.GetUnitWaypointBehavior().InitRange(this.WeaponRange)
        this.FireTimeHelper.TimeDelayHelper_().Delay = (int(weaponProperties.GetReloadTime()))
}


// throws Exception
func (this *UnitLayer) ProcessBuiltTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if !this.GetHealthInterface().IsAlive() {
            if this.IsReadyForExplosion() {
                var currentFrame int = this.DestroyAnimationInterface.GetFrame(); _ = currentFrame
                var size int = this.DestroyAnimationInterface.GetSize() - 1; _ = size
                if currentFrame == size {
                    if !this.GetHealthInterface().IsAlive() {
                        this.SetDestroyed(true)
                    }
                } else {
                    this.DestroyAnimationInterface.NextFrame()
                }
            } else {
                this.SetAnimationInterface(this.DestroyAnimationInterface)
                SecondaryPlayerQueueFactory_GetInstance().Add(ExplosionBasicSound_GetInstance())
                this.ShakeListener.OnSmallShakeEvent()
                this.Vibration.Vibrate(this.Duration, 0, 0)
                this.SetReadyForExplosion(true)
            }
        } else {
            this.AdvancedRTSGameLayer.ProcessBuiltTick(allBinaryLayerManager)
        }
        this.CaptionAnimationHelper.Tick()
        this.GetUnitWaypointBehavior().ProcessTick(allBinaryLayerManager)
}


func (this *UnitLayer) TeleportTo(geographicMapCellPosition IGeographicMapCellPosition) {
        var point IGPoint = geographicMapCellPosition.GetMidPoint(); _ = point
        this.SetPosition(point.GetX() - this.GetHalfWidth(), point.GetY() - this.GetHalfHeight(), this.Z)
}


// throws Exception
func (this *UnitLayer) GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition {
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var geographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(this.X + this.GetHalfWidth(), this.Y + this.GetHalfHeight()); _ = geographicMapCellPosition
        var raceTrackGeographicMap *RaceTrackGeographicMap = any(geographicMapInterface).(*RaceTrackGeographicMap); _ = raceTrackGeographicMap
        if !raceTrackGeographicMap.IsValid(geographicMapCellPosition) {
            panic(errors.New(fmt.Sprint("Position is not really on the map: " + fmt.Sprint(geographicMapCellPosition))))
        }
        return geographicMapCellPosition
}


// throws Exception
func (this *UnitLayer) Fire(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
        if this.FireTimeHelper.IsTimeTNT() {
            this.FireAll(layerManager)
        } else {
            this.Reload()
        }
}


// throws Exception
func (this *UnitLayer) Left() {
        this.InitResourceAnimation.PreviousRotation()
        this.DecalAnimation.PreviousRotation()
        this.RotationAnimationInterfaceP.PreviousRotation()
}


// throws Exception
func (this *UnitLayer) Right() {
        this.InitResourceAnimation.NextRotation()
        this.DecalAnimation.NextRotation()
        this.RotationAnimationInterfaceP.NextRotation()
}


func (this *UnitLayer) Down() {
        var velocityProperties *VelocityProperties = this.GetVehicleProperties().GetVelocityProperties(); _ = velocityProperties
        if !velocityProperties.IsOverXYMaxForwardVelocity() {
            this.Accelerate(this.DecelerationBasicDecimal)
        }
}


func (this *UnitLayer) Up() {
        var velocityProperties *VelocityProperties = this.GetVehicleProperties().GetVelocityProperties(); _ = velocityProperties
        if !velocityProperties.IsOverXYMaxForwardVelocity() {
            this.Accelerate(this.AccelerationBasicDecimal)
        }
}


func (this *UnitLayer) InitInputProcessors() {
        this.InputProcessorArray[Canvas_RIGHT] = NewSpecialRightGameInputProcessor(this)
        this.InputProcessorArray[Canvas_LEFT] = NewSpecialLeftGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_NUM0] = NewSpecialFireGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_POUND] = this.InputProcessorArray[Canvas_KEY_NUM0]
        this.InputProcessorArray[Canvas_DOWN] = NewSpecialDownGameInputProcessor(this)
        this.InputProcessorArray[Canvas_UP] = NewSpecialUpGameInputProcessor(this)
        this.AdvancedRTSGameLayer.InitInputProcessors()
}


// throws Exception
func (this *UnitLayer) ProcessInput(layerManager IAllBinaryLayerManager) {
        var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = list.Get(index); _ = anyType
            var key int = GameKeyEventUtil_GetKey(anyType); _ = key
            this.InputProcessorArray[key].ProcessEvent(layerManager, GameKeyEvent_NONE)
        }
        list.Clear()
        this.GroundFriction()
        this.Move()
}


func (this *UnitLayer) Accelerate(accelerate *BasicDecimal) {
        this.GetVehicleProperties().GetVelocityProperties().AddVelocityi(accelerate.GetUnscaled(), int(this.RotationAnimationInterfaceP.GetAngleInfoP().GetAngle()), 90)
}


// throws Exception
func (this *UnitLayer) FireAll(layerManager IAllBinaryLayerManager) {
        var angleInfo *AngleInfo = this.RotationAnimationInterfaceP.GetAngleInfoP(); _ = angleInfo
        var angle int = (int(angleInfo.GetAngle()) + this.SlightAngle); _ = angle
        this.Hashtable.Put(SmallIntegerSingletonFactory_GetInstance().GetAt(1), SmallIntegerSingletonFactory_GetInstance().GetAt(int(AngleFactory_GetInstance().GetAt(angle).GetValue())))
        var salvoInterface SalvoInterface = any(this.GetPartInterfaceArray()[0]).(SalvoInterface); _ = salvoInterface
        salvoInterface.Process(layerManager, int16(angle), int16(90))
}


func (this *UnitLayer) Downgrade() {
        if this.GetLevel() > 1 {
            this.AdvancedRTSGameLayer.Downgrade()
        }
}


func (this *UnitLayer) Upgrade() {
        this.AdvancedRTSGameLayer.Upgrade()
        this.InitRangeHack()
}


func (this *UnitLayer) GroundFriction() {
        var vehicleFrictionProperties *VehicleFrictionProperties = this.GetVehicleProperties().GetVehicleFrictionProperties(); _ = vehicleFrictionProperties
        this.GetVehicleProperties().GetVehicleFrictionProperties().Friction(this.GetVehicleProperties().GetVelocityProperties(), vehicleFrictionProperties.GetTireFrictionNominator())
}


// throws Exception
func (this *UnitLayer) TrackTo(reason string) {
        var waypointBehaviorBase IWaypointBehaviorBase = any(this.WaypointBehaviorBase).(IWaypointBehaviorBase); _ = waypointBehaviorBase
        var nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition = waypointBehaviorBase.GetNextUnvisitedPathGeographicMapCellPosition(); _ = nextUnvisitedPathGeographicMapCellPosition
        var point IGPoint = nextUnvisitedPathGeographicMapCellPosition.GetMidPoint(); _ = point
        var dx int = (this.GetXP() + this.GetHalfWidth()) - point.GetX(); _ = dx
        var dy int = (this.GetYP() + this.GetHalfHeight()) - point.GetY(); _ = dy
        this.RtsLogHelper.TrackTo(this, nextUnvisitedPathGeographicMapCellPosition, dx, dy, reason)
        this.TrackToDXY(dx, dy)
}


// throws Exception
func (this *UnitLayer) TrackToDXY(dx int, dy int) {
        var angleOfTarget int = 0; _ = angleOfTarget
        this.TrackToDXYTargetAngle(dx, dy, angleOfTarget)
}


// throws Exception
func (this *UnitLayer) TurnTo(dx int, dy int, targetAngle int) bool {
        var waypointBehaviorBase IWaypointBehaviorBase = any(this.WaypointBehaviorBase).(IWaypointBehaviorBase); _ = waypointBehaviorBase
        var nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition = waypointBehaviorBase.GetNextUnvisitedPathGeographicMapCellPosition(); _ = nextUnvisitedPathGeographicMapCellPosition
        var evading bool = false; _ = evading
        if this.GetUnitWaypointBehavior().GetSensorAction() == SensorActionFactory_GetInstance().EVADE {
            this.RtsLogHelper.Evade(this)
            evading = true
            targetAngle += 180
        }
        var angleInfo *AngleInfo = this.RotationAnimationInterfaceP.GetAngleInfoP(); _ = angleInfo
        var angle int = FrameUtil_GetInstance().AdjustAngleToFrameAngle(int(angleInfo.GetAngle()) - 270); _ = angle
        this.RtsLogHelper.TurnTo(this, dx, dy, angleInfo, angle, this.MovementAngle, evading, targetAngle)
        var gameKeyEventFactory *GameKeyEventFactory = GameKeyEventFactory_GetInstance(); _ = gameKeyEventFactory
        if dx == 0 && dy == 0 {
            this.RtsLogHelper.DoneMoving(this)
            return true
        } else if (int(this.MovementAngle.GetValue())) == angle {
            if dx > 0 && this.MovementAngle == this.AngleFactory.LEFT {
                this.RtsLogHelper.MovingLeft(this)
                return false
            }
            if dx < 0 && this.MovementAngle == this.AngleFactory.RIGHT {
                this.RtsLogHelper.MovingRight(this)
                return false
            }
            if dy > 0 && this.MovementAngle == this.AngleFactory.UP {
                this.RtsLogHelper.MovingUp(this)
                return false
            }
            if dy < 0 && this.MovementAngle == this.AngleFactory.DOWN {
                this.RtsLogHelper.MovingDown(this)
                return false
            }
            this.RtsLogHelper.CurrentMoveEnded(this)
            if this.MovementAngle == this.AngleFactory.LEFT || this.MovementAngle == this.AngleFactory.RIGHT {
                this.HandleDeltalY(dx, dy)
            } else if this.MovementAngle == this.AngleFactory.UP || this.MovementAngle == this.AngleFactory.DOWN {
                this.HandleDeltalX(dx, dy)
            }
            return true
        } else {
            if nextUnvisitedPathGeographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
                if this.SteeringInsideGeographicMapCellPosition != nextUnvisitedPathGeographicMapCellPosition {
                    if max(dx, -(dx)) > max(dy, -(dy)) && dy != 0 {
                        this.HandleDeltalY(dx, dy)
                    } else if dx != 0 {
                        this.HandleDeltalX(dx, dy)
                    } else {
                        this.HandleDeltalY(dx, dy)
                    }
                }
                var deltaAngle2 int = int(this.MovementAngle.GetValue()) - angle; _ = deltaAngle2
                if deltaAngle2 > 0 {
                    this.RtsLogHelper.RotateRight(this)
                    this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_RIGHT))
                } else {
                    this.RtsLogHelper.RotateLeft(this)
                    this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_LEFT))
                }
                return true
            } else {
                this.RtsLogHelper.NoRotation(this)
            }
            return false
        }
}


func (this *UnitLayer) HandleDeltalX(dx int, dy int) {
        var waypointBehaviorBase IWaypointBehaviorBase = any(this.WaypointBehaviorBase).(IWaypointBehaviorBase); _ = waypointBehaviorBase
        var nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition = waypointBehaviorBase.GetNextUnvisitedPathGeographicMapCellPosition(); _ = nextUnvisitedPathGeographicMapCellPosition
        if dx > 0 {
            this.MovementAngle = this.AngleFactory.LEFT
            this.SteeringInsideGeographicMapCellPosition = nextUnvisitedPathGeographicMapCellPosition
        } else {
            this.MovementAngle = this.AngleFactory.RIGHT
            this.SteeringInsideGeographicMapCellPosition = nextUnvisitedPathGeographicMapCellPosition
        }
        this.RtsLogHelper.Handle(this, this.MovementAngle)
}


func (this *UnitLayer) HandleDeltalY(dx int, dy int) {
        var waypointBehaviorBase IWaypointBehaviorBase = any(this.WaypointBehaviorBase).(IWaypointBehaviorBase); _ = waypointBehaviorBase
        var nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition = waypointBehaviorBase.GetNextUnvisitedPathGeographicMapCellPosition(); _ = nextUnvisitedPathGeographicMapCellPosition
        if dy > 0 {
            this.MovementAngle = this.AngleFactory.UP
            this.SteeringInsideGeographicMapCellPosition = nextUnvisitedPathGeographicMapCellPosition
        } else {
            this.MovementAngle = this.AngleFactory.DOWN
            this.SteeringInsideGeographicMapCellPosition = nextUnvisitedPathGeographicMapCellPosition
        }
        this.RtsLogHelper.Handle(this, this.MovementAngle)
}


// throws Exception
func (this *UnitLayer) TrackToDXYTargetAngle(dx int, dy int, targetAngle int) {
        var list IBasicArrayList = this.GetUnitWaypointBehavior().GetSteeringVisitorList(); _ = list
        if list.Size() > 0 {
            for index := int(list.Size() - 1); index >= 0; index-- {
                var steeringVisitor ISteeringVisitor = any(list.Get(index)).(ISteeringVisitor); _ = steeringVisitor
                var anyType any = steeringVisitor.Visit(this); _ = anyType
                if anyType == nil {
                    list.RemoveAt(index)
                }
            }
            this.FireOrMove()
        } else if !this.TurnTo(dx, dy, targetAngle) {
            this.FireOrMove()
        }
}


// throws Exception
func (this *UnitLayer) FireOrMove() {
        var gameKeyEventFactory *GameKeyEventFactory = GameKeyEventFactory_GetInstance(); _ = gameKeyEventFactory
        if this.GetUnitWaypointBehavior().NeedToMove() {
            this.RtsLayer2LogHelper.SteeringUp(this)
            if this.ShowMoreCaptionStates && !this.CaptionAnimationHelper.IsShowing() {
                this.CaptionAnimationHelper.Update(UnitLayer_MOVE, this.BasicColorFactory.GREEN)
            }
            this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_UP))
        } else {
            this.CaptionAnimationHelper.Update(CommonPhoneStrings_GetInstance().FIRE, this.BasicColorFactory.RED)
            this.RtsLayer2LogHelper.SteeringFireOrStop(this)
            this.AllStop()
            this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_KEY_NUM0))
            TrackingEventHandler_GetInstance().FireEvent(this.GetTrackingEvent())
        }
}


func (this *UnitLayer) Move() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "move", e)
                }
            }()
            var velocityProperties *VelocityProperties = this.GetVehicleProperties().GetVelocityProperties(); _ = velocityProperties
            var velocityXScaled int64 = int64(velocityProperties.GetVelocityXBasicDecimalP().GetScaled()); _ = velocityXScaled
            var velocityYScaled int64 = int64(velocityProperties.GetVelocityYBasicDecimalP().GetScaled()); _ = velocityYScaled
            this.GetUnitWaypointBehavior().Move()
            if velocityXScaled != 0 || velocityYScaled != 0 {
                this.GetUnitWaypointBehavior().SetMoving(true)
                var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
                var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
                this.LayerPartialCellPositionsUtil.GetAllDXY(geographicMapInterface, this, int(velocityXScaled), int(velocityYScaled), UnitLayer_GetPartialpositionlist())
                var cellPosition IGeographicMapCellPosition = any(DropCellPositionHistory_GetInstance().GetCellPositionWithDrop(UnitLayer_GetPartialpositionlist())).(IGeographicMapCellPosition); _ = cellPosition
                if cellPosition == this.CellPositionFactory.NONE {
                    var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
                    var tiledLayerUtil *TiledLayerUtil = TiledLayerUtil_GetInstance(); _ = tiledLayerUtil
                    var x int = this.X + int(velocityXScaled); _ = x
                    var y int = this.Y + int(velocityYScaled); _ = y
                    x = tiledLayerUtil.KeepOnMapX(tiledLayer, x, this.GetWidth())
                    y = tiledLayerUtil.KeepOnMapY(tiledLayer, y, this.GetHeight())
                    this.SetPosition(x, y, this.Z)
                } else {
                    var allbinaryLayer IAllBinaryLayer = DropCellPositionHistory_GetInstance().GetLayerInterface(cellPosition); _ = allbinaryLayer
                    this.GetUnitWaypointBehavior().SetMovingFromStopped(false)
                    this.GetUnitWaypointBehavior().AddBuildingChase(allbinaryLayer, cellPosition)
                }
            }
            if this.GetUnitWaypointBehavior().IsMoving() {
                TrackingEventHandler_GetInstance().FireEvent(this.GetTrackingEvent())
            } else {
                this.GetUnitWaypointBehavior().SetMovingFromStopped(false)
            }
            if this.GetUnitWaypointBehavior().IsMovingFromStopped() && this.IsVisible() {
                SecondaryPlayerQueueFactory_GetInstance().Add(this.MoveSoundInterface)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *UnitLayer) AllStop() {
        var velocityProperties *VelocityProperties = this.GetVehicleProperties().GetVelocityProperties(); _ = velocityProperties
        velocityProperties.GetVelocityXBasicDecimalP().Setint(0)
        velocityProperties.GetVelocityYBasicDecimalP().Setint(0)
}


func (this *UnitLayer) Paint(graphics IGraphics) {
        if this.IsVisible() {
            this.AdvancedRTSGameLayer.Paint(graphics)
            var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
            var viewX int = viewPosition.GetX(); _ = viewX
            var viewY int = viewPosition.GetY(); _ = viewY
            this.DecalAnimation.PaintXY(graphics, viewX, viewY)
            this.RangeAnimation.PaintXY(graphics, viewX, viewY)
            this.SensorRangeAnimation.PaintXY(graphics, viewX, viewY)
            this.DamageFloatersPaintableInterface.Paint(graphics)
            this.HealthBar.Paint(graphics)
            this.CaptionAnimationHelper.PaintXY(graphics, viewX, viewY)
            this.PathAnimation.PaintXY(graphics, viewX, viewY)
            this.ResourceAnimation.PaintXY(graphics, viewX, viewY)
        }
}


// throws Exception
func (this *UnitLayer) OnBuildingEvent(event *RTSLayerEvent) {
        var buildingLayer *BuildingLayer = any(event.GetRtsLayer()).(*BuildingLayer); _ = buildingLayer
        this.GetUnitWaypointBehavior().MoveAwayFromBuilding(buildingLayer)
}


func (this *UnitLayer) GetVehicleProperties() *VehicleProperties {
        return this.VehicleProperties
}


// throws Exception
func (this *UnitLayer) Damage(damage int, damageType int) {
        this.AdvancedRTSGameLayer.Damage(damage, damageType)
        this.DamageFloaters.Add(damage)
        if damage > 0 {
            this.GetHealthInterface().Damage(damage)
        }
}


// throws Exception
func (this *UnitLayer) GetDamage(damageType int) int {
        return 0
}


// throws Exception
func (this *UnitLayer) SetDestroyed(destroyed bool) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "setDestroyed")
        this.AdvancedRTSGameLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            WaypointEventHandlerFactory_GetInstance(this.GetGroupInterface()[0]).RemoveListener(this.GetUnitWaypointBehavior())
            TrackingEventHandler_GetInstance().RemoveListener(this)
            BuildingEventHandler_GetInstance().RemoveListener(this)
            if !this.GetHealthInterface().IsAlive() {
                var damage int = this.GetHealthInterface().GetMaxHealth(); _ = damage
                if damage > 10 {
                    SecondaryPlayerQueueFactory_GetInstance().Add(ExplosionBasicSound_GetInstance())
                    if damage < 100 {
                        this.ShakeListener.OnSmallShakeEvent()
                        this.Vibration.Vibrate(this.Duration, 0, 0)
                    } else if damage < 1000 {
                        this.ShakeListener.OnMediumShakeEvent()
                        this.Vibration.Vibrate(this.Duration * 2, 0, 0)
                    } else if damage < 3000 {
                        this.ShakeListener.OnLargeShakeEvent()
                        this.Vibration.Vibrate(this.Duration * 4, 0, 0)
                    }
                }
            }
        }
}


func (this *UnitLayer) GetLoad() int {
        return this.ResourceLoad
}


// throws Exception
func (this *UnitLayer) ClearResourceAnimation() {
        this.ResourceAnimation = any(NullIndexedAnimationFactory_GetFactoryInstance().GetInstance(0)).(IIndexedAnimation)
}


// throws Exception
func (this *UnitLayer) SetLoad(resource int) {
        if resource > 0 {
            this.ResourceAnimation = this.InitResourceAnimation
        } else {
            this.ClearResourceAnimation()
        }
        this.ResourceLoad = resource
}


func (this *UnitLayer) AddLoad(resource int) {
        this.ResourceLoad += resource
}


// throws Exception
func (this *UnitLayer) HandleCost(ownerLayer PathFindingLayerInterface) {
        if this.GetLoad() > 0 {
            this.CAPITAL_EVENT.SetValue(this.GetLoad())
            CapitalEventHandlerFactory_GetInstance(ownerLayer.GetGroupInterface()[0]).FireEvent(this.CAPITAL_EVENT)
            this.SetLoad(0)
        }
}


func (this *UnitLayer) CreateHudPaintable() ISelectionHudPaintable {
        var rtsLayerHudPaintable *RTSLayerHudPaintable = RTSLayerHudPaintable_GetInstance(); _ = rtsLayerHudPaintable
        rtsLayerHudPaintable.SetBasicColorP(this.AllBinaryGameLayerManagerP.GetForegroundBasicColor())
        rtsLayerHudPaintable.SetRtsLayer(this)
        return rtsLayerHudPaintable
}


func (this *UnitLayer) GetHudPaintable() ISelectionHudPaintable {
        var rtsLayerHudPaintable *RTSLayerHudPaintable = RTSLayerHudPaintable_GetInstance(); _ = rtsLayerHudPaintable
        return rtsLayerHudPaintable
}


func (this *UnitLayer) GetType() int {
        return UnitLayer_GetStaticType()
}


func UnitLayer_GetStaticType() int {
        return 1
}


func (this *UnitLayer) GetUnitWaypointBehavior() IUnitWaypointBehavior {
        return any(this.GetWaypointBehavior()).(IUnitWaypointBehavior)
}


func (this *UnitLayer) GetTrackingEvent() *TrackingEvent {
        return this.TrackingEvent
}


func (this *UnitLayer) GetCaptionAnimationHelper() ICaptionAnimationHelperBase {
        return this.CaptionAnimationHelper
}


func (this *UnitLayer) IsSelfUpgradeable() bool {
        return false
}


func (this *UnitLayer) GetMaxResourceLoad() int16 {
        return this.MaxResourceLoad
}


func UnitLayer_GetPartialpositionlist() IBasicArrayList {
        return UnitLayer_PartialPositionList
}

