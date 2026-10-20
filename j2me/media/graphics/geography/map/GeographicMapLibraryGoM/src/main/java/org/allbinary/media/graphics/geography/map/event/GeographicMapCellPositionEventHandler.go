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

type GeographicMapCellPositionEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var GeographicMapCellPositionEventHandler_Instance *GeographicMapCellPositionEventHandler = NewGeographicMapCellPositionEventHandler()

func (this *GeographicMapCellPositionEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func GeographicMapCellPositionEventHandler_GetInstance() *GeographicMapCellPositionEventHandler {
        return GeographicMapCellPositionEventHandler_Instance
}


func NewGeographicMapCellPositionEventHandler() *GeographicMapCellPositionEventHandler {
    this := &GeographicMapCellPositionEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *GeographicMapCellPositionEventHandler) AddListener(miniMapLayer IMiniMapLayer) {
        if !this.List.Contains(miniMapLayer) {
            this.List.Add(miniMapLayer)
        }
}


func (this *GeographicMapCellPositionEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *GeographicMapCellPositionEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *GeographicMapCellPositionEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var miniMapLayer IMiniMapLayer = any(this.List.Get(index)).(IMiniMapLayer); _ = miniMapLayer
                miniMapLayer.OnGeographicMapCellPositionEvent(any(eventObject).(*GeographicMapCellPositionEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
// synchronized
func (this *GeographicMapCellPositionEventHandler) FireRemoveEvent(eventObject IAllBinaryEventObject) {
        var list IBasicArrayList = this.EventListenerInterfaceList; _ = list
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var miniMapLayer IMiniMapLayer = any(this.List.Get(index)).(IMiniMapLayer); _ = miniMapLayer
                miniMapLayer.OnRemoveGeographicMapCellPositionEvent(any(eventObject).(*GeographicMapCellPositionEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        var eventListenerInterface EventListenerInterface; _ = eventListenerInterface
        var geographicMapCellPositionEventListenerInterface GeographicMapCellPositionEventListenerInterface; _ = geographicMapCellPositionEventListenerInterface
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                eventListenerInterface = any(list.Get(index)).(EventListenerInterface)
                geographicMapCellPositionEventListenerInterface = (any(eventListenerInterface).(GeographicMapCellPositionEventListenerInterface))
                geographicMapCellPositionEventListenerInterface.OnRemoveGeographicMapCellPositionEvent(any(eventObject).(*GeographicMapCellPositionEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
}


// throws Exception
func (this *GeographicMapCellPositionEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var geographicMapCellPositionEventListenerInterface GeographicMapCellPositionEventListenerInterface = (any(eventListenerInterface).(GeographicMapCellPositionEventListenerInterface)); _ = geographicMapCellPositionEventListenerInterface
        geographicMapCellPositionEventListenerInterface.OnGeographicMapCellPositionEvent(any(eventObject).(*GeographicMapCellPositionEvent))
}

