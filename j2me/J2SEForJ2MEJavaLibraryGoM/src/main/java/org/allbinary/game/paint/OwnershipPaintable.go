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

//J2SEForJ2ME
type OwnershipPaintable struct {
    Paintable
    LogUtil *LogUtil
    DisplayInfo *DisplayInfoSingleton
    MyFontProcessor IMyFontProcessor
    COPYRIGHT string
    BasicColor *BasicColor
    Color int
    Anchor int
    COPYRIGHT_Y int
    BeginWidth int
}

// Implements UpdateMyFontInterface


func (this *OwnershipPaintable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.COPYRIGHT = "AllBinary Copyright (c) 2011"
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    this.Color = this.BasicColor.IntValue()
    this.Anchor = Anchor_TOP_LEFT
    
    
}


func OwnershipPaintable_GetInstance() *OwnershipPaintable {
        return NewOwnershipPaintable()
}


func NewOwnershipPaintable() *OwnershipPaintable {
    this := &OwnershipPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
    return this
}


func (this *OwnershipPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.COPYRIGHT_Y = 2 * font.GetHeight()
        this.BeginWidth = (font.StringWidth(this.COPYRIGHT) >> 1)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *OwnershipPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
        this.Color = basicColor.IntValue()
}


func (this *OwnershipPaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *OwnershipPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        graphics.SetColor(this.Color)
        var halfWidth int = this.DisplayInfo.GetLastHalfWidth(); _ = halfWidth
        var height int = this.DisplayInfo.GetLastHeight(); _ = height
        graphics.DrawString(this.COPYRIGHT, halfWidth - this.BeginWidth, height - this.COPYRIGHT_Y, this.Anchor)
}

