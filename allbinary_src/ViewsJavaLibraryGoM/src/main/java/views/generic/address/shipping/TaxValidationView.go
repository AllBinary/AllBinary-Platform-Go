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

type TaxValidationView struct {
    ShippingAddressView
}

// Implements ValidationComponentInterface



func NewTaxValidationView(transformInfoInterface TransformInfoInterface) *TaxValidationView {
    this := &TaxValidationView{}
        this.ShippingAddressView = *NewShippingAddressView(transformInfoInterface)
    return this
}


func (this *TaxValidationView) IsValid() bool {
        var billingAddressesEntity *ShippingAddressesEntity = NewShippingAddressesEntity(this.GetWeblisketSession().GetUserName()); _ = billingAddressesEntity
        var streetAddressList IBasicArrayList = billingAddressesEntity.Get().(IBasicArrayList); _ = streetAddressList
        if streetAddressList == nil {
            return false
        }
        var size int = streetAddressList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var aStreetAddress *StreetAddress = any(streetAddressList.Get(index)).(*StreetAddress); _ = aStreetAddress
            if strings.Compare(aStreetAddress.GetName(), ShippingAddressData_TAX) == 0 {
                this.StreetAddress = aStreetAddress
                return true
            }
        }
        return false
}


func (this *TaxValidationView) ValidationInfo() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *TaxValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *TaxValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

