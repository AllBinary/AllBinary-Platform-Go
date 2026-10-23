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

type UserInterface interface {
    TableMappingInterface
    ValidationInterface
    SetUserName(value string)
    SetPrefixName(value string)
    SetFirstName(value string)
    SetLastName(value string)
    SetMiddleName(value string)
    SetSuffixName(value string)
    SetCompany(value string)
    SetPositionAtCompany(value string)
    SetMainEmail(value string)
    SetSecondaryEmail(value string)
    SetHomePhone(value string)
    SetCellPhone(value string)
    SetWorkPhone(value string)
    SetOtherContact(value string)
    SetElectronicContact(value string)
    SetFax(value string)
    SetRole(aRole IUserRole)
    SetEncryption(value string)
    SetSecret(value string)
    SetPassword(value string)
    SetPermissions(value string)
    Enable()
    SetEnable(enable string)
    GetUserName() string
    GetPrefixName() string
    GetFirstName() string
    GetLastName() string
    GetMiddleName() string
    GetSuffixName() string
    GetCompany() string
    GetPositionAtCompany() string
    GetMainEmail() string
    GetSecondaryEmail() string
    GetHomePhone() string
    GetCellPhone() string
    GetWorkPhone() string
    GetOtherContact() string
    GetElectronicContact() string
    GetFax() string
    GetRole() IUserRole
    GetEncryption() string
    GetSecret() string
    GetPassword() string
    GetPermissions() string
    GetUserConfigurationInterface() UserConfigurationInterface
    SetUserConfigurationInterface(userConfigurationInterface UserConfigurationInterface)
    ToPasswordHashMap() *HashMap
    ValidateSession(weblisketSession WeblisketSessionInterface)
    UpdateSession(weblisketSession WeblisketSessionInterface)
    IsSessionValid() bool
}


