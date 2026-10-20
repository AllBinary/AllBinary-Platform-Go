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

type OrderHelper struct {
    TagHelper
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
    Request *HttpServletRequest
}


func (this *OrderHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
}


func NewOrderHelper(propertiesHashMap *HashMap, pageContext *PageContext) *OrderHelper {
    this := &OrderHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.PropertiesHashMap = propertiesHashMap
        this.PageContext = pageContext
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        var storeName string = any(propertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if !StringValidationUtil_GetInstance().IsEmpty(storeName) {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        }
        this.WeblisketSession = NewWeblisketSession(this.PropertiesHashMap, this.PageContext)
    return this
}


func (this *OrderHelper) SetPaymentGateway() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Failed to set PaymentGateway for Order: ")
                if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 bool) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            ex := Java_ToError(recovered)
                            _ = ex
                    stringBuffer.Append(" Exception Getting")
                        }
                    }()
                    var orderInterface OrderInterface = this.WeblisketSession.GetOrder(); _ = orderInterface
                    if orderInterface != nil {
                        stringBuffer.Append(orderInterface.GetId())
                        stringBuffer.Append(" to: ")
                        stringBuffer.Append(orderInterface.GetPaymentMethod())
                    }
                    return
                }(); tryControl2 == 1 {
                    tryControl1, tryResult1 = 1, tryResult2; return
                }
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, "setPaymentGateway()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var paymentGatewayBoolean bool = false; _ = paymentGatewayBoolean
            var orderInterface OrderInterface = this.WeblisketSession.GetOrder(); _ = orderInterface
            var requestPaymentGateway string = any(this.Request.GetParameter(fmt.Sprint(PaymentGatewayData_NAME))).(string); _ = requestPaymentGateway
            if !StringValidationUtil_GetInstance().IsEmpty(requestPaymentGateway) {
                this.WeblisketSession.SetPaymentMethod(requestPaymentGateway)
                orderInterface.SetPaymentMethod(requestPaymentGateway)
                paymentGatewayBoolean = true
            }
            var paymentGatewayEntityInterface *PaymentGatewayEntity = any(PaymentGatewayEntityFactory_GetInstance()).(*PaymentGatewayEntity); _ = paymentGatewayEntityInterface
            var paymentTypeVector IBasicArrayList = paymentGatewayEntityInterface.FindPaymentTypeVectorByStore(this.WeblisketSession.GetStoreName()); _ = paymentTypeVector
            if paymentTypeVector.Size() == 1 {
                var paymentType *BasicPaymentType = any(paymentTypeVector.Get(0)).(*BasicPaymentType); _ = paymentType
                var paymentGatewayInterface PaymentGatewayInterface = any(paymentGatewayEntityInterface.GetPaymentGatewayInterface(this.WeblisketSession.GetStoreName(), paymentType)).(PaymentGatewayInterface); _ = paymentGatewayInterface
                var paymentGateway string = paymentGatewayInterface.GetName(); _ = paymentGateway
                this.WeblisketSession.SetPaymentMethod(paymentGateway)
                orderInterface.SetPaymentMethod(paymentGateway)
                paymentGatewayBoolean = true
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Successfully set PaymentGateway Order: ")
                stringBuffer.Append(orderInterface.GetId())
                stringBuffer.Append(" to: ")
                stringBuffer.Append(orderInterface.GetPaymentMethod())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "setPaymentGateway()")
            }
            tryControl1, tryResult1 = 1, paymentGatewayBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *OrderHelper) Process() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Failed to Process Order: ")
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        ex := Java_ToError(recovered)
                        _ = ex
                stringBuffer.Append(" Exception Getting Id")
                    }
                }()
                stringBuffer.Append(this.WeblisketSession.GetOrder().GetId())
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
            var error_ string = fmt.Sprint(stringBuffer); _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var order OrderInterface = this.WeblisketSession.GetOrder(); _ = order
            order.SetStoreName(this.StoreFrontInterface.GetName())
            var result string = OrderProcessorUtil_GetInstance().Process(this.WeblisketSession.GetUserName(), any(order).(IOrder)); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successfully Processed Order: " + result, this, "processOrder()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

