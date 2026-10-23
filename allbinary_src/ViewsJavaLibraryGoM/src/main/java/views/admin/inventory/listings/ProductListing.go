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

type ProductListing struct {
    LogUtil *LogUtil
    Directory *Directory
    BasicItemData *BasicItemData
    StoreFronts *StoreFrontsEntity
    StaticPages *StaticPagesEntity
    Inventory *InventoryEntity
    SearchRequest *SearchRequest
    INVENTORY string
}

// Implements ProductListingInterface


func (this *ProductListing) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Directory = Directory_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    this.INVENTORY = ".InventoryView"
}


func NewProductListing(searchRequest *SearchRequest) *ProductListing {
    this := &ProductListing{}
    this.initFields()
        this.SearchRequest = searchRequest
        this.StoreFronts = NewStoreFrontsEntity()
        this.StaticPages = NewStaticPagesEntity()
        this.Inventory = NewInventoryEntity()
    return this
}


// throws Exception
func (this *ProductListing) GetHashSet(storeFront StoreFrontInterface) *HashSet {
        var inventoryColumnUtil *InventoryColumnUtil = InventoryColumnUtil_GetInstance(); _ = inventoryColumnUtil
        var keywords IBasicArrayList = inventoryColumnUtil.GetColumnWhereLike(this.Inventory, storeFront.GetCategoryPath(), this.BasicItemData.KEYWORDS); _ = keywords
        var subStoreVector IBasicArrayList = storeFront.GetSubStores(); _ = subStoreVector
        var size int = subStoreVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var subStore string = any(subStoreVector.Get(index)).(string); _ = subStore
            var substoreKeywords IBasicArrayList = inventoryColumnUtil.GetColumnWhereLike(this.Inventory, fmt.Sprint(AbPathData_GetInstance().SEPARATOR) + subStore, this.BasicItemData.CATEGORY); _ = substoreKeywords
            keywords.AddAllList(substoreKeywords)
        }
        var uniqueTokens *UniqueTokens = NewUniqueTokens(); _ = uniqueTokens
        var keywordHashSet *HashSet = uniqueTokens.GetWhithoutDashesAndSkipNumberOnlyTokens(keywords); _ = keywordHashSet
        return keywordHashSet
}


// throws Exception
func (this *ProductListing) SavePage(file string, data string) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STATICPAGEGENERATIONLOGGING) {
            this.LogUtil.PutF("Creating File: " + file, this, "generateAll()")
        }
        var newFile *AbFile = AbFile_CreateAbFile(file); _ = newFile
        if newFile.Exists() {
            newFile.Delete()
        }
        newFile.CreateNewFile()
        if newFile.Exists() {
            var idOutData *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(newFile); _ = idOutData
            idOutData.WriteBytes(data)
            idOutData.Flush()
            StreamUtil_GetInstance().Close(idOutData)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STATICPAGEGENERATIONLOGGING) {
                this.LogUtil.PutF("Wrote Total Bytes: " + fmt.Sprint(newFile.Length()), this, "generateAll()")
            }
        } else {
            panic(errors.New(fmt.Sprint("Could Not Create: " + file)))
        }
}


// throws Exception
func (this *ProductListing) Create(keywordData string, keywordFilenameHashMap *HashMap, vector IBasicArrayList, staticPath IAbPath) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var hashMap *HashMap = SpecialCharacterUtil_GetHashMap(); _ = hashMap
        hashMap.Put(CommonSeps_GetInstance().SPACE, StringUtil_GetInstance().EMPTY_STRING)
        var pageName string = NewReplace(hashMap).All(keywordData); _ = pageName
        var searchParams *SearchParams = this.SearchRequest.GetParams(); _ = searchParams
        searchParams.Add(this.BasicItemData.KEYWORDS, keywordData)
        searchParams.SetStartPage(CommonPhoneStrings_GetInstance().ZERO)
        this.SearchRequest.SetParams(searchParams)
        this.SearchRequest.SetFileBaseName(pageName)
        var abeClientInformation AbeClientInformationInterface = ServiceClientInformationInterfaceFactory_GetInstance(); _ = abeClientInformation
        var inventorySearchUtil *InventorySearchUtil = InventorySearchUtil_GetInstance(); _ = inventorySearchUtil
        var productListingPages []string = inventorySearchUtil.Search(abeClientInformation, this.SearchRequest, vector); _ = productListingPages
        for index := int(0); index < len(productListingPages); index++ {
            if productListingPages[index] == "" {
                break
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STATICPAGEGENERATIONLOGGING) {
                this.LogUtil.PutF("Saving Listing: " + fmt.Sprint(index), this, "generateAll()")
            }
            var indexStr string = StringUtil_GetInstance().EMPTY_STRING; _ = indexStr
            if index > 0 {
                indexStr = fmt.Sprint(int(index))
            }
            keywordFilenameHashMap.Put(keywordData, pageName + indexStr)
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(fmt.Sprint(staticPath))
            stringBuffer.Append(pageName)
            stringBuffer.Append(indexStr)
            stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
            stringBuffer.Append(InputOutputTypeData_GetInstance().DEFAULT)
            var file string = fmt.Sprint(stringBuffer); _ = file
            if !this.Directory.Create(staticPath) {
                panic(errors.New(fmt.Sprint("Could Not Create Directory: " + fmt.Sprint(staticPath))))
            }
            this.SavePage(file, productListingPages[index])
        }
}


func (this *ProductListing) AddStaticPageInfoToDatabase(storeFront StoreFrontInterface, keywordFilenameHashMap *HashMap) {
        var keywordHashSet Set = keywordFilenameHashMap.KeySet(); _ = keywordHashSet
        var keywordArray []any = keywordHashSet.ToArray(); _ = keywordArray
        var size int = len(keywordArray); _ = size
        for index := int(0); index < size; index++ {
            var insertVector IBasicArrayList = NewBasicArrayListD(); _ = insertVector
            var keywordData string = any(keywordArray[index]).(string); _ = keywordData
            var fileName string = any(keywordFilenameHashMap.Get(keywordData)).(string); _ = fileName
            insertVector.Add(storeFront.GetName())
            insertVector.Add(keywordData)
            insertVector.Add(fileName)
            this.StaticPages.Insert(insertVector)
        }
}


// throws Exception
func (this *ProductListing) GenerateAll_1() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var storeFront StoreFrontInterface = this.SearchRequest.GetStoreFront(); _ = storeFront
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            stringBuffer.Append(storeFront.GetName())
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            stringBuffer.Append(storeFront.GetStaticPath())
            var staticPath IAbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = staticPath
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STATICPAGEGENERATIONLOGGING) {
                this.LogUtil.PutF("Store Static Pages Path: " + fmt.Sprint(staticPath), this, "generateAll()")
            }
            var keywordHashSet *HashSet = this.GetHashSet(storeFront); _ = keywordHashSet
            var keywordArray []any = keywordHashSet.ToArray(); _ = keywordArray
            var size int = len(keywordArray); _ = size
            var keywordFilenameHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keywordFilenameHashMap
            if size == 0 {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Products For ")
                stringBuffer.Append(storeFront.GetName())
                stringBuffer.Append(" Store Not Found")
                tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            }
            var inventorySearchUtil *InventorySearchUtil = InventorySearchUtil_GetInstance(); _ = inventorySearchUtil
            var vector IBasicArrayList = inventorySearchUtil.GetBasicItemIdColumn(this.SearchRequest); _ = vector
            for index := int(0); index < size; index++ {
                var keywordData string = any(keywordArray[index]).(string); _ = keywordData
                if len(keywordData) > 1 {
                    this.Create(keywordData, keywordFilenameHashMap, vector, staticPath)
                }
            }
            this.AddStaticPageInfoToDatabase(storeFront, keywordFilenameHashMap)
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("Static Files Generated Successfully For ")
            stringBuffer.Append(storeFront.GetName())
            stringBuffer.Append(" it used packages ")
            stringBuffer.Append(storeFront.GetPackageLocation())
            stringBuffer.Append(this.INVENTORY)
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of ProductListing.generateAll
func (this *ProductListing) GenerateAll(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GenerateAll_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GenerateAll_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of ProductListing.generateAll for the arguments")
}


// throws Exception
func (this *ProductListing) GenerateAll_2(storeName string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if storeName != "" {
                if strings.Compare(storeName, GLOBALS2.GENERATEALLSTORES) != 0 {
                    this.SearchRequest.SetStoreFront(this.StoreFronts.GetStoreFrontInterface(storeName))
                    stringBuffer.Append(this.GenerateAll())
                    stringBuffer.Append("<br />")
                } else {
                    var storeFrontVector IBasicArrayList = this.StoreFronts.GetStoreFrontNames(); _ = storeFrontVector
                    var size int = storeFrontVector.Size(); _ = size
                    for index := int(0); index < size; index++ {
                        storeName = any(storeFrontVector.Get(index)).(string)
                        this.SearchRequest.SetStoreFront(this.StoreFronts.GetStoreFrontInterface(storeName))
                        stringBuffer.Append(this.GenerateAll())
                        stringBuffer.Append("<br />")
                    }
                }
            } else {
                tryControl1, tryResult1 = 1, "Generation Failed No Store Specified<br/>"; return
            }
            stringBuffer.Append("All Static Pages Generated<br/>")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

