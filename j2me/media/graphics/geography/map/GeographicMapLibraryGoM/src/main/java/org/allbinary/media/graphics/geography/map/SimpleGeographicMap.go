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

type SimpleGeographicMap struct {
    CommonStrings *CommonStrings
    Id int
    Name string
    CellTypeIdToGeographicMapCellType []int
    TiledLayer IAllBinaryTiledLayer
    ForegroundBasicColor *BasicColor
    BackgroundBasicColor *BasicColor
    GeographicMapCellTypeFactory IGeographicMapCellTypeFactory
}

// The Java class SimpleGeographicMap has subclasses so its variables are this interface
type ISimpleGeographicMap interface {
    SimpleGeographicMap_() *SimpleGeographicMap
    GetId() int
    GetName() string
    Reset()
    GetAllBinaryTiledLayer() IAllBinaryTiledLayer
    GetCellTypeAt(cellPosition IGeographicMapCellPosition) IGeographicMapCellType
    GetCellTypeFromMapCellTypeInt(cellTypeId int) int
    GetForegroundBasicColor() *BasicColor
    SetForegroundBasicColor(foregroundBasicColor *BasicColor)
    GetBackgroundBasicColor() *BasicColor
    SetBackgroundBasicColor(backgroundBasicColor *BasicColor)
}

func (this *SimpleGeographicMap) SimpleGeographicMap_() *SimpleGeographicMap {
    return this
}


func (this *SimpleGeographicMap) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    this.ForegroundBasicColor = BasicColorFactory_GetInstance().NULL_COLOR
    this.BackgroundBasicColor = BasicColorFactory_GetInstance().NULL_COLOR
    this.GeographicMapCellTypeFactory = GeographicMapCellTypeFactory_GetInstance()
}


func NewSimpleGeographicMap(id int, name string, cellTypeIdToGeographicMapCellType []int, tiledLayer IAllBinaryTiledLayer, foregroundBasicColor *BasicColor, backgroundBasicColor *BasicColor) *SimpleGeographicMap {
    this := &SimpleGeographicMap{}
    this.initFields()
        this.Id = id
        this.Name = name
        this.CellTypeIdToGeographicMapCellType = cellTypeIdToGeographicMapCellType
        this.TiledLayer = tiledLayer
        this.SetBackgroundBasicColor(backgroundBasicColor)
        this.SetForegroundBasicColor(foregroundBasicColor)
    return this
}


func (this *SimpleGeographicMap) GetId() int {
        return this.Id
}


func (this *SimpleGeographicMap) GetName() string {
        return this.Name
}


func (this *SimpleGeographicMap) Reset() {
}


func (this *SimpleGeographicMap) GetAllBinaryTiledLayer() IAllBinaryTiledLayer {
        return this.TiledLayer
}


// throws Exception
func (this *SimpleGeographicMap) GetCellTypeAt(cellPosition IGeographicMapCellPosition) IGeographicMapCellType {
        var i_column int = cellPosition.GetColumn(); _ = i_column
        var i_row int = cellPosition.GetRow(); _ = i_row
        var cellTypeId int = this.TiledLayer.GetCell(i_column, i_row); _ = cellTypeId
        if cellTypeId < 0 {
            cellTypeId = this.TiledLayer.GetAnimatedTile(cellTypeId)
        }
        return this.GeographicMapCellTypeFactory.Get(this.CellTypeIdToGeographicMapCellType[cellTypeId])
}


func (this *SimpleGeographicMap) GetCellTypeFromMapCellTypeInt(cellTypeId int) int {
        return this.CellTypeIdToGeographicMapCellType[cellTypeId]
}


func (this *SimpleGeographicMap) GetForegroundBasicColor() *BasicColor {
        return this.ForegroundBasicColor
}


func (this *SimpleGeographicMap) SetForegroundBasicColor(foregroundBasicColor *BasicColor) {
        this.ForegroundBasicColor = foregroundBasicColor
}


func (this *SimpleGeographicMap) GetBackgroundBasicColor() *BasicColor {
        return this.BackgroundBasicColor
}


func (this *SimpleGeographicMap) SetBackgroundBasicColor(backgroundBasicColor *BasicColor) {
        this.BackgroundBasicColor = backgroundBasicColor
}

