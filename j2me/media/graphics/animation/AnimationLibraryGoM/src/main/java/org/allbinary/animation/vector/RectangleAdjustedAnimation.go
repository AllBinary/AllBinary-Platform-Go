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

type RectangleAdjustedAnimation struct {
    Animation
    Width int
    Height int
    OffsetX int
    OffsetY int
}

// Implements ColorCompositeInterface


func (this *RectangleAdjustedAnimation) initFields() {
    
    
    
    
}


func NewRectangleAdjustedAnimation(width int, height int, offsetX int, offsetY int, basicColor *BasicColor) *RectangleAdjustedAnimation {
    this := &RectangleAdjustedAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.Width = width
        this.Height = height
        this.OffsetX = offsetX
        this.OffsetY = offsetY
        this.SetBasicColorP(basicColor)
    return this
}


func (this *RectangleAdjustedAnimation) NextFrame() {
}


func (this *RectangleAdjustedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        graphics.DrawRect(x + this.OffsetX, y + this.OffsetY, this.Width, this.Height)
}


func (this *RectangleAdjustedAnimation) SetWidth(width int) {
        this.Width = width
}


func (this *RectangleAdjustedAnimation) SetHeight(height int) {
        this.Height = height
}


func (this *RectangleAdjustedAnimation) SetOffsetX(offsetX int) {
        this.OffsetX = offsetX
}


func (this *RectangleAdjustedAnimation) SetOffsetY(offsetY int) {
        this.OffsetY = offsetY
}

