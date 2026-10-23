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

type FrameUtil struct {
}

var FrameUtil_Instance *FrameUtil = NewFrameUtil()


func FrameUtil_GetInstance() *FrameUtil {
        return FrameUtil_Instance
}


func NewFrameUtil() *FrameUtil {
    this := &FrameUtil{}
    return this
}


func (this *FrameUtil) GetFrameForAngle(angle int16, angleIncrement int) int {
        return this.AdjustAngleToFrameAngle(int(angle)) / angleIncrement
}


func (this *FrameUtil) GetFrameAngle(frame int, angleIncrement int) int {
        var frameAngle int = angleIncrement * frame - 90; _ = frameAngle
        return this.AdjustAngleToFrameAngle(frameAngle)
}


func (this *FrameUtil) AdjustAngleToFrameAngle(currentAngle int) int {
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        if currentAngle > 359 {
            currentAngle = (currentAngle - int(angleFactory.TOTAL_ANGLE))
            for currentAngle > 359 {
                currentAngle = (currentAngle - int(angleFactory.TOTAL_ANGLE))
            }
        }
        if currentAngle < 0 {
            currentAngle = (currentAngle + int(angleFactory.TOTAL_ANGLE))
            for currentAngle < 0 {
                currentAngle = (currentAngle + int(angleFactory.TOTAL_ANGLE))
            }
        }
        return currentAngle
}

