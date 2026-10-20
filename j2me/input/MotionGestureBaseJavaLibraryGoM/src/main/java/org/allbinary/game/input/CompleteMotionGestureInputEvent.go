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

type CompleteMotionGestureInputEvent struct {
    AllBinaryEventObject
    LogUtil *LogUtil
    Name string
    MotionGestureInput *MotionGestureInput
}

// Implements CompleteMotionGestureInputInterface

// The Java class CompleteMotionGestureInputEvent has subclasses so its variables are this interface
type ICompleteMotionGestureInputEvent interface {
    IAllBinaryEventObject
    CompleteMotionGestureInputEvent_() *CompleteMotionGestureInputEvent
    SetName(name string)
    GetName() string
    SetMotionGestureInput(motionGestureInput *MotionGestureInput)
    GetMotionGestureInput() *MotionGestureInput
}

func (this *CompleteMotionGestureInputEvent) CompleteMotionGestureInputEvent_() *CompleteMotionGestureInputEvent {
    return this
}


func (this *CompleteMotionGestureInputEvent) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Name = StringUtil_GetInstance().EMPTY_STRING
    this.MotionGestureInput = TouchMotionGestureFactory_GetInstance().NO_MOTION
}


func NewCompleteMotionGestureInputEvent(name string, motionGestureInput *MotionGestureInput) *CompleteMotionGestureInputEvent {
    this := &CompleteMotionGestureInputEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(motionGestureInput)
        this.initFields()
        this.SetName(name)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "InputToGameKeyEventAction", e)
                }
            }()
            this.SetMotionGestureInput(motionGestureInput)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *CompleteMotionGestureInputEvent) SetName(name string) {
        this.Name = name
}


func (this *CompleteMotionGestureInputEvent) GetName() string {
        return this.Name
}


func (this *CompleteMotionGestureInputEvent) SetMotionGestureInput(motionGestureInput *MotionGestureInput) {
        this.MotionGestureInput = motionGestureInput
}


func (this *CompleteMotionGestureInputEvent) GetMotionGestureInput() *MotionGestureInput {
        return this.MotionGestureInput
}

