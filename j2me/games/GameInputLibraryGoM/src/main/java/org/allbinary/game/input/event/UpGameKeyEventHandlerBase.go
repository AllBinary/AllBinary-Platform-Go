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

type UpGameKeyEventHandlerBase struct {
    BasicEventHandler
    List IBasicArrayList
}

// The Java class UpGameKeyEventHandlerBase has subclasses so its variables are this interface
type IUpGameKeyEventHandlerBase interface {
    IBasicEventHandler
    UpGameKeyEventHandlerBase_() *UpGameKeyEventHandlerBase
    AddListener(playerGameInput IPlayerGameInput)
    RemoveAllListeners()
    RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface)
    RemoveListener(eventListenerInterface EventListenerInterface)
    FireEvent(eventObject IAllBinaryEventObject)
    Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface)
    ToString() string
}

func (this *UpGameKeyEventHandlerBase) UpGameKeyEventHandlerBase_() *UpGameKeyEventHandlerBase {
    return this
}

var UpGameKeyEventHandlerBase_TOTAL_LISTENERS string = " Total PlayerGameInput Listeners: "
var UpGameKeyEventHandlerBase_LISTENER_LABEL string = " PlayerGameInput Listener: "

func (this *UpGameKeyEventHandlerBase) initFields() {
    this.List = NewBasicArrayListD()
}


func NewUpGameKeyEventHandlerBase() *UpGameKeyEventHandlerBase {
    this := &UpGameKeyEventHandlerBase{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *UpGameKeyEventHandlerBase) AddListener(playerGameInput IPlayerGameInput) {
        if !this.List.Contains(playerGameInput) {
            this.List.Add(playerGameInput)
        }
}


func (this *UpGameKeyEventHandlerBase) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *UpGameKeyEventHandlerBase) RemoveListenerSingleThreaded(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListenerSingleThreaded(eventListenerInterface)
}


// synchronized
func (this *UpGameKeyEventHandlerBase) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *UpGameKeyEventHandlerBase) FireEvent(eventObject IAllBinaryEventObject) {
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
                playerGameInput.OnUpGameKeyEvent(any(eventObject).(*GameKeyEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *UpGameKeyEventHandlerBase) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var upGameKeyEventListenerInterface UpGameKeyEventListenerInterface = any(eventListenerInterface).(UpGameKeyEventListenerInterface); _ = upGameKeyEventListenerInterface
        upGameKeyEventListenerInterface.OnUpGameKeyEvent(any(eventObject).(*GameKeyEvent))
}


func (this *UpGameKeyEventHandlerBase) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var size int = this.List.Size(); _ = size
        stringBuffer.Append(this.BasicEventHandler.ToString())
        stringBuffer.Append(UpGameKeyEventHandlerBase_TOTAL_LISTENERS)
        stringBuffer.Appendint(size)
        for index := int(0); index < size; index++ {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.TOSTRING, e)
                    }
                }()
                var eventListenerInterface EventListenerInterface = any(this.List.Get(index)).(EventListenerInterface); _ = eventListenerInterface
                stringBuffer.Append(UpGameKeyEventHandlerBase_LISTENER_LABEL)
                stringBuffer.Append(fmt.Sprint(eventListenerInterface))
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        }
        return fmt.Sprint(stringBuffer)
}

func (this *UpGameKeyEventHandlerBase) String() string {
    return this.ToString()
}

