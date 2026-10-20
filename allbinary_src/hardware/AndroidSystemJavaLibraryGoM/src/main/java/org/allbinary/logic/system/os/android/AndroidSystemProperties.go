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

type AndroidSystemProperties struct {
    CommonStrings *CommonStrings
}

var AndroidSystemProperties_SINGLETON *AndroidSystemProperties = NewAndroidSystemProperties(ResourceUtil_GetInstance().GetContext())

func (this *AndroidSystemProperties) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}


func AndroidSystemProperties_GetInstance() *AndroidSystemProperties {
        return AndroidSystemProperties_SINGLETON
}


func NewAndroidSystemProperties(context *Context) *AndroidSystemProperties {
    this := &AndroidSystemProperties{}
    this.initFields()
    return this
}


func (this *AndroidSystemProperties) GetDeviceId() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *AndroidSystemProperties) GetDeviceSoftwareVersion() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetLine1Number() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetNetworkCountryIso() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetNetworkOperator() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetNetworkOperatorName() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetSimCountryIso() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetSimOperator() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetSimOperatorName() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetSimSerialNumber() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetSubscriberId() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetVoiceMailAlphaTag() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetVoiceMailNumber() string {
        return this.CommonStrings.DISABLE
}


func (this *AndroidSystemProperties) GetNetworkType() int {
        return -1
}


func (this *AndroidSystemProperties) GetPhoneType() int {
        return -1
}


func (this *AndroidSystemProperties) GetBoard() string {
        return Build_BOARD
}


func (this *AndroidSystemProperties) GetBrand() string {
        return Build_BRAND
}


func (this *AndroidSystemProperties) GetDevice() string {
        return Build_DEVICE
}


func (this *AndroidSystemProperties) GetFingerprint() string {
        return Build_FINGERPRINT
}


func (this *AndroidSystemProperties) GetHost() string {
        return Build_HOST
}


func (this *AndroidSystemProperties) GetId() string {
        return Build_ID
}


func (this *AndroidSystemProperties) GetModel() string {
        return Build_MODEL
}


func (this *AndroidSystemProperties) GetProduct() string {
        return Build_PRODUCT
}


func (this *AndroidSystemProperties) GetTags() string {
        return Build_TAGS
}


func (this *AndroidSystemProperties) GetTime() int64 {
        return Build_TIME
}


func (this *AndroidSystemProperties) GetType() string {
        return Build_TYPE
}


func (this *AndroidSystemProperties) GetUser() string {
        return Build_USER
}

