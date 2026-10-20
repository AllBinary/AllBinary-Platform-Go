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

type StreetAddressData struct {
}

var StreetAddressData_MAX int = 255
var StreetAddressData_MIN int = 0
var StreetAddressData_FORM string = "FORM"
var StreetAddressData_STATES string = "STATES"
var StreetAddressData_NUMBEROFADDRESSES string = "STREETADDRESS_NUMBEROFADDRESSES"
var StreetAddressData_ID string = "STREETADDRESS_ID"
var StreetAddressData_DEFAULT string = "STREETADDRESS_DEFAULT_ADDRESS"
var StreetAddressData_ADDRESS string = "STREETADDRESS_ADDRESS"
var StreetAddressData_NAME string = "STREETADDRESS_NAME"
var StreetAddressData_STREET string = "STREETADDRESS_STREET"
var StreetAddressData_CITY string = "STREETADDRESS_CITY"
var StreetAddressData_STATE string = "STREETADDRESS_STATE"
var StreetAddressData_CODE string = "STREETADDRESS_CODE"
var StreetAddressData_COUNTRY string = "STREETADDRESS_COUNTRY"


func NewStreetAddressData() *StreetAddressData {
    this := &StreetAddressData{}
    return this
}

