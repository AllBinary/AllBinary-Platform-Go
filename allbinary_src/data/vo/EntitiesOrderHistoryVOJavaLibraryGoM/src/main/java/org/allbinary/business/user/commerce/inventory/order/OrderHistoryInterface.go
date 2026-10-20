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

type OrderHistoryInterface interface {
    SetUserName(value string)
    SetDefaultShippingAddress(shippingAddress *StreetAddress)
    SetBillingAddress(billingAddress *StreetAddress)
    SetPaymentInfo(payment *Payment)
    SetShipDate(value string)
    SetOrderDate(value string)
    SetTransDate(value string)
    SetCancelDate(value string)
    SetStatus(value string)
    SetSubTotal(value *Money)
    SetShippingCost(value *Money)
    SetTaxes(value *Money)
    SetTotal(value *Money)
    GetUserName() string
    GetShippingAddress() *StreetAddress
    GetBillingAddress() *StreetAddress
    GetPaymentInfo() *Payment
    GetShipDate() string
    GetOrderDate() string
    GetTransDate() string
    GetCancelDate() string
    GetStatus() string
    GetSubTotal() *Money
    GetShippingCost() *Money
    GetTaxes() *Money
    GetTotal() *Money
    GetPaymentMethod() string
}


