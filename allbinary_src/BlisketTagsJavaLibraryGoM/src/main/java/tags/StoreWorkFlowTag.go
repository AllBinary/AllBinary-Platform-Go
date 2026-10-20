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

//and hidden field removal
type StoreWorkFlowTag struct {
    HelperTag
    LogUtil *LogUtil
    ViewName string
    ViewFile string
}

// The Java class StoreWorkFlowTag has subclasses so its variables are this interface
type IStoreWorkFlowTag interface {
    IHelperTag
    StoreWorkFlowTag_() *StoreWorkFlowTag
    SetName(value string)
    Process() int
    DoStartTag() int
}

func (this *StoreWorkFlowTag) StoreWorkFlowTag_() *StoreWorkFlowTag {
    return this
}


func (this *StoreWorkFlowTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewStoreWorkFlowTag() *StoreWorkFlowTag {
    this := &StoreWorkFlowTag{}
        this.HelperTag = *NewHelperTag(NewStoreTagWorkFlowHelperFactory())
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *StoreWorkFlowTag) SetName(value string) {
        this.ViewName = value
        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().NAME, this.ViewName)
}


// throws Exception
func (this *StoreWorkFlowTag) Process() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
            }
            var helperClass reflect.Type = reflect.TypeOf(this.GetHelper()); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.PROCESS, nil); _ = method
            var result int = any(method.Invoke(this.GetHelper(), nil)).(int); _ = result
            tryControl1, tryResult1 = 1, int(result); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *StoreWorkFlowTag) DoStartTag() int {
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
                this.LogUtil.PutF(this.CommonStrings.START, this, this.TagStrings.DO_START_TAG)
            }
            this.SetHelper()
            tryControl1, tryResult1 = 1, this.Process(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

