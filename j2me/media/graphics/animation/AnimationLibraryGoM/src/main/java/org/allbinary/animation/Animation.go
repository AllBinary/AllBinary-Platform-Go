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

type Animation struct {
    BasicSetColorUtil IBasicColorSetUtil
    BasicColor *BasicColor
    ColorP int
    ChangeBasicColorP *BasicColor
    ChangeColorP int
    AlphaP int
}

// Implements AnimationInterface, OpenGLSurfaceChangedInterface

// The Java class Animation has subclasses so its variables are this interface
type IAnimation interface {
    Animation_() *Animation
    SetAlpha(alpha int)
    NextFrame()
    PaintXY(graphics IGraphics, x int, y int)
    PaintThreedXYZ(graphics IGraphics, x int, y int, z int)
    IsThreed() bool
    GetBasicColorP() *BasicColor
    SetBasicColorP(basicColor *BasicColor)
    SetBackgroundBasicColorP(basicColor *BasicColor)
    GetChangeBasicColor() *BasicColor
    ChangeBasicColor(basicColor *BasicColor)
    GetChangeColor() int
    GetColor() int
    SetScale(scaleX float32, scaleY float32)
    SetMaxScale(maxScaleX float32, maxScaleY float32)
    Set(gl GL)
    SetDx(dx int)
    GetDx() int
    SetDy(dy int)
    GetDy() int
}

func (this *Animation) Animation_() *Animation {
    return this
}


func (this *Animation) initFields() {
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    
    this.ChangeBasicColorP = BasicColorFactory_GetInstance().NULL_COLOR
    
    
}


func NewAnimation() *Animation {
    this := &Animation{}
    this.initFields()
    return this
}


func (this *Animation) SetAlpha(alpha int) {
        this.AlphaP = alpha
}


// throws Exception
func (this *Animation) NextFrame() {
}


func (this *Animation) PaintXY(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
}


func (this *Animation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
}


func (this *Animation) IsThreed() bool {
        return false
}


func (this *Animation) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *Animation) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
        this.ColorP = this.BasicColor.IntValue()
}


func (this *Animation) SetBackgroundBasicColorP(basicColor *BasicColor) {
}


func (this *Animation) GetChangeBasicColor() *BasicColor {
        return this.ChangeBasicColorP
}


func (this *Animation) ChangeBasicColor(basicColor *BasicColor) {
        this.ChangeBasicColorP = basicColor
        this.ChangeColorP = this.ChangeBasicColorP.IntValue()
}


func (this *Animation) GetChangeColor() int {
        return this.ChangeColorP
}


func (this *Animation) GetColor() int {
        return this.ColorP
}


func (this *Animation) SetScale(scaleX float32, scaleY float32) {
}


func (this *Animation) SetMaxScale(maxScaleX float32, maxScaleY float32) {
}


// throws Exception
func (this *Animation) Set(gl GL) {
}


func (this *Animation) SetDx(dx int) {
}


func (this *Animation) GetDx() int {
        return 0
}


func (this *Animation) SetDy(dy int) {
}


func (this *Animation) GetDy() int {
        return 0
}

