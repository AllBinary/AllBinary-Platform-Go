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

type ScreenCapture struct {
}

// The Java class ScreenCapture has subclasses so its variables are this interface
type IScreenCapture interface {
    ScreenCapture_() *ScreenCapture
    SaveFrame()
    StartRecording()
    EndRecording()
}

func (this *ScreenCapture) ScreenCapture_() *ScreenCapture {
    return this
}


func NewScreenCapture() *ScreenCapture {
    this := &ScreenCapture{}
    return this
}


func (this *ScreenCapture) SaveFrame() {
}


func (this *ScreenCapture) StartRecording() {
}


func (this *ScreenCapture) EndRecording() {
}

