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

type SearchRequest struct {
    AbContext
    XslFile string
    PageName string
    ContentType string
    SearchParams *SearchParams
    StoreFront StoreFrontInterface
}


func (this *SearchRequest) initFields() {
    
    
    
    
    
}


func NewSearchRequest(storeFront StoreFrontInterface, searchParams *SearchParams, xslFile string, contentType string, propertiesHashMap *HashMap, pageContext *PageContext) *SearchRequest {
    this := &SearchRequest{}
        this.AbContext = *NewAbContext(propertiesHashMap, pageContext)
        this.initFields()
        this.StoreFront = storeFront
        this.SearchParams = searchParams
        this.XslFile = xslFile
        this.PageName = this.PageName
        this.ContentType = contentType
    return this
}


func (this *SearchRequest) SetStoreFront(storeFront StoreFrontInterface) {
        this.StoreFront = storeFront
}


func (this *SearchRequest) SetParams(searchParams *SearchParams) {
        this.SearchParams = searchParams
}


func (this *SearchRequest) SetXslFile(value string) {
        this.XslFile = value
}


func (this *SearchRequest) SetFileBaseName(value string) {
        this.PageName = value
}


func (this *SearchRequest) SetContentType(value string) {
        this.ContentType = value
}


func (this *SearchRequest) GetStoreFront() StoreFrontInterface {
        return this.StoreFront
}


func (this *SearchRequest) GetParams() *SearchParams {
        return this.SearchParams
}


func (this *SearchRequest) GetXslFile() string {
        return this.XslFile
}


func (this *SearchRequest) GetFileBaseName() string {
        return this.PageName
}


func (this *SearchRequest) GetContentType() string {
        return this.ContentType
}


// throws Exception
func (this *SearchRequest) GetFileBaseNameNode(document *Document) *Node {
        return ModDomHelper_CreateNameValueNodes(document, SearchData_FILENAME, this.GetFileBaseName())
}

