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

type SystemHardwareFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    HardwareInterface HardwareInterface
}

var SystemHardwareFactory_Instance *SystemHardwareFactory = NewSystemHardwareFactory()

func (this *SystemHardwareFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.HardwareInterface = NoHardware_GetInstance()
}


func SystemHardwareFactory_GetInstance() *SystemHardwareFactory {
        return SystemHardwareFactory_Instance
}


func NewSystemHardwareFactory() *SystemHardwareFactory {
    this := &SystemHardwareFactory{}
    this.initFields()
    return this
}


func (this *SystemHardwareFactory) GetInstance(operatingSystemInterface IGenericOperatingSystem) HardwareInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 HardwareInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            if this.HardwareInterface == NoHardware_GetInstance() {
                this.HardwareInterface = HardwareFactory_GetInstance().GetInstance(operatingSystemInterface)
            }
            this.LogUtil.PutF("Found Hardware", this, this.CommonStrings.CONSTRUCTOR)
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.HardwareInterface
}


func SystemHardwareFactory_Static_toString() string {
        var systemHardwareFactory *SystemHardwareFactory = SystemHardwareFactory_GetInstance(); _ = systemHardwareFactory
        var osBuffer *StringMaker = NewStringMaker(); _ = osBuffer
        osBuffer.Append("Hardware Info: \n")
        if systemHardwareFactory.HardwareInterface != nil {
            osBuffer.Append(fmt.Sprint(systemHardwareFactory.HardwareInterface))
        } else {
            osBuffer.Append(StringUtil_GetInstance().NULL_STRING)
        }
        return fmt.Sprint(osBuffer)
}

