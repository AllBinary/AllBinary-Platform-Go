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

type UserData struct {
}

var UserData_MAXLEN int = 250
var UserData_NAME string = "USER_NAME"
var UserData_USERNAME string = "CUSTOMER_USER_NAME"
var UserData_PASSWORD string = "CUSTOMER_PASSWORD"
var UserData_PREFIXNAME string = "CUSTOMER_PREFIX_NAME"
var UserData_FIRSTNAME string = "CUSTOMER_FIRST_NAME"
var UserData_LASTNAME string = "CUSTOMER_LAST_NAME"
var UserData_MIDDLENAME string = "CUSTOMER_MIDDLE_NAME"
var UserData_SUFFIXNAME string = "CUSTOMER_SUFFIX_NAME"
var UserData_COMPANY string = "CUSTOMER_COMPANY"
var UserData_POSITIONATCOMPANY string = "CUSTOMER_POSITION_AT_COMPANY"
var UserData_MAINEMAIL string = "CUSTOMER_MAIN_EMAIL"
var UserData_SECONDARYEMAIL string = "CUSTOMER_SECONDARY_EMAIL"
var UserData_HOMEPHONE string = "CUSTOMER_HOME_PHONE"
var UserData_CELLPHONE string = "CUSTOMER_CELL_PHONE"
var UserData_WORKPHONE string = "CUSTOMER_WORK_PHONE"
var UserData_OTHERCONTACT string = "CUSTOMER_OTHER_CONTACT"
var UserData_ELECTRONICDEVICE string = "CUSTOMER_ELECTRONIC_DEVICE"
var UserData_FAX string = "CUSTOMER_FAX"
var UserData_SECRET string = "CUSTOMER_SECRET"
var UserData_PERMISSIONS string = "CUSTOMER_PERMISSIONS"
var UserData_CONFIGURATION string = "CUSTOMER_CONFIGURATION"


func NewUserData() *UserData {
    this := &UserData{}
    return this
}

