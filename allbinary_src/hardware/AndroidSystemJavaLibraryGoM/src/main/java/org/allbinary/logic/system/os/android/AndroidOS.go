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

type AndroidOS struct {
    GenericOperatingSystem
    StringBuffer *StringMaker
    DEVICE_ID string
    DEVICE_SOFTWARE_VERSION string
    LINE_1_NUMBER string
    OUYA string
}


func (this *AndroidOS) initFields() {
    this.StringBuffer = NewStringMaker()
    this.DEVICE_ID = "DeviceId"
    this.DEVICE_SOFTWARE_VERSION = "DeviceSoftwareVersion"
    this.LINE_1_NUMBER = "Line1Number"
    this.OUYA = "ouya"
}


func NewAndroidOS() *AndroidOS {
    this := &AndroidOS{}
    this.GenericOperatingSystem = *NewGenericOperatingSystem()
    this.initFields()
        var SDK_VERSION int = AndroidInfoFactory_GetInstance().GetVersion(); _ = SDK_VERSION
        if SDK_VERSION > 10 {
            this.Scalable = true
        }
        var properties *AndroidSystemProperties = AndroidSystemProperties_GetInstance(); _ = properties
        var commonStrings *CommonSeps = CommonSeps_GetInstance(); _ = commonStrings
        this.StringBuffer.Append(this.DEVICE_ID)
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetDeviceId())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append(this.DEVICE_SOFTWARE_VERSION)
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetDeviceSoftwareVersion())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append(this.LINE_1_NUMBER)
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetLine1Number())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("NetworkCountryIso")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetNetworkCountryIso())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("NetworkOperator")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetNetworkOperator())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("NetworkOperatorName")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetNetworkOperatorName())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("NetworkType")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Appendint(properties.GetNetworkType())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("PhoneType")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Appendint(properties.GetPhoneType())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("SimCountryIso")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetSimCountryIso())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("SimOperator")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetSimOperator())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("SimOperatorName")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetSimOperatorName())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("SimSerialNumber")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetSimSerialNumber())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("SubscriberId")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetSubscriberId())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("VoiceMailAlphaTag")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetVoiceMailAlphaTag())
        this.StringBuffer.Append(commonStrings.SPACE)
        this.StringBuffer.Append("VoiceMailNumber")
        this.StringBuffer.Append(commonStrings.EQUALS)
        this.StringBuffer.Append(properties.GetVoiceMailNumber())
        this.StringBuffer.Append(commonStrings.SPACE)
    return this
}


func (this *AndroidOS) IsOverScan() bool {
        var properties *AndroidSystemProperties = AndroidSystemProperties_GetInstance(); _ = properties
        if strings.Index(strings.ToLower(properties.GetDevice()), string(this.OUYA)) >= 0 {
            return true
        }
        return false
}


func (this *AndroidOS) GetOverScanXPercent() int {
        return 90
}


func (this *AndroidOS) GetOverScanYPercent() int {
        return 90
}


func (this *AndroidOS) ToString() string {
        var osBuffer *StringMaker = NewStringMaker(); _ = osBuffer
        osBuffer.Append(this.GenericOperatingSystem.ToString())
        osBuffer.Append(CommonSeps_GetInstance().SPACE)
        osBuffer.Append("Other System Info: \n")
        osBuffer.Append(fmt.Sprint(this.StringBuffer))
        return fmt.Sprint(osBuffer)
}

func (this *AndroidOS) String() string {
    return this.ToString()
}

