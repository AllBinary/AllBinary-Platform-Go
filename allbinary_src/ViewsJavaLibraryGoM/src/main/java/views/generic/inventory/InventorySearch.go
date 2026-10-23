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

type InventorySearch struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BasicItemData *BasicItemData
    SearchRequest *SearchRequest
}

// Implements InventoryViewSearchInterface

// The Java class InventorySearch has subclasses so its variables are this interface
type IInventorySearch interface {
    InventorySearch_() *InventorySearch
    SearchSingleStaticPage() string
    SearchSingleDynamicPage() string
    Search() []string
    GetProduct(product string) string
}

func (this *InventorySearch) InventorySearch_() *InventorySearch {
    return this
}


func (this *InventorySearch) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
}


func NewInventorySearch(searchRequest *SearchRequest) *InventorySearch {
    this := &InventorySearch{}
        
        this.initFields()
        this.SearchRequest = searchRequest
    return this
}


// throws Exception
func (this *InventorySearch) SearchSingleStaticPage() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "searchSingleStaticPage", e)
            }
            panic(errors.New(fmt.Sprint("Failed retrieve Single Product Page Static")))
                }
            }()
            var storeFront StoreFrontInterface = this.SearchRequest.GetStoreFront(); _ = storeFront
            var searchParams *SearchParams = this.SearchRequest.GetParams(); _ = searchParams
            var columnValueHashMap *HashMap = searchParams.Get(); _ = columnValueHashMap
            var file string = NewStaticPagesEntity().GetFile(storeFront.GetName(), NewReplace("-", CommonSeps_GetInstance().SPACE).All(any(columnValueHashMap.Get(this.BasicItemData.KEYWORDS)).(string))); _ = file
            if StringValidationUtil_GetInstance().IsEmpty(file) {
                tryControl1, tryResult1 = 1, nil; return
            } else {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append(storeFront.GetCurrentHostName())
                stringBuffer.Append(storeFront.GetCurrentHostNamePath())
                stringBuffer.Append(storeFront.GetStaticPath())
                stringBuffer.Append(file)
                stringBuffer.Append(this.SearchRequest.GetParams().GetEndPage())
                stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
                stringBuffer.Append(InputOutputTypeData_GetInstance().DEFAULT)
                var filePath string = NewReplace(CommonSeps_GetInstance().SPACE, "%20").All(fmt.Sprint(stringBuffer)); _ = filePath
                tryControl1, tryResult1 = 1, filePath; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *InventorySearch) SearchSingleDynamicPage() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "searchSingleDynamicPage", e)
            }
            panic(errors.New(fmt.Sprint("Failed retrieve Single Product Page Dynamically")))
                }
            }()
            var str []string = this.Search(); _ = str
            tryControl1, tryResult1 = 1, str[int(this.SearchRequest.GetParams().GetStartPageInt())]; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *InventorySearch) Search() []string {
        var abeClientInformation AbeClientInformationInterface = ServiceClientInformationInterfaceFactory_GetInstance(); _ = abeClientInformation
        var inventorySearchUtil *InventorySearchUtil = InventorySearchUtil_GetInstance(); _ = inventorySearchUtil
        var vector IBasicArrayList = inventorySearchUtil.GetBasicItemIdColumn(this.SearchRequest); _ = vector
        return inventorySearchUtil.Search(abeClientInformation, this.SearchRequest, vector)
}


// throws Exception
func (this *InventorySearch) GetProduct(product string) string {
        var abeClientInformation AbeClientInformationInterface = ServiceClientInformationInterfaceFactory_GetInstance(); _ = abeClientInformation
        return InventorySearchProductUtil_GetInstance().GetProduct(abeClientInformation, this.SearchRequest, product)
}

