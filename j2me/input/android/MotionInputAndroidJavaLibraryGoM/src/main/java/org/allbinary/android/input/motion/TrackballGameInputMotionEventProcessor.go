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

type TrackballGameInputMotionEventProcessor struct {
    MotionEventProcessor
    LogUtil *LogUtil
    Listener *SingleKeyPressGameKeyFromMotionGestureEventListener
    CommonStrings *CommonStrings
}

var TrackballGameInputMotionEventProcessor_PROCESSOR *TrackballGameInputMotionEventProcessor = NewTrackballGameInputMotionEventProcessor()

func (this *TrackballGameInputMotionEventProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Listener = NewSingleKeyPressGameKeyFromMotionGestureEventListener()
    this.CommonStrings = CommonStrings_GetInstance()
}


func TrackballGameInputMotionEventProcessor_GetInstance() *TrackballGameInputMotionEventProcessor {
        return TrackballGameInputMotionEventProcessor_PROCESSOR
}


func NewTrackballGameInputMotionEventProcessor() *TrackballGameInputMotionEventProcessor {
    this := &TrackballGameInputMotionEventProcessor{}
    this.MotionEventProcessor = *NewMotionEventProcessor()
    this.initFields()
    return this
}


func (this *TrackballGameInputMotionEventProcessor) Process(motionEvent *MotionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            var x float32 = motionEvent.GetX(); _ = x
            var y float32 = motionEvent.GetY(); _ = y
            if x < 0.0 {
                this.Listener.OnCompleteMotionGestureInputEvent(LeftTrackballInputToGameKeyEventAction_GetInstance())
            } else if x > 0.0 {
                this.Listener.OnCompleteMotionGestureInputEvent(RightTrackballInputToGameKeyEventAction_GetInstance())
            }
            if y < 0.0 {
                this.Listener.OnCompleteMotionGestureInputEvent(UpTrackballInputToGameKeyEventAction_GetInstance())
            } else if y > 0.0 {
                this.Listener.OnCompleteMotionGestureInputEvent(DownTrackballInputToGameKeyEventAction_GetInstance())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

