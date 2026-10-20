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

type AuthorizeNetAimData struct {
}

var AuthorizeNetAimData_X_DELIM_DATA string = "x_delim_data"
var AuthorizeNetAimData_X_DELIM_CHAR string = "x_delim_char"
var AuthorizeNetAimData_X_ENCAP_CHAR string = "x_encap_char"
var AuthorizeNetAimData_X_AUTHENTICATION_INDICATOR string = "x_authentication_indicator"
var AuthorizeNetAimData_X_CARDHOLDER_AUTHENTICATION_VALUE string = "x_cardholder_authentication_value"
var AuthorizeNetAimData_X_DRIVERS_LICENSE_NUM string = "x_drivers_license_num"
var AuthorizeNetAimData_X_DRIVERS_LICENSE_STATE string = "x_drivers_license_state"
var AuthorizeNetAimData_X_DRIVERS_LICENSE_DOB string = "x_drivers_license_dob"


func NewAuthorizeNetAimData() *AuthorizeNetAimData {
    this := &AuthorizeNetAimData{}
    return this
}

