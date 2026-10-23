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

type AbeClientLicense struct {
    Hashtable *ABHashtable[any, any]
    Id string
    Servers IBasicArrayList
    Special string
    LicenseType *LicenseType
}

// Implements AbeLicenseInterface


func (this *AbeClientLicense) initFields() {
    
    
    
    
    
}


func NewAbeClientLicense(hashtable *ABHashtable[any, any]) *AbeClientLicense {
    this := &AbeClientLicense{}
    this.initFields()
        this.Hashtable = StdUtil_GetInstance().CreateHashtable()
        var abeClientInformationData *AbeClientInformationData = AbeClientInformationData_GetInstance(); _ = abeClientInformationData
        var keyValue any = hashtable.Get(abeClientInformationData.KEY); _ = keyValue
        if keyValue != nil {
            if func() bool { _, ok := any(keyValue).(string); return ok }() {
                this.Hashtable.Put(abeClientInformationData.KEY, keyValue)
            }
        }
        this.Id = any(hashtable.Get(abeClientInformationData.LICENSEID)).(string)
        var vector IBasicArrayList = any(hashtable.Get(abeClientInformationData.LICENSESERVERS)).(IBasicArrayList); _ = vector
        this.Servers = NewBasicArrayListD()
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            this.Servers.Add(vector.Get(index))
        }
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.SetSpecial(stringUtil.GetNonNull(any(hashtable.Get(abeClientInformationData.SPECIAL)).(string)))
        var licenseTypeString string = stringUtil.GetNonNull(any(hashtable.Get(abeClientInformationData.LICENSE_TYPE)).(string)); _ = licenseTypeString
        this.LicenseType = LicenseTypeFactory_GetInstance().GetInstance(licenseTypeString)
    return this
}


func AbeClientLicense_HasRequiredKeys(resultHashtable *ABHashtable[any, any]) bool {
        var abeClientInformationData *AbeClientInformationData = AbeClientInformationData_GetInstance(); _ = abeClientInformationData
        if resultHashtable.ContainsKey(abeClientInformationData.LICENSEID) && resultHashtable.ContainsKey(abeClientInformationData.LICENSESERVERS) && resultHashtable.ContainsKey(abeClientInformationData.KEY) && resultHashtable.ContainsKey(abeClientInformationData.LICENSE_TYPE) {
            return true
        } else {
            return false
        }
}


func (this *AbeClientLicense) HasKey() bool {
        return this.IsValid()
}


func (this *AbeClientLicense) GetKey(keyName string) string {
        return any(this.Hashtable.Get(keyName)).(string)
}


func (this *AbeClientLicense) GetLicenseId() string {
        return this.Id
}


func (this *AbeClientLicense) GetServers() IBasicArrayList {
        return this.Servers
}


func (this *AbeClientLicense) IsValid() bool {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsEmpty(this.GetKey(AbeClientInformationData_GetInstance().KEY)) {
            return true
        }
        return false
}


func (this *AbeClientLicense) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var BREAK string = "<br/>"; _ = BREAK
        stringBuffer.Append("License Id: ")
        stringBuffer.Append(this.GetLicenseId())
        stringBuffer.Append(BREAK)
        stringBuffer.Append("Is Valid: ")
        stringBuffer.Appendboolean(this.IsValid())
        stringBuffer.Append(BREAK)
        stringBuffer.Append("Keys: ")
        stringBuffer.Append(fmt.Sprint(this.Hashtable))
        stringBuffer.Append(BREAK)
        var serverVector IBasicArrayList = this.GetServers(); _ = serverVector
        var size int = serverVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var nextServerString string = any(serverVector.Get(index)).(string); _ = nextServerString
            stringBuffer.Append("Server: ")
            stringBuffer.Append(nextServerString)
            stringBuffer.Append(BREAK)
        }
        return fmt.Sprint(stringBuffer)
}

func (this *AbeClientLicense) String() string {
    return this.ToString()
}


func (this *AbeClientLicense) SetSpecial(special string) {
        this.Special = special
}


func (this *AbeClientLicense) GetSpecial() string {
        return this.Special
}


func (this *AbeClientLicense) GetLicenseType() *LicenseType {
        return this.LicenseType
}

