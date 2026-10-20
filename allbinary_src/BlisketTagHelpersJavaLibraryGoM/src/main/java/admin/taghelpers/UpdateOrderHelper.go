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

type UpdateOrderHelper struct {
    TagHelper
    LogUtil *LogUtil
    Request *HttpServletRequest
    EntryId string
    OrderId string
    UserName string
    StoreName string
    ShipDate string
    OrderDate string
    TransDate string
    CancelDate string
    Status string
    PaymentMethod string
    Name string
    Type string
    Expiration string
    Number string
    Encryption string
    BillingName string
    BillingStreet string
    BillingCity string
    BillingState string
    BillingCode string
    BillingCountry string
    ShippingName string
    ShippingStreet string
    ShippingCity string
    ShippingState string
    ShippingCode string
    ShippingCountry string
    ShippingMethodName string
    SubTotal string
    ShippingCost string
    Tax string
    Total string
    Special string
    UserComment string
    UserCancelComment string
    StoreComment string
    StoreCancelComment string
}


func (this *UpdateOrderHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewUpdateOrderHelper(hashMap *HashMap, pageContext *PageContext) *UpdateOrderHelper {
    this := &UpdateOrderHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *UpdateOrderHelper) GetFormData() {
        this.EntryId = this.Request.GetParameter(EntryData_GetInstance().ID)
        this.OrderId = this.Request.GetParameter(OrderData_ID)
        this.UserName = this.Request.GetParameter(UserData_USERNAME)
        this.StoreName = this.Request.GetParameter(StoreFrontData_GetInstance().NAME)
        this.ShipDate = this.Request.GetParameter(OrderHistoryData_SHIPPEDDATE)
        this.OrderDate = this.Request.GetParameter(OrderHistoryData_ORDERDATE)
        this.TransDate = this.Request.GetParameter(OrderHistoryData_TRANSDATE)
        this.CancelDate = this.Request.GetParameter(OrderHistoryData_CANCELDATE)
        this.Status = this.Request.GetParameter(OrderHistoryData_STATUS)
        this.PaymentMethod = this.Request.GetParameter(PaymentData_METHOD)
        this.Name = this.Request.GetParameter(PaymentData_NAME)
        this.Type = this.Request.GetParameter(PaymentData_TYPE)
        this.Expiration = this.Request.GetParameter(PaymentData_EXPIRATION)
        this.Number = this.Request.GetParameter(PaymentData_NUMBER)
        this.Encryption = this.Request.GetParameter(EntryData_GetInstance().ENCRYPTION)
        this.BillingName = this.Request.GetParameter(BillingAddressData_NAME)
        this.BillingStreet = this.Request.GetParameter(BillingAddressData_STREET)
        this.BillingCity = this.Request.GetParameter(BillingAddressData_CITY)
        this.BillingState = this.Request.GetParameter(BillingAddressData_STATE)
        this.BillingCode = this.Request.GetParameter(BillingAddressData_CODE)
        this.BillingCountry = this.Request.GetParameter(BillingAddressData_COUNTRY)
        this.ShippingName = this.Request.GetParameter(ShippingAddressData_NAME)
        this.ShippingStreet = this.Request.GetParameter(ShippingAddressData_STREET)
        this.ShippingCity = this.Request.GetParameter(ShippingAddressData_CITY)
        this.ShippingState = this.Request.GetParameter(ShippingAddressData_STATE)
        this.ShippingCode = this.Request.GetParameter(ShippingAddressData_CODE)
        this.ShippingCountry = this.Request.GetParameter(ShippingAddressData_COUNTRY)
        this.ShippingMethodName = this.Request.GetParameter(ShippingMethodData_NAME)
        this.SubTotal = this.Request.GetParameter(OrderHistoryData_SUBTOTAL)
        this.ShippingCost = this.Request.GetParameter(OrderHistoryData_SHIPPINGCOST)
        this.Tax = this.Request.GetParameter(OrderHistoryData_TAX)
        this.Total = this.Request.GetParameter(OrderHistoryData_TOTAL)
        this.Special = this.Request.GetParameter(EntryData_GetInstance().SPECIAL)
        this.UserComment = this.Request.GetParameter(OrderData_CUSTOMERCOMMENT)
        this.UserCancelComment = this.Request.GetParameter(OrderData_CUSTOMERCANCELCOMMENT)
        this.StoreComment = this.Request.GetParameter(OrderData_STORECOMMENT)
        this.StoreCancelComment = this.Request.GetParameter(OrderData_STORECANCELCOMMENT)
}


func (this *UpdateOrderHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update order table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Successful update"; _ = success
            var orderHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = orderHashMap
            var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
            orderHashMap.Put(EntryData_GetInstance().ID, this.EntryId)
            orderHashMap.Put(OrderData_ID, this.OrderId)
            orderHashMap.Put(UserData_USERNAME, this.UserName)
            orderHashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
            orderHashMap.Put(OrderHistoryData_SHIPPEDDATE, this.ShipDate)
            orderHashMap.Put(OrderHistoryData_ORDERDATE, this.OrderDate)
            orderHashMap.Put(OrderHistoryData_TRANSDATE, this.TransDate)
            orderHashMap.Put(OrderHistoryData_CANCELDATE, this.CancelDate)
            orderHashMap.Put(OrderHistoryData_STATUS, this.Status)
            orderHashMap.Put(PaymentData_METHOD, this.PaymentMethod)
            orderHashMap.Put(PaymentData_NAME, this.Name)
            orderHashMap.Put(PaymentData_TYPE, this.Type)
            orderHashMap.Put(PaymentData_EXPIRATION, this.Expiration)
            orderHashMap.Put(BillingAddressData_NAME, this.BillingName)
            orderHashMap.Put(BillingAddressData_STREET, this.BillingStreet)
            orderHashMap.Put(BillingAddressData_CITY, this.BillingCity)
            orderHashMap.Put(BillingAddressData_STATE, this.BillingState)
            orderHashMap.Put(BillingAddressData_CODE, this.BillingCode)
            orderHashMap.Put(BillingAddressData_COUNTRY, this.BillingCountry)
            orderHashMap.Put(ShippingAddressData_NAME, this.ShippingName)
            orderHashMap.Put(ShippingAddressData_STREET, this.ShippingStreet)
            orderHashMap.Put(ShippingAddressData_CITY, this.ShippingCity)
            orderHashMap.Put(ShippingAddressData_STATE, this.ShippingState)
            orderHashMap.Put(ShippingAddressData_CODE, this.ShippingCode)
            orderHashMap.Put(ShippingAddressData_COUNTRY, this.ShippingCountry)
            orderHashMap.Put(ShippingMethodData_NAME, this.ShippingMethodName)
            orderHashMap.Put(OrderHistoryData_SUBTOTAL, this.SubTotal)
            orderHashMap.Put(OrderHistoryData_SHIPPINGCOST, this.ShippingCost)
            orderHashMap.Put(OrderHistoryData_TAX, this.Tax)
            orderHashMap.Put(OrderHistoryData_TOTAL, this.Total)
            orderHashMap.Put(EntryData_GetInstance().SPECIAL, this.Special)
            orderHashMap.Put(OrderData_CUSTOMERCOMMENT, this.UserComment)
            orderHashMap.Put(OrderData_CUSTOMERCANCELCOMMENT, this.UserCancelComment)
            orderHashMap.Put(OrderData_STORECOMMENT, this.StoreComment)
            orderHashMap.Put(OrderData_STORECANCELCOMMENT, this.StoreCancelComment)
            whereHashMap.Put(OrderData_ID, this.OrderId)
            OrderHistoryEntityFactory_GetInstance().Update(whereHashMap, orderHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

