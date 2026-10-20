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

type HelpPaintable struct {
    Paintable
    DisplayInfo *DisplayInfoSingleton
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    Anchor int
    Title string
    InputInfo []string
    BasicColor *BasicColor
    TitleBeginWidth int
    BeginWidthArray []int
    CharHeight int
}

// Implements UpdateMyFontInterface

// The Java class HelpPaintable has subclasses so its variables are this interface
type IHelpPaintable interface {
    IPaintable
    HelpPaintable_() *HelpPaintable
    UpdateMeasurement(graphics IGraphics)
    SetInputInfoP(inputInfo []string)
    Paint(graphics IGraphics)
}

func (this *HelpPaintable) HelpPaintable_() *HelpPaintable {
    return this
}


func (this *HelpPaintable) initFields() {
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.Anchor = Anchor_TOP_LEFT
    
    this.InputInfo = StringUtil_GetInstance().GetArrayInstance()
    
    
    this.BeginWidthArray = NullUtil_GetInstance().NULL_INT_ARRAY
    
}


func NewHelpPaintable(title string, backgroundBasicColor *BasicColor, basicColor *BasicColor) *HelpPaintable {
    this := &HelpPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.Title = title
        this.BasicColor = basicColor
    return this
}


func (this *HelpPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.TitleBeginWidth = (font.StringWidth(this.Title) >> 1)
        this.CharHeight = font.GetHeight()
        var size int = len(this.InputInfo); _ = size
        for index := int(0); index < size; index++ {
            this.BeginWidthArray[index] = (font.StringWidth(this.InputInfo[index]) >> 1)
        }
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *HelpPaintable) SetInputInfoP(inputInfo []string) {
        this.InputInfo = inputInfo
        this.BeginWidthArray = make([]int, len(this.InputInfo))
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *HelpPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var halfWidth int = this.DisplayInfo.GetLastHalfWidth(); _ = halfWidth
        graphics.SetColor(this.BasicColor.IntValue())
        graphics.DrawString(this.Title, halfWidth - this.TitleBeginWidth, this.CharHeight, this.Anchor)
        var size int = len(this.InputInfo); _ = size
        for index := int(0); index < size; index++ {
            graphics.DrawString(this.InputInfo[index], halfWidth - this.BeginWidthArray[index], (index + 3) * this.CharHeight, this.Anchor)
        }
}

