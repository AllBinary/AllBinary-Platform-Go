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

type WaypointCellPositionHistory struct {
    PositionList IBasicArrayList
    LayerList IBasicArrayList
}

var WaypointCellPositionHistory_SINGLETON *WaypointCellPositionHistory = NewWaypointCellPositionHistory()

func (this *WaypointCellPositionHistory) initFields() {
    this.PositionList = NewBasicArrayListD()
    this.LayerList = NewBasicArrayListD()
}

func NewWaypointCellPositionHistory() *WaypointCellPositionHistory {
    this := &WaypointCellPositionHistory{}
    this.initFields()
    return this
}


func WaypointCellPositionHistory_GetInstance() *WaypointCellPositionHistory {
        return WaypointCellPositionHistory_SINGLETON
}


func (this *WaypointCellPositionHistory) Add(cellPosition ICellPosition, layerInterface IAllBinaryLayer) {
        this.PositionList.Add(cellPosition)
        this.LayerList.Add(layerInterface)
}


func (this *WaypointCellPositionHistory) AddList(list IBasicArrayList, layerInterface IAllBinaryLayer) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            this.Add(any(list.Get(index)).(ICellPosition), layerInterface)
        }
}


func (this *WaypointCellPositionHistory) Remove(cellPosition ICellPosition) {
        var index int = this.PositionList.IndexOf(cellPosition); _ = index
        if index >= 0 {
            this.PositionList.RemoveAt(index)
            this.LayerList.RemoveAt(index)
        }
}


func (this *WaypointCellPositionHistory) RemoveLayer(layerInterface IAllBinaryLayer) {
        var index int = 0; _ = index
        for index != -1 {
            index = this.LayerList.IndexOf(layerInterface)
            if index >= 0 {
                this.PositionList.RemoveAt(index)
                this.LayerList.RemoveAt(index)
            }
        }
}


func (this *WaypointCellPositionHistory) Clear() {
        this.PositionList.Clear()
        this.LayerList.Clear()
}


func (this *WaypointCellPositionHistory) IsCellPositionWithDrop(cellPosition ICellPosition) bool {
        return this.PositionList.Contains(cellPosition)
}


func (this *WaypointCellPositionHistory) AnyCellPositionWithDrop(list IBasicArrayList) bool {
        for index := int(list.Size() - 1); index >= 0; index-- {
            if this.IsCellPositionWithDrop(any(list.Get(index)).(ICellPosition)) {
                return true
            }
        }
        return false
}


func (this *WaypointCellPositionHistory) GetLayerInterface(cellPosition ICellPosition) IAllBinaryLayer {
        var index int = this.PositionList.IndexOf(cellPosition); _ = index
        if index >= 0 {
            return any(this.LayerList.Get(index)).(IAllBinaryLayer)
        } else {
            return AllBinaryLayer_NULL_ALLBINARY_LAYER
        }
}

