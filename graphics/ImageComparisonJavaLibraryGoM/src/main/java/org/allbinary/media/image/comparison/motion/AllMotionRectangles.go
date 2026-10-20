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

type AllMotionRectangles struct {
    MotionRectangles
}

var AllMotionRectangles_NAME string = "allMotionRectangles_"


func NewAllMotionRectangles(imageComparisonResult *ImageComparisonResult) *AllMotionRectangles {
    this := &AllMotionRectangles{}
        this.MotionRectangles = *NewMotionRectangles(AllMotionRectangles_NAME, imageComparisonResult)
        this.CreateMotionRectangles(imageComparisonResult)
    return this
}


// throws Exception
func (this *AllMotionRectangles) CreateMotionRectangles(imageComparisonInfo *ImageComparisonResult) {
        var vector IBasicArrayList = imageComparisonInfo.GetNonMatchingPixelVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var pixelDelta *PixelDelta = any(vector.Get(index)).(*PixelDelta); _ = pixelDelta
            this.AddPixelDeltaToAMotionRectangle(pixelDelta)
        }
}


// throws Exception
func (this *AllMotionRectangles) AddPixelDeltaToExistingMotionRectangle(pixelDelta *PixelDelta) bool {
        var vector IBasicArrayList = this.GetVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var motionRectangle *MotionRectangle = any(vector.Get(index)).(*MotionRectangle); _ = motionRectangle
            if this.ShouldPixelDeltaPartOfMotionRectangle(motionRectangle, pixelDelta) {
                var newWidth int = pixelDelta.GetPoint().GetX() - motionRectangle.GetRectangle().X; _ = newWidth
                var newHeight int = pixelDelta.GetPoint().GetY() - motionRectangle.GetRectangle().Y; _ = newHeight
                if newWidth < motionRectangle.GetRectangle().Width {
                    newWidth = motionRectangle.GetRectangle().Width
                }
                if newHeight < motionRectangle.GetRectangle().Height {
                    newHeight = motionRectangle.GetRectangle().Height
                }
                motionRectangle.GetRectangle().SetSize(newWidth, newHeight)
                return true
            }
        }
        return false
}


// throws Exception
func (this *AllMotionRectangles) AddPixelDeltaToAMotionRectangle(pixelDelta *PixelDelta) {
        if !this.AddPixelDeltaToExistingMotionRectangle(pixelDelta) {
            var newRectangle *Rectangle = NewRectangle(pixelDelta.GetPoint().GetX(), pixelDelta.GetPoint().GetY(), 1, 1); _ = newRectangle
            var motionRectangle *MotionRectangle = NewMotionRectangle(newRectangle); _ = motionRectangle
            motionRectangle.GetPixelDeltaVector().Add(pixelDelta)
            this.GetVector().Add(motionRectangle)
        }
}


// throws Exception
func (this *AllMotionRectangles) ShouldPixelDeltaPartOfMotionRectangle(motionRectangle *MotionRectangle, pixelDelta *PixelDelta) bool {
        var isPixelLeftOfTheRectangle bool = false; _ = isPixelLeftOfTheRectangle
        if motionRectangle.GetRectangle().Y > pixelDelta.GetPoint().GetY() {
            panic(errors.New(fmt.Sprint("PixelAboveTheRectangle")))
        }
        if motionRectangle.GetRectangle().X > pixelDelta.GetPoint().GetX() {
            isPixelLeftOfTheRectangle = true
        }
        if !isPixelLeftOfTheRectangle {
            return this.IsPixelDeltaLikelyPartOfMotionRectangle(motionRectangle, pixelDelta)
        }
        return false
}


// throws Exception
func (this *AllMotionRectangles) IsPixelDeltaLikelyPartOfMotionRectangle(motionRectangle *MotionRectangle, pixelDelta *PixelDelta) bool {
        var maxX int = motionRectangle.GetRectangle().X + motionRectangle.GetRectangle().Width; _ = maxX
        var maxY int = motionRectangle.GetRectangle().Y + motionRectangle.GetRectangle().Height; _ = maxY
        var horizontalDistance int = pixelDelta.GetPoint().GetX() - maxX; _ = horizontalDistance
        var verticalDistance int = pixelDelta.GetPoint().GetY() - maxY; _ = verticalDistance
        if horizontalDistance < this.MAX_PIXEL_DISTANCE_THRESHOLD && verticalDistance < this.MAX_PIXEL_DISTANCE_THRESHOLD {
            return true
        }
        return false
}

