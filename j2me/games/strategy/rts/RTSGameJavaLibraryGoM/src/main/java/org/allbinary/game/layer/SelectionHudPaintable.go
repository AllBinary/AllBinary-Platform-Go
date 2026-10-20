/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type SelectionHudPaintable struct {
    InitUpdatePaintable
    Y int
    PrimitiveLongUtil *PrimitiveLongUtil
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    X int
    TextX int
    ImageX int
    BasicColor *BasicColor
    Color int
    Width int
    Height int
    Name string
    AnimationInterface IAnimation
    BackgroundColor int
}

// Implements UpdateMyFontInterface

// The Java class SelectionHudPaintable has subclasses so its variables are this interface
type ISelectionHudPaintable interface {
    IInitUpdatePaintable
    SelectionHudPaintable_() *SelectionHudPaintable
    UpdateMeasurement(graphics IGraphics)
    Update()
    UpdateSelectionInfo()
    UpdateInfo()
    GetPrimitiveLongUtil() *PrimitiveLongUtil
    SetAnimationInterface(animationInterface IAnimation)
    GetAnimationInterface() IAnimation
    SetName(name string)
    GetName() string
    SetBasicColorP(basicColor *BasicColor)
    GetBasicColorP() *BasicColor
    SetColor(color int)
    GetColor() int
    GetWidth() int
    SetHeight(height int)
    GetHeight() int
    GetX() int
    Paint(graphics IGraphics)
}

func (this *SelectionHudPaintable) SelectionHudPaintable_() *SelectionHudPaintable {
    return this
}


func (this *SelectionHudPaintable) initFields() {
    this.Y = CommonButtons_GetInstance().STANDARD_BUTTON_SIZE + 17
    
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    
    
    
    this.BasicColor = BasicColorFactory_GetInstance().BLACK
    
    
    
    this.Name = StringUtil_GetInstance().EMPTY_STRING
    this.AnimationInterface = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
    this.BackgroundColor = BasicColorFactory_GetInstance().GREY.IntValue()
}


func NewSelectionHudPaintable() *SelectionHudPaintable {
    this := &SelectionHudPaintable{}
    this.InitUpdatePaintable = *NewInitUpdatePaintable()
    this.initFields()
        this.Update()
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(10000)
    return this
}


func (this *SelectionHudPaintable) UpdateMeasurement(graphics IGraphics) {
        var commonButtons *CommonButtons = CommonButtons_GetInstance(); _ = commonButtons
        var font Font = graphics.GetFont(); _ = font
        this.SetHeight(commonButtons.STANDARD_BUTTON_SIZE + font.GetHeight())
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *SelectionHudPaintable) Update() {
        var commonButtons *CommonButtons = CommonButtons_GetInstance(); _ = commonButtons
        var touchButtonLocationHelper *TouchButtonLocationHelper = NewTouchButtonLocationHelper(); _ = touchButtonLocationHelper
        var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
        this.X = commonButtons.STANDARD_BUTTON_SIZE + touchButtonLocationHelper.GetColumnsRemainderHalf()
        this.TextX = this.GetX() + 4
        this.Width = displayInfoSingleton.GetLastWidth() - this.GetX() * 2
        this.ImageX = this.GetWidth() + touchButtonLocationHelper.GetColumnsRemainderHalf() - 10
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *SelectionHudPaintable) UpdateSelectionInfo() {
}


func (this *SelectionHudPaintable) UpdateInfo() {
}


func (this *SelectionHudPaintable) GetPrimitiveLongUtil() *PrimitiveLongUtil {
        return this.PrimitiveLongUtil
}


func (this *SelectionHudPaintable) SetAnimationInterface(animationInterface IAnimation) {
        this.AnimationInterface = animationInterface
}


func (this *SelectionHudPaintable) GetAnimationInterface() IAnimation {
        return this.AnimationInterface
}


func (this *SelectionHudPaintable) SetName(name string) {
        this.Name = name
}


func (this *SelectionHudPaintable) GetName() string {
        return this.Name
}


func (this *SelectionHudPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
        this.SetColor(basicColor.IntValue())
}


func (this *SelectionHudPaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *SelectionHudPaintable) SetColor(color int) {
        this.Color = color
}


func (this *SelectionHudPaintable) GetColor() int {
        return this.Color
}


func (this *SelectionHudPaintable) GetWidth() int {
        return this.Width
}


func (this *SelectionHudPaintable) SetHeight(height int) {
        this.Height = height
}


func (this *SelectionHudPaintable) GetHeight() int {
        return this.Height
}


func (this *SelectionHudPaintable) GetX() int {
        return this.X
}


func (this *SelectionHudPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        graphics.SetColor(this.BackgroundColor)
        graphics.DrawRect(this.GetX(), this.Y, this.GetWidth(), this.GetHeight())
        graphics.SetColor(this.GetColor())
        graphics.DrawString(this.GetName(), this.TextX, this.Y, 0)
}

