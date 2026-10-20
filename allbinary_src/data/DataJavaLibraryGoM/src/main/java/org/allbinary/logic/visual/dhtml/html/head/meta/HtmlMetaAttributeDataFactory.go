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

type HtmlMetaAttributeDataFactory struct {
    NAME *HtmlMetaAttributeData
    HTTP_EQUIV *HtmlMetaAttributeData
    CONTENT *HtmlMetaAttributeData
}

var HtmlMetaAttributeDataFactory_Instance *HtmlMetaAttributeDataFactory = NewHtmlMetaAttributeDataFactory()

func (this *HtmlMetaAttributeDataFactory) initFields() {
    this.NAME = NewHtmlMetaAttributeData(0, "HTMLMETA_ATTRIBUTE_NAME")
    this.HTTP_EQUIV = NewHtmlMetaAttributeData(1, "HTMLMETA_ATTRIBUTE_HTTPEQUIV")
    this.CONTENT = NewHtmlMetaAttributeData(2, "HTMLMETA_ATTRIBUTE_CONTENT")
}


func HtmlMetaAttributeDataFactory_GetInstance() *HtmlMetaAttributeDataFactory {
        return HtmlMetaAttributeDataFactory_Instance
}


func NewHtmlMetaAttributeDataFactory() *HtmlMetaAttributeDataFactory {
    this := &HtmlMetaAttributeDataFactory{}
    this.initFields()
    return this
}

