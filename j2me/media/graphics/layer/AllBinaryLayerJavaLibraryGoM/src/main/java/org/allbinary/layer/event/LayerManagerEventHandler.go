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

type LayerManagerEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
    CREATE string
    DELETE string
}

var LayerManagerEventHandler_Instance *LayerManagerEventHandler = NewLayerManagerEventHandler()

func (this *LayerManagerEventHandler) initFields() {
    this.List = NewBasicArrayListD()
    this.CREATE = "Create"
    this.DELETE = "Delete"
}


func LayerManagerEventHandler_GetInstance() *LayerManagerEventHandler {
        return LayerManagerEventHandler_Instance
}


func NewLayerManagerEventHandler() *LayerManagerEventHandler {
    this := &LayerManagerEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *LayerManagerEventHandler) AddListener(layerManagerEventListener ILayerManagerEventListener) {
        if !this.List.Contains(layerManagerEventListener) {
            this.List.Add(layerManagerEventListener)
        }
}


func (this *LayerManagerEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *LayerManagerEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *LayerManagerEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var layerManagerEventListener ILayerManagerEventListener = any(this.List.BasicArrayList_().ObjectArray[index]).(ILayerManagerEventListener); _ = layerManagerEventListener
                layerManagerEventListener.OnCreateLayerManagerEvent(any(eventObject).(*LayerManagerEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
// synchronized
func (this *LayerManagerEventHandler) FireDeleteEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var layerManagerEventListener ILayerManagerEventListener = any(this.List.BasicArrayList_().ObjectArray[index]).(ILayerManagerEventListener); _ = layerManagerEventListener
                layerManagerEventListener.OnDeleteLayerManagerEvent(any(eventObject).(*LayerManagerEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        var eventListenerInterfaceList IBasicArrayList = this.EventListenerInterfaceList; _ = eventListenerInterfaceList
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var layerManagerEventListenerInterface LayerManagerEventListenerInterface; _ = layerManagerEventListenerInterface
        var index int = 0; _ = index
        for index < eventListenerInterfaceList.Size() {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(eventListenerInterfaceList.BasicArrayList_().ObjectArray[index]).(EventListenerInterface)
                layerManagerEventListenerInterface = (any(eventListenerInterface).(LayerManagerEventListenerInterface))
                layerManagerEventListenerInterface.OnDeleteLayerManagerEvent(any(eventObject).(*LayerManagerEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
            index++
        }
}


// throws Exception
func (this *LayerManagerEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var layerManagerEventListenerInterface LayerManagerEventListenerInterface = (any(eventListenerInterface).(LayerManagerEventListenerInterface)); _ = layerManagerEventListenerInterface
        layerManagerEventListenerInterface.OnCreateLayerManagerEvent(any(eventObject).(*LayerManagerEvent))
}

