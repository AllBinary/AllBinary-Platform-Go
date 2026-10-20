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

type InventoryHelper struct {
    TagHelper
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    WeblisketSession *WeblisketSession
    Response *HttpServletResponse
    Request *HttpServletRequest
    StreetAddress *StreetAddress
    Id string
    SearchRequest *SearchRequest
}


func (this *InventoryHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
}


func NewInventoryHelper(propertiesHashMap *HashMap, pageContext *PageContext) *InventoryHelper {
    this := &InventoryHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.Response = any(pageContext.GetResponse()).(*HttpServletResponse)
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        var storeName string = any(propertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        var xslFile string = any(propertiesHashMap.Get(TransformInfoData_GetInstance().TEMPLATEFILENAME)).(string); _ = xslFile
        if storeName != "" {
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(storeName); _ = storeFrontInterface
            var contentType string = AcceptableResponseGenerator_GetInstance().Get(this.Request); _ = contentType
            var searchParams *SearchParams = NewSearchParams(this.Request); _ = searchParams
            this.SearchRequest = NewSearchRequest(storeFrontInterface, searchParams, xslFile, contentType, propertiesHashMap, pageContext)
        }
        this.GetFormData()
    return this
}


func (this *InventoryHelper) GetFormData() {
        this.Id = this.Request.GetParameter(this.BasicItemData.ID)
}


func (this *InventoryHelper) ViewSummary() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to View Summary"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "viewSummary()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br>" + "Exception: " + fmt.Sprint(e) + "<br>"; return
                }
            }()
            var inventoryViewSearchInterface InventoryViewSearchInterface = InventoryViewFactory_GetInstance(this.SearchRequest); _ = inventoryViewSearchInterface
            var result string = inventoryViewSearchInterface.GetProduct(this.Id); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successfull Inventory View Summary", this, "viewSummary()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InventoryHelper) Search() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to search Inventory by keyword"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "searchSingleDynamicPage()", e)
            }
            tryControl1, tryResult1 = 1, error_ + "<br>" + "Exception: " + fmt.Sprint(e) + "<br>"; return
                }
            }()
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            var keywords string = any(this.SearchRequest.GetParams().Get().Get(this.BasicItemData.KEYWORDS)).(string); _ = keywords
            if keywords == "" || strings.Compare(keywords, stringUtil.EMPTY_STRING) == 0 {
                tryControl1, tryResult1 = 1, "Please enter a keyword in the search field above the Search button."; return
            }
            var inventoryViewSearchInterface InventoryViewSearchInterface = InventoryViewFactory_GetInstance(this.SearchRequest); _ = inventoryViewSearchInterface
            var result string = inventoryViewSearchInterface.SearchSingleDynamicPage(); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successful Inventory Dynamic Search", this, "searchSingleDynamicPage()")
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

