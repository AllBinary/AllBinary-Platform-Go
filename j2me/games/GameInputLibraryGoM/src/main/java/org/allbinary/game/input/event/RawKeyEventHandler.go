/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type RawKeyEventHandler struct {
    LogUtil *LogUtil
    EventListenerInterfaceList IBasicArrayList
}

var RawKeyEventHandler_Instance *RawKeyEventHandler = NewRawKeyEventHandler()
var RawKeyEventHandler_TOTAL_LISTENERS string = "Total Listeners: "
var RawKeyEventHandler_LISTENER_LABEL string = " Listener: "

func (this *RawKeyEventHandler) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func RawKeyEventHandler_GetInstance() *RawKeyEventHandler {
        return RawKeyEventHandler_Instance
}


func NewRawKeyEventHandler() *RawKeyEventHandler {
    this := &RawKeyEventHandler{}
    this.initFields()
        this.EventListenerInterfaceList = NewBasicArrayListD()
    return this
}


func (this *RawKeyEventHandler) RemoveAllListeners() {
        this.EventListenerInterfaceList = NewBasicArrayListD()
}


func (this *RawKeyEventHandler) AddListeners(vector IBasicArrayList) {
        var eventListenerInterface RawKeyEventListener; _ = eventListenerInterface
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            eventListenerInterface = any(vector.Get(index)).(RawKeyEventListener)
            this.AddListener(eventListenerInterface)
        }
}


func (this *RawKeyEventHandler) AddListenerSingleThreaded(eventListenerInterface RawKeyEventListener) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


func (this *RawKeyEventHandler) AddListener(eventListenerInterface RawKeyEventListener) {
        if !this.EventListenerInterfaceList.Contains(eventListenerInterface) {
            this.EventListenerInterfaceList.Add(eventListenerInterface)
        }
}


func (this *RawKeyEventHandler) RemoveListenerSingleThreaded(eventListenerInterface RawKeyEventListener) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


func (this *RawKeyEventHandler) RemoveListener(eventListenerInterface RawKeyEventListener) {
        this.EventListenerInterfaceList.Remove(eventListenerInterface)
}


// throws Exception
func (this *RawKeyEventHandler) FireEvent(keyCode int, deviceId int, repeated bool) {
        var eventListenerInterface RawKeyEventListener; _ = eventListenerInterface
        var index int = 0; _ = index
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
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(RawKeyEventListener)
                this.Process(keyCode, deviceId, repeated, eventListenerInterface)
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *RawKeyEventHandler) Process(keyCode int, deviceId int, repeated bool, eventListenerInterface RawKeyEventListener) {
        eventListenerInterface.OnEventRaw(keyCode, deviceId, repeated)
}


func (this *RawKeyEventHandler) GetEventListenerInterfaceListP() IBasicArrayList {
        return this.EventListenerInterfaceList
}


func (this *RawKeyEventHandler) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(RawKeyEventHandler_TOTAL_LISTENERS)
        stringBuffer.Appendint(this.EventListenerInterfaceList.Size())
        var eventListenerInterface RawKeyEventListener; _ = eventListenerInterface
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
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(RawKeyEventListener)
                stringBuffer.Append(RawKeyEventHandler_LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            index++
        }
        return fmt.Sprint(stringBuffer)
}

func (this *RawKeyEventHandler) String() string {
    return this.ToString()
}

