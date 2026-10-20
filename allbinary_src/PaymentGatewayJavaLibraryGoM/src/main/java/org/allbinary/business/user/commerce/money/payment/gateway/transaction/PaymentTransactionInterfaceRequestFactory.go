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

type PaymentTransactionInterfaceRequestFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

var PaymentTransactionInterfaceRequestFactory_Instance *PaymentTransactionInterfaceRequestFactory = NewPaymentTransactionInterfaceRequestFactory()

func (this *PaymentTransactionInterfaceRequestFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func PaymentTransactionInterfaceRequestFactory_GetInstance() *PaymentTransactionInterfaceRequestFactory {
        return PaymentTransactionInterfaceRequestFactory_Instance
}


func NewPaymentTransactionInterfaceRequestFactory() *PaymentTransactionInterfaceRequestFactory {
    this := &PaymentTransactionInterfaceRequestFactory{}
    this.initFields()
    return this
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *PaymentTransactionInterfaceRequestFactory) GetInstance(transformInfoInterface TransformInfoInterface) PaymentTransactionInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentTransactionInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getPaymentTransactionInterface()", e)
            }
            panic(e)
                }
            }()
            var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var pageContext *PageContext = httpTransformInfoInterface.GetPageContext(); _ = pageContext
            var httpServletRequest *HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest); _ = httpServletRequest
            var command string = httpServletRequest.GetParameter(org_Allbinary.Globals.GLOBALS2.ADMINCOMMAND); _ = command
            if command != "" && strings.Compare(command, org_Allbinary.Globals.GLOBALS2.AUTHORIZEFORMEVALBODYONERROR) == 0 {
                tryControl1, tryResult1 = 1, this.GenerateFromTestData(transformInfoInterface); return
            } else {
                var propertiesHashMap *HashMap = httpTransformInfoInterface.GetPropertiesHashMap(); _ = propertiesHashMap
                var weblisketSession *WeblisketSession = NewWeblisketSession(propertiesHashMap, pageContext); _ = weblisketSession
                var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(weblisketSession.GetStoreName()); _ = storeFrontInterface
                var order OrderInterface = weblisketSession.GetOrder(); _ = order
                var orderId string = order.GetId(); _ = orderId
                var orderHistoryEntityInterface OrderHistoryEntityInterface = OrderHistoryEntityFactory_GetInstance(); _ = orderHistoryEntityInterface
                var orderReview *OrderHistory = orderHistoryEntityInterface.GetOrder(orderId); _ = orderReview
                var orderPaymentInfo *Payment = orderReview.GetPaymentInfo(); _ = orderPaymentInfo
                orderPaymentInfo.SetTransactionType(fmt.Sprint(TransactionTypeFactory_GetInstance().SALE))
                orderPaymentInfo.SetTenderType(fmt.Sprint(TenderTypeFactory_GetInstance().CREDITCARD))
                orderReview.SetPaymentInfo(orderPaymentInfo)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                    this.LogUtil.PutF(orderReview.GetId(), this, "getPaymentTransactionInterface()")
                }
                var paymentTransactionInterface PaymentTransactionInterface = PaymentTransactionInterfaceFactory_GetInstance().GetInstance(orderReview); _ = paymentTransactionInterface
                tryControl1, tryResult1 = 1, paymentTransactionInterface; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *PaymentTransactionInterfaceRequestFactory) GenerateFromTestData(transformInfoInterface TransformInfoInterface) PaymentTransactionInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentTransactionInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENTERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generatePaymentTransactionInterfaceFromTestData()", e)
            }
            panic(e)
                }
            }()
            var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            var propertiesHashMap *HashMap = httpTransformInfoInterface.GetPropertiesHashMap(); _ = propertiesHashMap
            var pageContext *PageContext = httpTransformInfoInterface.GetPageContext(); _ = pageContext
            var weblisketSession *WeblisketSession = NewWeblisketSession(propertiesHashMap, pageContext); _ = weblisketSession
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(weblisketSession.GetStoreName()); _ = storeFrontInterface
            var httpServletRequest *HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest); _ = httpServletRequest
            var gatewayName string = any(weblisketSession.GetPaymentMethod()).(string); _ = gatewayName
            if StringValidationUtil_GetInstance().IsEmpty(gatewayName) {
                gatewayName = httpServletRequest.GetParameter(fmt.Sprint(PaymentGatewayData_NAME))
            }
            var transactionType string = httpServletRequest.GetParameter(PaymentData_TRANSACTIONTYPE); _ = transactionType
            var tenderType string = httpServletRequest.GetParameter(PaymentData_TENDERTYPE); _ = tenderType
            var payment *Payment = NewPayment(); _ = payment
            payment.SetTransactionType(transactionType)
            payment.SetTenderType(tenderType)
            var orderReview *OrderHistory = NewOrderHistory(NewBasket()); _ = orderReview
            orderReview.SetStoreName(storeFrontInterface.GetName())
            orderReview.SetPaymentMethod(gatewayName)
            orderReview.SetUserName("testing")
            var orderId string = httpServletRequest.GetParameter("orderNumber"); _ = orderId
            if orderId == "" || strings.Compare(orderId, StringUtil_GetInstance().EMPTY_STRING) == 0 {
                orderId = NewOrderIdGenerator().GetNext()
                orderReview.SetId(orderId)
                var account string = httpServletRequest.GetParameter("account"); _ = account
                var expirationDate string = httpServletRequest.GetParameter("expirationDate"); _ = expirationDate
                var checkNumber string = httpServletRequest.GetParameter("checkNumber"); _ = checkNumber
                var driversLicense string = httpServletRequest.GetParameter("driversLicense"); _ = driversLicense
                var magneticInkCheckReader string = httpServletRequest.GetParameter("magneticInkCheckReader"); _ = magneticInkCheckReader
                var aba string = httpServletRequest.GetParameter("aba"); _ = aba
                var accountType string = httpServletRequest.GetParameter("accountType"); _ = accountType
                var name string = httpServletRequest.GetParameter("name"); _ = name
                payment.SetName(name)
                payment.SetNumber(account)
                payment.SetExpiration(expirationDate)
                payment.SetCheckNumber(checkNumber)
                payment.SetDriversLicense(driversLicense)
                payment.SetMagneticInkCheckReader(magneticInkCheckReader)
                payment.SetAba(aba)
                payment.SetAccountType(accountType)
                var street string = httpServletRequest.GetParameter("street"); _ = street
                var city string = httpServletRequest.GetParameter("city"); _ = city
                var state string = httpServletRequest.GetParameter("state"); _ = state
                var zip string = httpServletRequest.GetParameter("zip"); _ = zip
                var streetAddress *StreetAddress = NewStreetAddress(); _ = streetAddress
                streetAddress.SetName(name)
                streetAddress.SetStreet(street)
                streetAddress.SetCity(city)
                streetAddress.SetState(state)
                streetAddress.SetCode(zip)
                var email string = httpServletRequest.GetParameter("email"); _ = email
                var amount string = httpServletRequest.GetParameter("amount"); _ = amount
                orderReview.SetTotal(NewMoney(amount))
                orderReview.SetBillingAddress(streetAddress)
            } else {
                orderReview.SetId(orderId)
            }
            orderReview.SetPaymentInfo(payment)
            var paymentTransactionInterface PaymentTransactionInterface = PaymentTransactionInterfaceFactory_GetInstance().GetInstance(orderReview); _ = paymentTransactionInterface
            tryControl1, tryResult1 = 1, paymentTransactionInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

