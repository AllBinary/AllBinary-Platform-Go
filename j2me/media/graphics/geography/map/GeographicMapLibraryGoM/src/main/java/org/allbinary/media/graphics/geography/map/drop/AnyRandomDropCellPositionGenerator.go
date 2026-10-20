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

type AnyRandomDropCellPositionGenerator struct {
    BaseDropCellPositionGenerator
    List IBasicArrayList
    GeographicMapInterface any
}


func (this *AnyRandomDropCellPositionGenerator) initFields() {
    this.List = NewBasicArrayListD()
    this.GeographicMapInterface = NullUtil_GetInstance().NULL_OBJECT
}

func NewAnyRandomDropCellPositionGenerator() *AnyRandomDropCellPositionGenerator {
    this := &AnyRandomDropCellPositionGenerator{}
    this.BaseDropCellPositionGenerator = *NewBaseDropCellPositionGenerator()
    this.initFields()
    return this
}


// throws Exception
func (this *AnyRandomDropCellPositionGenerator) Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap) {
        this.GeographicMapInterface = geographicMapInterface
        var basicGeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = basicGeographicMapCellPositionFactory
        var randomFactory *MyRandomFactory = MyRandomFactory_GetInstance(); _ = randomFactory
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        var total int = tiledLayer.GetColumns() * tiledLayer.GetRows(); _ = total
        var randomColumn int; _ = randomColumn
        var randomRow int; _ = randomRow
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        for index := int(total); Java_PreDecrement(&index) >= 0;  {
            randomColumn = randomFactory.GetAbsoluteNextInt(tiledLayer.GetColumns())
            randomRow = randomFactory.GetAbsoluteNextInt(tiledLayer.GetRows())
            geographicMapCellPosition = basicGeographicMapCellPositionFactory.GetAt(randomColumn, randomRow)
            if this.IsDropAllowedAt(geographicMapCellPosition) {
                this.List.Add(geographicMapCellPosition)
            }
        }
}


// throws Exception
func (this *AnyRandomDropCellPositionGenerator) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
}

