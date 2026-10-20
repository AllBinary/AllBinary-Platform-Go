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

type XmlValidationTransformTag struct {
    TransformTag
    LogUtil *LogUtil
}


func (this *XmlValidationTransformTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewXmlValidationTransformTag() *XmlValidationTransformTag {
    this := &XmlValidationTransformTag{}
    this.TransformTag = *NewTransformTag()
    this.initFields()
    return this
}


// throws Exception
func (this *XmlValidationTransformTag) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.IS_VALID, e)
            }
            panic(e)
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.IS_VALID, nil); _ = method
            var result bool = any(method.Invoke(this.GetHelper(), nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *XmlValidationTransformTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TransformTag_SKIP_BODY; return
                }
            }()
            this.SetHelper()
            if this.IsValid() {
                this.TransformTag.DoStartTag()
                tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
            } else {
                this.TransformTag.DoStartTag()
                tryControl1, tryResult1 = 1, this.SKIP_BODY; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

