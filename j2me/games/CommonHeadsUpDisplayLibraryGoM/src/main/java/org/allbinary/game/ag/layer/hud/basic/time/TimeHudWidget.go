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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type TimeHudWidget struct {
    BasicHud
    TIME_STRING string
    TIME_CHAR_ARRAY []rune
    Offset int
    String []rune
    TotalDigits int
    Timer ITimer
}


func (this *TimeHudWidget) initFields() {
    this.TIME_STRING = "Time "
    this.TIME_CHAR_ARRAY = []rune{'T', 'i', 'm', 'e', ' '}
    
    this.String = PrimitiveLongSingleton_GetInstance().ZERO
    this.TotalDigits = 1
    
}


// throws Exception
func TimeHudWidget_GetInstance(location int, direction int, basicColor *BasicColor, timer ITimer) *TimeHudWidget {
        return NewTimeHudWidget(location, direction, 2, basicColor, timer)
}


func NewTimeHudWidget(location int, direction int, bufferZone int, basicColor *BasicColor, timer ITimer) *TimeHudWidget {
    this := &TimeHudWidget{}
        this.BasicHud = *NewBasicHud(location, direction, bufferZone, basicColor)
        this.initFields()
        this.Timer = timer
        this.Set()
        this.UpdateMaxHeight = 14
    return this
}


func (this *TimeHudWidget) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.UpdateMaxWidth = font.GetSize() * 5
        this.Offset = font.StringWidth(this.TIME_STRING) + MyFontProcessor_DefaultStringWidth(font, 3)
        this.BasicHud.UpdateMeasurement(graphics)
}


func (this *TimeHudWidget) Update() {
        this.Timer.Update()
        this.Set()
}


func (this *TimeHudWidget) Set() {
        this.String = this.Timer.GetTimeChars()
        this.TotalDigits = this.Timer.GetCurrentTotalDigits()
}


func (this *TimeHudWidget) GetTimer() ITimer {
        return this.Timer
}


func (this *TimeHudWidget) Paint(graphics IGraphics) {
        this.BasicHud.PaintDX(graphics, this.TIME_CHAR_ARRAY, 0, len(this.TIME_CHAR_ARRAY), this.String, 0, this.TotalDigits, this.Offset)
}

