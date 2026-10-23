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

type WindowsOperatingSystemFactory struct {
    LogUtil *LogUtil
}

var WindowsOperatingSystemFactory_Instance *WindowsOperatingSystemFactory = NewWindowsOperatingSystemFactory()

func (this *WindowsOperatingSystemFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func WindowsOperatingSystemFactory_GetInstance() *WindowsOperatingSystemFactory {
        return WindowsOperatingSystemFactory_Instance
}


func NewWindowsOperatingSystemFactory() *WindowsOperatingSystemFactory {
    this := &WindowsOperatingSystemFactory{}
    this.initFields()
    return this
}


func (this *WindowsOperatingSystemFactory) GetOperatingSystemInstance() IGenericOperatingSystem {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IGenericOperatingSystem) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, NoOperatingSystem_NO_OPERATING_SYSTEM; return
                }
            }()
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            var osName string = SystemProperties_GetInstance().GetName(); _ = osName
            var operatingSystemInterface IGenericOperatingSystem; _ = operatingSystemInterface
            if strings.Compare(osName, operatingSystems.WINDOWS_NT) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else if strings.Compare(osName, operatingSystems.WINDOWS_10) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else if strings.Compare(osName, operatingSystems.WINDOWS_11) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else if strings.Compare(osName, operatingSystems.WINDOWS2000) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else if strings.Compare(osName, operatingSystems.WINDOWS_XP) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else if strings.Compare(osName, operatingSystems.WINDOWS_VISTA) == 0 {
                operatingSystemInterface = NewWindowsOS()
            } else {
                if operatingSystems.IsUnknownSpecificOSAllowed() {
                    operatingSystemInterface = NewWindowsOS()
                } else {
                    panic(errors.New(fmt.Sprint("Specific Windows OS Not Supported: " + osName)))
                }
            }
            tryControl1, tryResult1 = 1, operatingSystemInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

