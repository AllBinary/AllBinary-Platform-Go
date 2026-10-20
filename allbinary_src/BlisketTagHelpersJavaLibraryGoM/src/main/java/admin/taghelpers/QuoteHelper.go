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

type QuoteHelper struct {
    BasicTable
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    Portion *Portion
}


func (this *QuoteHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
}


func NewQuoteHelper(hashMap *HashMap, pageContext *PageContext) *QuoteHelper {
    this := &QuoteHelper{}
    this.BasicTable = *NewBasicTable()
    this.initFields()
        var storeName string = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        } else {
            this.StoreFrontInterface = NewStoreFront()
        }
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
        this.Portion = NewPortion(hashMap)
    return this
}


// throws Exception
func (this *QuoteHelper) EmailUser(quoteRequest *QuoteRequest) {
        var user UserInterface = UserEntityFactory_GetInstance().GetUser(quoteRequest.GetUserName()); _ = user
        var userEmailSubject string = "Quote Request Receipt"; _ = userEmailSubject
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("This is just a verification email. ")
        stringBuffer.Append("We usually respond to quote request within 24 hours.")
        stringBuffer.Append("\n\nThank You For Your Business.")
        var userEmailTextBody string = fmt.Sprint(stringBuffer); _ = userEmailTextBody
        var basicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(this.StoreFrontInterface, userEmailSubject, userEmailTextBody)).(IBasicEmailInfo); _ = basicEmailInfo
        var emailInfo *EmailInfo = NewEmailInfo(basicEmailInfo); _ = emailInfo
        var userEmailEventHandler *UserEmailEventHandler = UserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_QUOTEREQUEST, user); _ = userEmailEventHandler
        userEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_QUOTEREQUEST, emailInfo)
}


// throws Exception
func (this *QuoteHelper) EmailAdmins(quoteRequest *QuoteRequest) {
        var adminEmailSubject string = "Quote Request"; _ = adminEmailSubject
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("\nUserName: ")
        stringBuffer.Append(quoteRequest.GetUserName())
        stringBuffer.Append("\nProject Info: \n")
        stringBuffer.Append(quoteRequest.GetProjectInfo())
        stringBuffer.Append("\nUser Comments: \n")
        stringBuffer.Append(quoteRequest.GetUserComments())
        stringBuffer.Append("\nBudget: \n")
        stringBuffer.Append(quoteRequest.GetBudget())
        stringBuffer.Append("\nTime Frame: \n")
        stringBuffer.Append(quoteRequest.GetTimeFrame())
        stringBuffer.Append("\nComments: \n")
        stringBuffer.Append(quoteRequest.GetComments())
        var adminEmailTextBody string = fmt.Sprint(stringBuffer); _ = adminEmailTextBody
        var basicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(this.StoreFrontInterface, adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = basicEmailInfo
        var emailInfo *EmailInfo = NewEmailInfo(basicEmailInfo); _ = emailInfo
        var storeAdminUserEmailEventHandler *UserEmailEventHandler = AdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_QUOTEREQUEST); _ = storeAdminUserEmailEventHandler
        var adminUserEmailEventHandler *UserEmailEventHandler = StoreAdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(UserEmailEventNameData_QUOTEREQUEST, this.AbeClientInformation, this.StoreFrontInterface); _ = adminUserEmailEventHandler
        storeAdminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_QUOTEREQUEST, emailInfo)
        adminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_QUOTEREQUEST, emailInfo)
}


// throws Exception
func (this *QuoteHelper) Email() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "email", e)
            }
            tryControl1, tryResult1 = 1, "Thank You For Your Business.<p>"; return
                }
            }()
            var quoteRequestEntity *QuoteRequestEntity = QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance(); _ = quoteRequestEntity
            var userName string = this.WeblisketSession.GetUserName(); _ = userName
            var vector IBasicArrayList = quoteRequestEntity.GetIds(userName); _ = vector
            var id int = 0; _ = id
            var size int = vector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var nextId int = int((any(vector.Get(index)).(int))); _ = nextId
                if id < nextId {
                    id = nextId
                }
            }
            var quoteRequest *QuoteRequest = quoteRequestEntity.Get(userName, id); _ = quoteRequest
            if quoteRequest != nil {
                this.EmailUser(quoteRequest)
                this.EmailAdmins(quoteRequest)
                tryControl1, tryResult1 = 1, "Thank You For Your Business.<p/>"; return
            } else {
                panic(errors.New(fmt.Sprint("No Quote Request")))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteHelper) Drop() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to drop QuoteRequest table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.DROP, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance().DropTable(); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, this.CommonStrings.DROP)
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteHelper) Create() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to create new QuoteRequest table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "create()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance().CreateTable(); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "create()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteHelper) Restore() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to restore backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "restore()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Restore Successful"; _ = success
            var result string = AbSqlTableUtil_GetInstance().RestoreTable(QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance(), this.Portion); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "restore()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteHelper) Backup() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to make backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Restore Successful"; _ = success
            var result string = AbSqlTableUtil_GetInstance().BackupTable(QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance()); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "backup()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

