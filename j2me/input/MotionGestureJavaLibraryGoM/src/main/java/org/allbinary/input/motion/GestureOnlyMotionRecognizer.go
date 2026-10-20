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

type GestureOnlyMotionRecognizer struct {
    MotionRecognizer
    MotionGestureRecognizer *MotionGestureRecognizer
    Id int
}

var GestureOnlyMotionRecognizer_Index int = 0
var GestureOnlyMotionRecognizer_PointCircularPool *GPointCircularPool = NewGPointCircularPool(40)

func (this *GestureOnlyMotionRecognizer) initFields() {
    
    
}


func NewGestureOnlyMotionRecognizer() *GestureOnlyMotionRecognizer {
    this := &GestureOnlyMotionRecognizer{}
    this.MotionRecognizer = *NewMotionRecognizer()
    this.initFields()
        this.Id = Java_PostIncrement(&GestureOnlyMotionRecognizer_Index)
        this.MotionGestureRecognizer = NewMotionGestureRecognizer(this.Id)
    return this
}


// throws Exception
func (this *GestureOnlyMotionRecognizer) ProcessStartMotionEvent(x int, y int, deviceId int, modifiers int) {
        var point *CustomGPoint = any(GestureOnlyMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
        point.SetX(x)
        point.SetY(y)
        this.MotionGestureRecognizer.ProcessPressedMotionEvent(point, deviceId, modifiers)
}


// throws Exception
func (this *GestureOnlyMotionRecognizer) ProcessEndMotionEvent(x int, y int, deviceId int, modifiers int) {
        var point *CustomGPoint = any(GestureOnlyMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
        point.SetX(x)
        point.SetY(y)
        this.MotionGestureRecognizer.ProcessReleasedMotionEvent(point, deviceId, modifiers)
}


// throws Exception
func (this *GestureOnlyMotionRecognizer) ProcessDraggedMotionEvent(x int, y int, deviceId int, modifiers int) {
        var point *CustomGPoint = any(GestureOnlyMotionRecognizer_PointCircularPool.GetNextInstance()).(*CustomGPoint); _ = point
        point.SetX(x)
        point.SetY(y)
        this.MotionGestureRecognizer.ProcessDraggedMotionEvent(point, deviceId, modifiers)
}

