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

type PressGameKeyEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var PressGameKeyEventHandler_Instance *PressGameKeyEventHandler = NewPressGameKeyEventHandler()

func (this *PressGameKeyEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func PressGameKeyEventHandler_GetInstance() *PressGameKeyEventHandler {
        return PressGameKeyEventHandler_Instance
}


func NewPressGameKeyEventHandler() *PressGameKeyEventHandler {
    this := &PressGameKeyEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *PressGameKeyEventHandler) AddListener(playerGameInput IPlayerGameInput) {
        if !this.List.Contains(playerGameInput) {
            this.List.Add(playerGameInput)
        }
}


func (this *PressGameKeyEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *PressGameKeyEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *PressGameKeyEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
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
                playerGameInput.OnPressGameKeyEvent(any(eventObject).(*GameKeyEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *PressGameKeyEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var pressGameKeyEventListenerInterface PressGameKeyEventListenerInterface = (any(eventListenerInterface).(PressGameKeyEventListenerInterface)); _ = pressGameKeyEventListenerInterface
        pressGameKeyEventListenerInterface.OnPressGameKeyEvent(any(eventObject).(*GameKeyEvent))
}

