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

type StoreFrontsRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    PageContext *PageContext
    Request *HttpServletRequest
    WeblisketSession *WeblisketSession
    StoreName string
    ModifyingStoreFrontInterface StoreFrontInterface
    Portion *Portion
}


func (this *StoreFrontsRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
    
    
    
}


func NewStoreFrontsRequestHelper(hashMap *HashMap, pageContext *PageContext) *StoreFrontsRequestHelper {
    this := &StoreFrontsRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        this.PageContext = pageContext
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
        this.Portion = NewPortion(hashMap)
        this.GetFormData()
    return this
}


// throws Exception
func (this *StoreFrontsRequestHelper) GetFormData() {
        this.StoreName = this.WeblisketSession.GetStoreName()
        this.ModifyingStoreFrontInterface = any(NewStoreFront(this.Request)).(StoreFrontInterface)
}


func (this *StoreFrontsRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update storefronts table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Updated Successfully"; _ = success
            var hashMapData *HashMap = this.ModifyingStoreFrontInterface.ToHashMap(); _ = hashMapData
            StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance().Update(hashMapData)
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


// throws Exception
func (this *StoreFrontsRequestHelper) SendStoreCreatedEmails() {
        var storeManagerEmailSubject string = "Store Manager Nofification For Store: " + this.ModifyingStoreFrontInterface.GetName(); _ = storeManagerEmailSubject
        var adminEmailSubject string = "Admin Notification For Store: " + this.ModifyingStoreFrontInterface.GetName(); _ = adminEmailSubject
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Store Created: ")
        stringBuffer.Append(this.ModifyingStoreFrontInterface.GetName())
        stringBuffer.Append("\n\n")
        stringBuffer.Append("Click here for the Admin Interface:\n")
        stringBuffer.Append(this.ModifyingStoreFrontInterface.GetCurrentHomeHostName())
        stringBuffer.Append("/admin/login.jsp\n\n")
        stringBuffer.Append("Goto Your Store:\n")
        stringBuffer.Append(this.ModifyingStoreFrontInterface.GetCurrentHomeHostName())
        stringBuffer.Append(this.ModifyingStoreFrontInterface.GetName())
        stringBuffer.Append("/index.jsp")
        stringBuffer.Append("\n\n")
        stringBuffer.Append("UserName: ")
        stringBuffer.Append(this.WeblisketSession.GetUserName())
        var adminEmailTextBody string = fmt.Sprint(stringBuffer); _ = adminEmailTextBody
        var adminBasicEmailInfo IBasicEmailInfo = any(NewAdminEmailInfo(adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = adminBasicEmailInfo
        var storeAdminBasicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(this.ModifyingStoreFrontInterface, storeManagerEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = storeAdminBasicEmailInfo
        var storeAdminEmailInfo *EmailInfo = NewEmailInfo(storeAdminBasicEmailInfo); _ = storeAdminEmailInfo
        var adminEmailInfo *EmailInfo = NewEmailInfo(adminBasicEmailInfo); _ = adminEmailInfo
        var adminUserEmailEventHandler *UserEmailEventHandler = AdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_STORECREATED); _ = adminUserEmailEventHandler
        var storeAdminUserEmailEventHandler *UserEmailEventHandler = StoreAdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(UserEmailEventNameData_STORECREATED, this.AbeClientInformation, this.ModifyingStoreFrontInterface); _ = storeAdminUserEmailEventHandler
        storeAdminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_STORECREATED, storeAdminEmailInfo)
        adminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_STORECREATED, adminEmailInfo)
}


// synchronized
func (this *StoreFrontsRequestHelper) Install() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to Install storefront"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "install()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Intall StoreFront Successfully: ")
            stringBuffer.Appendint(int(this.Portion.GetCurrent()))
            stringBuffer.Append(" of ")
            stringBuffer.Appendint(int(this.Portion.GetTotal()))
            var success string = fmt.Sprint(stringBuffer); _ = success
            this.ModifyingStoreFrontInterface.Install(int(this.Portion.GetCurrent()), int(this.Portion.GetTotal()))
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "install()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *StoreFrontsRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add storefront"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insert()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Added Successfully"; _ = success
            var values IBasicArrayList = this.ModifyingStoreFrontInterface.ToVector(); _ = values
            StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance().Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "insert()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreFrontsRequestHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to delete storefronts table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Delete Successfully"; _ = success
            StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance().Delete(this.StoreName)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "delete()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

