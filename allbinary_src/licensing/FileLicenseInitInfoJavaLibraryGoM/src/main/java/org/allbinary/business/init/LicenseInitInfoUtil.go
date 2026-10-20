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
    LogUtil *LogUtil
    StringUtil *StringUtil
    INITFILENAME string
    ABOUT string
    PRIVACY_POLICY string
    FilePath string
}

var LicenseInitInfoUtil_Instance *LicenseInitInfoUtil = NewLicenseInitInfoUtil()

func (this *LicenseInitInfoUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.INITFILENAME = "licenseinitdata.dat"
    this.ABOUT = "about"
    this.PRIVACY_POLICY = "privacy_policy"
    this.FilePath = this.StringUtil.EMPTY_STRING
}

func NewLicenseInitInfoUtil() *LicenseInitInfoUtil {
    this := &LicenseInitInfoUtil{}
    this.initFields()
    return this
}


func LicenseInitInfoUtil_GetInstance() *LicenseInitInfoUtil {
        return LicenseInitInfoUtil_Instance
}


// synchronized
func (this *LicenseInitInfoUtil) SetFilePath(filePath string) {
        this.FilePath = filePath
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) Write(initData *LicenseInitInfo) {
        if this.FilePath == this.StringUtil.EMPTY_STRING {
            this.FilePath = URLGLOBALS.GetMainPath() + PATH_GLOBALS.GetInstance().INIT_PATH
        }
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Command Failed: " + this.INITFILENAME, this, "write", e)
            FileStreamFactory_GetInstance().Delete(this.FilePath, this.INITFILENAME)
            panic(e)
                }
            }()
            var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstance(this.FilePath, this.INITFILENAME); _ = dataOutputStream
            var licenseIdCrypted []int8 = tsUtil.GetByteArray(NewWeakCrypt(1).Encrypt(initData.GetLicenseId())); _ = licenseIdCrypted
            dataOutputStream.WriteUTF(DatabaseEncoder_Encode(licenseIdCrypted))
            var numberOfLicenseServers int = initData.GetNumberOfServers(); _ = numberOfLicenseServers
            dataOutputStream.WriteInt(numberOfLicenseServers)
            var licenseServerCrypted []int8; _ = licenseServerCrypted
            for index := int(0); index < numberOfLicenseServers; index++ {
                licenseServerCrypted = tsUtil.GetByteArray(NewWeakCrypt(3).Encrypt(initData.GetServer(index)))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(licenseServerCrypted))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) Read() *LicenseInitInfo {
        return this.ReadAgain(0)
}


// throws Exception
// synchronized
func (this *LicenseInitInfoUtil) ReadAgain(initializeCounter int) *LicenseInitInfo {
        var METHOD_NAME string = "readAgain"; _ = METHOD_NAME
        if this.FilePath == this.StringUtil.EMPTY_STRING {
            this.FilePath = URLGLOBALS.GetMainPath() + PATH_GLOBALS.GetInstance().INIT_PATH
        }
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *LicenseInitInfo) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 *LicenseInitInfo) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                this.LogUtil.Put("LicenseInitInfo Read Retry: " + this.INITFILENAME, this, "readAgain()", e2)
                    }
                }()
                this.LogUtil.Put("Command Failed: " + this.INITFILENAME, this, METHOD_NAME, e)
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
            panic(errors.New(fmt.Sprint("LicenseInitInfo Read Error: " + this.INITFILENAME)))
                }
            }()
            this.LogUtil.PutF("LicenseInitInfo File: " + this.INITFILENAME, this, METHOD_NAME)
            var fileStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileStreamFactory
            var iFile *AbFileInputStream = fileStreamFactory.GetFileInputStreamInstance(this.FilePath, this.INITFILENAME); _ = iFile
            if iFile != nil {
                var iData *AbDataInputStream = NewAbDataInputStream(iFile); _ = iData
                var initInfo *LicenseInitInfo = NewLicenseInitInfo(); _ = initInfo
                var decodedByteArray []int8 = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decodedByteArray
                var licenseIdDecoded string = string(decodedByteArray); _ = licenseIdDecoded
                initInfo.SetLicenseId(NewWeakCrypt(1).Decrypt(licenseIdDecoded))
                var numberOfLicenseServers int = iData.ReadInt(); _ = numberOfLicenseServers
                var NEXT_FILE string = fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().NEXT).Append(" License Server From File: ")); _ = NEXT_FILE
                var licenseServerDecoded string; _ = licenseServerDecoded
                for index := int(0); index < numberOfLicenseServers; index++ {
                    decodedByteArray = DatabaseEncoder_Decode(iData.ReadUTF())
                    licenseServerDecoded = string(decodedByteArray)
                    initInfo.SetServer(NewWeakCrypt(3).Decrypt(licenseServerDecoded), index)
                    this.LogUtil.PutF(NEXT_FILE + fmt.Sprint(initInfo.GetServer(index)), this, METHOD_NAME)
                }
                tryControl1, tryResult1 = 1, initInfo; return
            } else {
                panic(errors.New(fmt.Sprint("Could Not Load License InitInfo: " + this.INITFILENAME)))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *LicenseInitInfoUtil) GetFilePath() string {
        return this.FilePath
}

