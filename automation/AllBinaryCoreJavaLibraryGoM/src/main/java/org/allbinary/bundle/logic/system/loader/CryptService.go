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

type CryptService struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbCryptUtil *AbCryptUtil
}


func (this *CryptService) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AbCryptUtil = AbCryptUtil_GetInstance()
}


func NewCryptService() *CryptService {
    this := &CryptService{}
    this.initFields()
        Globals_GetInstance().Init(reflect.TypeOf(this).GetClassLoader(), "./")
        this.LogUtil.PutF("Set Globals: " + fmt.Sprint(URLGLOBALS.GetWebappPath()), this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *CryptService) GetDecryptedInputStream(abeClientInformation AbeClientInformationInterface, name string, inputStream IInputStream) IInputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IInputStream) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.ShowLicenseDialog(abeClientInformation, e)
                }
            }()
            var key string = AbKeys_GetInstance().GetKey(abeClientInformation, name); _ = key
            var decrypted []int8 = this.AbCryptUtil.Decrypt(inputStream, key); _ = decrypted
            tryControl1, tryResult1 = 1, NewByteArrayInputStream(decrypted); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *CryptService) ShowLicenseDialog(abeClientInformation AbeClientInformationInterface, e error) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e3 := Java_ToError(recovered)
                    _ = e3
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e3)
                }
            }()
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
            var basicTextJDialog *BasicTextJDialog = NewBasicTextJDialog(e.Error()); _ = basicTextJDialog
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e2)
                    }
                }()
                var abeLicenseInterface AbeLicenseInterface = AbeLicenseInterfaceFactory_GetInstance().GetLicenseInstance(abeClientInformation); _ = abeLicenseInterface
                if abeLicenseInterface != AbeNoLicense_GetInstance() {
                    if abeLicenseInterface.IsValid() {
                        basicTextJDialog.SetText("Subscription Invalid")
                    }
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            basicTextJDialog.AddCloseListener(NewExitCloseListener())
            basicTextJDialog.SetVisible(true)
            return
        }(); tryControl1 == 1 {
            return
        }
}

