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

type PressStartMenuPaintable struct {
    Paintable
    DisplayInfo *DisplayInfoSingleton
    PRESS_START string
    KEY_START string
    MENU_START string
    MyFontProcessor IMyFontProcessor
    StartString string
    TimeDelayHelper ITimeDelayHelper
    Flash bool
    Anchor int
    BeginWidth int
    Line int
}

// Implements UpdateMyFontInterface


func (this *PressStartMenuPaintable) initFields() {
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.PRESS_START = "Press Screen To Start"
    this.KEY_START = "Press or Click F2 To Begin"
    this.MENU_START = "Press Start From The Menu To Begin"
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.StartString = StringUtil_GetInstance().EMPTY_STRING
    this.TimeDelayHelper = NewTimeDelayHelper(1100)
    
    this.Anchor = Anchor_TOP_LEFT
    
    
}


func NewPressStartMenuPaintable() *PressStartMenuPaintable {
    this := &PressStartMenuPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        if TouchScreenFactory_GetInstance().IsTouch() {
            this.StartString = this.PRESS_START
        } else if AppletUtil_IsAppletLoader(this) {
            this.StartString = this.KEY_START
        } else {
            this.StartString = this.MENU_START
        }
    return this
}


func (this *PressStartMenuPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.BeginWidth = (graphics.GetFont().StringWidth(this.StartString) >> 1)
        this.Line = (4 * MyFontProcessor_DefaultCharWidth(font)) + (font.GetHeight() >> 1)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *PressStartMenuPaintable) Paint(graphics IGraphics) {
        if this.TimeDelayHelper.IsTimeTNT() {
            if this.IsFlash() {
                this.SetFlash(false)
            } else {
                this.SetFlash(true)
            }
        }
        if this.IsFlash() {
            this.MyFontProcessor.Process(graphics)
            graphics.DrawString(this.StartString, this.DisplayInfo.GetLastHalfWidth() - this.BeginWidth, this.DisplayInfo.GetLastHeight() - this.Line, this.Anchor)
        }
}


func (this *PressStartMenuPaintable) SetFlash(flash bool) {
        this.Flash = flash
}


func (this *PressStartMenuPaintable) IsFlash() bool {
        return this.Flash
}

