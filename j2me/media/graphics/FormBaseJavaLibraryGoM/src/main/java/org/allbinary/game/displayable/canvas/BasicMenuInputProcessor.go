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

type BasicMenuInputProcessor struct {
    PlayerGameInput
    MotionGestureEventList IBasicArrayList
    Canvas ICanvas
}

// Implements BaseMotionGestureEventListener

// The Java class BasicMenuInputProcessor has subclasses so its variables are this interface
type IBasicMenuInputProcessor interface {
    IPlayerGameInput
    BasicMenuInputProcessor_() *BasicMenuInputProcessor
    OnMotionGestureEvent(motionGestureEvent *MotionGestureEvent)
    OnScrolledMotionGestureEvent(motionGestureEvent *MotionGestureEvent)
    SetCanvas(canvas IMyCanvas)
    GetCanvas() IMyCanvas
    ProcessInputList() int
}

func (this *BasicMenuInputProcessor) BasicMenuInputProcessor_() *BasicMenuInputProcessor {
    return this
}


func (this *BasicMenuInputProcessor) initFields() {
    this.MotionGestureEventList = NewBasicArrayListD()
    this.Canvas = NullCanvas_NULL_CANVAS
}


func NewBasicMenuInputProcessor(gameKeyEventList IBasicArrayList, playerInputId int, gameCanvas ICanvas) *BasicMenuInputProcessor {
    this := &BasicMenuInputProcessor{}
        this.PlayerGameInput = *NewPlayerGameInput(gameKeyEventList, NewBasicArrayListD(), playerInputId)
        this.initFields()
        this.Canvas = gameCanvas
    return this
}


func (this *BasicMenuInputProcessor) OnMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
        this.MotionGestureEventList.Add(motionGestureEvent)
}


func (this *BasicMenuInputProcessor) OnScrolledMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
}


func (this *BasicMenuInputProcessor) SetCanvas(canvas IMyCanvas) {
        this.Canvas = canvas
}


func (this *BasicMenuInputProcessor) GetCanvas() IMyCanvas {
        return any(this.Canvas).(IMyCanvas)
}


// throws Exception
func (this *BasicMenuInputProcessor) ProcessInputList() int {
        return -1
}

