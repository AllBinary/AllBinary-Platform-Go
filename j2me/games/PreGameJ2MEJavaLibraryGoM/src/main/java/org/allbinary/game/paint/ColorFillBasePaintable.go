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

type ColorFillBasePaintable struct {
    Paintable
    BasicColor *BasicColor
    ColorP int
}

// Implements ColorCompositeInterface

// The Java class ColorFillBasePaintable has subclasses so its variables are this interface
type IColorFillBasePaintable interface {
    IPaintable
    ColorFillBasePaintable_() *ColorFillBasePaintable
    SetBasicColorP(basicColor *BasicColor)
    GetBasicColorP() *BasicColor
    GetColor() int
    SetColor(color int)
    Paint(graphics IGraphics)
}

func (this *ColorFillBasePaintable) ColorFillBasePaintable_() *ColorFillBasePaintable {
    return this
}


func (this *ColorFillBasePaintable) initFields() {
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    
}


func NewColorFillBasePaintable(basicColor *BasicColor) *ColorFillBasePaintable {
    this := &ColorFillBasePaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.SetBasicColorP(basicColor)
    return this
}


func (this *ColorFillBasePaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
        this.SetColor(this.GetBasicColorP().IntValue())
}


func (this *ColorFillBasePaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *ColorFillBasePaintable) GetColor() int {
        return this.ColorP
}


func (this *ColorFillBasePaintable) SetColor(color int) {
        this.ColorP = color
}


func (this *ColorFillBasePaintable) Paint(graphics IGraphics) {
}

