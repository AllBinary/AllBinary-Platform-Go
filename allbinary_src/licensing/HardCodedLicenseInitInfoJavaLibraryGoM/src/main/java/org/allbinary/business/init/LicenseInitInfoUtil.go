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

type LicenseInitInfoUtil struct {
    INITFILENAME string
    ABOUT string
    PRIVACY_POLICY string
    OFFICIAL_ALLBINARY_SERVICES_SERVER string
}

var LicenseInitInfoUtil_Instance *LicenseInitInfoUtil = NewLicenseInitInfoUtil()

func (this *LicenseInitInfoUtil) initFields() {
    this.INITFILENAME = "licenseinitdata.dat"
    this.ABOUT = "about"
    this.PRIVACY_POLICY = "privacy_policy"
    this.OFFICIAL_ALLBINARY_SERVICES_SERVER = "https://services.allbinary.com/LicServ/serverssl.php"
}


func LicenseInitInfoUtil_GetInstance() *LicenseInitInfoUtil {
        return LicenseInitInfoUtil_Instance
}


// synchronized
func (this *LicenseInitInfoUtil) SetFilePath(filePath string) {
}


func NewLicenseInitInfoUtil() *LicenseInitInfoUtil {
    this := &LicenseInitInfoUtil{}
    this.initFields()
    return this
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) Write(initData *LicenseInitInfo) {
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) Read() *LicenseInitInfo {
        return this.ReadAgain(0)
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) ReadAgain(initializeCounter int) *LicenseInitInfo {
        var licenseInitInfo *LicenseInitInfo = NewLicenseInitInfo(); _ = licenseInitInfo
        licenseInitInfo.ClearServers()
        licenseInitInfo.AddServer(this.OFFICIAL_ALLBINARY_SERVICES_SERVER)
        return licenseInitInfo
}

