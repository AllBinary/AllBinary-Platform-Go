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

type PaymentType struct {
    BasicPaymentType *BasicPaymentType
    PaymentTransactionInterfaceFactoryInterface PaymentTransactionInterfaceFactoryInterface
    PaymentProcessorInterfaceFactoryInterface PaymentProcessorInterfaceFactoryInterface
}


func (this *PaymentType) initFields() {
    
    
    
}


func NewPaymentType(basicPaymentType *BasicPaymentType) *PaymentType {
    this := &PaymentType{}
    this.initFields()
        this.BasicPaymentType = basicPaymentType
        this.PaymentTransactionInterfaceFactoryInterface = nil
        this.PaymentProcessorInterfaceFactoryInterface = nil
        PaymentTypeUtil_GetInstance().Add(this)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewPaymentType(basicPaymentType *BasicPaymentType, paymentTransactionInterfaceFactoryInterface PaymentTransactionInterfaceFactoryInterface, paymentProcessorInterfaceFactoryInterface PaymentProcessorInterfaceFactoryInterface) *PaymentType {
    this := &PaymentType{}
    this.initFields()
        this.BasicPaymentType = basicPaymentType
        this.PaymentTransactionInterfaceFactoryInterface = paymentTransactionInterfaceFactoryInterface
        this.PaymentProcessorInterfaceFactoryInterface = paymentProcessorInterfaceFactoryInterface
        PaymentTypeUtil_GetInstance().Add(this)
    return this
}


// throws Exception
func (this *PaymentType) GetPaymentTransactionInterfaceFactoryInterface() PaymentTransactionInterfaceFactoryInterface {
        if this.PaymentTransactionInterfaceFactoryInterface == nil {
            panic(errors.New(fmt.Sprint("PaymentTransactionInterfaceFactoryInterface is Null")))
        }
        return this.PaymentTransactionInterfaceFactoryInterface
}


// throws Exception
func (this *PaymentType) GetPaymentProcessorInterfaceFactoryInterface() PaymentProcessorInterfaceFactoryInterface {
        if this.PaymentProcessorInterfaceFactoryInterface == nil {
            panic(errors.New(fmt.Sprint("PaymentProcessorInterfaceFactoryInterface is Null")))
        }
        return this.PaymentProcessorInterfaceFactoryInterface
}


func (this *PaymentType) ToVector() *BasicArrayList {
        var vector *BasicArrayList = NewBasicArrayListD(); _ = vector
        vector.AddAllList(this.GetBasicPaymentType().ToVector())
        if this.PaymentTransactionInterfaceFactoryInterface != nil {
            vector.Add(reflect.TypeOf(this.PaymentTransactionInterfaceFactoryInterface).String())
        }
        if this.PaymentProcessorInterfaceFactoryInterface != nil {
            vector.Add(reflect.TypeOf(this.PaymentProcessorInterfaceFactoryInterface).String())
        }
        return vector
}


func (this *PaymentType) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.PutAll(this.GetBasicPaymentType().ToHashMap())
        if this.PaymentTransactionInterfaceFactoryInterface != nil {
            hashMap.Put(fmt.Sprint(PaymentGatewayData_PAYMENTTRANSACTIONINTERFACEFACTORYINTERFACE), reflect.TypeOf(this.PaymentTransactionInterfaceFactoryInterface).String())
        }
        if this.PaymentProcessorInterfaceFactoryInterface != nil {
            hashMap.Put(fmt.Sprint(PaymentGatewayData_PAYMENTPROCESSORINTERFACEFACTORYINTERFACE), reflect.TypeOf(this.PaymentProcessorInterfaceFactoryInterface).String())
        }
        return hashMap
}


func (this *PaymentType) GetBasicPaymentType() *BasicPaymentType {
        return this.BasicPaymentType
}

