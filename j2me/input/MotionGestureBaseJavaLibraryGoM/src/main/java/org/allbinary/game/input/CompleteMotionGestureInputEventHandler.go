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

type CompleteMotionGestureInputEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var CompleteMotionGestureInputEventHandler_SINGLETON *CompleteMotionGestureInputEventHandler = NewCompleteMotionGestureInputEventHandler()

func (this *CompleteMotionGestureInputEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func CompleteMotionGestureInputEventHandler_GetInstance() *CompleteMotionGestureInputEventHandler {
        return CompleteMotionGestureInputEventHandler_SINGLETON
}


func NewCompleteMotionGestureInputEventHandler() *CompleteMotionGestureInputEventHandler {
    this := &CompleteMotionGestureInputEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *CompleteMotionGestureInputEventHandler) AddListener(completeMotionGestureInputEventListener ICompleteMotionGestureInputEventListener) {
        if !this.List.Contains(completeMotionGestureInputEventListener) {
            this.List.Add(completeMotionGestureInputEventListener)
        }
}


func (this *CompleteMotionGestureInputEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *CompleteMotionGestureInputEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *CompleteMotionGestureInputEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var completeMotionGestureInputEventListener ICompleteMotionGestureInputEventListener = any(this.List.BasicArrayList_().ObjectArray[index]).(ICompleteMotionGestureInputEventListener); _ = completeMotionGestureInputEventListener
                completeMotionGestureInputEventListener.OnCompleteMotionGestureInputEvent(any(eventObject).(ICompleteMotionGestureInputEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *CompleteMotionGestureInputEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var completeMotionGestureInputEventListenerInterface CompleteMotionGestureInputEventListenerInterface = any(eventListenerInterface).(CompleteMotionGestureInputEventListenerInterface); _ = completeMotionGestureInputEventListenerInterface
        completeMotionGestureInputEventListenerInterface.OnCompleteMotionGestureInputEvent(any(eventObject).(ICompleteMotionGestureInputEvent))
}

