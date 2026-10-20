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

type MyGraphicItemEventService struct {
}

var MyGraphicItemEventService_ListenerList *EventListenerList = NewEventListenerList()
var MyGraphicItemEventService_DELETE string = "delete"
var MyGraphicItemEventService_DESELECT string = "deselect"
var MyGraphicItemEventService_SELECT string = "select"
var MyGraphicItemEventService_DUPLICATE string = "duplicate"
var MyGraphicItemEventService_ROTATE string = "rotate"


func NewMyGraphicItemEventService() *MyGraphicItemEventService {
    this := &MyGraphicItemEventService{}
    return this
}


// synchronized
func MyGraphicItemEventService_AddListener(listener MyGraphicItemEventListener) {
        MyGraphicItemEventService_ListenerList.Add(any(reflect.TypeOf(listener)).(reflect.Type), any(listener).(*EventListener))
}


// synchronized
func MyGraphicItemEventService_RemoveListener(listener MyGraphicItemEventListener) {
        MyGraphicItemEventService_ListenerList.Remove(any(reflect.TypeOf(listener)).(reflect.Type), any(listener).(*EventListener))
}


// throws Exception
// synchronized
func MyGraphicItemEventService_Fire(evt *MyGraphicItemEvent) {
        var listeners []any = MyGraphicItemEventService_ListenerList.GetListenerList(); _ = listeners
        for i := int(0); i < len(listeners); i += 2 {
            if listeners[i] == reflect.TypeOf(listeners[i + 1]) {
                if strings.Compare(evt.GetCommand(), MyGraphicItemEventService_SELECT) == 0 {
                    (any(listeners[i + 1]).(MyGraphicItemEventListener)).Highlight(evt)
                } else if strings.Compare(evt.GetCommand(), MyGraphicItemEventService_DESELECT) == 0 {
                    (any(listeners[i + 1]).(MyGraphicItemEventListener)).Deselect(evt)
                } else if strings.Compare(evt.GetCommand(), MyGraphicItemEventService_DELETE) == 0 {
                    (any(listeners[i + 1]).(MyGraphicItemEventListener)).Delete(evt)
                } else if strings.Compare(evt.GetCommand(), MyGraphicItemEventService_DUPLICATE) == 0 {
                    (any(listeners[i + 1]).(MyGraphicItemEventListener)).Duplicate(evt)
                } else if strings.Compare(evt.GetCommand(), MyGraphicItemEventService_ROTATE) == 0 {
                    (any(listeners[i + 1]).(MyGraphicItemEventListener)).Rotate(evt)
                }
            }
        }
}

