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

type MotionRectangle struct {
    Rectangle *Rectangle
    PixelDeltaVector IBasicArrayList
}

// Implements *Cloneable


func (this *MotionRectangle) initFields() {
    
    
}


func NewMotionRectangle(rectangle *Rectangle) *MotionRectangle {
    this := &MotionRectangle{}
    this.initFields()
        this.SetRectangle(rectangle)
        this.SetPixelDeltaVector(NewBasicArrayListD())
    return this
}


func (this *MotionRectangle) GetRectangle() *Rectangle {
        return this.Rectangle
}


func (this *MotionRectangle) SetRectangle(rectangle *Rectangle) {
        this.Rectangle = rectangle
}


func (this *MotionRectangle) GetPixelDeltaVector() IBasicArrayList {
        return this.PixelDeltaVector
}


func (this *MotionRectangle) SetPixelDeltaVector(pixelDeltaVector IBasicArrayList) {
        this.PixelDeltaVector = pixelDeltaVector
}


func (this *MotionRectangle) Clone() any {
        var motionRectangle *MotionRectangle = NewMotionRectangle(any(this.GetRectangle().Clone()).(*Rectangle)); _ = motionRectangle
        motionRectangle.SetPixelDeltaVector(any(this.GetPixelDeltaVector().Clone()).(IBasicArrayList))
        return motionRectangle
}

