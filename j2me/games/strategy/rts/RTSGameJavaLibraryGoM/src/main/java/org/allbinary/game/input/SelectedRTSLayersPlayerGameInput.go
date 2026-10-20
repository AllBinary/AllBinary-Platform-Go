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

type SelectedRTSLayersPlayerGameInput struct {
    PlayerGameInput
    InputProcessorArray []IGameInputProcessor
    List IBasicArrayList
    IsSingleKeyProcessing bool
    SelectedRTSLayersList IBasicArrayList
    PreSelectedRTSLayersList IBasicArrayList
    PaintSelectedRTSLayersList IBasicArrayList
    RtsPlayerLayerInterface RTSPlayerLayerInterface
    SelectRTSLayerVisitorInterface IVisitor
    UpgradeGameNotificationEvent *GameNotificationEvent
    NoMoneyGameNotificationEvent *GameNotificationEvent
    DowngradeGameNotificationEvent *GameNotificationEvent
}


func (this *SelectedRTSLayersPlayerGameInput) initFields() {
    this.InputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    
    this.IsSingleKeyProcessing = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS) || Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_PRESS)
    this.SelectedRTSLayersList = NewBasicArrayListD()
    this.PreSelectedRTSLayersList = NewBasicArrayListD()
    this.PaintSelectedRTSLayersList = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    
    
    
    
    
}


func NewSelectedRTSLayersPlayerGameInput(towerInfoPaintable *RTSLayerInfoPaintable, rtsPlayerLayerInterface RTSPlayerLayerInterface, list IBasicArrayList, playerInputId int, selectRTSLayerVisitorFactoryInterface SelectRTSLayerVisitorFactoryInterface) *SelectedRTSLayersPlayerGameInput {
    this := &SelectedRTSLayersPlayerGameInput{}
        this.PlayerGameInput = *NewPlayerGameInput(list, NewBasicArrayListD(), playerInputId)
        this.initFields()
        this.InitInputProcessors()
        this.RtsPlayerLayerInterface = rtsPlayerLayerInterface
        this.List = list
        this.SelectRTSLayerVisitorInterface = selectRTSLayerVisitorFactoryInterface.Create(this)
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.UpgradeGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().UPGRADE, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.NoMoneyGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NO_MONEY, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        this.DowngradeGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().DOWNGRADE, smallIntegerSingletonFactory.GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
    return this
}


// throws Exception
func (this *SelectedRTSLayersPlayerGameInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.UpgradeGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.NoMoneyGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
        this.DowngradeGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


func (this *SelectedRTSLayersPlayerGameInput) IsSelected(rtsLayer IRTSLayer) bool {
        if this.SelectedRTSLayersList.Contains(rtsLayer) {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *SelectedRTSLayersPlayerGameInput) Upgrade() {
        var anyChanged bool = false; _ = anyChanged
        for index := int(this.SelectedRTSLayersList.Size() - 1); index >= 0; index-- {
            var rtsLayer IRTSLayer = any(this.SelectedRTSLayersList.Get(index)).(IRTSLayer); _ = rtsLayer
            if rtsLayer.IsUpgradeable() {
                var capital ICapital = this.RtsPlayerLayerInterface.GetCapital(); _ = capital
                var upgradeCost int = rtsLayer.GetUpgradeCost(); _ = upgradeCost
                if upgradeCost <= capital.GetTotalMoney() {
                    anyChanged = true
                    this.RtsPlayerLayerInterface.Add(UpgradeSound_GetInstance())
                    rtsLayer.Upgrade()
                    capital.RemoveMoney(upgradeCost)
                    if !this.RtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                        GameNotificationEventHandler_GetInstance().FireEvent(this.UpgradeGameNotificationEvent)
                    }
                } else {
                    this.RtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                    if !this.RtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                        GameNotificationEventHandler_GetInstance().FireEvent(this.NoMoneyGameNotificationEvent)
                    }
                }
            }
        }
        if anyChanged {
            var rtsPlayerGameInput IRTSPlayerGameInput = (any(this.RtsPlayerLayerInterface.GetPlayerGameInput()).(IRTSPlayerGameInput)); _ = rtsPlayerGameInput
            rtsPlayerGameInput.UpdatePaintable()
        }
}


// throws Exception
func (this *SelectedRTSLayersPlayerGameInput) Downgrade() {
        var anyChanged bool = false; _ = anyChanged
        for index := int(this.SelectedRTSLayersList.Size()); Java_PreDecrement(&index) >= 0;  {
            var rtsLayer IRTSLayer = any(this.SelectedRTSLayersList.Get(index)).(IRTSLayer); _ = rtsLayer
            if rtsLayer.IsDowngradeable() {
                anyChanged = true
                this.RtsPlayerLayerInterface.Add(DowngradeSound_GetInstance())
                var downgradeCost int = rtsLayer.GetDowngradeCost(); _ = downgradeCost
                rtsLayer.Downgrade()
                var capital ICapital = this.RtsPlayerLayerInterface.GetCapital(); _ = capital
                capital.AddMoney(downgradeCost)
                if !this.RtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.DowngradeGameNotificationEvent)
                }
            }
        }
        if anyChanged {
            var rtsPlayerGameInput IRTSPlayerGameInput = (any(this.RtsPlayerLayerInterface.GetPlayerGameInput()).(IRTSPlayerGameInput)); _ = rtsPlayerGameInput
            rtsPlayerGameInput.UpdatePaintable()
        }
}


func (this *SelectedRTSLayersPlayerGameInput) InitInputProcessors() {
        this.InputProcessorArray[Canvas_KEY_NUM1] = NewSelectedRTSLayersPlayerUpgradeGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_NUM3] = NewSelectedRTSLayersPlayerDowngradeGameInputProcessor(this)
        GameInputProcessorUtil_Init(this.InputProcessorArray)
}


// throws Exception
func (this *SelectedRTSLayersPlayerGameInput) ProcessInputKey(key int) {
        if this.SelectedRTSLayersList != nil {
            this.InputProcessorArray[key].ProcessEvent(AllBinaryGameLayerManager_GetNullInstance(), GameKeyEvent_NONE)
        }
}


// throws Exception
func (this *SelectedRTSLayersPlayerGameInput) ProcessInput(layerManager IAllBinaryLayerManager) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.PROCESS_INPUT, e)
                }
            }()
            var size int = this.List.Size(); _ = size
            var key int = 0; _ = key
            for index := int(0); index < size; index++ {
                var gameKeyEvent *GameKeyEvent = any(this.List.Get(index)).(*GameKeyEvent); _ = gameKeyEvent
                key = gameKeyEvent.GetKey()
                this.ProcessInputKey(key)
            }
            if this.IsSingleKeyProcessing {
                this.Clear()
            } else {
                this.Update()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *SelectedRTSLayersPlayerGameInput) GetSelectedBasicArrayList() IBasicArrayList {
        return this.SelectedRTSLayersList
}


func (this *SelectedRTSLayersPlayerGameInput) AddSelectedRTSLayer(selectedLayer IRTSLayer) {
        this.PaintSelectedRTSLayersList = BasicArrayListUtil_GetInstance().GetImmutableInstance()
        if selectedLayer == CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.DeselectAll()
            this.SelectedRTSLayersList.Clear()
        } else {
            if !this.SelectedRTSLayersList.Contains(selectedLayer) {
                this.SelectedRTSLayersList.Add(selectedLayer)
            }
        }
        this.PaintSelectedRTSLayersList = this.SelectedRTSLayersList
}


func (this *SelectedRTSLayersPlayerGameInput) SetSelectedRTSLayer(selectedLayer ICollidableDestroyableDamageableLayer) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Selected Layer: ")
        if selectedLayer != nil {
            stringBuffer.Append(selectedLayer.GetName())
        }
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "setSelectedRTSLayer")
        this.PaintSelectedRTSLayersList = BasicArrayListUtil_GetInstance().GetImmutableInstance()
        this.SelectRTSLayerVisitorInterface.Visit(selectedLayer)
        this.DeselectAll()
        if this.SelectedRTSLayersList.Size() > 0 {
            this.GetPreSelectedRTSLayersList().Clear()
            var tempList IBasicArrayList = this.GetPreSelectedRTSLayersList(); _ = tempList
            this.PreSelectedRTSLayersList = this.SelectedRTSLayersList
            this.SelectedRTSLayersList = tempList
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Preselected: ").Append(fmt.Sprint(this.PreSelectedRTSLayersList))), this, "setSelectedRTSLayer")
        }
        if selectedLayer != CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.SelectedRTSLayersList.Add(selectedLayer)
        }
        this.PaintSelectedRTSLayersList = this.SelectedRTSLayersList
}


func (this *SelectedRTSLayersPlayerGameInput) GetLastSelectedRtsLayer() ICollidableDestroyableDamageableLayer {
        if this.IsAnyRTSLayerSelected() {
            return any(this.SelectedRTSLayersList.Get(this.SelectedRTSLayersList.Size() - 1)).(ICollidableDestroyableDamageableLayer)
        } else {
            return CollidableDestroyableDamageableLayer_GetNullInstance()
        }
}


func (this *SelectedRTSLayersPlayerGameInput) SelectAllPreselected() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Select all Preselected: ")
        stringBuffer.Append(fmt.Sprint(this.PreSelectedRTSLayersList))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selectAllPreselected")
        for index := int(this.PreSelectedRTSLayersList.Size() - 1); index >= 0; index-- {
            var rtsLayer IRTSLayer = any(this.PreSelectedRTSLayersList.Get(index)).(IRTSLayer); _ = rtsLayer
            rtsLayer.Select()
        }
}


func (this *SelectedRTSLayersPlayerGameInput) DeselectAllPreselected() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Deselect all Preselected: ")
        stringBuffer.Append(fmt.Sprint(this.PreSelectedRTSLayersList))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "deselectAllPreselected")
        for index := int(this.PreSelectedRTSLayersList.Size() - 1); index >= 0; index-- {
            var rtsLayer IRTSLayer = any(this.PreSelectedRTSLayersList.Get(index)).(IRTSLayer); _ = rtsLayer
            rtsLayer.Deselect()
        }
        this.PreSelectedRTSLayersList.Clear()
}


func (this *SelectedRTSLayersPlayerGameInput) DeselectAll() {
        for index := int(this.SelectedRTSLayersList.Size() - 1); index >= 0; index-- {
            var rtsLayer IRTSLayer = any(this.SelectedRTSLayersList.Get(index)).(IRTSLayer); _ = rtsLayer
            rtsLayer.Deselect()
        }
}


func (this *SelectedRTSLayersPlayerGameInput) IsAnyRTSLayerSelected() bool {
        if this.SelectedRTSLayersList.Size() == 0 {
            return false
        }
        return true
}


func (this *SelectedRTSLayersPlayerGameInput) GetPaintSelectedRTSLayersList() IBasicArrayList {
        return this.PaintSelectedRTSLayersList
}


func (this *SelectedRTSLayersPlayerGameInput) GetPreSelectedRTSLayersList() IBasicArrayList {
        return this.PreSelectedRTSLayersList
}

