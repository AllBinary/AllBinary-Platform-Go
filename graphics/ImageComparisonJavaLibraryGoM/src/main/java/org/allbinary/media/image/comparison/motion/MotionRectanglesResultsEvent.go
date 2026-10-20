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

type MotionRectanglesResultsEvent struct {
    AllBinaryEventObject
    Frame int64
    MotionRectangles IMotionRectangles
}


func (this *MotionRectanglesResultsEvent) initFields() {
    
    
}


func NewMotionRectanglesResultsEvent(anyType any, frame int64, motionRectangles IMotionRectangles) *MotionRectanglesResultsEvent {
    this := &MotionRectanglesResultsEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
        this.Frame = frame
        this.SetMotionRectangles(motionRectangles)
    return this
}


func (this *MotionRectanglesResultsEvent) GetMotionRectangles() IMotionRectangles {
        return this.MotionRectangles
}


func (this *MotionRectanglesResultsEvent) SetMotionRectangles(motionRectangles IMotionRectangles) {
        this.MotionRectangles = motionRectangles
}


func (this *MotionRectanglesResultsEvent) GetFrame() int64 {
        return this.Frame
}

