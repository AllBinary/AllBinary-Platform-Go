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

type CellPositionFactory struct {
    LogUtil *LogUtil
    NONE ICellPosition
    CellPositionArray [][]ICellPosition
    Columns int
    Rows int
}

var CellPositionFactory_SINGLETON *CellPositionFactory = NewCellPositionFactory()

func (this *CellPositionFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NONE = NewCellPosition(-1, -1, -1, -1)
    this.CellPositionArray = func() [][]ICellPosition { array := make([][]ICellPosition, 0); for index := range array { array[index] = make([]ICellPosition, 0) }; return array }()
    
    
}


func CellPositionFactory_GetInstance() *CellPositionFactory {
        return CellPositionFactory_SINGLETON
}


func NewCellPositionFactory() *CellPositionFactory {
    this := &CellPositionFactory{}
    this.initFields()
    return this
}


func (this *CellPositionFactory) Init(columns int, rows int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Init: columns: ")
        stringBuffer.Appendint(columns)
        stringBuffer.Append(" rows: ")
        stringBuffer.Appendint(rows)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, commonStrings.INIT)
        this.CellPositionArray = func() [][]ICellPosition { array := make([][]ICellPosition, columns); for index := range array { array[index] = make([]ICellPosition, rows) }; return array }()
        this.Columns = columns
        this.Rows = rows
        for column := int(0); column < columns; column++ {
            for row := int(0); row < rows; row++ {
                this.CreateInstance(column, row)
            }
        }
}


func (this *CellPositionFactory) GetInstanceColRow(i_column int, i_row int) ICellPosition {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 ICellPosition) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append("columns: ").Appendint(this.GetColumns()).Append(" rows: ").Appendint(this.GetRows()).Append(" col: ").Appendint(i_column).Append(" row: ").Appendint(i_row)), this, commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, this.NONE; return
                }
            }()
            tryControl1, tryResult1 = 1, this.CellPositionArray[i_column][i_row]; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CellPositionFactory) CreateInstance(i_column int, i_row int) ICellPosition {
        var cellPositionCanBeNull ICellPosition = this.CellPositionArray[i_column][i_row]; _ = cellPositionCanBeNull
        if cellPositionCanBeNull == nil {
            cellPositionCanBeNull = NewCellPosition(i_column, i_row, this.Columns, this.Rows)
            this.CellPositionArray[i_column][i_row] = cellPositionCanBeNull
        }
        return cellPositionCanBeNull
}


func (this *CellPositionFactory) GetColumns() int {
        return this.Columns
}


func (this *CellPositionFactory) GetRows() int {
        return this.Rows
}

