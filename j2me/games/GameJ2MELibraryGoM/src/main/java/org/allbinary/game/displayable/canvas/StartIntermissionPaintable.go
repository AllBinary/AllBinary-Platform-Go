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

type StartIntermissionPaintable struct {
    InitUpdatePaintable
    FontDebugFactory *FontDebugFactory
    BasicSetColorUtil IBasicColorSetUtil
    GameCanvas IAllBinaryGameCanvas
    StringArray []string
    FontSize int
    Font Font
    LastWidth []int
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    BasicColor *BasicColor
    Color int
    LineYOffsetArray []int
    HasChanged bool
    Anchor int
    FontHeightP int
    BEGIN_LEVEL string
    EMPTY_STRING string
}

// Implements UpdateMyFontInterface


func (this *StartIntermissionPaintable) initFields() {
    this.FontDebugFactory = FontDebugFactory_GetInstance()
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    
    
    
    
    
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.BasicColor = BasicColorFactory_GetInstance().BLACK
    
    this.LineYOffsetArray = NullUtil_GetInstance().NULL_INT_ARRAY
    this.HasChanged = true
    this.Anchor = Anchor_TOP_LEFT
    
    this.BEGIN_LEVEL = "Begin Level "
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
}


func NewStartIntermissionPaintable(gameCanvas IAllBinaryGameCanvas, stringArray []string, basicColor *BasicColor, font Font) *StartIntermissionPaintable {
    this := &StartIntermissionPaintable{}
    this.InitUpdatePaintable = *NewInitUpdatePaintable()
    this.initFields()
        this.GameCanvas = gameCanvas
        this.StringArray = stringArray
        this.LastWidth = make([]int, len(this.StringArray))
        this.SetBasicColorP(basicColor)
        this.Color = basicColor.IntValue()
        this.FontSize = font.GetSize()
        this.Font = font
    return this
}


func (this *StartIntermissionPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeightP = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *StartIntermissionPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var existingFont Font = graphics.GetFont(); _ = existingFont
        this.FontDebugFactory.SetFont(this.Font, graphics)
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.BasicColor, this.Color)
        var beginWidth int; _ = beginWidth
        for index := int(len(this.LineYOffsetArray) - 1); index >= 0; index-- {
            if this.HasChanged {
                this.LastWidth[index] = (graphics.GetFont().StringWidth(this.StringArray[index]) >> 1)
            }
            beginWidth = this.LastWidth[index]
            graphics.DrawString(this.StringArray[index], displayInfo.GetLastHalfWidth() - beginWidth, displayInfo.GetLastHalfHeight() - this.LineYOffsetArray[index], this.Anchor)
        }
        this.HasChanged = false
        this.FontDebugFactory.SetFont(existingFont, graphics)
}


func (this *StartIntermissionPaintable) Update() {
        var level int = this.GameCanvas.GetLayerManager().GetGameInfo().GetCurrentLevel(); _ = level
        this.StringArray[0] = fmt.Sprint(NewStringMaker().Append(this.BEGIN_LEVEL).Appendint(level))
        for index := int(len(this.StringArray) - 1); index >= 1; index-- {
            this.StringArray[index] = this.EMPTY_STRING
        }
}


func (this *StartIntermissionPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *StartIntermissionPaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}

