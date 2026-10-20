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

type UserEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
    COMMAND_SUCCESS_FOR_USER string
    PASSWORD_LABEL string
    EQUALS string
    COMMAND_SUCCESS_BUT_LOGIN_FAILED string
    INVALID_PASSWORD_LABEL string
    NOT_EQUAL string
    END_QUOTES string
}

// Implements UserEntityInterface


func (this *UserEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "user"
    this.COMMAND_SUCCESS_FOR_USER = "Command Success for user: "
    this.PASSWORD_LABEL = " Password: "
    this.EQUALS = "=="
    this.COMMAND_SUCCESS_BUT_LOGIN_FAILED = "Command Success but login failed for user: "
    this.INVALID_PASSWORD_LABEL = " Password: \n\""
    this.NOT_EQUAL = "\"!=\""
    this.END_QUOTES = CommonSeps_GetInstance().QUOTE
}


func NewUserEntity() *UserEntity {
    this := &UserEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *UserEntity) Insert(values *BasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, AbSqlBean_INSERT, e)
            }
                }
            }()
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *UserEntity) GetAdministrators() *BasicArrayList {
        return this.GetUsersWithRole(UserRoleFactory_GetInstance().ADMINISTRATOR)
}


// throws Exception
func (this *UserEntity) GetStoreManagers(storeFrontInterface StoreFrontInterface) *BasicArrayList {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        var usersVector *BasicArrayList = NewBasicArrayListD(); _ = usersVector
        keysAndValues.Put(fmt.Sprint(UserRoleData_NAME), fmt.Sprint(UserRoleFactory_GetInstance().STOREMANAGER))
        keysAndValues.Put(UserData_PERMISSIONS, storeFrontInterface.GetName())
        var usersHashMapVector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = usersHashMapVector
        var size int = usersHashMapVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var userHashMap *HashMap = any(usersHashMapVector.Get(i)).(*HashMap); _ = userHashMap
            if userHashMap != nil {
                usersVector.Add(NewUser(userHashMap))
            }
        }
        return usersVector
}


// throws Exception
func (this *UserEntity) GetCustomers() *BasicArrayList {
        return this.GetUsersWithRole(UserRoleFactory_GetInstance().CUSTOMER)
}


// throws Exception
func (this *UserEntity) GetUsersWithRole(userRole *UserRole) *BasicArrayList {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        var usersVector *BasicArrayList = NewBasicArrayListD(); _ = usersVector
        keysAndValues.Put(fmt.Sprint(UserRoleData_NAME), fmt.Sprint(userRole))
        var usersHashMapVector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = usersHashMapVector
        var size int = usersHashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userHashMap *HashMap = any(usersHashMapVector.Get(index)).(*HashMap); _ = userHashMap
            if userHashMap != nil {
                usersVector.Add(NewUser(userHashMap))
            }
        }
        return usersVector
}


// throws Exception
func (this *UserEntity) GetUsers(storeFrontInterface StoreFrontInterface) *BasicArrayList {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        var usersVector *BasicArrayList = NewBasicArrayListD(); _ = usersVector
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeFrontInterface.GetName())
        var usersHashMapVector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = usersHashMapVector
        var size int = usersHashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userHashMap *HashMap = any(usersHashMapVector.Get(index)).(*HashMap); _ = userHashMap
            if userHashMap != nil {
                usersVector.Add(NewUser(userHashMap))
            }
        }
        return usersVector
}


// throws Exception
func (this *UserEntity) GetUser(userName string) UserInterface {
        var row *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = row
        row.Put(UserData_USERNAME, userName)
        var userHashMap *HashMap = this.AbSqlBean.GetRow(row); _ = userHashMap
        if userHashMap != nil {
            return CreateUserFactory_GetInstance(userHashMap)
        } else {
            return nil
        }
}


func (this *UserEntity) DeleteWhere(key string, value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "deleteWhere", e)
            }
                }
            }()
            this.AbSqlBean.DeleteWhere(key, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "deleteWhere")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *UserEntity) Login(userName string, password string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "login", e)
            }
            tryControl1, tryResult1 = 1, "Error"; return
                }
            }()
            var isUserNameAndPasswordCorrect int = 0; _ = isUserNameAndPasswordCorrect
            var result string = this.AbSqlBean.GetField(UserData_USERNAME, userName, UserData_PASSWORD); _ = result
            var encryption string = this.AbSqlBean.GetField(UserData_USERNAME, userName, EntryData_GetInstance().ENCRYPTION); _ = encryption
            if encryption != "" && strings.Compare(encryption, this.StringUtil.EMPTY_STRING) != 0 {
                isUserNameAndPasswordCorrect = strings.Compare(result, NewSuperCrypt(int(int(encryption))).Encrypt(password))
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF("Command Success but login failed for user: " + userName + " because user did not exist", this, "login")
                }
                tryControl1, tryResult1 = 1, AbSqlBean_GLOBALS2.LOGINFAILED; return
            }
            if isUserNameAndPasswordCorrect == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(this.COMMAND_SUCCESS_FOR_USER)
                    stringBuffer.Append(userName)
                    stringBuffer.Append(this.PASSWORD_LABEL)
                    stringBuffer.Append(password)
                    stringBuffer.Append(this.EQUALS)
                    stringBuffer.Append(result)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "login")
                }
                tryControl1, tryResult1 = 1, AbSqlBean_GLOBALS2.LOGINSUCCESS; return
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(this.COMMAND_SUCCESS_BUT_LOGIN_FAILED)
                    stringBuffer.Append(userName)
                    stringBuffer.Append(this.INVALID_PASSWORD_LABEL)
                    stringBuffer.Append(NewSuperCrypt(int(int(encryption))).Encrypt(password))
                    stringBuffer.Append(this.NOT_EQUAL)
                    stringBuffer.Append(result)
                    stringBuffer.Append(this.END_QUOTES)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "login")
                }
                tryControl1, tryResult1 = 1, AbSqlBean_GLOBALS2.LOGINFAILED; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *UserEntity) Update(userName string, updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(UserData_USERNAME, userName, updatedValues)
}


func (this *UserEntity) DropTables() string {
        return this.AbSqlBean.DropTable()
}


func (this *UserEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(UserData_USERNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(UserData_PREFIXNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_FIRSTNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_LASTNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_MIDDLENAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_SUFFIXNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_COMPANY).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_POSITIONATCOMPANY).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_MAINEMAIL).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(UserData_SECONDARYEMAIL).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_HOMEPHONE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_CELLPHONE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_WORKPHONE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_OTHERCONTACT).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_ELECTRONICDEVICE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_FAX).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(fmt.Sprint(UserRoleData_NAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(UserData_CONFIGURATION).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(UserData_PERMISSIONS).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(entryData.ENCRYPTION).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(UserData_SECRET).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(UserData_PASSWORD).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(entryData.ENABLE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(entryData.TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(entryData.LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(UserData_USERNAME).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *UserEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}

