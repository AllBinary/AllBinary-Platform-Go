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

type BasicEventHandler struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    EventStrings *EventStrings
    EventListenerInterfaceList IBasicArrayList
}

// Implements BasicEventHandlerInterface

// The Java class BasicEventHandler has subclasses so its variables are this interface
type IBasicEventHandler interface {
    BasicEventHandler_() *BasicEventHandler
    RemoveAllListeners()
    AddListeners(vector IBasicArrayList)
    RemoveListeners(vector IBasicArrayList)
    AddListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    AddListenerInterface(eventListenerInterface EventListenerInterface)
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEvent(eventObject IAllBinaryEventObject)
    Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface)
    GetEventListenerInterfaceListP() IBasicArrayList
    ToString() string
}

func (this *BasicEventHandler) BasicEventHandler_() *BasicEventHandler {
    return this
}


func (this *BasicEventHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.EventStrings = EventStrings_GetInstance()
    
}


func NewBasicEventHandler() *BasicEventHandler {
    this := &BasicEventHandler{}
    this.initFields()
        this.EventListenerInterfaceList = NewBasicArrayListD()
    return this
}


func (this *BasicEventHandler) RemoveAllListeners() {
        this.EventListenerInterfaceList = NewBasicArrayListD()
}


func (this *BasicEventHandler) AddListeners(vector IBasicArrayList) {
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            eventListenerInterface = any(vector.Get(index)).(EventListenerInterface)
            this.AddListenerInterface(eventListenerInterface)
        }
}


func (this *BasicEventHandler) RemoveListeners(vector IBasicArrayList) {
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            eventListenerInterface = any(vector.Get(index)).(EventListenerInterface)
            this.RemoveListener(eventListenerInterface)
        }
}


func (this *BasicEventHandler) AddListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


func (this *BasicEventHandler) AddListenerInterface(eventListenerInterface EventListenerInterface) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


func (this *BasicEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


func (this *BasicEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


// throws Exception
func (this *BasicEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var index int = 0; _ = index
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.EventStrings.FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(EventListenerInterface)
                this.Process(eventObject, eventListenerInterface)
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *BasicEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        eventListenerInterface.OnEvent(eventObject)
}


func (this *BasicEventHandler) GetEventListenerInterfaceListP() IBasicArrayList {
        return this.EventListenerInterfaceList
}


func (this *BasicEventHandler) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.EventStrings.TOTAL_LISTENERS)
        stringBuffer.Appendint(this.EventListenerInterfaceList.Size())
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var index int = 0; _ = index
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.TOSTRING, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(EventListenerInterface)
                stringBuffer.Append(this.EventStrings.LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            index++
        }
        return fmt.Sprint(stringBuffer)
}

func (this *BasicEventHandler) String() string {
    return this.ToString()
}
