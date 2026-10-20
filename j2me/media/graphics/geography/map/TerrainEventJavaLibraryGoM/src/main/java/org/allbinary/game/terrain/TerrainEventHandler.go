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

type TerrainEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var TerrainEventHandler_Hashtable *ABHashtable[any, any] = StdUtil_GetInstance().NULL_TABLE

func (this *TerrainEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func TerrainEventHandler_Init() {
        TerrainEventHandler_Hashtable = StdUtil_GetInstance().CreateHashtable()
}


func TerrainEventHandler_GetInstance(anyType any) *TerrainEventHandler {
        var eventHandlerCanBeNull any = TerrainEventHandler_Hashtable.Get(anyType); _ = eventHandlerCanBeNull
        if eventHandlerCanBeNull == nil {
            eventHandlerCanBeNull = NewTerrainEventHandler()
            TerrainEventHandler_Hashtable.Put(anyType, eventHandlerCanBeNull)
        }
        return any(eventHandlerCanBeNull).(*TerrainEventHandler)
}


func NewTerrainEventHandler() *TerrainEventHandler {
    this := &TerrainEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *TerrainEventHandler) AddListener(terrainEventListener *TerrainEventListener) {
        if !this.List.Contains(terrainEventListener) {
            this.List.Add(terrainEventListener)
        }
}


func (this *TerrainEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *TerrainEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


// throws Exception
func (this *TerrainEventHandler) FireEvent(eventObject IAllBinaryEventObject) {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var terrainEventListener *TerrainEventListener = any(this.List.Get(index)).(*TerrainEventListener); _ = terrainEventListener
                terrainEventListener.OnTerrainEvent(any(eventObject).(*TerrainEvent))
                return
            }(); tryControl1 == 1 {
                return
            }
        }
        this.BasicEventHandler.FireEvent(eventObject)
}


// throws Exception
func (this *TerrainEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var terrainEventListenerInterface TerrainEventListenerInterface = any(eventListenerInterface).(TerrainEventListenerInterface); _ = terrainEventListenerInterface
        terrainEventListenerInterface.OnTerrainEvent(any(eventObject).(*TerrainEvent))
}

