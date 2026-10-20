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

type WorkWaypoint struct {
    Waypoint
}

var WorkWaypoint_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(25)


func NewWorkWaypoint(ownerLayer PathFindingLayerInterface) *WorkWaypoint {
    this := &WorkWaypoint{}
        this.Waypoint = *NewWaypoint(ownerLayer, WorkSound_GetInstance())
    return this
}


// throws Exception
func (this *WorkWaypoint) Visit(unitLayer *UnitLayer) {
        if unitLayer.GetGroupInterface()[0] != this.OwnerLayer.GetGroupInterface()[0] {
            return
        }
        var size int = this.GetConnectedWaypointList().Size(); _ = size
        if size > 0 {
            var occupyList IBasicArrayList = this.OwnerLayer.GetGeographicMapCellPositionArea().GetOccupyingGeographicMapCellPositionList(); _ = occupyList
            var outOfResources bool = true; _ = outOfResources
            var geologicalGeographicMapCellPosition *GeologicalGeographicMapCellPosition; _ = geologicalGeographicMapCellPosition
            var geologicalResource *GeologicalResource; _ = geologicalResource
            for index := int(occupyList.Size()); Java_PreDecrement(&index) >= 0;  {
                geologicalGeographicMapCellPosition = any(occupyList.Get(index)).(*GeologicalGeographicMapCellPosition)
                geologicalResource = geologicalGeographicMapCellPosition.GetGeologicalResource()
                if geologicalResource.GetTotal() >= 1 {
                    outOfResources = false
                    var maxResourceLoad int = int(unitLayer.GetMaxResourceLoad()); _ = maxResourceLoad
                    geologicalResource.Remove(maxResourceLoad)
                    unitLayer.SetLoad(maxResourceLoad)
                }
            }
            var waypointInfoHudPaintable IWaypointInfoHudPaintable = any(this.OwnerLayer.GetHudPaintable()).(IWaypointInfoHudPaintable); _ = waypointInfoHudPaintable
            if this.OwnerLayer == waypointInfoHudPaintable.GetRtsLayer() {
                waypointInfoHudPaintable.UpdateSelectionInfo()
            }
            if !outOfResources {
                var unitWaypointBehavior IUnitWaypointBehavior = any(unitLayer.GetWaypointBehavior()).(IUnitWaypointBehavior); _ = unitWaypointBehavior
                var rtsLayer IRTSLayer; _ = rtsLayer
                for this.GetConnectedWaypointList().Size() > 0 {
                    rtsLayer = any(this.GetConnectedWaypointList().Get(0)).(IRTSLayer)
                    if rtsLayer.IsDestroyed() {
                        this.GetConnectedWaypointList().Remove(rtsLayer)
                    } else {
                        unitWaypointBehavior.InsertWaypoint(0, rtsLayer)
                        break
                    }
                }
            }
        }
}


func (this *WorkWaypoint) GetType() int {
        return 2
}

