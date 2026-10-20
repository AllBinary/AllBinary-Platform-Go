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

type TileLayerItem struct {
    ABCustomItem
    TiledLayer IAllBinaryTiledLayer
}

// Implements ABCustomItemInterface


func (this *TileLayerItem) initFields() {
    
}


func NewTileLayerItem(label string, tiledLayer IAllBinaryTiledLayer, layout int, altText string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *TileLayerItem {
    this := &TileLayerItem{}
        this.ABCustomItem = *NewABCustomItem(label, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.TiledLayer = tiledLayer
    return this
}


func (this *TileLayerItem) SetOwner(owner IScreen) {
}


func (this *TileLayerItem) GetMinimumWidth() int {
        return this.TiledLayer.GetWidth()
}


func (this *TileLayerItem) GetMinimumHeight() int {
        return this.TiledLayer.GetHeight()
}


func (this *TileLayerItem) GetMinContentHeight() int {
        return this.TiledLayer.GetHeight()
}


func (this *TileLayerItem) GetMinContentWidth() int {
        return this.TiledLayer.GetWidth()
}


func (this *TileLayerItem) GetPrefContentHeight(width int) int {
        return this.TiledLayer.GetHeight()
}


func (this *TileLayerItem) GetPrefContentWidth(height int) int {
        return this.TiledLayer.GetWidth()
}


func (this *TileLayerItem) PaintXY(graphics IGraphics, x int, y int) {
        this.TiledLayer.SetPosition(x, y, this.TiledLayer.GetZP())
        this.TiledLayer.Paint(graphics)
}


func (this *TileLayerItem) PaintUnselected(graphics IGraphics, x int, y int) {
}

