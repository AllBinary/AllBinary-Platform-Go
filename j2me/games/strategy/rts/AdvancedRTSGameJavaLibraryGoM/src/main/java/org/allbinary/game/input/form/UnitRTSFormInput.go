/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
 *  
 * By agreeing to this license you and any business entity you represent are 
 * legally bound to the AllBinary Open License Version 1 legal agreement. 
 *  
 * You may obtain the AllBinary Open License Version 1 legal agreement from 
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository. 
 *  
 * Created By: Travis Berthelot  
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

type UnitRTSFormInput struct {
    RTSFormInput
    NoMoneyGameNotificationEvent *GameNotificationEvent
    NewUnitGameNotificationEvent *GameNotificationEvent
}

var UnitRTSFormInput_DECAL_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(23)

func (this *UnitRTSFormInput) initFields() {
    
    
}


func NewUnitRTSFormInput(groupInterface []*Group) *UnitRTSFormInput {
    this := &UnitRTSFormInput{}
        this.RTSFormInput = *NewRTSFormInput(groupInterface)
        this.initFields()
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.NoMoneyGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NO_MONEY, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.NewUnitGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NEW_UNIT, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
        this.GetHashtable().Put(WorkWaypoint_ID, smallIntegerSingletonFactory.GetAt(50))
    return this
}


// throws Exception
func (this *UnitRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.RTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.NoMoneyGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.NewUnitGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


// throws Exception
func (this *UnitRTSFormInput) ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int) {
        this.RTSFormInput.Process(layerManager)
        var associatedRtsLayer2 IRTSLayer = any(associatedRtsLayer).(IRTSLayer); _ = associatedRtsLayer2
        var geographicMapCellPositionArea IGeographicMapCellPositionAreaBase = associatedRtsLayer2.RTSLayer_().GeographicMapCellPositionAreaBase; _ = geographicMapCellPositionArea
        var geographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionArea.GetNextSurroundingGeographicMapCellPosition(); _ = geographicMapCellPosition
        this.GetHashtable().Put(Layer_ID, associatedRtsLayer)
        this.GetHashtable().Put(UnitRTSFormInput_DECAL_ID, (any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface)).GetDecalBasicColor())
        if this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] == CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.NewUnconstructedRTSLayerInterfaceArray[itemIndex] = this.GetInstance(layerManager, item, geographicMapCellPosition)
        } else {
            var cellPoint IGPoint = geographicMapCellPosition.GetPoint(); _ = cellPoint
            var rtsLayer IRTSLayer = any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer); _ = rtsLayer
            rtsLayer.SetPosition(cellPoint.GetX(), cellPoint.GetY(), rtsLayer.GetZP())
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(rtsLayer.RTSLayer_().AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            rtsLayer.RTSLayer_().GeographicMapCellPositionAreaBase.Update(geographicMapInterface)
        }
        var cellPoint IGPoint = geographicMapCellPosition.GetMidPoint(); _ = cellPoint
        var rtsLayer IRTSLayer = any(this.NewUnconstructedRTSLayerInterfaceArray[itemIndex]).(IRTSLayer); _ = rtsLayer
        rtsLayer.SetPosition(cellPoint.GetX() - rtsLayer.GetHalfWidth(), cellPoint.GetY() - rtsLayer.GetHalfHeight(), rtsLayer.GetZP())
        this.AttemptBuild(associatedRtsLayer2, rtsPlayerLayerInterface, layerManager, any(rtsLayer).(IRTSLayer), itemIndex)
}


// throws Exception
func (this *UnitRTSFormInput) AttemptBuild(associatedRtsLayer IRTSLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer, itemIndex int) {
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
            rtsPlayerLayerInterface.Add(BuildingSound_GetInstance())
            capital.RemoveMoney(cost)
            AssignWaypointsUtil_GetInstance().Set(any(layerInterface).(*UnitLayer), any(associatedRtsLayer).(IAdvancedRTSGameLayer))
            layerManager.AppendAt(layerInterface, PlayersSingletonFactory_Total)
            var advancedRTSPlayerLayerInterface AdvancedRTSPlayerLayerInterface = any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface); _ = advancedRTSPlayerLayerInterface
            var advancedPlayerOwnedRTSLayers *AdvancedPlayerOwnedRTSLayers = advancedRTSPlayerLayerInterface.GetAdvancedPlayerOwnedRTSLayers(); _ = advancedPlayerOwnedRTSLayers
            advancedPlayerOwnedRTSLayers.AddUnit(layerInterface)
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.NewUnitGameNotificationEvent)
            }
        } else {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.NoMoneyGameNotificationEvent)
            }
        }
}

