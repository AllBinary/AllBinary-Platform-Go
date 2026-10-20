/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type CustomTextBoxIndexedAnimation struct {
    IndexedAnimation
    CustomTextBox ICustomTextBox
    LastText string
    HasChanged bool
    TextChangeListener *TextChangeListener
}

// Implements RawKeyEventListener, UpdateMyFontInterface, TextInterface


func (this *CustomTextBoxIndexedAnimation) initFields() {
    
    this.LastText = StringUtil_GetInstance().INIT_STRING
    this.HasChanged = true
    this.TextChangeListener = TextChangeListener_GetInstance()
}


func NewCustomTextBoxIndexedAnimation(customTextBox ICustomTextBox) *CustomTextBoxIndexedAnimation {
    this := &CustomTextBoxIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(AnimationBehavior_GetInstance())
        this.initFields()
        this.CustomTextBox = customTextBox
    return this
}


func (this *CustomTextBoxIndexedAnimation) SetBasicColorP(basicColor *BasicColor) {
        this.CustomTextBox.GetTextFieldItem().SetForegroundBasicColorP(basicColor)
}


func (this *CustomTextBoxIndexedAnimation) SetBackgroundBasicColorP(basicColor *BasicColor) {
        this.CustomTextBox.GetTextFieldItem().SetBackgroundBasicColorP(basicColor)
}


func (this *CustomTextBoxIndexedAnimation) SetTextWithOnMeasure(text string, textChangeListener *TextChangeListener) {
        this.SetText(text)
        this.TextChangeListener = textChangeListener
}


func (this *CustomTextBoxIndexedAnimation) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.TextChangeListener.OnMeasure()
        this.TextChangeListener = TextChangeListener_GetInstance()
}


func (this *CustomTextBoxIndexedAnimation) SetText(text string) {
        if this.LastText != text {
            this.HasChanged = true
            this.LastText = text
            this.CustomTextBox.GetTextFieldItem().SetString(text)
        }
}


func (this *CustomTextBoxIndexedAnimation) GetText() string {
        return this.CustomTextBox.GetTextFieldItem().GetString()
}


func (this *CustomTextBoxIndexedAnimation) OnEventRaw(keyCode int, deviceId int, repeated bool) {
        this.CustomTextBox.OnEventRaw(keyCode, deviceId, repeated)
}


func (this *CustomTextBoxIndexedAnimation) GetFontHeight() int {
        return this.CustomTextBox.GetTextFieldItem().GetFontHeight()
}


func (this *CustomTextBoxIndexedAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.CustomTextBox.PaintXY(graphics, x, y)
}


func (this *CustomTextBoxIndexedAnimation) PaintThreedXYZ(graphics IGraphics, x int, y int, z int) {
}

