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

type BasicGeographicMapUtil struct {
    LogUtil *LogUtil
}

var BasicGeographicMapUtil_Instance *BasicGeographicMapUtil = NewBasicGeographicMapUtil()

func (this *BasicGeographicMapUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewBasicGeographicMapUtil() *BasicGeographicMapUtil {
    this := &BasicGeographicMapUtil{}
    this.initFields()
    return this
}


func BasicGeographicMapUtil_GetInstance() *BasicGeographicMapUtil {
        return BasicGeographicMapUtil_Instance
}


// throws Exception
func (this *BasicGeographicMapUtil) GetBorderingRow(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) int {
        var LOCATION_CHANGE int = 1; _ = LOCATION_CHANGE
        switch direction {
            case 0:
                return oldGeographicMapCellPosition.GetRow()
            case 1:
                return oldGeographicMapCellPosition.GetRow()
            case 2:
                return oldGeographicMapCellPosition.GetRow() + LOCATION_CHANGE
            case 3:
                return oldGeographicMapCellPosition.GetRow() - LOCATION_CHANGE
            default:
                panic(errors.New(fmt.Sprint("Only Four Directions")))
        }
}


// throws Exception
func (this *BasicGeographicMapUtil) GetBorderingColumn(direction int, oldGeographicMapCellPosition IGeographicMapCellPosition) int {
        var LOCATION_CHANGE int = 1; _ = LOCATION_CHANGE
        switch direction {
            case 0:
                return oldGeographicMapCellPosition.GetColumn() - LOCATION_CHANGE
            case 1:
                return oldGeographicMapCellPosition.GetColumn() + LOCATION_CHANGE
            case 2:
                return oldGeographicMapCellPosition.GetColumn()
            case 3:
                return oldGeographicMapCellPosition.GetColumn()
            default:
                panic(errors.New(fmt.Sprint("Only Four Directions")))
        }
}


// throws Exception
func (this *BasicGeographicMapUtil) IsSameCellPosition(fromGeographicMapCellPosition IGeographicMapCellPosition, toGeographicMapCellPosition IGeographicMapCellPosition) bool {
        var fromColumn int = fromGeographicMapCellPosition.GetColumn(); _ = fromColumn
        var fromRow int = fromGeographicMapCellPosition.GetRow(); _ = fromRow
        var goColumn int = toGeographicMapCellPosition.GetColumn(); _ = goColumn
        var goRow int = toGeographicMapCellPosition.GetRow(); _ = goRow
        if fromColumn == goColumn && fromRow == goRow {
            return true
        } else {
            return false
        }
}


func (this *BasicGeographicMapUtil) CreateAllBinaryTiledLayerArray(geographicMapInterfaceArray []IBasicGeographicMap) []ILayer {
        var tiledLayerArray []ILayer = make([]ILayer, len(geographicMapInterfaceArray)); _ = tiledLayerArray
        return this.GetAllBinaryTiledLayerArray(geographicMapInterfaceArray, tiledLayerArray, 0)
}


func (this *BasicGeographicMapUtil) GetAllBinaryTiledLayerArray(geographicMapInterfaceArray []IBasicGeographicMap, tiledLayerArray []ILayer, startIndex int) []ILayer {
        var geographicMapInterface IBasicGeographicMap; _ = geographicMapInterface
        var size int = len(geographicMapInterfaceArray); _ = size
        var count int = 0; _ = count
        for index := int(size + startIndex); Java_PreDecrement(&index) >= startIndex;  {
            geographicMapInterface = geographicMapInterfaceArray[index]
            tiledLayerArray[Java_PostIncrement(&count)] = geographicMapInterface.GetAllBinaryTiledLayer()
        }
        return tiledLayerArray
}


func (this *BasicGeographicMapUtil) Move(geographicMapInterfaceArray []IBasicGeographicMap, dx int, dy int) {
        var geographicMapInterface IBasicGeographicMap; _ = geographicMapInterface
        var size int = len(geographicMapInterfaceArray); _ = size
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            geographicMapInterface = geographicMapInterfaceArray[index]
            geographicMapInterface.GetAllBinaryTiledLayer().MoveDXY(dx, dy)
        }
}


func (this *BasicGeographicMapUtil) SetPosition(geographicMapInterfaceArray []IBasicGeographicMap, x int, y int) {
        var geographicMapInterface IBasicGeographicMap; _ = geographicMapInterface
        var size int = len(geographicMapInterfaceArray); _ = size
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            geographicMapInterface = geographicMapInterfaceArray[index]
            geographicMapInterface.GetAllBinaryTiledLayer().SetPosition(x, y, geographicMapInterface.GetAllBinaryTiledLayer().GetZP())
        }
}

