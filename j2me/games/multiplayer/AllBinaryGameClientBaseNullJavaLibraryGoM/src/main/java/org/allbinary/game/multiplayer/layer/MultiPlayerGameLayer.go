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

type MultiPlayerGameLayer struct {
    CollidableDestroyableDamageableLayer
}

// Implements MultiPlayerGameLayerInterface

// The Java class MultiPlayerGameLayer has subclasses so its variables are this interface
type IMultiPlayerGameLayer interface {
    ICollidableDestroyableDamageableLayer
    MultiPlayerGameLayer_() *MultiPlayerGameLayer
    GetMultiplayerBehavior() IMultiplayerBehavior
    SetPingInfo(show bool)
    SetFiring(firing int16)
    GetFiring() int16
}

func (this *MultiPlayerGameLayer) MultiPlayerGameLayer_() *MultiPlayerGameLayer {
    return this
}



func NewMultiPlayerGameLayer(remoteInfo *RemoteInfo, groupInterface []*Group, name string, layerInfo *Rectangle, viewPosition IViewPositionBase) *MultiPlayerGameLayer {
    this := &MultiPlayerGameLayer{}
        this.CollidableDestroyableDamageableLayer = *NewCollidableDestroyableDamageableLayer(groupInterface, name, layerInfo, viewPosition)
    return this
}


func (this *MultiPlayerGameLayer) GetMultiplayerBehavior() IMultiplayerBehavior {
        return MultiplayerBehavior_NULL_MULTIPLAYER_BEHAVIOR
}


func (this *MultiPlayerGameLayer) SetPingInfo(show bool) {
}


func (this *MultiPlayerGameLayer) SetFiring(firing int16) {
}


func (this *MultiPlayerGameLayer) GetFiring() int16 {
        return 0
}

