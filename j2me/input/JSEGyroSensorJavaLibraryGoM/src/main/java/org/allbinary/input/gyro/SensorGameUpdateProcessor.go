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

type SensorGameUpdateProcessor struct {
    InputSensor IAllBinarySensor
}

// The Java class SensorGameUpdateProcessor has subclasses so its variables are this interface
type ISensorGameUpdateProcessor interface {
    SensorGameUpdateProcessor_() *SensorGameUpdateProcessor
    SetInputSensor(inputSensor IAllBinarySensor)
    GetInputSensor() IAllBinarySensor
    Process(anyType any)
    SendNotifications(anyType any)
    IsAnySensor() bool
    SetNoSensors()
}

func (this *SensorGameUpdateProcessor) SensorGameUpdateProcessor_() *SensorGameUpdateProcessor {
    return this
}


func (this *SensorGameUpdateProcessor) initFields() {
    
}

func NewSensorGameUpdateProcessor() *SensorGameUpdateProcessor {
    this := &SensorGameUpdateProcessor{}
    this.initFields()
    return this
}


func (this *SensorGameUpdateProcessor) SetInputSensor(inputSensor IAllBinarySensor) {
        this.InputSensor = inputSensor
}


func (this *SensorGameUpdateProcessor) GetInputSensor() IAllBinarySensor {
        return this.InputSensor
}


// throws Exception
func (this *SensorGameUpdateProcessor) Process(anyType any) {
}


// throws Exception
func (this *SensorGameUpdateProcessor) SendNotifications(anyType any) {
}


func (this *SensorGameUpdateProcessor) IsAnySensor() bool {
        if this.InputSensor == NoAllBinarySensor_GetInstance() {
            return false
        } else {
            return true
        }
}


func (this *SensorGameUpdateProcessor) SetNoSensors() {
        GyroSensorFactory_GetInstance().SetCompleteMotionGestureInputEventListener(NoCompleteMotionGestureInputEventListener_GetInstance())
        AccelerometerSensorFactory_GetInstance().SetCompleteMotionGestureInputEventListener(NoCompleteMotionGestureInputEventListener_GetInstance())
        this.SetInputSensor(NoAllBinarySensor_GetInstance())
}

