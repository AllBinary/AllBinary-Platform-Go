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

type PaymentGatewayViewAbstract struct {
    TransformInfoHttpComposite
    PaymentGatewayPrimaryKey *PaymentGatewayPrimaryKey
}


func (this *PaymentGatewayViewAbstract) initFields() {
    
}


func NewPaymentGatewayViewAbstract(transformInfoInterface TransformInfoInterface) *PaymentGatewayViewAbstract {
    this := &PaymentGatewayViewAbstract{}
        this.TransformInfoHttpComposite = *NewTransformInfoHttpComposite(transformInfoInterface)
        this.initFields()
        this.GetFormData(any(this.GetPageContext().GetRequest()).(*HttpServletRequest))
    return this
}


func (this *PaymentGatewayViewAbstract) GetFormData(request *HttpServletRequest) {
        var storeName string = this.GetWeblisketSession().GetStoreName(); _ = storeName
        var gatewayName string = request.GetParameter(fmt.Sprint(PaymentGatewayData_NAME)); _ = gatewayName
        this.SetPaymentGatewayPrimaryKey(NewPaymentGatewayPrimaryKey(storeName, gatewayName))
}


func (this *PaymentGatewayViewAbstract) GetPaymentGatewayPrimaryKey() *PaymentGatewayPrimaryKey {
        return this.PaymentGatewayPrimaryKey
}


func (this *PaymentGatewayViewAbstract) SetPaymentGatewayPrimaryKey(paymentGatewayPrimaryKey *PaymentGatewayPrimaryKey) {
        this.PaymentGatewayPrimaryKey = paymentGatewayPrimaryKey
}

