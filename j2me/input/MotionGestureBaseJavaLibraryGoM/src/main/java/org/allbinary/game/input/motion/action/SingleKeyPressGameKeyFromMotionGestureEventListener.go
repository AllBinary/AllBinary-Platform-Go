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

type SingleKeyPressGameKeyFromMotionGestureEventListener struct {
    CompleteMotionGestureInputEventListener
    LogUtil *LogUtil
}


func (this *SingleKeyPressGameKeyFromMotionGestureEventListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewSingleKeyPressGameKeyFromMotionGestureEventListener() *SingleKeyPressGameKeyFromMotionGestureEventListener {
    this := &SingleKeyPressGameKeyFromMotionGestureEventListener{}
    this.CompleteMotionGestureInputEventListener = *NewCompleteMotionGestureInputEventListener()
    this.initFields()
    return this
}


func (this *SingleKeyPressGameKeyFromMotionGestureEventListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *SingleKeyPressGameKeyFromMotionGestureEventListener) OnCompleteMotionGestureInputEvent(completeMotionGestureInputEvent ICompleteMotionGestureInputEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "onCompleteMotionGestureInputEvent", e)
                }
            }()
            var gameKeyCompleteMotionGestureInputEvent IGameKeyCompleteMotionGestureInputEvent = any(completeMotionGestureInputEvent).(IGameKeyCompleteMotionGestureInputEvent); _ = gameKeyCompleteMotionGestureInputEvent
            var gameKeyEvent *GameKeyEvent = gameKeyCompleteMotionGestureInputEvent.GetGameKeyEvent(); _ = gameKeyEvent
            PressGameKeyEventHandler_GetInstance().FireEvent(gameKeyEvent)
            return
        }(); tryControl1 == 1 {
            return
        }
}

