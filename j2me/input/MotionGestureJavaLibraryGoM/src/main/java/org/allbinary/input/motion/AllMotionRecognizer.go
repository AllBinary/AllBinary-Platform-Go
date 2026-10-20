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

type AllMotionRecognizer struct {
    MotionRecognizer
    LogUtil *LogUtil
    MotionGestureRecognizer *MotionGestureRecognizer
    TouchButtonProcessing bool
    TouchButtonRecognizer *TouchButtonRecognizer
    Id int
    LastX int
    LastY int
}

var AllMotionRecognizer_Index int = 0
var AllMotionRecognizer_PointCircularPool *GPointCircularPool = NewGPointCircularPool(40)

func (this *AllMotionRecognizer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.TouchButtonProcessing = false
    
    
    
    
}


func NewAllMotionRecognizer() *AllMotionRecognizer {
    this := &AllMotionRecognizer{}
    this.MotionRecognizer = *NewMotionRecognizer()
    this.initFields()
        this.Id = Java_PostIncrement(&AllMotionRecognizer_Index)
        this.MotionGestureRecognizer = NewMotionGestureRecognizer(this.Id)
        this.TouchButtonRecognizer = NewTouchButtonRecognizer()
    return this
}


// throws Exception
func (this *AllMotionRecognizer) ProcessStartMotionEvent(x int, y int, deviceId int, modifiers int) {
        if this.TouchButtonRecognizer.PressTouchButtonInput(x, y, deviceId) {
            this.TouchButtonProcessing = true
        }
        if !this.TouchButtonProcessing {
            var point *CustomGPoint = any(AllMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
            point.SetX(x)
            point.SetY(y)
            this.MotionGestureRecognizer.ProcessPressedMotionEvent(point, deviceId, modifiers)
        }
}


// throws Exception
func (this *AllMotionRecognizer) ProcessEndMotionEvent(x int, y int, deviceId int, modifiers int) {
        if this.TouchButtonRecognizer.ReleaseTouchButtonInput(x, y, deviceId) {
        }
        if !this.TouchButtonProcessing {
            var point *CustomGPoint = any(AllMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
            point.SetX(x)
            point.SetY(y)
            this.MotionGestureRecognizer.ProcessReleasedMotionEvent(point, deviceId, modifiers)
        }
        if this.TouchButtonProcessing {
            this.TouchButtonProcessing = false
        }
}


// throws Exception
func (this *AllMotionRecognizer) ProcessDraggedMotionEvent(x int, y int, deviceId int, modifiers int) {
        if this.TouchButtonProcessing {
            this.TouchButtonRecognizer.PressTouchButtonInput(x, y, deviceId)
        }
        if !this.TouchButtonProcessing {
            var point *CustomGPoint = any(AllMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
            point.SetX(x)
            point.SetY(y)
            this.MotionGestureRecognizer.ProcessDraggedMotionEvent(point, deviceId, modifiers)
        }
}


// throws Exception
func (this *AllMotionRecognizer) ProcessMovedMotionEvent(x int, y int, deviceId int, modifiers int) {
        if x != this.LastX || y != this.LastY {
            this.LastX = x
            this.LastY = y
            var point *CustomGPoint = any(AllMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
            point.SetX(x)
            point.SetY(y)
            this.MotionGestureRecognizer.ProcessMovedMotionEvent(point, deviceId, modifiers)
        }
}


// throws Exception
func (this *AllMotionRecognizer) ProcessScrolledMotionEvent(x int, y int, deviceId int, modifiers int) {
        var point *CustomGPoint = any(AllMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
        point.SetX(x)
        point.SetY(y)
        this.MotionGestureRecognizer.ProcessScrolledMotionEvent(point, deviceId, modifiers)
}


func (this *AllMotionRecognizer) GetMotionGestureRecognizer() *MotionGestureRecognizer {
        return this.MotionGestureRecognizer
}

