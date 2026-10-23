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

type DbConnectionInfo struct {
    JdbcDriver string
    Name string
    UserName string
    Password string
    Schema string
    Server string
    Port string
    Url string
    Host string
    SCHEMA_SEP string
    USER_NAME_KEY string
    PASSWORD_KEY string
}

// Implements DatabaseConnectionInfoInterface

// The Java class DbConnectionInfo has subclasses so its variables are this interface
type IDbConnectionInfo interface {
    DbConnectionInfo_() *DbConnectionInfo
    UpdateUrl()
    GetUrl() string
    GetJdbcDriver() string
    GetName() string
    GetUserName() string
    GetPassword() string
    UpdateHost()
    GetHost() string
    GetSchema() string
    GetServer() string
    GetPort() string
    SetJdbcDriver(value string)
    SetName(value string)
    SetUserName(value string)
    SetPassword(value string)
    SetSchema(value string)
    SetServer(value string)
    SetPort(value string)
    GetUserNameKey() string
    GetPasswordKey() string
}

func (this *DbConnectionInfo) DbConnectionInfo_() *DbConnectionInfo {
    return this
}


func (this *DbConnectionInfo) initFields() {
    
    
    
    
    
    
    
    
    
    this.SCHEMA_SEP = "://"
    this.USER_NAME_KEY = "?user="
    this.PASSWORD_KEY = "&password="
}


func NewDbConnectionInfo() *DbConnectionInfo {
    this := &DbConnectionInfo{}
    this.initFields()
    return this
}


func (this *DbConnectionInfo) UpdateUrl() {
        this.UpdateHost()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.GetHost())
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(this.GetUserNameKey())
        stringBuffer.Append(this.GetUserName())
        stringBuffer.Append(this.GetPasswordKey())
        stringBuffer.Append(this.GetPassword())
        this.Url = fmt.Sprint(stringBuffer)
}


func (this *DbConnectionInfo) GetUrl() string {
        return this.Url
}


func (this *DbConnectionInfo) GetJdbcDriver() string {
        return this.JdbcDriver
}


func (this *DbConnectionInfo) GetName() string {
        return this.Name
}


func (this *DbConnectionInfo) GetUserName() string {
        return this.UserName
}


func (this *DbConnectionInfo) GetPassword() string {
        return this.Password
}


func (this *DbConnectionInfo) UpdateHost() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.GetSchema())
        stringBuffer.Append(this.SCHEMA_SEP)
        stringBuffer.Append(this.GetServer())
        if this.GetPort() != nil && len(this.GetPort()) > 1 {
            stringBuffer.Append(CommonSeps_GetInstance().COLON)
            stringBuffer.Append(this.GetPort())
        }
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        this.Host = fmt.Sprint(stringBuffer)
}


func (this *DbConnectionInfo) GetHost() string {
        return this.Host
}


func (this *DbConnectionInfo) GetSchema() string {
        return this.Schema
}


func (this *DbConnectionInfo) GetServer() string {
        return this.Server
}


func (this *DbConnectionInfo) GetPort() string {
        return this.Port
}


func (this *DbConnectionInfo) SetJdbcDriver(value string) {
        this.JdbcDriver = value
}


func (this *DbConnectionInfo) SetName(value string) {
        this.Name = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) SetUserName(value string) {
        this.UserName = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) SetPassword(value string) {
        this.Password = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) SetSchema(value string) {
        this.Schema = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) SetServer(value string) {
        this.Server = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) SetPort(value string) {
        this.Port = value
        this.UpdateUrl()
}


func (this *DbConnectionInfo) GetUserNameKey() string {
        return this.USER_NAME_KEY
}


func (this *DbConnectionInfo) GetPasswordKey() string {
        return this.PASSWORD_KEY
}

