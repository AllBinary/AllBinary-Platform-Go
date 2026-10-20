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

type CustomTextBox struct {
    GameCommandCanvas
    TextFieldItem *ABTextFieldItem
    InputFactory *InputFactory
}

// Implements RawKeyEventListener

// The Java class CustomTextBox has subclasses so its variables are this interface
type ICustomTextBox interface {
    IGameCommandCanvas
    CustomTextBox_() *CustomTextBox
    Submit()
    OnEventRaw(keyCode int, deviceId int, repeated bool)
    KeyPressed(keyCode int)
    KeyReleased(keyCode int)
    KeyRepeated(keyCode int)
    KeyPressedByDevice(keyCode int, deviceId int)
    KeyReleasedByDevice(keyCode int, deviceId int)
    Paint(graphics IGraphics)
    PaintXY(graphics IGraphics, x int, y int)
    GetTextFieldItem() *ABTextFieldItem
}

func (this *CustomTextBox) CustomTextBox_() *CustomTextBox {
    return this
}


func (this *CustomTextBox) initFields() {
    
    this.InputFactory = InputFactory_GetInstance()
}


func NewCustomTextBox(cmdListener CommandListener, label string, text string, maxSize int, constraints int, font Font, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *CustomTextBox {
    this := &CustomTextBox{}
        this.GameCommandCanvas = *NewGameCommandCanvas(cmdListener, label, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var textFieldItem *ABTextFieldItem = NewABTextFieldItem(this, NewTextItemVisitor(), stringUtil.EMPTY_STRING, stringUtil.EMPTY_STRING, maxSize, 0, stringUtil.EMPTY_STRING, font, backgroundBasicColor, foregroundBasicColor); _ = textFieldItem
        textFieldItem.SetString(text)
        this.TextFieldItem = textFieldItem
        this.SetTitle(label)
    return this
}


func (this *CustomTextBox) Submit() {
}


func (this *CustomTextBox) OnEventRaw(keyCode int, deviceId int, repeated bool) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Appendint(keyCode)), this, "onEvent")
        this.KeyPressedByDevice(keyCode, deviceId)
}


func (this *CustomTextBox) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *CustomTextBox) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *CustomTextBox) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *CustomTextBox) KeyPressedByDevice(keyCode int, deviceId int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.KEY_PRESSED, e)
                }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Appendint(keyCode)), this, this.GameInputStrings.KEY_PRESSED)
            var platformKeyFactory *PlatformKeyFactory = PlatformKeyFactory_GetInstance(); _ = platformKeyFactory
            var input IInput = this.InputFactory.GetInstanceById(keyCode); _ = input
            if platformKeyFactory.IsSubmission(input) {
                this.Submit()
            } else {
                this.TextFieldItem.KeyPressed(keyCode)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CustomTextBox) KeyReleasedByDevice(keyCode int, deviceId int) {
}


func (this *CustomTextBox) Paint(graphics IGraphics) {
        graphics.SetColor(this.BackgroundColor)
        graphics.FillRect(0, 0, this.DisplayInfo.GetLastWidth(), this.DisplayInfo.GetLastHeight())
        graphics.SetColor(this.ForegroundColor)
        if !J2MEUtil_IsJ2ME() {
            graphics.DrawString(this.GetTitle(), 8, 1, 0)
        }
        this.PaintXY(graphics, 8, graphics.GetFont().GetHeight() + 2)
        this.GameCommandCanvas.Paint(graphics)
}


func (this *CustomTextBox) PaintXY(graphics IGraphics, x int, y int) {
        this.TextFieldItem.PaintXY(graphics, x, y)
}


func (this *CustomTextBox) GetTextFieldItem() *ABTextFieldItem {
        return this.TextFieldItem
}

