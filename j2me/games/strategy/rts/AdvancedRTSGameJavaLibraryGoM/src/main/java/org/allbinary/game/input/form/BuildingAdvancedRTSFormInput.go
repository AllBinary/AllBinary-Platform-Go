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

type BuildingAdvancedRTSFormInput struct {
    BuildingRTSFormInput
    AtLeastOneHouseGameNotificationEvent *GameNotificationEvent
    BUILD_BUILDING_RTS_LAYER_EVENT *RTSLayerEvent
}


func (this *BuildingAdvancedRTSFormInput) initFields() {
    
    this.BUILD_BUILDING_RTS_LAYER_EVENT = NewRTSLayerEvent(NullPathFindingLayer_NULL_PATH_FINDING_LAYER)
}


func NewBuildingAdvancedRTSFormInput(groupInterface []*Group, isUnitProducer bool) *BuildingAdvancedRTSFormInput {
    this := &BuildingAdvancedRTSFormInput{}
        this.BuildingRTSFormInput = *NewBuildingRTSFormInput(groupInterface, isUnitProducer)
        this.initFields()
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.AtLeastOneHouseGameNotificationEvent = NewGameNotificationEvent(this, "Build At Least One House First", SmallIntegerSingletonFactory_GetInstance().GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
    return this
}


// throws Exception
func (this *BuildingAdvancedRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.BuildingRTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.AtLeastOneHouseGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


func (this *BuildingAdvancedRTSFormInput) IsPositionBlocked() bool {
        if DropCellPositionHistory_GetInstance().IsCellPositionWithDrop(this.GetSelectedGeographicCellPosition()) || WaypointCellPositionHistory_GetInstance().IsCellPositionWithDrop(this.GetSelectedGeographicCellPosition()) {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *BuildingAdvancedRTSFormInput) IsBuildAttemptValid(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerInterface IRTSLayer) bool {
        var advancedRTSPlayerLayerInterface AdvancedRTSPlayerLayerInterface = any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface); _ = advancedRTSPlayerLayerInterface
        var advancedPlayerOwnedRTSLayers *AdvancedPlayerOwnedRTSLayers = advancedRTSPlayerLayerInterface.GetAdvancedPlayerOwnedRTSLayers(); _ = advancedPlayerOwnedRTSLayers
        var totalHouses int = advancedPlayerOwnedRTSLayers.GetHouseList().Size(); _ = totalHouses
        if totalHouses < 1 && strings.Index(layerInterface.GetName(), string(" House")) < 0 {
            rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
            if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                GameNotificationEventHandler_GetInstance().FireEvent(this.AtLeastOneHouseGameNotificationEvent)
            }
            return false
        }
        return this.BuildingRTSFormInput.IsBuildAttemptValid(rtsPlayerLayerInterface, layerInterface)
}


// throws Exception
func (this *BuildingAdvancedRTSFormInput) Add(rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, layerInterface IRTSLayer) {
        this.BUILD_BUILDING_RTS_LAYER_EVENT.SetRtsLayer(layerInterface)
        BuildingEventHandler_GetInstance().FireEvent(this.BUILD_BUILDING_RTS_LAYER_EVENT)
        if (any(rtsPlayerLayerInterface).(AdvancedRTSPlayerLayerInterface)).IsLocalPlayer() {
            LocalPlayerBuildingEventHandler_GetInstance().FireEvent(this.BUILD_BUILDING_RTS_LAYER_EVENT)
        }
        this.BuildingRTSFormInput.Add(rtsPlayerLayerInterface, layerManager, layerInterface)
}

