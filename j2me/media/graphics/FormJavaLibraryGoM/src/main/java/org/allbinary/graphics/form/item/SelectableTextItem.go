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

type SelectableTextItem struct {
    ABTextItem
}



func NewSelectableTextItem(label string, layout int, altText string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *SelectableTextItem {
    this := &SelectableTextItem{}
        this.ABTextItem = *NewABTextItem(label, layout, altText, backgroundBasicColor, foregroundBasicColor)
    return this
}


func (this *SelectableTextItem) IsFocusable() bool {
        return true
}


func (this *SelectableTextItem) PaintXY(graphics IGraphics, x int, y int) {
        this.MyFontProcessor.Process(graphics)
        var stringComponent *ABStringComponent = this.GetLabelStringComponent(); _ = stringComponent
        if stringComponent != nil {
            graphics.SetColor(stringComponent.GetForegroundBasicColor().IntValue())
        }
        graphics.DrawString(this.GetLabel(), x + 2, y, 0)
}

