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

type MotionRectangleConstraints struct {
    MinDimension *Dimension
    MaxDimension *Dimension
    MaxMotionRectangles int
    MinArea int
}

// Implements MotionRectangleConstraintsInterface

// The Java class MotionRectangleConstraints has subclasses so its variables are this interface
type IMotionRectangleConstraints interface {
    MotionRectangleConstraints_() *MotionRectangleConstraints
    GetMinDimension() *Dimension
    GetMaxDimension() *Dimension
    IsTooSmall(rectangle *Rectangle) bool
    IsAreaTooSmall(rectangle *Rectangle) bool
    IsTooBig(rectangle *Rectangle) bool
    IsValid(frame int64, bufferedImage *BufferedImage, rectangle *Rectangle) bool
    SetMinDimension(minDimension *Dimension)
    SetMaxDimension(maxDimension *Dimension)
    GetMaxMotionRectangles() int
    SetMaxMotionRectangles(maxMotionRectangles int)
    GetMinArea() int
    SetMinArea(minArea int)
}

func (this *MotionRectangleConstraints) MotionRectangleConstraints_() *MotionRectangleConstraints {
    return this
}


func (this *MotionRectangleConstraints) initFields() {
    
    
    
    
}


func NewMotionRectangleConstraints() *MotionRectangleConstraints {
    this := &MotionRectangleConstraints{}
    this.initFields()
    return this
}


func (this *MotionRectangleConstraints) GetMinDimension() *Dimension {
        return this.MinDimension
}


func (this *MotionRectangleConstraints) GetMaxDimension() *Dimension {
        return this.MaxDimension
}


func (this *MotionRectangleConstraints) IsTooSmall(rectangle *Rectangle) bool {
        if rectangle.Width < this.GetMinDimension().Width || rectangle.Height < this.GetMinDimension().Height {
            return true
        } else {
            return false
        }
}


func (this *MotionRectangleConstraints) IsAreaTooSmall(rectangle *Rectangle) bool {
        if rectangle.Width * rectangle.Height < this.GetMinArea() {
            return true
        } else {
            return false
        }
}


func (this *MotionRectangleConstraints) IsTooBig(rectangle *Rectangle) bool {
        if rectangle.Width > this.GetMaxDimension().Width || rectangle.Height > this.GetMaxDimension().Height {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *MotionRectangleConstraints) IsValid(frame int64, bufferedImage *BufferedImage, rectangle *Rectangle) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *MotionRectangleConstraints) SetMinDimension(minDimension *Dimension) {
        this.MinDimension = minDimension
}


func (this *MotionRectangleConstraints) SetMaxDimension(maxDimension *Dimension) {
        this.MaxDimension = maxDimension
}


func (this *MotionRectangleConstraints) GetMaxMotionRectangles() int {
        return this.MaxMotionRectangles
}


func (this *MotionRectangleConstraints) SetMaxMotionRectangles(maxMotionRectangles int) {
        this.MaxMotionRectangles = maxMotionRectangles
}


func (this *MotionRectangleConstraints) GetMinArea() int {
        return this.MinArea
}


func (this *MotionRectangleConstraints) SetMinArea(minArea int) {
        this.MinArea = minArea
}

