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

type ValidationTransformTag struct {
    TransformTag
    LogUtil *LogUtil
    XslFile string
    Logic bool
}

// The Java class ValidationTransformTag has subclasses so its variables are this interface
type IValidationTransformTag interface {
    ITransformTag
    ValidationTransformTag_() *ValidationTransformTag
    SetXsl(value string)
    SetLogic(logic bool)
    IsValid() bool
    ValidationInfo() string
    DoStartTag() int
    DoEndTag() int
}

func (this *ValidationTransformTag) ValidationTransformTag_() *ValidationTransformTag {
    return this
}


func (this *ValidationTransformTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewValidationTransformTag() *ValidationTransformTag {
    this := &ValidationTransformTag{}
        this.TransformTag = *NewTransformTag(NewValidatedViewHelperFactory())
        this.initFields()
        this.Logic = true
    return this
}


func (this *ValidationTransformTag) SetXsl(value string) {
        this.XslFile = value
        this.SetTemplateFile(this.XslFile)
}


func (this *ValidationTransformTag) SetLogic(logic bool) {
        this.Logic = logic
}


// throws Exception
func (this *ValidationTransformTag) IsValid() bool {
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


// throws Exception
func (this *ValidationTransformTag) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "validationInfo()", e)
            }
            panic(e)
                }
            }()
            var addressHelperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("validationInfo", nil); _ = method
            var result string = any(method.Invoke(this.GetHelper(), nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *ValidationTransformTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TransformTag_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                var request *HttpServletRequest = any(this.PageContext.GetRequest()).(*HttpServletRequest); _ = request
                stringBuffer.Append("ValidationViewTag Start For: ")
                stringBuffer.Append(this.GetName())
                stringBuffer.Append("\nView FIle: ")
                stringBuffer.Append(this.GetObjectFile())
                stringBuffer.Append("\nRequest URI: ")
                stringBuffer.Append(request.GetRequestURI())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
            }
            this.SetHelper()
            if this.IsValid() {
                this.TransformTag.DoStartTag()
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("View File: ")
                    stringBuffer.Append(this.GetObjectFile())
                    stringBuffer.Append("\nLogic includes body if true=")
                    stringBuffer.Appendboolean(this.Logic)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
                }
                if this.Logic {
                    tryControl1, tryResult1 = 1, TransformTag_EVAL_BODY_INCLUDE; return
                } else {
                    tryControl1, tryResult1 = 1, TransformTag_SKIP_BODY; return
                }
            } else {
                this.PageContext.GetOut().Print(this.ValidationInfo())
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("View File: ")
                    stringBuffer.Append(this.GetObjectFile())
                    stringBuffer.Append("\nisValid()=false")
                    stringBuffer.Append("\nLogic skips body if true=")
                    stringBuffer.Appendboolean(this.Logic)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
                }
                if this.Logic {
                    tryControl1, tryResult1 = 1, TransformTag_SKIP_BODY; return
                } else {
                    tryControl1, tryResult1 = 1, TransformTag_EVAL_BODY_INCLUDE; return
                }
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ValidationTransformTag) DoEndTag() int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("Tag Ended", this, "doEndTag")
        }
        this.Logic = true
        return this.TransformTag.DoEndTag()
}

