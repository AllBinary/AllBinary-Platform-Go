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

type GameLevelDisplayChangeEventListenersFactory struct {
    List IBasicArrayList
}

var GameLevelDisplayChangeEventListenersFactory_Instance *GameLevelDisplayChangeEventListenersFactory = NewGameLevelDisplayChangeEventListenersFactory()

func (this *GameLevelDisplayChangeEventListenersFactory) initFields() {
    this.List = NewBasicArrayListD()
}

func NewGameLevelDisplayChangeEventListenersFactory() *GameLevelDisplayChangeEventListenersFactory {
    this := &GameLevelDisplayChangeEventListenersFactory{}
    this.initFields()
    return this
}


func GameLevelDisplayChangeEventListenersFactory_GetInstance() *GameLevelDisplayChangeEventListenersFactory {
        return GameLevelDisplayChangeEventListenersFactory_Instance
}


func (this *GameLevelDisplayChangeEventListenersFactory) Add(displayChangeEventListener DisplayChangeEventListener) {
        this.List.Add(displayChangeEventListener)
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(displayChangeEventListener)
}


func (this *GameLevelDisplayChangeEventListenersFactory) Clear() {
        var displayChangeEventHandler IBasicEventHandler = DisplayChangeEventHandler_GetInstance(); _ = displayChangeEventHandler
        var displayChangeEventListener DisplayChangeEventListener; _ = displayChangeEventListener
        for index := int(this.List.Size()); Java_PreDecrement(&index) >= 0;  {
            displayChangeEventListener = any(this.List.BasicArrayList_().ObjectArray[index]).(DisplayChangeEventListener)
            displayChangeEventHandler.RemoveListener(displayChangeEventListener)
        }
        this.List.Clear()
}

