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

type UpKeyEventHandlerBase struct {
    BasicEventHandler
    List IBasicArrayList
}

// The Java class UpKeyEventHandlerBase has subclasses so its variables are this interface
type IUpKeyEventHandlerBase interface {
    IBasicEventHandler
    UpKeyEventHandlerBase_() *UpKeyEventHandlerBase
    AddListener(playerGameInput IPlayerGameInput)
    RemoveAllListeners()
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEventForInteger(eventObject int)
    ProcessForInteger(eventObject int, eventListenerInterface EventListenerInterface)
    ToString() string
}

func (this *UpKeyEventHandlerBase) UpKeyEventHandlerBase_() *UpKeyEventHandlerBase {
    return this
}

var UpKeyEventHandlerBase_TOTAL_LISTENERS string = " Total PlayerGameInput Listeners: "
var UpKeyEventHandlerBase_LISTENER_LABEL string = " PlayerGameInput Listener: "

func (this *UpKeyEventHandlerBase) initFields() {
    this.List = NewBasicArrayListD()
}


func NewUpKeyEventHandlerBase() *UpKeyEventHandlerBase {
    this := &UpKeyEventHandlerBase{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *UpKeyEventHandlerBase) AddListener(playerGameInput IPlayerGameInput) {
        if !this.List.Contains(playerGameInput) {
            this.List.Add(playerGameInput)
        }
}


func (this *UpKeyEventHandlerBase) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *UpKeyEventHandlerBase) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
}


// synchronized
func (this *UpKeyEventHandlerBase) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *UpKeyEventHandlerBase) FireEventForInteger(eventObject int) {
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
                playerGameInput.OnUpKeyEvent(eventObject)
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
                this.ProcessForInteger(eventObject, eventListenerInterface)
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *UpKeyEventHandlerBase) ProcessForInteger(eventObject int, eventListenerInterface EventListenerInterface) {
        var upKeyEventListenerInterface UpKeyEventListenerInterface = any(eventListenerInterface).(UpKeyEventListenerInterface); _ = upKeyEventListenerInterface
        upKeyEventListenerInterface.OnUpKeyEvent(eventObject)
}


func (this *UpKeyEventHandlerBase) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicEventHandler.ToString())
        stringBuffer.Append(UpKeyEventHandlerBase_TOTAL_LISTENERS)
        stringBuffer.Appendint(this.List.Size())
        for index := int(0); index < this.List.Size(); index++ {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.TOSTRING, e)
                    }
                }()
                var eventListenerInterface EventListenerInterface = any(this.List.Get(index)).(EventListenerInterface); _ = eventListenerInterface
                stringBuffer.Append(UpKeyEventHandlerBase_LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        }
        return fmt.Sprint(stringBuffer)
}

func (this *UpKeyEventHandlerBase) String() string {
    return this.ToString()
}

