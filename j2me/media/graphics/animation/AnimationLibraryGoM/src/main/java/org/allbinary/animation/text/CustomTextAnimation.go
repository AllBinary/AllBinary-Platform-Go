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

type CustomTextAnimation struct {
    TextAnimation
    FontDebugFactory *FontDebugFactory
    FontSize int
    Font Font
    LastText string
    HasChanged bool
    TextWidth int
}

// Implements TextInterface

// The Java class CustomTextAnimation has subclasses so its variables are this interface
type ICustomTextAnimation interface {
    ITextAnimation
    CustomTextAnimation_() *CustomTextAnimation
    PaintXY(graphics IGraphics, x int, y int)
    SetScale(scaleX float32, scaleY float32)
    SetText(text string)
    GetText() string
    GetWidth() int
    GetFontHeight() int
}

func (this *CustomTextAnimation) CustomTextAnimation_() *CustomTextAnimation {
    return this
}


func (this *CustomTextAnimation) initFields() {
    this.FontDebugFactory = FontDebugFactory_GetInstance()
    
    
    this.LastText = StringUtil_GetInstance().INIT_STRING
    this.HasChanged = true
    
}


func NewCustomTextAnimation(text string, fontSize int, animationBehavior IAnimationBehavior) *CustomTextAnimation {
    this := &CustomTextAnimation{}
        this.TextAnimation = *NewTextAnimation(text, animationBehavior)
        this.initFields()
        this.FontSize = fontSize
        this.Font = Font_GetFont(Font_FACE_SYSTEM, Font_STYLE_PLAIN, fontSize)
    return this
}


func (this *CustomTextAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        var existingFont Font = graphics.GetFont(); _ = existingFont
        this.FontDebugFactory.SetFont(this.Font, graphics)
        this.TextAnimation.PaintXYNoUpdate(graphics, x, y)
        this.FontDebugFactory.SetFont(existingFont, graphics)
}


func (this *CustomTextAnimation) SetScale(scaleX float32, scaleY float32) {
        if scaleX != scaleY {
            panic(errors.New(fmt.Sprint()))
        }
        this.HasChanged = true
        this.Font = Font_GetFont(Font_FACE_SYSTEM, Font_STYLE_PLAIN, int((float32(this.FontSize) * scaleX)))
}


func (this *CustomTextAnimation) SetText(text string) {
        if this.LastText != text {
            this.HasChanged = true
            this.LastText = text
            this.TextAnimation.SetText(text)
        }
}


func (this *CustomTextAnimation) GetText() string {
        return this.LastText
}


func (this *CustomTextAnimation) GetWidth() int {
        if this.HasChanged {
            this.TextWidth = this.Font.StringWidth(this.TextArrayP[0])
            this.HasChanged = false
        }
        return this.TextWidth
}


func (this *CustomTextAnimation) GetFontHeight() int {
        return this.Font.GetHeight()
}

