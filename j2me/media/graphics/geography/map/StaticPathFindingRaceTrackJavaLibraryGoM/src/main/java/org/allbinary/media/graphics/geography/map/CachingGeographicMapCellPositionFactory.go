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

type CachingGeographicMapCellPositionFactory struct {
    BasicGeographicMapCellPositionFactory
}



func NewCachingGeographicMapCellPositionFactory(geographicMapInterface IBasicGeographicMap) *CachingGeographicMapCellPositionFactory {
    this := &CachingGeographicMapCellPositionFactory{}
        this.BasicGeographicMapCellPositionFactory = *NewBasicGeographicMapCellPositionFactory(geographicMapInterface)
    return this
}


// throws Exception
func (this *CachingGeographicMapCellPositionFactory) CreateInstance(i_column int, i_row int, width int, height int) IGeographicMapCellPosition {
        var hashtable *ABHashtable[any, any] = GeographicMapCellPositionFactory_GetHashtable(); _ = hashtable
        var cellPositionKey string = CellPosition_ToStringColRow(i_column, i_row); _ = cellPositionKey
        var cellPositionCanBeNull any = hashtable.Get(cellPositionKey); _ = cellPositionCanBeNull
        if cellPositionCanBeNull == nil {
            cellPositionCanBeNull = this.GeographicMapCellPositionFactoryInterface.GetInstance(this.GeographicMapInterface, i_column, i_row, this.GetColumns(), this.GetRows(), width, height)
            this.GeographicMapCellPositionArray[i_row][i_column] = any(cellPositionCanBeNull).(IGeographicMapCellPosition)
            hashtable.Put(cellPositionKey, cellPositionCanBeNull)
        }
        return any(cellPositionCanBeNull).(IGeographicMapCellPosition)
}

