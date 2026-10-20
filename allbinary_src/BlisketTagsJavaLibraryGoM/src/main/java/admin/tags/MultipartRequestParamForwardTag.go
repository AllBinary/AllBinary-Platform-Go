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

type MultipartRequestParamForwardTag struct {
    CustomTagSupport
    Page string
    Command string
    ParamVector IBasicArrayList
}


func (this *MultipartRequestParamForwardTag) initFields() {
    
    
    
}

func NewMultipartRequestParamForwardTag() *MultipartRequestParamForwardTag {
    this := &MultipartRequestParamForwardTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *MultipartRequestParamForwardTag) SetPage(page string) {
        this.Page = page
}


func (this *MultipartRequestParamForwardTag) SetCommand(command string) {
        this.Command = command
}


func (this *MultipartRequestParamForwardTag) SetParamVector(paramVector IBasicArrayList) {
        this.ParamVector = paramVector
}


// throws JspTagException
func (this *MultipartRequestParamForwardTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, Tag_SKIP_BODY; return
                }
            }()
            var parentTag IHelperTag = any(this.GetParent()).(IHelperTag); _ = parentTag
            ParentInventoryTagHelper_GetInstance().IsValid(this, parentTag)
            var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
            hashMap.Put(AbTagData_PARENT, parentTag)
            var multipartRequestParamForwardHelper *MultipartRequestParamForwardHelper = NewMultipartRequestParamForwardHelper(parentTag, this.Page, this.Command, this.ParamVector, this.PageContext); _ = multipartRequestParamForwardHelper
            multipartRequestParamForwardHelper.Forward()
            tryControl1, tryResult1 = 1, Tag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

