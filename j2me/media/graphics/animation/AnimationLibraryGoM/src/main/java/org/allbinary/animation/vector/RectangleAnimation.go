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

type RectangleAnimation struct {
    Animation
    Width int
    Height int
}

// Implements ColorCompositeInterface


func (this *RectangleAnimation) initFields() {
    
    
}


func NewRectangleAnimation(width int, height int, basicColor *BasicColor) *RectangleAnimation {
    this := &RectangleAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.Width = width
        this.Height = height
        this.SetBasicColorP(basicColor)
    return this
}


func (this *RectangleAnimation) NextFrame() {
}


func (this *RectangleAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        graphics.DrawRect(x, y, this.Width, this.Height)
}


func (this *RectangleAnimation) SetWidth(width int) {
        this.Width = width
}


func (this *RectangleAnimation) SetHeight(height int) {
        this.Height = height
}

