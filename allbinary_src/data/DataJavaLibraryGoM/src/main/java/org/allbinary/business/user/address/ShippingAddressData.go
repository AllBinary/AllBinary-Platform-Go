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

type ShippingAddressData struct {
}

var ShippingAddressData_SETTOBILLINGADDRESS string = "Set To Billing Address"
var ShippingAddressData_SELECT string = "Select Shipping Address"
var ShippingAddressData_SELECTNAME string = "SELECTSHIPPINGADDRESS"
var ShippingAddressData_MULTIPLE string = "SHIPPINGADDRESSES"
var ShippingAddressData_SHIPPINGADDRESSES string = "SHIPPINGADDRESSES"
var ShippingAddressData_SHIPPINGADDRESS string = "SHIPPINGADDRESS"
var ShippingAddressData_ID string = "SHIPPINGADDRESS_ID"
var ShippingAddressData_ADDRESS string = "SHIPPINGADDRESS_ADDRESS"
var ShippingAddressData_NAME string = "SHIPPINGADDRESS_NAME"
var ShippingAddressData_STREET string = "SHIPPINGADDRESS_STREET"
var ShippingAddressData_CITY string = "SHIPPINGADDRESS_CITY"
var ShippingAddressData_STATE string = "SHIPPINGADDRESS_STATE"
var ShippingAddressData_CODE string = "SHIPPINGADDRESS_CODE"
var ShippingAddressData_COUNTRY string = "SHIPPINGADDRESS_COUNTRY"
var ShippingAddressData_TAX string = "SHIPPINGADDRESS_TAX"


func NewShippingAddressData() *ShippingAddressData {
    this := &ShippingAddressData{}
    return this
}

