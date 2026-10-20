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

type Grid struct {
    Grid IGPoint
    IsGridOn bool
    IsGridPossible bool
    Zoom int
    IsChanged bool
}


func (this *Grid) initFields() {
    
    this.IsGridOn = true
    this.IsGridPossible = false
    this.Zoom = 3
    this.IsChanged = true
}




// Java overloads of new Grid
func NewGrid(args ...any) *Grid {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewGrid_1()
        }
        if len(args) == 1 && Java_IsArg[*Grid](args[0], pass) {
            return NewGrid_2(Java_Arg[*Grid](args[0]))
        }
    }
    panic("No Java overload of new Grid for the arguments")
}
func NewGrid_1() *Grid {
    this := &Grid{}
    this.initFields()
    return this
}


func NewGrid_2(grid *Grid) *Grid {
    this := &Grid{}
    this.initFields()
        this.Grid = grid.GPoint_().Grid
        this.ShowGrid(grid.GPoint_().IsGridOn)
        this.IsGridPossible = grid.GPoint_().IsGridPossible
        this.SetZoom(grid.GetZoom())
    return this
}


func (this *Grid) IsGridOn() bool {
        return this.IsGridOn
}


func (this *Grid) IsGridPossible() bool {
        return this.IsGridPossible
}


func (this *Grid) GetGrid() IGPoint {
        return this.Grid
}


func (this *Grid) GetZoom() int {
        return this.Zoom
}


func (this *Grid) ShowGrid(value bool) {
        this.IsGridOn = value
}


func (this *Grid) SetZoom(zoom int) {
        this.Zoom = zoom
        if this.Zoom < 1 {
            zoom = 1
        }
        this.IsChanged = true
}

