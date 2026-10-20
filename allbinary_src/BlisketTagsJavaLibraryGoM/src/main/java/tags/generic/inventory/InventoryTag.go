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

type InventoryTag struct {
    StoreValidationTransformTag
    LogUtil *LogUtil
}


func (this *InventoryTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewInventoryTag() *InventoryTag {
    this := &InventoryTag{}
        this.StoreValidationTransformTag = *NewStoreValidationTransformTag()
        this.initFields()
    return this
}


// throws LicensingException
func (this *InventoryTag) Search() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewInventoryHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var addressHelperClass reflect.Type = reflect.TypeOf(anyType); _ = addressHelperClass
            var method *Method = addressHelperClass.GetMethod("search", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            this.PageContext.GetOut().Print(result)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws LicensingException
func (this *InventoryTag) ViewSummary() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewInventoryHelperFactory().GetInstance(this.GetPropertiesHashMap(), this.PageContext); _ = anyType
            var helperClass reflect.Type = reflect.TypeOf(anyType); _ = helperClass
            var method *Method = helperClass.GetMethod("viewSummary", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            this.PageContext.GetOut().Print(result)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws JspTagException
func (this *InventoryTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, StoreValidationTransformTag_SKIP_BODY; return
                }
            }()
            if this.GetCommand() != nil {
                if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.SEARCH) == 0 {
                    this.SetName("Basic Search Inventory View")
                    this.SetObjectFile("views.generic.inventory.AutoSearchValidationInventoryView")
                    this.Search()
                    tryControl1, tryResult1 = 1, StoreValidationTransformTag_EVAL_BODY_INCLUDE; return
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.VIEWSUMMARY) == 0 {
                    this.SetName("Basic Inventory View")
                    this.SetObjectFile("views.generic.inventory.AutoValidationInventoryView")
                    this.ViewSummary()
                    tryControl1, tryResult1 = 1, StoreValidationTransformTag_EVAL_BODY_INCLUDE; return
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.VIEW) == 0 {
                    this.SetName("Basic Inventory View")
                    this.SetObjectFile("views.generic.inventory.AutoValidationInventoryView")
                    tryControl1, tryResult1 = 1, this.StoreValidationTransformTag.DoStartTag(); return
                } else {
                    panic(errors.New(fmt.Sprint("No Such View Command: " + fmt.Sprint(this.GetCommand()))))
                }
            }
            panic(errors.New(fmt.Sprint("Command Null")))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

