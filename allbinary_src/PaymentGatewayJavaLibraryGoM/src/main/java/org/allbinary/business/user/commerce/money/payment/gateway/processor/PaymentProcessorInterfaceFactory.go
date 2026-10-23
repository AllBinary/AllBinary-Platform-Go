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

type PaymentProcessorInterfaceFactory struct {
    LogUtil *LogUtil
}

var PaymentProcessorInterfaceFactory_Instance *PaymentProcessorInterfaceFactory = NewPaymentProcessorInterfaceFactory()

func (this *PaymentProcessorInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func PaymentProcessorInterfaceFactory_GetInstance() *PaymentProcessorInterfaceFactory {
        return PaymentProcessorInterfaceFactory_Instance
}


func NewPaymentProcessorInterfaceFactory() *PaymentProcessorInterfaceFactory {
    this := &PaymentProcessorInterfaceFactory{}
    this.initFields()
    return this
}


// throws Exception
// synchronized
func (this *PaymentProcessorInterfaceFactory) GetInstance(transformInfoInterface TransformInfoInterface) PaymentProcessorInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentProcessorInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var paymentTransactionInterface PaymentTransactionInterface = PaymentTransactionInterfaceRequestFactory_GetInstance().GetInstance(transformInfoInterface); _ = paymentTransactionInterface
            var gatewayName string = paymentTransactionInterface.GetOrderHistory().GetPaymentMethod(); _ = gatewayName
            if !StringValidationUtil_GetInstance().IsEmpty(gatewayName) {
                var paymentType *PaymentType = PaymentTypeUtil_GetInstance().Get(gatewayName); _ = paymentType
                var paymentProcessorInterfaceFactoryInterface PaymentProcessorInterfaceFactoryInterface = paymentType.GetPaymentProcessorInterfaceFactoryInterface(); _ = paymentProcessorInterfaceFactoryInterface
                tryControl1, tryResult1 = 1, any(paymentProcessorInterfaceFactoryInterface.GetInstance(transformInfoInterface, paymentTransactionInterface)).(PaymentProcessorInterface); return
            }
            panic(errors.New(fmt.Sprint("Error Getting PaymentProcessorInterface")))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

