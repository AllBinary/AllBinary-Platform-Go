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

//SWTToJ2ME
type AboutPaintable struct {
    Paintable
    DisplayInfoSingleton *DisplayInfoSingleton
    ABOUT string
    Info []string
    Developers []string
    PaintableArray []IPaintable
    MyFontProcessor IMyFontProcessor
    CharHeight int
    AboutBeginWidth int
    InfoBeginWidth []int
    DevelopersBeginWidth []int
    Anchor int
}

// Implements UpdateMyFontInterface


func (this *AboutPaintable) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.ABOUT = CanvasStrings_GetInstance().ABOUT
    
    
    this.PaintableArray = []IPaintable{this}
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    
    
    
    
    this.Anchor = Anchor_TOP_LEFT
}


func AboutPaintable_GetInstance(info []string, developers []string) *AboutPaintable {
        return NewAboutPaintable(info, developers)
}


func NewAboutPaintable(info []string, developers []string) *AboutPaintable {
    this := &AboutPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.Info = info
        this.InfoBeginWidth = make([]int, len(this.Info))
        this.Developers = developers
        this.DevelopersBeginWidth = make([]int, len(this.Developers))
    return this
}


func (this *AboutPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.CharHeight = font.GetHeight()
        this.AboutBeginWidth = (font.StringWidth(this.ABOUT) >> 1)
        var infoSize int = len(this.Info); _ = infoSize
        for index := int(0); index < infoSize; index++ {
            this.InfoBeginWidth[index] = (font.StringWidth(this.Info[index]) >> 1)
        }
        var size int = len(this.Developers); _ = size
        for index := int(0); index < size; index++ {
            this.DevelopersBeginWidth[index] = (font.StringWidth(this.Developers[index]) >> 1)
        }
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *AboutPaintable) GetPaintableArrayInstance() []IPaintable {
        return this.PaintableArray
}


func (this *AboutPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var halfWidth int = this.DisplayInfoSingleton.GetLastHalfWidth(); _ = halfWidth
        graphics.DrawString(this.ABOUT, halfWidth - this.AboutBeginWidth, 2 * this.CharHeight, this.Anchor)
        var infoSize int = len(this.Info); _ = infoSize
        for index := int(0); index < infoSize; index++ {
            graphics.DrawString(this.Info[index], halfWidth - this.InfoBeginWidth[index], (4 + index) * this.CharHeight, this.Anchor)
        }
        var size int = len(this.Developers); _ = size
        for index := int(0); index < size; index++ {
            graphics.DrawString(this.Developers[index], halfWidth - this.DevelopersBeginWidth[index], (5 + infoSize + index) * this.CharHeight, this.Anchor)
        }
}

