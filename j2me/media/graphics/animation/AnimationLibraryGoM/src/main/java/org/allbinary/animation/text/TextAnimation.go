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

type TextAnimation struct {
    IndexedAnimation
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    TextArrayP []string
    Anchor int
    FontHeight int
    OffsetY int
    TextChangeListener *TextChangeListener
}

// Implements UpdateMyFontInterface

// The Java class TextAnimation has subclasses so its variables are this interface
type ITextAnimation interface {
    IIndexedAnimation
    TextAnimation_() *TextAnimation
    UpdateMeasurement(graphics IGraphics)
    NextFrame()
    PaintXY(graphics IGraphics, x int, y int)
    PaintXYNoUpdate(graphics IGraphics, x int, y int)
    SetTextWithOnMeasure(text string, textChangeListener *TextChangeListener)
    SetText(text string)
    GetTextArray() []string
    GetFontHeight() int
}

func (this *TextAnimation) TextAnimation_() *TextAnimation {
    return this
}


func (this *TextAnimation) initFields() {
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.TextArrayP = StringUtil_GetInstance().ONE_EMPTY_STRING_ARRAY
    this.Anchor = Anchor_TOP_LEFT
    this.FontHeight = 0
    this.OffsetY = 0
    this.TextChangeListener = TextChangeListener_GetInstance()
}


func NewTextAnimation(text string, animationBehavior IAnimationBehavior) *TextAnimation {
    this := &TextAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.SetText(text)
    return this
}


func (this *TextAnimation) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        var adjustedFontHeight int = int(-(float64(this.FontHeight) * 2.00)); _ = adjustedFontHeight
        this.OffsetY = func() int { if SWTJOGLProcessor_GetInstance().IsJOGL() { return adjustedFontHeight }; return 0 }()
        this.TextChangeListener.OnMeasure()
        this.TextChangeListener = TextChangeListener_GetInstance()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


// throws Exception
func (this *TextAnimation) NextFrame() {
}


func (this *TextAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        this.PaintXYNoUpdate(graphics, x, y)
}


func (this *TextAnimation) PaintXYNoUpdate(graphics IGraphics, x int, y int) {
        this.BasicSetColorUtil.SetBasicColorP3(graphics, this.GetBasicColorP(), this.GetColor())
        var size int = len(this.TextArrayP); _ = size
        for index := int(0); index < size; index++ {
            graphics.DrawString(this.TextArrayP[index], x, y + this.OffsetY + (index * this.FontHeight), this.Anchor)
        }
}


func (this *TextAnimation) SetTextWithOnMeasure(text string, textChangeListener *TextChangeListener) {
        this.SetText(text)
        this.TextChangeListener = textChangeListener
}


func (this *TextAnimation) SetText(text string) {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        if text != "" && len(text) > 0 {
            var index int = 0; _ = index
            var startIndex int = 0; _ = startIndex
            var endIndex int = 0; _ = endIndex
            for index >= 0 {
                startIndex = index
                index = text.IndexOf('\n', startIndex)
                endIndex = index
                if index < 0 {
                    endIndex = len(text)
                }
                list.Add(text[startIndex:endIndex])
                if index < 0 {
                    break
                }
                index++
            }
        }
        if list.Size() > 0 {
            var textArray []string = make([]string, list.Size()); _ = textArray
            var size int = list.Size(); _ = size
            for index := int(0); index < size; index++ {
                textArray[index] = any(list.Get(index)).(string)
            }
            this.TextArrayP = textArray
        } else {
            this.TextArrayP = StringUtil_GetInstance().ONE_EMPTY_STRING_ARRAY
        }
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *TextAnimation) GetTextArray() []string {
        return this.TextArrayP
}


func (this *TextAnimation) GetFontHeight() int {
        return this.FontHeight
}

