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

type ViewCustomizersTag struct {
    StoreValidationTransformTag
}



func NewViewCustomizersTag() *ViewCustomizersTag {
    this := &ViewCustomizersTag{}
    this.StoreValidationTransformTag = *NewStoreValidationTransformTag()
    return this
}


// throws JspTagException
func (this *ViewCustomizersTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, StoreValidationTransformTag_SKIP_BODY; return
                }
            }()
            this.SetName("Add Validation Customizers View")
            this.SetObjectFile("views.admin.customizers.CustomizersValidationView")
            tryControl1, tryResult1 = 1, this.StoreValidationTransformTag.DoStartTag(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

