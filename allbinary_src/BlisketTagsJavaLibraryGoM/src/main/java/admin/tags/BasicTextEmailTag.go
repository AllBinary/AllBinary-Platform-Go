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

type BasicTextEmailTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Subject string
    Body string
}


func (this *BasicTextEmailTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewBasicTextEmailTag() *BasicTextEmailTag {
    this := &BasicTextEmailTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *BasicTextEmailTag) SetSubject(value string) {
        this.Subject = value
}


func (this *BasicTextEmailTag) SetBody(value string) {
        this.Body = value
}


func (this *BasicTextEmailTag) Send() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to Send Email."; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "send()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var propertiesHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = propertiesHashMap
            propertiesHashMap.Put("Subject", this.Subject)
            propertiesHashMap.Put("Body", this.Body)
            var anyType any = NewBasicTextEmailHelperFactory().GetInstance(propertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("send", nil); _ = method
            method.Invoke(anyType, nil)
            tryControl1, tryResult1 = 1, "Email Sent"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *BasicTextEmailTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                }
            }()
            this.PageContext.GetOut().Print(this.Send())
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

