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

type SetLicenseInitInfo struct {
    Vector IBasicArrayList
}

var SetLicenseInitInfo_Stdinput *BufferedReader = NewBufferedReader(NewInputStreamReader(System_In))

func (this *SetLicenseInitInfo) initFields() {
    this.Vector = NewBasicArrayListD()
}


func NewSetLicenseInitInfo(options []string) *SetLicenseInitInfo {
    this := &SetLicenseInitInfo{}
    this.initFields()
        LicenseInitInfoUtil_GetInstance().SetFilePath(PrintLicenseInitInfo_PATH + PATH_GLOBALS.GetInstance().INIT_PATH)
        for index := int(1); index < len(options); index++ {
            this.Vector.Add(options[index])
        }
    return this
}


// throws Exception
func (this *SetLicenseInitInfo) GetLicenseInitInfo() *LicenseInitInfo {
        var licenseInitInfoUtil *LicenseInitInfoUtil = LicenseInitInfoUtil_GetInstance(); _ = licenseInitInfoUtil
        if FileFactory_GetInstance().IsFile(licenseInitInfoUtil.INITFILENAME) {
            return licenseInitInfoUtil.Read()
        } else {
            return NewLicenseInitInfo()
        }
}


func (this *SetLicenseInitInfo) CreateNewLicenseInitInfo(licenceServerVector IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Print("Error: " + fmt.Sprint(e))
            fmt.Println(e)
                }
            }()
            if licenceServerVector == nil {
                panic(errors.New(fmt.Sprint("null BasicArrayList Passed - Just Make Empty if you want")))
            }
            var licenseInitInfo *LicenseInitInfo = this.GetLicenseInitInfo(); _ = licenseInitInfo
            fmt.Println("License Serv Example: http://www.dartstatistics.com/LicServ/serverssl.php")
            licenseInitInfo.ClearServers()
            var file *File = NewFile(fmt.Sprint(PrintLicenseInitInfo_PATH) + "licenseservers.txt"); _ = file
            var reader *FileReader = NewFileReader(file); _ = reader
            var lineNumberReader *LineNumberReader = NewLineNumberReader(reader); _ = lineNumberReader
            for true {
                var line string = lineNumberReader.ReadLine(); _ = line
                if line != "" {
                    licenseInitInfo.AddServer(line)
                } else {
                    break
                }
            }
            LicenseInitInfoUtil_GetInstance().Write(licenseInitInfo)
            LicenseInitInfoUtil_GetInstance().Read()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func SetLicenseInitInfo_Main(args []string) {
        var initInfo *SetLicenseInitInfo = NewSetLicenseInitInfo(args); _ = initInfo
        initInfo.CreateNewLicenseInitInfo(NewBasicArrayListD())
}

