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

type TableTag struct {
    PropertiesTag
    LogUtil *LogUtil
    Enabled string
    TagHelperFactoryInterface TagHelperFactoryInterface
    TagRequestHelperFactoryInterface TagHelperFactoryInterface
    RequestObject any
    Current int
    Total int
}

// The Java class TableTag has subclasses so its variables are this interface
type ITableTag interface {
    IPropertiesTag
    TableTag_() *TableTag
    SetIsSelected(enabled string)
    IsEnabled() bool
    SetTagHelperFactory(tagHelperFactoryInterface TagHelperFactoryInterface)
    GetTagHelperFactoryInterface() TagHelperFactoryInterface
    SetTagRequestHelperFactory(tagHelperFactoryInterface TagHelperFactoryInterface)
    GetTagRequestHelperFactoryInterface() TagHelperFactoryInterface
    SetCurrent(current int)
    GetCurrent() int
    SetTotal(total int)
    GetTotal() int
    Drop() string
    Create() string
    Restore() string
    Backup() string
    Insert() string
    Delete() string
    Update() string
    DoStartTag() int
}

func (this *TableTag) TableTag_() *TableTag {
    return this
}


func (this *TableTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}


func NewTableTag() *TableTag {
    this := &TableTag{}
    this.PropertiesTag = *NewPropertiesTag()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *TableTag) SetIsSelected(enabled string) {
        this.Enabled = enabled
}


// throws Exception
func (this *TableTag) IsEnabled() bool {
        if !StringValidationUtil_GetInstance().IsEmpty(this.Enabled) {
            if BooleanUtil_GetInstance().GetFromString(this.Enabled) || strings.Compare(this.Enabled, "on") == 0 {
                return true
            }
        }
        return false
}


func (this *TableTag) SetTagHelperFactory(tagHelperFactoryInterface TagHelperFactoryInterface) {
        this.TagHelperFactoryInterface = tagHelperFactoryInterface
}


func (this *TableTag) GetTagHelperFactoryInterface() TagHelperFactoryInterface {
        return this.TagHelperFactoryInterface
}


func (this *TableTag) SetTagRequestHelperFactory(tagHelperFactoryInterface TagHelperFactoryInterface) {
        this.TagRequestHelperFactoryInterface = tagHelperFactoryInterface
}


func (this *TableTag) GetTagRequestHelperFactoryInterface() TagHelperFactoryInterface {
        return this.TagRequestHelperFactoryInterface
}


func (this *TableTag) SetCurrent(current int) {
        this.Current = current
}


func (this *TableTag) GetCurrent() int {
        return this.Current
}


func (this *TableTag) SetTotal(total int) {
        this.Total = total
}


func (this *TableTag) GetTotal() int {
        return this.Total
}


// throws LicensingException
func (this *TableTag) Drop() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to drop table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.DROP, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var anyType any = this.GetTagHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("drop", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Create() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to create table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "create()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var anyType any = this.GetTagHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("create", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Restore() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to restore backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "restore()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var anyType any = this.GetTagHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("restore", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Backup() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to make backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var anyType any = this.GetTagHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("backup", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Insert() string {
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
            var helperClass reflect.Type = reflect.TypeOf(this.RequestObject); _ = helperClass
            var method *Method = helperClass.GetMethod("insert", nil); _ = method
            var result string = any(method.Invoke(this.RequestObject, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Delete() string {
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
            var helperClass reflect.Type = reflect.TypeOf(this.RequestObject); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.Delete, nil); _ = method
            var result string = any(method.Invoke(this.RequestObject, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *TableTag) Update() string {
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
            var helperClass reflect.Type = reflect.TypeOf(this.RequestObject); _ = helperClass
            var method *Method = helperClass.GetMethod("update", nil); _ = method
            var result string = any(method.Invoke(this.RequestObject, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *TableTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, PropertiesTag_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                var request *HttpServletRequest = any(this.PageContext.GetRequest()).(*HttpServletRequest); _ = request
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("TableTag Helper: ")
                stringBuffer.Append(reflect.TypeOf(this.GetTagHelperFactoryInterface()).String())
                stringBuffer.Append(" Request URI: ")
                stringBuffer.Append(request.GetRequestURI())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
            }
            if this.GetTagRequestHelperFactoryInterface() != nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                    this.LogUtil.PutF("TableTag RequestHelper: " + reflect.TypeOf(this.GetTagRequestHelperFactoryInterface()).String(), this, this.TagStrings.DO_START_TAG)
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                    this.LogUtil.PutF("TableTag RequestHelper: null", this, this.TagStrings.DO_START_TAG)
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("TableTag this.getCommand(): " + fmt.Sprint(this.GetCommand()), this, this.TagStrings.DO_START_TAG)
            }
            if this.IsEnabled() {
                if this.GetCommand() != nil {
                    if strings.Compare(this.GetCommand(), this.CommonStrings.INSERT) == 0 || strings.Compare(this.GetCommand(), SqlStrings_GetInstance().INSERT_INTO) == 0 || strings.Compare(this.GetCommand(), this.CommonStrings.DELETE) == 0 || strings.Compare(this.GetCommand(), this.CommonStrings.UPDATE) == 0 || strings.Compare(this.GetCommand(), SqlStrings_GetInstance().UPDATE) == 0 {
                        this.RequestObject = this.GetTagRequestHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext)
                    }
                    if strings.Compare(this.GetCommand(), this.CommonStrings.INSERT) == 0 || strings.Compare(this.GetCommand(), SqlStrings_GetInstance().INSERT_INTO) == 0 {
                        this.Insert()
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.DELETE) == 0 {
                        this.Delete()
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.UPDATE) == 0 || strings.Compare(this.GetCommand(), SqlStrings_GetInstance().UPDATE) == 0 {
                        this.Update()
                    }
                    if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.BACKUP) == 0 {
                        var output string = this.Backup(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                        tryControl1, tryResult1 = 1, PropertiesTag_SKIP_BODY; return
                    } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.RESTORE) == 0 {
                        this.GetPropertiesHashMap().Put("current", this.GetCurrent())
                        this.GetPropertiesHashMap().Put("total", this.GetTotal())
                        var output string = this.Restore(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                        tryControl1, tryResult1 = 1, PropertiesTag_SKIP_BODY; return
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.DROP) == 0 {
                        var output string = this.Drop(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                        tryControl1, tryResult1 = 1, PropertiesTag_SKIP_BODY; return
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.CREATE) == 0 {
                        var output string = this.Create(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                        tryControl1, tryResult1 = 1, PropertiesTag_SKIP_BODY; return
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.INSERT) == 0 {
                        var output string = this.Insert(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.DELETE) == 0 {
                        var output string = this.Delete(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                    } else if strings.Compare(this.GetCommand(), this.CommonStrings.UPDATE) == 0 {
                        var output string = this.Update(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                    }
                }
            }
            tryControl1, tryResult1 = 1, PropertiesTag_EVAL_BODY_INCLUDE; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

