/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
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

type DropCellPositionHistory struct {
    CellPositionFactory *CellPositionFactory
    PositionList IBasicArrayList
    LayerList IBasicArrayList
}

var DropCellPositionHistory_SINGLETON *DropCellPositionHistory = NewDropCellPositionHistory()

func (this *DropCellPositionHistory) initFields() {
    this.CellPositionFactory = CellPositionFactory_GetInstance()
    this.PositionList = NewBasicArrayListD()
    this.LayerList = NewBasicArrayListD()
}

func NewDropCellPositionHistory() *DropCellPositionHistory {
    this := &DropCellPositionHistory{}
    this.initFields()
    return this
}


func DropCellPositionHistory_GetInstance() *DropCellPositionHistory {
        return DropCellPositionHistory_SINGLETON
}


func (this *DropCellPositionHistory) Add(cellPosition ICellPosition, layerInterface IAllBinaryLayer) {
        this.PositionList.Add(cellPosition)
        this.LayerList.Add(layerInterface)
}


func (this *DropCellPositionHistory) AddAll(list IBasicArrayList, layerInterface IAllBinaryLayer) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            this.Add(any(list.Get(index)).(ICellPosition), layerInterface)
        }
}


func (this *DropCellPositionHistory) Remove(cellPosition ICellPosition) {
        var index int = this.PositionList.IndexOf(cellPosition); _ = index
        if index >= 0 {
            this.PositionList.RemoveAt(index)
            this.LayerList.RemoveAt(index)
        }
}


func (this *DropCellPositionHistory) RemoveAll(layerInterface IAllBinaryLayer) {
        var index int = 0; _ = index
        for index != -1 {
            index = this.LayerList.IndexOf(layerInterface)
            if index >= 0 {
                this.PositionList.RemoveAt(index)
                this.LayerList.RemoveAt(index)
            }
        }
}


func (this *DropCellPositionHistory) Clear() {
        this.PositionList.Clear()
        this.LayerList.Clear()
}


func (this *DropCellPositionHistory) IsCellPositionWithDrop(cellPosition ICellPosition) bool {
        return this.PositionList.Contains(cellPosition)
}


func (this *DropCellPositionHistory) AnyCellPositionWithDrop(list IBasicArrayList) bool {
        for index := int(list.Size() - 1); index >= 0; index-- {
            if this.IsCellPositionWithDrop(any(list.Get(index)).(ICellPosition)) {
                return true
            }
        }
        return false
}


func (this *DropCellPositionHistory) GetCellPositionWithDrop(list IBasicArrayList) ICellPosition {
        var cellPosition ICellPosition; _ = cellPosition
        for index := int(list.Size() - 1); index >= 0; index-- {
            cellPosition = any(list.Get(index)).(ICellPosition)
            if this.IsCellPositionWithDrop(cellPosition) {
                return cellPosition
            }
        }
        return this.CellPositionFactory.NONE
}


func (this *DropCellPositionHistory) GetLayerInterface(cellPosition ICellPosition) IAllBinaryLayer {
        var index int = this.PositionList.IndexOf(cellPosition); _ = index
        if index >= 0 {
            return any(this.LayerList.Get(index)).(IAllBinaryLayer)
        } else {
            return AllBinaryLayer_NULL_ALLBINARY_LAYER
        }
}

