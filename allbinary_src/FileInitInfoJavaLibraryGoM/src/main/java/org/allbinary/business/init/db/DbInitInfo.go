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

type DbInitInfo struct {
    DbConnectionInfo
    CommonStrings *CommonStrings
    Directory *Directory
    HttpData *HttpData
    GET_URL string
    PACKAGE string
    InitFileName string
    HasRead bool
}

// The Java class DbInitInfo has subclasses so its variables are this interface
type IDbInitInfo interface {
    IDbConnectionInfo
    DbInitInfo_() *DbInitInfo
    GetUrl() string
    Write()
    Load()
    SetHasRead(value bool)
    UpdateIfNeeded()
    GetName() string
    GetUserName() string
    GetPassword() string
    GetSchema() string
    GetServer() string
    GetPort() string
}

func (this *DbInitInfo) DbInitInfo_() *DbInitInfo {
    return this
}


func (this *DbInitInfo) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Directory = Directory_GetInstance()
    this.HttpData = HttpData_GetInstance()
    this.GET_URL = "getUrl"
    this.PACKAGE = DbConnectionInfo_FREEBLISKET_PATH_GLOBALS.GetInstance().DBINITPATH
    
    this.HasRead = false
}


func NewDbInitInfo(initFileName string, read bool) *DbInitInfo {
    this := &DbInitInfo{}
        this.DbConnectionInfo = *NewDbConnectionInfo()
        this.initFields()
        this.InitFileName = initFileName
        if read {
            this.UpdateIfNeeded()
        } else {
            this.SetHasRead(true)
        }
    return this
}


func (this *DbInitInfo) GetUrl() string {
        var url string = this.DbConnectionInfo.GetUrl(); _ = url
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
            PreLogUtil_Put(fmt.Sprint(this.HttpData.URL_LABEL) + url, this, this.GET_URL)
        }
        return url
}


// throws Exception
func (this *DbInitInfo) Write() {
        var PATH IAbPath = NewAbPath(fmt.Sprint(DbConnectionInfo_URLGLOBALS.GetWebappPath()) + this.PACKAGE, StringUtil_GetInstance().EMPTY_STRING); _ = PATH
        this.Directory.Create(PATH)
        var FILEABPATH IAbPath = NewAbPath(fmt.Sprint(DbConnectionInfo_URLGLOBALS.GetWebappPath()) + this.PACKAGE, this.InitFileName); _ = FILEABPATH
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("Failed Write: " + fmt.Sprint(FILEABPATH), this, "write", e)
            }
                }
            }()
            var newFile *AbFile = AbFile_CreateAbFileFromAbPath(FILEABPATH); _ = newFile
            newFile.CreateNewFile()
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(newFile); _ = dataOutputStream
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                StreamUtil_GetInstance().Close(dataOutputStream)
                }()
                var cryptedJdbcDriver []int8 = Java_GetBytes(NewWeakCrypt(1).Encrypt(this.GetJdbcDriver())); _ = cryptedJdbcDriver
                var cryptedName []int8 = Java_GetBytes(NewWeakCrypt(2).Encrypt(this.GetName())); _ = cryptedName
                var cryptedUserName []int8 = Java_GetBytes(NewWeakCrypt(3).Encrypt(this.GetUserName())); _ = cryptedUserName
                var cryptedPassword []int8 = Java_GetBytes(NewWeakCrypt(4).Encrypt(this.GetPassword())); _ = cryptedPassword
                var cryptedSchema []int8 = Java_GetBytes(NewWeakCrypt(5).Encrypt(this.GetSchema())); _ = cryptedSchema
                var cryptedServer []int8 = Java_GetBytes(NewWeakCrypt(6).Encrypt(this.GetServer())); _ = cryptedServer
                var cryptedPort []int8 = Java_GetBytes(NewWeakCrypt(7).Encrypt(this.GetPort())); _ = cryptedPort
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedJdbcDriver))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedName))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedUserName))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedPassword))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedSchema))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedServer))
                dataOutputStream.WriteUTF(DatabaseEncoder_Encode(cryptedPort))
                this.HasRead = false
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *DbInitInfo) Load() {
        var FILEABPATH IAbPath = NewAbPath(fmt.Sprint(DbConnectionInfo_URLGLOBALS.GetWebappPath()) + this.PACKAGE, this.InitFileName); _ = FILEABPATH
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("Failed Loading: " + fmt.Sprint(FILEABPATH), this, this.CommonStrings.LOAD, e)
            }
                }
            }()
            var file *AbFile = AbFile_CreateAbFileFromAbPath(FILEABPATH); _ = file
            if file.IsFile() {
                var iFile IAbFileInputStream = NewAbFileInputStream(file); _ = iFile
                var iData *AbDataInputStream = NewAbDataInputStream(iFile); _ = iData
                if tryControl2 := func() (tryControl2 int) {
                    defer func() {
                    StreamUtil_GetInstance().Close(iData)
                    }()
                    var decryptedJdbcDriver string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedJdbcDriver
                    var decryptedName string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedName
                    var decryptedUserName string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedUserName
                    var decryptedPassword string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedPassword
                    var decryptedSchema string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedSchema
                    var decryptedServer string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedServer
                    var decryptedPort string = DatabaseEncoder_Decode(iData.ReadUTF()); _ = decryptedPort
                    this.SetJdbcDriver(NewWeakCrypt(1).Decrypt(decryptedJdbcDriver))
                    this.SetName(NewWeakCrypt(2).Decrypt(decryptedName))
                    this.SetUserName(NewWeakCrypt(3).Decrypt(decryptedUserName))
                    this.SetPassword(NewWeakCrypt(4).Decrypt(decryptedPassword))
                    this.SetSchema(NewWeakCrypt(5).Decrypt(decryptedSchema))
                    this.SetServer(NewWeakCrypt(6).Decrypt(decryptedServer))
                    this.SetPort(NewWeakCrypt(7).Decrypt(decryptedPort))
                    var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
                    return
                }(); tryControl2 == 1 {
                    tryControl1 = 1; return
                }
            } else {
                this.HasRead = false
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                    PreLogUtil_Put("Not a File - Failed Loading: " + fmt.Sprint(FILEABPATH), this, this.CommonStrings.LOAD)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DbInitInfo) SetHasRead(value bool) {
        this.HasRead = value
}


func (this *DbInitInfo) UpdateIfNeeded() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, "updateIfNeeded", e)
            }
                }
            }()
            if !this.HasRead {
                this.HasRead = true
                this.Load()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DbInitInfo) GetName() string {
        this.UpdateIfNeeded()
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
            PreLogUtil_Put("Name: " + this.DbConnectionInfo.GetName(), this, "getName")
        }
        return this.DbConnectionInfo.GetName()
}


func (this *DbInitInfo) GetUserName() string {
        this.UpdateIfNeeded()
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
            PreLogUtil_Put("Name: " + fmt.Sprint(this.DbConnectionInfo.GetUserName()), this, "getUserName")
        }
        return this.DbConnectionInfo.GetUserName()
}


func (this *DbInitInfo) GetPassword() string {
        this.UpdateIfNeeded()
        return this.DbConnectionInfo.GetPassword()
}


func (this *DbInitInfo) GetSchema() string {
        this.UpdateIfNeeded()
        return this.DbConnectionInfo.GetSchema()
}


func (this *DbInitInfo) GetServer() string {
        this.UpdateIfNeeded()
        return this.DbConnectionInfo.GetServer()
}


func (this *DbInitInfo) GetPort() string {
        this.UpdateIfNeeded()
        return this.DbConnectionInfo.GetPort()
}

