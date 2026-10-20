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

type ShippingAddressTag struct {
    TableTag
    LogUtil *LogUtil
    Value string
}


func (this *ShippingAddressTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewShippingAddressTag() *ShippingAddressTag {
    this := &ShippingAddressTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewShippingAddressHelperFactory())
        this.SetTagRequestHelperFactory(NewShippingAddressHelperFactory())
    return this
}


func (this *ShippingAddressTag) SetValue(value string) {
        this.Value = value
}


// throws LicensingException
func (this *ShippingAddressTag) Set() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewShippingAddressHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("set", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *ShippingAddressTag) SetToBillingAddress() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewShippingAddressHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("setToBillingAddress", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *ShippingAddressTag) DoStartTag() int {
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
                    if strings.Compare(this.GetCommand(), ShippingAddressData_SELECT) == 0 {
                        this.Set()
                        tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
                    } else if strings.Compare(this.GetCommand(), ShippingAddressData_SETTOBILLINGADDRESS) == 0 {
                        this.SetToBillingAddress()
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

