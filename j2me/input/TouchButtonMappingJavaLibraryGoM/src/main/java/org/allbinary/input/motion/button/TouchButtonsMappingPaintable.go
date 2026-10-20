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

type TouchButtonsMappingPaintable struct {
    Paintable
    LogUtil *LogUtil
    ForegroundColor int
    PaintableTable [][]IPaintable
    TouchButtonLocationHelper *TouchButtonLocationHelper
}


func (this *TouchButtonsMappingPaintable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.PaintableTable = func() [][]IPaintable { array := make([][]IPaintable, 0); for index := range array { array[index] = make([]IPaintable, 0) }; return array }()
    this.TouchButtonLocationHelper = NewTouchButtonLocationHelper()
}


func NewTouchButtonsMappingPaintable(basicColor *BasicColor) *TouchButtonsMappingPaintable {
    this := &TouchButtonsMappingPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.ForegroundColor = basicColor.IntValue()
        this.Init()
    return this
}


func (this *TouchButtonsMappingPaintable) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "updateRectangle", e)
                }
            }()
            this.PaintableTable = this.CreatePaintableTable()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *TouchButtonsMappingPaintable) CreatePaintableTable() [][]IPaintable {
        var totalColumns int = this.TouchButtonLocationHelper.GetTotalColumns(); _ = totalColumns
        var totalRows int = this.TouchButtonLocationHelper.GetTotalRows(); _ = totalRows
        var paintableTable [][]IPaintable = func() [][]IPaintable { array := make([][]IPaintable, totalColumns); for index := range array { array[index] = make([]IPaintable, totalRows) }; return array }(); _ = paintableTable
        var cellPositionFactory *CellPositionFactory = CellPositionFactory_GetInstance(); _ = cellPositionFactory
        var commonButtons *CommonButtons = CommonButtons_GetInstance(); _ = commonButtons
        for index := int(totalColumns - 1); index >= 0; index-- {
            for rowIndex := int(totalRows - 1); rowIndex >= 0; rowIndex-- {
                this.PaintableTable[index][rowIndex] = TouchButton_CreateButton(BasicTouchInputFactory_GetInstance().NONE, TouchButtonBlankResource_GetInstance(), commonButtons.NORMAL_BUTTON, cellPositionFactory.GetInstanceColRow(index, rowIndex), this.TouchButtonLocationHelper.GetColumnsRemainderHalf(), this.TouchButtonLocationHelper.GetRowsRemainderHalf())
            }
        }
        return paintableTable
}


func (this *TouchButtonsMappingPaintable) Paint(graphics IGraphics) {
        var totalColumns int = this.TouchButtonLocationHelper.GetTotalColumns(); _ = totalColumns
        var totalRows int = this.TouchButtonLocationHelper.GetTotalRows(); _ = totalRows
        for index := int(totalColumns - 1); index >= 0; index-- {
            for rowIndex := int(totalRows - 1); rowIndex >= 0; rowIndex-- {
                this.PaintableTable[index][rowIndex].Paint(graphics)
            }
        }
}

