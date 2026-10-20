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

type BasketTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Command string
    StoreName string
    PropertiesHashMap *HashMap
}


func (this *BasketTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewBasketTag() *BasketTag {
    this := &BasketTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *BasketTag) SetCommand(command string) {
        this.Command = command
}


func (this *BasketTag) SetStoreName(value string) {
        this.StoreName = value
}


// throws LicensingException
func (this *BasketTag) IsBasketEmpty() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewBasketHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("isBasketEmpty", nil); _ = method
            var emptyBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = emptyBoolean
            tryControl1, tryResult1 = 1, emptyBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *BasketTag) AddItemToBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewBasketRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("addItemToBasket", nil); _ = method
            var emptyBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = emptyBoolean
            tryControl1, tryResult1 = 1, emptyBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *BasketTag) RemoveItemFromBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewBasketRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("removeItemFromBasket", nil); _ = method
            var emptyBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = emptyBoolean
            tryControl1, tryResult1 = 1, emptyBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *BasketTag) AdjustBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewBasketRequestHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("adjustBasket", nil); _ = method
            var emptyBoolean bool = any(method.Invoke(anyType, nil)).(bool); _ = emptyBoolean
            tryControl1, tryResult1 = 1, emptyBoolean; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *BasketTag) DoStartTag() int {
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
                if strings.Compare(this.Command, BasketData_INSERT) == 0 {
                    if !this.AddItemToBasket() {
                        this.PageContext.GetOut().Println("Item is not currently being sold.<p/>")
                    }
                } else if strings.Compare(this.Command, BasketData_DELETE) == 0 {
                    this.RemoveItemFromBasket()
                } else if strings.Compare(this.Command, BasketData_ADJUST) == 0 {
                    this.AdjustBasket()
                } else if strings.Compare(this.Command, BasketData_ISEMPTY) == 0 {
                    if this.IsBasketEmpty() {
                        tryControl1, tryResult1 = 1, this.EVAL_BODY_INCLUDE; return
                    } else {
                        tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                    }
                }
            }
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

