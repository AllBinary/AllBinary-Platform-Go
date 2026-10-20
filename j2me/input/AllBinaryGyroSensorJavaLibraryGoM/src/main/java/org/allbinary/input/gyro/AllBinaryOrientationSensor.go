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

type AllBinaryOrientationSensor struct {
    AllBinarySensor
    SensorDataFloatArray []float32
    Listener CompleteMotionGestureInputEventListenerInterface
}

// The Java class AllBinaryOrientationSensor has subclasses so its variables are this interface
type IAllBinaryOrientationSensor interface {
    IAllBinarySensor
    AllBinaryOrientationSensor_() *AllBinaryOrientationSensor
    SetCompleteMotionGestureInputEventListener(completeMotionGestureInputEventListenerInterface CompleteMotionGestureInputEventListenerInterface)
    GetOrientation() *Orientation
}

func (this *AllBinaryOrientationSensor) AllBinaryOrientationSensor_() *AllBinaryOrientationSensor {
    return this
}


func (this *AllBinaryOrientationSensor) initFields() {
    this.SensorDataFloatArray = PrimitiveFloatUtil_GetArrayInstance()
    this.Listener = NoCompleteMotionGestureInputEventListener_GetInstance()
}


func (this *AllBinaryOrientationSensor) SetCompleteMotionGestureInputEventListener(completeMotionGestureInputEventListenerInterface CompleteMotionGestureInputEventListenerInterface) {
        this.Listener = completeMotionGestureInputEventListenerInterface
}


func NewAllBinaryOrientationSensor() *AllBinaryOrientationSensor {
    this := &AllBinaryOrientationSensor{}
    this.AllBinarySensor = *NewAllBinarySensor()
    this.initFields()
    return this
}


// throws Exception
func (this *AllBinaryOrientationSensor) GetOrientation() *Orientation {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}

