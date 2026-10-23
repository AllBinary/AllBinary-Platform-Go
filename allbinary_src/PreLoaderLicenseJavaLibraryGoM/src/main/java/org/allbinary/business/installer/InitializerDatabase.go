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

type InitializerDatabase struct {
    LogUtil *LogUtil
    AdminDbUserName string
    AdminDbPassword string
    AdminJdbcDriver string
    AdminSchema string
    AdminServer string
    AdminPort string
    InitDb *DynamicInitDb
}

var InitializerDatabase_MAXDB int = 30
var InitializerDatabase_MAX int = 16
var InitializerDatabase_MIN int = 4
var InitializerDatabase_MINPASSWORD int = 0

func (this *InitializerDatabase) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
}




// Java overloads of new InitializerDatabase
func NewInitializerDatabase(args ...any) *InitializerDatabase {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[Map](args[1], pass) {
            return NewInitializerDatabase_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[Map](args[1]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewInitializerDatabase_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of new InitializerDatabase for the arguments")
}
func NewInitializerDatabase_1(abeClientInformation AbeClientInformationInterface, map_ Map) *InitializerDatabase {
    this := &InitializerDatabase{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var keys Set = map_.KeySet(); _ = keys
        var keyArray []any = keys.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        for index := int(0); index < size; index++ {
            var key string = any(keyArray[index]).(string); _ = key
            var values []string = any(map_.Get(key)).([]string); _ = values
            hashMap.Put(key, values[0])
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("key: ")
            stringBuffer.Append(key)
            stringBuffer.Append(" Value: ")
            stringBuffer.Append(values[0])
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getFormData()")
        }
        this.GetFormData(abeClientInformation, hashMap)
    return this
}


func NewInitializerDatabase_2(abeClientInformation AbeClientInformationInterface, initHashMap *HashMap) *InitializerDatabase {
    this := &InitializerDatabase{}
    this.initFields()
        this.GetFormData(abeClientInformation, initHashMap)
    return this
}


func (this *InitializerDatabase) GetFormData(abeClientInformation AbeClientInformationInterface, hashMap *HashMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to get form data", this, "getFormData()", e)
                }
            }()
            var initializerData *InitializerData = InitializerData_GetInstance(); _ = initializerData
            this.SetAdminDbUserName(any(hashMap.Get(initializerData.DBUSER)).(string))
            this.SetAdminDbPassword(any(hashMap.Get(initializerData.DBPASSWORD)).(string))
            this.SetAdminJdbcDriver(any(hashMap.Get(initializerData.ADMINJDBCDRIVER)).(string))
            this.SetAdminSchema(any(hashMap.Get(initializerData.ADMINSCHEMA)).(string))
            this.SetAdminServer(any(hashMap.Get(initializerData.ADMINSERVER)).(string))
            this.SetAdminPort(any(hashMap.Get(initializerData.ADMINPORT)).(string))
            var dbConnectionInfo IDbConnectionInfo = NewDbConnectionInfo(); _ = dbConnectionInfo
            var adminDbName string = StringUtil_GetInstance().EMPTY_STRING; _ = adminDbName
            dbConnectionInfo.SetJdbcDriver(this.GetAdminJdbcDriver())
            dbConnectionInfo.SetName(adminDbName)
            dbConnectionInfo.SetUserName(this.GetAdminDbUserName())
            dbConnectionInfo.SetPassword(this.GetAdminDbPassword())
            dbConnectionInfo.SetSchema(this.GetAdminSchema())
            dbConnectionInfo.SetServer(this.GetAdminServer())
            dbConnectionInfo.SetPort(this.GetAdminPort())
            this.InitDb = NewDynamicInitDb(abeClientInformation, any(dbConnectionInfo).(DatabaseConnectionInfoInterface))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InitializerDatabase) IsJdbcDriverValid(jdbcDriverClassPathString string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADERERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            Class_ForName(jdbcDriverClassPathString).NewInstance()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitializerDatabase) IsValid() bool {
        var isValid bool = true; _ = isValid
        if !this.IsJdbcDriverValid(this.AdminJdbcDriver) {
            isValid = false
        }
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.AdminDbUserName, InitializerDatabase_MIN, InitializerDatabase_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.AdminDbPassword, InitializerDatabase_MINPASSWORD, InitializerDatabase_MAX) {
            isValid = false
        }
        return isValid
}


func (this *InitializerDatabase) GetJdbcDriverValidationInfo(jdbcDriver string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("The JDBC driver (")
        stringBuffer.Append(jdbcDriver)
        stringBuffer.Append(") you have provided is not valid.<br/>")
        return fmt.Sprint(stringBuffer)
}


func (this *InitializerDatabase) GetInvalidInfo() string {
        var isValid bool = true; _ = isValid
        var isJdbcDriversValid bool = true; _ = isJdbcDriversValid
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if !this.IsJdbcDriverValid(this.AdminJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.GetAdminJdbcDriver()))
        }
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.AdminDbUserName, InitializerDatabase_MIN, InitializerDatabase_MAX) {
            stringBuffer.Append("Admin username should be < " + fmt.Sprint(InitializerDatabase_MAX) + " and > " + fmt.Sprint(InitializerDatabase_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.AdminDbPassword, InitializerDatabase_MINPASSWORD, InitializerDatabase_MAX) {
            stringBuffer.Append("Admin password should be < " + fmt.Sprint(InitializerDatabase_MAX) + " and > " + fmt.Sprint(InitializerDatabase_MINPASSWORD) + " characters in length.<br />")
        }
        if !isJdbcDriversValid {
            stringBuffer.Append(InitializerData_GetInstance().GetJdbcDriverSolutionInfo())
        }
        return fmt.Sprint(stringBuffer)
}


func (this *InitializerDatabase) CreateUsers() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to Create Users", this, "createUsers()", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.InitDb.AddUsers()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitializerDatabase) CreateDatabases() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to Create Databases", this, "createDatabases()", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.InitDb.AddDatabases()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitializerDatabase) CreateTables() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to Create Tables", this, "createTables()", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.InitDb.AddTables()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitializerDatabase) GetAdminDbUserName() string {
        return this.AdminDbUserName
}


func (this *InitializerDatabase) SetAdminDbUserName(adminDbUserName string) {
        this.AdminDbUserName = adminDbUserName
}


func (this *InitializerDatabase) GetAdminDbPassword() string {
        return this.AdminDbPassword
}


func (this *InitializerDatabase) SetAdminDbPassword(adminDbPassword string) {
        this.AdminDbPassword = adminDbPassword
}


func (this *InitializerDatabase) GetAdminJdbcDriver() string {
        return this.AdminJdbcDriver
}


func (this *InitializerDatabase) SetAdminJdbcDriver(adminJdbcDriver string) {
        this.AdminJdbcDriver = adminJdbcDriver
}


func (this *InitializerDatabase) GetAdminSchema() string {
        return this.AdminSchema
}


func (this *InitializerDatabase) SetAdminSchema(adminSchema string) {
        this.AdminSchema = adminSchema
}


func (this *InitializerDatabase) GetAdminServer() string {
        return this.AdminServer
}


func (this *InitializerDatabase) SetAdminServer(adminServer string) {
        this.AdminServer = adminServer
}


func (this *InitializerDatabase) GetAdminPort() string {
        return this.AdminPort
}


func (this *InitializerDatabase) SetAdminPort(adminPort string) {
        this.AdminPort = adminPort
}

