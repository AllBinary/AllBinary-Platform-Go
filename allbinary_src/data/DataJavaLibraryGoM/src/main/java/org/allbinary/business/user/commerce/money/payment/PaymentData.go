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

type PaymentData struct {
}

var PaymentData_MIN int = 4
var PaymentData_MAX int = 19
var PaymentData_ADDNEW string = "Add New Payment"
var PaymentData_VIEW string = "View Payment"
var PaymentData_SELECT string = "Select Payment"
var PaymentData_CHANGE string = "Change Payment"
var PaymentData_PAYMENT string = "PAYMENT"
var PaymentData_METHOD string = "PAYMENT_METHOD"
var PaymentData_ID string = "PAYMENT_ID"
var PaymentData_NUMBER string = "PAYMENT_NUMBER"
var PaymentData_NAME string = "PAYMENT_NAME"
var PaymentData_TYPE string = "PAYMENT_TYPE"
var PaymentData_EXPIRATION string = "PAYMENT_EXPIRATION"
var PaymentData_TENDERTYPE string = "PAYMENT_TENDERTYPE"
var PaymentData_TRANSACTIONTYPE string = "PAYMENT_TRANSACTIONTYPE"
var PaymentData_ABA string = "PAYMENT_ABA"
var PaymentData_ACCOUNT string = "PAYMENT_ACCOUNT"
var PaymentData_ACCOUNTTYPE string = "PAYMENT_ACCOUNTTYPE"
var PaymentData_AUTHORIZATIONCODE string = "PAYMENT_AUTHORIZATIONCODE"
var PaymentData_CHECKNUMBER string = "PAYMENT_CHECKNUMBER"
var PaymentData_CHECKTYPE string = "PAYMENT_CHECKTYPE"
var PaymentData_DRIVERSLICENSE string = "PAYMENT_DRIVERSLICENSE"
var PaymentData_MAGNETICINKCHECKREADER string = "PAYMENT_MAGNETICINKCHECKREADER"

func NewPaymentData() *PaymentData {
    this := &PaymentData{}
    return this
}

