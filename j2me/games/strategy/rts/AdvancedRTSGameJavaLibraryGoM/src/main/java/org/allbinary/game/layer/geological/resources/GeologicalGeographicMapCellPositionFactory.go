/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type GeologicalGeographicMapCellPositionFactory struct {
    ResourcePerMapCellType []int
}

// Implements GeographicMapCellPositionFactoryInterface


func (this *GeologicalGeographicMapCellPositionFactory) initFields() {
    
}


func NewGeologicalGeographicMapCellPositionFactory(resourcePerMapCellType []int) *GeologicalGeographicMapCellPositionFactory {
    this := &GeologicalGeographicMapCellPositionFactory{}
    this.initFields()
        this.ResourcePerMapCellType = resourcePerMapCellType
    return this
}


// throws Exception
func (this *GeologicalGeographicMapCellPositionFactory) GetInstance(geographicMapInterface IBasicGeographicMap, i_column int, i_row int, columns int, rows int, width int, height int) IGeographicMapCellPosition {
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        var resources int = this.ResourcePerMapCellType[tiledLayer.GetCell(i_column, i_row)] + MyRandomFactory_GetInstance().GetAbsoluteNextInt(1000); _ = resources
        return NewGeologicalGeographicMapCellPosition(i_column, i_row, columns, rows, width, height, NewGeologicalResource(resources))
}

