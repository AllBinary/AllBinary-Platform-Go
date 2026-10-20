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

type OrderHistoryTag struct {
    TableTag
    LogUtil *LogUtil
    Status string
}


func (this *OrderHistoryTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewOrderHistoryTag() *OrderHistoryTag {
    this := &OrderHistoryTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewOrderHistoryHelperFactory())
        this.SetTagRequestHelperFactory(NewOrderHistoryRequestHelperFactory())
    return this
}


func (this *OrderHistoryTag) SetStatus(value string) {
        this.Status = value
}


// throws LicensingException
func (this *OrderHistoryTag) SetOrderStatus() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var request *HttpServletRequest = any(this.PageContext.GetRequest()).(*HttpServletRequest); _ = request
            var anyType any = NewOrderHistoryRequestHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            if this.Status == "" {
                var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
                var method *Method = helperClass.GetMethod("setOrderStatus", nil); _ = method
                var result string = any(method.Invoke(anyType, nil)).(string); _ = result
                tryControl1, tryResult1 = 1, result; return
            } else {
                var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
                var methodParams []reflect.Type = []reflect.Type{reflect.TypeOf(this.Status)}; _ = methodParams
                var method *Method = helperClass.GetMethod("setOrderStatus", methodParams); _ = method
                var methodArgs []any = []any{this.Status}; _ = methodArgs
                var result string = any(method.Invoke(anyType, methodArgs)).(string); _ = result
                tryControl1, tryResult1 = 1, result; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *OrderHistoryTag) DoStartTag() int {
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
                    this.GetPropertiesHashMap().Put(OrderHistoryData_STATUS, this.Status)
                    if strings.Compare(this.GetCommand(), OrderHistoryData_SETSTATUS) == 0 {
                        this.SetOrderStatus()
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

