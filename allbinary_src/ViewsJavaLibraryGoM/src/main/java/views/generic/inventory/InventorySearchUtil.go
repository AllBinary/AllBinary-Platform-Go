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

type InventorySearchUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonPhoneStrings *CommonPhoneStrings
    BasicItemData *BasicItemData
}

var InventorySearchUtil_Instance *InventorySearchUtil = NewInventorySearchUtil()

func (this *InventorySearchUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonPhoneStrings = CommonPhoneStrings_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
}

func NewInventorySearchUtil() *InventorySearchUtil {
    this := &InventorySearchUtil{}
    this.initFields()
    return this
}


func InventorySearchUtil_GetInstance() *InventorySearchUtil {
        return InventorySearchUtil_Instance
}


// throws Exception
func (this *InventorySearchUtil) GetBasicItemIdColumn(searchRequest *SearchRequest) *BasicArrayList {
        var inventoryEntityInterface *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntityInterface
        var storeFrontInterface StoreFrontInterface = searchRequest.GetStoreFront(); _ = storeFrontInterface
        var inventorySearchUtil *InventoryColumnUtil = InventoryColumnUtil_GetInstance(); _ = inventorySearchUtil
        var column *BasicArrayList = inventorySearchUtil.GetColumnWhereLike(inventoryEntityInterface, storeFrontInterface.GetCategoryPath(), this.BasicItemData.ID); _ = column
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF("Number Of Items Found: " + fmt.Sprint(column.Size()), this, "search")
        }
        var subStoreVector *BasicArrayList = storeFrontInterface.GetSubStores(); _ = subStoreVector
        var size int = subStoreVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var subStore string = any(subStoreVector.Get(index)).(string); _ = subStore
            var substoreIdColumn *BasicArrayList = inventorySearchUtil.GetColumnWhereLike(inventoryEntityInterface, subStore, this.BasicItemData.ID); _ = substoreIdColumn
            column.AddAllList(substoreIdColumn)
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF("Number Of Items Found Including SubStores: " + fmt.Sprint(column.Size()), this, "search")
        }
        return column
}


// throws Exception
func (this *InventorySearchUtil) GetNoResults(viewDocumentInterface TransformDocumentInterface, inventoryNode *Node) string {
        inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_TOTAL_NUMBER_PAGES, this.CommonPhoneStrings.ZERO))
        inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_TOTAL_NUMBER_ITEMS, this.CommonPhoneStrings.ZERO))
        var success string = XmlDocumentHelper_ToString(viewDocumentInterface.GetDoc()); _ = success
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF("No Results Xml: " + success, this, "search")
        }
        return success
}


func (this *InventorySearchUtil) Search(abeClientInformation AbeClientInformationInterface, searchRequest *SearchRequest, column *BasicArrayList) []string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "search", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var inventoryEntityInterface *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntityInterface
            var searchParams *SearchParams = searchRequest.GetParams(); _ = searchParams
            var startPage int = int(searchParams.GetStartPageInt()); _ = startPage
            var endPage int = int(searchParams.GetEndPageInt()); _ = endPage
            var pageLength int = int(searchParams.GetLengthInt()); _ = pageLength
            var savedPagesInRange int = 0; _ = savedPagesInRange
            var columnValueHashMap *HashMap = searchParams.Get(); _ = columnValueHashMap
            var keyword string = NewReplace("-", CommonSeps_GetInstance().SPACE).All(any(columnValueHashMap.Get(this.BasicItemData.KEYWORDS)).(string)); _ = keyword
            if StringValidationUtil_GetInstance().IsEmpty(keyword) {
                tryControl1, tryResult1 = 1, nil; return
            }
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if startPage >= 0 {
                if pageLength == 0 {
                    pageLength = SearchData_DEFAULT_PAGE_LENGTH
                }
                var MAXPAGES int = 100; _ = MAXPAGES
                var productListingPages []string = make([]string, MAXPAGES); _ = productListingPages
                var documents []*Document = make([]*Document, MAXPAGES); _ = documents
                var inventoryNodes []*Node = make([]*Node, MAXPAGES); _ = inventoryNodes
                keyword = strings.ToUpper(keyword)
                var lastPage int = -1; _ = lastPage
                var numberOfResultsOnCurrentPage int = 0; _ = numberOfResultsOnCurrentPage
                var numberOfResults int = 0; _ = numberOfResults
                var viewDocumentInterface TransformDocumentInterface = TransformStoreDocumentFactory_GetInstance(searchRequest); _ = viewDocumentInterface
                var inventoryNode *Node = viewDocumentInterface.GetDoc().CreateElement(InventoryData_INVENTORY); _ = inventoryNode
                viewDocumentInterface.GetBaseNode().AppendChild(inventoryNode)
                inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_PAGE, fmt.Sprint(int(startPage))))
                if searchRequest.GetFileBaseName() != nil {
                    inventoryNode.AppendChild(searchRequest.GetFileBaseNameNode(viewDocumentInterface.GetDoc()))
                }
                var currentPage int = -1; _ = currentPage
                var size int = column.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var product string = any(column.Get(index)).(string); _ = product
                    var itemInterface ItemInterface = inventoryEntityInterface.GetItem(product); _ = itemInterface
                    var keywords string = itemInterface.GetKeywords(); _ = keywords
                    keywords = strings.ToUpper(keywords)
                    if strings.Index(keywords, string(keyword)) != -1 {
                        if currentPage + 1 >= startPage && currentPage + 1 <= endPage {
                            var itemNode *Node = NewBasicItemView(itemInterface, NewBasicArrayListD()).ToXmlNode(viewDocumentInterface.GetDoc()); _ = itemNode
                            itemNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), BasketData_ITEMTOTALINBASKET, this.CommonPhoneStrings.ONE))
                            inventoryNode.AppendChild(itemNode.CloneNode(true))
                        }
                        currentPage = ((numberOfResults + 1) / pageLength) - 1
                        numberOfResultsOnCurrentPage++
                        if lastPage < currentPage {
                            if currentPage >= startPage && currentPage <= endPage {
                                inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_TOTAL_NUMBER_ITEMS_ON_THIS_PAGE, fmt.Sprint(int(numberOfResultsOnCurrentPage))))
                                documents[currentPage] = viewDocumentInterface.GetDoc()
                                inventoryNodes[currentPage] = inventoryNode
                                savedPagesInRange++
                                viewDocumentInterface = TransformStoreDocumentFactory_GetInstance(searchRequest)
                                inventoryNode = viewDocumentInterface.GetDoc().CreateElement(InventoryData_INVENTORY)
                                viewDocumentInterface.GetBaseNode().AppendChild(inventoryNode)
                                inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_PAGE, fmt.Sprint(int(currentPage + 1))))
                                if searchRequest.GetFileBaseName() != nil {
                                    inventoryNode.AppendChild(searchRequest.GetFileBaseNameNode(viewDocumentInterface.GetDoc()))
                                }
                            }
                            lastPage = currentPage
                            numberOfResultsOnCurrentPage = 0
                        }
                        numberOfResults++
                        if lastPage >= MAXPAGES - 1 {
                            break
                        }
                    }
                }
                if numberOfResultsOnCurrentPage > 0 && (currentPage + 1 >= startPage && currentPage + 1 <= endPage) {
                    inventoryNode.AppendChild(ModDomHelper_CreateNameValueNodes(viewDocumentInterface.GetDoc(), SearchData_TOTAL_NUMBER_ITEMS_ON_THIS_PAGE, fmt.Sprint(int(numberOfResultsOnCurrentPage - 1))))
                    documents[lastPage + 1] = viewDocumentInterface.GetDoc()
                    inventoryNodes[lastPage + 1] = inventoryNode
                    savedPagesInRange++
                }
                if numberOfResultsOnCurrentPage > 0 {
                    lastPage++
                }
                if endPage > MAXPAGES - 1 {
                    endPage = MAXPAGES - 1
                }
                for index := int(startPage); index <= endPage; index++ {
                    if index <= lastPage {
                        var tempDocument *Document = documents[index]; _ = tempDocument
                        inventoryNodes[index].AppendChild(ModDomHelper_CreateNameValueNodes(tempDocument, SearchData_TOTAL_NUMBER_PAGES, fmt.Sprint(int(lastPage + 1))))
                        inventoryNodes[index].AppendChild(ModDomHelper_CreateNameValueNodes(tempDocument, SearchData_TOTAL_NUMBER_ITEMS, fmt.Sprint(int(numberOfResults))))
                        for pageIndex := int(0); pageIndex <= lastPage; pageIndex++ {
                            stringBuffer.Delete(0, stringBuffer.Length())
                            stringBuffer.Append(searchRequest.GetFileBaseName())
                            if stringBuffer.Length() > 0 {
                                if pageIndex != 0 {
                                    stringBuffer.Append(fmt.Sprint(pageIndex))
                                }
                                inventoryNodes[index].AppendChild(ModDomHelper_CreateNameValueNodes(tempDocument, SearchData_PAGE_INFO, fmt.Sprint(pageIndex), fmt.Sprint(stringBuffer)))
                            } else {
                                inventoryNodes[index].AppendChild(ModDomHelper_CreateNameValueNodes(tempDocument, SearchData_PAGE_INFO, fmt.Sprint(pageIndex), SearchData_SEARCH))
                            }
                        }
                        var success string = XmlDocumentHelper_ToString(tempDocument); _ = success
                        productListingPages[index] = NewStoreTransformer(abeClientInformation, any(NewTransformInfoHttpSearch(searchRequest)).(TransformInfoInterface)).Translate(success)
                    }
                }
                if savedPagesInRange == 0 {
                    var result string = this.GetNoResults(viewDocumentInterface, inventoryNode); _ = result
                    productListingPages[0] = NewStoreTransformer(abeClientInformation, NewTransformInfoHttpSearch(searchRequest)).Translate(result)
                }
                tryControl1, tryResult1 = 1, productListingPages; return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

