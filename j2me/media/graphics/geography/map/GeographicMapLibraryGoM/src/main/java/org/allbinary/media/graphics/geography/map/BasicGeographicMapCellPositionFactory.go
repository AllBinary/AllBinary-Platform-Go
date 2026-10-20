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

type BasicGeographicMapCellPositionFactory struct {
    LogUtil *LogUtil
    GeographicMapCellPositionArray [][]IGeographicMapCellPosition
    GeographicMapInterface IBasicGeographicMap
    GeographicMapCellPositionFactoryInterface GeographicMapCellPositionFactoryInterface
    Columns int
    Rows int
    TiledLayer IAllBinaryTiledLayer
}

// The Java class BasicGeographicMapCellPositionFactory has subclasses so its variables are this interface
type IBasicGeographicMapCellPositionFactory interface {
    BasicGeographicMapCellPositionFactory_() *BasicGeographicMapCellPositionFactory
    Init()
    Visit(geographicMapCelPositionFactoryInitVisitorInterface GeographicMapCellPositionFactoryInitVisitorInterface)
    GetAt(i_column int, i_row int) IGeographicMapCellPosition
    GetInstance(anotherMapGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition
    CreateInstance(i_column int, i_row int, width int, height int) IGeographicMapCellPosition
    GetColumns() int
    GetRows() int
}

func (this *BasicGeographicMapCellPositionFactory) BasicGeographicMapCellPositionFactory_() *BasicGeographicMapCellPositionFactory {
    return this
}


func (this *BasicGeographicMapCellPositionFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}


func NewBasicGeographicMapCellPositionFactory(geographicMapInterface IBasicGeographicMap) *BasicGeographicMapCellPositionFactory {
    this := &BasicGeographicMapCellPositionFactory{}
    this.initFields()
        this.GeographicMapInterface = geographicMapInterface
        this.GeographicMapCellPositionFactoryInterface = this.GeographicMapInterface.GetGeographicMapCellPositionFactoryInterface()
        this.TiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer()
        this.Columns = this.TiledLayer.GetColumns()
        this.Rows = this.TiledLayer.GetRows()
        this.GeographicMapCellPositionArray = func() [][]IGeographicMapCellPosition { array := make([][]IGeographicMapCellPosition, this.Rows); for index := range array { array[index] = make([]IGeographicMapCellPosition, this.Columns) }; return array }()
        this.Init()
    return this
}


// throws Exception
func (this *BasicGeographicMapCellPositionFactory) Init() {
        var tiledLayer IAllBinaryTiledLayer = this.TiledLayer; _ = tiledLayer
        var columns int = this.GetColumns(); _ = columns
        var rows int = this.GetRows(); _ = rows
        var width int = tiledLayer.GetCellWidth(); _ = width
        var height int = tiledLayer.GetCellHeight(); _ = height
        for column := int(0); column < columns; column++ {
            for row := int(0); row < rows; row++ {
                this.CreateInstance(column, row, width, height)
            }
        }
}


// throws Exception
func (this *BasicGeographicMapCellPositionFactory) Visit(geographicMapCelPositionFactoryInitVisitorInterface GeographicMapCellPositionFactoryInitVisitorInterface) {
        var tiledLayer IAllBinaryTiledLayer = this.TiledLayer; _ = tiledLayer
        var columns int = this.GetColumns(); _ = columns
        var rows int = this.GetRows(); _ = rows
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append(commonSeps.BRACKET_OPEN).Appendint(rows).Append(commonSeps.BRACKET_CLOSE).Append(commonSeps.BRACKET_OPEN).Appendint(columns).Append(commonSeps.BRACKET_CLOSE)), this, "visit", e)
            panic(e)
                }
            }()
            var cellPosition IGeographicMapCellPosition; _ = cellPosition
            for column := int(0); column < columns; column++ {
                for row := int(0); row < rows; row++ {
                    cellPosition = this.GetAt(column, row)
                    geographicMapCelPositionFactoryInitVisitorInterface.Visit(tiledLayer, cellPosition)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *BasicGeographicMapCellPositionFactory) GetAt(i_column int, i_row int) IGeographicMapCellPosition {
        return this.GeographicMapCellPositionArray[i_row][i_column]
}


// throws Exception
func (this *BasicGeographicMapCellPositionFactory) GetInstance(anotherMapGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        return this.GetAt(anotherMapGeographicMapCellPosition.GetColumn(), anotherMapGeographicMapCellPosition.GetRow())
}


// throws Exception
func (this *BasicGeographicMapCellPositionFactory) CreateInstance(i_column int, i_row int, width int, height int) IGeographicMapCellPosition {
        var cellPositionCanBeNull any = this.GeographicMapCellPositionArray[i_row][i_column]; _ = cellPositionCanBeNull
        if cellPositionCanBeNull == nil {
            cellPositionCanBeNull = this.GeographicMapCellPositionFactoryInterface.GetInstance(this.GeographicMapInterface, i_column, i_row, this.Columns, this.Rows, width, height)
            this.GeographicMapCellPositionArray[i_row][i_column] = any(cellPositionCanBeNull).(IGeographicMapCellPosition)
        }
        return any(cellPositionCanBeNull).(IGeographicMapCellPosition)
}


func (this *BasicGeographicMapCellPositionFactory) GetColumns() int {
        return this.Columns
}


func (this *BasicGeographicMapCellPositionFactory) GetRows() int {
        return this.Rows
}

