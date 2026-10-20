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

type GameMotionGestureListener struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    MotionGestureCollection IBasicArrayList
    TouchGestureCollection IBasicArrayList
    Signed CompleteMotionGestureListenerInterface
}

// Implements MotionGestureEventListener


func (this *GameMotionGestureListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.MotionGestureCollection = NewBasicArrayListD()
    this.TouchGestureCollection = NewBasicArrayListD()
    
}


func NewGameMotionGestureListener(signed CompleteMotionGestureListenerInterface) *GameMotionGestureListener {
    this := &GameMotionGestureListener{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.Signed = signed
    return this
}


func (this *GameMotionGestureListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


func (this *GameMotionGestureListener) OnUpMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnDownMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnDiagonalDownRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnDiagonalDownLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnDiagonalUpRightMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnDiagonalUpLeftMotionGestureEvent(ev *MotionGestureEvent) {
        this.OnMotionGestureEvent(ev)
}


func (this *GameMotionGestureListener) OnPressedMotionGestureEvent(ev *MotionGestureEvent) {
        this.TouchGestureCollection.Add(ev.GetMotionGesture())
}


func (this *GameMotionGestureListener) Released(ev *MotionGestureEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.CommonStrings.EXCEPTION_LABEL)
            stringBuffer.Append(StringUtil_GetInstance().ToString(ev.GetMotionGesture()))
            this.LogUtil.Put(fmt.Sprint(stringBuffer), this, "release", e)
                }
            }()
            this.TouchGestureCollection.Add(ev.GetMotionGesture())
            this.Signed.OnMotionGestureCompleted(this.TouchGestureCollection)
            this.TouchGestureCollection.Clear()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameMotionGestureListener) OnScrolledMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
}


func (this *GameMotionGestureListener) OnMotionGestureEvent(ev *MotionGestureEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.CommonStrings.EXCEPTION_LABEL)
            stringBuffer.Append(StringUtil_GetInstance().ToString(ev.GetMotionGesture()))
            this.LogUtil.Put(fmt.Sprint(stringBuffer), this, "onMotionGestureEvent", e)
                }
            }()
            var motionGestureInput *MotionGestureInput = ev.GetMotionGesture(); _ = motionGestureInput
            if motionGestureInput == TouchMotionGestureFactory_GetInstance().PRESSED {
                this.OnPressedMotionGestureEvent(ev)
            } else if motionGestureInput == TouchMotionGestureFactory_GetInstance().RELEASED {
                this.Released(ev)
            } else {
                this.MotionGestureCollection.Add(motionGestureInput)
                this.Signed.OnMotionGestureCompleted(this.MotionGestureCollection)
                this.MotionGestureCollection.Clear()
                this.TouchGestureCollection.Clear()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

