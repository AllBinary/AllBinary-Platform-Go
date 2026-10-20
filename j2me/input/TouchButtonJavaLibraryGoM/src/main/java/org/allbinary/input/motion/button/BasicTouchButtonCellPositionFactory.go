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

type BasicTouchButtonCellPositionFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TOP_RIGHT ICellPosition
    TOP_LEFT ICellPosition
    SECOND_FROM_TOP_LEFT ICellPosition
    SECOND_FROM_TOP_RIGHT ICellPosition
    TOP_SECOND_FROM_LEFT ICellPosition
    TOP_SECOND_FROM_RIGHT ICellPosition
    BOTTOM_RIGHT ICellPosition
    BOTTOM_SECOND_FROM_RIGHT ICellPosition
    SECOND_FROM_BOTTOM_RIGHT ICellPosition
    THIRD_FROM_BOTTOM_RIGHT ICellPosition
    FOURTH_FROM_BOTTOM_RIGHT ICellPosition
    SECOND_FROM_BOTTOM_SECOND_FROM_RIGHT ICellPosition
    SECOND_FROM_BOTTOM_THIRD_FROM_RIGHT ICellPosition
    SECOND_FROM_BOTTOM_FOURTH_FROM_RIGHT ICellPosition
    BOTTOM_THIRD_FROM_RIGHT ICellPosition
    BOTTOM_FOURTH_FROM_RIGHT ICellPosition
    BOTTOM_SECOND_FROM_LEFT ICellPosition
    SECOND_FROM_BOTTOM_LEFT ICellPosition
    THIRD_FROM_BOTTOM_LEFT ICellPosition
    BOTTOM_LEFT ICellPosition
    SECOND_FROM_BOTTOM_SECOND_FROM_LEFT ICellPosition
}


func (this *BasicTouchButtonCellPositionFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewBasicTouchButtonCellPositionFactory() *BasicTouchButtonCellPositionFactory {
    this := &BasicTouchButtonCellPositionFactory{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        var touchButtonLocationHelper *TouchButtonLocationHelper = NewTouchButtonLocationHelper(); _ = touchButtonLocationHelper
        var cellPositionFactory *CellPositionFactory = CellPositionFactory_GetInstance(); _ = cellPositionFactory
        cellPositionFactory.Init(touchButtonLocationHelper.GetTotalColumns(), touchButtonLocationHelper.GetTotalRows())
        this.TOP_LEFT = cellPositionFactory.GetInstanceColRow(0, 0)
        this.TOP_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, 0)
        this.BOTTOM_LEFT = cellPositionFactory.GetInstanceColRow(0, touchButtonLocationHelper.GetTotalRows() - 1)
        this.BOTTOM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, touchButtonLocationHelper.GetTotalRows() - 1)
        this.BOTTOM_SECOND_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 2, touchButtonLocationHelper.GetTotalRows() - 1)
        this.SECOND_FROM_BOTTOM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, touchButtonLocationHelper.GetTotalRows() - 2)
        this.SECOND_FROM_BOTTOM_SECOND_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 2, touchButtonLocationHelper.GetTotalRows() - 2)
        this.BOTTOM_SECOND_FROM_LEFT = cellPositionFactory.GetInstanceColRow(1, touchButtonLocationHelper.GetTotalRows() - 1)
        this.SECOND_FROM_BOTTOM_LEFT = cellPositionFactory.GetInstanceColRow(0, touchButtonLocationHelper.GetTotalRows() - 2)
        this.SECOND_FROM_BOTTOM_SECOND_FROM_LEFT = cellPositionFactory.GetInstanceColRow(1, touchButtonLocationHelper.GetTotalRows() - 2)
        if touchButtonLocationHelper.GetTotalRows() > 2 {
            this.SECOND_FROM_TOP_LEFT = cellPositionFactory.GetInstanceColRow(0, 1)
            this.SECOND_FROM_TOP_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, 1)
        } else {
            this.SECOND_FROM_TOP_LEFT = cellPositionFactory.NONE
            this.SECOND_FROM_TOP_RIGHT = cellPositionFactory.NONE
        }
        if touchButtonLocationHelper.GetTotalRows() >= 4 {
            this.THIRD_FROM_BOTTOM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, touchButtonLocationHelper.GetTotalRows() - 3)
            this.FOURTH_FROM_BOTTOM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 1, touchButtonLocationHelper.GetTotalRows() - 4)
            this.THIRD_FROM_BOTTOM_LEFT = cellPositionFactory.GetInstanceColRow(0, touchButtonLocationHelper.GetTotalRows() - 3)
        } else {
            this.THIRD_FROM_BOTTOM_RIGHT = cellPositionFactory.NONE
            this.FOURTH_FROM_BOTTOM_RIGHT = cellPositionFactory.NONE
            this.THIRD_FROM_BOTTOM_LEFT = cellPositionFactory.NONE
        }
        if touchButtonLocationHelper.GetTotalColumns() >= 4 {
            this.SECOND_FROM_BOTTOM_THIRD_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 3, touchButtonLocationHelper.GetTotalRows() - 2)
            this.SECOND_FROM_BOTTOM_FOURTH_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 4, touchButtonLocationHelper.GetTotalRows() - 2)
            this.BOTTOM_THIRD_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 3, touchButtonLocationHelper.GetTotalRows() - 1)
            this.BOTTOM_FOURTH_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 4, touchButtonLocationHelper.GetTotalRows() - 1)
            this.TOP_SECOND_FROM_LEFT = cellPositionFactory.GetInstanceColRow(1, 0)
            this.TOP_SECOND_FROM_RIGHT = cellPositionFactory.GetInstanceColRow(touchButtonLocationHelper.GetTotalColumns() - 2, 0)
        } else {
            this.SECOND_FROM_BOTTOM_THIRD_FROM_RIGHT = cellPositionFactory.NONE
            this.SECOND_FROM_BOTTOM_FOURTH_FROM_RIGHT = cellPositionFactory.NONE
            this.BOTTOM_THIRD_FROM_RIGHT = cellPositionFactory.NONE
            this.BOTTOM_FOURTH_FROM_RIGHT = cellPositionFactory.NONE
            this.TOP_SECOND_FROM_LEFT = cellPositionFactory.NONE
            this.TOP_SECOND_FROM_RIGHT = cellPositionFactory.NONE
        }
    return this
}

