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

type AbeClientInformationData struct {
    KEY string
    NAME string
    VERSION string
    SPECIALNAME string
    LICENSEID string
    LICENSE_TYPE string
    PREVIOUSLICENSEID string
    OSNAME string
    OSARCH string
    OSVERSION string
    OS string
    HARDWARE string
    LICENSESERVERS string
    NEWLICENSE string
    ISNEW string
    SPECIAL string
}

var AbeClientInformationData_Instance *AbeClientInformationData = NewAbeClientInformationData()

func (this *AbeClientInformationData) initFields() {
    this.KEY = "KEY"
    this.NAME = "NAME"
    this.VERSION = "VERSION"
    this.SPECIALNAME = "SPECIALNAME"
    this.LICENSEID = "LICENSEID"
    this.LICENSE_TYPE = "LICENSETYPE"
    this.PREVIOUSLICENSEID = "PREVIOUSLICENSEID"
    this.OSNAME = "OSNAME"
    this.OSARCH = "OSARCH"
    this.OSVERSION = "OSVERSION"
    this.OS = "OS"
    this.HARDWARE = "HARDWARE"
    this.LICENSESERVERS = "LICENSESERVERS"
    this.NEWLICENSE = "NEWLICENSE"
    this.ISNEW = "ISNEW"
    this.SPECIAL = "SPECIAL"
}


func AbeClientInformationData_GetInstance() *AbeClientInformationData {
        return AbeClientInformationData_Instance
}


func NewAbeClientInformationData() *AbeClientInformationData {
    this := &AbeClientInformationData{}
    this.initFields()
    return this
}

