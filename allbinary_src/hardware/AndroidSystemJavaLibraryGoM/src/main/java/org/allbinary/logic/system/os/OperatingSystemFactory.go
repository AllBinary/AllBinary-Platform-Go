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

type OperatingSystemFactory struct {
    LogUtil *LogUtil
    GenericOperatingSystem *GenericOperatingSystem
    HasDetected bool
}

var OperatingSystemFactory_Instance *OperatingSystemFactory = NewOperatingSystemFactory()

func (this *OperatingSystemFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GenericOperatingSystem = NoOperatingSystem_NO_OPERATING_SYSTEM
    this.HasDetected = false
}


func OperatingSystemFactory_GetInstance() *OperatingSystemFactory {
        return OperatingSystemFactory_Instance
}


func NewOperatingSystemFactory() *OperatingSystemFactory {
    this := &OperatingSystemFactory{}
    this.initFields()
    return this
}


// synchronized
func (this *OperatingSystemFactory) GetOperatingSystemInstance() *GenericOperatingSystem {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *GenericOperatingSystem) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.GenericOperatingSystem = NoOperatingSystem_NO_OPERATING_SYSTEM
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
                }
            }()
            var systemProperties *SystemProperties = SystemProperties_GetInstance(); _ = systemProperties
            var osName string = systemProperties.GetName(); _ = osName
            if !this.HasDetected {
                this.HasDetected = true
                if strings.Index(osName, string(OperatingSystems_GetInstance().ANDROID)) >= 0 {
                    this.LogUtil.PutF("Found a Android OS", this, commonStrings.GET_INSTANCE)
                    this.GenericOperatingSystem = AndroidOperatingSystemFactory_GetInstance().GetOperatingSystemInstance()
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Operating System Info: ").Append(fmt.Sprint(this.GenericOperatingSystem))), this, commonStrings.GET_INSTANCE)
                } else if strings.Index(osName, string(OperatingSystems_GetInstance().HARMONY)) >= 0 {
                    this.LogUtil.PutF("Found a Harmony OS", this, commonStrings.GET_INSTANCE)
                    this.GenericOperatingSystem = AndroidOperatingSystemFactory_GetInstance().GetOperatingSystemInstance()
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Operating System Info: ").Append(fmt.Sprint(this.GenericOperatingSystem))), this, commonStrings.GET_INSTANCE)
                } else {
                    panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("OS Not Supported: ").Append(osName)))))
                }
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.GenericOperatingSystem
}

