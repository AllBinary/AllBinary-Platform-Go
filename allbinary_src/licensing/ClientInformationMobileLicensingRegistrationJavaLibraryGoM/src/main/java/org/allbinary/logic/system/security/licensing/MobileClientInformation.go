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

type MobileClientInformation struct {
    AbeClientInformation
}

var MobileClientInformation_DESC string = "Mobile"
var MobileClientInformation_ANDROID_DESC string = "Android" + fmt.Sprint(MobileClientInformation_DESC)


func NewMobileClientInformation(name string, version string, specialName string, shortName string) *MobileClientInformation {
    this := &MobileClientInformation{}
        this.AbeClientInformation = *NewAbeClientInformation(name, version, fmt.Sprint(NewStringMaker().Append(specialName).Append(CommonSeps_GetInstance().SPACE).Append(PartnerIdentifierFileUtil_GetInstance().Get())), shortName)
        var string_ string = this.GetSpecialName(); _ = string_
        var number string = string_[len(string_) - 1:]; _ = number
        if Integer_GetInteger(number) == nil {
            this.SetHardSale(false)
        }
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if this.IsHardSale() {
            this.LogUtil.PutF("Use Hard Coded Selling", this, commonStrings.CONSTRUCTOR)
        } else {
            this.LogUtil.PutF("Don't Use Hard Coded Selling", this, commonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *MobileClientInformation) ToHashtable() *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = this.AbeClientInformation.ToHashtable(); _ = hashtable
        var registrationConfiguration *RegistrationConfiguration = RegistrationConfiguration_GetInstance(); _ = registrationConfiguration
        hashtable.Put(registrationConfiguration.NAME, registrationConfiguration.GetRegistrationCode())
        return hashtable
}

