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

type AndroidOperatingSystemFactory struct {
    LogUtil *LogUtil
    TsUtil *TsUtil
}

var AndroidOperatingSystemFactory_Instance *AndroidOperatingSystemFactory = NewAndroidOperatingSystemFactory()

func (this *AndroidOperatingSystemFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
}


func AndroidOperatingSystemFactory_GetInstance() *AndroidOperatingSystemFactory {
        return AndroidOperatingSystemFactory_Instance
}


func NewAndroidOperatingSystemFactory() *AndroidOperatingSystemFactory {
    this := &AndroidOperatingSystemFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *AndroidOperatingSystemFactory) GetOperatingSystemInstance() IGenericOperatingSystem {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IGenericOperatingSystem) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put("Failed to get OperatingSystem returning NoOperatingSystem", this, commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, NoOperatingSystem_NO_OPERATING_SYSTEM; return
                }
            }()
            var GenericOperatingSystem IGenericOperatingSystem; _ = GenericOperatingSystem
            var systemProperties *SystemProperties = SystemProperties_GetInstance(); _ = systemProperties
            var osName string = systemProperties.GetName(); _ = osName
            if this.TsUtil.CompareTo(osName, OperatingSystems_GetInstance().ANDROID) == 0 {
                GenericOperatingSystem = NewAndroidOS()
            } else {
                if OperatingSystems_GetInstance().IsUnknownSpecificOSAllowed() {
                    GenericOperatingSystem = NewAndroidOS()
                } else {
                    panic(errors.New(fmt.Sprint("Specific Android OS Not Supported: " + osName)))
                }
            }
            tryControl1, tryResult1 = 1, GenericOperatingSystem; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

