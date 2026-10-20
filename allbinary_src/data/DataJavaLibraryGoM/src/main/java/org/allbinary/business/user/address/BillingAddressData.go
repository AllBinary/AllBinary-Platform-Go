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

type BillingAddressData struct {
}

var BillingAddressData_MAX int = 255
var BillingAddressData_MIN int = 0
var BillingAddressData_MULTIPLE string = "BILLINGADDRESSES"
var BillingAddressData_STATES string = "BILLINGADDRESS_STATES"
var BillingAddressData_SETTOSHIPPINGADDRESS string = "Set To Shipping Address"
var BillingAddressData_SELECT string = "Select Billing Address"
var BillingAddressData_SETTOSHIPPINGADDRESSNAME string = "SETTOSHIPPINGADDRESS"
var BillingAddressData_BILLINGADDRESSES string = "BILLINGADDRESSES"
var BillingAddressData_BILLINGADDRESS string = "BILLINGADDRESS"
var BillingAddressData_ID string = "BILLINGADDRESS_ID"
var BillingAddressData_ADDRESS string = "BILLINGADDRESS_ADDRESS"
var BillingAddressData_NAME string = "BILLINGADDRESS_NAME"
var BillingAddressData_STREET string = "BILLINGADDRESS_STREET"
var BillingAddressData_CITY string = "BILLINGADDRESS_CITY"
var BillingAddressData_STATE string = "BILLINGADDRESS_STATE"
var BillingAddressData_CODE string = "BILLINGADDRESS_CODE"
var BillingAddressData_COUNTRY string = "BILLINGADDRESS_COUNTRY"


func NewBillingAddressData() *BillingAddressData {
    this := &BillingAddressData{}
    return this
}

