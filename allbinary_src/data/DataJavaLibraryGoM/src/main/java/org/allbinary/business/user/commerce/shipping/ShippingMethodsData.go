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

type ShippingMethodsData struct {
}

var ShippingMethodsData_SHIPPINGMETHODSORDERSUMMARIES string = "SHIPPINGMETHODS_ORDERSUMMARIES"
var ShippingMethodsData_SHIPPINGMETHODORDERSUMMARY string = "SHIPPINGMETHOD_ORDERSUMMARY"
var ShippingMethodsData_SHIPPINGMETHODS string = "SHIPPINGMETHODS"
var ShippingMethodsData_NAME string = "SHIPPINGMETHODS_NAME"
var ShippingMethodsData_SHIPPINGCOST string = "SHIPPINGMETHOD_SHIPPINGCOST"
var ShippingMethodsData_TAXRATE string = "SHIPPINGMETHOD_TAXRATE"
var ShippingMethodsData_SUBTOTAL string = "SHIPPINGMETHOD_SUBTOTAL"
var ShippingMethodsData_TAX string = "SHIPPINGMETHOD_TAX"
var ShippingMethodsData_TOTAL string = "SHIPPINGMETHOD_TOTAL"

func NewShippingMethodsData() *ShippingMethodsData {
    this := &ShippingMethodsData{}
    return this
}

