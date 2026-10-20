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

type VelocityWidget struct {
    BasicHud
    KILOMETERS_PER_HOUR_STR []rune
    TotalChars int
    Velocity int
    MaxVelocity int
    String []rune
    TotalDigits int
    PowerOfTenVelocity int
    PrimitiveLongUtil *PrimitiveLongUtil
    Offset int
    Offset2 int
}


func (this *VelocityWidget) initFields() {
    this.KILOMETERS_PER_HOUR_STR = []rune{' ', 'k', 'm', '/', 'h'}
    this.TotalChars = len(this.KILOMETERS_PER_HOUR_STR)
    
    
    this.String = PrimitiveLongSingleton_GetInstance().ZERO
    this.TotalDigits = 1
    
    
    this.Offset = 0
    this.Offset2 = 0
}


func NewVelocityWidget(powerOfTenVelocity int, location int, direction int, basicColor *BasicColor) *VelocityWidget {
    this := &VelocityWidget{}
        this.BasicHud = *NewBasicHud(location, direction, 2, basicColor)
        this.initFields()
        this.PowerOfTenVelocity = powerOfTenVelocity
        this.MaxVelocity = powerOfTenVelocity
        this.Velocity = 0
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(powerOfTenVelocity)
        this.UpdateMaxHeight = 14
    return this
}


func (this *VelocityWidget) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.UpdateMaxWidth = font.GetSize() * (5 + MathUtil_GetInstance().GetTotalDigits(this.PowerOfTenVelocity) + 1)
        this.BasicHud.UpdateMeasurement(graphics)
        this.Offset = MyFontProcessor_DefaultStringWidth(font, this.PrimitiveLongUtil.GetMaxDigits()) + MyFontProcessor_DefaultStringWidth(font, 2)
        this.Offset2 = this.Offset - MyFontProcessor_DefaultStringWidth(font, this.TotalDigits) - MyFontProcessor_DefaultStringWidth(font, 2)
}


func (this *VelocityWidget) Get() int {
        return this.Velocity
}


func (this *VelocityWidget) Add(value int) {
        this.Set(this.Velocity + value)
}


func (this *VelocityWidget) Set(value int) {
        var lastVelocity int = this.Velocity; _ = lastVelocity
        this.Velocity = value
        if this.Velocity > this.MaxVelocity {
            this.Velocity = 0
        }
        if lastVelocity != this.Velocity {
            this.String = this.PrimitiveLongUtil.GetCharArray(this.Velocity * 18)
            if OpenGLFeatureUtil_GetInstance().IsAnyThreed() {
                this.TotalDigits = this.PrimitiveLongUtil.GetCurrentTotalDigits() + 1
            } else {
                this.TotalDigits = this.PrimitiveLongUtil.GetCurrentTotalDigits()
            }
            this.MyFontProcessor = this.UpdateMyFontProcessor
        }
}


func (this *VelocityWidget) Reduce(value int) {
        this.Set(this.Velocity - value)
}


func (this *VelocityWidget) Paint(graphics IGraphics) {
        this.BasicHud.PaintDXY(graphics, this.String, 0, this.TotalDigits, this.KILOMETERS_PER_HOUR_STR, 0, this.TotalChars, this.Offset2, this.Offset)
}

