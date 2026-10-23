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

// This is the information sent to the license server
type AbeClientInformation struct {
    ClientInformation
    LogUtil *LogUtil
}

// The Java class AbeClientInformation has subclasses so its variables are this interface
type IAbeClientInformation interface {
    IClientInformation
    AbeClientInformation_() *AbeClientInformation
    Init()
}

func (this *AbeClientInformation) AbeClientInformation_() *AbeClientInformation {
    return this
}


func (this *AbeClientInformation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewAbeClientInformation(name string, version string, specialName string, shortName string) *AbeClientInformation {
    this := &AbeClientInformation{}
        this.ClientInformation = *NewClientInformation(name, version, specialName, shortName)
        this.initFields()
    return this
}


func (this *AbeClientInformation) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            this.SetOperatingSystemInterface(OperatingSystemFactory_GetInstance().GetOperatingSystemInstance())
            var licenseInitInfo *LicenseInitInfo = LicenseInitInfoUtil_GetInstance().Read(); _ = licenseInitInfo
            this.SetLicenseId(licenseInitInfo.GetLicenseId())
            this.SetLicenseServers(licenseInitInfo.GetServerList())
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if stringValidationUtil.IsEmpty(this.GetLicenseId()) {
                var NONE string = "No License Id"; _ = NONE
                this.SetLicenseId(NONE)
            }
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Special Name: ").Append(this.GetSpecialName())), this, commonStrings.CONSTRUCTOR)
            return
        }(); tryControl1 == 1 {
            return
        }
}

