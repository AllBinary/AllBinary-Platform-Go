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

type GeologicalResourceInfoHudPaintable struct {
    WaypointInfoHudPaintable
    KeyvalueDrawString *KeyValueDrawCharArray
    FontHeight int
}

var GeologicalResourceInfoHudPaintable_RESOURCES string = "Resources: "

func (this *GeologicalResourceInfoHudPaintable) initFields() {
    
    this.FontHeight = 0
}


func NewGeologicalResourceInfoHudPaintable() *GeologicalResourceInfoHudPaintable {
    this := &GeologicalResourceInfoHudPaintable{}
    this.WaypointInfoHudPaintable = *NewWaypointInfoHudPaintable()
    this.initFields()
        this.KeyvalueDrawString = NewKeyValueDrawCharArray(GeologicalResourceInfoHudPaintable_RESOURCES, this.TextX)
    return this
}


func (this *GeologicalResourceInfoHudPaintable) UpdateMeasurement(graphics IGraphics) {
        this.WaypointInfoHudPaintable.UpdateMeasurement(graphics)
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = (2 * font.GetHeight())
}


func (this *GeologicalResourceInfoHudPaintable) UpdateSelectionInfo() {
        this.WaypointInfoHudPaintable.UpdateSelectionInfo()
        var rtsLayer IRTSLayer = any(this.RtsLayerP).(IRTSLayer); _ = rtsLayer
        var list IBasicArrayList = rtsLayer.RTSLayer_().GeographicMapCellPositionAreaBase.GetOccupyingGeographicMapCellPositionList(); _ = list
        var total int = 0; _ = total
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            var geographicMapCellPosition *GeologicalGeographicMapCellPosition = any(list.Get(index)).(*GeologicalGeographicMapCellPosition); _ = geographicMapCellPosition
            total += geographicMapCellPosition.GetGeologicalResource().GetTotal()
        }
        this.KeyvalueDrawString.Update(this.GetPrimitiveLongUtil().GetCharArray(total), this.GetPrimitiveLongUtil().GetCurrentTotalDigits())
}


func (this *GeologicalResourceInfoHudPaintable) Paint(graphics IGraphics) {
        this.WaypointInfoHudPaintable.Paint(graphics)
        this.KeyvalueDrawString.Paint(graphics, (this.Y + this.FontHeight))
}

