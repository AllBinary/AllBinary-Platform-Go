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

type AdjustedCircleAnimation struct {
    CircleAnimation
    XAdjustment int
    YAdjustment int
}


func (this *AdjustedCircleAnimation) initFields() {
    
    
}


func AdjustedCircleAnimation_CreateAnimation(width int, height int, xAdjustment int, yAdjustment int, basicColor *BasicColor) *AdjustedCircleAnimation {
        var adjustedCircleAnimation *AdjustedCircleAnimation = NewAdjustedCircleAnimation(width, height, basicColor); _ = adjustedCircleAnimation
        adjustedCircleAnimation.InitXY(xAdjustment, yAdjustment)
        return adjustedCircleAnimation
}


func AdjustedCircleAnimation_CreateW(width int, height int, innerWidth int, basicColor *BasicColor) *AdjustedCircleAnimation {
        var adjustedCircleAnimation *AdjustedCircleAnimation = NewAdjustedCircleAnimation(width, height, basicColor); _ = adjustedCircleAnimation
        adjustedCircleAnimation.Init(innerWidth)
        return adjustedCircleAnimation
}


func NewAdjustedCircleAnimation(width int, height int, basicColor *BasicColor) *AdjustedCircleAnimation {
    this := &AdjustedCircleAnimation{}
        this.CircleAnimation = *NewCircleAnimation(width, height, basicColor)
        this.initFields()
    return this
}


func (this *AdjustedCircleAnimation) InitXY(xAdjustment int, yAdjustment int) {
        this.XAdjustment = xAdjustment
        this.YAdjustment = yAdjustment
}


func (this *AdjustedCircleAnimation) Init(innerWidth int) {
        if AndroidUtil_IsAndroid() {
            var widthPortion int = (((innerWidth / 3) * 10) / 7); _ = widthPortion
            this.XAdjustment = -((this.Width >> 1) - widthPortion)
            this.YAdjustment = this.XAdjustment
        } else {
            this.XAdjustment = -((this.Width >> 1) - innerWidth)
            this.YAdjustment = this.XAdjustment
        }
}


func (this *AdjustedCircleAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        graphics.DrawArc(x + this.XAdjustment, y + this.YAdjustment, this.Width, this.Height, 0, this.TOTAL_ANGLE)
}

