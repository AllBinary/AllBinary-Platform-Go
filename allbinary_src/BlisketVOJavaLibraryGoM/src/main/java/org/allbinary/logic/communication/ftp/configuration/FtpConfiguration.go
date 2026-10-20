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

type FtpConfiguration struct {
    FtpServer string
    FtpPath string
    FtpUserName string
    FtpPassword string
}

// Implements FtpConfigurationInterface


func (this *FtpConfiguration) initFields() {
    
    
    
    
}




// Java overloads of new FtpConfiguration
func NewFtpConfiguration(args ...any) *FtpConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewFtpConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewFtpConfiguration_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new FtpConfiguration for the arguments")
}
func NewFtpConfiguration_1() *FtpConfiguration {
    this := &FtpConfiguration{}
    this.initFields()
    return this
}


func NewFtpConfiguration_2(hashMap *HashMap) *FtpConfiguration {
    this := &FtpConfiguration{}
    this.initFields()
        this.FtpServer = any(hashMap.Get(FtpConfigurationData_SERVER)).(string)
        this.FtpPath = any(hashMap.Get(FtpConfigurationData_PATH)).(string)
        this.FtpUserName = any(hashMap.Get(FtpConfigurationData_USERNAME)).(string)
        this.FtpPassword = any(hashMap.Get(FtpConfigurationData_PASSWORD)).(string)
    return this
}


func (this *FtpConfiguration) GetServer() string {
        return this.FtpServer
}


func (this *FtpConfiguration) GetPath() string {
        return this.FtpPath
}


func (this *FtpConfiguration) GetUserName() string {
        return this.FtpUserName
}


func (this *FtpConfiguration) GetPassword() string {
        return this.FtpPassword
}


func (this *FtpConfiguration) SetServer(value string) {
        this.FtpServer = value
}


func (this *FtpConfiguration) SetPath(value string) {
        this.FtpPath = value
}


func (this *FtpConfiguration) SetUserName(value string) {
        this.FtpUserName = value
}


func (this *FtpConfiguration) SetPassword(value string) {
        this.FtpPassword = value
}

