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

type VisibleCellPositionsSingleton struct {
    StationaryVisibleCellPositions [][]int16
    VisibleCellPositions [][]int16
    CurrentlyVisibleCellPositions [][]int16
    ALL_VISIBLE_TILEDLAYER *SimpleTiledLayer
    SimpleTiledLayer *SimpleTiledLayer
    PaintSimpleTiledLayer *SimpleTiledLayer
    CurrentIndex int
}

// Implements BuildingEventListenerInterface, ScrollMapEventListenerInterface

var VisibleCellPositionsSingleton_Instance *VisibleCellPositionsSingleton = NewVisibleCellPositionsSingleton()

func (this *VisibleCellPositionsSingleton) initFields() {
    this.StationaryVisibleCellPositions = NullUtil_GetInstance().NULL_SHORT_ARRAY_ARRAY
    this.VisibleCellPositions = NullUtil_GetInstance().NULL_SHORT_ARRAY_ARRAY
    this.CurrentlyVisibleCellPositions = NullUtil_GetInstance().NULL_SHORT_ARRAY_ARRAY
    this.ALL_VISIBLE_TILEDLAYER = NewSimpleTiledLayer(0, 0, 0, 0)
    this.SimpleTiledLayer = this.ALL_VISIBLE_TILEDLAYER
    this.PaintSimpleTiledLayer = this.ALL_VISIBLE_TILEDLAYER
    
}


func VisibleCellPositionsSingleton_GetInstance() *VisibleCellPositionsSingleton {
        return VisibleCellPositionsSingleton_Instance
}


func NewVisibleCellPositionsSingleton() *VisibleCellPositionsSingleton {
    this := &VisibleCellPositionsSingleton{}
    this.initFields()
        LocalPlayerBuildingEventHandler_GetInstance().AddListenerInterface(this)
        ScrollMapEventHandler_GetInstance().AddListenerInterface(this)
    return this
}


func (this *VisibleCellPositionsSingleton) Init(simpleTiledLayer *SimpleTiledLayer) {
        this.PaintSimpleTiledLayer = this.ALL_VISIBLE_TILEDLAYER
        this.CurrentIndex = 0
        var rows int = simpleTiledLayer.GetRows(); _ = rows
        var columns int = simpleTiledLayer.GetColumns(); _ = columns
        this.StationaryVisibleCellPositions = func() [][]int16 { array := make([][]int16, rows); for index := range array { array[index] = make([]int16, columns) }; return array }()
        this.VisibleCellPositions = func() [][]int16 { array := make([][]int16, rows); for index := range array { array[index] = make([]int16, columns) }; return array }()
        this.CurrentlyVisibleCellPositions = func() [][]int16 { array := make([][]int16, rows); for index := range array { array[index] = make([]int16, columns) }; return array }()
        this.SetSimpleTiledLayer(simpleTiledLayer)
}


func (this *VisibleCellPositionsSingleton) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *VisibleCellPositionsSingleton) OnBuildingEvent(event *RTSLayerEvent) {
        this.PaintSimpleTiledLayer = this.SimpleTiledLayer
}


// throws Exception
func (this *VisibleCellPositionsSingleton) OnMoveEvent(scrollMapEvent *ScrollMapEvent) {
        this.Move(scrollMapEvent.GetDx(), scrollMapEvent.GetDy())
}


func (this *VisibleCellPositionsSingleton) AddStationaryCellPositions(list IBasicArrayList) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            var cellPosition ICellPosition = any(list.Get(index)).(ICellPosition); _ = cellPosition
            this.StationaryVisibleCellPositions[cellPosition.GetRow()][cellPosition.GetColumn()]++
        }
}


func (this *VisibleCellPositionsSingleton) RemoveStationaryCellPositions(list IBasicArrayList) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            var cellPosition ICellPosition = any(list.Get(index)).(ICellPosition); _ = cellPosition
            this.StationaryVisibleCellPositions[cellPosition.GetRow()][cellPosition.GetColumn()]--
        }
}


func (this *VisibleCellPositionsSingleton) Update() {
        if this.CurrentIndex == 0 {
            var temp [][]int16 = this.CurrentlyVisibleCellPositions; _ = temp
            this.CurrentlyVisibleCellPositions = this.VisibleCellPositions
            this.VisibleCellPositions = temp
            for index := int(len(this.VisibleCellPositions) - 1); index >= 0; index-- {
                for index2 := int(len(this.VisibleCellPositions[0]) - 1); index2 >= 0; index2-- {
                    this.VisibleCellPositions[index][index2] = this.StationaryVisibleCellPositions[index][index2]
                }
            }
        }
        this.CurrentIndex++
        if this.CurrentIndex > 10 {
            this.CurrentIndex = 0
        }
}


func (this *VisibleCellPositionsSingleton) ShouldProcess() bool {
        if this.CurrentIndex == 0 {
            return true
        } else {
            return false
        }
}


func (this *VisibleCellPositionsSingleton) AddCellPositions(list IBasicArrayList) {
        for index := int(list.Size() - 1); index >= 0; index-- {
            var cellPosition ICellPosition = any(list.Get(index)).(ICellPosition); _ = cellPosition
            this.VisibleCellPositions[cellPosition.GetRow()][cellPosition.GetColumn()]++
        }
}


func (this *VisibleCellPositionsSingleton) IsVisible(cellPosition ICellPosition) bool {
        if this.VisibleCellPositions[cellPosition.GetRow()][cellPosition.GetColumn()] > 0 {
            return true
        } else {
            return false
        }
}


func (this *VisibleCellPositionsSingleton) Move(dx int, dy int) {
        this.GetSimpleTiledLayer().MoveDXY(dx, dy)
}


func (this *VisibleCellPositionsSingleton) Paint(graphics IGraphics) {
        this.PaintSimpleTiledLayer.Paint(graphics, this.CurrentlyVisibleCellPositions)
}


func (this *VisibleCellPositionsSingleton) SetSimpleTiledLayer(simpleTiledLayer *SimpleTiledLayer) {
        this.SimpleTiledLayer = simpleTiledLayer
}


func (this *VisibleCellPositionsSingleton) GetSimpleTiledLayer() *SimpleTiledLayer {
        return this.SimpleTiledLayer
}

