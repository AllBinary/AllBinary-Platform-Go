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

type BasicColor struct {
    Name string
    Value int
    RedComponent float32
    GreenComponent float32
    BlueComponent float32
    AlphaComponent float32
    Red int16
    Green int16
    Blue int16
    Alpha int16
}


func (this *BasicColor) initFields() {
    
    
    
    
    
    
    
    
    
    
}


func NewBasicColor(alphaValue int, value int, name string) *BasicColor {
    this := &BasicColor{}
    this.initFields()
        this.Name = name
        var tempValue int; _ = tempValue
        var basicColorUtil *BasicColorUtil = BasicColorUtil_GetInstance(); _ = basicColorUtil
        if basicColorUtil.IsAlpha {
            if basicColorUtil.FfOpaque {
                tempValue = alphaValue | value
            } else {
                tempValue = 0x00000000 | value
            }
        } else {
            tempValue = value
        }
        var alphaInt int = ((tempValue >> 24) & 255); _ = alphaInt
        this.Alpha = int16(alphaInt)
        this.AlphaComponent = (float32(alphaInt)) / 255.0
        var redInt int = ((tempValue >> 16) & 255); _ = redInt
        this.Red = int16(redInt)
        this.RedComponent = (float32(redInt)) / 255.0
        var greenInt int = ((tempValue >> 8) & 255); _ = greenInt
        this.Green = int16(greenInt)
        this.GreenComponent = (float32(greenInt)) / 255.0
        var blueInt int = ((tempValue) & 255); _ = blueInt
        this.Blue = int16(blueInt)
        this.BlueComponent = (float32(blueInt)) / 255.0
        this.Value = tempValue
    return this
}


func (this *BasicColor) IntValue() int {
        return this.Value
}


func (this *BasicColor) ToInt() int {
        return this.Value
}


func (this *BasicColor) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("BasicColor: ")
        stringBuffer.Append("Alpha: ")
        stringBuffer.Appendshort(this.Alpha)
        stringBuffer.Append(" Red: ")
        stringBuffer.Appendshort(this.Red)
        stringBuffer.Append(" Green: ")
        stringBuffer.Appendshort(this.Green)
        stringBuffer.Append(" Blue: ")
        stringBuffer.Appendshort(this.Blue)
        stringBuffer.Append(" Value: ")
        stringBuffer.Appendint(this.Value)
        return fmt.Sprint(stringBuffer)
}

func (this *BasicColor) String() string {
    return this.ToString()
}


func (this *BasicColor) GetName() string {
        return this.Name
}


func (this *BasicColor) GetRedComponent() float32 {
        return this.RedComponent
}


func (this *BasicColor) GetGreenComponent() float32 {
        return this.GreenComponent
}


func (this *BasicColor) GetBlueComponent() float32 {
        return this.BlueComponent
}


func (this *BasicColor) GetAlphaComponent() float32 {
        return this.AlphaComponent
}
