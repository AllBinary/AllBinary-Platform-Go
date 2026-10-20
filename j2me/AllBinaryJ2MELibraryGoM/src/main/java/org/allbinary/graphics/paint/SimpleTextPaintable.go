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

type SimpleTextPaintable struct {
    Paintable
    DisplayInfoSingleton *DisplayInfoSingleton
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    Text string
    BasicColor *BasicColor
    TopScoresWidth int
    FontHeight int
    Anchor int
}

// Implements UpdateMyFontInterface


func (this *SimpleTextPaintable) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    
    
    this.TopScoresWidth = 0
    this.FontHeight = 0
    this.Anchor = Anchor_TOP_LEFT
}


func NewSimpleTextPaintable(text string, basicColor *BasicColor) *SimpleTextPaintable {
    this := &SimpleTextPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.Text = text
        this.BasicColor = basicColor
    return this
}


func (this *SimpleTextPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.TopScoresWidth = (font.StringWidth(this.Text) >> 1)
        this.FontHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *SimpleTextPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var width int = this.DisplayInfoSingleton.GetLast()[this.DisplayInfoSingleton.WIDTH]; _ = width
        graphics.SetColor(this.GetBasicColorP().IntValue())
        graphics.DrawString(this.Text, (width >> 1) - this.TopScoresWidth, this.FontHeight, this.Anchor)
}


func (this *SimpleTextPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *SimpleTextPaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *SimpleTextPaintable) SetText(text string) {
        this.Text = text
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *SimpleTextPaintable) GetText() string {
        return this.Text
}

