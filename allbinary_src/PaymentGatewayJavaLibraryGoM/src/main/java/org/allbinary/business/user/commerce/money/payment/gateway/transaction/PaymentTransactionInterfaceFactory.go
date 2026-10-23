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

type PaymentTransactionInterfaceFactory struct {
    LogUtil *LogUtil
}

var PaymentTransactionInterfaceFactory_Instance *PaymentTransactionInterfaceFactory = NewPaymentTransactionInterfaceFactory()

func (this *PaymentTransactionInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func PaymentTransactionInterfaceFactory_GetInstance() *PaymentTransactionInterfaceFactory {
        return PaymentTransactionInterfaceFactory_Instance
}


func NewPaymentTransactionInterfaceFactory() *PaymentTransactionInterfaceFactory {
    this := &PaymentTransactionInterfaceFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *PaymentTransactionInterfaceFactory) GetInstance(orderReview *OrderHistory) PaymentTransactionInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentTransactionInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put("Failed to generate test data", this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var gatewayName string = orderReview.GetPaymentMethod(); _ = gatewayName
            if gatewayName != "" {
                var paymentType *PaymentType = PaymentTypeUtil_GetInstance().Get(gatewayName); _ = paymentType
                var paymentTransactionInterfaceFactoryInterface PaymentTransactionInterfaceFactoryInterface = paymentType.GetPaymentTransactionInterfaceFactoryInterface(); _ = paymentTransactionInterfaceFactoryInterface
                var paymentTransactionInterface PaymentTransactionInterface = paymentTransactionInterfaceFactoryInterface.GetInstance(orderReview); _ = paymentTransactionInterface
                tryControl1, tryResult1 = 1, paymentTransactionInterface; return
            }
            panic(errors.New(fmt.Sprint("Failed to create PaymentTransactionInterface: " + "GatewayName is: " + gatewayName)))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

