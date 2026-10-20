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

type MyGraphicItemEvent struct {
    EventObject
}



func NewMyGraphicItemEvent(source any) *MyGraphicItemEvent {
    this := &MyGraphicItemEvent{}
        this.EventObject = *NewEventObject(source)
    return this
}


func (this *MyGraphicItemEvent) GetSource() any {
        return this.EventObject.GetSource()
}


func (this *MyGraphicItemEvent) ToString() string {
        return this.EventObject.ToString()
}

func (this *MyGraphicItemEvent) String() string {
    return this.ToString()
}


func (this *MyGraphicItemEvent) GetCommand() string {
        var myEventSource *MyGraphicItemEventSource = any(this.GetSource()).(*MyGraphicItemEventSource); _ = myEventSource
        return myEventSource.GetCommand()
}


func (this *MyGraphicItemEvent) GetTreeNode() *MutableTreeNode {
        var myEventSource *MyGraphicItemEventSource = any(this.GetSource()).(*MyGraphicItemEventSource); _ = myEventSource
        return myEventSource.GetTreeNode()
}


func (this *MyGraphicItemEvent) GetAngle() float64 {
        var myEventSource *MyGraphicItemEventSource = any(this.GetSource()).(*MyGraphicItemEventSource); _ = myEventSource
        return myEventSource.GetAngle()
}

