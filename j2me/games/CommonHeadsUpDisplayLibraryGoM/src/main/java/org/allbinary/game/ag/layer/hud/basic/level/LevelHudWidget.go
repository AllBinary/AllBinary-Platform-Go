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

type LevelHudWidget struct {
    BasicHud
    Level int
    Maxlevel int
    LevelString []rune
    LevelNumberCharArray []rune
    LevelNumberTotalDigits int
    Offset int
    PrimitiveLongUtil *PrimitiveLongUtil
}

// Implements PaintableInterface


func (this *LevelHudWidget) initFields() {
    
    
    
    this.LevelNumberCharArray = NullUtil_GetInstance().NULL_CHAR_ARRAY
    
    
    
}


// throws Exception
func LevelHudWidget_CreateHud(maxlevel int, location int, direction int) *LevelHudWidget {
        return NewLevelHudWidget(maxlevel, location, direction)
}


func NewLevelHudWidget(maxlevel int, location int, direction int) *LevelHudWidget {
    this := &LevelHudWidget{}
        this.BasicHud = *NewBasicHud(location, direction, 2, BasicColorFactory_GetInstance().GREY)
        this.initFields()
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(1000)
        var LEVEL string = "Lv "; _ = LEVEL
        this.LevelString = []rune(LEVEL)
        this.Maxlevel = maxlevel
        this.Level = maxlevel
        this.Update()
        this.UpdateMaxHeight = 14
    return this
}


func (this *LevelHudWidget) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.UpdateMaxWidth = font.GetSize() * 4
        this.Offset = font.CharsWidth(this.LevelString, 0, len(this.LevelString)) + font.GetSize()
        this.BasicHud.UpdateMeasurement(graphics)
}


func (this *LevelHudWidget) Update() {
        this.LevelNumberCharArray = this.PrimitiveLongUtil.GetCharArray(this.Level)
        this.LevelNumberTotalDigits = this.PrimitiveLongUtil.GetCurrentTotalDigits()
}


func (this *LevelHudWidget) SetLevel(level int) {
        this.Level = level
        this.Update()
}


func (this *LevelHudWidget) NextLevel() {
        this.Level++
        this.Update()
}


func (this *LevelHudWidget) PreviousLevel() {
        this.Level--
        if this.Level < 0 {
            this.Level = 0
        }
        this.Update()
}


func (this *LevelHudWidget) IsComplete() bool {
        if this.Level <= this.Maxlevel {
            return false
        } else {
            return true
        }
}


func (this *LevelHudWidget) Paint(graphics IGraphics) {
        this.BasicHud.PaintDX(graphics, this.LevelString, 0, len(this.LevelString), this.LevelNumberCharArray, 0, this.LevelNumberTotalDigits, this.Offset)
}


func (this *LevelHudWidget) PaintThreed(graphics IGraphics) {
}

