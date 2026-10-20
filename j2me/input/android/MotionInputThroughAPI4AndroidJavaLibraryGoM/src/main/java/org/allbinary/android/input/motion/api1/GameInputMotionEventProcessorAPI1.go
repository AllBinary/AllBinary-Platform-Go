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

type GameInputMotionEventProcessorAPI1 struct {
    BaseGameInputMotionEventProcessor
    LogUtil *LogUtil
    MotionRecognizer IMotionRecognizer
}

var GameInputMotionEventProcessorAPI1_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameInputMotionEventProcessorAPI1) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MotionRecognizer = NewAllMotionRecognizer()
}


func GameInputMotionEventProcessorAPI1_GetInstance() *GameInputMotionEventProcessorAPI1 {
        if GameInputMotionEventProcessorAPI1_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameInputMotionEventProcessorAPI1_Instance = NewGameInputMotionEventProcessorAPI1()
        }
        return any(GameInputMotionEventProcessorAPI1_Instance).(*GameInputMotionEventProcessorAPI1)
}


func NewGameInputMotionEventProcessorAPI1() *GameInputMotionEventProcessorAPI1 {
    this := &GameInputMotionEventProcessorAPI1{}
    this.BaseGameInputMotionEventProcessor = *NewBaseGameInputMotionEventProcessor()
    this.initFields()
    return this
}


func (this *GameInputMotionEventProcessorAPI1) Process(motionEvent *MotionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            var x int = int(motionEvent.GetX()); _ = x
            var y int = int(motionEvent.GetY()); _ = y
            var action int = motionEvent.GetAction(); _ = action
            if action == MotionEvent_ACTION_DOWN {
                this.MotionRecognizer.ProcessStartMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
            } else if action == MotionEvent_ACTION_UP || action == MotionEvent_ACTION_CANCEL {
                this.MotionRecognizer.ProcessEndMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
            } else if action == MotionEvent_ACTION_MOVE {
                this.MotionRecognizer.ProcessDraggedMotionEvent(x, y, motionEvent.GetDeviceId(), 0)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

