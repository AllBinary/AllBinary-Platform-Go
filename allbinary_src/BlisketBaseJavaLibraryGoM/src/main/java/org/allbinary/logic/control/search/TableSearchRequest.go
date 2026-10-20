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

type TableSearchRequest struct {
    XslFile string
    PageName string
    ContentType string
    Keywords string
    SearchParams *SearchParams
}


func (this *TableSearchRequest) initFields() {
    
    
    
    
    
}


func NewTableSearchRequest(keywords string, searchParams *SearchParams, xslFile string, contentType string) *TableSearchRequest {
    this := &TableSearchRequest{}
    this.initFields()
        this.Keywords = keywords
        this.SearchParams = searchParams
        this.XslFile = xslFile
        this.PageName = this.PageName
        this.ContentType = contentType
    return this
}


func (this *TableSearchRequest) SetKeywords(value string) {
        this.Keywords = value
}


func (this *TableSearchRequest) SetParams(searchParams *SearchParams) {
        this.SearchParams = searchParams
}


func (this *TableSearchRequest) SetXslFile(value string) {
        this.XslFile = value
}


func (this *TableSearchRequest) SetFileBaseName(value string) {
        this.PageName = value
}


func (this *TableSearchRequest) SetContentType(value string) {
        this.ContentType = value
}


func (this *TableSearchRequest) GetKeywords() string {
        return this.Keywords
}


func (this *TableSearchRequest) GetParams() *SearchParams {
        return this.SearchParams
}


// throws Exception
func (this *TableSearchRequest) GetXslFile() string {
        var path string = fmt.Sprint(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH) + this.XslFile; _ = path
        return path
}


func (this *TableSearchRequest) GetFileBaseName() string {
        return this.PageName
}


func (this *TableSearchRequest) GetContentType() string {
        return this.ContentType
}


// throws Exception
func (this *TableSearchRequest) GetFileBaseNameNode(document *Document) *Node {
        return ModDomHelper_CreateNameValueNodes(document, SearchData_FILENAME, this.GetFileBaseName())
}

