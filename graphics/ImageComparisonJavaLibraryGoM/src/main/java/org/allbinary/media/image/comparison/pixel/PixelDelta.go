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

type PixelDelta struct {
    Point IGPoint
    ColorDelta *ColorDelta
    Key any
}

// Implements CacheableInterface


func (this *PixelDelta) initFields() {
    
    
    
}


func NewPixelDelta(point IGPoint, colorDelta *ColorDelta) *PixelDelta {
    this := &PixelDelta{}
    this.initFields()
        this.SetPoint(point)
        this.SetColorDelta(colorDelta)
        this.Key = PixelDelta_GetKey(this.GetPoint(), this.GetColorDelta())
    return this
}


func (this *PixelDelta) GetKey() any {
        return this.Key
}


func PixelDelta_GetKey(point IGPoint, colorDelta *ColorDelta) any {
        return fmt.Sprint(NewStringMaker().Appendint(TsUtil_GetInstance().HashCode(point)).Append(CommonSeps_GetInstance().UNDERSCORE).Append(fmt.Sprint(colorDelta.GetKey())))
}


func (this *PixelDelta) GetPoint() IGPoint {
        return this.Point
}


func (this *PixelDelta) SetPoint(point IGPoint) {
        this.Point = point
}


func (this *PixelDelta) GetColorDelta() *ColorDelta {
        return this.ColorDelta
}


func (this *PixelDelta) SetColorDelta(colorDelta *ColorDelta) {
        this.ColorDelta = colorDelta
}


func (this *PixelDelta) ToString() string {
        return fmt.Sprint(NewStringMaker().Append("PixelDelta: Point: ").Append(fmt.Sprint(this.Point)).Append(CommonSeps_GetInstance().SPACE).Append(fmt.Sprint(this.ColorDelta)))
}

func (this *PixelDelta) String() string {
    return this.ToString()
}

