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

type RaceTrackPickupLayer struct {
    PickupLayer
    ViewPositionEventHandler *ViewPositionEventHandler
}

var RaceTrackPickupLayer_NAME string = "RaceTrackPickupLayer"

func (this *RaceTrackPickupLayer) initFields() {
    this.ViewPositionEventHandler = ViewPositionEventHandler_GetInstance()
}


func NewRaceTrackPickupLayer() *RaceTrackPickupLayer {
    this := &RaceTrackPickupLayer{}
        this.PickupLayer = *NewPickupLayer(RaceTrackPickupLayer_NAME, RemoteInfo_REMOTE_INFO, 0, CountedPickedUpLayerInterfaceFactory_NULL_COUNTED_PICKUP_LAYER_FACTORY, NullAnimationFactory_GetFactoryInstance().GetInstance(0), NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 0, 0), NewTileLayerPositionIntoViewPosition())
        this.initFields()
    return this
}


func (this *RaceTrackPickupLayer) SetTiledLayer(tiledLayer IAllBinaryTiledLayer) {
        var viewPosition *TileLayerPositionIntoViewPosition = any(this.GetViewPosition()).(*TileLayerPositionIntoViewPosition); _ = viewPosition
        viewPosition.SetTiledLayer(tiledLayer)
        this.ViewPositionEventHandler.AddListener(any(this).(IAllBinaryLayer))
}


func (this *RaceTrackPickupLayer) SetDestroyed(destroyed bool) {
        this.PickupLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            this.ViewPositionEventHandler.RemoveListener(this)
        }
}

