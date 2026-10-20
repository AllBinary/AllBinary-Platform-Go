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

type TileLayerPositionIntoViewPosition struct {
    ViewPosition
    TiledLayer ILayer
}


func (this *TileLayerPositionIntoViewPosition) initFields() {
    this.TiledLayer = NullLayer_GetInstance()
}


func NewTileLayerPositionIntoViewPosition() *TileLayerPositionIntoViewPosition {
    this := &TileLayerPositionIntoViewPosition{}
        this.ViewPosition = *NewViewPosition(0, 0, 0)
        this.initFields()
    return this
}


func (this *TileLayerPositionIntoViewPosition) GetX() int {
        return this.ViewPosition.GetX() - this.TiledLayer.GetXP()
}


func (this *TileLayerPositionIntoViewPosition) GetY() int {
        return this.ViewPosition.GetY() - this.TiledLayer.GetYP()
}


func (this *TileLayerPositionIntoViewPosition) GetZ() int {
        return this.ViewPosition.GetZ() - this.TiledLayer.GetZP()
}


func (this *TileLayerPositionIntoViewPosition) SetTiledLayer(tiledLayer IAllBinaryTiledLayer) {
        this.TiledLayer = tiledLayer
}

