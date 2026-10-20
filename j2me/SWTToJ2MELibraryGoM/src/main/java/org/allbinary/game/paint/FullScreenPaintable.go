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

//SwtToJ2ME
type FullScreenPaintable struct {
    Paintable
    DisplayInfo *DisplayInfoSingleton
    MyFontProcessor IMyFontProcessor
    FULLSCREEN_TEXT string
    Anchor int
    Y int
    BeginWidth int
}

// Implements UpdateMyFontInterface


func (this *FullScreenPaintable) initFields() {
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.FULLSCREEN_TEXT = "F11 - Toggle Fullscreen"
    this.Anchor = Anchor_TOP_LEFT
    
    
}


func FullScreenPaintable_GetInstance() *FullScreenPaintable {
        return NewFullScreenPaintable()
}


func NewFullScreenPaintable() *FullScreenPaintable {
    this := &FullScreenPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
    return this
}


func (this *FullScreenPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.Y = 4 * font.GetHeight()
        this.BeginWidth = (font.StringWidth(this.FULLSCREEN_TEXT) >> 1)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *FullScreenPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var halfWidth int = this.DisplayInfo.GetLastHalfWidth(); _ = halfWidth
        var height int = this.DisplayInfo.GetLastHeight(); _ = height
        graphics.DrawString(this.FULLSCREEN_TEXT, halfWidth - this.BeginWidth, height - this.Y, this.Anchor)
}

