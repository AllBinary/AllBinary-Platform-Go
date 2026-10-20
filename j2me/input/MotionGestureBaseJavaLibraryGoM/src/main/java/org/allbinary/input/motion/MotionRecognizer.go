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

type MotionRecognizer struct {
}

// The Java class MotionRecognizer has subclasses so its variables are this interface
type IMotionRecognizer interface {
    MotionRecognizer_() *MotionRecognizer
    ProcessStartMotionEvent(x int, y int, deviceId int, modifiers int)
    ProcessEndMotionEvent(x int, y int, deviceId int, modifiers int)
    ProcessDraggedMotionEvent(x int, y int, deviceId int, modifiers int)
}

func (this *MotionRecognizer) MotionRecognizer_() *MotionRecognizer {
    return this
}


func NewMotionRecognizer() *MotionRecognizer {
    this := &MotionRecognizer{}
    return this
}


// throws Exception
func (this *MotionRecognizer) ProcessStartMotionEvent(x int, y int, deviceId int, modifiers int) {
}


// throws Exception
func (this *MotionRecognizer) ProcessEndMotionEvent(x int, y int, deviceId int, modifiers int) {
}


// throws Exception
func (this *MotionRecognizer) ProcessDraggedMotionEvent(x int, y int, deviceId int, modifiers int) {
}

