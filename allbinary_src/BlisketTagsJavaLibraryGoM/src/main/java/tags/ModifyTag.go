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

type ModifyTag struct {
    HelperTag
    LogUtil *LogUtil
}

// The Java class ModifyTag has subclasses so its variables are this interface
type IModifyTag interface {
    IHelperTag
    ModifyTag_() *ModifyTag
    Insert() string
    Delete() string
    Update() string
    DoStartTag() int
}

func (this *ModifyTag) ModifyTag_() *ModifyTag {
    return this
}


func (this *ModifyTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewModifyTag(tagHelperFactoryInterface TagHelperFactoryInterface) *ModifyTag {
    this := &ModifyTag{}
        this.HelperTag = *NewHelperTag(tagHelperFactoryInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws LicensingException
func (this *ModifyTag) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to insert"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insert()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = helperClass
            var method *Method = helperClass.GetMethod("insert", nil); _ = method
            var result string = any(method.Invoke(this.GetHelper(), nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *ModifyTag) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to delete"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.Delete, nil); _ = method
            var result string = any(method.Invoke(this.GetHelper(), nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *ModifyTag) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = helperClass
            var method *Method = helperClass.GetMethod("update", nil); _ = method
            var result string = any(method.Invoke(this.GetHelper(), nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *ModifyTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, HelperTag_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                var request *HttpServletRequest = any(this.PageContext.GetRequest()).(*HttpServletRequest); _ = request
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Command: ")
                stringBuffer.Append(this.GetCommand())
                stringBuffer.Append(" Request URI: ")
                stringBuffer.Append(request.GetRequestURI())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            }
            if this.GetCommand() != nil {
                this.SetHelper()
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                if strings.Compare(this.GetCommand(), commonStrings.INSERT) == 0 {
                    var output string = this.Insert(); _ = output
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                        this.PageContext.GetOut().Print(output + "<br />")
                    }
                } else if strings.Compare(this.GetCommand(), commonStrings.DELETE) == 0 {
                    var output string = this.Delete(); _ = output
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                        this.PageContext.GetOut().Print(output + "<br />")
                    }
                } else if strings.Compare(this.GetCommand(), commonStrings.UPDATE) == 0 {
                    var output string = this.Update(); _ = output
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                        this.PageContext.GetOut().Print(output + "<br />")
                    }
                }
            }
            tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

