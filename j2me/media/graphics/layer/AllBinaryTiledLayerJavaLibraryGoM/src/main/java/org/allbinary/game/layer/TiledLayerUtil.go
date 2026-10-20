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

type TiledLayerUtil struct {
}

var TiledLayerUtil_Instance *TiledLayerUtil = NewTiledLayerUtil()

func NewTiledLayerUtil() *TiledLayerUtil {
    this := &TiledLayerUtil{}
    return this
}


func TiledLayerUtil_GetInstance() *TiledLayerUtil {
        return TiledLayerUtil_Instance
}


func (this *TiledLayerUtil) KeepOnBorderdMapX(terrainTiledLayer IAllBinaryTiledLayer, x int, width int, mapBorderBufferWidth int) int {
        var newX int = x; _ = newX
        var maxRight int = terrainTiledLayer.GetWidth() - width - mapBorderBufferWidth; _ = maxRight
        if x + width > maxRight {
            newX = maxRight
        }
        if x < mapBorderBufferWidth {
            newX = mapBorderBufferWidth
        }
        return newX
}


func (this *TiledLayerUtil) KeepOnBorderdMapY(terrainTiledLayer IAllBinaryTiledLayer, y int, height int, mapBorderBufferHeight int) int {
        var newY int = y; _ = newY
        var maxBottom int = terrainTiledLayer.GetHeight() - height - mapBorderBufferHeight; _ = maxBottom
        if y + height > maxBottom {
            newY = maxBottom
        }
        if y < mapBorderBufferHeight {
            newY = mapBorderBufferHeight
        }
        return newY
}


func (this *TiledLayerUtil) KeepOnMapX(terrainTiledLayer IAllBinaryTiledLayer, x int, width int) int {
        var newX int = x; _ = newX
        var maxRight int = terrainTiledLayer.GetWidth() - width; _ = maxRight
        if x >= maxRight {
            newX = maxRight
        }
        if x < 0 {
            newX = 0
        }
        return newX
}


func (this *TiledLayerUtil) KeepOnMapY(terrainTiledLayer IAllBinaryTiledLayer, y int, height int) int {
        var newY int = y; _ = newY
        var maxBottom int = terrainTiledLayer.GetHeight() - height; _ = maxBottom
        if y >= maxBottom {
            newY = maxBottom
        }
        if y < 0 {
            newY = 0
        }
        return newY
}

