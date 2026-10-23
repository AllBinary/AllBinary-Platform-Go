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

type BasicPaymentTypeUtil struct {
    PaymentTypeVector IBasicArrayList
}

var BasicPaymentTypeUtil_Instance *BasicPaymentTypeUtil = NewBasicPaymentTypeUtil()

func (this *BasicPaymentTypeUtil) initFields() {
    this.PaymentTypeVector = NewBasicArrayListD()
}


func init() {
}


func BasicPaymentTypeUtil_GetInstance() *BasicPaymentTypeUtil {
        return BasicPaymentTypeUtil_Instance
}


func NewBasicPaymentTypeUtil() *BasicPaymentTypeUtil {
    this := &BasicPaymentTypeUtil{}
    this.initFields()
    return this
}


func (this *BasicPaymentTypeUtil) Add(paymentType *BasicPaymentType) {
        this.PaymentTypeVector.Add(paymentType)
}


// throws Exception
func (this *BasicPaymentTypeUtil) Get(paymentTypeString string) *BasicPaymentType {
        var size int = this.PaymentTypeVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var paymentType *BasicPaymentType = any(this.PaymentTypeVector.Get(i)).(*BasicPaymentType); _ = paymentType
            if strings.Compare(paymentType.GetName(), paymentTypeString) == 0 {
                return paymentType
            }
        }
        panic(errors.New(fmt.Sprint("Unknown PaymentType: " + paymentTypeString)))
}


func (this *BasicPaymentTypeUtil) IsContain(paymentType *BasicPaymentType) bool {
        return this.PaymentTypeVector.Contains(paymentType)
}


func (this *BasicPaymentTypeUtil) Difference(a_PaymentTypeVector IBasicArrayList) IBasicArrayList {
        var diff IBasicArrayList = NewBasicArrayListD(); _ = diff
        var size int = this.PaymentTypeVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var paymentType *BasicPaymentType = any(this.PaymentTypeVector.Get(i)).(*BasicPaymentType); _ = paymentType
            if !a_PaymentTypeVector.Contains(paymentType) {
                diff.Add(paymentType)
            }
        }
        return diff
}

