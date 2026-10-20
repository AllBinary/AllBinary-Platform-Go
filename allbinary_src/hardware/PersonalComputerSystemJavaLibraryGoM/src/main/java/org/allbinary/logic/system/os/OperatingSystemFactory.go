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
    CommonStrings *CommonStrings
    GenericOperatingSystem IGenericOperatingSystem
    HasDetected bool
}

var OperatingSystemFactory_Instance *OperatingSystemFactory = NewOperatingSystemFactory()

func (this *OperatingSystemFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
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
func (this *OperatingSystemFactory) GetOperatingSystemInstance() IGenericOperatingSystem {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IGenericOperatingSystem) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, NoOperatingSystem_NO_OPERATING_SYSTEM; return
                }
            }()
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            var systemProperties *SystemProperties = SystemProperties_GetInstance(); _ = systemProperties
            var osName string = systemProperties.GetName(); _ = osName
            var osArch string = systemProperties.GetArch(); _ = osArch
            var osVersion string = systemProperties.GetVersion(); _ = osVersion
            if !this.HasDetected {
                this.LogUtil.PutF("osName: " + osName, this, this.CommonStrings.GET_INSTANCE)
                this.HasDetected = true
                if strings.Index(osName, string(operatingSystems.LINUX)) >= 0 {
                    if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                        this.LogUtil.PutF("Found a Linux OS", this, this.CommonStrings.GET_INSTANCE)
                    }
                    this.GenericOperatingSystem = LinuxOperatingSystemFactory_GetInstance().GetOperatingSystemInstance()
                } else if strings.Index(osName, string(operatingSystems.WINDOWS)) >= 0 {
                    if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                        this.LogUtil.PutF("Found a Windows OS", this, this.CommonStrings.GET_INSTANCE)
                    }
                    this.GenericOperatingSystem = WindowsOperatingSystemFactory_GetInstance().GetOperatingSystemInstance()
                } else if strings.Index(osName, string(operatingSystems.SOLARIS)) >= 0 {
                    if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                        this.LogUtil.PutF("Found a Solaris OS", this, this.CommonStrings.GET_INSTANCE)
                    }
                    this.GenericOperatingSystem = SolarisOperatingSystemFactory_GetInstance().GetOperatingSystemInstance()
                } else {
                    panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("OS Not Supported: ").Append(osName)))))
                }
                var log *Log = LogFactory_GetInstanceF(fmt.Sprint(NewStringMaker().Append("OperatingSystem Info: ").Append(StringUtil_GetInstance().ToString(this.GenericOperatingSystem))), this, this.CommonStrings.GET_INSTANCE); _ = log
                fmt.Println(fmt.Sprint(log))
                this.LogUtil.PutL(log)
            }
            tryControl1, tryResult1 = 1, this.GenericOperatingSystem; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}
