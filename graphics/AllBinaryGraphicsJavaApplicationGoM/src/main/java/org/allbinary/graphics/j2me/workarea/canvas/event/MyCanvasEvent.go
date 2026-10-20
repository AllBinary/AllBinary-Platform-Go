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

type MyCanvasEvent struct {
    EventObject
}



func NewMyCanvasEvent(source any) *MyCanvasEvent {
    this := &MyCanvasEvent{}
        this.EventObject = *NewEventObject(source)
    return this
}


func (this *MyCanvasEvent) GetSource() any {
        return this.EventObject.GetSource()
}


func (this *MyCanvasEvent) ToString() string {
        return this.EventObject.ToString()
}

func (this *MyCanvasEvent) String() string {
    return this.ToString()
}


func (this *MyCanvasEvent) GetCommand() string {
        var myEventSource *MyCanvasEventSource = any(this.GetSource()).(*MyCanvasEventSource); _ = myEventSource
        return myEventSource.GetCommand()
}


func (this *MyCanvasEvent) GetTreeNode() *MutableTreeNode {
        var myEventSource *MyCanvasEventSource = any(this.GetSource()).(*MyCanvasEventSource); _ = myEventSource
        return myEventSource.GetTreeNode()
}


func (this *MyCanvasEvent) GetAngle() float64 {
        var myEventSource *MyCanvasEventSource = any(this.GetSource()).(*MyCanvasEventSource); _ = myEventSource
        return myEventSource.GetAngle()
}

