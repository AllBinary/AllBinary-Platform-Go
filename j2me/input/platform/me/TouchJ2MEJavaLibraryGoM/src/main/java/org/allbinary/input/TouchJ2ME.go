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

//TouchJ2ME
type TouchJ2ME struct {
    LogUtil *LogUtil
    DEVICE_ID int
    MotionRecognizer *AllMotionRecognizer
}


func (this *TouchJ2ME) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.DEVICE_ID = 0
    this.MotionRecognizer = NewAllMotionRecognizer()
}


func NewTouchJ2ME() *TouchJ2ME {
    this := &TouchJ2ME{}
    this.initFields()
        var motionGesturesHandler *BasicMotionGesturesHandler = this.MotionRecognizer.GetMotionGestureRecognizer().GetMotionGesturesHandler(); _ = motionGesturesHandler
        motionGesturesHandler.AddListenerInterface(NewGameMotionGestureListener(MotionGestureReceiveInterfaceFactory_GetInstance()))
    return this
}


func (this *TouchJ2ME) PointerDragged(x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "pointerDragged", e)
                }
            }()
            this.MotionRecognizer.ProcessDraggedMotionEvent(x, y, this.DEVICE_ID, 0)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchJ2ME) PointerPressed(x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "pointerPressed", e)
                }
            }()
            this.MotionRecognizer.ProcessStartMotionEvent(x, y, this.DEVICE_ID, 0)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchJ2ME) PointerReleased(x int, y int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "pointerReleased", e)
                }
            }()
            this.MotionRecognizer.ProcessEndMotionEvent(x, y, this.DEVICE_ID, 0)
            return
        }(); tryControl1 == 1 {
            return
        }
}

