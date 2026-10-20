/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
 *  
 * By agreeing to this license you and any business entity you represent are 
 * legally bound to the AllBinary Open License Version 1 legal agreement. 
 *  
 * You may obtain the AllBinary Open License Version 1 legal agreement from 
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository. 
 *  
 * Created By: Travis Berthelot  
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

type VisualWaypointLayer struct {
    WaypointLayer
}



func NewVisualWaypointLayer(remoteInfo *RemoteInfo, parentLayer PathFindingLayerInterface, groupInterface []*Group, rootName string, name string, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, x int, y int) *VisualWaypointLayer {
    this := &VisualWaypointLayer{}
        this.WaypointLayer = *NewWaypointLayer(remoteInfo, parentLayer, groupInterface, rootName, name, animationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface, rectangle, x, y)
    return this
}


func (this *VisualWaypointLayer) CreateHudPaintable() ISelectionHudPaintable {
        var rtsLayerHudPaintable IWaypointInfoHudPaintable = WaypointInfoHudPaintableSingleton_GetInstance(); _ = rtsLayerHudPaintable
        rtsLayerHudPaintable.SetBasicColorP(this.AllBinaryGameLayerManagerP.GetForegroundBasicColor())
        rtsLayerHudPaintable.SetRtsLayer(this)
        return rtsLayerHudPaintable
}


func (this *VisualWaypointLayer) GetHudPaintable() ISelectionHudPaintable {
        var rtsLayerHudPaintable IWaypointInfoHudPaintable = WaypointInfoHudPaintableSingleton_GetInstance(); _ = rtsLayerHudPaintable
        return rtsLayerHudPaintable
}

