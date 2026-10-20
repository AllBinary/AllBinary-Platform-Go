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

type ReleaseControlledMotionGestureListener struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CurrentMotionGesture *MotionGestureInput
    IsMouseGestureOccurring bool
    MotionGestureCollection IBasicArrayList
    Signed CompleteMotionGestureListenerInterface
}

// Implements MotionGestureEventListener


func (this *ReleaseControlledMotionGestureListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CurrentMotionGesture = TouchMotionGestureFactory_GetInstance().NO_MOTION
    this.IsMouseGestureOccurring = false
    this.MotionGestureCollection = NewBasicArrayListD()
    
}


func NewReleaseControlledMotionGestureListener(signed CompleteMotionGestureListenerInterface) *ReleaseControlledMotionGestureListener {
    this := &ReleaseControlledMotionGestureListener{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.Signed = signed
    return this
}


func (this *ReleaseControlledMotionGestureListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


func (this *ReleaseControlledMotionGestureListener) OnUpMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnDownMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnDiagonalDownRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnDiagonalDownLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnDiagonalUpRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnDiagonalUpLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnPressedMotionGestureEvent(ev *MotionGestureEvent) {
        this.IsMouseGestureOccurring = true
        this.CurrentMotionGesture = TouchMotionGestureFactory_GetInstance().NO_MOTION
        this.OnMotionGestureEvent(ev)
}


func (this *ReleaseControlledMotionGestureListener) OnScrolledMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
}


func (this *ReleaseControlledMotionGestureListener) Released(ev *MotionGestureEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(StringUtil_GetInstance().ToString(ev.GetMotionGesture()))), this, "release", e)
                }
            }()
            if this.IsMouseGestureOccurring == false {
                tryControl1 = 1; return
            }
            this.IsMouseGestureOccurring = false
            this.Signed.OnMotionGestureCompleted(this.MotionGestureCollection)
            this.MotionGestureCollection.Clear()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ReleaseControlledMotionGestureListener) OnMotionGestureEvent(ev *MotionGestureEvent) {
        if this.IsMouseGestureOccurring == false {
            return
        }
        var motionGestureInput *MotionGestureInput = ev.GetMotionGesture(); _ = motionGestureInput
        if motionGestureInput == TouchMotionGestureFactory_GetInstance().PRESSED {
            this.OnPressedMotionGestureEvent(ev)
        } else if motionGestureInput == TouchMotionGestureFactory_GetInstance().RELEASED {
            this.Released(ev)
        } else {
            if this.CurrentMotionGesture != motionGestureInput {
                this.CurrentMotionGesture = motionGestureInput
                this.MotionGestureCollection.Add(ev.GetMotionGesture())
            }
        }
}

