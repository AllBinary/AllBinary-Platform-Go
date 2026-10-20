/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type BuildingEventHandler struct {
    BasicEventHandler
}

var BuildingEventHandler_EventHandler *BuildingEventHandler = NewBuildingEventHandler()


func NewBuildingEventHandler() *BuildingEventHandler {
    this := &BuildingEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    return this
}


func BuildingEventHandler_GetInstance() *BuildingEventHandler {
        return BuildingEventHandler_EventHandler
}


// throws Exception
func (this *BuildingEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var buildingEventListenerInterface BuildingEventListenerInterface = (any(eventListenerInterface).(BuildingEventListenerInterface)); _ = buildingEventListenerInterface
        buildingEventListenerInterface.OnBuildingEvent(any(eventObject).(*RTSLayerEvent))
}

