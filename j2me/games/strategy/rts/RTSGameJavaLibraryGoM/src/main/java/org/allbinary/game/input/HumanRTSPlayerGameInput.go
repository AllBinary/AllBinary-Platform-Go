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

type HumanRTSPlayerGameInput struct {
    RTSPlayerGameInput
    IsDragging bool
    GameStateFactory *GameStateFactory
    RectangleCollisionUtil *RectangleCollisionUtil
    TouchMotionGestureFactory *TouchMotionGestureFactory
    PRESSED *MotionGestureInput
    RELEASED *MotionGestureInput
    TouchButtonsBuilderFactory *TouchButtonsBuilderFactory
    NotYoursGameNotificationEvent *GameNotificationEvent
    SelectedRtsLayer ICollidableDestroyableDamageableLayer
    StartPoint IGPoint
    EndPoint IGPoint
    POSSIBLE string
    AT string
    METHOD string
    ADDING string
    SPACE string
    MultiSelectPaintable *MultiSelectPaintable
}

// Implements BaseMotionGestureEventListener


func (this *HumanRTSPlayerGameInput) initFields() {
    this.IsDragging = false
    this.GameStateFactory = GameStateFactory_GetInstance()
    this.RectangleCollisionUtil = RectangleCollisionUtil_GetInstance()
    this.TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance()
    this.PRESSED = this.TouchMotionGestureFactory.PRESSED
    this.RELEASED = this.TouchMotionGestureFactory.RELEASED
    
    
    this.SelectedRtsLayer = CollidableDestroyableDamageableLayer_GetNullInstance()
    this.StartPoint = PointFactory_GetInstance().ZERO_ZERO
    this.EndPoint = PointFactory_GetInstance().ZERO_ZERO
    this.POSSIBLE = "Possible: "
    this.AT = "At: "
    this.METHOD = "makeSelection"
    this.ADDING = "Adding: "
    this.SPACE = CommonSeps_GetInstance().SPACE
    this.MultiSelectPaintable = NewMultiSelectPaintable()
}


func NewHumanRTSPlayerGameInput(gameCanvas IAllBinaryGameCanvas, inputList IBasicArrayList, playerInputId int, towerInfoPaintable *RTSLayerInfoPaintable, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerPositionFinderInterface LayerPositionFinderInterface, selectRTSLayerVisitorFactoryInterface SelectRTSLayerVisitorFactoryInterface, touchButtonsBuilderFactory *TouchButtonsBuilderFactory) *HumanRTSPlayerGameInput {
    this := &HumanRTSPlayerGameInput{}
        this.RTSPlayerGameInput = *NewRTSPlayerGameInput(gameCanvas, inputList, playerInputId, towerInfoPaintable, rtsPlayerLayerInterface, layerPositionFinderInterface, selectRTSLayerVisitorFactoryInterface)
        this.initFields()
        this.TouchButtonsBuilderFactory = touchButtonsBuilderFactory
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.NotYoursGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NOT_YOURS, SmallIntegerSingletonFactory_GetInstance().GetAt(2), basicColorFactory.PINK, BooleanFactory_GetInstance().FALSE)
        BasicMotionGesturesHandler_GetInstance().AddListenerInterface(this)
    return this
}


// throws Exception
func (this *HumanRTSPlayerGameInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.NotYoursGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


// throws Exception
func (this *HumanRTSPlayerGameInput) ProcessDraggingMotionInput(layerManager IAllBinaryLayerManager) bool {
        var motionGestureEvent *MotionGestureEvent = any(this.GetMotionGestureInputList().Get(this.GetMotionGestureInputList().Size() - 1)).(*MotionGestureEvent); _ = motionGestureEvent
        var motionGestureInput *MotionGestureInput = motionGestureEvent.GetMotionGesture(); _ = motionGestureInput
        var rtsFormInput IRTSFormInput = this.GetSelectedRtsFormInput(); _ = rtsFormInput
        if motionGestureInput == this.PRESSED {
            var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
            if this.GetRtsPlayerLayerInterface().GetCurrentScrollSelectionForm().IsInForm(point) {
                this.IsDragging = rtsFormInput.ProcessSticky(this.GetSelectedRtsLayer(), this.GetRtsPlayerLayerInterface(), layerManager, point)
            }
        } else if motionGestureInput == this.RELEASED {
            var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
            if this.GetRtsPlayerLayerInterface().GetCurrentScrollSelectionForm().IsInForm(point) {
                rtsFormInput.ProcessAtPoint(this.GetSelectedRtsLayer(), this.GetRtsPlayerLayerInterface(), layerManager, point)
            } else if rtsFormInput.IsStickyItemSelected() {
                var previousRtsFormInput IRTSFormInput = rtsFormInput; _ = previousRtsFormInput
                this.Select(motionGestureEvent)
                this.GetSelectedBuildingPlayerGameInput().SelectAllPreselected()
                previousRtsFormInput.ProcessAtPoint(this.GetSelectedRtsLayer(), this.GetRtsPlayerLayerInterface(), layerManager, point)
                this.GetSelectedBuildingPlayerGameInput().DeselectAllPreselected()
            } else {
                this.Select(motionGestureEvent)
            }
            rtsFormInput.SetStickyItemSelected(false)
            rtsFormInput.SetSelectedStickyItem(ABCustomItem_GetNullInstance())
            rtsFormInput.SetSelectedStickyItemIndex(-1)
            return true
        }
        return false
}


func (this *HumanRTSPlayerGameInput) MakeSelection() {
        var rtsLayerList IBasicArrayList = this.GetRtsPlayerLayerInterface().GetPlayerOwnedRTSLayers().PlayerOwnedRTSLayers_().RtsLayerList; _ = rtsLayerList
        var rectX1 int = this.StartPoint.GetX(); _ = rectX1
        var rectY1 int = this.StartPoint.GetY(); _ = rectY1
        var rectX2 int = this.EndPoint.GetX(); _ = rectX2
        var rectY2 int = this.EndPoint.GetY(); _ = rectY2
        if rectX1 > rectX2 {
            rectX2 = rectX1
            rectX1 = this.EndPoint.GetX()
        }
        if rectY1 > rectY2 {
            rectY2 = rectY1
            rectY1 = this.EndPoint.GetY()
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.AT)
        stringBuffer.Appendint(rectX1)
        stringBuffer.Append(this.SPACE)
        stringBuffer.Appendint(rectY1)
        stringBuffer.Append(this.SPACE)
        stringBuffer.Appendint(rectX2)
        stringBuffer.Append(this.SPACE)
        stringBuffer.Appendint(rectY2)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD)
        var rtsLayer IRTSLayer; _ = rtsLayer
        for index := int(rtsLayerList.Size() - 1); index >= 0; index-- {
            rtsLayer = any(rtsLayerList.Get(index)).(IRTSLayer)
            stringBuffer.Delete(0, stringBuffer.Length())
            var viewPosition IViewPositionBase = rtsLayer.GetViewPosition(); _ = viewPosition
            stringBuffer.Append(this.POSSIBLE)
            stringBuffer.Append(rtsLayer.GetName())
            stringBuffer.Append(this.SPACE)
            stringBuffer.Append(this.AT)
            stringBuffer.Appendint((viewPosition.GetX() + rtsLayer.GetHalfWidth()))
            stringBuffer.Append(this.SPACE)
            stringBuffer.Appendint((viewPosition.GetY() + rtsLayer.GetHalfHeight()))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD)
            if this.RectangleCollisionUtil.IsInside(rectX1, rectY1, rectX2, rectY2, viewPosition.GetX() + rtsLayer.GetHalfWidth(), viewPosition.GetY() + rtsLayer.GetHalfHeight()) {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.ADDING).Append(rtsLayer.GetName())), this, this.METHOD)
                rtsLayer.Select()
                this.GetSelectedBuildingPlayerGameInput().AddSelectedRTSLayer(rtsLayer)
            }
        }
        var gameLayer ICollidableDestroyableDamageableLayer = this.GetSelectedBuildingPlayerGameInput().GetLastSelectedRtsLayer(); _ = gameLayer
        if gameLayer != CollidableDestroyableDamageableLayer_GetNullInstance() {
            var lastRTSLayer IRTSLayer = any(gameLayer).(IRTSLayer); _ = lastRTSLayer
            if lastRTSLayer.GetScrollSelectionForm() != ScrollSelectionForm_GetNullScrollSelectionForm() {
                this.UpdateForm(lastRTSLayer.GetScrollSelectionForm(), lastRTSLayer.GetRTSFormInput())
                this.SetSelectedRtsLayer(lastRTSLayer)
            }
        }
        this.UpdatePaintable()
}


// throws Exception
func (this *HumanRTSPlayerGameInput) ProcessSelectionBoxMotionInput(layerManager IAllBinaryLayerManager) {
        for index := int(this.GetMotionGestureInputList().Size() - 1); index >= 0; index-- {
            var motionGestureEvent *MotionGestureEvent = any(this.GetMotionGestureInputList().Get(index)).(*MotionGestureEvent); _ = motionGestureEvent
            var motionGestureInput *MotionGestureInput = motionGestureEvent.GetMotionGesture(); _ = motionGestureInput
            var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
            if motionGestureInput == this.TouchMotionGestureFactory.PRESSED {
                this.StartPoint = point
                this.EndPoint = point
            } else if motionGestureInput == this.TouchMotionGestureFactory.RELEASED {
                this.EndPoint = point
                this.MakeSelection()
                this.StartPoint = PointFactory_GetInstance().ZERO_ZERO
                this.EndPoint = PointFactory_GetInstance().ZERO_ZERO
            } else {
                this.EndPoint = point
            }
        }
}


// throws Exception
func (this *HumanRTSPlayerGameInput) ProcessMotionInput(layerManager IAllBinaryLayerManager) {
        if this.GetMotionGestureInputList().Size() > 0 {
            var endDrag bool = this.ProcessDraggingMotionInput(layerManager); _ = endDrag
            if !this.IsDragging {
                this.ProcessSelectionBoxMotionInput(layerManager)
            }
            if endDrag {
                this.IsDragging = false
            }
        }
        this.GetMotionGestureInputList().Clear()
}


func (this *HumanRTSPlayerGameInput) SetSelectedRtsFormInput(selectedRtsFormInput IRTSFormInput) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START).Append(StringUtil_GetInstance().ToString(selectedRtsFormInput))), this, "setSelectedRtsFormInput")
        this.RTSPlayerGameInput.SetSelectedRtsFormInput(selectedRtsFormInput)
        this.SetSelectedRtsLayer(CollidableDestroyableDamageableLayer_GetNullInstance())
}


func (this *HumanRTSPlayerGameInput) UpdateFormForLayer(rtsLayer IRTSLayer) {
        var scrollSelectionForm IScrollSelectionForm = rtsLayer.GetScrollSelectionForm(); _ = scrollSelectionForm
        if scrollSelectionForm != ScrollSelectionForm_GetNullScrollSelectionForm() && this.GetRtsPlayerLayerInterface().GetGroupInterface()[0] == rtsLayer.GetGroupInterface()[0] {
            this.UpdateForm(scrollSelectionForm, rtsLayer.GetRTSFormInput())
            this.GetSelectedBuildingPlayerGameInput().SetSelectedRTSLayer(CollidableDestroyableDamageableLayer_GetNullInstance())
            this.SetSelectedRtsLayer(rtsLayer)
        } else {
            this.GetRtsPlayerLayerInterface().SetCurrentScrollSelectionForm(this.GetRtsPlayerLayerInterface().GetBuildingScrollSelectionForm())
            this.SetSelectedRtsFormInput(this.GetRtsPlayerLayerInterface().GetRTSFormInput())
            this.GetSelectedBuildingPlayerGameInput().SetSelectedRTSLayer(CollidableDestroyableDamageableLayer_GetNullInstance())
        }
}


func (this *HumanRTSPlayerGameInput) UpdateForm(scrollSelectionForm IScrollSelectionForm, rtfFormInput IRTSFormInput) {
        this.GetRtsPlayerLayerInterface().SetCurrentScrollSelectionForm(scrollSelectionForm)
        this.SetSelectedRtsFormInput(rtfFormInput)
}


// throws Exception
func (this *HumanRTSPlayerGameInput) SetSelectedRTSLayer(rtSLayer ICollidableDestroyableDamageableLayer, geographicMapCellPosition IGeographicMapCellPosition) {
        if rtSLayer != CollidableDestroyableDamageableLayer_GetNullInstance() && this.GetRtsPlayerLayerInterface().GetGroupInterface()[0] != rtSLayer.GetGroupInterface()[0] {
            this.GetRtsPlayerLayerInterface().Add(ErrorSound_GetInstance())
            GameNotificationEventHandler_GetInstance().FireEvent(this.NotYoursGameNotificationEvent)
            return
        }
        var gameCanvas IAllBinaryGameCanvas = this.GetGameCanvas(); _ = gameCanvas
        if rtSLayer != CollidableDestroyableDamageableLayer_GetNullInstance() {
            this.GetRtsPlayerLayerInterface().Add(SelectBuildingSound_GetInstance())
            var rtsLayer IRTSLayer = any(rtSLayer).(IRTSLayer); _ = rtsLayer
            rtsLayer.Select()
            this.UpdateFormForLayer(rtsLayer)
            if gameCanvas.GetGameState() == this.GameStateFactory.PLAYING_GAME_STATE {
                if rtsLayer.IsSelfUpgradeable() {
                    gameCanvas.UpdateCurrentTouchInputFactory(NewRTSTouchButtonsBuilder())
                } else {
                    gameCanvas.UpdateCurrentTouchInputFactory(NewNoBuildingSelectedTouchButtonsBuilder())
                }
            } else {
                var endLevelTouchButtonsBuilder *EndLevelTouchButtonsBuilder = NewEndLevelTouchButtonsBuilder(); _ = endLevelTouchButtonsBuilder
                if this.TouchButtonsBuilderFactory != nil {
                    endLevelTouchButtonsBuilder.Add(this.TouchButtonsBuilderFactory.GetInstance())
                }
                gameCanvas.UpdateCurrentTouchInputFactory(endLevelTouchButtonsBuilder)
            }
        } else {
            this.LogUtil.PutF("Set Player Default Form", this, "setSelectedRTSLayer")
            this.GetRtsPlayerLayerInterface().SetCurrentScrollSelectionForm(this.GetRtsPlayerLayerInterface().GetBuildingScrollSelectionForm())
            this.SetSelectedRtsFormInput(this.GetRtsPlayerLayerInterface().GetRTSFormInput())
            this.GetSelectedBuildingPlayerGameInput().SetSelectedRTSLayer(CollidableDestroyableDamageableLayer_GetNullInstance())
            if gameCanvas.GetGameState() == this.GameStateFactory.PLAYING_GAME_STATE {
                gameCanvas.UpdateCurrentTouchInputFactory(NewNoBuildingSelectedTouchButtonsBuilder())
            } else {
                var endLevelNoBuildingSelectedTouchButtonsBuilder *EndLevelNoBuildingSelectedTouchButtonsBuilder = NewEndLevelNoBuildingSelectedTouchButtonsBuilder(); _ = endLevelNoBuildingSelectedTouchButtonsBuilder
                if this.TouchButtonsBuilderFactory != nil {
                    endLevelNoBuildingSelectedTouchButtonsBuilder.Add(this.TouchButtonsBuilderFactory.GetInstance())
                }
                gameCanvas.UpdateCurrentTouchInputFactory(endLevelNoBuildingSelectedTouchButtonsBuilder)
            }
        }
        this.GetSelectedRtsFormInput().SetSelectedGeographicCellPosition(geographicMapCellPosition)
        this.GetSelectedBuildingPlayerGameInput().SetSelectedRTSLayer(rtSLayer)
        this.UpdatePaintable()
}


func (this *HumanRTSPlayerGameInput) UpdatePaintable() {
        var list IBasicArrayList = this.GetSelectedBuildingPlayerGameInput().GetSelectedBasicArrayList(); _ = list
        if list.Size() > 1 {
            var geographicMapCompositeInterface GeographicMapCompositeInterface = any(this.GetGameCanvas()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
            var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
            this.MultiSelectPaintable.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
            this.MultiSelectPaintable.Update(list)
            this.GetRTSLayerInfoPaintable().UpdateRTSLayerInfoSelection(this.MultiSelectPaintable)
        } else if list.Size() == 1 {
            var rtsLayer IRTSLayer = any(list.Get(0)).(IRTSLayer); _ = rtsLayer
            this.GetRTSLayerInfoPaintable().UpdateRTSLayerInfoSelection(rtsLayer.CreateHudPaintable())
        } else {
            this.GetRTSLayerInfoPaintable().UpdateRTSLayerInfoInit(NullInitUpdatePaintable_GetInstance())
        }
}


func (this *HumanRTSPlayerGameInput) Paint(graphics IGraphics) {
        this.RTSPlayerGameInput.Paint(graphics)
        if this.StartPoint != PointFactory_GetInstance().ZERO_ZERO && this.EndPoint != PointFactory_GetInstance().ZERO_ZERO {
            graphics.SetColor(BasicColorFactory_GetInstance().RED.IntValue())
            var rectX1 int = this.StartPoint.GetX(); _ = rectX1
            var rectY1 int = this.StartPoint.GetY(); _ = rectY1
            var rectX2 int = this.EndPoint.GetX(); _ = rectX2
            var rectY2 int = this.EndPoint.GetY(); _ = rectY2
            if rectX1 > rectX2 {
                rectX2 = rectX1
                rectX1 = this.EndPoint.GetX()
            }
            if rectY1 > rectY2 {
                rectY2 = rectY1
                rectY1 = this.EndPoint.GetY()
            }
            graphics.DrawRect(rectX1, rectY1, rectX2 - rectX1, rectY2 - rectY1)
        }
}


func (this *HumanRTSPlayerGameInput) OnMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
        this.GetMotionGestureInputList().Add(motionGestureEvent)
}


func (this *HumanRTSPlayerGameInput) OnScrolledMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
}


func (this *HumanRTSPlayerGameInput) GetSelectedRtsLayer() ICollidableDestroyableDamageableLayer {
        return this.SelectedRtsLayer
}


func (this *HumanRTSPlayerGameInput) SetSelectedRtsLayer(selectedRtsLayer ICollidableDestroyableDamageableLayer) {
        this.SelectedRtsLayer = selectedRtsLayer
}

