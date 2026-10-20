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

type LayerCoveringCellPositionsUtil struct {
    CellPositionsUtil *CellPositionsUtil
}

var LayerCoveringCellPositionsUtil_Instance *LayerCoveringCellPositionsUtil = NewLayerCoveringCellPositionsUtil()

func (this *LayerCoveringCellPositionsUtil) initFields() {
    this.CellPositionsUtil = CellPositionsUtil_GetInstance()
}

func NewLayerCoveringCellPositionsUtil() *LayerCoveringCellPositionsUtil {
    this := &LayerCoveringCellPositionsUtil{}
    this.initFields()
    return this
}


func LayerCoveringCellPositionsUtil_GetInstance() *LayerCoveringCellPositionsUtil {
        return LayerCoveringCellPositionsUtil_Instance
}


// throws Exception
func (this *LayerCoveringCellPositionsUtil) GetAllXY(geographicMapInterface IBasicGeographicMap, layerInterface IAllBinaryLayer, x int, y int, reusableList IBasicArrayList) IBasicArrayList {
        var topLeftGeographicMapCellPosition IGeographicMapCellPosition = geographicMapInterface.GetCellPositionAtXY(x, y); _ = topLeftGeographicMapCellPosition
        return this.GetAll(geographicMapInterface, topLeftGeographicMapCellPosition, layerInterface, reusableList)
}


// throws Exception
func (this *LayerCoveringCellPositionsUtil) GetAll(geographicMapInterface IBasicGeographicMap, topLeftGeographicMapCellPosition IGeographicMapCellPosition, layerInterface IAllBinaryLayer, reusableList IBasicArrayList) IBasicArrayList {
        reusableList.Clear()
        var columns int = layerInterface.GetWidth() / geographicMapInterface.GetAllBinaryTiledLayer().GetCellWidth(); _ = columns
        var rows int = layerInterface.GetHeight() / geographicMapInterface.GetAllBinaryTiledLayer().GetCellHeight(); _ = rows
        if columns == 0 {
            columns = 1
        }
        if rows == 0 {
            rows = 1
        }
        return this.CellPositionsUtil.GetAll(geographicMapInterface, topLeftGeographicMapCellPosition, columns, rows, reusableList)
}

