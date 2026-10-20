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

type AbeLicenseInterfaceFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Time int64
    Check bool
    CheckPeriod int
    AbeLicenseInterface AbeLicenseInterface
}

var AbeLicenseInterfaceFactory_SINGLETON *AbeLicenseInterfaceFactory = NewAbeLicenseInterfaceFactory()

func (this *AbeLicenseInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Time = 0
    this.Check = false
    this.CheckPeriod = 36000000
    this.AbeLicenseInterface = nil
}


func AbeLicenseInterfaceFactory_GetInstance() *AbeLicenseInterfaceFactory {
        return AbeLicenseInterfaceFactory_SINGLETON
}


func NewAbeLicenseInterfaceFactory() *AbeLicenseInterfaceFactory {
    this := &AbeLicenseInterfaceFactory{}
    this.initFields()
    return this
}


// throws LicensingException
func (this *AbeLicenseInterfaceFactory) GetLicenseInstance(abeClientInformation AbeClientInformationInterface) AbeLicenseInterface {
        if this.IsTimeToGetKey() {
            return this.Get(abeClientInformation)
        } else {
            return this.AbeLicenseInterface
        }
}


// throws LicensingException
func (this *AbeLicenseInterfaceFactory) Get(abeClientInformation AbeClientInformationInterface) AbeLicenseInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 AbeLicenseInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Licensing IO Error", this, this.CommonStrings.GET, e)
            panic(errors.New(fmt.Sprint("License Server Connection Error")))
                }
            }()
            this.LogUtil.PutF("Getting Keys", this, this.CommonStrings.GET)
            this.AbeLicenseInterface = AbeNoLicense_GetInstance()
            var licenseClient *AbeLicenseClient = NewAbeLicenseClient(); _ = licenseClient
            this.AbeLicenseInterface = licenseClient.Get(abeClientInformation)
            if this.AbeLicenseInterface != nil {
                this.LogUtil.PutF("Default Key: " + fmt.Sprint(this.AbeLicenseInterface.GetKey(AbeClientInformationData_GetInstance().KEY)), this, this.CommonStrings.GET)
            }
            tryControl1, tryResult1 = 1, this.AbeLicenseInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbeLicenseInterfaceFactory) IsTimeToGetKey() bool {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var currentTime int64 = calendar.GetTimeInMillis(); _ = currentTime
        if this.AbeLicenseInterface == nil || this.AbeLicenseInterface == AbeNoLicense_GetInstance() || !this.AbeLicenseInterface.HasKey() || this.IsCheck() || currentTime - int64(this.CheckPeriod) > this.Time {
            this.AbeLicenseInterface = nil
            this.Time = currentTime
            return true
        } else {
            return false
        }
}


func (this *AbeLicenseInterfaceFactory) SetCheck(check bool) {
        this.Check = check
}


func (this *AbeLicenseInterfaceFactory) IsCheck() bool {
        return this.Check
}

