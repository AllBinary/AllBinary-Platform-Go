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

type AllBinarySensorManager struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ORIENTATION_SENSOR_INPUT string
    ORIENTATION_TYPE string
    List IBasicArrayList
}

var AllBinarySensorManager_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *AllBinarySensorManager) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ORIENTATION_SENSOR_INPUT = "Orientation Sensor Input"
    this.ORIENTATION_TYPE = "Orientation Type"
    this.List = NewBasicArrayListD()
}


func AllBinarySensorManager_GetInstance() *AllBinarySensorManager {
        if AllBinarySensorManager_Instance == NullUtil_GetInstance().NULL_OBJECT {
            AllBinarySensorManager_Instance = NewAllBinarySensorManager()
        }
        return any(AllBinarySensorManager_Instance).(*AllBinarySensorManager)
}


func NewAllBinarySensorManager() *AllBinarySensorManager {
    this := &AllBinarySensorManager{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            GyroSensorFactory_Init()
            AccelerometerSensorFactory_Init()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *AllBinarySensorManager) Init() {
}


func (this *AllBinarySensorManager) GetSensorNamesList() IBasicArrayList {
        return this.List
}


func (this *AllBinarySensorManager) Shutdown() {
}

