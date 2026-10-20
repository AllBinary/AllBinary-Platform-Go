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

type CompositeRTSFormInput struct {
    RTSFormInput
    RtsFormInputArray []IRTSFormInput
    ItemIndex []int
    IsPrimaryWaypointCreator bool
}


func (this *CompositeRTSFormInput) initFields() {
    
    
    
}


func NewCompositeRTSFormInput(groupInterface []*Group, isPrimaryWaypointCreator bool, itemIndex []int) *CompositeRTSFormInput {
    this := &CompositeRTSFormInput{}
        this.RTSFormInput = *NewRTSFormInput(groupInterface)
        this.initFields()
        this.ItemIndex = itemIndex
        this.IsPrimaryWaypointCreator = isPrimaryWaypointCreator
        this.RtsFormInputArray = make([]IRTSFormInput, 2)
        this.RtsFormInputArray[0] = NewWaypointRTSFormInput(this.GroupInterfaceArray, isPrimaryWaypointCreator)
        this.RtsFormInputArray[1] = NewUnitRTSFormInput(this.GroupInterfaceArray)
    return this
}


// throws Exception
func (this *CompositeRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.RTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var size int = len(this.RtsFormInputArray); _ = size
        for index := int(0); index < size; index++ {
            this.RtsFormInputArray[index].SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        }
}


// throws Exception
func (this *CompositeRTSFormInput) ProcessAtPoint(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) {
        var scrollSelectionForm IScrollSelectionForm = rtsPlayerLayerInterface.GetCurrentScrollSelectionForm(); _ = scrollSelectionForm
        var index int = scrollSelectionForm.GetSelectedIndexForPoint(point); _ = index
        if this.IsStickyItemSelected() && associatedRtsLayer == CollidableDestroyableDamageableLayer_GetNullInstance() {
            if this.GetSelectedStickyItemIndex() <= this.ItemIndex[0] {
                this.RtsFormInputArray[0].ProcessAtPoint(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, point)
            }
        } else if index > this.ItemIndex[0] {
            this.RtsFormInputArray[1].ProcessAtPoint(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, point)
        }
}


// throws Exception
func (this *CompositeRTSFormInput) ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().INDEX_LABEL + index) + " > " + fmt.Sprint(this.ItemIndex[0]), this, commonStrings.PROCESS)
        if this.IsStickyItemSelected() && associatedRtsLayer == CollidableDestroyableDamageableLayer_GetNullInstance() {
            if this.GetSelectedStickyItemIndex() <= this.ItemIndex[0] {
                this.RtsFormInputArray[0].ProcessGameSpecific(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, item, index)
            }
        } else if index > this.ItemIndex[0] {
            this.RtsFormInputArray[1].ProcessGameSpecific(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, item, index)
        }
}


// throws Exception
func (this *CompositeRTSFormInput) ProcessSticky(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, point IGPoint) bool {
        return this.RtsFormInputArray[0].ProcessSticky(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, point)
}


// throws Exception
func (this *CompositeRTSFormInput) ProcessStickyGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, index int) {
        this.RtsFormInputArray[0].ProcessStickyGameSpecific(associatedRtsLayer, rtsPlayerLayerInterface, layerManager, item, index)
}


func (this *CompositeRTSFormInput) GetSelectedStickyItemIndex() int {
        return this.RtsFormInputArray[0].GetSelectedStickyItemIndex()
}


func (this *CompositeRTSFormInput) IsStickyItemSelected() bool {
        return this.RtsFormInputArray[0].IsStickyItemSelected()
}


func (this *CompositeRTSFormInput) SetStickyItemSelected(stickyItemSelected bool) {
        this.RtsFormInputArray[0].SetStickyItemSelected(stickyItemSelected)
}


func (this *CompositeRTSFormInput) GetSelectedStickyItem() IABCustomItem {
        return this.RtsFormInputArray[0].GetSelectedStickyItem()
}


func (this *CompositeRTSFormInput) SetSelectedStickyItem(selectedStickyItem IABCustomItem) {
        this.RtsFormInputArray[0].SetSelectedStickyItem(selectedStickyItem)
}

