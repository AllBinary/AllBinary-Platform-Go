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

type LicenseRegistrationUtil struct {
}



func NewLicenseRegistrationUtil() *LicenseRegistrationUtil {
    this := &LicenseRegistrationUtil{}
    return this
}


// synchronized
func LicenseRegistrationUtil_Process(abeClientInformation AbeClientInformationInterface, registrationId string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("License Registration Exception", "LicenseRegistrationUtil", "License Registration", e)
                }
            }()
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_Put(fmt.Sprint(CommonLabels_GetInstance().START_LABEL) + "License Registration", "LicenseRegistrationUtil", commonStrings.PROCESS)
            var registrationConfiguration *RegistrationConfiguration = RegistrationConfiguration_GetInstance(); _ = registrationConfiguration
            registrationConfiguration.SetRegistrationCode(registrationId)
            registrationConfiguration.Write()
            var hashtable *ABHashtable[any, any] = abeClientInformation.ToHashtable(); _ = hashtable
            hashtable.Put(RegistrationConfiguration_GetInstance().NAME, registrationId)
            hashtable.Put("message", SpecialMessageUtil_GetInstance().Get())
            NewXmlRpcRemoteLicenseRegistrationClient(abeClientInformation).Get(hashtable)
            return
        }(); tryControl1 == 1 {
            return
        }
}

