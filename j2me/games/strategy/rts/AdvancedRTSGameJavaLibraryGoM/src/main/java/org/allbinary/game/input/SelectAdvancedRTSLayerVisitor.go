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

type SelectAdvancedRTSLayerVisitor struct {
    Visitor
    SelectedRTSLayersPlayerGameInput *SelectedRTSLayersPlayerGameInput
}


func (this *SelectAdvancedRTSLayerVisitor) initFields() {
    
}


func NewSelectAdvancedRTSLayerVisitor(selectedRTSLayersPlayerGameInput *SelectedRTSLayersPlayerGameInput) *SelectAdvancedRTSLayerVisitor {
    this := &SelectAdvancedRTSLayerVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
        this.SelectedRTSLayersPlayerGameInput = selectedRTSLayersPlayerGameInput
    return this
}


func (this *SelectAdvancedRTSLayerVisitor) Visit(anyType any) any {
        var selectedLayer ICollidableDestroyableDamageableLayer = any(anyType).(ICollidableDestroyableDamageableLayer); _ = selectedLayer
        if selectedLayer != CollidableDestroyableDamageableLayer_GetNullInstance() && selectedLayer.GetType() == BuildingLayer_GetStaticType() {
            var rtsGameLayer IAdvancedRTSGameLayer = any(selectedLayer).(IAdvancedRTSGameLayer); _ = rtsGameLayer
            var assignWaypointsUtil *AssignWaypointsUtil = AssignWaypointsUtil_GetInstance(); _ = assignWaypointsUtil
            var list IBasicArrayList = this.SelectedRTSLayersPlayerGameInput.GetSelectedBasicArrayList(); _ = list
            var currentRTSLayer IRTSLayer; _ = currentRTSLayer
            var unitLayer *UnitLayer; _ = unitLayer
            for index := int(list.Size() - 1); index >= 0; index-- {
                currentRTSLayer = any(list.Get(index)).(IRTSLayer)
                if currentRTSLayer.GetType() == UnitLayer_GetStaticType() {
                    unitLayer = any(currentRTSLayer).(*UnitLayer)
                    unitLayer.SetParentLayer(rtsGameLayer)
                    assignWaypointsUtil.Set(unitLayer, rtsGameLayer)
                }
            }
        }
        return NullUtil_GetInstance().NULL_OBJECT
}

