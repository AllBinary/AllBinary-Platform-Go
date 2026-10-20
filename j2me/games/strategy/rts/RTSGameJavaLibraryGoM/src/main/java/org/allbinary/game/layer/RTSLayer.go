/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSLayer struct {
    MultiPlayerGameLayer
    Debug bool
    ShowMoreCaptionStates bool
    InputProcessorArray []IGameInputProcessor
    RtsLayer2LogHelper IRTSLayer2LogHelper
    WaypointLogHelperP IWaypointLogHelper
    Waypoint2LogHelperP IWaypoint2LogHelper
    WaypointRunnableLogHelperP IWaypointRunnableLogHelper
    RtsFormInput IRTSFormInput
    VerticleBuildAnimationInterface IIndexedAnimation
    RangeAnimation IAnimation
    InitRangeAnimation IAnimation
    SensorRangeAnimation IAnimation
    InitSensorRangeAnimation IAnimation
    BaseAnimationInterface IAnimation
    BuildFrameTimeHelper ITimeDelayHelper
    BuildAnimationInterface IAnimation
    AnimationInterface IAnimation
    InitAnimationInterface IIndexedAnimation
    IndexedButShouldBeRotationAnimationInterface IIndexedAnimation
    EmptyAnimationInterface IIndexedAnimation
    DestroyAnimationInterface IIndexedAnimation
    FireTimeHelper ITimeDelayHelper
    RootName string
    Hashtable *ABHashtable[any, any]
    BUILD_VALUE int
    GeographicMapCellPositionAreaBase IGeographicMapCellPositionAreaBase
    HealthInterface *Health
    Level int
    MaxLevel int
    HackVerticleBuild int
    SlightAngle int
    PercentCompleteP int
    Destroyed bool
    Selected bool
    RtsLayerUtil *RTSLayerUtil
    LayerDistanceUtil *LayerDistanceUtil
}

// Implements TickableInterface, GameInputInterface, GameKeyEventSourceInterface, RTSInterface, PathFindingLayerInterface

// The Java class RTSLayer has subclasses so its variables are this interface
type IRTSLayer interface {
    IMultiPlayerGameLayer
    RTSLayer_() *RTSLayer
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    UpdateWaypointBehavior(geographicMapInterface IBasicGeographicMap)
    Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface)
    Select()
    Deselect()
    SetSelected(selected bool)
    IsSelected() bool
    GetPercentComplete() int
    GetScrollSelectionForm() IScrollSelectionForm
    InitInputProcessors()
    ProcessBuiltTick(allBinaryLayerManager IAllBinaryLayerManager)
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    ProcessInput(layerManager IAllBinaryLayerManager)
    Reload()
    OnMovement(trackingEvent *TrackingEvent)
    OnMovementFound(trackingEvent *TrackingEvent)
    Paint(graphics IGraphics)
    Damage(damage int, damageType int)
    GetSourceId() int
    Downgrade()
    Upgrade()
    IsCompleted() bool
    Build()
    IsSelfUpgradeable() bool
    GetCost() int
    GetDowngradeCost() int
    GetUpgradeCost() int
    IsUpgradeable() bool
    IsDowngradeable() bool
    GetLevel() int
    SetLevel(level int)
    GetVerticleBuildAnimationInterface() IAnimation
    GetAnimationInterface() IAnimation
    SetAnimationInterface(animationInterface IAnimation)
    SetTarget(targetGameLayer PathFindingLayerInterface)
    CreateHudPaintable() ISelectionHudPaintable
    GetHudPaintable() ISelectionHudPaintable
    GetMaxLevel() int
    SetMaxLevel(maxLevel int)
    GetHealthInterface() *Health
    SetHealthInterface(healthInterface *Health)
    GetRTSFormInput() IRTSFormInput
    IsDestroyed() bool
    SetDestroyed(destroyed bool)
    GetEndGeographicMapCellPositionList() IBasicArrayList
    GetGeographicMapCellPositionArea() IGeographicMapCellPositionAreaBase
    ShouldHandleStartSameAsEnd() bool
    HandleCost(ownerLayer PathFindingLayerInterface)
    GetWaypointBehavior() IWaypointBehaviorBase
    GetParentLayer() PathFindingLayerInterface
    GetRTSLayer2LogHelper() IRTSLayer2LogHelper
    GetWaypointLogHelper() IWaypointLogHelper
    GetWaypoint2LogHelper() IWaypoint2LogHelper
    GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper
    ShouldAddWaypointFromBuilding() bool
    GetCaptionAnimationHelper() ICaptionAnimationHelperBase
    IsShowMoreCaptionStates() bool
    Init(geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPositionBasicArrayList IBasicArrayList)
    GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition
    GetTopLeftGeographicMapCellPosition() IGeographicMapCellPosition
    GetMoveOutOfBuildAreaPath(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList
    SetClosestGeographicMapCellHistory(pathsList IBasicArrayList)
    TeleportTo(geographicMapCellPosition IGeographicMapCellPosition)
    SetLoad(resource int)
    GetSurroundingGeographicMapCellPositionList() IBasicArrayList
    TrackTo(reason string)
    TrackToDXY(dx int, dy int)
    IsWaypointListEmptyOrOnlyTargets() bool
    GetTrackingEvent() *TrackingEvent
    BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool
    AllStop()
    ImplmentsTickableInterface() bool
    ImplmentsGameInputInterface() bool
    GetRootName() string
    GetRotationAnimationInterface() IRotationAnimation
    SetFrame(index int)
    SetRotationAnimationInterface(rotationAnimationInterface IRotationAnimation)
}

func (this *RTSLayer) RTSLayer_() *RTSLayer {
    return this
}

var RTSLayer_Id int = 0
var RTSLayer_MAIN_INDEX int = 0
var RTSLayer_SECONDARY_INDEX int = 1

func (this *RTSLayer) initFields() {
    this.Debug = true
    this.ShowMoreCaptionStates = this.Debug
    this.InputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    this.RtsLayer2LogHelper = RTSLayer2LogHelper_GetInstance()
    this.WaypointLogHelperP = WaypointLogHelper_GetInstance()
    this.Waypoint2LogHelperP = Waypoint2LogHelper_GetInstance()
    this.WaypointRunnableLogHelperP = WaypointRunnableLogHelper_GetInstance()
    
    
    
    
    
    
    
    this.BuildFrameTimeHelper = NewTimeDelayHelper(50)
    
    
    
    
    
    
    this.FireTimeHelper = NewTimeDelayHelper(700)
    
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    this.BUILD_VALUE = 63
    
    this.HealthInterface = Health_NULL_HEALTH
    this.Level = 1
    this.MaxLevel = 1
    this.HackVerticleBuild = this.BUILD_VALUE
    this.SlightAngle = 0
    
    this.Destroyed = false
    this.Selected = false
    this.RtsLayerUtil = RTSLayerUtil_GetInstance()
    this.LayerDistanceUtil = LayerDistanceUtil_GetInstance()
}


// throws Exception
func RTSLayer_CreateSimulatedInstance(remoteInfo *RemoteInfo) IRTSLayer {
        var nullAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullAnimationFactory_GetFactoryInstance(); _ = nullAnimationInterfaceFactoryInterface
        var nullIndexedAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullIndexedAnimationFactory_GetFactoryInstance(); _ = nullIndexedAnimationInterfaceFactoryInterface
        return NewRTSLayer(RemoteInfo_REMOTE_INFO, GroupFactory_GetInstance().NULL_GROUP_ARRAY, StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING, Health_NULL_HEALTH, NullRTSFormInputFactory_GetInstance(), nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, NullIndexedAnimationFactory_GetFactoryInstance(), RectangleFactory_SINGLETON, 0, 0, NewTileLayerPositionIntoViewPosition())
}


func NewRTSLayer(remoteInfo *RemoteInfo, groupInterface []*Group, rootName string, name string, healthInterface *Health, rtsFormInput IRTSFormInput, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, x int, y int, viewPosition IViewPositionBase) *RTSLayer {
    this := &RTSLayer{}
        this.MultiPlayerGameLayer = *NewMultiPlayerGameLayer(remoteInfo, groupInterface, name, rectangle, viewPosition)
        this.initFields()
        this.InitInputProcessors()
        this.RootName = rootName
        this.HealthInterface = healthInterface
        this.RtsFormInput = rtsFormInput
        this.SetPosition(x, y, this.Z)
        this.VerticleBuildAnimationInterface = any(verticleBuildAnimationInterfaceFactoryInterface.GetInstance(0)).(IIndexedAnimation)
        this.BuildAnimationInterface = buildAnimationInterfaceFactoryInterface.GetInstance(0)
        this.BaseAnimationInterface = baseAnimationInterfaceFactoryInterface.GetInstance(0)
        this.EmptyAnimationInterface = any(emptyAnimationInterfaceFactoryInterface.GetInstance(0)).(IIndexedAnimation)
        this.IndexedButShouldBeRotationAnimationInterface = any(animationInterfaceFactoryInterface.GetInstance(0)).(IIndexedAnimation)
        this.InitAnimationInterface = this.IndexedButShouldBeRotationAnimationInterface
        this.DestroyAnimationInterface = any(proceduralAnimationInterfaceFactoryInterface.GetInstanceAnimation(this.IndexedButShouldBeRotationAnimationInterface)).(IIndexedAnimation)
        var animation IAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0); _ = animation
        this.RangeAnimation = animation
        this.InitRangeAnimation = animation
        this.SensorRangeAnimation = animation
        this.InitSensorRangeAnimation = animation
        this.AnimationInterface = this.InitAnimationInterface
        this.GeographicMapCellPositionAreaBase = NewGeographicMapCellPositionArea(this)
    return this
}


// throws Exception
func (this *RTSLayer) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.MultiPlayerGameLayer.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        var viewPosition2 *TileLayerPositionIntoViewPosition = any(this.GetViewPosition()).(*TileLayerPositionIntoViewPosition); _ = viewPosition2
        viewPosition2.SetTiledLayer(tiledLayer)
        this.UpdateWaypointBehavior(geographicMapInterface)
}


// throws Exception
func (this *RTSLayer) UpdateWaypointBehavior(geographicMapInterface IBasicGeographicMap) {
        this.GeographicMapCellPositionAreaBase.Update(geographicMapInterface)
}


// throws Exception
func (this *RTSLayer) Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        ViewPositionEventHandler_GetInstance().AddListener(any(this).(IAllBinaryLayer))
        this.AnimationInterface = this.BuildAnimationInterface
        this.HackVerticleBuild = 0
}


func (this *RTSLayer) Select() {
        this.RangeAnimation = this.InitRangeAnimation
        this.SensorRangeAnimation = this.InitSensorRangeAnimation
        this.SetSelected(true)
}


func (this *RTSLayer) Deselect() {
        var animation IAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0); _ = animation
        this.RangeAnimation = animation
        this.SensorRangeAnimation = animation
        this.SetSelected(false)
}


func (this *RTSLayer) SetSelected(selected bool) {
        this.Selected = selected
}


func (this *RTSLayer) IsSelected() bool {
        return this.Selected
}


func (this *RTSLayer) GetPercentComplete() int {
        return this.PercentCompleteP
}


func (this *RTSLayer) GetScrollSelectionForm() IScrollSelectionForm {
        return ScrollSelectionForm_GetNullScrollSelectionForm()
}


func (this *RTSLayer) InitInputProcessors() {
        GameInputProcessorUtil_Init(this.InputProcessorArray)
}


// throws Exception
func (this *RTSLayer) ProcessBuiltTick(allBinaryLayerManager IAllBinaryLayerManager) {
        this.AnimationInterface = this.IndexedButShouldBeRotationAnimationInterface
}


func (this *RTSLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processTick", e)
                }
            }()
            if this.IsCompleted() {
                this.ProcessBuiltTick(allBinaryLayerManager)
            } else {
                this.Build()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *RTSLayer) ProcessInput(layerManager IAllBinaryLayerManager) {
}


func (this *RTSLayer) Reload() {
}


func (this *RTSLayer) OnMovement(trackingEvent *TrackingEvent) {
}


// throws Exception
func (this *RTSLayer) OnMovementFound(trackingEvent *TrackingEvent) {
}


func (this *RTSLayer) Paint(graphics IGraphics) {
        this.MultiPlayerGameLayer.PaintFirst(graphics)
        var viewPosition IViewPositionBase = this.GetViewPosition(); _ = viewPosition
        var viewX int = viewPosition.GetX(); _ = viewX
        var viewY int = viewPosition.GetY(); _ = viewY
        this.GetAnimationInterface().PaintXY(graphics, viewX, viewY)
}


// throws Exception
func (this *RTSLayer) Damage(damage int, damageType int) {
}


func (this *RTSLayer) GetSourceId() int {
        return RTSLayer_Id
}


func (this *RTSLayer) Downgrade() {
        this.RtsLayerUtil.Downgrade(this)
}


func (this *RTSLayer) Upgrade() {
        this.RtsLayerUtil.Upgrade(this)
}


func (this *RTSLayer) IsCompleted() bool {
        if this.HackVerticleBuild < this.BUILD_VALUE {
            return false
        } else {
            return true
        }
}


// throws Exception
func (this *RTSLayer) Build() {
        if this.BuildFrameTimeHelper.IsTimeTNT() {
            this.AnimationInterface = this.BuildAnimationInterface
            this.VerticleBuildAnimationInterface.NextFrame()
            this.HackVerticleBuild++
        }
        this.PercentCompleteP = 100 * this.HackVerticleBuild / this.BUILD_VALUE
        this.GetHudPaintable().UpdateInfo()
}


func (this *RTSLayer) IsSelfUpgradeable() bool {
        return true
}


func (this *RTSLayer) GetCost() int {
        return this.RtsLayerUtil.GetCost(this)
}


func (this *RTSLayer) GetDowngradeCost() int {
        return this.RtsLayerUtil.GetDowngradeCost(this)
}


func (this *RTSLayer) GetUpgradeCost() int {
        return this.RtsLayerUtil.GetUpgradeCost(this)
}


func (this *RTSLayer) IsUpgradeable() bool {
        if this.GetLevel() < this.GetMaxLevel() {
            return true
        } else {
            return false
        }
}


func (this *RTSLayer) IsDowngradeable() bool {
        if this.GetLevel() > 1 {
            return true
        } else {
            return false
        }
}


func (this *RTSLayer) GetLevel() int {
        return this.Level
}


func (this *RTSLayer) SetLevel(level int) {
        this.Level = level
}


func (this *RTSLayer) GetVerticleBuildAnimationInterface() IAnimation {
        return this.VerticleBuildAnimationInterface
}


func (this *RTSLayer) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *RTSLayer) SetAnimationInterface(animationInterface IAnimation) {
        this.AnimationInterface = animationInterface
}


// throws Exception
func (this *RTSLayer) SetTarget(targetGameLayer PathFindingLayerInterface) {
        var anotherTargetDistance int = this.LayerDistanceUtil.GetDistance(any(this).(IAllBinaryLayer), any(targetGameLayer).(IAllBinaryLayer)); _ = anotherTargetDistance
        var waypointBehaviorBase IWaypointBehaviorBase = this.GetWaypointBehavior(); _ = waypointBehaviorBase
        waypointBehaviorBase.SetTargetWithDistance(any(targetGameLayer).(PathFindingLayerInterface), anotherTargetDistance)
}


func (this *RTSLayer) CreateHudPaintable() ISelectionHudPaintable {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) GetHudPaintable() ISelectionHudPaintable {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) GetMaxLevel() int {
        return this.MaxLevel
}


func (this *RTSLayer) SetMaxLevel(maxLevel int) {
        this.MaxLevel = maxLevel
}


func (this *RTSLayer) GetHealthInterface() *Health {
        return this.HealthInterface
}


func (this *RTSLayer) SetHealthInterface(healthInterface *Health) {
        this.HealthInterface = healthInterface
}


func (this *RTSLayer) GetRTSFormInput() IRTSFormInput {
        return this.RtsFormInput
}


// throws Exception
func (this *RTSLayer) IsDestroyed() bool {
        return this.Destroyed
}


// throws Exception
func (this *RTSLayer) SetDestroyed(destroyed bool) {
        this.Destroyed = destroyed
        if this.IsDestroyed() {
            DestroyedLayerProcessor_GetInstance().Add(this)
            ViewPositionEventHandler_GetInstance().RemoveListener(this)
        }
}


func (this *RTSLayer) GetEndGeographicMapCellPositionList() IBasicArrayList {
        return this.GeographicMapCellPositionAreaBase.GetOccupyingGeographicMapCellPositionList()
}


func (this *RTSLayer) GetGeographicMapCellPositionArea() IGeographicMapCellPositionAreaBase {
        return this.GeographicMapCellPositionAreaBase
}


func (this *RTSLayer) ShouldHandleStartSameAsEnd() bool {
        return true
}


// throws Exception
func (this *RTSLayer) HandleCost(ownerLayer PathFindingLayerInterface) {
}


func (this *RTSLayer) GetWaypointBehavior() IWaypointBehaviorBase {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) GetParentLayer() PathFindingLayerInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) GetRTSLayer2LogHelper() IRTSLayer2LogHelper {
        return this.RtsLayer2LogHelper
}


func (this *RTSLayer) GetWaypointLogHelper() IWaypointLogHelper {
        return this.WaypointLogHelperP
}


func (this *RTSLayer) GetWaypoint2LogHelper() IWaypoint2LogHelper {
        return this.Waypoint2LogHelperP
}


func (this *RTSLayer) GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper {
        return this.WaypointRunnableLogHelperP
}


func (this *RTSLayer) ShouldAddWaypointFromBuilding() bool {
        return false
}


func (this *RTSLayer) GetCaptionAnimationHelper() ICaptionAnimationHelperBase {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) IsShowMoreCaptionStates() bool {
        return this.ShowMoreCaptionStates
}


// throws Exception
func (this *RTSLayer) Init(geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPositionBasicArrayList IBasicArrayList) {
}


// throws Exception
func (this *RTSLayer) GetCurrentGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *RTSLayer) GetTopLeftGeographicMapCellPosition() IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *RTSLayer) GetMoveOutOfBuildAreaPath(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *RTSLayer) SetClosestGeographicMapCellHistory(pathsList IBasicArrayList) {
}


func (this *RTSLayer) TeleportTo(geographicMapCellPosition IGeographicMapCellPosition) {
}


// throws Exception
func (this *RTSLayer) SetLoad(resource int) {
}


// throws Exception
func (this *RTSLayer) GetSurroundingGeographicMapCellPositionList() IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *RTSLayer) TrackTo(reason string) {
}


// throws Exception
func (this *RTSLayer) TrackToDXY(dx int, dy int) {
}


func (this *RTSLayer) IsWaypointListEmptyOrOnlyTargets() bool {
        return false
}


func (this *RTSLayer) GetTrackingEvent() *TrackingEvent {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *RTSLayer) BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool {
        return false
}


func (this *RTSLayer) AllStop() {
}


func RTSLayer_GetMinStaticType() int {
        return 1
}


func RTSLayer_GetMaxStaticType() int {
        return 4
}


func RTSLayer_IsRTSLayer(layerInterface IAllBinaryLayer) bool {
        if layerInterface.GetType() >= RTSLayer_GetMinStaticType() && layerInterface.GetType() <= RTSLayer_GetMaxStaticType() {
            return true
        }
        return false
}


func (this *RTSLayer) ImplmentsTickableInterface() bool {
        return true
}


func (this *RTSLayer) ImplmentsGameInputInterface() bool {
        return true
}


func (this *RTSLayer) GetRootName() string {
        return this.RootName
}


func (this *RTSLayer) GetRotationAnimationInterface() IRotationAnimation {
        return any(this.IndexedButShouldBeRotationAnimationInterface).(IRotationAnimation)
}


func (this *RTSLayer) SetFrame(index int) {
        this.GetRotationAnimationInterface().SetFrame(index)
}


func (this *RTSLayer) SetRotationAnimationInterface(rotationAnimationInterface IRotationAnimation) {
        this.IndexedButShouldBeRotationAnimationInterface = rotationAnimationInterface
}

