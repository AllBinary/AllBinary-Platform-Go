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

type GPoint struct {
    X int
    Y int
    Z int
}

// The Java class GPoint has subclasses so its variables are this interface
type IGPoint interface {
    GPoint_() *GPoint
    GetX() int
    GetY() int
    GetZ() int
    GetRawX() int
    GetRawY() int
    GetRawZ() int
    ToString() string
}

func (this *GPoint) GPoint_() *GPoint {
    return this
}

var GPoint_POINT_LABEL string = "Point: "

func (this *GPoint) initFields() {
    
    
    
}


func GPoint_GetInstance(point IGPoint) IGPoint {
        return NewGPoint(point.GetX(), point.GetY(), point.GetZ())
}


func NewGPoint(x int, y int, z int) *GPoint {
    this := &GPoint{}
    this.initFields()
        this.X = x
        this.Y = y
        this.Z = z
    return this
}


func (this *GPoint) GetX() int {
        return this.X
}


func (this *GPoint) GetY() int {
        return this.Y
}


func (this *GPoint) GetZ() int {
        return this.Z
}


func (this *GPoint) GetRawX() int {
        return this.X
}


func (this *GPoint) GetRawY() int {
        return this.Y
}


func (this *GPoint) GetRawZ() int {
        return this.Z
}


func (this *GPoint) ToString() string {
        return GPoint_ToStringStatic(this.GetX(), this.GetY(), this.GetZ())
}

func (this *GPoint) String() string {
    return this.ToString()
}


func GPoint_ToStringStatic(x int, y int, z int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        stringBuffer.Append(GPoint_POINT_LABEL)
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Appendint(x)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Appendint(y)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Appendint(z)
        return fmt.Sprint(stringBuffer)
}
