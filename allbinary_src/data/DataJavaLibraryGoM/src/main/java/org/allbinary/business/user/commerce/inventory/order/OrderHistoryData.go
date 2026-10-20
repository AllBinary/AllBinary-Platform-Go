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

type OrderHistoryData struct {
}

var OrderHistoryData_MAXIDLEN int = 19
var OrderHistoryData_DATEDELTA string = "DATE_DELTA"
var OrderHistoryData_DATETYPE string = "DATE_TYPE"
var OrderHistoryData_TYPELONG string = "TYPE_LONG"
var OrderHistoryData_TYPECAESAR string = "TYPE_CAESAR"
var OrderHistoryData_TODATE string = "TO_DATE"
var OrderHistoryData_FROMHOUR string = "FROM_HOUR"
var OrderHistoryData_FROMDAY string = "FROM_DAY"
var OrderHistoryData_FROMMONTH string = "FROM_MONTH"
var OrderHistoryData_FROMYEAR string = "FROM_YEAR"
var OrderHistoryData_TOHOUR string = "TO_HOUR"
var OrderHistoryData_TODAY string = "TO_DAY"
var OrderHistoryData_TOMONTH string = "TO_MONTH"
var OrderHistoryData_TOYEAR string = "TO_YEAR"
var OrderHistoryData_ORDERHISTORY string = "ORDERHISTORY"
var OrderHistoryData_SHIPPEDDATE string = "ORDERHISTORY_SHIPPED_DATE"
var OrderHistoryData_ORDERDATE string = "ORDERHISTORY_ORDER_DATE"
var OrderHistoryData_TRANSDATE string = "ORDERHISTORY_TRANS_DATE"
var OrderHistoryData_CANCELDATE string = "ORDERHISTORY_CANCEL_DATE"
var OrderHistoryData_SHIPPEDDATEFORMATTED string = "ORDERHISTORY_SHIPPED_DATE_FORMATTED"
var OrderHistoryData_ORDERDATEFORMATTED string = "ORDERHISTORY_ORDER_DATE_FORMATTED"
var OrderHistoryData_TRANSDATEFORMATTED string = "ORDERHISTORY_TRANS_DATE_FORMATTED"
var OrderHistoryData_CANCELDATEFORMATTED string = "ORDERHISTORY_CANCEL_DATE_FORMATTED"
var OrderHistoryData_SUBTOTAL string = "ORDERHISTORY_SUB_TOTAL"
var OrderHistoryData_SHIPPINGCOST string = "ORDERHISTORY_SHIPPING_COST"
var OrderHistoryData_TAX string = "ORDERHISTORY_TAX"
var OrderHistoryData_TOTAL string = "ORDERHISTORY_TOTAL"
var OrderHistoryData_CANCELINFO string = "ORDERHISTORY_CANCEL_INFO"
var OrderHistoryData_CANCELTYPE string = "ORDERHISTORY_CANCEL_TYPE"
var OrderHistoryData_STATUS string = "ORDERHISTORY_STATUS"
var OrderHistoryData_PREPROCESSINGNAME string = "PREPROCESSING"
var OrderHistoryData_PROCESSINGNAME string = "PROCESSING"
var OrderHistoryData_CANCELLEDNAME string = "CANCELLED"
var OrderHistoryData_PARTIALLYSHIPPEDNAME string = "PARTIALLYSHIPPED"
var OrderHistoryData_SHIPPEDNAME string = "SHIPPED"
var OrderHistoryData_PREPROCESSING string = "Order Preprocessing"
var OrderHistoryData_PROCESSING string = "Order Processing"
var OrderHistoryData_CANCELLED string = "Order Cancelled"
var OrderHistoryData_PARTIALLYSHIPPED string = "Order Partially Shipped"
var OrderHistoryData_SHIPPED string = "Order Shipped"
var OrderHistoryData_SETSTATUS string = "Set Status"

func NewOrderHistoryData() *OrderHistoryData {
    this := &OrderHistoryData{}
    return this
}

