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

type PrintLicenseInitInfo struct {
}

var PrintLicenseInitInfo_PATH string = "G:\\mnt\\bc\\mydev\\work\\allbinary_src\\licensing\\data\\"


func NewPrintLicenseInitInfo(options []string) *PrintLicenseInitInfo {
    this := &PrintLicenseInitInfo{}
        LicenseInitInfoUtil_GetInstance().SetFilePath(PrintLicenseInitInfo_PATH + PATH_GLOBALS.GetInstance().INIT_PATH)
    return this
}


// throws Exception
func (this *PrintLicenseInitInfo) GetLicenseInitInfo() *LicenseInitInfo {
        var licenseInitInfoUtil *LicenseInitInfoUtil = LicenseInitInfoUtil_GetInstance(); _ = licenseInitInfoUtil
        if FileFactory_GetInstance().IsFile(licenseInitInfoUtil.GetFilePath()) {
            return licenseInitInfoUtil.Read()
        } else {
            fmt.Println("No Licence File at: " + fmt.Sprint(licenseInitInfoUtil.GetFilePath()))
            return NewLicenseInitInfo()
        }
}


func (this *PrintLicenseInitInfo) PrintInitInfo() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Print("Error: " + fmt.Sprint(e))
            fmt.Println(e)
                }
            }()
            var licenseInitInfo *LicenseInitInfo = this.GetLicenseInitInfo(); _ = licenseInitInfo
            fmt.Println("License ID: " + fmt.Sprint(licenseInitInfo.GetLicenseId()))
            var list IBasicArrayList = licenseInitInfo.GetServerList(); _ = list
            var size int = list.Size(); _ = size
            for index := int(0); index < size; index++ {
                fmt.Println("Lic Serv: " + fmt.Sprint(list.Get(index)))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func PrintLicenseInitInfo_Main(args []string) {
        var initInfo *PrintLicenseInitInfo = NewPrintLicenseInitInfo(args); _ = initInfo
        initInfo.PrintInitInfo()
}

