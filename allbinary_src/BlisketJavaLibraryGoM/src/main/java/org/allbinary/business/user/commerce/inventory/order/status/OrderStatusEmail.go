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

type OrderStatusEmail struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    OrderHistory *OrderHistory
    StoreFrontInterface StoreFrontInterface
    User UserInterface
}


func (this *OrderStatusEmail) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
}


func NewOrderStatusEmail(abeClientInformation AbeClientInformationInterface, orderHistory *OrderHistory) *OrderStatusEmail {
    this := &OrderStatusEmail{}
    this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.OrderHistory = orderHistory
        var storeName string = orderHistory.GetStoreName(); _ = storeName
        this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        this.User = UserEntityFactory_GetInstance().GetUser(orderHistory.GetUserName())
    return this
}


// throws Exception
func (this *OrderStatusEmail) Process() {
        this.NotifyStoreAdmin()
        this.NotifyUser()
}


// throws Exception
func (this *OrderStatusEmail) NotifyStoreAdmin() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "emailAdmin", e)
            }
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Order Email", this, "emailStoreAdmin")
            }
            var adminEmailSubject string = "Order " + fmt.Sprint(this.OrderHistory.GetStatus()) + " Attempt"; _ = adminEmailSubject
            var adminEmailTextBody string = "Order attempt: " + fmt.Sprint(this.OrderHistory.GetId()) + " Preprocessing by: " + fmt.Sprint(this.User.GetUserName()); _ = adminEmailTextBody
            var adminBasicEmailInfo IBasicEmailInfo = any(NewAdminEmailInfo(adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = adminBasicEmailInfo
            var storeAdminBasicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(this.StoreFrontInterface, adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = storeAdminBasicEmailInfo
            var storeAdminEmailInfo *EmailInfo = NewEmailInfo(storeAdminBasicEmailInfo); _ = storeAdminEmailInfo
            var adminEmailInfo *EmailInfo = NewEmailInfo(adminBasicEmailInfo); _ = adminEmailInfo
            var userEmailEventNameData *UserEmailEventNameData = UserEmailEventNameData_GetInstance(this.OrderHistory.GetStatus()); _ = userEmailEventNameData
            var adminUserEmailEventHandler *UserEmailEventHandler = AdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, userEmailEventNameData); _ = adminUserEmailEventHandler
            var storeAdminUserEmailEventHandler *UserEmailEventHandler = StoreAdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(userEmailEventNameData, this.AbeClientInformation, this.StoreFrontInterface); _ = storeAdminUserEmailEventHandler
            storeAdminUserEmailEventHandler.ReceiveEmailInfo(userEmailEventNameData, storeAdminEmailInfo)
            adminUserEmailEventHandler.ReceiveEmailInfo(userEmailEventNameData, adminEmailInfo)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *OrderStatusEmail) NotifyUser() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "notifyUser", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Email User", this, "notifyUser()")
            }
            var subject string = "Order " + fmt.Sprint(this.OrderHistory.GetStatus()) + " Attempt"; _ = subject
            var textBody string = "Order attempt: " + fmt.Sprint(this.OrderHistory.GetId()) + " Preprocessing by: " + fmt.Sprint(this.User.GetUserName()); _ = textBody
            var basicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(this.StoreFrontInterface, subject, textBody)).(IBasicEmailInfo); _ = basicEmailInfo
            var emailInfo *EmailInfo = NewEmailInfo(basicEmailInfo); _ = emailInfo
            var userEmailEventNameData *UserEmailEventNameData = UserEmailEventNameData_GetInstance(this.OrderHistory.GetStatus()); _ = userEmailEventNameData
            var userEmailEventHandler *UserEmailEventHandler = UserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, userEmailEventNameData, this.User); _ = userEmailEventHandler
            userEmailEventHandler.ReceiveEmailInfo(userEmailEventNameData, emailInfo)
            return
        }(); tryControl1 == 1 {
            return
        }
}

