/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type ViewPosition struct {
    ViewPositionBase
    AllbinaryLayer IAllBinaryLayer
}

// The Java class ViewPosition has subclasses so its variables are this interface
type IViewPosition interface {
    IViewPositionBase
    ViewPosition_() *ViewPosition
    GetX() int
    GetY() int
    GetZ() int
    GetX2() int
    GetY2() int
    GetZ2() int
    SetAllbinaryLayer(allbinaryLayer any)
}

func (this *ViewPosition) ViewPosition_() *ViewPosition {
    return this
}


func (this *ViewPosition) initFields() {
    this.AllbinaryLayer = AllBinaryLayer_NULL_ALLBINARY_LAYER
}


func ViewPosition_GetInstanceD() IViewPosition {
        return NewViewPosition(0, 0, 0)
}


func NewViewPosition(x int, y int, z int) *ViewPosition {
    this := &ViewPosition{}
        this.ViewPositionBase = *NewViewPositionBase(x, y, z)
        this.initFields()
    return this
}


func (this *ViewPosition) GetX() int {
        return this.AllbinaryLayer.GetXP()
}


func (this *ViewPosition) GetY() int {
        return this.AllbinaryLayer.GetYP()
}


func (this *ViewPosition) GetZ() int {
        return this.AllbinaryLayer.GetZP()
}


func (this *ViewPosition) GetX2() int {
        return this.GetX() + this.AllbinaryLayer.GetWidth()
}


func (this *ViewPosition) GetY2() int {
        return this.GetY() + this.AllbinaryLayer.GetHeight()
}


func (this *ViewPosition) GetZ2() int {
        return int((int64(this.GetZ()) + this.AllbinaryLayer.GetDepth()))
}


func (this *ViewPosition) SetAllbinaryLayer(allbinaryLayer any) {
        this.AllbinaryLayer = any(allbinaryLayer).(IAllBinaryLayer)
}

