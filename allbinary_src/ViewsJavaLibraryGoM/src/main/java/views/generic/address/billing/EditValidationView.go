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

type EditValidationView struct {
    BillingAddressView
    Value string
}

// Implements ValidationComponentInterface


func (this *EditValidationView) initFields() {
    
}


func NewEditValidationView(transformInfoInterface TransformInfoInterface) *EditValidationView {
    this := &EditValidationView{}
        this.BillingAddressView = *NewBillingAddressView(transformInfoInterface)
        this.initFields()
        this.Value = this.GetPageContext().GetRequest().GetParameter(StreetAddressData_ID)
    return this
}


func (this *EditValidationView) IsValid() bool {
        if this.GetWeblisketSession().GetUserName() != nil {
            return false
        }
        var billingAddressesEntity *BillingAddressesEntity = NewBillingAddressesEntity(this.GetWeblisketSession().GetUserName()); _ = billingAddressesEntity
        this.StreetAddress = billingAddressesEntity.Get(int(this.Value))
        if this.StreetAddress == nil {
            return false
        }
        return this.StreetAddress.IsValid()
}


func (this *EditValidationView) ValidationInfo() string {
        return this.StreetAddress.ValidationInfo()
}


func (this *EditValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

