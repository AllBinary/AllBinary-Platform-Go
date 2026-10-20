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

type InventoryItemViewParentTagHelper struct {
}

var InventoryItemViewParentTagHelper_Instance *InventoryItemViewParentTagHelper = NewInventoryItemViewParentTagHelper()

func NewInventoryItemViewParentTagHelper() *InventoryItemViewParentTagHelper {
    this := &InventoryItemViewParentTagHelper{}
    return this
}


func InventoryItemViewParentTagHelper_GetInstance() *InventoryItemViewParentTagHelper {
        return InventoryItemViewParentTagHelper_Instance
}


// throws Exception
func (this *InventoryItemViewParentTagHelper) GetInventoryItemView(inventoryTag IHelperTag) IInventoryItemView {
        if inventoryTag != nil {
            this.CheckForValidationViewHelper(inventoryTag.GetHelper())
            var validationViewHelper *ValidationViewHelper = any(inventoryTag.GetHelper()).(*ValidationViewHelper); _ = validationViewHelper
            var viewObject TransformInterface = validationViewHelper.GetViewObject(); _ = viewObject
            if viewObject.GetTypeId() != InventoryItemView_TYPE_ID {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Must have ")
                stringBuffer.Append("InventoryItemView and not: ")
                stringBuffer.Append(reflect.TypeOf(inventoryTag.GetHelper()).String())
                panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
            }
            return any(viewObject).(IInventoryItemView)
        } else {
            panic(errors.New(fmt.Sprint("Must have parent tag.")))
        }
}


// throws Exception
func (this *InventoryItemViewParentTagHelper) GetDownloadableInventoryItemView(inventoryTag IHelperTag) IDownloadableInventoryItemView {
        if inventoryTag != nil {
            this.CheckForValidationViewHelper(inventoryTag.GetHelper())
            var validationViewHelper *ValidationViewHelper = any(inventoryTag.GetHelper()).(*ValidationViewHelper); _ = validationViewHelper
            var viewObject TransformInterface = validationViewHelper.GetViewObject(); _ = viewObject
            if viewObject.GetTypeId() != DownloadableInventoryItemView_TYPE_ID {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Must have ")
                stringBuffer.Append("DownloadableInventoryItemView and not: ")
                stringBuffer.Append(reflect.TypeOf(inventoryTag.GetHelper()).String())
                panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
            }
            return any(viewObject).(IDownloadableInventoryItemView)
        } else {
            panic(errors.New(fmt.Sprint("Must have parent tag.")))
        }
}


// throws Exception
func (this *InventoryItemViewParentTagHelper) GetRequestMapInterface(inventoryTag IHelperTag) RequestMapInterface {
        if inventoryTag != nil {
            this.CheckForValidationViewHelper(inventoryTag.GetHelper())
            var validationViewHelper *ValidationViewHelper = any(inventoryTag.GetHelper()).(*ValidationViewHelper); _ = validationViewHelper
            var viewObject TransformInterface = validationViewHelper.GetViewObject(); _ = viewObject
            if viewObject.GetTypeId() != InventoryItemView_TYPE_ID && viewObject.GetTypeId() != DownloadableInventoryItemView_TYPE_ID {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Must have ")
                stringBuffer.Append("specific TYPE_ID and not: ")
                stringBuffer.Append(reflect.TypeOf(inventoryTag.GetHelper()).String())
                panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
            }
            return any(viewObject).(RequestMapInterface)
        } else {
            panic(errors.New(fmt.Sprint("Must have parent tag.")))
        }
}


// throws Exception
func (this *InventoryItemViewParentTagHelper) CheckForValidationViewHelper(helperObject any) {
        if !(func() bool { _, ok := any(helperObject).(*ValidationViewHelper); return ok }()) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Must have ")
            stringBuffer.Append("ValidationViewHelper")
            stringBuffer.Append("and not: ")
            stringBuffer.Append(reflect.TypeOf(helperObject).String())
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
}

