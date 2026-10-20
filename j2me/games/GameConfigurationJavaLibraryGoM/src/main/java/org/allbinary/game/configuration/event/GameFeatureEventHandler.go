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

type GameFeatureEventHandler struct {
    BasicEventHandler
}

var GameFeatureEventHandler_GameKeyEventHandler *GameFeatureEventHandler = NewGameFeatureEventHandler()


func NewGameFeatureEventHandler() *GameFeatureEventHandler {
    this := &GameFeatureEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    return this
}


func GameFeatureEventHandler_GetInstance() *GameFeatureEventHandler {
        return GameFeatureEventHandler_GameKeyEventHandler
}


// throws Exception
func (this *GameFeatureEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        var gameFeatureListenerInterface GameFeatureListenerInterface = (any(eventListenerInterface).(GameFeatureListenerInterface)); _ = gameFeatureListenerInterface
        gameFeatureListenerInterface.OnGameFeatureChange(any(eventObject).(*GameFeatureEvent))
}

