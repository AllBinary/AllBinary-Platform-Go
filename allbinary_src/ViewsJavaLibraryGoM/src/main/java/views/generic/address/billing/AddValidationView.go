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

type AddValidationView struct {
    BillingAddressView
}

// Implements ValidationComponentInterface



func NewAddValidationView(transformInfoInterface TransformInfoInterface) *AddValidationView {
    this := &AddValidationView{}
        this.BillingAddressView = *NewBillingAddressView(transformInfoInterface)
    return this
}


// throws Exception
func (this *AddValidationView) IsValid() bool {
        this.StreetAddress = NewStreetAddress(this.GetRequest())
        if this.StreetAddress.IsValid() == false {
            return false
        }
        return true
}


func (this *AddValidationView) ValidationInfo() string {
        return this.StreetAddress.ValidationInfo()
}


func (this *AddValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *AddValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

