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

type OrderData struct {
}

var OrderData_ORDERS string = "ORDERS"
var OrderData_ORDER string = "ORDER"
var OrderData_ID string = "ORDER_ID"
var OrderData_CUSTOMERCOMMENT string = "ORDER_CUSTOMER_COMMENT"
var OrderData_CUSTOMERCANCELCOMMENT string = "ORDER_CUSTOMER_CANCEL_COMMENT"
var OrderData_STORECOMMENT string = "ORDER_STORE_COMMENT"
var OrderData_STORECANCELCOMMENT string = "ORDER_STORE_CANCEL_COMMENT"

func NewOrderData() *OrderData {
    this := &OrderData{}
    return this
}

