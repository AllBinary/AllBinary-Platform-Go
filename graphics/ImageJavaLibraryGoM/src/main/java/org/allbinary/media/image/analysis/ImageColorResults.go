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

type ImageColorResults struct {
    ColorRange *ColorRange
    ColorAverage *ColorAverage
}


func (this *ImageColorResults) initFields() {
    
    
}


func NewImageColorResults() *ImageColorResults {
    this := &ImageColorResults{}
    this.initFields()
        this.SetColorRange(NewColorRange())
        this.GetColorRange().SetMaxRed(-1)
        this.GetColorRange().SetMaxGreen(-1)
        this.GetColorRange().SetMaxBlue(-1)
        this.GetColorRange().SetMinRed(256)
        this.GetColorRange().SetMinGreen(256)
        this.GetColorRange().SetMinBlue(256)
        this.SetColorAverage(NewColorAverage())
    return this
}


func (this *ImageColorResults) GetColorRange() *ColorRange {
        return this.ColorRange
}


func (this *ImageColorResults) SetColorRange(colorRange *ColorRange) {
        this.ColorRange = colorRange
}


func (this *ImageColorResults) GetColorAverage() *ColorAverage {
        return this.ColorAverage
}


func (this *ImageColorResults) SetColorAverage(colorAverage *ColorAverage) {
        this.ColorAverage = colorAverage
}

