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

type VirtualKeyboardEventHandler struct {
    BasicEventHandler
    SHOW_EVENT *VirtualKeyboardEvent
    HIDE_EVENT *VirtualKeyboardEvent
}

var VirtualKeyboardEventHandler_GameKeyEventHandler *VirtualKeyboardEventHandler = NewVirtualKeyboardEventHandler()

func (this *VirtualKeyboardEventHandler) initFields() {
    this.SHOW_EVENT = NewVirtualKeyboardEvent(BooleanFactory_GetInstance().TRUE)
    this.HIDE_EVENT = NewVirtualKeyboardEvent(BooleanFactory_GetInstance().FALSE)
}


func NewVirtualKeyboardEventHandler() *VirtualKeyboardEventHandler {
    this := &VirtualKeyboardEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func VirtualKeyboardEventHandler_GetInstance() *VirtualKeyboardEventHandler {
        return VirtualKeyboardEventHandler_GameKeyEventHandler
}


// throws Exception
func (this *VirtualKeyboardEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var virtualKeyboardEventListenerInterface VirtualKeyboardEventListenerInterface = any(eventListenerInterface).(VirtualKeyboardEventListenerInterface); _ = virtualKeyboardEventListenerInterface
        virtualKeyboardEventListenerInterface.OnVirtualKeyboardEvent(any(eventObject).(*VirtualKeyboardEvent))
}


func (this *VirtualKeyboardEventHandler) Open() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "open", e)
                }
            }()
            time.Sleep(time.Duration(120) * time.Millisecond)
            this.FireEvent(this.SHOW_EVENT)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *VirtualKeyboardEventHandler) Close() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CLOSE, e)
                }
            }()
            time.Sleep(time.Duration(120) * time.Millisecond)
            this.FireEvent(this.HIDE_EVENT)
            return
        }(); tryControl1 == 1 {
            return
        }
}

