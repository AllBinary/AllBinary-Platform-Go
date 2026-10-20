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

type OrientationToGameKeyEvent struct {
    AllBinaryGyroSensor IAllBinaryOrientationSensor
    MIN int
}


func (this *OrientationToGameKeyEvent) initFields() {
    
    this.MIN = 15
}


func NewOrientationToGameKeyEvent() *OrientationToGameKeyEvent {
    this := &OrientationToGameKeyEvent{}
    this.initFields()
        if Features_GetInstance().IsFeature(SensorFeatureFactory_GetInstance().NO_ORIENTATION) {
            this.AllBinaryGyroSensor = NewNoGyroSensor()
        } else {
            panic(errors.New(fmt.Sprint("Not Such Feature")))
        }
    return this
}


func (this *OrientationToGameKeyEvent) SetAllBinaryGyroSensor(allBinaryGyroSensor IAllBinaryOrientationSensor) {
        this.AllBinaryGyroSensor = allBinaryGyroSensor
}


func (this *OrientationToGameKeyEvent) GetAllBinaryGyroSensor() IAllBinaryOrientationSensor {
        return this.AllBinaryGyroSensor
}


// throws Exception
func (this *OrientationToGameKeyEvent) Process() {
}


// throws Exception
func (this *OrientationToGameKeyEvent) UpdateGameKeyEvents(orientation *Orientation, yaw int, pitch int, roll int) {
}

