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

type DownGameKeyEventHandlerBase struct {
    BasicEventHandler
    List IBasicArrayList
}

// The Java class DownGameKeyEventHandlerBase has subclasses so its variables are this interface
type IDownGameKeyEventHandlerBase interface {
    IBasicEventHandler
    DownGameKeyEventHandlerBase_() *DownGameKeyEventHandlerBase
    AddListener(playerGameInput IPlayerGameInput)
    RemoveAllListeners()
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEvent(eventObject IAllBinaryEventObject)
    Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface)
    ToString() string
}

func (this *DownGameKeyEventHandlerBase) DownGameKeyEventHandlerBase_() *DownGameKeyEventHandlerBase {
    return this
}

var DownGameKeyEventHandlerBase_TOTAL_LISTENERS string = " Total PlayerGameInput Listeners: "
var DownGameKeyEventHandlerBase_LISTENER_LABEL string = " PlayerGameInput Listener: "

func (this *DownGameKeyEventHandlerBase) initFields() {
    this.List = NewBasicArrayListD()
}


func NewDownGameKeyEventHandlerBase() *DownGameKeyEventHandlerBase {
    this := &DownGameKeyEventHandlerBase{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *DownGameKeyEventHandlerBase) AddListener(playerGameInput IPlayerGameInput) {
        if !this.List.Contains(playerGameInput) {
            this.List.Add(playerGameInput)
        }
}


func (this *DownGameKeyEventHandlerBase) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *DownGameKeyEventHandlerBase) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
}


// synchronized
func (this *DownGameKeyEventHandlerBase) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *DownGameKeyEventHandlerBase) FireEvent(eventObject IAllBinaryEventObject) {
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
                playerGameInput.OnDownGameKeyEvent(any(eventObject).(*GameKeyEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *DownGameKeyEventHandlerBase) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var downGameKeyEventListenerInterface DownGameKeyEventListenerInterface = any(eventListenerInterface).(DownGameKeyEventListenerInterface); _ = downGameKeyEventListenerInterface
        downGameKeyEventListenerInterface.OnDownGameKeyEvent(any(eventObject).(*GameKeyEvent))
}


func (this *DownGameKeyEventHandlerBase) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicEventHandler.ToString())
        stringBuffer.Append(DownGameKeyEventHandlerBase_TOTAL_LISTENERS)
        stringBuffer.Appendint(this.List.Size())
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var size int = this.List.Size(); _ = size
        var index int = 0; _ = index
        for index < size {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.TOSTRING, e)
                    }
                }()
                eventListenerInterface = any(this.List.Get(index)).(EventListenerInterface)
                stringBuffer.Append(DownGameKeyEventHandlerBase_LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            index++
        }
        return fmt.Sprint(stringBuffer)
}

func (this *DownGameKeyEventHandlerBase) String() string {
    return this.ToString()
}

