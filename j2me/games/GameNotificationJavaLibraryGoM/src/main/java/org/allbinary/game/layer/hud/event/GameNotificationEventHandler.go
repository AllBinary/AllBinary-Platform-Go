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

type GameNotificationEventHandler struct {
    BasicEventHandler
    Enabled bool
}

var GameNotificationEventHandler_Instance *GameNotificationEventHandler = NewGameNotificationEventHandler()

func (this *GameNotificationEventHandler) initFields() {
    this.Enabled = true
}


func NewGameNotificationEventHandler() *GameNotificationEventHandler {
    this := &GameNotificationEventHandler{}
    this.BasicEventHandler = *NewBasicEventHandler()
    this.initFields()
    return this
}


// synchronized
func (this *GameNotificationEventHandler) RemoveAllListeners() {
        var eventListenerInterface EventListenerInterface = NullEventListener_NULL_EVENT_LISTENER; _ = eventListenerInterface
        var list IBasicArrayList = this.EventListenerInterfaceList; _ = list
        if list.Size() > 0 {
            eventListenerInterface = any(list.BasicArrayList_().ObjectArray[0]).(EventListenerInterface)
        }
        this.BasicEventHandler.RemoveAllListeners()
        if eventListenerInterface != NullEventListener_NULL_EVENT_LISTENER {
            this.AddListenerInterface(eventListenerInterface)
        }
}


func GameNotificationEventHandler_GetInstance() *GameNotificationEventHandler {
        return GameNotificationEventHandler_Instance
}


// throws Exception
func (this *GameNotificationEventHandler) Process(eventObject IAllBinaryEventObject, eventListenerInterface EventListenerInterface) {
        if !this.Enabled {
            return
        }
        var gameNotificationListenerInterface GameNotificationListenerInterface = any(eventListenerInterface).(GameNotificationListenerInterface); _ = gameNotificationListenerInterface
        gameNotificationListenerInterface.OnGameNotificationEvent(any(eventObject).(*GameNotificationEvent))
}

