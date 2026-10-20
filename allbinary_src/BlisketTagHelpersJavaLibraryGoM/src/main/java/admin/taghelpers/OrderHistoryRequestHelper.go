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

type OrderHistoryRequestHelper struct {
    TagHelper
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    Request *HttpServletRequest
    Id string
    GroupId string
    Status string
}


func (this *OrderHistoryRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
    
}


func NewOrderHistoryRequestHelper(hashMap *HashMap, pageContext *PageContext) *OrderHistoryRequestHelper {
    this := &OrderHistoryRequestHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *OrderHistoryRequestHelper) GetFormData() {
        this.Id = this.Request.GetParameter(OrderData_ID)
        this.GroupId = this.Request.GetParameter(ShippingMethodData_GROUP)
        this.Status = this.Request.GetParameter(OrderHistoryData_STATUS)
}


func (this *OrderHistoryRequestHelper) SetOrderStatus_1() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to set order status"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setOrderStatus()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            if strings.Compare(this.GroupId, CommonPhoneStrings_GetInstance().ZERO) != 0 {
                if OrderItemsEntityFactory_GetInstance().IsEverythingShipped(this.Id) {
                    this.Status = OrderHistoryData_SHIPPED
                } else {
                    this.Status = OrderHistoryData_PARTIALLYSHIPPED
                }
            }
            var success string = "Status successfully set to: " + this.Status; _ = success
            OrderHistoryEntityFactory_GetInstance().SetStatus(this.Id, this.Status)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Success", this, "setOrderStatus()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of OrderHistoryRequestHelper.setOrderStatus
func (this *OrderHistoryRequestHelper) SetOrderStatus(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.SetOrderStatus_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.SetOrderStatus_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of OrderHistoryRequestHelper.setOrderStatus for the arguments")
}


func (this *OrderHistoryRequestHelper) SetOrderStatus_2(newStatus string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view order table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setOrderStatus(newStatus)", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            if newStatus == "" {
                newStatus = this.Status
            }
            var success string = "Status successfully set to: " + newStatus; _ = success
            OrderHistoryEntityFactory_GetInstance().SetStatus(this.Id, newStatus)
            var orderHistory *OrderHistory = OrderHistoryEntityFactory_GetInstance().GetOrder(this.Id); _ = orderHistory
            NewOrderStatusEmail(this.AbeClientInformation, orderHistory).Process()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Success", this, "setOrderStatus()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

