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

type HtmlMetaAttributesData struct {
    NAME_VALUE string
    HTTP_EQUIV_VALUE string
}

var HtmlMetaAttributesData_Instance *HtmlMetaAttributesData = NewHtmlMetaAttributesData()

func (this *HtmlMetaAttributesData) initFields() {
    this.NAME_VALUE = "name"
    this.HTTP_EQUIV_VALUE = "http-equiv"
}


func HtmlMetaAttributesData_GetInstance() *HtmlMetaAttributesData {
        return HtmlMetaAttributesData_Instance
}


func NewHtmlMetaAttributesData() *HtmlMetaAttributesData {
    this := &HtmlMetaAttributesData{}
    this.initFields()
    return this
}

