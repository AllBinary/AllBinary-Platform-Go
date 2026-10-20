/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type AllBinaryJ2METiledLayer struct {
    AllBinaryTiledLayer
    TiledLayerP *TiledLayer
    DebugColor int
    Font2 Font
}


func (this *AllBinaryJ2METiledLayer) initFields() {
    
    
    this.Font2 = Font_GetFont(Font_FACE_SYSTEM, Font_STYLE_PLAIN, 4)
}


func NewAllBinaryJ2METiledLayer(dataId int, tiledLayer *TiledLayer, mapTwoDArray [][]int, debugColor int) *AllBinaryJ2METiledLayer {
    this := &AllBinaryJ2METiledLayer{}
        this.AllBinaryTiledLayer = *NewAllBinaryTiledLayer(dataId, tiledLayer.GetWidth(), tiledLayer.GetHeight(), tiledLayer.GetCellWidth(), tiledLayer.GetCellHeight())
        this.initFields()
        this.TiledLayerP = tiledLayer
        this.DebugColor = debugColor
        this.SetCells(mapTwoDArray)
    return this
}


func (this *AllBinaryJ2METiledLayer) Paint(graphics IGraphics) {
        this.TiledLayerP.Paint(graphics)
}


func (this *AllBinaryJ2METiledLayer) PaintDebug(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        graphics.SetFont(this.Font2)
        graphics.SetColor(this.DebugColor)
        var x int = this.TiledLayerP.GetX(); _ = x
        var y int = this.TiledLayerP.GetY(); _ = y
        var firstColumn int = 0; _ = firstColumn
        var firstRow int = 0; _ = firstRow
        var totalColumns int = this.TiledLayerP.GetColumns(); _ = totalColumns
        var totalRows int = this.TiledLayerP.GetRows(); _ = totalRows
        var x0 int = x; _ = x0
        var tile int = 0; _ = tile
        var tileHeight int = this.TiledLayerP.GetCellHeight(); _ = tileHeight
        var tileWidth int = this.TiledLayerP.GetCellWidth(); _ = tileWidth
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        for rowIndex := int(firstRow); rowIndex < totalRows; rowIndex++, y += tileHeight {
            x = x0
            for columnIndex := int(firstColumn); columnIndex < totalColumns; columnIndex++, x += tileWidth {
                tile = this.TiledLayerP.GetCell(columnIndex, rowIndex)
                if tile < 0 {
                    tile = this.TiledLayerP.GetAnimatedTile(tile)
                }
                if tile == 0 {
                    continue
                }
                tile--
                stringMaker.Delete(0, stringMaker.Length())
                graphics.DrawString(fmt.Sprint(stringMaker.Appendint(columnIndex).Append(commonSeps.COMMA).Appendint(rowIndex).Append(commonSeps.COLON).Appendint(tile)), x + 5, y + 10, 0)
            }
        }
        graphics.SetFont(font)
}


func (this *AllBinaryJ2METiledLayer) MoveDXY(dx int, dy int) {
        this.TiledLayerP.Move(dx, dy)
        this.AllBinaryTiledLayer.MoveDXY(-dx, -dy)
}


func (this *AllBinaryJ2METiledLayer) SetPosition(x int, y int, z int) {
        this.TiledLayerP.SetPosition(x, y)
        this.AllBinaryTiledLayer.SetPosition(-x, -y, z)
}


func (this *AllBinaryJ2METiledLayer) GetTiledLayer() *TiledLayer {
        return this.TiledLayerP
}


func (this *AllBinaryJ2METiledLayer) GetCellWidth() int {
        return this.TiledLayerP.GetCellWidth()
}


func (this *AllBinaryJ2METiledLayer) GetCellHeight() int {
        return this.TiledLayerP.GetCellHeight()
}


func (this *AllBinaryJ2METiledLayer) GetCell(col int, row int) int {
        return this.TiledLayerP.GetCell(col, row)
}


func (this *AllBinaryJ2METiledLayer) GetColumns() int {
        return this.TiledLayerP.GetColumns()
}


func (this *AllBinaryJ2METiledLayer) GetRows() int {
        return this.TiledLayerP.GetRows()
}


func (this *AllBinaryJ2METiledLayer) SetCell(col int, row int, index int) {
        this.TiledLayerP.SetCell(col, row, index)
}


func (this *AllBinaryJ2METiledLayer) GetAnimatedTile(animationTileIndex int) int {
        return this.TiledLayerP.GetAnimatedTile(animationTileIndex)
}

