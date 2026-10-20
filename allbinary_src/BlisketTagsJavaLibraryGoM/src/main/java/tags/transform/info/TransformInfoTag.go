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

type TransformInfoTag struct {
    ModifyTag
    LogUtil *LogUtil
    Name string
    StoreName string
    ObjectFile string
    AnyType string
    ObjectConfigFile string
    ObjectConfig string
    TemplateFile string
    Template string
    DataFile string
    Data string
    Type string
}

// The Java class TransformInfoTag has subclasses so its variables are this interface
type ITransformInfoTag interface {
    IModifyTag
    TransformInfoTag_() *TransformInfoTag
    SetName(value string)
    SetStoreName(value string)
    SetObjectFile(value string)
    SetObject(value string)
    SetTemplateFile(value string)
    SetTemplate(value string)
    SetDataFile(value string)
    SetData(value string)
    SetObjectConfigFile(value string)
    SetObjectConfig(value string)
    SetType(value string)
    GetName() string
    GetObjectFile() string
    DoStartTag() int
}

func (this *TransformInfoTag) TransformInfoTag_() *TransformInfoTag {
    return this
}


func (this *TransformInfoTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
}




// Java overloads of new TransformInfoTag
func NewTransformInfoTag(args ...any) *TransformInfoTag {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTransformInfoTag_1()
        }
        if len(args) == 1 && Java_IsArg[TagHelperFactoryInterface](args[0], pass) {
            return NewTransformInfoTag_2(Java_Arg[TagHelperFactoryInterface](args[0]))
        }
    }
    panic("No Java overload of new TransformInfoTag for the arguments")
}
func NewTransformInfoTag_1() *TransformInfoTag {
    this := &TransformInfoTag{}
        this.ModifyTag = *NewModifyTag(NewTransformInfoRequestHelperFactory())
        this.initFields()
    return this
}


func NewTransformInfoTag_2(tagHelperFactoryInterface TagHelperFactoryInterface) *TransformInfoTag {
    this := &TransformInfoTag{}
        this.ModifyTag = *NewModifyTag(tagHelperFactoryInterface)
        this.initFields()
    return this
}


func (this *TransformInfoTag) SetName(value string) {
        this.Name = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().NAME, this.Name)
}


func (this *TransformInfoTag) SetStoreName(value string) {
        this.StoreName = value
        this.GetPropertiesHashMap().Put(StoreFrontData_GetInstance().NAME, this.StoreName)
}


func (this *TransformInfoTag) SetObjectFile(value string) {
        this.ObjectFile = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().OBJECTFILENAME, this.ObjectFile)
}


func (this *TransformInfoTag) SetObject(value string) {
        this.AnyType = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().OBJECT, this.AnyType)
}


func (this *TransformInfoTag) SetTemplateFile(value string) {
        this.TemplateFile = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().TEMPLATEFILENAME, this.TemplateFile)
}


func (this *TransformInfoTag) SetTemplate(value string) {
        this.Template = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().TEMPLATE, this.Template)
}


func (this *TransformInfoTag) SetDataFile(value string) {
        this.DataFile = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().DATAFILENAME, this.DataFile)
}


func (this *TransformInfoTag) SetData(value string) {
        this.Data = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().DATA, this.Data)
}


func (this *TransformInfoTag) SetObjectConfigFile(value string) {
        this.ObjectConfigFile = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().OBJECTCONFIGFILENAME, this.ObjectConfigFile)
}


func (this *TransformInfoTag) SetObjectConfig(value string) {
        this.ObjectConfig = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().OBJECTCONFIG, this.ObjectConfig)
}


func (this *TransformInfoTag) SetType(value string) {
        this.Type = value
        this.GetPropertiesHashMap().Put(OutputTypeData_GetInstance().NAME, this.Type)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("TransformInfoTag type set: " + this.Type, this, "setType")
        }
}


func (this *TransformInfoTag) GetName() string {
        return this.Name
}


func (this *TransformInfoTag) GetObjectFile() string {
        return this.ObjectFile
}


// throws JspTagException
func (this *TransformInfoTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, ModifyTag_SKIP_BODY; return
                }
            }()
            tryControl1, tryResult1 = 1, this.ModifyTag.DoStartTag(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

