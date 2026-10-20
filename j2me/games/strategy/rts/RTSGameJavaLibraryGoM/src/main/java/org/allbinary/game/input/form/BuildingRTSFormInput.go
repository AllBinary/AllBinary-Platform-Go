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

type BuildingRTSFormInput struct {
    RTSFormInput
    IsUnitProducer bool
    BuildOnPathGameNotificationEvent *GameNotificationEvent
    SpotTakenGameNotificationEvent *GameNotificationEvent
    SelectBuildSpotGameNotificationEvent *GameNotificationEvent
    RoadCollisionGameNotificationEvent *GameNotificationEvent
    NoMoneyGameNotificationEvent *GameNotificationEvent
    BuildingCollisionGameNotificationEvent *GameNotificationEvent
    StructureToCloseCollisionGameNotificationEvent *GameNotificationEvent
    BuildingGameNotificationEvent *GameNotificationEvent
    MapEdgeGameNotificationEvent *GameNotificationEvent
    DropCellPositionHistory *DropCellPositionHistory
}

// The Java class BuildingRTSFormInput has subclasses so its variables are this interface
type IBuildingRTSFormInput interface {
    IRTSFormInput
    BuildingRTSFormInput_() *BuildingRTSFormInput
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int)
    IsPositionBlocked() bool
    IsDemoLocked(item IABCustomItem) bool
    BuildFromMotionInput(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int)
    AttemptBuild(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer, itemIndex int) bool
    IsBuildAttemptValid(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerInterface IRTSLayer) bool
    Add(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer)
    IsSurroundingCellsOffMap(layerInterface IRTSLayer) bool
}

func (this *BuildingRTSFormInput) BuildingRTSFormInput_() *BuildingRTSFormInput {
    return this
}


func (this *BuildingRTSFormInput) initFields() {
    
    
    
    
    
    
    
    
    
    
    this.DropCellPositionHistory = DropCellPositionHistory_GetInstance()
}


func NewBuildingRTSFormInput(groupInterface []*Group, isUnitProducer bool) *BuildingRTSFormInput {
    this := &BuildingRTSFormInput{}
        this.RTSFormInput = *NewRTSFormInput(groupInterface)
        this.initFields()
        this.IsUnitProducer = isUnitProducer
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.BuildOnPathGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().BUILD_ON_PATH, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.SpotTakenGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().SPOT_TAKEN, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.SelectBuildSpotGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().SELECT_BUILD_SPOT, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.RoadCollisionGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().ROAD_COLLISION, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.NoMoneyGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NO_MONEY, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.BuildingCollisionGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().BUILDING_COLLISION, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.StructureToCloseCollisionGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().STRUCTURE_TO_CLOSE, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.BuildingGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().BUILDING, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.MapEdgeGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().MAP_EDGE, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
    return this
}


// throws Exception
func (this *BuildingRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.RTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.BuildOnPathGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.SpotTakenGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.SelectBuildSpotGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.RoadCollisionGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.NoMoneyGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.BuildingCollisionGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.StructureToCloseCollisionGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.BuildingGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.MapEdgeGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


// throws Exception
func (this *BuildingRTSFormInput) ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int) {
        this.RTSFormInput.Process(layerManager)
        this.BuildFromMotionInput(rtsPlayerLayerInterface, layerManager, item, itemIndex)
}


func (this *BuildingRTSFormInput) IsPositionBlocked() bool {
        return this.DropCellPositionHistory.IsCellPositionWithDrop(this.GetSelectedGeographicCellPosition())
}


func (this *BuildingRTSFormInput) IsDemoLocked(item IABCustomItem) bool {
        var factoryItem *LayerInterfaceFactoryImageItem = any(item).(*LayerInterfaceFactoryImageItem); _ = factoryItem
        var layerInterfaceFactoryInterface DemoLockedWithCostLayerInterfaceFactoryInterface = any(factoryItem.GetLayerInterfaceFactoryInterface()).(DemoLockedWithCostLayerInterfaceFactoryInterface); _ = layerInterfaceFactoryInterface
        if LockedUtil_GetInstance().IsLocked() && layerInterfaceFactoryInterface.IsDemoLocked() {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *BuildingRTSFormInput) BuildFromMotionInput(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(commonStrings.START, this, "buildFromMotionInput")
        if layerManager == nil {
            panic(errors.New(fmt.Sprint()))
        }
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(layerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        if this.IsDemoLocked(item) {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            LockedFeatureNotificationUtil_GetInstance().FireWithColor(geographicMapInterface.GetForegroundBasicColor())
            return
        }
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetSelectedGeographicCellPosition(); _ = geographicMapCellPosition
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            if !this.IsPositionBlocked() {
                var geographicMapCellType IGeographicMapCellType = geographicMapInterface.GetCellTypeAt(geographicMapCellPosition); _ = geographicMapCellType
                var raceTrackGeographicMapCellTypeFactory IRaceTrackGeographicMapCellTypeFactory = any(geographicMapInterface.GetGeographicMapCellTypeFactory()).(IRaceTrackGeographicMapCellTypeFactory); _ = raceTrackGeographicMapCellTypeFactory
                if !raceTrackGeographicMapCellTypeFactory.IsPath(geographicMapCellType) {
                    if this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] == CollidableDestroyableDamageableLayer_GetNullInstance() {
                        this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] = this.GetInstance(layerManager, item, geographicMapCellPosition)
                    } else {
                        var rtsLayer IRTSLayer = any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer); _ = rtsLayer
                        var cellPoint IGPoint = geographicMapCellPosition.GetPoint(); _ = cellPoint
                        rtsLayer.SetPosition(cellPoint.GetX(), cellPoint.GetY(), cellPoint.GetZ())
                        rtsLayer.RTSLayer_().GeographicMapCellPositionAreaBase.Update(geographicMapInterface)
                    }
                    var rtsLayer IRTSLayer = any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer); _ = rtsLayer
                    this.AttemptBuild(rtsPlayerLayerInterface, layerManager, rtsLayer, itemIndex)
                } else {
                    rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                    if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                        GameNotificationEventHandler_GetInstance().FireEvent(this.BuildOnPathGameNotificationEvent)
                    }
                }
            } else {
                rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.SpotTakenGameNotificationEvent)
                }
            }
        } else {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.SelectBuildSpotGameNotificationEvent)
            }
        }
}


// throws Exception
func (this *BuildingRTSFormInput) AttemptBuild(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer, itemIndex int) bool {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Layer: ").Append(StringUtil_GetInstance().ToString(layerInterface))), this, "attemptBuild")
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = layerInterface.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var occupyList IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        if !this.IsBuildAttemptValid(rtsPlayerLayerInterface, layerInterface) {
            return false
        }
        var geographicMapCellType IGeographicMapCellType; _ = geographicMapCellType
        for index := int(occupyList.Size() - 1); index >= 0; index-- {
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(layerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            geographicMapCellType = geographicMapInterface.GetCellTypeAt(any(occupyList.Get(index)).(IGeographicMapCellPosition))
            var raceTrackGeographicMapCellTypeFactory IRaceTrackGeographicMapCellTypeFactory = any(geographicMapInterface.GetGeographicMapCellTypeFactory()).(IRaceTrackGeographicMapCellTypeFactory); _ = raceTrackGeographicMapCellTypeFactory
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
            this.Add(rtsPlayerLayerInterface, layerManager, layerInterface)
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
func (this *BuildingRTSFormInput) IsBuildAttemptValid(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerInterface IRTSLayer) bool {
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = layerInterface.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var occupyList IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        if this.DropCellPositionHistory.AnyCellPositionWithDrop(occupyList) {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.BuildingCollisionGameNotificationEvent)
            }
            return false
        }
        if this.IsUnitProducer && this.DropCellPositionHistory.AnyCellPositionWithDrop(geographicMapCellPositionArea.GetSurroundingGeographicMapCellPositionList()) {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.StructureToCloseCollisionGameNotificationEvent)
            }
            return false
        } else {
            if this.IsSurroundingCellsOffMap(layerInterface) {
                rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.MapEdgeGameNotificationEvent)
                }
                return false
            }
        }
        return true
}


// throws Exception
func (this *BuildingRTSFormInput) Add(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer) {
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = layerInterface.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var occupyList IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        this.DropCellPositionHistory.AddAll(occupyList, layerInterface)
        var rtsPlayerGameInput IRTSPlayerGameInput = any(rtsPlayerLayerInterface.GetPlayerGameInput()).(IRTSPlayerGameInput); _ = rtsPlayerGameInput
        rtsPlayerGameInput.SetSelectedRTSLayer(layerInterface, this.GetSelectedGeographicCellPosition())
        layerManager.Append(layerInterface)
        rtsPlayerLayerInterface.Add(BuildingSound_GetInstance())
        if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
            GameNotificationEventHandler_GetInstance().FireEvent(this.BuildingGameNotificationEvent)
        }
}


func (this *BuildingRTSFormInput) IsSurroundingCellsOffMap(layerInterface IRTSLayer) bool {
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = layerInterface.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var occupyList IBasicArrayList = geographicMapCellPositionArea.GetOccupyingGeographicMapCellPositionList(); _ = occupyList
        var surroundList IBasicArrayList = geographicMapCellPositionArea.GetSurroundingGeographicMapCellPositionList(); _ = surroundList
        var occupySize int = occupyList.Size(); _ = occupySize
        var surroundSize int = surroundList.Size(); _ = surroundSize
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("occupySize: ")
        stringBuffer.Appendint(occupySize)
        stringBuffer.Append(" surroundSize: ")
        stringBuffer.Appendint(surroundSize)
        stringBuffer.Append(" surroundList: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(surroundList))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "isSurroundingCellsOffMap")
        var isSurroundOffMap bool = false; _ = isSurroundOffMap
        if occupySize == 1 && surroundSize != 8 {
            isSurroundOffMap = true
        } else if occupySize == 4 && surroundSize != 12 {
            isSurroundOffMap = true
        }
        if occupySize == 0 || isSurroundOffMap {
            return true
        } else {
            return false
        }
}

