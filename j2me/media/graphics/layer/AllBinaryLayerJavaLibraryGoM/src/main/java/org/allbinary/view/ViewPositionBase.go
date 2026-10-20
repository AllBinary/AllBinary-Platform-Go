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

type ViewPositionBase struct {
    GPoint
}

// The Java class ViewPositionBase has subclasses so its variables are this interface
type IViewPositionBase interface {
    IGPoint
    ViewPositionBase_() *ViewPositionBase
    GetX2() int
    GetY2() int
    GetZ2() int
    SetAllbinaryLayer(allbinaryLayer any)
}

func (this *ViewPositionBase) ViewPositionBase_() *ViewPositionBase {
    return this
}

var ViewPositionBase_NULL_VIEW_POSITION IViewPositionBase = NewViewPositionBase(0, 0, 0)


func NewViewPositionBase(x int, y int, z int) *ViewPositionBase {
    this := &ViewPositionBase{}
        this.GPoint = *NewGPoint(x, y, z)
    return this
}


func (this *ViewPositionBase) GetX2() int {
        return this.GetX()
}


func (this *ViewPositionBase) GetY2() int {
        return this.GetY()
}


func (this *ViewPositionBase) GetZ2() int {
        return this.GetZ()
}


func (this *ViewPositionBase) SetAllbinaryLayer(allbinaryLayer any) {
}

