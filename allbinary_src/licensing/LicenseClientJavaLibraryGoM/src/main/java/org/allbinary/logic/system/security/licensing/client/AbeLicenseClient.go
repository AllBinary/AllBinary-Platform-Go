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

type AbeLicenseClient struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    MINSERVERS int
}


func (this *AbeLicenseClient) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.MINSERVERS = 3
}


func NewAbeLicenseClient() *AbeLicenseClient {
    this := &AbeLicenseClient{}
    this.initFields()
    return this
}


// throws Exception
// synchronized
func (this *AbeLicenseClient) Get(abeClientInformation AbeClientInformationInterface) AbeLicenseInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 AbeLicenseInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.GET)
            var xmlRpcAbeLicenseClient IXmlRpcAbeClient = NewXmlRpcAbeLicenseRetrievalClient(abeClientInformation); _ = xmlRpcAbeLicenseClient
            var abeLicenseInterface AbeLicenseInterface = any(xmlRpcAbeLicenseClient.Get(nil)).(AbeLicenseInterface); _ = abeLicenseInterface
            var licenseId string = abeLicenseInterface.GetLicenseId(); _ = licenseId
            var servers IBasicArrayList = abeLicenseInterface.GetServers(); _ = servers
            var isNewLicenseId bool = false; _ = isNewLicenseId
            var isBetterServerList bool = false; _ = isBetterServerList
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if stringValidationUtil.IsEmpty(abeClientInformation.GetLicenseId()) && !abeClientInformation.IsSameId(licenseId) {
                isNewLicenseId = true
            }
            if servers.Size() >= this.MINSERVERS && abeClientInformation.IsLargerOrDifferentServerList(servers) {
                isBetterServerList = true
            }
            if isBetterServerList || isNewLicenseId {
                var initInfo *LicenseInitInfo = LicenseInitInfoUtil_GetInstance().Read(); _ = initInfo
                initInfo.SetLicenseId(licenseId)
                initInfo.SetServerList(servers)
                LicenseInitInfoUtil_GetInstance().Write(initInfo)
                abeClientInformation.Init()
            }
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.GET)
            tryControl1, tryResult1 = 1, abeLicenseInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

