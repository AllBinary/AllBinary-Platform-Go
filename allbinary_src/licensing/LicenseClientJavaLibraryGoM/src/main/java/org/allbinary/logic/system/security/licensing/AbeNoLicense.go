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

type AbeNoLicense struct {
    ServerVector *BasicArrayList
}

// Implements AbeLicenseInterface

var AbeNoLicense_AbeNoLicense *AbeNoLicense = NewAbeNoLicense()

func (this *AbeNoLicense) initFields() {
    this.ServerVector = NewBasicArrayListD()
}


func NewAbeNoLicense() *AbeNoLicense {
    this := &AbeNoLicense{}
    this.initFields()
    return this
}


func AbeNoLicense_GetInstance() *AbeNoLicense {
        return AbeNoLicense_AbeNoLicense
}


func (this *AbeNoLicense) HasKey() bool {
        return false
}


func (this *AbeNoLicense) GetKey(keyName string) string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *AbeNoLicense) GetLicenseId() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *AbeNoLicense) GetSpecial() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *AbeNoLicense) GetServers() *BasicArrayList {
        return this.ServerVector
}


func (this *AbeNoLicense) IsValid() bool {
        return false
}


func (this *AbeNoLicense) GetLicenseType() *LicenseType {
        return LicenseTypeFactory_GetInstance().UNKNOWN
}

