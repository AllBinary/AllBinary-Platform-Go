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

// This is the information sent to the license server
type ClientInformation struct {
    OperatingSystemInterface *GenericOperatingSystem
    TsUtil *TsUtil
    Name string
    ShortName string
    Version string
    SpecialName string
    HardSale bool
    LicenseId string
    LicenseServers *BasicArrayList
}

// Implements AbeClientInformationInterface


func (this *ClientInformation) initFields() {
    this.OperatingSystemInterface = NoOperatingSystem_NO_OPERATING_SYSTEM
    this.TsUtil = TsUtil_GetInstance()
    
    
    
    
    this.HardSale = true
    this.LicenseId = StringUtil_GetInstance().EMPTY_STRING
    this.LicenseServers = BasicArrayListUtil_GetInstance().GetImmutableInstance()
}


func NewClientInformation(name string, version string, specialName string, shortName string) *ClientInformation {
    this := &ClientInformation{}
    this.initFields()
        this.Name = name
        this.Version = version
        this.SpecialName = specialName
        this.ShortName = fmt.Sprint(NewStringMaker().Append(shortName).Append(CommonSeps_GetInstance().UNDERSCORE).Append(version))
        this.Init()
    return this
}


func (this *ClientInformation) Init() {
}


func (this *ClientInformation) GetOperatingSystemInterface() *GenericOperatingSystem {
        return this.OperatingSystemInterface
}


func (this *ClientInformation) GetName() string {
        return this.Name
}


func (this *ClientInformation) GetSpecialName() string {
        return this.SpecialName
}


func (this *ClientInformation) GetVersion() string {
        return this.Version
}


func (this *ClientInformation) GetLicenseId() string {
        return this.LicenseId
}


func (this *ClientInformation) GetLicenseServers() *BasicArrayList {
        return this.LicenseServers
}


func (this *ClientInformation) GetLicenseServer(index int) string {
        return any(this.LicenseServers.ObjectArray[index]).(string)
}


func (this *ClientInformation) GetNumberOfLicenseServers() int {
        return this.LicenseServers.Size()
}


func (this *ClientInformation) ToHashtable() *ABHashtable[any, any] {
        var abeClientInformationData *AbeClientInformationData = AbeClientInformationData_GetInstance(); _ = abeClientInformationData
        var genericOperatingSystem *GenericOperatingSystem = this.GetOperatingSystemInterface(); _ = genericOperatingSystem
        var clientInfoHashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = clientInfoHashtable
        clientInfoHashtable.Put(abeClientInformationData.NAME, this.Name)
        clientInfoHashtable.Put(abeClientInformationData.VERSION, this.Version)
        clientInfoHashtable.Put(abeClientInformationData.SPECIALNAME, this.SpecialName)
        clientInfoHashtable.Put(abeClientInformationData.LICENSEID, this.GetLicenseId())
        clientInfoHashtable.Put(abeClientInformationData.OSNAME, genericOperatingSystem.GetName())
        clientInfoHashtable.Put(abeClientInformationData.OSARCH, genericOperatingSystem.GetArch())
        clientInfoHashtable.Put(abeClientInformationData.OSVERSION, genericOperatingSystem.GetVersion())
        clientInfoHashtable.Put(abeClientInformationData.OS, fmt.Sprint(genericOperatingSystem))
        var hardwareInterface HardwareInterface = SystemHardwareFactory_GetInstance().GetInstance(genericOperatingSystem); _ = hardwareInterface
        var hardwareString string = fmt.Sprint(hardwareInterface); _ = hardwareString
        clientInfoHashtable.Put(abeClientInformationData.HARDWARE, hardwareString)
        return clientInfoHashtable
}


func (this *ClientInformation) IsSameId(alicenseId string) bool {
        if this.TsUtil.CompareTo(this.GetLicenseId(), alicenseId) == 0 {
            return true
        } else {
            return false
        }
}


func (this *ClientInformation) IsLargerOrDifferentServerList(list *BasicArrayList) bool {
        var size int = list.Size() - 1; _ = size
        var currentSize int = this.LicenseServers.Size() - 1; _ = currentSize
        if size > currentSize {
            return true
        }
        for index := int(size); index >= 0; index-- {
            var newServerName string = any(list.ObjectArray[index]).(string); _ = newServerName
            for index2 := int(currentSize); index2 >= 0; index2-- {
                var nextOldServerName string = any(this.LicenseServers.ObjectArray[index]).(string); _ = nextOldServerName
                if this.TsUtil.CompareTo(newServerName, nextOldServerName) != 0 {
                    return true
                }
            }
        }
        return false
}


func (this *ClientInformation) SetHardSale(hardSale bool) {
        this.HardSale = hardSale
}


func (this *ClientInformation) IsHardSale() bool {
        return this.HardSale
}


func (this *ClientInformation) SetOperatingSystemInterface(operatingSystemInterface *GenericOperatingSystem) {
        this.OperatingSystemInterface = operatingSystemInterface
}


func (this *ClientInformation) SetLicenseId(licenseId string) {
        this.LicenseId = licenseId
}


func (this *ClientInformation) SetLicenseServers(licenseServers *BasicArrayList) {
        this.LicenseServers = licenseServers
}


func (this *ClientInformation) ToShortString() string {
        return this.ShortName
}

