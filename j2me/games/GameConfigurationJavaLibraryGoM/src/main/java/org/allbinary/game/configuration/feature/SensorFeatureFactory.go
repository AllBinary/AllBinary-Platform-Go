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

type SensorFeatureFactory struct {
    NO_ORIENTATION *SensorFeature
    ORIENTATION_SENSORS *SensorFeature
    SIMULATED_ORIENTATION_SENSORS *SensorFeature
    YAW *SensorFeature
    YAW_MINUS_RIGHT_PLUS_LEFT *SensorFeature
    YAW_MINUS_LEFT_PLUS_RIGHT *SensorFeature
    PITCH *SensorFeature
    ROLL *SensorFeature
}

var SensorFeatureFactory_Instance *SensorFeatureFactory = NewSensorFeatureFactory()

func (this *SensorFeatureFactory) initFields() {
    this.NO_ORIENTATION = NewSensorFeature("No Orientation")
    this.ORIENTATION_SENSORS = NewSensorFeature("Orientation Sensors")
    this.SIMULATED_ORIENTATION_SENSORS = NewSensorFeature("Simulated Orientation Sensors")
    this.YAW = NewSensorFeature("Yaw")
    this.YAW_MINUS_RIGHT_PLUS_LEFT = NewSensorFeature("Yaw -Right +Left")
    this.YAW_MINUS_LEFT_PLUS_RIGHT = NewSensorFeature("Yaw -Left +Right ")
    this.PITCH = NewSensorFeature("Pitch")
    this.ROLL = NewSensorFeature("Roll")
}


func SensorFeatureFactory_GetInstance() *SensorFeatureFactory {
        return SensorFeatureFactory_Instance
}


func NewSensorFeatureFactory() *SensorFeatureFactory {
    this := &SensorFeatureFactory{}
    this.initFields()
    return this
}

