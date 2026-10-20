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

type MotionEventProcessor struct {
}

// The Java class MotionEventProcessor has subclasses so its variables are this interface
type IMotionEventProcessor interface {
    MotionEventProcessor_() *MotionEventProcessor
    Process(motionEvent *MotionEvent)
}

func (this *MotionEventProcessor) MotionEventProcessor_() *MotionEventProcessor {
    return this
}



func NewMotionEventProcessor() *MotionEventProcessor {
    this := &MotionEventProcessor{}
    return this
}


// throws Exception
func (this *MotionEventProcessor) Process(motionEvent *MotionEvent) {
}

