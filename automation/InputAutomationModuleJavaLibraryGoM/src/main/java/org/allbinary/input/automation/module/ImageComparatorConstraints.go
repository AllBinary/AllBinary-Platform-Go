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

type ImageComparatorConstraints struct {
    LogUtil *LogUtil
    AvoidVector IBasicArrayList
    MaxNonMatchingPixelDeltas int
    DoImageComparisonEveryNthFrame int
}

// Implements ImageComparatorConstraintsInterface

// The Java class ImageComparatorConstraints has subclasses so its variables are this interface
type IImageComparatorConstraints interface {
    ImageComparatorConstraints_() *ImageComparatorConstraints
    GetAvoidVector() IBasicArrayList
    IsColorAllowed(frame int, point IGPoint, color *Color) bool
    IsCollisionWithAvoidRectangles(args ...any) bool
    SetAvoidVector(avoidVector IBasicArrayList)
    GetMaxNonMatchingPixelDeltas() int
    SetMaxNonMatchingPixelDeltas(maxNonMatchingPixelDeltas int)
    IsFrameAllowed(frame int) bool
    IsImageValid(bufferedImage *BufferedImage) bool
    Log()
}

func (this *ImageComparatorConstraints) ImageComparatorConstraints_() *ImageComparatorConstraints {
    return this
}


func (this *ImageComparatorConstraints) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewImageComparatorConstraints(doImageComparisonEveryNthFrame int) *ImageComparatorConstraints {
    this := &ImageComparatorConstraints{}
    this.initFields()
        this.DoImageComparisonEveryNthFrame = doImageComparisonEveryNthFrame
        this.SetAvoidVector(NewBasicArrayListD())
    return this
}


func (this *ImageComparatorConstraints) GetAvoidVector() IBasicArrayList {
        return this.AvoidVector
}


func (this *ImageComparatorConstraints) IsColorAllowed(frame int, point IGPoint, color *Color) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *ImageComparatorConstraints) IsCollisionWithAvoidRectangles_1(rectangle *Rectangle) bool {
        var isCollsionWithAvoidRectangles bool = false; _ = isCollsionWithAvoidRectangles
        var avoidVector IBasicArrayList = this.GetAvoidVector(); _ = avoidVector
        var size int = avoidVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var avoidRectangle *Rectangle = any(avoidVector.Get(index)).(*Rectangle); _ = avoidRectangle
            if RectangleCollisionUtil_IsCollision(avoidRectangle, rectangle) {
                isCollsionWithAvoidRectangles = true
            }
        }
        return isCollsionWithAvoidRectangles
}

// Java overloads of ImageComparatorConstraints.isCollisionWithAvoidRectangles
func (this *ImageComparatorConstraints) IsCollisionWithAvoidRectangles(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Rectangle](args[0], pass) {
            return this.IsCollisionWithAvoidRectangles_1(Java_Arg[*Rectangle](args[0]))
        }
        if len(args) == 1 && Java_IsArg[IGPoint](args[0], pass) {
            return this.IsCollisionWithAvoidRectangles_2(Java_Arg[IGPoint](args[0]))
        }
    }
    panic("No Java overload of ImageComparatorConstraints.isCollisionWithAvoidRectangles for the arguments")
}


func (this *ImageComparatorConstraints) IsCollisionWithAvoidRectangles_2(point IGPoint) bool {
        var isCollsionWithAvoidRectangles bool = false; _ = isCollsionWithAvoidRectangles
        var avoidVector IBasicArrayList = this.GetAvoidVector(); _ = avoidVector
        var size int = avoidVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var avoidRectangle *Rectangle = any(avoidVector.Get(index)).(*Rectangle); _ = avoidRectangle
            if RectangleCollisionUtil_IsCollision(avoidRectangle, point) {
                isCollsionWithAvoidRectangles = true
                break
            }
        }
        return isCollsionWithAvoidRectangles
}


func (this *ImageComparatorConstraints) SetAvoidVector(avoidVector IBasicArrayList) {
        this.AvoidVector = avoidVector
}


func (this *ImageComparatorConstraints) GetMaxNonMatchingPixelDeltas() int {
        return this.MaxNonMatchingPixelDeltas
}


func (this *ImageComparatorConstraints) SetMaxNonMatchingPixelDeltas(maxNonMatchingPixelDeltas int) {
        this.MaxNonMatchingPixelDeltas = maxNonMatchingPixelDeltas
}


func (this *ImageComparatorConstraints) IsFrameAllowed(frame int) bool {
        var remainder int = ((frame + 1) % this.DoImageComparisonEveryNthFrame); _ = remainder
        this.LogUtil.PutF(" Frame: " + fmt.Sprint(frame) + " remainder: " + fmt.Sprint(remainder) + " this.doImageComparisonEveryNthFrame: " + fmt.Sprint(this.DoImageComparisonEveryNthFrame), this, "isCollisionWithAvoidRectangles")
        if remainder == 0 {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *ImageComparatorConstraints) IsImageValid(bufferedImage *BufferedImage) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *ImageComparatorConstraints) Log() {
        var avoidVector IBasicArrayList = this.GetAvoidVector(); _ = avoidVector
        var size int = avoidVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var avoidRectangle *Rectangle = any(avoidVector.Get(index)).(*Rectangle); _ = avoidRectangle
            this.LogUtil.PutF("Avoid Rectangle: " + fmt.Sprint(avoidRectangle), this, "log")
        }
}

