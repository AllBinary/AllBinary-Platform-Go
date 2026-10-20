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

type ConsolidateMotionRectangles struct {
    MotionRectangles
    LogUtil *LogUtil
}

var ConsolidateMotionRectangles_NAME string = "consolidatedMotionRectangles_"

func (this *ConsolidateMotionRectangles) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewConsolidateMotionRectangles(motionRectangles IMotionRectangles) *ConsolidateMotionRectangles {
    this := &ConsolidateMotionRectangles{}
        this.MotionRectangles = *NewMotionRectangles(ConsolidateMotionRectangles_NAME, motionRectangles.GetImageComparisonResult())
        this.initFields()
        var vector IBasicArrayList = motionRectangles.GetVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var motionRectangle *MotionRectangle = any(vector.Get(index)).(*MotionRectangle); _ = motionRectangle
            vector.Add(motionRectangle.Clone())
        }
        this.ConsolidateMotionRectangles()
    return this
}


func (this *ConsolidateMotionRectangles) ConsolidateMotionRectangles() {
        var vector IBasicArrayList = this.GetVector(); _ = vector
        this.LogUtil.PutF("Start - Size Before: " + fmt.Sprint(vector.Size()), this, "consolidateMotionRectangleConstraints")
        var removeVector IBasicArrayList = NewBasicArrayListD(); _ = removeVector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var motionRectangle *MotionRectangle = any(vector.Get(index)).(*MotionRectangle); _ = motionRectangle
            for i := int(index + 1); i < vector.Size(); i++ {
                var motionRectangle2 *MotionRectangle = any(vector.Get(i)).(*MotionRectangle); _ = motionRectangle2
                var rectangle *Rectangle = motionRectangle.GetRectangle(); _ = rectangle
                var rectangle2 *Rectangle = motionRectangle2.GetRectangle(); _ = rectangle2
                var x2 int = rectangle.X + rectangle.Width; _ = x2
                var y2 int = rectangle.Y + rectangle.Height; _ = y2
                var rightX int = rectangle2.X + rectangle2.Width; _ = rightX
                var bottomY int = rectangle2.Y + rectangle2.Height; _ = bottomY
                if rectangle.X <= rectangle2.X && rectangle.Y <= rectangle2.Y && x2 >= rightX && y2 >= bottomY {
                    removeVector.Add(motionRectangle2)
                }
            }
        }
        this.Remove(removeVector)
}


func (this *ConsolidateMotionRectangles) Remove(removeVector IBasicArrayList) {
        var vector IBasicArrayList = this.GetVector(); _ = vector
        var size int = removeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            vector.Remove(removeVector.Get(index))
        }
}

