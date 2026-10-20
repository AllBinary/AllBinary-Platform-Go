/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2009 AllBinary
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

type BaseGameInputMotionEventProcessor struct {
    MotionEventProcessor
    CommonStrings *CommonStrings
    GameMotionGestureListener *GameMotionGestureListener
}

// The Java class BaseGameInputMotionEventProcessor has subclasses so its variables are this interface
type IBaseGameInputMotionEventProcessor interface {
    IMotionEventProcessor
    BaseGameInputMotionEventProcessor_() *BaseGameInputMotionEventProcessor
    Process(motionEvent *MotionEvent)
    AddListener()
}

func (this *BaseGameInputMotionEventProcessor) BaseGameInputMotionEventProcessor_() *BaseGameInputMotionEventProcessor {
    return this
}


func (this *BaseGameInputMotionEventProcessor) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewBaseGameInputMotionEventProcessor() *BaseGameInputMotionEventProcessor {
    this := &BaseGameInputMotionEventProcessor{}
    this.MotionEventProcessor = *NewMotionEventProcessor()
    this.initFields()
        this.GameMotionGestureListener = NewGameMotionGestureListener(MotionGestureReceiveInterfaceFactory_GetInstance())
        this.AddListener()
    return this
}


func (this *BaseGameInputMotionEventProcessor) Process(motionEvent *MotionEvent) {
}


func (this *BaseGameInputMotionEventProcessor) AddListener() {
        var motionGesturesHandler *BasicMotionGesturesHandler = BasicMotionGesturesHandler_GetInstance(); _ = motionGesturesHandler
        motionGesturesHandler.AddListenerInterface(this.GameMotionGestureListener)
}

