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

type CustomGPoint struct {
    GPoint
    Cx int
    Cy int
    Cz int
}

var CustomGPoint_NULL_CUSTOM_POINT *CustomGPoint = NewCustomGPoint(0, 0, 0)

func (this *CustomGPoint) initFields() {
    
    
    
}


func CustomGPoint_GetInstance(point IGPoint) *CustomGPoint {
        return NewCustomGPoint(point.GetX(), point.GetY(), point.GetZ())
}


func CustomGPoint_GetInstance3(x int, y int) *CustomGPoint {
        return NewCustomGPoint(x, y, 3)
}


func NewCustomGPoint(x int, y int, z int) *CustomGPoint {
    this := &CustomGPoint{}
        this.GPoint = *NewGPoint(x, y, z)
        this.initFields()
    return this
}


func (this *CustomGPoint) SetX(x int) {
        this.Cx = x
}


func (this *CustomGPoint) GetX() int {
        return this.Cx
}


func (this *CustomGPoint) SetY(y int) {
        this.Cy = y
}


func (this *CustomGPoint) GetY() int {
        return this.Cy
}


func (this *CustomGPoint) SetZ(z int) {
        this.Cz = z
}


func (this *CustomGPoint) GetZ() int {
        return this.Cz
}

