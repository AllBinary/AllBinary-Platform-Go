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

type PaymentTag struct {
    TableTag
    LogUtil *LogUtil
    StoreName string
    PropertiesHashMap *HashMap
}


func (this *PaymentTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewPaymentTag() *PaymentTag {
    this := &PaymentTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewPaymentHelperFactory())
        this.SetTagRequestHelperFactory(NewPaymentHelperFactory())
    return this
}


func (this *PaymentTag) SetStoreName(value string) {
        this.StoreName = value
}


// throws LicensingException
func (this *PaymentTag) Select() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewPaymentHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("select", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *PaymentTag) DoStartTag() int {
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
                this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
                this.PropertiesHashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
                if this.GetCommand() != nil {
                    if strings.Compare(this.GetCommand(), PaymentData_SELECT) == 0 {
                        var output string = this.Select(); _ = output
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGEXTRAOUTPUT) {
                            this.PageContext.GetOut().Print(output + "<br />")
                        }
                        tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
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

