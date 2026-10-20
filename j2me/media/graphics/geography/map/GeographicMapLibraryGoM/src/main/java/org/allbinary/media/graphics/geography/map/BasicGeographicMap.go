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

type BasicGeographicMap struct {
    SimpleGeographicMap
    MathUtil *MathUtil
    GeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory
    GeographicMapCellPositionFactoryInterface GeographicMapCellPositionFactoryInterface
    GeographicMapCellTypeFactory IGeographicMapCellTypeFactory
}

// Implements GeographicMapInterface

// The Java class BasicGeographicMap has subclasses so its variables are this interface
type IBasicGeographicMap interface {
    ISimpleGeographicMap
    BasicGeographicMap_() *BasicGeographicMap
    GetCellPosition(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition
    GetCellPositionNoThrow(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition
    IsOfFourDirections(oldGeographicMapCellPosition IGeographicMapCellPosition, newGeographicMapCellPosition IGeographicMapCellPosition) bool
    GetCellPositionAtXY(x int, y int) IGeographicMapCellPosition
    GetCellPositionAtXYNoThrow(x int, y int) IGeographicMapCellPosition
    IsOnMap(geographicMapCellPosition IGeographicMapCellPosition) bool
    GetCellPositionAtNoThrow(x int, y int, x2 int, y2 int, geographicMapCellPositionList IBasicArrayList) IBasicArrayList
    GetCellPositionsAt(layer ILayer, currentCellPositionArray [][]IGeographicMapCellPosition, cellPositionArray [][]IGeographicMapCellPosition) bool
    GetCellTypeAtXY(x int, y int) IGeographicMapCellType
    GetGeographicMapCellPositionFactory() IBasicGeographicMapCellPositionFactory
    GetGeographicMapCellPositionFactoryInterface() GeographicMapCellPositionFactoryInterface
    GetGeographicMapCellTypeFactory() IGeographicMapCellTypeFactory
}

func (this *BasicGeographicMap) BasicGeographicMap_() *BasicGeographicMap {
    return this
}

var BasicGeographicMap_NULL_BASIC_GEOGRAPHIC_MAP_ARRAY []IBasicGeographicMap = make([]IBasicGeographicMap, 0)

func (this *BasicGeographicMap) initFields() {
    this.MathUtil = MathUtil_GetInstance()
    
    
    
}


func NewBasicGeographicMap(id int, name string, cellTypeIdToGeographicMapCellType []int, tiledLayer IAllBinaryTiledLayer, foregroundBasicColor *BasicColor, backgroundBasicColor *BasicColor, geographicMapCellPositionFactoryInterface GeographicMapCellPositionFactoryInterface, geographicMapCellPositionBaseFactory IGeographicMapCellPositionBaseFactory, geographicMapCellTypeFactory IGeographicMapCellTypeFactory) *BasicGeographicMap {
    this := &BasicGeographicMap{}
        this.SimpleGeographicMap = *NewSimpleGeographicMap(id, name, cellTypeIdToGeographicMapCellType, tiledLayer, foregroundBasicColor, backgroundBasicColor)
        this.initFields()
        this.GeographicMapCellPositionFactoryInterface = geographicMapCellPositionFactoryInterface
        this.GeographicMapCellPositionFactory = geographicMapCellPositionBaseFactory.GetInstance(this)
        this.GeographicMapCellTypeFactory = geographicMapCellTypeFactory
    return this
}


// throws Exception
func (this *BasicGeographicMap) GetCellPosition(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        switch direction {
            case 0:
                return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn() - 1, oldGeographicMapCellPosition.GetRow())
            case 1:
                return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn() + 1, oldGeographicMapCellPosition.GetRow())
            case 2:
                return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn(), oldGeographicMapCellPosition.GetRow() - 1)
            case 3:
                return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn(), oldGeographicMapCellPosition.GetRow() + 1)
            default:
                panic(errors.New(fmt.Sprint("Only Four Directions")))
        }
}


// throws Exception
func (this *BasicGeographicMap) GetCellPositionNoThrow(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellPosition {
        switch direction {
            case 0:
                if oldGeographicMapCellPosition.GetColumn() - 1 >= 0 {
                    return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn() - 1, oldGeographicMapCellPosition.GetRow())
                } else {
                    return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                }
            fallthrough
            case 1:
                if oldGeographicMapCellPosition.GetColumn() + 1 < this.GetAllBinaryTiledLayer().GetColumns() {
                    return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn() + 1, oldGeographicMapCellPosition.GetRow())
                } else {
                    return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                }
            fallthrough
            case 2:
                if oldGeographicMapCellPosition.GetRow() - 1 >= 0 {
                    return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn(), oldGeographicMapCellPosition.GetRow() - 1)
                } else {
                    return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                }
            fallthrough
            case 3:
                if oldGeographicMapCellPosition.GetRow() + 1 < this.GetAllBinaryTiledLayer().GetRows() {
                    return this.GeographicMapCellPositionFactory.GetAt(oldGeographicMapCellPosition.GetColumn(), oldGeographicMapCellPosition.GetRow() + 1)
                } else {
                    return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
                }
            fallthrough
            default:
                return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
        }
}


// throws Exception
func (this *BasicGeographicMap) IsOfFourDirections(oldGeographicMapCellPosition IGeographicMapCellPosition, newGeographicMapCellPosition IGeographicMapCellPosition) bool {
        for index := int(0); index < 4; index++ {
            if newGeographicMapCellPosition == this.GetCellPositionNoThrow(index, oldGeographicMapCellPosition) {
                return true
            }
        }
        return false
}


// throws Exception
func (this *BasicGeographicMap) GetCellPositionAtXY(x int, y int) IGeographicMapCellPosition {
        var allBinaryTiledLayer IAllBinaryTiledLayer = this.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var i_column int = this.MathUtil.Abs(x / allBinaryTiledLayer.GetCellHeight()); _ = i_column
        var i_row int = this.MathUtil.Abs(y / allBinaryTiledLayer.GetCellWidth()); _ = i_row
        return this.GeographicMapCellPositionFactory.GetAt(i_column, i_row)
}


func (this *BasicGeographicMap) GetCellPositionAtXYNoThrow(x int, y int) IGeographicMapCellPosition {
        var allBinaryTiledLayer IAllBinaryTiledLayer = this.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var i_column int = this.MathUtil.Abs(x / allBinaryTiledLayer.GetCellHeight()); _ = i_column
        var i_row int = this.MathUtil.Abs(y / allBinaryTiledLayer.GetCellWidth()); _ = i_row
        if allBinaryTiledLayer.GetColumns() > i_column && allBinaryTiledLayer.GetRows() > i_row {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IGeographicMapCellPosition) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                tryControl1, tryResult1 = 1, SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION; return
                    }
                }()
                tryControl1, tryResult1 = 1, this.GeographicMapCellPositionFactory.GetAt(i_column, i_row); return
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        } else {
            return SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
        }
}


// throws Exception
func (this *BasicGeographicMap) IsOnMap(geographicMapCellPosition IGeographicMapCellPosition) bool {
        var allBinaryTiledLayer IAllBinaryTiledLayer = this.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var i_column int = geographicMapCellPosition.GetColumn(); _ = i_column
        var i_row int = geographicMapCellPosition.GetRow(); _ = i_row
        if allBinaryTiledLayer.GetColumns() > i_column && allBinaryTiledLayer.GetRows() > i_row {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *BasicGeographicMap) GetCellPositionAtNoThrow(x int, y int, x2 int, y2 int, geographicMapCellPositionList IBasicArrayList) IBasicArrayList {
        geographicMapCellPositionList.Clear()
        var allBinaryTiledLayer IAllBinaryTiledLayer = this.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var i_columnMin int = this.MathUtil.Abs(x / allBinaryTiledLayer.GetCellHeight()); _ = i_columnMin
        var i_rowMin int = this.MathUtil.Abs(y / allBinaryTiledLayer.GetCellWidth()); _ = i_rowMin
        var i_columnMax int = this.MathUtil.Abs(x2 / allBinaryTiledLayer.GetCellHeight()) + 1; _ = i_columnMax
        var i_rowMax int = this.MathUtil.Abs(y2 / allBinaryTiledLayer.GetCellWidth()) + 1; _ = i_rowMax
        for columnIndex := int(i_columnMin); columnIndex < i_columnMax; columnIndex++ {
            for rowIndex := int(i_rowMin); rowIndex < i_rowMax; rowIndex++ {
                if allBinaryTiledLayer.GetColumns() > columnIndex && allBinaryTiledLayer.GetRows() > rowIndex {
                    geographicMapCellPositionList.Add(this.GeographicMapCellPositionFactory.GetAt(columnIndex, rowIndex))
                }
            }
        }
        return geographicMapCellPositionList
}


// throws Exception
func (this *BasicGeographicMap) GetCellPositionsAt(layer ILayer, currentCellPositionArray [][]IGeographicMapCellPosition, cellPositionArray [][]IGeographicMapCellPosition) bool {
        var hasChanged bool = false; _ = hasChanged
        var size int = len(cellPositionArray); _ = size
        var size2 int = len(cellPositionArray[0]); _ = size2
        var xPortion int = layer.GetXP() / (size - 1); _ = xPortion
        var yPortion int = layer.GetYP() / (size - 1); _ = yPortion
        for index := int(0); index < size; index++ {
            for index2 := int(0); index2 < size2; index2++ {
                var x int = xPortion * index; _ = x
                var y int = yPortion * index; _ = y
                cellPositionArray[index][index2] = this.GetCellPositionAtXY(x, y)
                if currentCellPositionArray[index][index2] != cellPositionArray[index][index2] {
                    hasChanged = true
                }
            }
        }
        return hasChanged
}


// throws Exception
func (this *BasicGeographicMap) GetCellTypeAtXY(x int, y int) IGeographicMapCellType {
        var cellPosition IGeographicMapCellPosition = this.GetCellPositionAtXY(x, y); _ = cellPosition
        return this.GetCellTypeAt(cellPosition)
}


func (this *BasicGeographicMap) GetGeographicMapCellPositionFactory() IBasicGeographicMapCellPositionFactory {
        return this.GeographicMapCellPositionFactory
}


func (this *BasicGeographicMap) GetGeographicMapCellPositionFactoryInterface() GeographicMapCellPositionFactoryInterface {
        return this.GeographicMapCellPositionFactoryInterface
}


func (this *BasicGeographicMap) GetGeographicMapCellTypeFactory() IGeographicMapCellTypeFactory {
        return this.GeographicMapCellTypeFactory
}

