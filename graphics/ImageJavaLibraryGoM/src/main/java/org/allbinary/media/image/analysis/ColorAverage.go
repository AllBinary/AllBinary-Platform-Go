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

type ColorAverage struct {
    AvgRed float32
    AvgGreen float32
    AvgBlue float32
}


func (this *ColorAverage) initFields() {
    
    
    
}


func NewColorAverage() *ColorAverage {
    this := &ColorAverage{}
    this.initFields()
        this.SetAvgRed(0)
        this.SetAvgGreen(0)
        this.SetAvgBlue(0)
    return this
}


func (this *ColorAverage) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("getAvgRed(): ")
        stringBuffer.Appendfloat(this.GetAvgRed())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getAvgRed(): ")
        stringBuffer.Appendfloat(this.GetAvgRed())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getAvgGreen(): ")
        stringBuffer.Appendfloat(this.GetAvgGreen())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        return fmt.Sprint(stringBuffer)
}

func (this *ColorAverage) String() string {
    return this.ToString()
}


func (this *ColorAverage) GetAvgRed() float32 {
        return this.AvgRed
}


func (this *ColorAverage) SetAvgRed(avgRed float32) {
        this.AvgRed = avgRed
}


func (this *ColorAverage) GetAvgGreen() float32 {
        return this.AvgGreen
}


func (this *ColorAverage) SetAvgGreen(avgGreen float32) {
        this.AvgGreen = avgGreen
}


func (this *ColorAverage) GetAvgBlue() float32 {
        return this.AvgBlue
}


func (this *ColorAverage) SetAvgBlue(avgBlue float32) {
        this.AvgBlue = avgBlue
}

