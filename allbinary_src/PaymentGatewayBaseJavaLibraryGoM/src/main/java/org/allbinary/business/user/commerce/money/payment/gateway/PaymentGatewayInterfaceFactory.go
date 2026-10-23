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

type PaymentGatewayInterfaceFactory struct {
    LogUtil *LogUtil
}

// Implements PaymentGatewayInterfaceFactoryInterface


func (this *PaymentGatewayInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPaymentGatewayInterfaceFactory() *PaymentGatewayInterfaceFactory {
    this := &PaymentGatewayInterfaceFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *PaymentGatewayInterfaceFactory) GetInstance_1(httpServletRequest *HttpServletRequest) PaymentGatewayInterface {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentGatewayInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "getInstance(HttpServletRequest)", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                this.LogUtil.PutF("Getting", this, commonStrings.GET_INSTANCE)
            }
            tryControl1, tryResult1 = 1, this.GetInstance(NewRequestParams(httpServletRequest).ToHashMap()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of PaymentGatewayInterfaceFactory.getInstance
func (this *PaymentGatewayInterfaceFactory) GetInstance(args ...any) PaymentGatewayInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return this.GetInstance_1(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*BasicPaymentType](args[0], pass) {
            return this.GetInstance_2(Java_Arg[*BasicPaymentType](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return this.GetInstance_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of PaymentGatewayInterfaceFactory.getInstance for the arguments")
}


// throws Exception
func (this *PaymentGatewayInterfaceFactory) GetInstance_2(paymentType *BasicPaymentType) PaymentGatewayInterface {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(fmt.Sprint(PaymentGatewayData_NAME), paymentType.GetName())
        return this.GetInstance(hashMap)
}


// throws Exception
func (this *PaymentGatewayInterfaceFactory) GetInstance_3(hashMap *HashMap) PaymentGatewayInterface {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentGatewayInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            panic(e)
                }
            }()
            var gatewayName string = any(hashMap.Get(fmt.Sprint(PaymentGatewayData_NAME))).(string); _ = gatewayName
            var paymentType *BasicPaymentType = BasicPaymentTypeUtil_GetInstance().Get(gatewayName); _ = paymentType
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                this.LogUtil.PutF("Getting PaymentGatewayInterface for PaymentType: " + paymentType.GetName() + " with: " + fmt.Sprint(hashMap), this, commonStrings.GET_INSTANCE)
            }
            var paymentGatewayInterfaceFactoryInterface PaymentGatewayInterfaceFactoryInterface = paymentType.GetPaymentGatewayInterfaceFactoryInterface(); _ = paymentGatewayInterfaceFactoryInterface
            tryControl1, tryResult1 = 1, paymentGatewayInterfaceFactoryInterface.GetInstance(hashMap); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

