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

type MultiSelectPaintable struct {
    SelectionHudPaintable
    RootNameList IBasicArrayList
    TotalCharArray []rune
    RootNamesString string
    TOTAL string
    BackgroundColor int
    TotalWidth int
    TextLine2Y int
}


func (this *MultiSelectPaintable) initFields() {
    this.RootNameList = NewBasicArrayListD()
    this.TotalCharArray = NullUtil_GetInstance().NULL_CHAR_ARRAY
    this.RootNamesString = StringUtil_GetInstance().EMPTY_STRING
    this.TOTAL = "Total Selected: "
    this.BackgroundColor = BasicColorFactory_GetInstance().GREY.IntValue()
    
    
}


func NewMultiSelectPaintable() *MultiSelectPaintable {
    this := &MultiSelectPaintable{}
    this.SelectionHudPaintable = *NewSelectionHudPaintable()
    this.initFields()
    return this
}


func (this *MultiSelectPaintable) UpdateMeasurement(graphics IGraphics) {
        this.SelectionHudPaintable.UpdateMeasurement(graphics)
        var font Font = graphics.GetFont(); _ = font
        this.TotalWidth = font.StringWidth(this.TOTAL)
        this.TextLine2Y = (this.Y + font.GetHeight())
}


func (this *MultiSelectPaintable) Update(list IBasicArrayList) {
        this.Clear()
        var size int = list.Size(); _ = size
        this.TotalCharArray = this.GetPrimitiveLongUtil().GetCharArray(size)
        for index := int(list.Size() - 1); index >= 0; index-- {
            var rtsLayer IRTSLayer = any(list.Get(index)).(IRTSLayer); _ = rtsLayer
            if !this.RootNameList.Contains(rtsLayer.GetRootName()) {
                this.RootNameList.Add(rtsLayer.GetRootName())
            }
        }
        var COMMA_SEP string = CommonSeps_GetInstance().COMMA_SEP; _ = COMMA_SEP
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(this.RootNameList.Size() - 1); index >= 0; index-- {
            var rootName string = any(this.RootNameList.Get(index)).(string); _ = rootName
            stringBuffer.Append(rootName)
            if index != 0 {
                stringBuffer.Append(COMMA_SEP)
            }
        }
        this.RootNamesString = fmt.Sprint(stringBuffer)
}


func (this *MultiSelectPaintable) Clear() {
        this.RootNameList.Clear()
}


func (this *MultiSelectPaintable) Paint(graphics IGraphics) {
        graphics.SetColor(this.BackgroundColor)
        graphics.DrawRect(this.GetX(), this.Y, this.GetWidth(), this.GetHeight())
        graphics.SetColor(this.GetColor())
        graphics.DrawString(this.TOTAL, this.TextX, this.Y, 0)
        graphics.DrawChars(this.TotalCharArray, 0, this.GetPrimitiveLongUtil().GetCurrentTotalDigits(), this.TextX + this.TotalWidth, this.Y, 0)
        graphics.DrawString(this.RootNamesString, this.TextX, this.TextLine2Y, 0)
}

