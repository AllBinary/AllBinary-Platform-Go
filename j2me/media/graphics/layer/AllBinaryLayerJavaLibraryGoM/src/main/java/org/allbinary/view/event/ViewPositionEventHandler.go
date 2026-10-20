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

type ViewPositionEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var ViewPositionEventHandler_SINGLETON *ViewPositionEventHandler = NewViewPositionEventHandler()

func (this *ViewPositionEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func ViewPositionEventHandler_GetInstance() *ViewPositionEventHandler {
        return ViewPositionEventHandler_SINGLETON
}


func NewViewPositionEventHandler() *ViewPositionEventHandler {
    this := &ViewPositionEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *ViewPositionEventHandler) AddListener(layerInterface IAllBinaryLayer) {
        if !this.List.Contains(layerInterface) {
            this.List.Add(layerInterface)
        }
}


func (this *ViewPositionEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *ViewPositionEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *ViewPositionEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var layerInterface IAllBinaryLayer = any(this.List.BasicArrayList_().ObjectArray[index]).(IAllBinaryLayer); _ = layerInterface
                layerInterface.OnChangeEvent(any(eventObject).(*ViewPositionEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *ViewPositionEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var viewPositionEventListenerInterface ViewPositionEventListenerInterface = (any(eventListenerInterface).(ViewPositionEventListenerInterface)); _ = viewPositionEventListenerInterface
        viewPositionEventListenerInterface.OnChangeEvent(any(eventObject).(*ViewPositionEvent))
}

