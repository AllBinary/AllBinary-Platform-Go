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

type StaticPagesRequestHelper struct {
    AbContext
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    Request *HttpServletRequest
    StoreName string
    SearchParams *SearchParams
    XslFile string
}

// Implements TagHelperInterface


func (this *StaticPagesRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
    
}


func NewStaticPagesRequestHelper(propertiesHashMap *HashMap, pageContext *PageContext) *StaticPagesRequestHelper {
    this := &StaticPagesRequestHelper{}
        this.AbContext = *NewAbContext(propertiesHashMap, pageContext)
        this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.XslFile = any(propertiesHashMap.Get(TransformInfoData_GetInstance().TEMPLATEFILENAME)).(string)
        this.GetFormData()
    return this
}


func (this *StaticPagesRequestHelper) GetFormData() {
        this.StoreName = this.Request.GetParameter(StoreFrontData_GetInstance().NAME)
        this.SearchParams = NewSearchParams(this.Request)
}


// throws Exception
func (this *StaticPagesRequestHelper) Email() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "email", e)
            }
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Generated Static Pages Notification Email", this, "email")
            }
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.StoreName); _ = storeFrontInterface
            var adminEmailSubject string = "Generated Static Pages Email Notification"; _ = adminEmailSubject
            var adminEmailTextBody string = "Generated static pages for store: " + this.StoreName; _ = adminEmailTextBody
            var adminBasicEmailInfo IBasicEmailInfo = any(NewAdminEmailInfo(adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = adminBasicEmailInfo
            var storeAdminBasicEmailInfo IBasicEmailInfo = any(NewStoreEmailInfo(storeFrontInterface, adminEmailSubject, adminEmailTextBody)).(IBasicEmailInfo); _ = storeAdminBasicEmailInfo
            var storeAdminEmailInfo *EmailInfo = NewEmailInfo(storeAdminBasicEmailInfo); _ = storeAdminEmailInfo
            var adminEmailInfo *EmailInfo = NewEmailInfo(adminBasicEmailInfo); _ = adminEmailInfo
            var adminUserEmailEventHandler *UserEmailEventHandler = AdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(this.AbeClientInformation, UserEmailEventNameData_STOREGENERATINGSTATICPAGES); _ = adminUserEmailEventHandler
            var storeAdminUserEmailEventHandler *UserEmailEventHandler = StoreAdminUserEmailEventHandlerSingletons_GetInstance().GetInstance(UserEmailEventNameData_STOREGENERATINGSTATICPAGES, this.AbeClientInformation, storeFrontInterface); _ = storeAdminUserEmailEventHandler
            storeAdminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_STOREGENERATINGSTATICPAGES, storeAdminEmailInfo)
            adminUserEmailEventHandler.ReceiveEmailInfo(UserEmailEventNameData_STOREGENERATINGSTATICPAGES, adminEmailInfo)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StaticPagesRequestHelper) GenerateStaticPages() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to generate staticpages table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "generateStaticPages()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var contentType string = AcceptableResponseGenerator_GetInstance().Get(this.Request); _ = contentType
            var searchRequest *SearchRequest = NewSearchRequest(nil, this.SearchParams, this.XslFile, contentType, this.GetPropertiesHashMap(), this.GetPageContext()); _ = searchRequest
            var success string = ProductListingFactory_GetInstance(searchRequest).GenerateAll(this.StoreName); _ = success
            this.Email()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "generateStaticPages()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StaticPagesRequestHelper) MakePublic() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to makePublic"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "makePublic()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.StoreName); _ = storeFrontInterface
            var fromAbPath IAbPath = NewAbPath(storeFrontInterface.GetTestHtmlPath() + storeFrontInterface.GetStaticPath(), StringUtil_GetInstance().EMPTY_STRING); _ = fromAbPath
            var toAbPath IAbPath = NewAbPath(fmt.Sprint(AbContext_URLGLOBALS.GetWebappPath()) + storeFrontInterface.GetName() + fmt.Sprint(AbPathData_GetInstance().SEPARATOR) + fmt.Sprint(storeFrontInterface.GetStaticPath()), StringUtil_GetInstance().EMPTY_STRING); _ = toAbPath
            FileUtil_GetInstance().Copy(fromAbPath, toAbPath)
            var success string = "Made Public"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "makePublic()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

