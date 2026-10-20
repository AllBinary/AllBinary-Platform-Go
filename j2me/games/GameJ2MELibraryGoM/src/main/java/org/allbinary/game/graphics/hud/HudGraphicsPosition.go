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

type HudGraphicsPosition struct {
    Anchor int
    Point IGPoint
}

var HudGraphicsPosition_NULL_HUD_GRAPHICS_POSITION *HudGraphicsPosition = NewHudGraphicsPosition(PointFactory_GetInstance().ZERO_ZERO, 9)

func (this *HudGraphicsPosition) initFields() {
    
    this.Point = PointFactory_GetInstance().ZERO_ZERO
}


func NewHudGraphicsPosition(point IGPoint, anchor int) *HudGraphicsPosition {
    this := &HudGraphicsPosition{}
    this.initFields()
        this.SetAnchor(anchor)
        this.SetPoint(point)
    return this
}


func (this *HudGraphicsPosition) GetAnchor() int {
        return this.Anchor
}


func (this *HudGraphicsPosition) SetAnchor(anchor int) {
        this.Anchor = anchor
}


func (this *HudGraphicsPosition) GetPoint() IGPoint {
        return this.Point
}


func (this *HudGraphicsPosition) SetPoint(point IGPoint) {
        this.Point = point
}

