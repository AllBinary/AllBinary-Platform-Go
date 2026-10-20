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

type LayerBounds struct {
    LogUtil *LogUtil
    Rectangle *Rectangle
}

// Implements BoundsChangeEventListener

// The Java class LayerBounds has subclasses so its variables are this interface
type ILayerBounds interface {
    LayerBounds_() *LayerBounds
    GetRectangleP() *Rectangle
    Visit(boundsVisitorInterface BoundsVisitorInterface)
    OnEvent(eventObject IAllBinaryEventObject)
    OnBoundsChangeEvent(boundsChangeEvent *BoundsChangeEvent)
}

func (this *LayerBounds) LayerBounds_() *LayerBounds {
    return this
}


func (this *LayerBounds) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewLayerBounds(rectangle *Rectangle) *LayerBounds {
    this := &LayerBounds{}
    this.initFields()
        this.Rectangle = rectangle
    return this
}


func (this *LayerBounds) GetRectangleP() *Rectangle {
        return this.Rectangle
}


func (this *LayerBounds) Visit(boundsVisitorInterface BoundsVisitorInterface) {
}


func (this *LayerBounds) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *LayerBounds) OnBoundsChangeEvent(boundsChangeEvent *BoundsChangeEvent) {
        this.Rectangle = boundsChangeEvent.GetRectangle()
}

