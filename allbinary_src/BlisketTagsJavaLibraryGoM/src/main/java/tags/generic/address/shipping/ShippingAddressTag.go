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
    StoreValidationTransformTag
    Value string
}


func (this *ShippingAddressTag) initFields() {
    
}


func NewShippingAddressTag() *ShippingAddressTag {
    this := &ShippingAddressTag{}
    this.StoreValidationTransformTag = *NewStoreValidationTransformTag()
    this.initFields()
    return this
}


// throws JspTagException
func (this *ShippingAddressTag) DoStartTag() int {
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
                if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.VIEW) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.ViewShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.EDIT) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.EditShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.NEW) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.NewShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.INSERT) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.AddShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), this.CommonStrings.DELETE) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.DeleteShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), ShippingAddressData_SELECT) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.SelectShippingAddressesView")
                } else if strings.Compare(this.GetCommand(), ShippingAddressData_SETTOBILLINGADDRESS) == 0 {
                    this.SetName("Basic Shipping Address View")
                    this.SetObjectFile("views.generic.address.shipping.SetShippingAddressesView")
                } else {
                    panic(errors.New(fmt.Sprint("No Such View Command: " + fmt.Sprint(this.GetCommand()))))
                }
                tryControl1, tryResult1 = 1, this.StoreValidationTransformTag.DoStartTag(); return
            }
            panic(errors.New(fmt.Sprint("Command Null")))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

