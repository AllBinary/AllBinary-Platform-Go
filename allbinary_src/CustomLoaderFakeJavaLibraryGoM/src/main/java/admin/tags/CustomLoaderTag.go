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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setCustomLoaderWebappPath()", e)
            }
                }
            }()
            var urlGlobals *AppUrlGlobals = NewAppUrlGlobals(); _ = urlGlobals
            urlGlobals.SetWebappPath(this.WebappPath)
            CustomTagSupport_URLGLOBALS.Init(urlGlobals)
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
                if strings.Compare(this.Command, CustomTagSupport_GLOBALS2.SET) == 0 {
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

