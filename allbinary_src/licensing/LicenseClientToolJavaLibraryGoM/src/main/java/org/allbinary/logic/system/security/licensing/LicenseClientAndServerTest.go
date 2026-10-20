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

type LicenseClientAndServerTest struct {
    LogUtil *LogUtil
    Vector *BasicArrayList
}


func (this *LicenseClientAndServerTest) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Vector = NewBasicArrayListD()
}


func NewLicenseClientAndServerTest(options []string) *LicenseClientAndServerTest {
    this := &LicenseClientAndServerTest{}
    this.initFields()
        LicenseInitInfoUtil_GetInstance().SetFilePath(PrintLicenseInitInfo_PATH + PATH_GLOBALS.GetInstance().INIT_PATH)
        for index := int(1); index < len(options); index++ {
            this.Vector.Add(options[index])
        }
        if len(options) < 2 {
            var file *File = NewFile(fmt.Sprint(PrintLicenseInitInfo_PATH) + "testlicenseservers.txt"); _ = file
            var reader *FileReader = NewFileReader(file); _ = reader
            var lineNumberReader *LineNumberReader = NewLineNumberReader(reader); _ = lineNumberReader
            for true {
                var line string = lineNumberReader.ReadLine(); _ = line
                if line != "" {
                    this.Vector.Add(line)
                } else {
                    break
                }
            }
        }
    return this
}


// synchronized
func (this *LicenseClientAndServerTest) Test() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Test Error: " + fmt.Sprint(stringBuffer), this, "test", e)
                }
            }()
            var list *BasicArrayList = NewBasicArrayListD(); _ = list
            var TESTING string = "Testing with new licence:<br/>"; _ = TESTING
            var PROCESSING string = "Processing: <br/>"; _ = PROCESSING
            var END string = "<br/>End<br/>"; _ = END
            var RESULT string = "Test Result: "; _ = RESULT
            var size int = this.Vector.Size(); _ = size
            for index := int(0); index < size; index++ {
                stringBuffer.Append(TESTING)
                var serverName string = any(this.Vector.Get(index)).(string); _ = serverName
                if tryControl2 := func() (tryControl2 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    list.Add(serverName)
                    this.LogUtil.Put("Test Error: " + fmt.Sprint(stringBuffer), this, "test", e)
                        }
                    }()
                    var licenseInitInfo *LicenseInitInfo = NewLicenseInitInfo(); _ = licenseInitInfo
                    licenseInitInfo.AddServer(serverName)
                    LicenseInitInfoUtil_GetInstance().Write(licenseInitInfo)
                    var abeClientInformation AbeClientInformationInterface = AbeClientInformationInterfaceFactory_GetInstance(); _ = abeClientInformation
                    var xmlRpcAbeLicenseClient *XmlRpcAbeLicenseRetrievalClient = NewXmlRpcAbeLicenseRetrievalClient(abeClientInformation); _ = xmlRpcAbeLicenseClient
                    var abeLicenseInterface AbeLicenseInterface = any(xmlRpcAbeLicenseClient.Get(nil)).(AbeLicenseInterface); _ = abeLicenseInterface
                    if abeLicenseInterface != nil {
                        stringBuffer.Append(PROCESSING)
                        stringBuffer.Append(fmt.Sprint(abeLicenseInterface))
                    }
                    stringBuffer.Append(END)
                    this.LogUtil.PutF(RESULT + fmt.Sprint(stringBuffer), this, "test")
                    stringBuffer.Delete(0, stringBuffer.Length())
                    return
                }(); tryControl2 == 1 {
                    tryControl1 = 1; return
                }
            }
            this.LogUtil.PutF("Total Servers Failed: " + fmt.Sprint(list.Size()) + " Failed Servers: " + fmt.Sprint(list), this, "test")
            return
        }(); tryControl1 == 1 {
            return
        }
}


func LicenseClientAndServerTest_Main(args []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, "main()", "main()", e)
                }
            }()
            var licensingTest *LicenseClientAndServerTest = NewLicenseClientAndServerTest(args); _ = licensingTest
            licensingTest.Test()
            return
        }(); tryControl1 == 1 {
            return
        }
}

