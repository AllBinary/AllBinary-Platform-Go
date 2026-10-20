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

type GeographicMapEventHandler struct {
    BasicEventHandler
    List IBasicArrayList
}

var GeographicMapEventHandler_Instance *GeographicMapEventHandler = NewGeographicMapEventHandler()

func (this *GeographicMapEventHandler) initFields() {
    this.List = NewBasicArrayListD()
}


func GeographicMapEventHandler_GetInstance() *GeographicMapEventHandler {
        return GeographicMapEventHandler_Instance
}


func NewGeographicMapEventHandler() *GeographicMapEventHandler {
    this := &GeographicMapEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


func (this *GeographicMapEventHandler) AddListener(gameLayer IAllBinaryGameLayer) {
        if !this.List.Contains(gameLayer) {
            this.List.Add(gameLayer)
        }
}


func (this *GeographicMapEventHandler) RemoveAllListeners() {
        this.List.Clear()
        this.BasicEventHandler.RemoveAllListeners()
}


func (this *GeographicMapEventHandler) RemoveListener(eventListenerInterface EventListenerInterface) {
        this.List.Remove(eventListenerInterface)
        this.BasicEventHandler.RemoveListener(eventListenerInterface)
}


func (this *GeographicMapEventHandler) FireEvent() {
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, EventStrings_GetInstance().FIRE_EVENT, e)
                    }
                }()
                var gameLayer IAllBinaryGameLayer = any(this.List.Get(index)).(IAllBinaryGameLayer); _ = gameLayer
                gameLayer.Move()
                return
            }(); tryControl1 == 1 {
                return
            }
        }
}

