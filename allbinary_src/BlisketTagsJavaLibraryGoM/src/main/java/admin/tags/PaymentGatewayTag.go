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

type PaymentGatewayTag struct {
    TableTag
    LogUtil *LogUtil
}


func (this *PaymentGatewayTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPaymentGatewayTag() *PaymentGatewayTag {
    this := &PaymentGatewayTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewPaymentGatewayHelperFactory())
        this.SetTagRequestHelperFactory(NewPaymentGatewayHelperFactory())
    return this
}


// throws LicensingException, Exception
func (this *PaymentGatewayTag) Process() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("LicensingException", this, this.CommonStrings.PROCESS, e)
            panic(e)
                }
            }()
            var anyType any = NewPaymentGatewayHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.PROCESS, nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *PaymentGatewayTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
                }
            }()
            if this.IsEnabled() {
                if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.AUTHORIZEORDEREVALBODYONERROR) == 0 || strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.AUTHORIZEORDERANDEVALBODY) == 0 || strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.AUTHORIZEFORMEVALBODYONERROR) == 0 {
                    var output string = this.Process(); _ = output
                    {
                        this.PageContext.GetOut().Print(output + "<br />")
                    }
                    if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.AUTHORIZEORDEREVALBODYONERROR) == 0 {
                        tryControl1, tryResult1 = 1, this.SKIP_BODY; return
                    } else {
                        tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
                    }
                } else {
                    tryControl1, tryResult1 = 1, this.TableTag.DoStartTag(); return
                }
            }
            tryControl1, tryResult1 = 1, TableTag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

