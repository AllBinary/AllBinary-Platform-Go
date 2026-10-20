/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
 *  
 * By agreeing to this license you and any business entity you represent are 
 * legally bound to the AllBinary Open License Version 1 legal agreement. 
 *  
 * You may obtain the AllBinary Open License Version 1 legal agreement from 
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository. 
 *  
 * Created By: Travis Berthelot  
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

type TechEventHandler struct {
    BasicEventHandler
}

var TechEventHandler_EventHandler *TechEventHandler = NewTechEventHandler()


func NewTechEventHandler() *TechEventHandler {
    this := &TechEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    return this
}


func TechEventHandler_GetInstance() *TechEventHandler {
        return TechEventHandler_EventHandler
}


// throws Exception
func (this *TechEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var techEventListenerInterface TechEventListenerInterface = (any(eventListenerInterface).(TechEventListenerInterface)); _ = techEventListenerInterface
        techEventListenerInterface.OnTechEvent(eventObject)
}

