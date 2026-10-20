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

type NoMotionRectangleConstraints struct {
    MotionRectangleConstraints
}



func NewNoMotionRectangleConstraints() *NoMotionRectangleConstraints {
    this := &NoMotionRectangleConstraints{}
    this.MotionRectangleConstraints = *NewMotionRectangleConstraints()
        this.SetMinDimension(NewDimension(0, 0))
        this.SetMinArea(0)
        this.SetMaxDimension(NewDimension(Integer_MAX_VALUE, Integer_MAX_VALUE))
        this.SetMaxMotionRectangles(Integer_MAX_VALUE)
    return this
}


func (this *NoMotionRectangleConstraints) IsValid(frame int64, bufferedImage *BufferedImage, rectangle *Rectangle) bool {
        return true
}

