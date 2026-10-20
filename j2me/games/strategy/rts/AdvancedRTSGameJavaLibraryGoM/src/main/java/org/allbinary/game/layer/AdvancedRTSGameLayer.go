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

type AdvancedRTSGameLayer struct {
    RTSLayer
    ShakeListener IShakeAnimationListener
    Vibration IAllBinaryVibrationME
    Duration int
    ParentLayer PathFindingLayerInterface
    WaypointBehaviorBase TickableInterface
}

// Implements DestroyedEventListenerInterface

// The Java class AdvancedRTSGameLayer has subclasses so its variables are this interface
type IAdvancedRTSGameLayer interface {
    IRTSLayer
    AdvancedRTSGameLayer_() *AdvancedRTSGameLayer
    UpdateWaypointBehavior(geographicMapInterface IBasicGeographicMap)
    InitVisibility(rtsPlayerLayerInterface RTSPlayerLayerInterface)
    Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface)
    SetDestroyed(destroyed bool)
    GetParentLayer() PathFindingLayerInterface
    SetParentLayer(parentLayer PathFindingLayerInterface)
    GetWaypointBehavior() IWaypointBehaviorBase
    SetWaypointBehavior(unitWaypointHelper IWaypointBehaviorBase)
    GetWaypointLogHelper() IWaypointLogHelper
    GetWaypoint2LogHelper() IWaypoint2LogHelper
    GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper
    ShouldAddWaypointFromBuilding() bool
    IsWaypointListEmptyOrOnlyTargets() bool
    BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool
    GetSurroundingGeographicMapCellPositionList() IBasicArrayList
    SetSelected(selected bool)
    OnDestroyed(destroyedEvent *DestroyedEvent)
}

func (this *AdvancedRTSGameLayer) AdvancedRTSGameLayer_() *AdvancedRTSGameLayer {
    return this
}


func (this *AdvancedRTSGameLayer) initFields() {
    
    
    
    this.ParentLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
    this.WaypointBehaviorBase = NullTickable_GetInstance()
}


// throws Exception
func AdvancedRTSGameLayer_CreateSimulated() IAdvancedRTSGameLayer {
        var nullAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullAnimationFactory_GetFactoryInstance(); _ = nullAnimationInterfaceFactoryInterface
        var nullIndexedAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface = NullIndexedAnimationFactory_GetFactoryInstance(); _ = nullIndexedAnimationInterfaceFactoryInterface
        return NewAdvancedRTSGameLayer(RemoteInfo_REMOTE_INFO, NullPathFindingLayer_NULL_PATH_FINDING_LAYER, AdvancedRTSProperties_CreatePropertiesSimulated(), GroupFactory_GetInstance().NULL_GROUP_ARRAY, StringUtil_GetInstance().EMPTY_STRING, StringUtil_GetInstance().EMPTY_STRING, Health_NULL_HEALTH, NullRTSFormInputFactory_GetInstance(), nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullAnimationInterfaceFactoryInterface, nullIndexedAnimationInterfaceFactoryInterface, NullIndexedAnimationFactory_GetFactoryInstance(), RectangleFactory_SINGLETON, 0, 0, NewTileLayerPositionIntoViewPosition())
}


func NewAdvancedRTSGameLayer(remoteInfo *RemoteInfo, parentLayer PathFindingLayerInterface, advancedRTSProperties *AdvancedRTSProperties, groupInterface []*Group, rootName string, name string, healthInterface *Health, rtsFormInput IRTSFormInput, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, x int, y int, viewPosition IViewPositionBase) *AdvancedRTSGameLayer {
    this := &AdvancedRTSGameLayer{}
        this.RTSLayer = *NewRTSLayer(remoteInfo, groupInterface, rootName, name, healthInterface, rtsFormInput, animationInterfaceFactoryInterface, emptyAnimationInterfaceFactoryInterface, baseAnimationInterfaceFactoryInterface, buildAnimationInterfaceFactoryInterface, verticleBuildAnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface, rectangle, x, y, viewPosition)
        this.initFields()
        this.SetVisible(false)
        this.SetWaypointBehavior(NewWaypointBehaviorBase())
        this.ShakeListener = advancedRTSProperties.ShakeListener
        this.Vibration = advancedRTSProperties.Vibration
        this.Duration = advancedRTSProperties.Duration
        this.SetParentLayer(parentLayer)
    return this
}


// throws Exception
func (this *AdvancedRTSGameLayer) UpdateWaypointBehavior(geographicMapInterface IBasicGeographicMap) {
        this.RTSLayer.UpdateWaypointBehavior(geographicMapInterface)
        this.GetWaypointBehavior().GetWaypoint().SetAllBinaryGameLayerManager(this.AllBinaryGameLayerManagerP)
}


func (this *AdvancedRTSGameLayer) InitVisibility(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        var advancedRTSPlayerLayerInterface AdvancedRTSPlayerLayerInterface = any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface); _ = advancedRTSPlayerLayerInterface
        if advancedRTSPlayerLayerInterface.IsLocalPlayer() || advancedRTSPlayerLayerInterface.GetGameType() == GameTypeFactory_GetInstance().BOT {
            this.SetVisible(true)
        } else {
            this.SetVisible(false)
        }
}


// throws Exception
func (this *AdvancedRTSGameLayer) Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        this.RTSLayer.Construct(rtsPlayerLayerInterface)
        this.InitVisibility(rtsPlayerLayerInterface)
        BuildingEventHandler_GetInstance().AddListenerInterface(this.GetWaypointBehavior().GetWaypoint())
}


// throws Exception
func (this *AdvancedRTSGameLayer) SetDestroyed(destroyed bool) {
        this.RTSLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            BuildingEventHandler_GetInstance().RemoveListener(this.GetWaypointBehavior().GetWaypoint())
            BuildingEventHandler_GetInstance().RemoveListener(this)
            DestroyedEventHandler_GetInstance().AddListener(this)
            RTSLayerUtil_GetInstance().DestroyAndClear(this.GetWaypointBehavior().GetOwnedWaypointList())
        }
}


func (this *AdvancedRTSGameLayer) GetParentLayer() PathFindingLayerInterface {
        return this.ParentLayer
}


func (this *AdvancedRTSGameLayer) SetParentLayer(parentLayer PathFindingLayerInterface) {
        this.ParentLayer = parentLayer
}


func (this *AdvancedRTSGameLayer) GetWaypointBehavior() IWaypointBehaviorBase {
        return any(this.WaypointBehaviorBase).(IWaypointBehaviorBase)
}


func (this *AdvancedRTSGameLayer) SetWaypointBehavior(unitWaypointHelper IWaypointBehaviorBase) {
        this.WaypointBehaviorBase = unitWaypointHelper
}


func (this *AdvancedRTSGameLayer) GetWaypointLogHelper() IWaypointLogHelper {
        return this.WaypointLogHelperP
}


func (this *AdvancedRTSGameLayer) GetWaypoint2LogHelper() IWaypoint2LogHelper {
        return this.Waypoint2LogHelperP
}


func (this *AdvancedRTSGameLayer) GetWaypointRunnableLogHelper() IWaypointRunnableLogHelper {
        return this.WaypointRunnableLogHelperP
}


func (this *AdvancedRTSGameLayer) ShouldAddWaypointFromBuilding() bool {
        if this.ParentLayer != NullPathFindingLayer_NULL_PATH_FINDING_LAYER {
            var parentAdvancedRTSGameLayer IAdvancedRTSGameLayer = any(this.ParentLayer).(IAdvancedRTSGameLayer); _ = parentAdvancedRTSGameLayer
            if parentAdvancedRTSGameLayer.GetType() != UnitLayer_GetStaticType() {
                return true
            }
        }
        return false
}


func (this *AdvancedRTSGameLayer) IsWaypointListEmptyOrOnlyTargets() bool {
        if this.GetType() != UnitLayer_GetStaticType() {
            return true
        }
        return false
}


// throws Exception
func (this *AdvancedRTSGameLayer) BuildingChase(allbinaryLayer IAllBinaryLayer, cellPosition IGeographicMapCellPosition) bool {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
        var buildingDirection *Direction = GeographicMapDirectionUtil_GetInstance().GetDirectionFromCellPositionToAdjacentCellPosition(geographicMapCellPosition, cellPosition); _ = buildingDirection
        var angleInfo *AngleInfo = this.GetRotationAnimationInterface().GetAngleInfoP(); _ = angleInfo
        var angle int = FrameUtil_GetInstance().AdjustAngleToFrameAngle(int(angleInfo.GetAngle()) + 90); _ = angle
        var gameKeyEventFactory *GameKeyEventFactory = GameKeyEventFactory_GetInstance(); _ = gameKeyEventFactory
        if buildingDirection == DirectionFactory_GetInstance().UP {
            this.RtsLayer2LogHelper.BuildingAbove(this)
            if angle > 180 && angle < 360 {
                this.RtsLayer2LogHelper.BuildingChaseLeft(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_LEFT))
                if angle <= 190 {
                    return true
                }
            } else {
                this.RtsLayer2LogHelper.BuildingChaseRight(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_RIGHT))
                if angle >= 170 {
                    return true
                }
            }
        } else if buildingDirection == DirectionFactory_GetInstance().DOWN {
            this.RtsLayer2LogHelper.BuildingDown(this)
            if angle > 180 && angle < 360 {
                this.RtsLayer2LogHelper.BuildingChaseLeft(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_RIGHT))
                if angle > 315 {
                    return true
                }
            } else {
                this.RtsLayer2LogHelper.BuildingChaseRight(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_LEFT))
                if angle < 45 {
                    return true
                }
            }
        } else if buildingDirection == DirectionFactory_GetInstance().LEFT {
            this.RtsLayer2LogHelper.BuildingLeft()
            if angle > 90 && angle < 270 {
                this.RtsLayer2LogHelper.BuildingChaseDown(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_LEFT))
                if angle < 135 {
                    return true
                }
            } else {
                this.RtsLayer2LogHelper.BuildingChaseUp(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_RIGHT))
                if angle > 45 {
                    return true
                }
            }
        } else if buildingDirection == DirectionFactory_GetInstance().RIGHT {
            this.RtsLayer2LogHelper.BuildingRight()
            if angle > 90 && angle < 270 {
                this.RtsLayer2LogHelper.BuildingChaseDown(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_RIGHT))
                if angle > 225 {
                    return true
                }
            } else {
                this.RtsLayer2LogHelper.BuildingChaseUp(this, angle)
                this.GetGameKeyEventList().Add(gameKeyEventFactory.GetInstanceForKey(this, Canvas_LEFT))
                if angle < 315 && angle > 180 {
                    return true
                }
            }
        } else {
            return true
        }
        return false
}


// throws Exception
func (this *AdvancedRTSGameLayer) GetSurroundingGeographicMapCellPositionList() IBasicArrayList {
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.GeographicMapCellPositionAreaBase.Update(geographicMapInterface)
        return this.GeographicMapCellPositionAreaBase.GetSurroundingGeographicMapCellPositionList()
}


func (this *AdvancedRTSGameLayer) SetSelected(selected bool) {
        this.RTSLayer.SetSelected(selected)
        if this.Debug {
            if selected {
                this.WaypointLogHelperP = WaypointSelectedLogHelper_GetInstance()
                this.Waypoint2LogHelperP = Waypoint2SelectedLogHelper_GetInstance()
                this.WaypointRunnableLogHelperP = WaypointRunnableSelectedLogHelper_GetInstance()
            } else {
                this.WaypointLogHelperP = WaypointLogHelper_GetInstance()
                this.Waypoint2LogHelperP = Waypoint2LogHelper_GetInstance()
                this.WaypointRunnableLogHelperP = WaypointRunnableLogHelper_GetInstance()
            }
        }
}


func (this *AdvancedRTSGameLayer) OnDestroyed(destroyedEvent *DestroyedEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onDestroyed", e)
                }
            }()
            var layerInterface IAllBinaryLayer = destroyedEvent.GetLayerInterface(); _ = layerInterface
            if layerInterface == this {
                DestroyedEventHandler_GetInstance().RemoveListener(this)
                var waypoint IWaypointBase = any(this.GetWaypointBehavior().GetWaypoint()).(IWaypointBase); _ = waypoint
                waypoint.Reset()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

