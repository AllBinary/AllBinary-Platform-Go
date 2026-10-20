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

type CellPositionsUtil struct {
    ReusableSingleThreadedSurroundingList IBasicArrayList
    SurroundArray [][]int
}

var CellPositionsUtil_Instance *CellPositionsUtil = NewCellPositionsUtil()

func (this *CellPositionsUtil) initFields() {
    this.ReusableSingleThreadedSurroundingList = NewBasicArrayListS(8)
    this.SurroundArray = [][]int{[]int{-1, -1}, []int{0, -1}, []int{1, -1}, []int{-1, 1}, []int{0, 1}, []int{1, 1}, []int{1, 0}, []int{-1, 0}}
}

func NewCellPositionsUtil() *CellPositionsUtil {
    this := &CellPositionsUtil{}
    this.initFields()
    return this
}


func CellPositionsUtil_GetInstance() *CellPositionsUtil {
        return CellPositionsUtil_Instance
}


// throws Exception
func (this *CellPositionsUtil) GetAll(geographicMapInterface IBasicGeographicMap, topRightGeographicMapCellPosition IGeographicMapCellPosition, columns int, rows int, reusableList IBasicArrayList) IBasicArrayList {
        reusableList.Clear()
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var lastColumn int = topRightGeographicMapCellPosition.GetColumn() + columns; _ = lastColumn
        var lastRow int = topRightGeographicMapCellPosition.GetRow() + rows; _ = lastRow
        if (columns > 1 && lastColumn > geographicMapInterface.GetAllBinaryTiledLayer().GetColumns()) || (rows > 1 && lastRow > geographicMapInterface.GetAllBinaryTiledLayer().GetRows()) {
            return reusableList
        }
        for rowIndex := int(0); rowIndex < rows; rowIndex++ {
            for columnIndex := int(0); columnIndex < columns; columnIndex++ {
                var geographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(topRightGeographicMapCellPosition.GetColumn() + columnIndex, topRightGeographicMapCellPosition.GetRow() + rowIndex); _ = geographicMapCellPosition
                reusableList.Add(geographicMapCellPosition)
            }
        }
        return reusableList
}


// throws Exception
func (this *CellPositionsUtil) GetAllSurrounding(geographicMapInterface IBasicGeographicMap, occupyList IBasicArrayList, reusableList IBasicArrayList) IBasicArrayList {
        reusableList.Clear()
        for index := int(occupyList.Size() - 1); index >= 0; index-- {
            var layerGeographicMapCellPosition IGeographicMapCellPosition = any(occupyList.Get(index)).(IGeographicMapCellPosition); _ = layerGeographicMapCellPosition
            var surroundingGeographicMapCellPositionList IBasicArrayList = this.GetAllSurroundingAt(geographicMapInterface, layerGeographicMapCellPosition, this.ReusableSingleThreadedSurroundingList); _ = surroundingGeographicMapCellPositionList
            for index2 := int(surroundingGeographicMapCellPositionList.Size() - 1); index2 >= 0; index2-- {
                var geographicMapCellPosition IGeographicMapCellPosition = any(surroundingGeographicMapCellPositionList.Get(index2)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
                if !reusableList.Contains(geographicMapCellPosition) && !occupyList.Contains(geographicMapCellPosition) {
                    reusableList.Add(geographicMapCellPosition)
                }
            }
        }
        return reusableList
}


// throws Exception
func (this *CellPositionsUtil) GetAllSurroundingAt(geographicMapInterface IBasicGeographicMap, layerGeographicMapCellPosition IGeographicMapCellPosition, reusableSurroundingList IBasicArrayList) IBasicArrayList {
        reusableSurroundingList.Clear()
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        var row int; _ = row
        var column int; _ = column
        for index := int(0); index < 8; index++ {
            column = layerGeographicMapCellPosition.GetColumn() + this.SurroundArray[index][0]
            row = layerGeographicMapCellPosition.GetRow() + this.SurroundArray[index][1]
            if tiledLayer.IsOnTileLayer(column, row) {
                var geographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(column, row); _ = geographicMapCellPosition
                reusableSurroundingList.Add(geographicMapCellPosition)
            }
        }
        return reusableSurroundingList
}

