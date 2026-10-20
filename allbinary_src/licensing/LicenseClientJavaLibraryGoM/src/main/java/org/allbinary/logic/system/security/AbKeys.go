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

type AbKeys struct {
    LogUtil *LogUtil
}

var AbKeys_Instance *AbKeys = NewAbKeys()

func (this *AbKeys) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func AbKeys_GetInstance() *AbKeys {
        return AbKeys_Instance
}


func NewAbKeys() *AbKeys {
    this := &AbKeys{}
    this.initFields()
    return this
}


// throws LicensingException
// synchronized
func (this *AbKeys) GetKey(abeClientInformation AbeClientInformationInterface, keyName string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.LogUtil.PutF("Getting Key: " + keyName, this, "getKey")
            if strings.Compare(keyName, "DirectX 8") == 0 || strings.Compare(keyName, "DirectX") == 0 || strings.Compare(keyName, "Low Level") == 0 || strings.Compare(keyName, "System Drivers") == 0 {
                tryControl1, tryResult1 = 1, "Temp For Input Library"; return
            }
            tryControl1, tryResult1 = 1, AbeLicenseInterfaceFactory_GetInstance().GetLicenseInstance(abeClientInformation).GetKey(keyName); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

