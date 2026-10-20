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

type LicenseInitInfo struct {
    Id string
    ServerList *BasicArrayList
}


func (this *LicenseInitInfo) initFields() {
    this.Id = StringUtil_GetInstance().EMPTY_STRING
    this.ServerList = NewBasicArrayListD()
}


func NewLicenseInitInfo() *LicenseInitInfo {
    this := &LicenseInitInfo{}
    this.initFields()
    return this
}


func (this *LicenseInitInfo) GetLicenseId() string {
        return this.Id
}


func (this *LicenseInitInfo) GetServer(index int) string {
        return any(this.ServerList.Get(index)).(string)
}


func (this *LicenseInitInfo) GetNumberOfServers() int {
        return this.ServerList.Size()
}


func (this *LicenseInitInfo) SetLicenseId(value string) {
        this.Id = value
}


func (this *LicenseInitInfo) ClearServers() {
        this.ServerList.Clear()
}


func (this *LicenseInitInfo) SetServer(value string, index int) {
        this.ServerList.AddAt(index, value)
}


func (this *LicenseInitInfo) AddServer(value string) {
        this.ServerList.Add(value)
}


func (this *LicenseInitInfo) GetServerList() *BasicArrayList {
        return this.ServerList
}


func (this *LicenseInitInfo) SetServerList(servers *BasicArrayList) {
        if DebugFactory_GetInstance() == NoDebug_GetInstance() {
            this.ServerList = servers
        }
}

