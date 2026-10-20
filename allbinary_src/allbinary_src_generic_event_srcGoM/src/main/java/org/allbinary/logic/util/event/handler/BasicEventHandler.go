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
    EventListenerInterfaceList IBasicArrayList
}

// Implements BasicEventHandlerInterface

// The Java class BasicEventHandler has subclasses so its variables are this interface
type IBasicEventHandler interface {
    BasicEventHandler_() *BasicEventHandler
    ToString() string
    RemoveAllListeners()
    AddListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    AddListenerInterface(eventListenerInterface EventListenerInterface)
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEvent(eventObject IAllBinaryEventObject)
    Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface)
    GetEventListenerInterfaceListP() IBasicArrayList
}

func (this *BasicEventHandler) BasicEventHandler_() *BasicEventHandler {
    return this
}

var BasicEventHandler_PERFORMANCE_MESSAGE string = "Use Custom onEvent Methods for needed optimization"

func (this *BasicEventHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewBasicEventHandler() *BasicEventHandler {
    this := &BasicEventHandler{}
    this.initFields()
        this.EventListenerInterfaceList = NewBasicArrayListD()
    return this
}


func (this *BasicEventHandler) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var LISTENER string = " Listener: "; _ = LISTENER
        stringBuffer.Append("Total Listeners: ")
        stringBuffer.Appendint(this.EventListenerInterfaceList.Size())
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var index int = 0; _ = index
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.TOSTRING, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.BasicArrayList_().ObjectArray[index]).(EventListenerInterface)
                stringBuffer.Append(LISTENER)
                stringBuffer.Append(StringUtil_GetInstance().ToString(eventListenerInterface))
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


// synchronized
func (this *BasicEventHandler) RemoveAllListeners() {
        this.EventListenerInterfaceList = NewBasicArrayListD()
}


func (this *BasicEventHandler) AddListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


// synchronized
func (this *BasicEventHandler) AddListenerInterface(eventListenerInterface EventListenerInterface) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


func (this *BasicEventHandler) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


// synchronized
func (this *BasicEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


// throws Exception
// synchronized
func (this *BasicEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        var index int = 0; _ = index
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.BasicArrayList_().ObjectArray[index]).(EventListenerInterface)
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

