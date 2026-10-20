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

type RTSPlayerGameInput struct {
    PlayerGameInput
    InputProcessorArray []IGameInputProcessor
    RemoveInputProcessorArray []IGameInputProcessor
    InputList IBasicArrayList
    IsSingleKeyProcessing bool
    GameCanvas IAllBinaryGameCanvas
    MotionGestureInputList IBasicArrayList
    ScrollPlayerGameInput *ScrollMapPlayerGameInput
    SelectedRTSLayerPlayerGameInput *SelectedRTSLayersPlayerGameInput
    TowerInfoPaintable *RTSLayerInfoPaintable
    RtsPlayerLayerInterface RTSPlayerLayerInterface
    LayerPositionFinderInterface LayerPositionFinderInterface
    SelectedRtsFormInput IRTSFormInput
}

// The Java class RTSPlayerGameInput has subclasses so its variables are this interface
type IRTSPlayerGameInput interface {
    IPlayerGameInput
    RTSPlayerGameInput_() *RTSPlayerGameInput
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent)
    Left()
    Right()
    InitInputProcessors()
    ProcessInput(layerManager IAllBinaryLayerManager)
    ProcessMotionInput(layerManager IAllBinaryLayerManager)
    Select(motionGestureEvent *MotionGestureEvent)
    SetSelectedRTSLayer(rtsLayer ICollidableDestroyableDamageableLayer, geographicMapCellPosition IGeographicMapCellPosition)
    Paint(graphics IGraphics)
    UpdatePaintable()
    GetSelectedBuildingPlayerGameInput() *SelectedRTSLayersPlayerGameInput
    GetSelectedRtsFormInput() IRTSFormInput
    SetSelectedRtsFormInput(selectedRtsFormInput IRTSFormInput)
    GetRtsPlayerLayerInterface() RTSPlayerLayerInterface
    GetRTSLayerInfoPaintable() *RTSLayerInfoPaintable
    GetMotionGestureInputList() IBasicArrayList
    GetGameCanvas() IAllBinaryGameCanvas
    IsIsSingleKeyProcessing() bool
    GetScrollPlayerGameInput() *ScrollMapPlayerGameInput
}

func (this *RTSPlayerGameInput) RTSPlayerGameInput_() *RTSPlayerGameInput {
    return this
}


func (this *RTSPlayerGameInput) initFields() {
    this.InputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    this.RemoveInputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    
    this.IsSingleKeyProcessing = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS) || Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_PRESS)
    
    this.MotionGestureInputList = NewBasicArrayListD()
    
    
    
    
    
    this.SelectedRtsFormInput = NullRTSFormInputFactory_GetInstance()
}


func NewRTSPlayerGameInput(gameCanvas IAllBinaryGameCanvas, inputList IBasicArrayList, playerInputId int, towerInfoPaintable *RTSLayerInfoPaintable, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerPositionFinderInterface LayerPositionFinderInterface, selectRTSLayerVisitorFactoryInterface SelectRTSLayerVisitorFactoryInterface) *RTSPlayerGameInput {
    this := &RTSPlayerGameInput{}
        this.PlayerGameInput = *NewPlayerGameInput(inputList, NewBasicArrayListD(), playerInputId)
        this.initFields()
        this.InitInputProcessors()
        this.GameCanvas = gameCanvas
        this.InputList = inputList
        this.TowerInfoPaintable = towerInfoPaintable
        this.RtsPlayerLayerInterface = rtsPlayerLayerInterface
        this.SelectedRTSLayerPlayerGameInput = NewSelectedRTSLayersPlayerGameInput(this.GetRTSLayerInfoPaintable(), this.GetRtsPlayerLayerInterface(), this.InputList, playerInputId, selectRTSLayerVisitorFactoryInterface)
        if this.RtsPlayerLayerInterface != NullRTSLayer_NULL_RTS_LAYER {
            this.SetSelectedRtsFormInput(this.RtsPlayerLayerInterface.GetRTSFormInput())
        }
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.GameCanvas.GetLayerManager()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.ScrollPlayerGameInput = NewScrollMapPlayerGameInput(geographicMapInterface, this.InputList, playerInputId)
        this.LayerPositionFinderInterface = layerPositionFinderInterface
    return this
}


// throws Exception
func (this *RTSPlayerGameInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.SelectedRTSLayerPlayerGameInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        if this.SelectedRtsFormInput != NullRTSFormInputFactory_GetInstance() {
            this.SelectedRtsFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        }
}


func (this *RTSPlayerGameInput) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onDisplayChangeEvent", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onDisplayChangeEvent")
            this.GetRTSLayerInfoPaintable().Update()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *RTSPlayerGameInput) Left() {
        this.RtsPlayerLayerInterface.GetCurrentScrollSelectionForm().ProcessInputKey(Canvas_LEFT)
}


// throws Exception
func (this *RTSPlayerGameInput) Right() {
        this.RtsPlayerLayerInterface.GetCurrentScrollSelectionForm().ProcessInputKey(Canvas_RIGHT)
}


func (this *RTSPlayerGameInput) InitInputProcessors() {
        this.InputProcessorArray[Canvas_LEFT] = NewRTSPlayerLeftGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_NUM0] = this.InputProcessorArray[Canvas_LEFT]
        this.InputProcessorArray[Canvas_RIGHT] = NewRTSPlayerRightGameInputProcessor(this)
        this.InputProcessorArray[Canvas_KEY_POUND] = this.InputProcessorArray[Canvas_RIGHT]
        this.RemoveInputProcessorArray[Canvas_KEY_NUM1] = NewRTSPlayerFireGameInputProcessor(this)
        this.RemoveInputProcessorArray[Canvas_KEY_NUM3] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        this.RemoveInputProcessorArray[Canvas_KEY_NUM3] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        this.RemoveInputProcessorArray[Canvas_LEFT] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        this.RemoveInputProcessorArray[Canvas_RIGHT] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        this.RemoveInputProcessorArray[Canvas_KEY_NUM0] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        this.RemoveInputProcessorArray[Canvas_KEY_POUND] = this.RemoveInputProcessorArray[Canvas_KEY_NUM1]
        GameInputProcessorUtil_Init(this.InputProcessorArray)
        GameInputProcessorUtil_Init(this.RemoveInputProcessorArray)
}


// throws Exception
func (this *RTSPlayerGameInput) ProcessInput(layerManager IAllBinaryLayerManager) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.PROCESS_INPUT, e)
                }
            }()
            this.ProcessMotionInput(layerManager)
            var size int = this.InputList.Size(); _ = size
            var key int = 0; _ = key
            for index := int(0); index < size; index++ {
                var gameKeyEvent *GameKeyEvent = any(this.InputList.Get(index)).(*GameKeyEvent); _ = gameKeyEvent
                key = gameKeyEvent.GetKey()
                this.GetScrollPlayerGameInput().ProcessInputKey(key)
                this.GetSelectedBuildingPlayerGameInput().ProcessInputKey(key)
                this.InputProcessorArray[key].ProcessEvent(layerManager, gameKeyEvent)
                this.RemoveInputProcessorArray[key].ProcessEvent(layerManager, gameKeyEvent)
            }
            if this.IsIsSingleKeyProcessing() {
                this.Clear()
            } else {
                this.Update()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *RTSPlayerGameInput) ProcessMotionInput(layerManager IAllBinaryLayerManager) {
}


// throws Exception
func (this *RTSPlayerGameInput) Select(motionGestureEvent *MotionGestureEvent) {
        var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.GameCanvas.GetLayerManager()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var x int = point.GetX() + allBinaryTiledLayer.GetXP(); _ = x
        var y int = point.GetY() + allBinaryTiledLayer.GetYP(); _ = y
        var geographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXYNoThrow(x, y); _ = geographicMapCellPosition
        if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            SecondaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
            var layer IAllBinaryLayer = this.LayerPositionFinderInterface.GetLayerInterface(geographicMapCellPosition); _ = layer
            if layer == AllBinaryLayer_NULL_ALLBINARY_LAYER {
                layer = CollidableDestroyableDamageableLayer_GetNullInstance()
            } else {
                geographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(layer.GetXP(), layer.GetYP())
            }
            var foundRTSLayer ICollidableDestroyableDamageableLayer = any(layer).(ICollidableDestroyableDamageableLayer); _ = foundRTSLayer
            this.SetSelectedRTSLayer(foundRTSLayer, geographicMapCellPosition)
        } else {
            var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Off Of Map -").Append(commonLabels.WIDTH_LABEL).Appendint(allBinaryTiledLayer.GetWidth()).Append(commonLabels.HEIGHT_LABEL).Appendint(allBinaryTiledLayer.GetHeight())), this, "select")
        }
}


// throws Exception
func (this *RTSPlayerGameInput) SetSelectedRTSLayer(rtsLayer ICollidableDestroyableDamageableLayer, geographicMapCellPosition IGeographicMapCellPosition) {
        this.GetSelectedBuildingPlayerGameInput().SetSelectedRTSLayer(rtsLayer)
}


func (this *RTSPlayerGameInput) Paint(graphics IGraphics) {
        var geographicMapCellPosition IGeographicMapCellPosition = this.GetSelectedRtsFormInput().GetSelectedGeographicCellPosition(); _ = geographicMapCellPosition
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.GameCanvas.GetLayerManager()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        graphics.SetColor(BasicColorFactory_GetInstance().GREEN.IntValue())
        var list IBasicArrayList = this.GetSelectedBuildingPlayerGameInput().GetPaintSelectedRTSLayersList(); _ = list
        var width int = 0; _ = width
        var height int = 0; _ = height
        if list.Size() > 0 {
            for index := int(list.Size() - 1); index >= 0; index-- {
                var rtsLayer IRTSLayer = any(list.Get(index)).(IRTSLayer); _ = rtsLayer
                width = rtsLayer.GetWidth()
                height = rtsLayer.GetHeight()
                graphics.DrawRect(rtsLayer.GetXP() - allBinaryTiledLayer.GetXP(), rtsLayer.GetYP() - allBinaryTiledLayer.GetYP(), width, height)
            }
        } else if geographicMapCellPosition != SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            var point IGPoint = geographicMapCellPosition.GetPoint(); _ = point
            width = allBinaryTiledLayer.GetCellWidth()
            height = allBinaryTiledLayer.GetCellHeight()
            graphics.DrawRect(point.GetX() - allBinaryTiledLayer.GetXP(), point.GetY() - allBinaryTiledLayer.GetYP(), width, height)
        }
}


func (this *RTSPlayerGameInput) UpdatePaintable() {
}


func (this *RTSPlayerGameInput) GetSelectedBuildingPlayerGameInput() *SelectedRTSLayersPlayerGameInput {
        return this.SelectedRTSLayerPlayerGameInput
}


func (this *RTSPlayerGameInput) GetSelectedRtsFormInput() IRTSFormInput {
        return this.SelectedRtsFormInput
}


func (this *RTSPlayerGameInput) SetSelectedRtsFormInput(selectedRtsFormInput IRTSFormInput) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("RTSFormInput: ").Append(StringUtil_GetInstance().ToString(selectedRtsFormInput))), this, "setSelectedRtsFormInput")
        this.SelectedRtsFormInput = selectedRtsFormInput
}


func (this *RTSPlayerGameInput) GetRtsPlayerLayerInterface() RTSPlayerLayerInterface {
        return this.RtsPlayerLayerInterface
}


func (this *RTSPlayerGameInput) GetRTSLayerInfoPaintable() *RTSLayerInfoPaintable {
        return this.TowerInfoPaintable
}


func (this *RTSPlayerGameInput) GetMotionGestureInputList() IBasicArrayList {
        return this.MotionGestureInputList
}


func (this *RTSPlayerGameInput) GetGameCanvas() IAllBinaryGameCanvas {
        return this.GameCanvas
}


func (this *RTSPlayerGameInput) IsIsSingleKeyProcessing() bool {
        return this.IsSingleKeyProcessing
}


func (this *RTSPlayerGameInput) GetScrollPlayerGameInput() *ScrollMapPlayerGameInput {
        return this.ScrollPlayerGameInput
}

