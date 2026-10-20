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

type Rectangle struct {
    Point IGPoint
    Width int
    Height int
}


func (this *Rectangle) initFields() {
    
    
    
}


func NewRectangle(point IGPoint, width int, height int) *Rectangle {
    this := &Rectangle{}
    this.initFields()
        this.Point = point
        this.Width = width
        this.Height = height
    return this
}


func (this *Rectangle) GetPoint() IGPoint {
        return this.Point
}


func (this *Rectangle) SetPoint(point IGPoint) {
        this.Point = point
}


func (this *Rectangle) GetMaxX() int {
        return this.Point.GetX() + this.Width
}


func (this *Rectangle) GetMaxY() int {
        return this.Point.GetY() + this.Height
}


func (this *Rectangle) GetWidth() int {
        return this.Width
}


func (this *Rectangle) SetWidth(width int) {
        this.Width = width
}


func (this *Rectangle) GetHeight() int {
        return this.Height
}


func (this *Rectangle) SetHeight(height int) {
        this.Height = height
}


func (this *Rectangle) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(this.GetPoint()))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(CommonLabels_GetInstance().WIDTH_LABEL)
        stringBuffer.Appendint(this.GetWidth())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(CommonLabels_GetInstance().HEIGHT_LABEL)
        stringBuffer.Appendint(this.GetHeight())
        stringBuffer.Append(" MaxX: ")
        stringBuffer.Appendint(this.GetMaxX())
        stringBuffer.Append(" MaxY: ")
        stringBuffer.Appendint(this.GetMaxY())
        return fmt.Sprint(stringBuffer)
}

func (this *Rectangle) String() string {
    return this.ToString()
}

