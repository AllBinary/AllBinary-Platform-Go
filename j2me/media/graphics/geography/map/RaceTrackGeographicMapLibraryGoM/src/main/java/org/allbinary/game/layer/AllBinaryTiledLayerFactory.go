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

type AllBinaryTiledLayerFactory struct {
}

// Implements AllBinaryTiledLayerFactoryInterface


func NewAllBinaryTiledLayerFactory() *AllBinaryTiledLayerFactory {
    this := &AllBinaryTiledLayerFactory{}
    return this
}


// throws Exception
func (this *AllBinaryTiledLayerFactory) GetInstance(raceTrackInfo *RaceTrackInfo, raceTrackData *RaceTrackData) IAllBinaryTiledLayer {
        var allbinaryTiledLayer IAllBinaryTiledLayer; _ = allbinaryTiledLayer
        var tiledLayer *TiledLayer = NewTiledLayerFactory().GetInstance(raceTrackData); _ = tiledLayer
        allbinaryTiledLayer = NewAllBinaryJ2METiledLayer(raceTrackData.GetId(), tiledLayer, raceTrackData.GetMapArray(), BasicColorFactory_GetInstance().RED.IntValue())
        allbinaryTiledLayer.SetCells(raceTrackData.GetMapArray())
        return allbinaryTiledLayer
}


// throws Exception
func (this *AllBinaryTiledLayerFactory) GetMiniInstance(raceTrackData *RaceTrackData) IAllBinaryTiledLayer {
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var tiledLayer *TiledLayer = NewTiledLayerFactory().GetMiniInstance(raceTrackData); _ = tiledLayer
        var allbinaryTiledLayer IAllBinaryTiledLayer = NewAllBinaryJ2METiledLayer(smallIntegerSingletonFactory.GetAt(int(raceTrackData.GetId()) + 100 + 1), tiledLayer, raceTrackData.GetMapArray(), BasicColorFactory_GetInstance().RED.IntValue()); _ = allbinaryTiledLayer
        allbinaryTiledLayer.SetCells(raceTrackData.GetMapArray())
        return allbinaryTiledLayer
}

