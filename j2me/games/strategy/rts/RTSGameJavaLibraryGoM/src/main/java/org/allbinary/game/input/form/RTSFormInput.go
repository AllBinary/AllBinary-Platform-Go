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

type RTSFormInput struct {
    LogUtil *LogUtil
    GroupCommonFactory *GroupCommonFactory
    Hashtable *ABHashtable[any, any]
    NewUnconstructedRTSLayerInterfaceArray []ICollidableDestroyableDamageableLayer
    GroupInterfaceArray []*Group
    SelectedGeographicCellPosition IGeographicMapCellPosition
    StickyItemSelected bool
    SelectedStickyItemIndex int
    SelectedStickyItem IABCustomItem
}

// The Java class RTSFormInput has subclasses so its variables are this interface
type IRTSFormInput interface {
    RTSFormInput_() *RTSFormInput
    SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager)
    ProcessAtPoint(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint)
    Process(layerManager IAllBinaryLayerManager)
    ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int)
    ProcessSticky(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) bool
    ProcessStickyGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int)
    GetItemAt(rtsPlayerLayerInterface RTSPlayerLayerInterface, point IGPoint) IABCustomItem
    GetIndexAt(rtsPlayerLayerInterface RTSPlayerLayerInterface, point IGPoint) int
    GetInstance(layerManager IAllBinaryLayerManager, aItem IABCustomItem, geographicMapCellPosition IGeographicMapCellPosition) ICollidableDestroyableDamageableLayer
    GetSelectedGeographicCellPosition() IGeographicMapCellPosition
    SetSelectedGeographicCellPosition(selectedGeographicCellPosition IGeographicMapCellPosition)
    IsStickyItemSelected() bool
    SetStickyItemSelected(stickyItemSelected bool)
    GetSelectedStickyItem() IABCustomItem
    SetSelectedStickyItem(selectedStickyItem IABCustomItem)
    GetSelectedStickyItemIndex() int
    SetSelectedStickyItemIndex(selectedStickyItemIndex int)
    GetHashtable() *ABHashtable[any, any]
}

func (this *RTSFormInput) RTSFormInput_() *RTSFormInput {
    return this
}


func (this *RTSFormInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GroupCommonFactory = GroupCommonFactory_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    this.NewUnconstructedRTSLayerInterfaceArray = make([]ICollidableDestroyableDamageableLayer, 7)
    
    this.SelectedGeographicCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    
    
    this.SelectedStickyItem = ABCustomItem_GetNullInstance()
}


func NewRTSFormInput(groupInterfaceArray []*Group) *RTSFormInput {
    this := &RTSFormInput{}
    this.initFields()
        var size int = len(this.NewUnconstructedRTSLayerInterfaceArray); _ = size
        for index := int(0); index < size; index++ {
            this.NewUnconstructedRTSLayerInterfaceArray[index] = CollidableDestroyableDamageableLayer_GetNullInstance()
        }
        this.GroupInterfaceArray = groupInterfaceArray
        this.Hashtable.Put(DirectionFactory_GetInstance().NAME, DirectionFactory_GetInstance().DOWN)
        if groupInterfaceArray != nil {
            this.Hashtable.Put(this.GroupCommonFactory.ID, groupInterfaceArray)
        }
    return this
}


// throws Exception
func (this *RTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
}


// throws Exception
func (this *RTSFormInput) ProcessAtPoint(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) {
        var index int = this.GetIndexAt(rtsPlayerLayerInterface, point); _ = index
        if index < 0 {
            return
        }
        var scrollSelectionForm IScrollSelectionForm = rtsPlayerLayerInterface.GetCurrentScrollSelectionForm(); _ = scrollSelectionForm
        var item IABCustomItem = scrollSelectionForm.Get(index); _ = item
        if item == nil {
            return
        }
        this.ProcessGameSpecific(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, item, index)
}


// throws Exception
func (this *RTSFormInput) Process(layerManager IAllBinaryLayerManager) {
        this.SetAllBinaryGameLayerManager(any(layerManager).(IAllBinaryGameLayerManager))
}


// throws Exception
func (this *RTSFormInput) ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int) {
}


// throws Exception
func (this *RTSFormInput) ProcessSticky(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) bool {
        var index int = this.GetIndexAt(rtsPlayerLayerInterface, point); _ = index
        if index < 0 {
            return false
        }
        var scrollSelectionForm IScrollSelectionForm = rtsPlayerLayerInterface.GetCurrentScrollSelectionForm(); _ = scrollSelectionForm
        var item IABCustomItem = scrollSelectionForm.Get(index); _ = item
        if item == nil {
            return false
        }
        this.ProcessStickyGameSpecific(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, item, index)
        return true
}


// throws Exception
func (this *RTSFormInput) ProcessStickyGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int) {
}


// throws Exception
func (this *RTSFormInput) GetItemAt(rtsPlayerLayerInterface RTSPlayerLayerInterface, point IGPoint) IABCustomItem {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Point: ").Append(StringUtil_GetInstance().ToString(point))), this, "getItemAt")
        var scrollSelectionForm IScrollSelectionForm = rtsPlayerLayerInterface.GetCurrentScrollSelectionForm(); _ = scrollSelectionForm
        return scrollSelectionForm.GetSelectedItem(point)
}


// throws Exception
func (this *RTSFormInput) GetIndexAt(rtsPlayerLayerInterface RTSPlayerLayerInterface, point IGPoint) int {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Point: ").Append(StringUtil_GetInstance().ToString(point))), this, "getItemAt")
        var scrollSelectionForm IScrollSelectionForm = rtsPlayerLayerInterface.GetCurrentScrollSelectionForm(); _ = scrollSelectionForm
        return scrollSelectionForm.GetSelectedIndexForPoint(point)
}


// throws Exception
func (this *RTSFormInput) GetInstance(layerManager IAllBinaryLayerManager, aItem IABCustomItem, geographicMapCellPosition IGeographicMapCellPosition) ICollidableDestroyableDamageableLayer {
        var item *LayerInterfaceFactoryImageItem = any(aItem).(*LayerInterfaceFactoryImageItem); _ = item
        var layerInterfaceFactoryInterface LayerInterfaceFactoryInterface = item.GetLayerInterfaceFactoryInterface(); _ = layerInterfaceFactoryInterface
        var cellPoint IGPoint = geographicMapCellPosition.GetPoint(); _ = cellPoint
        if layerInterfaceFactoryInterface != nil {
            this.Hashtable.Put(AllBinaryGameLayerManager_ID, layerManager)
            return any(layerInterfaceFactoryInterface.GetNextInstance(this.GetHashtable(), cellPoint.GetX(), cellPoint.GetY(), cellPoint.GetZ())).(IRTSLayer)
        }
        return CollidableDestroyableDamageableLayer_GetNullInstance()
}


func (this *RTSFormInput) GetSelectedGeographicCellPosition() IGeographicMapCellPosition {
        return this.SelectedGeographicCellPosition
}


func (this *RTSFormInput) SetSelectedGeographicCellPosition(selectedGeographicCellPosition IGeographicMapCellPosition) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Selected GeographicMapCellPosition: ").Append(fmt.Sprint(selectedGeographicCellPosition))), this, "setSelectedGeographicCellPosition")
        this.SelectedGeographicCellPosition = selectedGeographicCellPosition
}


func (this *RTSFormInput) IsStickyItemSelected() bool {
        return this.StickyItemSelected
}


func (this *RTSFormInput) SetStickyItemSelected(stickyItemSelected bool) {
        this.StickyItemSelected = stickyItemSelected
}


func (this *RTSFormInput) GetSelectedStickyItem() IABCustomItem {
        return this.SelectedStickyItem
}


func (this *RTSFormInput) SetSelectedStickyItem(selectedStickyItem IABCustomItem) {
        this.SelectedStickyItem = selectedStickyItem
}


func (this *RTSFormInput) GetSelectedStickyItemIndex() int {
        return this.SelectedStickyItemIndex
}


func (this *RTSFormInput) SetSelectedStickyItemIndex(selectedStickyItemIndex int) {
        this.SelectedStickyItemIndex = selectedStickyItemIndex
}


func (this *RTSFormInput) GetHashtable() *ABHashtable[any, any] {
        return this.Hashtable
}

