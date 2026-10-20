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

type CircleAnimation struct {
    Animation
    TOTAL_ANGLE int
    Width int
    Height int
}

// Implements ColorCompositeInterface

// The Java class CircleAnimation has subclasses so its variables are this interface
type ICircleAnimation interface {
    IAnimation
    CircleAnimation_() *CircleAnimation
    NextFrame()
    PaintXY(graphics IGraphics, x int, y int)
}

func (this *CircleAnimation) CircleAnimation_() *CircleAnimation {
    return this
}


func (this *CircleAnimation) initFields() {
    this.TOTAL_ANGLE = int(AngleFactory_GetInstance().TOTAL_ANGLE)
    
    
}


func NewCircleAnimation(width int, height int, basicColor *BasicColor) *CircleAnimation {
    this := &CircleAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.Width = width
        this.Height = height
        this.SetBasicColorP(basicColor)
    return this
}


func (this *CircleAnimation) NextFrame() {
}


func (this *CircleAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        graphics.DrawArc(x, y, this.Width, this.Height, 0, this.TOTAL_ANGLE)
}

