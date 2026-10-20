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

type ColorDelta struct {
    Rgb1 int
    Rgb2 int
    Key any
}

// Implements CacheableInterface


func (this *ColorDelta) initFields() {
    
    
    
}


func NewColorDelta(rgb1 int, rgb2 int) *ColorDelta {
    this := &ColorDelta{}
    this.initFields()
        this.SetRgb1(rgb1)
        this.SetRgb2(rgb2)
        this.Key = ColorDelta_GetKey(rgb1, rgb2)
    return this
}


func (this *ColorDelta) GetKey() any {
        return this.Key
}


func ColorDelta_GetKey(rgb1 int, rgb2 int) any {
        return fmt.Sprint(NewStringMaker().Append(fmt.Sprint(rgb1)).Append(CommonSeps_GetInstance().UNDERSCORE).Append(fmt.Sprint(rgb2)))
}


func (this *ColorDelta) GetRgb1() int {
        return this.Rgb1
}


func (this *ColorDelta) SetRgb1(rgb1 int) {
        this.Rgb1 = rgb1
}


func (this *ColorDelta) GetRgb2() int {
        return this.Rgb2
}


func (this *ColorDelta) SetRgb2(rgb2 int) {
        this.Rgb2 = rgb2
}


func (this *ColorDelta) ToString() string {
        return fmt.Sprint(NewStringMaker().Append("ColorDelta: ").Append(StringUtil_GetInstance().ToString(this.GetKey())).Append(" RGB1: ").Appendint(this.GetRgb1()).Append(" RGB2: ").Appendint(this.GetRgb2()))
}

func (this *ColorDelta) String() string {
    return this.ToString()
}

