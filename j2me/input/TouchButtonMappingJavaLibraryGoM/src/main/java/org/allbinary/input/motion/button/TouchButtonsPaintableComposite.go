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

type TouchButtonsPaintableComposite struct {
    ProcessPaintable
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputMappingInterface InputMappingInterface
    TouchButtonsPaintable *TouchButtonsMappingPaintable
    TimeHelper ITimeDelayHelper
    Released bool
    METHOD_NAME string
    RELEASE string
    FAST_REPEAT string
    IGNORE string
}

// Implements CompleteMotionGestureInputEventListenerInterface


func (this *TouchButtonsPaintableComposite) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.TimeHelper = NewTimeDelayHelper(250)
    this.Released = true
    this.METHOD_NAME = "onCompleteMotionGestureInputEvent"
    this.RELEASE = "Ignoring: MotionGestureInput Release"
    this.FAST_REPEAT = "Ignoring: MotionGestureInput Repeated To Quickly"
    this.IGNORE = "Ignoring: Until Released"
}


func NewTouchButtonsPaintableComposite(inputMappingInterface InputMappingInterface, basicColor *BasicColor) *TouchButtonsPaintableComposite {
    this := &TouchButtonsPaintableComposite{}
    this.ProcessPaintable = *NewProcessPaintable()
    this.initFields()
        this.InputMappingInterface = inputMappingInterface
        CompleteMotionGestureInputEventHandler_GetInstance().AddListenerInterface(this)
        this.TouchButtonsPaintable = NewTouchButtonsMappingPaintable(basicColor)
    return this
}


func (this *TouchButtonsPaintableComposite) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *TouchButtonsPaintableComposite) OnCompleteMotionGestureInputEvent(completeMotionGestureInputEvent ICompleteMotionGestureInputEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.METHOD_NAME, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "onCompleteMotionGestureInputEvent")
            var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
            var gameKeyCompleteMotionGestureInputEvent IGameKeyCompleteMotionGestureInputEvent = any(completeMotionGestureInputEvent).(IGameKeyCompleteMotionGestureInputEvent); _ = gameKeyCompleteMotionGestureInputEvent
            var gameKey *GameKey = gameKeyCompleteMotionGestureInputEvent.GetGameKey(); _ = gameKey
            var motionGestureInput *MotionGestureInput = completeMotionGestureInputEvent.GetMotionGestureInput(); _ = motionGestureInput
            if motionGestureInput == touchMotionGestureFactory.RELEASED {
                this.LogUtil.PutF(this.RELEASE, this, this.METHOD_NAME)
                this.Released = true
                tryControl1 = 1; return
            }
            if motionGestureInput == TouchMotionGestureFactory_GetInstance().NO_MOTION && !this.TimeHelper.IsTimeTNT() {
                this.LogUtil.PutF(this.FAST_REPEAT, this, this.METHOD_NAME)
                tryControl1 = 1; return
            }
            if !this.Released {
                this.LogUtil.PutF(this.IGNORE, this, this.METHOD_NAME)
                tryControl1 = 1; return
            }
            this.Released = false
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("GameKey: ").Append(StringUtil_GetInstance().ToString(gameKey)).Append(" MotionGestureInput: ").Append(StringUtil_GetInstance().ToString(motionGestureInput))), this, this.METHOD_NAME)
            this.InputMappingInterface.ProcessInputMapping(gameKey, motionGestureInput)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchButtonsPaintableComposite) Process() {
        CompleteMotionGestureInputEventHandler_GetInstance().RemoveListener(this)
}

