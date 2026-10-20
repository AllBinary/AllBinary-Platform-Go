/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2009 AllBinary
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

type GameInputMotionEventProcessorAPI5 struct {
    BaseGameInputMotionEventProcessor
    LogUtil *LogUtil
    DisplayInfo *DisplayInfoSingleton
    MotionRecognizerArray []IMotionRecognizer
    TOTAL_MINUS_ONE int
}

var GameInputMotionEventProcessorAPI5_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameInputMotionEventProcessorAPI5) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    
    
}


func GameInputMotionEventProcessorAPI5_GetInstance() *GameInputMotionEventProcessorAPI5 {
        if GameInputMotionEventProcessorAPI5_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameInputMotionEventProcessorAPI5_Instance = NewGameInputMotionEventProcessorAPI5()
        }
        return any(GameInputMotionEventProcessorAPI5_Instance).(*GameInputMotionEventProcessorAPI5)
}


func NewGameInputMotionEventProcessorAPI5() *GameInputMotionEventProcessorAPI5 {
    this := &GameInputMotionEventProcessorAPI5{}
    this.BaseGameInputMotionEventProcessor = *NewBaseGameInputMotionEventProcessor()
    this.initFields()
        var size int = GameInputMotionInfoAPI5_GetInstance().MAX_POINTERS; _ = size
        this.MotionRecognizerArray = make([]*AllMotionRecognizer, size)
        this.TOTAL_MINUS_ONE = size
        for index := int(0); index < size; index++ {
            this.MotionRecognizerArray[index] = NewAllMotionRecognizer()
        }
    return this
}


func (this *GameInputMotionEventProcessorAPI5) IsValid(pointerId int) bool {
        if pointerId > this.TOTAL_MINUS_ONE {
            var message string = "Log that the pointer Ids are not what would be expected or Android has more than what I set my stuff to support pointerId: "; _ = message
            ForcedLogUtil_Log(message + fmt.Sprint(pointerId), this)
            return false
        } else {
            return true
        }
}


func (this *GameInputMotionEventProcessorAPI5) Process(motionEvent *MotionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            var action int = motionEvent.GetAction(); _ = action
            var total int = len(this.MotionRecognizerArray); _ = total
            if total > motionEvent.GetPointerCount() {
                total = motionEvent.GetPointerCount()
            } else {
                var message string = "Log that Android has more pointers than I have expected so increase it motionEvent.getPointerCount(): "; _ = message
                ForcedLogUtil_Log(message + fmt.Sprint(motionEvent.GetPointerCount()), this)
            }
            var totalMinusOne int = total - 1; _ = totalMinusOne
            if action == MotionEvent_ACTION_DOWN {
                for pointerIndex := int(totalMinusOne); pointerIndex >= 0; pointerIndex-- {
                    var pointerId int = motionEvent.GetPointerId(pointerIndex); _ = pointerId
                    if this.IsValid(pointerId) {
                        var scalableListener ScalableListener = this.DisplayInfo.GetScalableListener(); _ = scalableListener
                        var x int = int(((motionEvent.GetX(pointerIndex) - this.DisplayInfo.GetLeft()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = x
                        var y int = int(((motionEvent.GetY(pointerIndex) - this.DisplayInfo.GetTop()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = y
                        this.MotionRecognizerArray[pointerId].ProcessStartMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
                    }
                }
            } else if action == MotionEvent_ACTION_UP || action == MotionEvent_ACTION_CANCEL {
                for pointerIndex := int(totalMinusOne); pointerIndex >= 0; pointerIndex-- {
                    var pointerId int = motionEvent.GetPointerId(pointerIndex); _ = pointerId
                    if this.IsValid(pointerId) {
                        var scalableListener ScalableListener = this.DisplayInfo.GetScalableListener(); _ = scalableListener
                        var x int = int(((motionEvent.GetX(pointerIndex) - this.DisplayInfo.GetLeft()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = x
                        var y int = int(((motionEvent.GetY(pointerIndex) - this.DisplayInfo.GetTop()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = y
                        this.MotionRecognizerArray[pointerId].ProcessEndMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
                    }
                }
            } else if action == MotionEvent_ACTION_MOVE {
                for pointerIndex := int(totalMinusOne); pointerIndex >= 0; pointerIndex-- {
                    var pointerId int = motionEvent.GetPointerId(pointerIndex); _ = pointerId
                    if this.IsValid(pointerId) {
                        var scalableListener ScalableListener = this.DisplayInfo.GetScalableListener(); _ = scalableListener
                        var x int = int(((motionEvent.GetX(pointerIndex) - this.DisplayInfo.GetLeft()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = x
                        var y int = int(((motionEvent.GetY(pointerIndex) - this.DisplayInfo.GetTop()) / scalableListener.GetRatio(this.DisplayInfo.GetRatio()))); _ = y
                        this.MotionRecognizerArray[pointerId].ProcessDraggedMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

