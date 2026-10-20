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

type NumberStringHud struct {
    BasicHud
    PrimitiveLongUtil *PrimitiveLongUtil
    PrependString string
    PREPEND_STRING []rune
    Value int
    Max int
    Offset int
    ValueString []rune
    ValueTotalDigits int
}

// Implements PaintableInterface

// The Java class NumberStringHud has subclasses so its variables are this interface
type INumberStringHud interface {
    IBasicHud
    NumberStringHud_() *NumberStringHud
    UpdateMeasurement(graphics IGraphics)
    Get() int
    Add(value int)
    Set(value int)
    Reduce(value int)
    Paint(graphics IGraphics)
    PaintXY(graphics IGraphics, x int, y int)
    PaintThreed(graphics IGraphics)
}

func (this *NumberStringHud) NumberStringHud_() *NumberStringHud {
    return this
}


func (this *NumberStringHud) initFields() {
    
    
    
    
    
    
    
    
}


func NewNumberStringHud(prependString string, max_ int, location int, direction int, bufferZone int, basicColor *BasicColor) *NumberStringHud {
    this := &NumberStringHud{}
        this.BasicHud = *NewBasicHud(location, direction, bufferZone, basicColor)
        this.initFields()
        this.PrependString = prependString
        this.PREPEND_STRING = []rune(prependString)
        this.ValueString = PrimitiveLongSingleton_GetInstance().NUMBER_CHAR_ARRAYS[0]
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(max_ + 1)
        this.Max = max_
        this.Value = 0
        if direction == 0 {
            panic(errors.New(fmt.Sprint(BasicHudFactory_GetInstance().DIRECTION_EXCEPTION)))
        }
    return this
}


func (this *NumberStringHud) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.Offset = font.StringWidth(this.PrependString) + MyFontProcessor_DefaultCharWidth(font)
        this.BasicHud.UpdateMeasurement(graphics)
}


func (this *NumberStringHud) Get() int {
        return this.Value
}


func (this *NumberStringHud) Add(value int) {
        this.Set(this.Value + value)
}


func (this *NumberStringHud) Set(value int) {
        this.Value = value
        if this.Value > this.Max {
            this.Value = 0
        }
        this.ValueString = this.PrimitiveLongUtil.GetCharArray(this.Value)
        this.ValueTotalDigits = this.PrimitiveLongUtil.GetCurrentTotalDigits()
}


func (this *NumberStringHud) Reduce(value int) {
        this.Set(this.Value - value)
}


func (this *NumberStringHud) Paint(graphics IGraphics) {
        this.BasicHud.PaintDX(graphics, this.PREPEND_STRING, 0, len(this.PREPEND_STRING), this.ValueString, 0, this.ValueTotalDigits, this.Offset)
}


func (this *NumberStringHud) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        var charArray []rune = this.PREPEND_STRING; _ = charArray
        var len_ int = len(this.PREPEND_STRING); _ = len_
        var charArray2 []rune = this.ValueString; _ = charArray2
        var len2 int = this.ValueTotalDigits; _ = len2
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GetBasicColorP())
        graphics.DrawChars(charArray, 0, len_, x, y, 0)
        graphics.DrawChars(charArray2, 0, len2, x - this.Offset, y, 0)
}


func (this *NumberStringHud) PaintThreed(graphics IGraphics) {
}

