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

type ConstrainedMotionRectangles struct {
    MotionRectangles
    LogUtil *LogUtil
    MotionRectangleConstraintsInterface MotionRectangleConstraintsInterface
}

var ConstrainedMotionRectangles_NAME string = "constrainedMotionRectangles_"

func (this *ConstrainedMotionRectangles) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewConstrainedMotionRectangles(motionRectangleConstraintsInterface MotionRectangleConstraintsInterface, motionRectangles IMotionRectangles) *ConstrainedMotionRectangles {
    this := &ConstrainedMotionRectangles{}
        this.MotionRectangles = *NewMotionRectangles(ConstrainedMotionRectangles_NAME, motionRectangles.GetImageComparisonResult())
        this.initFields()
        this.SetMotionRectangleConstraintsInterface(motionRectangleConstraintsInterface)
    return this
}


// throws Exception
func (this *ConstrainedMotionRectangles) ApplyMotionRectangleConstraints(motionRectangles IMotionRectangles) {
        this.LogUtil.PutF("Start - Size Before: " + fmt.Sprint(motionRectangles.GetVector().Size()), this, "applyMotionRectangleConstraints")
        var bufferedImageArray []*BufferedImage = motionRectangles.GetImageComparisonResult().GetBufferedImages(); _ = bufferedImageArray
        var bufferedImage *BufferedImage = bufferedImageArray[1]; _ = bufferedImage
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var vector2 IBasicArrayList = motionRectangles.GetVector(); _ = vector2
        var size int = vector2.Size(); _ = size
        for index := int(0); index < size; index++ {
            var motionRectangle *MotionRectangle = any(vector2.Get(index)).(*MotionRectangle); _ = motionRectangle
            var rectangle *Rectangle = motionRectangle.GetRectangle(); _ = rectangle
            var isTooSmall bool = this.GetMotionRectangleConstraintsInterface().IsTooSmall(rectangle); _ = isTooSmall
            if isTooSmall {
                continue
            }
            var isAreaTooSmall bool = this.GetMotionRectangleConstraintsInterface().IsAreaTooSmall(rectangle); _ = isAreaTooSmall
            if isAreaTooSmall {
                continue
            }
            var isTooBig bool = this.GetMotionRectangleConstraintsInterface().IsTooBig(rectangle); _ = isTooBig
            if isTooBig {
                continue
            }
            var isValid bool = this.GetMotionRectangleConstraintsInterface().IsValid(this.GetImageComparisonResult().GetFrameTwo(), bufferedImage, motionRectangle.GetRectangle()); _ = isValid
            if !isValid {
                continue
            }
            if !isTooSmall && !isAreaTooSmall && !isTooBig && isValid {
                vector.Add(motionRectangle)
            }
        }
        this.SetVector(vector)
}


func (this *ConstrainedMotionRectangles) GetMotionRectangleConstraintsInterface() MotionRectangleConstraintsInterface {
        return this.MotionRectangleConstraintsInterface
}


func (this *ConstrainedMotionRectangles) SetMotionRectangleConstraintsInterface(motionRectangleConstraintsInterface MotionRectangleConstraintsInterface) {
        this.MotionRectangleConstraintsInterface = motionRectangleConstraintsInterface
}

