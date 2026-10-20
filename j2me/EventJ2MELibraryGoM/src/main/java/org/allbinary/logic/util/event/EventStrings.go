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

type EventStrings struct {
    PERFORMANCE_MESSAGE string
    TOTAL_LISTENERS string
    LISTENER_LABEL string
    FIRE_EVENT string
}

var EventStrings_Instance *EventStrings = NewEventStrings()

func (this *EventStrings) initFields() {
    this.PERFORMANCE_MESSAGE = "Use Custom onEvent Methods for needed optimization"
    this.TOTAL_LISTENERS = "Total Listeners: "
    this.LISTENER_LABEL = " Listener: "
    this.FIRE_EVENT = "fireEvent"
}

func NewEventStrings() *EventStrings {
    this := &EventStrings{}
    this.initFields()
    return this
}


func EventStrings_GetInstance() *EventStrings {
        return EventStrings_Instance
}
