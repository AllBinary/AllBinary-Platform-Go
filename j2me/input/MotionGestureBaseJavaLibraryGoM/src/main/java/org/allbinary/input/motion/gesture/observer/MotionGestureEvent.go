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

type MotionGestureEvent struct {
    AllBinaryEventObject
    MotionGesture *MotionGestureInput
    Id int
    PreviousPoint IGPoint
    CurrentPoint IGPoint
    CommonLabels *CommonLabels
}

var MotionGestureEvent_PREVIOUS string = " Previous: "
var MotionGestureEvent_HASHCODE string = " hashcode: "

func (this *MotionGestureEvent) initFields() {
    this.MotionGesture = TouchMotionGestureFactory_GetInstance().NO_MOTION
    
    this.PreviousPoint = PointFactory_GetInstance().ZERO_ZERO
    this.CurrentPoint = PointFactory_GetInstance().ZERO_ZERO
    this.CommonLabels = CommonLabels_GetInstance()
}


func NewMotionGestureEvent(source any, id int, motionGesture *MotionGestureInput) *MotionGestureEvent {
    this := &MotionGestureEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(source)
        this.initFields()
        this.Id = id
        this.SetMotionGesture(motionGesture)
    return this
}


func (this *MotionGestureEvent) GetMotionGesture() *MotionGestureInput {
        return this.MotionGesture
}


func (this *MotionGestureEvent) SetMotionGesture(motionGesture *MotionGestureInput) {
        this.MotionGesture = motionGesture
}


func (this *MotionGestureEvent) GetPreviousPoint() IGPoint {
        return this.PreviousPoint
}


func (this *MotionGestureEvent) SetPreviousPoint(previousPoint IGPoint) {
        this.PreviousPoint = previousPoint
}


func (this *MotionGestureEvent) GetCurrentPoint() IGPoint {
        return this.CurrentPoint
}


func (this *MotionGestureEvent) SetCurrentPoint(currentPoint IGPoint) {
        this.CurrentPoint = currentPoint
}


func (this *MotionGestureEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        stringBuffer.Append(stringUtil.ToString(this.MotionGesture))
        stringBuffer.Append(this.CommonLabels.CURRENT)
        stringBuffer.Append(stringUtil.ToString(this.CurrentPoint))
        stringBuffer.Append(MotionGestureEvent_PREVIOUS)
        stringBuffer.Append(stringUtil.ToString(this.PreviousPoint))
        stringBuffer.Append(MotionGestureEvent_HASHCODE)
        stringBuffer.Appendint(TsUtil_GetInstance().HashCode(this))
        return fmt.Sprint(stringBuffer)
}

func (this *MotionGestureEvent) String() string {
    return this.ToString()
}


func (this *MotionGestureEvent) GetId() int {
        return this.Id
}

