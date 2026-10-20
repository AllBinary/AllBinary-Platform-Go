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

type AllBinaryVisualDebugVibration struct {
    AllBinaryVibrationME
    DisplayInfoSingleton *DisplayInfoSingleton
    TimeDelayHelper ITimeDelayHelper
    VIBRATING string
    MyFontProcessor IMyFontProcessor
    Anchor int
    Width int
}

// Implements UpdateMyFontInterface


func (this *AllBinaryVisualDebugVibration) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(Integer_MAX_VALUE)
    this.VIBRATING = "Vibrating"
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.Anchor = Anchor_TOP_LEFT
    
}

func NewAllBinaryVisualDebugVibration() *AllBinaryVisualDebugVibration {
    this := &AllBinaryVisualDebugVibration{}
    this.AllBinaryVibrationME = *NewAllBinaryVibrationME()
    this.initFields()
    return this
}


func (this *AllBinaryVisualDebugVibration) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.Width = font.StringWidth(this.VIBRATING)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *AllBinaryVisualDebugVibration) Vibrate(duration int, type_ int, volume int) {
        this.TimeDelayHelper.TimeDelayHelper_().Delay = duration
}


func (this *AllBinaryVisualDebugVibration) Paint(graphics IGraphics) {
        if this.TimeDelayHelper.IsTimeTNT() {
            this.TimeDelayHelper.TimeDelayHelper_().Delay = Integer_MAX_VALUE
            this.MyFontProcessor.Process(graphics)
            graphics.DrawString(this.VIBRATING, this.DisplayInfoSingleton.GetLastHalfWidth() - (this.Width >> 1), 0, this.Anchor)
        }
}

