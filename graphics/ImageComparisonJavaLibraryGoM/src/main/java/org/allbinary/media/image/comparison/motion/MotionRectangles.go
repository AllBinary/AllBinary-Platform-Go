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

type MotionRectangles struct {
    Name string
    RectangleVector IBasicArrayList
    ImageComparisonResult *ImageComparisonResult
    MAX_PIXEL_DISTANCE_THRESHOLD int
}

// The Java class MotionRectangles has subclasses so its variables are this interface
type IMotionRectangles interface {
    MotionRectangles_() *MotionRectangles
    GetVector() IBasicArrayList
    SetVector(rectangleVector IBasicArrayList)
    GetName() string
    GetImageComparisonResult() *ImageComparisonResult
    SetImageComparisonResult(imageComparisonResult *ImageComparisonResult)
}

func (this *MotionRectangles) MotionRectangles_() *MotionRectangles {
    return this
}


func (this *MotionRectangles) initFields() {
    
    
    
    this.MAX_PIXEL_DISTANCE_THRESHOLD = 2
}


func NewMotionRectangles(name string, imageComparisonResult *ImageComparisonResult) *MotionRectangles {
    this := &MotionRectangles{}
    this.initFields()
        this.Name = name
        this.SetImageComparisonResult(imageComparisonResult)
        this.SetVector(NewBasicArrayListD())
    return this
}


func (this *MotionRectangles) GetVector() IBasicArrayList {
        return this.RectangleVector
}


func (this *MotionRectangles) SetVector(rectangleVector IBasicArrayList) {
        this.RectangleVector = rectangleVector
}


func (this *MotionRectangles) GetName() string {
        return this.Name
}


func (this *MotionRectangles) GetImageComparisonResult() *ImageComparisonResult {
        return this.ImageComparisonResult
}


func (this *MotionRectangles) SetImageComparisonResult(imageComparisonResult *ImageComparisonResult) {
        this.ImageComparisonResult = imageComparisonResult
}

