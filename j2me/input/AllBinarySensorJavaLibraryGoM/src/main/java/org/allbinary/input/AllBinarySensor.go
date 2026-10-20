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

type AllBinarySensor struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements AllBinarySensorListener

// The Java class AllBinarySensor has subclasses so its variables are this interface
type IAllBinarySensor interface {
    AllBinarySensor_() *AllBinarySensor
    Init()
    Update()
    GetId() int
    Shutdown()
    OnSensorChange()
}

func (this *AllBinarySensor) AllBinarySensor_() *AllBinarySensor {
    return this
}

var AllBinarySensor_NULL_ALLBINARY_SENSOR IAllBinarySensor = NewAllBinarySensor()
var AllBinarySensor_Max int = 100

func (this *AllBinarySensor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewAllBinarySensor() *AllBinarySensor {
    this := &AllBinarySensor{}
    this.initFields()
    return this
}


func (this *AllBinarySensor) Init() {
}


// throws Exception
func (this *AllBinarySensor) Update() {
}


func AllBinarySensor_SetMax(max_ int) {
        AllBinarySensor_Max = max_
}


func AllBinarySensor_GetMax() int {
        return AllBinarySensor_Max
}


func (this *AllBinarySensor) GetId() int {
        return AllBinarySensor_GetMax()
}


func (this *AllBinarySensor) Shutdown() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "shutdown")
}


func (this *AllBinarySensor) OnSensorChange() {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}

