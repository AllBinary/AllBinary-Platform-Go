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

type OrderHistoryHelper struct {
    TagHelper
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
    Request *HttpServletRequest
    StreetAddress *StreetAddress
    Id string
    GroupId string
    Status string
    NewStatus string
}


func (this *OrderHistoryHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
}


func NewOrderHistoryHelper(propertiesHashMap *HashMap, pageContext *PageContext) *OrderHistoryHelper {
    this := &OrderHistoryHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.PropertiesHashMap = propertiesHashMap
        this.PageContext = pageContext
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        var storeName string = any(propertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        this.NewStatus = any(propertiesHashMap.Get(OrderHistoryData_STATUS)).(string)
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        }
        this.GetFormData()
    return this
}


func (this *OrderHistoryHelper) GetFormData() {
        this.Id = this.Request.GetParameter(OrderData_ID)
        this.GroupId = this.Request.GetParameter(ShippingMethodData_GROUP)
        if this.NewStatus == "" {
            this.NewStatus = this.Request.GetParameter(OrderHistoryData_STATUS)
        }
}


func (this *OrderHistoryHelper) SetOrderStatus() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set order status"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setOrderStatus()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br/>" + "Exception: " + fmt.Sprint(e) + "<br/>"; return
                }
            }()
            var success string = "Status successfully set to: " + this.Status; _ = success
            var orderReviewInterface OrderHistoryInterface = OrderHistoryFactory_GetInstance(this.Id); _ = orderReviewInterface
            var status string = orderReviewInterface.GetStatus(); _ = status
            if strings.Compare(status, OrderHistoryData_CANCELLED) == 0 {
                tryControl1, tryResult1 = 1, "Order " + this.Id + " Already Cancelled<br/>"; return
            } else if strings.Compare(status, OrderHistoryData_SHIPPED) == 0 {
                tryControl1, tryResult1 = 1, "Order " + this.Id + " Already Shipped. You can't cancel a shipped order.<br/>"; return
            } else if strings.Compare(status, OrderHistoryData_PREPROCESSING) == 0 {
                tryControl1, tryResult1 = 1, "Order " + this.Id + " has been cancelled.<br/>"; return
            } else if strings.Compare(status, OrderHistoryData_PARTIALLYSHIPPED) == 0 {
                tryControl1, tryResult1 = 1, "Order " + this.Id + " Partially Shipped. You can cancel unshipped items, but not the whole order.<br/>"; return
            } else if strings.Compare(status, OrderHistoryData_PROCESSING) == 0 {
                var orderHistoryEntityInterface OrderHistoryEntityInterface = OrderHistoryEntityFactory_GetInstance(); _ = orderHistoryEntityInterface
                orderHistoryEntityInterface.SetStatus(this.Id, this.NewStatus)
                tryControl1, tryResult1 = 1, "Order " + this.Id + " has been cancelled.<br/>"; return
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "setOrderStatus()")
            }
            tryControl1, tryResult1 = 1, "Error Setting Order Status"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

