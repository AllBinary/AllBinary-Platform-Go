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

type ColorRange struct {
    MinRed int
    MaxRed int
    MinGreen int
    MaxGreen int
    MinBlue int
    MaxBlue int
}

// Implements ColorRangeInterface


func (this *ColorRange) initFields() {
    
    
    
    
    
    
}


func NewColorRange() *ColorRange {
    this := &ColorRange{}
    this.initFields()
        this.MinRed = 0
        this.MaxRed = 0
        this.MinGreen = 0
        this.MaxGreen = 0
        this.MinBlue = 0
        this.MaxBlue = 0
    return this
}


func (this *ColorRange) GetMinRed() int {
        return this.MinRed
}


func (this *ColorRange) SetMinRed(minRed int) {
        this.MinRed = minRed
}


func (this *ColorRange) GetMaxRed() int {
        return this.MaxRed
}


func (this *ColorRange) SetMaxRed(maxRed int) {
        this.MaxRed = maxRed
}


func (this *ColorRange) GetMinGreen() int {
        return this.MinGreen
}


func (this *ColorRange) SetMinGreen(minGreen int) {
        this.MinGreen = minGreen
}


func (this *ColorRange) GetMaxGreen() int {
        return this.MaxGreen
}


func (this *ColorRange) SetMaxGreen(maxGreen int) {
        this.MaxGreen = maxGreen
}


func (this *ColorRange) GetMinBlue() int {
        return this.MinBlue
}


func (this *ColorRange) SetMinBlue(minBlue int) {
        this.MinBlue = minBlue
}


func (this *ColorRange) GetMaxBlue() int {
        return this.MaxBlue
}


func (this *ColorRange) SetMaxBlue(maxBlue int) {
        this.MaxBlue = maxBlue
}


func (this *ColorRange) IsInRange(color *Color) bool {
        if color.GetRed() >= this.GetMinRed() && color.GetRed() <= this.GetMaxRed() && color.GetGreen() >= this.GetMinGreen() && color.GetGreen() <= this.GetMaxGreen() && color.GetBlue() >= this.GetMinBlue() && color.GetBlue() <= this.GetMaxBlue() {
            return true
        } else {
            return false
        }
}


func (this *ColorRange) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("getMinRed(): ")
        stringBuffer.Appendint(this.GetMinRed())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getMaxRed(): ")
        stringBuffer.Appendint(this.GetMaxRed())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getMinGreen(): ")
        stringBuffer.Appendint(this.GetMinGreen())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getMaxGreen(): ")
        stringBuffer.Appendint(this.GetMaxGreen())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getMinBlue(): ")
        stringBuffer.Appendint(this.GetMinGreen())
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("getMaxBlue(): ")
        stringBuffer.Appendint(this.GetMaxBlue())
        return fmt.Sprint(stringBuffer)
}

func (this *ColorRange) String() string {
    return this.ToString()
}

