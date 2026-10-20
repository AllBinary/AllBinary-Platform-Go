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

type MyCanvasEventService struct {
}

var MyCanvasEventService_ListenerList *EventListenerList = NewEventListenerList()
var MyCanvasEventService_DELETE string = "delete"
var MyCanvasEventService_DUPLICATE string = "duplicate"
var MyCanvasEventService_SELECT string = "select"
var MyCanvasEventService_ROTATE string = "rotate"
var MyCanvasEventService_EXPLODE string = "explode"
var MyCanvasEventService_AUTOEXPLODE string = "autoExplode"
var MyCanvasEventService_CENTER string = "center"


func NewMyCanvasEventService() *MyCanvasEventService {
    this := &MyCanvasEventService{}
    return this
}


// synchronized
func MyCanvasEventService_AddListener(listener MyCanvasEventListener) {
        MyCanvasEventService_ListenerList.Add(any(reflect.TypeOf(listener)).(reflect.Type), any(listener).(*EventListener))
}


// synchronized
func MyCanvasEventService_RemoveListener(listener MyCanvasEventListener) {
        MyCanvasEventService_ListenerList.Remove(any(reflect.TypeOf(listener)).(reflect.Type), any(listener).(*EventListener))
}


// throws Exception
// synchronized
func MyCanvasEventService_Fire(evt *MyCanvasEvent) {
        var listeners []any = MyCanvasEventService_ListenerList.GetListenerList(); _ = listeners
        for i := int(0); i < len(listeners); i += 2 {
            if listeners[i] == reflect.TypeOf(listeners[i + 1]) {
                if strings.Compare(evt.GetCommand(), MyCanvasEventService_DELETE) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Delete(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_DUPLICATE) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Duplicate(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_SELECT) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Select(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_ROTATE) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Rotate(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_EXPLODE) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Explode(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_AUTOEXPLODE) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).AutoExplode(evt)
                } else if strings.Compare(evt.GetCommand(), MyCanvasEventService_CENTER) == 0 {
                    (any(listeners[i + 1]).(MyCanvasEventListener)).Center(evt)
                }
            }
        }
}

