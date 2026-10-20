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

type CommandCurrentSelectionForm struct {
    ScrollCurrentSelectionForm
    SelectedAnimationArray []IAnimation
    UnSelectedAnimationArray []IAnimation
}

// Implements UpdateMyFontInterface


func (this *CommandCurrentSelectionForm) initFields() {
    this.SelectedAnimationArray = make([]IAnimation, 16)
    this.UnSelectedAnimationArray = make([]IAnimation, 16)
}


func NewCommandCurrentSelectionForm(title string, items []IABCustomItem, border int, moveForSmallScreen bool, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *CommandCurrentSelectionForm {
    this := &CommandCurrentSelectionForm{}
        this.ScrollCurrentSelectionForm = *NewScrollCurrentSelectionForm(title, items, ItemPaintableFactory_GetInstance(), border, -3, moveForSmallScreen, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.InitAnimations()
        this.AddAll(items)
    return this
}


func (this *CommandCurrentSelectionForm) UpdateMeasurement(graphics IGraphics) {
        this.UpdateAll(graphics, this.GetAllitems())
        this.ScrollCurrentSelectionForm.UpdateMeasurement(graphics)
}


func (this *CommandCurrentSelectionForm) InitAnimations() {
        var nullAnimation IAnimation = NullAnimationFactory_GetFactoryInstance().GetInstance(0); _ = nullAnimation
        for index := int(len(this.SelectedAnimationArray)); Java_PreDecrement(&index) >= 0;  {
            this.SelectedAnimationArray[index] = nullAnimation
        }
        for index := int(len(this.UnSelectedAnimationArray)); Java_PreDecrement(&index) >= 0;  {
            this.UnSelectedAnimationArray[index] = nullAnimation
        }
}


func (this *CommandCurrentSelectionForm) AddAll(items []IABCustomItem) {
        for index := int(len(items)); Java_PreDecrement(&index) >= 0;  {
            this.AddAt(index, items[index])
        }
}


func (this *CommandCurrentSelectionForm) AddAt(index int, item IABCustomItem) {
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        var buttonColor *BasicColor = basicColorFactory.TRANSPARENT_GREY; _ = buttonColor
        var selectedButtonColor *BasicColor = basicColorFactory.TRANSPARENT_RED; _ = selectedButtonColor
        var width int = item.GetMinimumWidth(); _ = width
        var height int = item.GetMinimumHeight(); _ = height
        var adjustedBorder int = 3; _ = adjustedBorder
        var offset int = -(this.HalfBorder + adjustedBorder); _ = offset
        var rectWidth int = width + this.Border - adjustedBorder; _ = rectWidth
        var rectHeight int = height + this.Border - adjustedBorder; _ = rectHeight
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
            this.SelectedAnimationArray[index] = NewRectangleAdjustedAnimation(rectWidth, rectHeight, offset, offset, selectedButtonColor)
        } else {
            this.SelectedAnimationArray[index] = NewRectangleFilledAdjustedAnimation(rectWidth, rectHeight, offset, offset, selectedButtonColor)
        }
        adjustedBorder = 4
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
            this.UnSelectedAnimationArray[index] = NewRectangleAdjustedAnimation(rectWidth, rectHeight, offset, offset, buttonColor)
        } else {
            this.UnSelectedAnimationArray[index] = NewRectangleFilledAdjustedAnimation(rectWidth, rectHeight, offset, offset, buttonColor)
        }
}


func (this *CommandCurrentSelectionForm) UpdateAll(graphics IGraphics, items []IABCustomItem) {
        for index := int(len(items)); Java_PreDecrement(&index) >= 0;  {
            items[index].PreMeasurement(graphics)
            this.UpdateAt(index, items[index])
        }
}


func (this *CommandCurrentSelectionForm) UpdateAt(index int, item IABCustomItem) {
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        var buttonColor *BasicColor = basicColorFactory.TRANSPARENT_GREY; _ = buttonColor
        var selectedButtonColor *BasicColor = basicColorFactory.TRANSPARENT_RED; _ = selectedButtonColor
        var width int = item.GetMinimumWidth(); _ = width
        var height int = item.GetMinimumHeight(); _ = height
        var adjustedBorder int = 3; _ = adjustedBorder
        var offset int = -(this.HalfBorder + adjustedBorder); _ = offset
        var rectWidth int = width + this.Border - adjustedBorder; _ = rectWidth
        var rectHeight int = height + this.Border - adjustedBorder; _ = rectHeight
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
            var rectangleAdjustedAnimation *RectangleAdjustedAnimation = any(this.SelectedAnimationArray[index]).(*RectangleAdjustedAnimation); _ = rectangleAdjustedAnimation
            rectangleAdjustedAnimation.SetWidth(rectWidth)
            rectangleAdjustedAnimation.SetHeight(rectHeight)
            rectangleAdjustedAnimation.SetOffsetX(offset)
            rectangleAdjustedAnimation.SetOffsetY(offset)
            rectangleAdjustedAnimation.SetBasicColorP(selectedButtonColor)
        } else {
            var rectangleAdjustedAnimation *RectangleFilledAdjustedAnimation = any(this.SelectedAnimationArray[index]).(*RectangleFilledAdjustedAnimation); _ = rectangleAdjustedAnimation
            rectangleAdjustedAnimation.SetWidth(rectWidth)
            rectangleAdjustedAnimation.SetHeight(rectHeight)
            rectangleAdjustedAnimation.SetOffsetX(offset)
            rectangleAdjustedAnimation.SetOffsetY(offset)
            rectangleAdjustedAnimation.SetBasicColorP(selectedButtonColor)
        }
        adjustedBorder = 4
        if J2MEUtil_IsJ2ME() || J2MEUtil_IsHTML() {
            var rectangleAdjustedAnimation *RectangleAdjustedAnimation = any(this.UnSelectedAnimationArray[index]).(*RectangleAdjustedAnimation); _ = rectangleAdjustedAnimation
            rectangleAdjustedAnimation.SetWidth(rectWidth)
            rectangleAdjustedAnimation.SetHeight(rectHeight)
            rectangleAdjustedAnimation.SetOffsetX(offset)
            rectangleAdjustedAnimation.SetOffsetY(offset)
            rectangleAdjustedAnimation.SetBasicColorP(buttonColor)
        } else {
            var rectangleAdjustedAnimation *RectangleFilledAdjustedAnimation = any(this.UnSelectedAnimationArray[index]).(*RectangleFilledAdjustedAnimation); _ = rectangleAdjustedAnimation
            rectangleAdjustedAnimation.SetWidth(rectWidth)
            rectangleAdjustedAnimation.SetHeight(rectHeight)
            rectangleAdjustedAnimation.SetOffsetX(offset)
            rectangleAdjustedAnimation.SetOffsetY(offset)
            rectangleAdjustedAnimation.SetBasicColorP(buttonColor)
        }
}


func (this *CommandCurrentSelectionForm) GetSelectedCommand() ICommand {
        var index int = this.ScrollCurrentSelectionForm.GetSelectedIndex(); _ = index
        var commandTextItem *CommandTextItem = any(this.Get(index)).(*CommandTextItem); _ = commandTextItem
        return commandTextItem.GetCommand()
}


func (this *CommandCurrentSelectionForm) Append(item IABCustomItem) int {
        var result int = this.ScrollCurrentSelectionForm.Append(item); _ = result
        this.AddAt(result, item)
        return result
}


func (this *CommandCurrentSelectionForm) Delete(itemNum int) {
        this.ScrollCurrentSelectionForm.Delete(itemNum)
}


func (this *CommandCurrentSelectionForm) DeleteAll() {
        this.InitAnimations()
        this.ScrollCurrentSelectionForm.DeleteAll()
}


func (this *CommandCurrentSelectionForm) Insert(itemNum int, item IABCustomItem) {
        this.ScrollCurrentSelectionForm.Insert(itemNum, item)
}


func (this *CommandCurrentSelectionForm) Set(itemNum int, item IABCustomItem) {
        this.ScrollCurrentSelectionForm.Set(itemNum, item)
}


func (this *CommandCurrentSelectionForm) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        this.ScrollCurrentSelectionForm.Paint(graphics)
}


// throws Exception
func (this *CommandCurrentSelectionForm) PaintItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        this.SelectedAnimationArray[index].PaintXY(graphics, x, y)
        return this.ScrollCurrentSelectionForm.PaintItem(graphics, index, item, x, y)
}


// throws Exception
func (this *CommandCurrentSelectionForm) PaintUnselectedItem(graphics IGraphics, index int, item IABCustomItem, x int, y int) int {
        this.UnSelectedAnimationArray[index].PaintXY(graphics, x, y)
        return this.ScrollCurrentSelectionForm.PaintUnselectedItem(graphics, index, item, x, y)
}

