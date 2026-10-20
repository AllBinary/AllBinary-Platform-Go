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

type ImageColorRangeResults struct {
    ColorRangeInterface ColorRangeInterface
    MatchingPixelsChecked int64
    TotalPixelsChecked int64
}


func (this *ImageColorRangeResults) initFields() {
    
    
    
}


func NewImageColorRangeResults() *ImageColorRangeResults {
    this := &ImageColorRangeResults{}
    this.initFields()
        this.SetColorRangeInterface(any(NewColorRange()).(ColorRangeInterface))
    return this
}


func (this *ImageColorRangeResults) GetColorRangeInterface() ColorRangeInterface {
        return this.ColorRangeInterface
}


func (this *ImageColorRangeResults) SetColorRangeInterface(colorRangeInterface ColorRangeInterface) {
        this.ColorRangeInterface = colorRangeInterface
}


func (this *ImageColorRangeResults) GetMatchingPixelsChecked() int64 {
        return this.MatchingPixelsChecked
}


func (this *ImageColorRangeResults) SetMatchingPixelsChecked(matchingPixelsChecked int64) {
        this.MatchingPixelsChecked = matchingPixelsChecked
}


func (this *ImageColorRangeResults) AddMatchingPixelsChecked() {
        this.MatchingPixelsChecked++
}


func (this *ImageColorRangeResults) GetTotalPixelsChecked() int64 {
        return this.TotalPixelsChecked
}


func (this *ImageColorRangeResults) AddTotalPixelsChecked() {
        this.TotalPixelsChecked++
}


func (this *ImageColorRangeResults) SetTotalPixelsChecked(totalPixelsChecked int64) {
        this.TotalPixelsChecked = totalPixelsChecked
}


func (this *ImageColorRangeResults) GetPercent() float32 {
        return (float32(this.GetMatchingPixelsChecked()) / float32(this.GetTotalPixelsChecked()))
}

