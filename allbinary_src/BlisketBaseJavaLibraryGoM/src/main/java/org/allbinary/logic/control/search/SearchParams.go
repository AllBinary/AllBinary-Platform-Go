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

type SearchParams struct {
    ColumnsAndSearchValues *HashMap
    Order string
    SortBy string
    ListLength string
    StartPage string
    EndPage string
}


func (this *SearchParams) initFields() {
    
    
    
    
    
    
}


func NewSearchParams(request *HttpServletRequest) *SearchParams {
    this := &SearchParams{}
    this.initFields()
        this.ColumnsAndSearchValues = StdUtil_GetInstance().CreateHashMap()
        this.SetLength(request.GetParameter(SearchData_LENGTH))
        this.SetOrder(request.GetParameter(SearchData_ORDER))
        var page string = request.GetParameter(SearchData_PAGE); _ = page
        this.SetStartPage(page)
        this.SetEndPage(page)
        this.SetSortBy(request.GetParameter(SearchData_SORTBY))
        var index int = 0; _ = index
        var columnName string = request.GetParameter(fmt.Sprint(SearchData_COLUMNNAME) + "[0]"); _ = columnName
        var columnValue string = request.GetParameter(fmt.Sprint(SearchData_COLUMNVALUE) + "[0]"); _ = columnValue
        for columnName != "" {
            this.Add(columnName, columnValue)
            index++
            columnName = request.GetParameter(fmt.Sprint(SearchData_COLUMNNAME) + "[" + fmt.Sprint(index) + "]")
            columnValue = request.GetParameter(fmt.Sprint(SearchData_COLUMNVALUE) + "[" + fmt.Sprint(index) + "]")
        }
    return this
}


func (this *SearchParams) Add(column string, value string) {
        this.ColumnsAndSearchValues.Put(column, value)
}


func (this *SearchParams) SetOrder(value string) {
        this.Order = value
}


func (this *SearchParams) SetSortBy(value string) {
        this.SortBy = value
}


func (this *SearchParams) SetLength(value string) {
        this.ListLength = value
}


func (this *SearchParams) SetStartPage(value string) {
        this.StartPage = value
}


func (this *SearchParams) SetEndPage(value string) {
        this.EndPage = value
}


func (this *SearchParams) Get() *HashMap {
        return this.ColumnsAndSearchValues
}


func (this *SearchParams) GetOrder() string {
        return this.Order
}


func (this *SearchParams) GetSortBy() string {
        return this.SortBy
}


func (this *SearchParams) GetLength() string {
        return this.ListLength
}


func (this *SearchParams) GetLengthInt() int {
        if this.ListLength != "" {
            return int(this.ListLength)
        } else {
            return int(0)
        }
}


func (this *SearchParams) GetStartPage() string {
        return this.StartPage
}


func (this *SearchParams) GetEndPage() string {
        return this.EndPage
}


func (this *SearchParams) GetStartPageInt() int {
        if this.StartPage != "" {
            return int(this.StartPage)
        } else {
            return int(0)
        }
}


func (this *SearchParams) GetEndPageInt() int {
        if this.EndPage != "" {
            return int(this.EndPage)
        } else {
            return int(0)
        }
}


// throws Exception
func (this *SearchParams) GetParamsNode(document *Document) *Node {
        var paramsNode *Node = ModDomHelper_CreateNameValueNodes(document, SearchData_PARAMS, SearchData_PARAMS); _ = paramsNode
        paramsNode.AppendChild(this.GetFieldsNode(document))
        paramsNode.AppendChild(this.GetOrderNode(document))
        paramsNode.AppendChild(this.GetSortByNode(document))
        paramsNode.AppendChild(this.GetLengthNode(document))
        return paramsNode
}


// throws Exception
func (this *SearchParams) GetFieldsNode(document *Document) *Node {
        var fieldsNode *Node = ModDomHelper_CreateNameValueNodes(document, SearchData_FIELDS, fmt.Sprint(int(this.ColumnsAndSearchValues.Size()))); _ = fieldsNode
        var set Set = this.ColumnsAndSearchValues.KeySet(); _ = set
        var searchValueArray []any = set.ToArray(); _ = searchValueArray
        var size int = len(searchValueArray); _ = size
        for index := int(0); index < size; index++ {
            var key string = any(searchValueArray[index]).(string); _ = key
            var searchValue string = any(this.ColumnsAndSearchValues.Get(key)).(string); _ = searchValue
            fieldsNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, SearchData_FIELD, key, ToDomHelper_ConvertNull(searchValue)))
        }
        return fieldsNode
}


// throws Exception
func (this *SearchParams) GetOrderNode(document *Document) *Node {
        return ModDomHelper_CreateNameValueNodes(document, SearchData_ORDER, ToDomHelper_ConvertNull(this.GetOrder()))
}


// throws Exception
func (this *SearchParams) GetSortByNode(document *Document) *Node {
        return ModDomHelper_CreateNameValueNodes(document, SearchData_SORTBY, ToDomHelper_ConvertNull(this.GetSortBy()))
}


// throws Exception
func (this *SearchParams) GetLengthNode(document *Document) *Node {
        return ModDomHelper_CreateNameValueNodes(document, SearchData_LENGTH, ToDomHelper_ConvertNull(this.GetLength()))
}

