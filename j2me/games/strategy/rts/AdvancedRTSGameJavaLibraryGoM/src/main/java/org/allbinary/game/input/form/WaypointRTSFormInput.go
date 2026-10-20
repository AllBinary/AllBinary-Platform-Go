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

type WaypointRTSFormInput struct {
    RTSFormInput
    WAYPOINT_EVENT *RTSLayerEvent
    DragToSpotGameNotificationEvent *GameNotificationEvent
    SpotTakenGameNotificationEvent *GameNotificationEvent
    BuildingCollisionGameNotificationEvent *GameNotificationEvent
    RoadCollisionGameNotificationEvent *GameNotificationEvent
    NewWaypointGameNotificationEvent *GameNotificationEvent
    NoMoneyGameNotificationEvent *GameNotificationEvent
    IsPrimaryWaypointCreator bool
    StickyAssociatedRtsLayer PathFindingLayerInterface
}


func (this *WaypointRTSFormInput) initFields() {
    this.WAYPOINT_EVENT = NewRTSLayerEvent(NullRTSLayer_NULL_RTS_LAYER)
    
    
    
    
    
    
    
    this.StickyAssociatedRtsLayer = NullPathFindingLayer_NULL_PATH_FINDING_LAYER
}


func NewWaypointRTSFormInput(groupInterface []*Group, isPrimaryWaypointCreator bool) *WaypointRTSFormInput {
    this := &WaypointRTSFormInput{}
        this.RTSFormInput = *NewRTSFormInput(groupInterface)
        this.initFields()
        this.IsPrimaryWaypointCreator = isPrimaryWaypointCreator
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.DragToSpotGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().DRAG_TO_SPOT, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.SpotTakenGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().SPOT_TAKEN, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.BuildingCollisionGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().BUILDING_COLLISION, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.RoadCollisionGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().ROAD_COLLISION, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.NewWaypointGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NEW_WAYPOINT, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.NoMoneyGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NO_MONEY, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
    return this
}


// throws Exception
func (this *WaypointRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.RTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.DragToSpotGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.SpotTakenGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.BuildingCollisionGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.RoadCollisionGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.NewWaypointGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.NoMoneyGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


// throws Exception
func (this *WaypointRTSFormInput) ProcessAtPoint(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) {
        this.RTSFormInput.Process(layerManager)
        if this.IsStickyItemSelected() || associatedRtsLayer == CollidableDestroyableDamageableLayer_GetNullInstance() {
            if this.IsStickyItemSelected() && associatedRtsLayer == CollidableDestroyableDamageableLayer_GetNullInstance() {
                var geographicMapCompositeInterface GeographicMapCompositeInterface = any(layerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
                var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
                var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
                var x int = point.GetX() + allBinaryTiledLayer.GetXP(); _ = x
                var y int = point.GetY() + allBinaryTiledLayer.GetYP(); _ = y
                var geographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(x, y); _ = geographicMapCellPosition
                this.ProcessCellPosition(rtsPlayerLayerInterface, layerManager, geographicMapCellPosition)
            } else {
                if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.DragToSpotGameNotificationEvent)
                }
            }
        } else {
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.SpotTakenGameNotificationEvent)
            }
        }
}


// throws Exception
func (this *WaypointRTSFormInput) ProcessCellPosition(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, geographicMapCellPosition IGeographicMapCellPosition) {
        var itemIndex int = this.GetSelectedStickyItemIndex(); _ = itemIndex
        if this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] == CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] = this.GetInstance(layerManager, this.GetSelectedStickyItem(), geographicMapCellPosition)
        } else {
            var cellPoint IGPoint = geographicMapCellPosition.GetPoint(); _ = cellPoint
            var rtsLayer IRTSLayer = any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer); _ = rtsLayer
            rtsLayer.SetPosition(cellPoint.GetX(), cellPoint.GetY(), rtsLayer.GetZP())
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(rtsLayer.RTSLayer_().AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            rtsLayer.RTSLayer_().GeographicMapCellPositionAreaBase.Update(geographicMapInterface)
        }
        this.AttemptBuild(rtsPlayerLayerInterface, layerManager, any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer), itemIndex)
}


// throws Exception
func (this *WaypointRTSFormInput) ProcessStickyGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int) {
        this.LogUtil.PutF("Set Sticking Item: " + fmt.Sprint(item), this, "processSticky")
        this.SetSelectedStickyItem(item)
        this.SetSelectedStickyItemIndex(index)
        this.SetStickyItemSelected(true)
        this.StickyAssociatedRtsLayer = any(associatedRtsLayer).(IRTSLayer)
        this.GetHashtable().Put(Layer_ID, associatedRtsLayer)
}


// throws Exception
func (this *WaypointRTSFormInput) AttemptBuild(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer, itemIndex int) bool {
        this.LogUtil.PutF("Layer: " + fmt.Sprint(layerInterface), this, "attemptBuild")
        if layerInterface == nil {
            this.LogUtil.Put("Layer was null", this, "attemptBuild", errors.New(fmt.Sprint()))
            return false
        }
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = layerInterface.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var list IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = list
        if DropCellPositionHistory_GetInstance().AnyCellPositionWithDrop(list) || WaypointCellPositionHistory_GetInstance().AnyCellPositionWithDrop(list) {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.BuildingCollisionGameNotificationEvent)
            }
            return false
        }
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(layerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var raceTrackGeographicMapCellTypeFactory IRaceTrackGeographicMapCellTypeFactory = any(geographicMapInterface.GetGeographicMapCellTypeFactory()).(IRaceTrackGeographicMapCellTypeFactory); _ = raceTrackGeographicMapCellTypeFactory
        var geographicMapCellType IGeographicMapCellType; _ = geographicMapCellType
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            geographicMapCellType = geographicMapInterface.GetCellTypeAt(any(list.Get(index)).(IGeographicMapCellPosition))
            if raceTrackGeographicMapCellTypeFactory.IsPath(geographicMapCellType) {
                rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.RoadCollisionGameNotificationEvent)
                }
                return false
            }
        }
        var cost int = layerInterface.GetCost(); _ = cost
        var capital ICapital = rtsPlayerLayerInterface.GetCapital(); _ = capital
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Trying to Build: ")
        stringBuffer.Append(layerInterface.GetName())
        stringBuffer.Append(" for: $")
        stringBuffer.Appendint(cost)
        stringBuffer.Append(" with ")
        stringBuffer.Appendint(capital.GetTotalMoney())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "attemptBuild")
        if cost <= capital.GetTotalMoney() {
            layerInterface.Construct(rtsPlayerLayerInterface)
            this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] = CollidableDestroyableDamageableLayer_GetNullInstance()
            capital.RemoveMoney(cost)
            WaypointCellPositionHistory_GetInstance().AddList(list, layerInterface)
            var geographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(layerInterface.GetXP(), layerInterface.GetYP()); _ = geographicMapCellPosition
            var waypointLayer IWaypointLayer = any(layerInterface).(IWaypointLayer); _ = waypointLayer
            this.AddWayPoint(waypointLayer)
            var rtsPlayerGameInput IRTSPlayerGameInput = any(rtsPlayerLayerInterface.GetPlayerGameInput()).(IRTSPlayerGameInput); _ = rtsPlayerGameInput
            rtsPlayerGameInput.SetSelectedRTSLayer(layerInterface, geographicMapCellPosition)
            var advancedRTSPlayerLayerInterface AdvancedRTSPlayerLayerInterface = any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface); _ = advancedRTSPlayerLayerInterface
            advancedRTSPlayerLayerInterface.GetAdvancedPlayerOwnedRTSLayers().AddWaypoint(layerInterface)
            layerManager.Append(layerInterface)
            var waypointAvancedRTSGameLayer IAdvancedRTSGameLayer = any(waypointLayer).(IAdvancedRTSGameLayer); _ = waypointAvancedRTSGameLayer
            rtsPlayerLayerInterface.Add(waypointAvancedRTSGameLayer.GetWaypointBehavior().GetWaypoint().GetSound())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.NewWaypointGameNotificationEvent)
            }
            return true
        } else {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.NoMoneyGameNotificationEvent)
            }
            return false
        }
}


// throws Exception
func (this *WaypointRTSFormInput) AddWayPoint(layerInterface IWaypointLayer) {
        var stickyAssociatedAdvandedRtsLayer IAdvancedRTSGameLayer = any(this.StickyAssociatedRtsLayer).(IAdvancedRTSGameLayer); _ = stickyAssociatedAdvandedRtsLayer
        var list IBasicArrayList = stickyAssociatedAdvandedRtsLayer.GetWaypointBehavior().GetOwnedWaypointList(); _ = list
        RTSLayerUtil_GetInstance().DestroyAndClear(list)
        list.Add(layerInterface)
        if this.IsPrimaryWaypointCreator {
            PrimaryWaypointHelper_GetInstance().SetWaypointLayer(layerInterface)
        }
        this.WAYPOINT_EVENT.SetRtsLayer(layerInterface)
        WaypointEventHandlerFactory_GetInstance(layerInterface.GetGroupInterface()[0]).FireEvent(this.WAYPOINT_EVENT)
}

