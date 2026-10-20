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

//super.paint(graphics, positionString, currentPostString, offset2, offset);
type PositionWidget struct {
    BasicHud
    Position int
    PositionString string
}


func (this *PositionWidget) initFields() {
    
    this.PositionString = StringUtil_GetInstance().EMPTY_STRING
}


func NewPositionWidget(position int, location int, direction int, basicColor *BasicColor) *PositionWidget {
    this := &PositionWidget{}
        this.BasicHud = *NewBasicHud(location, direction, 2, basicColor)
        this.initFields()
        this.UpdateMaxHeight = 10
        this.UpdateMaxWidth = 40
    return this
}


func (this *PositionWidget) Get() int {
        return this.Position
}


func (this *PositionWidget) Set(position int) {
        this.Position = position
        this.PositionString = PositionWidget_ToString(this.Get())
}


func (this *PositionWidget) Paint(graphics IGraphics) {
        this.BasicHud.PaintString(graphics, this.PositionString)
}


func PositionWidget_ToString(position int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(position))
        if position == 1 {
            stringBuffer.Append("st")
        } else if position == 2 {
            stringBuffer.Append("nd")
        } else if position == 3 {
            stringBuffer.Append("rd")
        } else {
            stringBuffer.Append("th")
        }
        return fmt.Sprint(stringBuffer)
}

