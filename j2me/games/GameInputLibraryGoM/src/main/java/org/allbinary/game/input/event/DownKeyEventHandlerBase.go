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

type DownKeyEventHandlerBase struct {
    BasicEventHandler
    List IBasicArrayList
}

// The Java class DownKeyEventHandlerBase has subclasses so its variables are this interface
type IDownKeyEventHandlerBase interface {
    IBasicEventHandler
    DownKeyEventHandlerBase_() *DownKeyEventHandlerBase
    AddListener(playerGameInput IPlayerGameInput)
    RemoveAllListeners()
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEventI(eventObject int)
    FireEventForGameKeyEvent(eventObject *GameKeyEvent)
    ProcessI(eventObject int, eventListenerInterface EventListenerInterface)
    ProcessEvent(eventObject *GameKeyEvent, eventListenerInterface EventListenerInterface)
    ToString() string
}

func (this *DownKeyEventHandlerBase) DownKeyEventHandlerBase_() *DownKeyEventHandlerBase {
    return this
}

var DownKeyEventHandlerBase_TOTAL_LISTENERS string = " Total PlayerGameInput Listeners: "
var DownKeyEventHandlerBase_LISTENER_LABEL string = " PlayerGameInput Listener: "

func (this *DownKeyEventHandlerBase) initFields() {
    this.List = NewBasicArrayListD()
}


func NewDownKeyEventHandlerBase() *DownKeyEventHandlerBase {
    this := &DownKeyEventHandlerBase{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *DownKeyEventHandlerBase) AddListener(playerGameInput IPlayerGameInput) {
        if !this.List.Contains(playerGameInput) {
            this.List.Add(playerGameInput)
        }
}


func (this *DownKeyEventHandlerBase) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *DownKeyEventHandlerBase) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
}


// synchronized
func (this *DownKeyEventHandlerBase) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *DownKeyEventHandlerBase) FireEventI(eventObject int) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var playerGameInput IPlayerGameInput = any(this.List.BasicArrayList_().ObjectArray[index]).(IPlayerGameInput); _ = playerGameInput
                playerGameInput.OnDownKey(eventObject)
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        var index int = 0; _ = index
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(EventListenerInterface)
                this.ProcessI(eventObject, eventListenerInterface)
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *DownKeyEventHandlerBase) FireEventForGameKeyEvent(eventObject *GameKeyEvent) {
        var size int = this.List.Size(); _ = size
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var playerGameInput IPlayerGameInput = any(this.List.BasicArrayList_().ObjectArray[index]).(IPlayerGameInput); _ = playerGameInput
                playerGameInput.OnDownKeyEvent(eventObject)
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        var index int = 0; _ = index
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        for index < this.EventListenerInterfaceList.Size() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(this.EventListenerInterfaceList.Get(index)).(EventListenerInterface)
                this.ProcessEvent(eventObject, eventListenerInterface)
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *DownKeyEventHandlerBase) ProcessI(eventObject int, eventListenerInterface EventListenerInterface) {
        var downKeyEventListenerInterface DownKeyEventListenerInterface = any(eventListenerInterface).(DownKeyEventListenerInterface); _ = downKeyEventListenerInterface
        downKeyEventListenerInterface.OnDownKey(eventObject)
}


// throws Exception
func (this *DownKeyEventHandlerBase) ProcessEvent(eventObject *GameKeyEvent, eventListenerInterface EventListenerInterface) {
        var downKeyEventListenerInterface DownKeyEventListenerInterface = any(eventListenerInterface).(DownKeyEventListenerInterface); _ = downKeyEventListenerInterface
        downKeyEventListenerInterface.OnDownKeyEvent(eventObject)
}


func (this *DownKeyEventHandlerBase) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicEventHandler.ToString())
        stringBuffer.Append(DownKeyEventHandlerBase_TOTAL_LISTENERS)
        stringBuffer.Appendint(this.List.Size())
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        for index := int(0); index < this.List.Size(); index++ {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.TOSTRING, e)
                    }
                }()
                eventListenerInterface = any(this.List.Get(index)).(EventListenerInterface)
                stringBuffer.Append(DownKeyEventHandlerBase_LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        }
        return fmt.Sprint(stringBuffer)
}

func (this *DownKeyEventHandlerBase) String() string {
    return this.ToString()
}

