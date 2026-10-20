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

type CustomLoaderTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Command string
    WebappPath string
    PropertiesHashMap *HashMap
}


func (this *CustomLoaderTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewCustomLoaderTag() *CustomLoaderTag {
    this := &CustomLoaderTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *CustomLoaderTag) SetCommand(command string) {
        this.Command = command
}


func (this *CustomLoaderTag) SetWebappPath(value string) {
        this.WebappPath = value
}


// throws LicensingException
func (this *CustomLoaderTag) SetCustomLoaderWebappPath() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewCustomLoaderHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var methodParams []reflect.Type = []reflect.Type{reflect.TypeOf(this.WebappPath)}; _ = methodParams
            var methodArgs []any = []any{this.WebappPath}; _ = methodArgs
            var setMethod *Method = helperClass.GetMethod("setWebappPath", methodParams); _ = setMethod
            setMethod.Invoke(anyType, methodArgs)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws JspTagException
func (this *CustomLoaderTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.TagStrings.DO_START_TAG)
            }
            if this.Command != "" {
                this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
                if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.SET) == 0 {
                    this.SetCustomLoaderWebappPath()
                }
            }
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

