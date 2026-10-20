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

type OrderTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Command string
    StoreName string
    PropertiesHashMap *HashMap
}


func (this *OrderTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewOrderTag() *OrderTag {
    this := &OrderTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *OrderTag) SetCommand(command string) {
        this.Command = command
}


func (this *OrderTag) SetStoreName(value string) {
        this.StoreName = value
}


// throws LicensingException
func (this *OrderTag) Process() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewOrderHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
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


// throws LicensingException
func (this *OrderTag) SetPaymentGateway() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewOrderHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("setPaymentGateway", nil); _ = method
            var result bool = any(method.Invoke(anyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *OrderTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                }
            }()
            if this.Command != "" {
                this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
                this.PropertiesHashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
                if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.PROCESS) == 0 {
                    this.Process()
                } else if strings.Compare(this.Command, fmt.Sprint(PaymentGatewayData_NAME)) == 0 {
                    if this.SetPaymentGateway() == true {
                        tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                    }
                }
            }
            tryControl1, tryResult1 = 1, TagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

