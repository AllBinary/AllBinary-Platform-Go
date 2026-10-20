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

type BuildingLayer struct {
    AdvancedRTSGameLayer
    BuildingLevelCost int
    Productivity int
    Efficiency int
    EfficiencyPerLevel int
    TrackingEvent *TrackingEvent
    DamageFloaters IDamageFloaters
    DamageFloatersPaintableInterface IPaintable
    HealthBar IPaintable
    PathsHashtable *ABHashtable[any, any]
    Local bool
    TimeDelayHelper ITimeDelayHelper
}

// Implements RotationAnimationInterfaceCompositeInterface, TrackingEventListenerInterface


func (this *BuildingLayer) initFields() {
    
    
    
    
    
    
    
    
    
    
    this.TimeDelayHelper = NewTimeDelayHelper(3000)
}


// throws Exception
func BuildingLayer_CreateSimulated() *BuildingLayer {
        var nullAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullAnimationFactory_GetFactoryInstance(); _ = nullAnimationInterfaceFactoryInterface
        var nullIndexedAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullIndexedAnimationFactory_GetFactoryInstance(); _ = nullIndexedAnimationInterfaceFactoryInterface
        return NewBuildingLayer(RemoteInfo_REMOTE_INFO, SimulatedBuildingPropertiesFactory_GetInstance(), AdvancedRTSProperties_CreatePropertiesSimulated(), GroupFactory_GetInstance().NULL_GROUP_ARRAY, StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING, Health_NULL_HEALTH, NullRTSFormInputFactory_GetInstance(), nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, NullIndexedAnimationFactory_GetFactoryInstance(), RectangleFactory_SINGLETON, 0, 0)
}


func NewBuildingLayer(remoteInfo *RemoteInfo, buildingPropertiesFactory IBuildingPropertiesFactory, advancedRTSProperties *AdvancedRTSProperties, groupInterface []*Group, rootName string, name string, healthInterface *Health, rtsFormInput IRTSFormInput, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, x int, y int) *BuildingLayer {
    this := &BuildingLayer{}
        this.AdvancedRTSGameLayer = *NewAdvancedRTSGameLayer(remoteInfo, NullPathFindingLayer_NULL_PATH_FINDING_LAYER, advancedRTSProperties, groupInterface, rootName, name, healthInterface, rtsFormInput, animationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface, rectangle, x, y, NewTileLayerPositionIntoViewPosition())
        this.initFields()
        this.SetCollidableInferface(NewCollidableRTSBehavior(true))
        this.GetWaypointBehavior().SetWaypoint(NewWaypoint(this, SelectSound_GetInstance()))
        this.DamageFloaters = buildingPropertiesFactory.GetDamageFloaters(this)
        this.DamageFloatersPaintableInterface = buildingPropertiesFactory.BuildingPropertiesFactory_().DamageFloatersPaintableInterface
        this.HealthBar = buildingPropertiesFactory.GetHealthBar(this)
        this.PathsHashtable = buildingPropertiesFactory.GetHashtable()
        this.SetMaxLevel(30)
        this.SetProductivity(1)
        this.SetEfficiency(this.CalculateEfficiency())
        this.EfficiencyPerLevel = buildingPropertiesFactory.GetEfficiencyPerLevel(this)
        this.Efficiency = this.EfficiencyPerLevel
        this.GenerateMoveOutOfBuildAreaPaths()
        this.TrackingEvent = buildingPropertiesFactory.GetTrackingEvent(this)
    return this
}


func (this *BuildingLayer) InitVisibility(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        var advancedRTSPlayerLayerInterface AdvancedRTSPlayerLayerInterface = any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface); _ = advancedRTSPlayerLayerInterface
        if advancedRTSPlayerLayerInterface.IsLocalPlayer() {
            this.Local = true
            this.AddVisibility()
        } else {
            this.Local = false
        }
        this.AdvancedRTSGameLayer.InitVisibility(rtsPlayerLayerInterface)
}


// throws Exception
func (this *BuildingLayer) Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        this.AdvancedRTSGameLayer.Construct(rtsPlayerLayerInterface)
        TrackingEventHandler_GetInstance().AddListenerInterface(this)
}


func (this *BuildingLayer) OnMovement(trackingEvent *TrackingEvent) {
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
                layerInterface.OnMovementFound(this.TrackingEvent)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *BuildingLayer) ProcessBuiltTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.GetHealthInterface().IsDamaged() {
            this.SetAnimationInterface(this.DestroyAnimationInterface)
        }
        if !this.GetHealthInterface().IsAlive() {
            if this.IsReadyForExplosion() {
                var currentFrame int = this.DestroyAnimationInterface.GetFrame(); _ = currentFrame
                var size int = this.DestroyAnimationInterface.GetSize() - 1; _ = size
                if currentFrame == size && !this.TimeDelayHelper.IsTimeTNT() {
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
                this.TimeDelayHelper.SetStartTimeTNT()
                this.SetReadyForExplosion(true)
            }
        } else {
            this.AdvancedRTSGameLayer.ProcessBuiltTick(allBinaryLayerManager)
        }
        this.IndexedButShouldBeRotationAnimationInterface.NextFrame()
}


func (this *BuildingLayer) CalculateEfficiency() int {
        return this.EfficiencyPerLevel
}


func (this *BuildingLayer) GetCost() int {
        var total int64 = RTSLayerUtil_GetInstance().GetCostExponential(int64((this.GetLevel() * this.GetBuildingLevelCost()))); _ = total
        return int(total)
}


func (this *BuildingLayer) GetDowngradeCost() int {
        var downgradeCost int64 = RTSLayerUtil_GetInstance().GetCostExponential(int64(((this.GetLevel() - 1) * this.GetBuildingLevelCost()))); _ = downgradeCost
        this.LogUtil.PutF("Cost: " + fmt.Sprint(downgradeCost), this, "getDowngradeCost")
        return int(downgradeCost) * 9 / 10
}


func (this *BuildingLayer) GetUpgradeCost() int {
        var upgradeCost int64 = RTSLayerUtil_GetInstance().GetCostExponential(int64(((this.GetLevel() + 1) * this.GetBuildingLevelCost()))); _ = upgradeCost
        return int(upgradeCost)
}


func (this *BuildingLayer) Downgrade() {
        this.AdvancedRTSGameLayer.Downgrade()
        this.SetProductivity(this.GetProductivity() - 1)
        this.SetEfficiency(this.GetEfficiency() - this.CalculateEfficiency())
        this.GetHealthInterface().SetMaxHealth(this.GetHealthInterface().GetMaxHealth() - ((this.GetLevel() + 1) * 100))
}


func (this *BuildingLayer) Upgrade() {
        this.AdvancedRTSGameLayer.Upgrade()
        this.SetProductivity(this.GetProductivity() + 1)
        this.SetEfficiency(this.GetEfficiency() + this.CalculateEfficiency())
        this.GetHealthInterface().SetMaxHealth(this.GetHealthInterface().GetMaxHealth() + (this.GetLevel() * 100))
}


func (this *BuildingLayer) GetBuildingLevelCost() int {
        return this.BuildingLevelCost
}


func (this *BuildingLayer) SetBuildingLevelCost(buildingLevelCost int) {
        this.BuildingLevelCost = buildingLevelCost
}


func (this *BuildingLayer) GetProductivity() int {
        return this.Productivity
}


func (this *BuildingLayer) SetProductivity(productivity int) {
        this.Productivity = productivity
}


func (this *BuildingLayer) GetEfficiency() int {
        return this.Efficiency
}


func (this *BuildingLayer) SetEfficiency(efficiency int) {
        this.Efficiency = efficiency
}


// throws Exception
func (this *BuildingLayer) GenerateMoveOutOfBuildAreaPaths() {
        var occupyList IBasicArrayList = this.GeographicMapCellPositionAreaBase.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        var surroundList IBasicArrayList = this.GeographicMapCellPositionAreaBase.GetSurroundingGeographicMapCellPositionList(); _ = surroundList
        var NO_DIRECTION *Direction = DirectionFactory_GetInstance().NO_DIRECTION; _ = NO_DIRECTION
        var geographicMapDirectionUtil *GeographicMapDirectionUtil = GeographicMapDirectionUtil_GetInstance(); _ = geographicMapDirectionUtil
        var pathsList IBasicArrayList; _ = pathsList
        var occupyGeographicMapCellPosition IGeographicMapCellPosition; _ = occupyGeographicMapCellPosition
        var surroundGeographicMapCellPosition IGeographicMapCellPosition; _ = surroundGeographicMapCellPosition
        var list IBasicArrayList; _ = list
        for index2 := int(occupyList.Size() - 1); index2 >= 0; index2-- {
            pathsList = NewBasicArrayListD()
            occupyGeographicMapCellPosition = any(occupyList.Get(index2)).(IGeographicMapCellPosition)
            for index := int(surroundList.Size() - 1); index >= 0; index-- {
                surroundGeographicMapCellPosition = any(surroundList.Get(index)).(IGeographicMapCellPosition)
                if geographicMapDirectionUtil.GetEightDirectionFromCellPositionToAdjacentCellPosition(surroundGeographicMapCellPosition, occupyGeographicMapCellPosition) != NO_DIRECTION {
                    list = NewBasicArrayListS(1)
                    list.Add(surroundGeographicMapCellPosition)
                    pathsList.Add(list)
                }
            }
            this.PathsHashtable.Put(occupyGeographicMapCellPosition, pathsList)
        }
}


func (this *BuildingLayer) GetMoveOutOfBuildAreaPath(geographicMapCellPosition IGeographicMapCellPosition) IBasicArrayList {
        var pathsList IBasicArrayList = any(this.PathsHashtable.Get(geographicMapCellPosition)).(IBasicArrayList); _ = pathsList
        return pathsList
}


func (this *BuildingLayer) GetEndGeographicMapCellPositionList() IBasicArrayList {
        return this.GeographicMapCellPositionAreaBase.GetSurroundingGeographicMapCellPositionList()
}


func (this *BuildingLayer) ShouldHandleStartSameAsEnd() bool {
        return false
}


func (this *BuildingLayer) Paint(graphics IGraphics) {
        if this.IsVisible() {
            this.AdvancedRTSGameLayer.Paint(graphics)
            this.DamageFloatersPaintableInterface.Paint(graphics)
            this.HealthBar.Paint(graphics)
        }
}


// throws Exception
func (this *BuildingLayer) Damage(damage int, damageType int) {
        this.AdvancedRTSGameLayer.Damage(damage, damageType)
        this.DamageFloaters.Add(damage)
        if damage > 0 {
            this.GetHealthInterface().Damage(damage)
        }
}


// throws Exception
func (this *BuildingLayer) GetDamage(damageType int) int {
        return 0
}


func (this *BuildingLayer) AddVisibility() {
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = this.GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var occupyList IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        VisibleCellPositionsSingleton_GetInstance().AddStationaryCellPositions(occupyList)
        var surroundList IBasicArrayList = geographicMapCellPositionArea.GetSurroundingGeographicMapCellPositionList(); _ = surroundList
        VisibleCellPositionsSingleton_GetInstance().AddStationaryCellPositions(surroundList)
}


func (this *BuildingLayer) RemoveVisibility() {
        var occupyList IBasicArrayList = this.GeographicMapCellPositionAreaBase.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        VisibleCellPositionsSingleton_GetInstance().RemoveStationaryCellPositions(occupyList)
        var surroundList IBasicArrayList = this.GeographicMapCellPositionAreaBase.GetSurroundingGeographicMapCellPositionList(); _ = surroundList
        VisibleCellPositionsSingleton_GetInstance().RemoveStationaryCellPositions(surroundList)
}


// throws Exception
func (this *BuildingLayer) SetDestroyed(destroyed bool) {
        this.AdvancedRTSGameLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            DropCellPositionHistory_GetInstance().RemoveAll(this)
            TrackingEventHandler_GetInstance().RemoveListener(this)
            if this.Local {
                this.RemoveVisibility()
            }
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


func (this *BuildingLayer) CreateHudPaintable() ISelectionHudPaintable {
        var buildingInfoHudPaintable *BuildingInfoHudPaintable = BuildingInfoHudPaintable_GetInstance(); _ = buildingInfoHudPaintable
        buildingInfoHudPaintable.SetBasicColorP(this.AllBinaryGameLayerManagerP.GetForegroundBasicColor())
        buildingInfoHudPaintable.SetRtsLayer(this)
        return buildingInfoHudPaintable
}


func (this *BuildingLayer) GetHudPaintable() ISelectionHudPaintable {
        var buildingInfoHudPaintable *BuildingInfoHudPaintable = BuildingInfoHudPaintable_GetInstance(); _ = buildingInfoHudPaintable
        return buildingInfoHudPaintable
}


func (this *BuildingLayer) GetType() int {
        return BuildingLayer_GetStaticType()
}


func BuildingLayer_GetStaticType() int {
        return 2
}

