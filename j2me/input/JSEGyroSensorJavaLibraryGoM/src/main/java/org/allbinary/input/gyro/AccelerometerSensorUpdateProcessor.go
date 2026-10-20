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

type AccelerometerSensorUpdateProcessor struct {
    Processor
}


func NewAccelerometerSensorUpdateProcessor() *AccelerometerSensorUpdateProcessor {
    this := &AccelerometerSensorUpdateProcessor{}
    this.Processor = *NewProcessor()
    return this
}


// throws Exception
func (this *AccelerometerSensorUpdateProcessor) Process() {
        var sensorFeatureFactory *SensorFeatureFactory = SensorFeatureFactory_GetInstance(); _ = sensorFeatureFactory
        if Features_GetInstance().IsFeature(sensorFeatureFactory.ORIENTATION_SENSORS) {
            AccelerometerSensorFactory_GetInstance().SetCompleteMotionGestureInputEventListener(NewSingleKeyPressGameKeyFromMotionGestureEventListener())
        } else if Features_GetInstance().IsFeature(sensorFeatureFactory.SIMULATED_ORIENTATION_SENSORS) {
            AccelerometerSensorFactory_GetInstance().SetCompleteMotionGestureInputEventListener(NewSingleKeyPressGameKeyFromMotionGestureEventListener())
        } else if Features_GetInstance().IsFeature(sensorFeatureFactory.NO_ORIENTATION) {
            AccelerometerSensorFactory_GetInstance().SetCompleteMotionGestureInputEventListener(NoCompleteMotionGestureInputEventListener_GetInstance())
        } else {
            panic(errors.New(fmt.Sprint("Sensor Feature Not Set")))
        }
}

