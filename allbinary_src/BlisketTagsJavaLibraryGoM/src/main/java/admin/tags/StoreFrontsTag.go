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

type StoreFrontsTag struct {
    TableTag
    LogUtil *LogUtil
}


func (this *StoreFrontsTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewStoreFrontsTag() *StoreFrontsTag {
    this := &StoreFrontsTag{}
    this.TableTag = *NewTableTag()
    this.initFields()
        this.SetTagHelperFactory(NewStoreFrontsHelperFactory())
        this.SetTagRequestHelperFactory(NewStoreFrontsRequestHelperFactory())
    return this
}


// throws LicensingException
func (this *StoreFrontsTag) GenerateSelect() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = this.GetTagHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("generateSelect", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *StoreFrontsTag) Install() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = this.GetTagRequestHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("install", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *StoreFrontsTag) SendStoreCreatedEmails() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = this.GetTagRequestHelperFactoryInterface().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("sendStoreCreatedEmails", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *StoreFrontsTag) DoStartTag() int {
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
                    var storeFrontData *StoreFrontData = StoreFrontData_GetInstance(); _ = storeFrontData
                    if strings.Compare(this.GetCommand(), storeFrontData.SELECT) == 0 {
                        this.PageContext.GetOut().Print(this.GenerateSelect())
                    } else if strings.Compare(this.GetCommand(), storeFrontData.INSTALL) == 0 {
                        this.GetPropertiesHashMap().Put("current", this.GetCurrent())
                        this.GetPropertiesHashMap().Put("total", this.GetTotal())
                        if int(this.GetCurrent()) == int(this.GetTotal()) {
                            tryControl1, tryResult1 = 1, TagSupport_EVAL_BODY_INCLUDE; return
                        }
                        this.Install()
                    } else if strings.Compare(this.GetCommand(), storeFrontData.INSTALL_COMPLETE) == 0 {
                        this.SendStoreCreatedEmails()
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

