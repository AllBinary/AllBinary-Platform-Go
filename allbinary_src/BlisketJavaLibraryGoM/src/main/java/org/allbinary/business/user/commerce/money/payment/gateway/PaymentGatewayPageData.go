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

type PaymentGatewayPageData struct {
    MAKEPAYMENT string
    CHECKOUT string
    SHIPPING string
    SHIPPINGADDRESS string
    SHIPPINGADDRESSACTION string
    BILLINGADDRESS string
    PAYMENT string
    AUTHORIZE string
    FINISH string
}

var PaymentGatewayPageData_EXTENSION string = ".jsp"
var PaymentGatewayPageData_NAME string = StringUtil_GetInstance().EMPTY_STRING
var PaymentGatewayPageData_FORWARDMAKEPAYMENT string = "MakePayment.jsp"
var PaymentGatewayPageData_PAYMENTOPTIONS string = "PaymentOptions" + fmt.Sprint(PaymentGatewayPageData_EXTENSION)
var PaymentGatewayPageData_STARTCHECKOUT string = "StartCheckout" + fmt.Sprint(PaymentGatewayPageData_EXTENSION)

func (this *PaymentGatewayPageData) initFields() {
    
    
    
    
    
    
    
    
    
}


func NewPaymentGatewayPageData(location string, paymentMethod string) *PaymentGatewayPageData {
    this := &PaymentGatewayPageData{}
    this.initFields()
        var paymentMethodString string = fmt.Sprint(PaymentGatewayPageData_NAME) + paymentMethod; _ = paymentMethodString
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(location)
        stringBuffer.Append("MakePayment")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.MAKEPAYMENT = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("Checkout")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.CHECKOUT = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("Shipping")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.SHIPPING = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("ShippingAddress")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.SHIPPINGADDRESS = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("ShippingAddressAction")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.SHIPPINGADDRESSACTION = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("BillingAddress")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.BILLINGADDRESS = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("Payment")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.PAYMENT = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("Authorize")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.AUTHORIZE = fmt.Sprint(stringBuffer)
        stringBuffer.Delete(0, stringBuffer.Length())
        stringBuffer.Append(paymentMethodString)
        stringBuffer.Append("PaymentFinish")
        stringBuffer.Append(PaymentGatewayPageData_EXTENSION)
        this.FINISH = fmt.Sprint(stringBuffer)
    return this
}

