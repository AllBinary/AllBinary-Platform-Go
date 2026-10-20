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

type StaticPagesTag struct {
    TableTag
    LogUtil *LogUtil
    XslFile string
}


func (this *StaticPagesTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewStaticPagesTag() *StaticPagesTag {
    this := &StaticPagesTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewStaticPagesHelperFactory())
        this.SetTagRequestHelperFactory(NewStaticPagesRequestHelperFactory())
    return this
}


func (this *StaticPagesTag) SetXsl(value string) {
        this.XslFile = value
}


// throws LicensingException
func (this *StaticPagesTag) GenerateStaticPages() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewStaticPagesRequestHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var method *Method = helperClass.GetMethod("generateStaticPages", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *StaticPagesTag) MakePublic() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewStaticPagesRequestHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var method *Method = helperClass.GetMethod("makePublic", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *StaticPagesTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
                }
            }()
            if this.IsEnabled() {
                if this.GetCommand() != nil {
                    if strings.Compare(this.GetCommand(), SearchData_GENERATESTATICPAGES) == 0 {
                        this.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().TEMPLATEFILENAME, this.XslFile)
                        var output string = this.GenerateStaticPages(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output)
                        }
                    } else if strings.Compare(this.GetCommand(), SearchData_MAKEPUBLIC) == 0 {
                        var output string = this.MakePublic(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output)
                        }
                    } else {
                        tryControl1, tryResult1 = 1, this.TableTag.DoStartTag(); return
                    }
                }
            }
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

