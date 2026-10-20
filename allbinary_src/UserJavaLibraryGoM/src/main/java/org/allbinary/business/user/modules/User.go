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

type User struct {
    LogUtil *LogUtil
    UserName string
    PrefixName string
    FirstName string
    LastName string
    MiddleName string
    SuffixName string
    Company string
    PositionAtCompany string
    MainEmail string
    SecondaryEmail string
    HomePhone string
    CellPhone string
    WorkPhone string
    OtherContact string
    ElectronicDevice string
    Fax string
    Role *UserRole
    Encryption string
    Secret string
    Password *Password
    Enable string
    Permissions string
    UserConfigurationInterface UserConfigurationInterface
}

// Implements UserInterface


func (this *User) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewUser() *User {
    this := &User{}
    this.initFields()
        this.UserConfigurationInterface = UserConfigurationInterfaceFactory_GetInstance()
        this.Password = NewPassword(StringUtil_GetInstance().EMPTY_STRING)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUser(request *HttpServletRequest) *User {
    this := &User{}
    this.initFields()
        this.GetFormData(NewRequestParams(request).ToHashMap())
        this.UserConfigurationInterface = UserConfigurationInterfaceFactory_GetInstance(this.GetRole())
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUser(userHashMap *HashMap) *User {
    this := &User{}
    this.initFields()
        this.GetFormData(userHashMap)
        this.UserConfigurationInterface = UserConfigurationInterfaceFactory_GetInstance(this.GetRole())
    return this
}


// throws Exception
func (this *User) GetFormData(userHashMap *HashMap) {
        this.UserName = NewUserName(userHashMap).Get()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var passwordString string = stringUtil.GetNonNull(any(userHashMap.Get(UserData_PASSWORD)).(string)); _ = passwordString
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(this.UserName) && stringValidationUtil.IsEmpty(passwordString) {
            this.UserName = stringUtil.GetNonNull(any(userHashMap.Get(WeblisketSessionData_REMOVABLEUSERNAME)).(string))
            passwordString = stringUtil.GetNonNull(any(userHashMap.Get(WeblisketSessionData_REMOVABLEPASSWORD)).(string))
        }
        var encryption string = stringUtil.GetNonNull(any(userHashMap.Get(EntryData_GetInstance().ENCRYPTION)).(string)); _ = encryption
        var secret string = stringUtil.GetNonNull(any(userHashMap.Get(UserData_SECRET)).(string)); _ = secret
        this.Password = NewPassword(passwordString)
        this.PrefixName = stringUtil.GetNonNull(any(userHashMap.Get(UserData_PREFIXNAME)).(string))
        this.FirstName = stringUtil.GetNonNull(any(userHashMap.Get(UserData_FIRSTNAME)).(string))
        this.LastName = stringUtil.GetNonNull(any(userHashMap.Get(UserData_LASTNAME)).(string))
        this.MiddleName = stringUtil.GetNonNull(any(userHashMap.Get(UserData_MIDDLENAME)).(string))
        this.SuffixName = stringUtil.GetNonNull(any(userHashMap.Get(UserData_SUFFIXNAME)).(string))
        this.Company = stringUtil.GetNonNull(any(userHashMap.Get(UserData_COMPANY)).(string))
        this.PositionAtCompany = stringUtil.GetNonNull(any(userHashMap.Get(UserData_POSITIONATCOMPANY)).(string))
        this.MainEmail = stringUtil.GetNonNull(any(userHashMap.Get(UserData_MAINEMAIL)).(string))
        this.SecondaryEmail = stringUtil.GetNonNull(any(userHashMap.Get(UserData_SECONDARYEMAIL)).(string))
        this.HomePhone = stringUtil.GetNonNull(any(userHashMap.Get(UserData_HOMEPHONE)).(string))
        this.CellPhone = stringUtil.GetNonNull(any(userHashMap.Get(UserData_CELLPHONE)).(string))
        this.WorkPhone = stringUtil.GetNonNull(any(userHashMap.Get(UserData_WORKPHONE)).(string))
        this.OtherContact = stringUtil.GetNonNull(any(userHashMap.Get(UserData_OTHERCONTACT)).(string))
        this.ElectronicDevice = stringUtil.GetNonNull(any(userHashMap.Get(UserData_ELECTRONICDEVICE)).(string))
        this.Fax = stringUtil.GetNonNull(any(userHashMap.Get(UserData_FAX)).(string))
        this.Role = UserRoleB_GetRole(any(userHashMap.Get(fmt.Sprint(UserRoleData_NAME))).(string))
        this.Permissions = stringUtil.GetNonNull(any(userHashMap.Get(UserData_PERMISSIONS)).(string))
        if !stringValidationUtil.IsEmpty(this.Permissions) && strings.Compare(this.Permissions, StoreFrontData_GetInstance().NAME) == 0 {
            this.Permissions = stringUtil.GetNonNull(any(userHashMap.Get(StoreFrontData_GetInstance().NAME)).(string))
        } else if this.Permissions == "" {
            this.Permissions = "No Permissions"
        }
        this.Enable = stringUtil.GetNonNull(any(userHashMap.Get(EntryData_GetInstance().ENABLE)).(string))
}


func (this *User) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATIONERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put("Failed to validate form", this, commonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if !UserName_GetInstance().IsValid(this.UserName) {
                valid = false
            }
            if !this.Password.IsValid() {
                valid = false
            }
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if !stringValidationUtil.IsValidRequired(this.FirstName, 1, UserData_MAXLEN) {
                valid = false
            }
            if !stringValidationUtil.IsValidRequired(this.LastName, 1, UserData_MAXLEN) {
                valid = false
            }
            if !stringValidationUtil.IsValidRequired(this.MainEmail, 1, UserData_MAXLEN) || strings.Index(this.MainEmail, string("@")) == -1 {
                valid = false
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *User) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(UserName_GetValidationInfo(this.UserName))
            stringBuffer.Append(this.Password.GetValidationInfo())
            if !stringValidationUtil.IsValidRequired(this.FirstName, 1, UserData_MAXLEN) {
                stringBuffer.Append("Please enter a valid First Name.<br />")
            }
            if !stringValidationUtil.IsValidRequired(this.LastName, 1, UserData_MAXLEN) {
                stringBuffer.Append("Please enter a valid Last Name.<br />")
            }
            if !stringValidationUtil.IsValidRequired(this.MainEmail, 1, UserData_MAXLEN) || strings.Index(this.MainEmail, string("@")) == -1 {
                stringBuffer.Append("Please enter a valid email address.<br />")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *User) SetUserName(value string) {
        this.UserName = value
}


func (this *User) SetPrefixName(value string) {
        this.PrefixName = value
}


func (this *User) SetFirstName(value string) {
        this.FirstName = value
}


func (this *User) SetLastName(value string) {
        this.LastName = value
}


func (this *User) SetMiddleName(value string) {
        this.MiddleName = value
}


func (this *User) SetSuffixName(value string) {
        this.SuffixName = value
}


func (this *User) SetCompany(value string) {
        this.Company = value
}


func (this *User) SetPositionAtCompany(value string) {
        this.PositionAtCompany = value
}


func (this *User) SetMainEmail(value string) {
        this.MainEmail = value
}


func (this *User) SetSecondaryEmail(value string) {
        this.SecondaryEmail = value
}


func (this *User) SetHomePhone(value string) {
        this.HomePhone = value
}


func (this *User) SetCellPhone(value string) {
        this.CellPhone = value
}


func (this *User) SetWorkPhone(value string) {
        this.WorkPhone = value
}


func (this *User) SetOtherContact(value string) {
        this.OtherContact = value
}


func (this *User) SetElectronicContact(value string) {
        this.ElectronicDevice = value
}


func (this *User) SetFax(value string) {
        this.Fax = value
}


func (this *User) SetRole(role *UserRole) {
        this.Role = role
}


func (this *User) SetPermissions(value string) {
        this.Permissions = value
}


func (this *User) SetEncryption(value string) {
        this.Encryption = value
}


func (this *User) SetSecret(value string) {
        this.Secret = value
}


func (this *User) SetPassword(value string) {
        this.Password.Set(value)
}


func (this *User) Enable() {
        this.SetEnable("Yes")
}


func (this *User) Disable() {
        this.SetEnable("No")
}


func (this *User) SetEnable(enable string) {
        this.Enable = enable
}


func (this *User) GetUserName() string {
        return this.UserName
}


func (this *User) GetPrefixName() string {
        return this.PrefixName
}


func (this *User) GetFirstName() string {
        return this.FirstName
}


func (this *User) GetLastName() string {
        return this.LastName
}


func (this *User) GetMiddleName() string {
        return this.MiddleName
}


func (this *User) GetSuffixName() string {
        return this.SuffixName
}


func (this *User) GetCompany() string {
        return this.Company
}


func (this *User) GetPositionAtCompany() string {
        return this.PositionAtCompany
}


func (this *User) GetMainEmail() string {
        return this.MainEmail
}


func (this *User) GetSecondaryEmail() string {
        return this.SecondaryEmail
}


func (this *User) GetHomePhone() string {
        return this.HomePhone
}


func (this *User) GetCellPhone() string {
        return this.CellPhone
}


func (this *User) GetWorkPhone() string {
        return this.WorkPhone
}


func (this *User) GetOtherContact() string {
        return this.OtherContact
}


func (this *User) GetElectronicContact() string {
        return this.ElectronicDevice
}


func (this *User) GetFax() string {
        return this.Fax
}


func (this *User) GetRole() *UserRole {
        return this.Role
}


func (this *User) GetPermissions() string {
        return this.Permissions
}


func (this *User) GetEncryption() string {
        return this.Encryption
}


func (this *User) GetSecret() string {
        return this.Secret
}


func (this *User) GetPassword() string {
        return this.Password.Get()
}


// throws Exception
func (this *User) ToVector() *BasicArrayList {
        var values *BasicArrayList = NewBasicArrayListD(); _ = values
        values.Add(this.UserName)
        values.Add(this.PrefixName)
        values.Add(this.FirstName)
        values.Add(this.LastName)
        values.Add(this.MiddleName)
        values.Add(this.SuffixName)
        values.Add(this.Company)
        values.Add(this.PositionAtCompany)
        values.Add(this.MainEmail)
        values.Add(this.SecondaryEmail)
        values.Add(this.HomePhone)
        values.Add(this.CellPhone)
        values.Add(this.WorkPhone)
        values.Add(this.OtherContact)
        values.Add(this.ElectronicDevice)
        values.Add(this.Fax)
        values.Add(fmt.Sprint(this.GetRole()))
        var userConfigurationDomDocumentMapping *UserConfigurationDomDocumentMapping = NewUserConfigurationDomDocumentMapping(this.GetUserConfigurationInterface()); _ = userConfigurationDomDocumentMapping
        values.Add(userConfigurationDomDocumentMapping.ToDomDocumentString())
        values.Add(this.Permissions)
        values.AddAllList(this.Password.ToVector(this.Secret))
        values.Add(this.Enable)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Add(time)
        values.Add(time)
        return values
}


// throws Exception
func (this *User) ToHashMap() *HashMap {
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(UserData_USERNAME, this.UserName)
        values.Put(UserData_PREFIXNAME, this.PrefixName)
        values.Put(UserData_FIRSTNAME, this.FirstName)
        values.Put(UserData_LASTNAME, this.LastName)
        values.Put(UserData_MIDDLENAME, this.MiddleName)
        values.Put(UserData_SUFFIXNAME, this.SuffixName)
        values.Put(UserData_COMPANY, this.Company)
        values.Put(UserData_POSITIONATCOMPANY, this.PositionAtCompany)
        values.Put(UserData_MAINEMAIL, this.MainEmail)
        values.Put(UserData_SECONDARYEMAIL, this.SecondaryEmail)
        values.Put(UserData_HOMEPHONE, this.HomePhone)
        values.Put(UserData_CELLPHONE, this.CellPhone)
        values.Put(UserData_WORKPHONE, this.WorkPhone)
        values.Put(UserData_OTHERCONTACT, this.OtherContact)
        values.Put(UserData_ELECTRONICDEVICE, this.ElectronicDevice)
        values.Put(UserData_FAX, this.Fax)
        if this.GetRole() != nil {
            values.Put(fmt.Sprint(UserRoleData_NAME), fmt.Sprint(this.GetRole()))
        }
        var userConfigurationDomDocumentMapping *UserConfigurationDomDocumentMapping = NewUserConfigurationDomDocumentMapping(this.GetUserConfigurationInterface()); _ = userConfigurationDomDocumentMapping
        values.Put(UserData_CONFIGURATION, userConfigurationDomDocumentMapping.ToDomDocumentString())
        values.Put(UserData_PERMISSIONS, this.Permissions)
        values.PutAll(this.Password.ToHashMap(this.Secret))
        values.Put(EntryData_GetInstance().ENABLE, this.Enable)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *User) ToPasswordHashMap() *HashMap {
        return this.Password.ToHashMap(this.Secret)
}


func (this *User) ValidateSession(weblisketSession WeblisketSessionInterface) {
        weblisketSession.SetAuthenticated()
        weblisketSession.SetRole(this.GetRole())
        weblisketSession.SetUserName(this.GetUserName())
}


func (this *User) UpdateSession(weblisketSession WeblisketSessionInterface) {
}


func (this *User) IsSessionValid() bool {
        return true
}


// throws Exception
func (this *User) GetKey() any {
        return any(this.UserName).(any)
}


func (this *User) ToValidationInfoDoc() *Document {
        return nil
}


func (this *User) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *User) GetUserConfigurationInterface() UserConfigurationInterface {
        return this.UserConfigurationInterface
}


func (this *User) SetUserConfigurationInterface(userConfigurationInterface UserConfigurationInterface) {
        this.UserConfigurationInterface = userConfigurationInterface
}

