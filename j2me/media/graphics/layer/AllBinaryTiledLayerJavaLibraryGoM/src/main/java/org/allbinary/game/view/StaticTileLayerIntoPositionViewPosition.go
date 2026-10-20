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

type StaticTileLayerIntoPositionViewPosition struct {
    ViewPosition
}

var StaticTileLayerIntoPositionViewPosition_TiledLayer ILayer = NullLayer_GetInstance()
var StaticTileLayerIntoPositionViewPosition_Layer IAllBinaryLayer = AllBinaryLayer_NULL_ALLBINARY_LAYER


func NewStaticTileLayerIntoPositionViewPosition() *StaticTileLayerIntoPositionViewPosition {
    this := &StaticTileLayerIntoPositionViewPosition{}
        this.ViewPosition = *NewViewPosition(0, 0, 0)
    return this
}


func (this *StaticTileLayerIntoPositionViewPosition) GetX() int {
        return this.ViewPosition.GetX() - StaticTileLayerIntoPositionViewPosition_TiledLayer.GetXP()
}


func (this *StaticTileLayerIntoPositionViewPosition) GetY() int {
        return this.ViewPosition.GetY() - StaticTileLayerIntoPositionViewPosition_TiledLayer.GetYP()
}


func (this *StaticTileLayerIntoPositionViewPosition) GetZ() int {
        return this.ViewPosition.GetZ() - StaticTileLayerIntoPositionViewPosition_TiledLayer.GetZP()
}


func StaticTileLayerIntoPositionViewPosition_SetTiledLayer(tiledLayer IAllBinaryTiledLayer) {
        StaticTileLayerIntoPositionViewPosition_TiledLayer = tiledLayer
}

