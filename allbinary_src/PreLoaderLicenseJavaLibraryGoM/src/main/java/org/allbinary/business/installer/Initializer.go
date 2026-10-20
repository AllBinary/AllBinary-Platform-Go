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

type Initializer struct {
    LogUtil *LogUtil
    UserJdbcDriver string
    UserName string
    UserUserName string
    UserPassword string
    UserSchema string
    UserServer string
    UserPort string
    HistoryJdbcDriver string
    HistoryName string
    HistoryUserName string
    HistoryPassword string
    HistorySchema string
    HistoryServer string
    HistoryPort string
    LogJdbcDriver string
    LogName string
    LogUserName string
    LogPassword string
    LogSchema string
    LogServer string
    LogPort string
    InventoryJdbcDriver string
    InventoryName string
    InventoryUserName string
    InventoryPassword string
    InventorySchema string
    InventoryServer string
    InventoryPort string
    StaticPagesJdbcDriver string
    StaticPagesName string
    StaticPagesUserName string
    StaticPagesPassword string
    StaticPagesSchema string
    StaticPagesServer string
    StaticPagesPort string
}

var Initializer_MAXDB int = 30
var Initializer_MAX int = 16
var Initializer_MIN int = 4
var Initializer_MINPASSWORD int = 0

func (this *Initializer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewInitializer(map_ Map) *Initializer {
    this := &Initializer{}
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
        this.GetFormData(hashMap)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewInitializer(initHashMap *HashMap) *Initializer {
    this := &Initializer{}
    this.initFields()
        this.GetFormData(initHashMap)
    return this
}


func (this *Initializer) GetFormData(hashMap *HashMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Unable to get form data", this, "getFormData()", e)
                }
            }()
            var initializerData *InitializerData = InitializerData_GetInstance(); _ = initializerData
            this.UserJdbcDriver = any(hashMap.Get(initializerData.CUSTOMERJDBCDRIVER)).(string)
            this.UserName = any(hashMap.Get(initializerData.CUSTOMERNAME)).(string)
            this.UserUserName = any(hashMap.Get(initializerData.CUSTOMERUSERNAME)).(string)
            this.UserPassword = any(hashMap.Get(initializerData.CUSTOMERPASSWORD)).(string)
            this.UserSchema = any(hashMap.Get(initializerData.CUSTOMERSCHEMA)).(string)
            this.UserServer = any(hashMap.Get(initializerData.CUSTOMERSERVER)).(string)
            this.UserPort = any(hashMap.Get(initializerData.CUSTOMERPORT)).(string)
            this.HistoryJdbcDriver = any(hashMap.Get(initializerData.HISTORYJDBCDRIVER)).(string)
            this.HistoryName = any(hashMap.Get(initializerData.HISTORYNAME)).(string)
            this.HistoryUserName = any(hashMap.Get(initializerData.HISTORYUSERNAME)).(string)
            this.HistoryPassword = any(hashMap.Get(initializerData.HISTORYPASSWORD)).(string)
            this.HistorySchema = any(hashMap.Get(initializerData.HISTORYSCHEMA)).(string)
            this.HistoryServer = any(hashMap.Get(initializerData.HISTORYSERVER)).(string)
            this.HistoryPort = any(hashMap.Get(initializerData.HISTORYPORT)).(string)
            this.LogJdbcDriver = any(hashMap.Get(initializerData.LOGJDBCDRIVER)).(string)
            this.LogName = any(hashMap.Get(initializerData.LOGNAME)).(string)
            this.LogUserName = any(hashMap.Get(initializerData.LOGUSERNAME)).(string)
            this.LogPassword = any(hashMap.Get(initializerData.LOGPASSWORD)).(string)
            this.LogSchema = any(hashMap.Get(initializerData.LOGSCHEMA)).(string)
            this.LogServer = any(hashMap.Get(initializerData.LOGSERVER)).(string)
            this.LogPort = any(hashMap.Get(initializerData.LOGPORT)).(string)
            this.InventoryJdbcDriver = any(hashMap.Get(initializerData.INVENTORYJDBCDRIVER)).(string)
            this.InventoryName = any(hashMap.Get(initializerData.INVENTORYNAME)).(string)
            this.InventoryUserName = any(hashMap.Get(initializerData.INVENTORYUSERNAME)).(string)
            this.InventoryPassword = any(hashMap.Get(initializerData.INVENTORYPASSWORD)).(string)
            this.InventorySchema = any(hashMap.Get(initializerData.INVENTORYSCHEMA)).(string)
            this.InventoryServer = any(hashMap.Get(initializerData.INVENTORYSERVER)).(string)
            this.InventoryPort = any(hashMap.Get(initializerData.INVENTORYPORT)).(string)
            this.StaticPagesJdbcDriver = any(hashMap.Get(initializerData.STATICPAGESJDBCDRIVER)).(string)
            this.StaticPagesName = any(hashMap.Get(initializerData.STATICPAGESNAME)).(string)
            this.StaticPagesUserName = any(hashMap.Get(initializerData.STATICPAGESUSERNAME)).(string)
            this.StaticPagesPassword = any(hashMap.Get(initializerData.STATICPAGESPASSWORD)).(string)
            this.StaticPagesSchema = any(hashMap.Get(initializerData.STATICPAGESSCHEMA)).(string)
            this.StaticPagesServer = any(hashMap.Get(initializerData.STATICPAGESSERVER)).(string)
            this.StaticPagesPort = any(hashMap.Get(initializerData.STATICPAGESPORT)).(string)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *Initializer) IsJdbcDriverValid(jdbcDriverClassPathString string) bool {
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


func (this *Initializer) IsValid() bool {
        var isValid bool = true; _ = isValid
        if !this.IsJdbcDriverValid(this.UserJdbcDriver) {
            isValid = false
        }
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.UserName, Initializer_MIN, Initializer_MAXDB) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.UserUserName, Initializer_MIN, Initializer_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.UserPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            isValid = false
        }
        if !this.IsJdbcDriverValid(this.HistoryJdbcDriver) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.HistoryName, Initializer_MIN, Initializer_MAXDB) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.HistoryUserName, Initializer_MIN, Initializer_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.HistoryPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            isValid = false
        }
        if !this.IsJdbcDriverValid(this.LogJdbcDriver) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.LogName, Initializer_MIN, Initializer_MAXDB) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.LogUserName, Initializer_MIN, Initializer_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.LogPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            isValid = false
        }
        if !this.IsJdbcDriverValid(this.InventoryJdbcDriver) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.InventoryName, Initializer_MIN, Initializer_MAXDB) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.InventoryUserName, Initializer_MIN, Initializer_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.InventoryPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            isValid = false
        }
        if !this.IsJdbcDriverValid(this.StaticPagesJdbcDriver) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.StaticPagesName, Initializer_MIN, Initializer_MAXDB) {
            isValid = false
        }
        if !stringValidationUtil.IsValidRequired(this.StaticPagesUserName, Initializer_MIN, Initializer_MAX) {
            isValid = false
        }
        if !stringValidationUtil.IsValidNotRequired(this.StaticPagesPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            isValid = false
        }
        return isValid
}


func (this *Initializer) GetJdbcDriverValidationInfo(jdbcDriver string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("The JDBC driver (")
        stringBuffer.Append(jdbcDriver)
        stringBuffer.Append(") you have provided is not valid.<br/>")
        return fmt.Sprint(stringBuffer)
}


func (this *Initializer) GetJdbcDriverSolutionInfo() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("The following describes the possible problems and solutions regarding the Jdbc Driver error(s):<p/>")
        stringBuffer.Append("1. A JDBC driver you specified is not in your classpath.<br/>")
        stringBuffer.Append("Solution 1: Move the JDBC driver into any directory specified in the existing classpath. <br/>")
        stringBuffer.Append("Solution 2: Add the directory that contains the JDBC driver to the classpath.<br/>")
        stringBuffer.Append("Solution 3: Add the JDBC driver to the WEB-INF/lib directory where you installed this webapp.<br/>")
        stringBuffer.Append("2. The JDBC driver you specified does not exit.<br/>")
        stringBuffer.Append("Solution: Get a JDBC driver. <br/>")
        stringBuffer.Append("3. The JDBC driver you specified is not valid.<br/>")
        stringBuffer.Append("Solution: Use a valid JDBC driver.<p/>")
        return fmt.Sprint(stringBuffer)
}


func (this *Initializer) GetInvalidInfo() string {
        var isJdbcDriversValid bool = true; _ = isJdbcDriversValid
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if !this.IsJdbcDriverValid(this.UserJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.UserJdbcDriver))
        }
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if !stringValidationUtil.IsValidRequired(this.UserName, Initializer_MIN, Initializer_MAXDB) {
            stringBuffer.Append("User db name should be < " + fmt.Sprint(Initializer_MAXDB) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidRequired(this.UserUserName, Initializer_MIN, Initializer_MAX) {
            stringBuffer.Append("User db username should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.UserPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            stringBuffer.Append("User DB password should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MINPASSWORD) + " characters in length.<br />")
        }
        if !this.IsJdbcDriverValid(this.HistoryJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.HistoryJdbcDriver))
        }
        if !stringValidationUtil.IsValidRequired(this.HistoryName, Initializer_MIN, Initializer_MAXDB) {
            stringBuffer.Append("History db should be < " + fmt.Sprint(Initializer_MAXDB) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidRequired(this.HistoryUserName, Initializer_MIN, Initializer_MAX) {
            stringBuffer.Append("History db username should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.HistoryPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            stringBuffer.Append("History db password should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MINPASSWORD) + " characters in length.<br />")
        }
        if !this.IsJdbcDriverValid(this.LogJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.LogJdbcDriver))
        }
        if !stringValidationUtil.IsValidRequired(this.LogName, Initializer_MIN, Initializer_MAXDB) {
            stringBuffer.Append("Log db should be < " + fmt.Sprint(Initializer_MAXDB) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidRequired(this.LogUserName, Initializer_MIN, Initializer_MAX) {
            stringBuffer.Append("Log db username should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.LogPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            stringBuffer.Append("Log db password should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MINPASSWORD) + " characters in length.<br />")
        }
        if !this.IsJdbcDriverValid(this.InventoryJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.InventoryJdbcDriver))
        }
        if !stringValidationUtil.IsValidRequired(this.InventoryName, Initializer_MIN, Initializer_MAXDB) {
            stringBuffer.Append("Inventory db should be < " + fmt.Sprint(Initializer_MAXDB) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidRequired(this.InventoryUserName, Initializer_MIN, Initializer_MAX) {
            stringBuffer.Append("Inventory db username should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.InventoryPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            stringBuffer.Append("Inventory DB password should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MINPASSWORD) + " characters in length.<br />")
        }
        if !this.IsJdbcDriverValid(this.StaticPagesJdbcDriver) {
            isJdbcDriversValid = false
            stringBuffer.Append(this.GetJdbcDriverValidationInfo(this.StaticPagesJdbcDriver))
        }
        if !stringValidationUtil.IsValidRequired(this.StaticPagesName, Initializer_MIN, Initializer_MAXDB) {
            stringBuffer.Append("Static Pages db should be < " + fmt.Sprint(Initializer_MAXDB) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidRequired(this.StaticPagesUserName, Initializer_MIN, Initializer_MAX) {
            stringBuffer.Append("Static Pages db username should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MIN) + " characters in length.<br />")
        }
        if !stringValidationUtil.IsValidNotRequired(this.StaticPagesPassword, Initializer_MINPASSWORD, Initializer_MAX) {
            stringBuffer.Append("Static Pages DB password should be < " + fmt.Sprint(Initializer_MAX) + " and > " + fmt.Sprint(Initializer_MINPASSWORD) + " characters in length.<br />")
        }
        if !isJdbcDriversValid {
            stringBuffer.Append(this.GetJdbcDriverSolutionInfo())
        }
        return fmt.Sprint(stringBuffer)
}


// throws Exception
func (this *Initializer) Set() {
        this.LogUtil.PutF("Creating DB connection files", this, "set()")
        var userDbInitInfo *UserDbInitInfo = NewUserDbInitInfo(false); _ = userDbInitInfo
        userDbInitInfo.SetJdbcDriver(this.UserJdbcDriver)
        userDbInitInfo.SetName(this.UserName)
        userDbInitInfo.SetUserName(this.UserUserName)
        userDbInitInfo.SetPassword(this.UserPassword)
        userDbInitInfo.SetSchema(this.UserSchema)
        userDbInitInfo.SetServer(this.UserServer)
        userDbInitInfo.SetPort(this.UserPort)
        userDbInitInfo.SetHasRead(true)
        userDbInitInfo.Write()
        var historyDbInitInfo *HistoryDbInitInfo = NewHistoryDbInitInfo(false); _ = historyDbInitInfo
        historyDbInitInfo.SetJdbcDriver(this.HistoryJdbcDriver)
        historyDbInitInfo.SetName(this.HistoryName)
        historyDbInitInfo.SetUserName(this.HistoryUserName)
        historyDbInitInfo.SetPassword(this.HistoryPassword)
        historyDbInitInfo.SetSchema(this.HistorySchema)
        historyDbInitInfo.SetServer(this.HistoryServer)
        historyDbInitInfo.SetPort(this.HistoryPort)
        historyDbInitInfo.SetHasRead(true)
        historyDbInitInfo.Write()
        var logDbInitInfo *LogDbInitInfo = NewLogDbInitInfo(false); _ = logDbInitInfo
        logDbInitInfo.SetJdbcDriver(this.LogJdbcDriver)
        logDbInitInfo.SetName(this.LogName)
        logDbInitInfo.SetUserName(this.LogUserName)
        logDbInitInfo.SetPassword(this.LogPassword)
        logDbInitInfo.SetSchema(this.LogSchema)
        logDbInitInfo.SetServer(this.LogServer)
        logDbInitInfo.SetPort(this.LogPort)
        logDbInitInfo.SetHasRead(true)
        logDbInitInfo.Write()
        var staticPagesDbInitInfo *StaticPagesDbInitInfo = NewStaticPagesDbInitInfo(false); _ = staticPagesDbInitInfo
        staticPagesDbInitInfo.SetJdbcDriver(this.StaticPagesJdbcDriver)
        staticPagesDbInitInfo.SetName(this.StaticPagesName)
        staticPagesDbInitInfo.SetUserName(this.StaticPagesUserName)
        staticPagesDbInitInfo.SetPassword(this.StaticPagesPassword)
        staticPagesDbInitInfo.SetSchema(this.StaticPagesSchema)
        staticPagesDbInitInfo.SetServer(this.StaticPagesServer)
        staticPagesDbInitInfo.SetPort(this.StaticPagesPort)
        staticPagesDbInitInfo.SetHasRead(true)
        staticPagesDbInitInfo.Write()
        var inventoryDbInitInfo *InventoryDbInitInfo = NewInventoryDbInitInfo(false); _ = inventoryDbInitInfo
        inventoryDbInitInfo.SetJdbcDriver(this.InventoryJdbcDriver)
        inventoryDbInitInfo.SetName(this.InventoryName)
        inventoryDbInitInfo.SetUserName(this.InventoryUserName)
        inventoryDbInitInfo.SetPassword(this.InventoryPassword)
        inventoryDbInitInfo.SetSchema(this.InventorySchema)
        inventoryDbInitInfo.SetServer(this.InventoryServer)
        inventoryDbInitInfo.SetPort(this.InventoryPort)
        inventoryDbInitInfo.SetHasRead(true)
        inventoryDbInitInfo.Write()
        this.LogUtil.PutF("Created DB connection files", this, "set()")
}

