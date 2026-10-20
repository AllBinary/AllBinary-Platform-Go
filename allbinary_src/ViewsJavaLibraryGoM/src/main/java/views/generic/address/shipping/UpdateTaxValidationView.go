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

type UpdateTaxValidationView struct {
    ShippingAddressView
}

// Implements ValidationComponentInterface



func NewUpdateTaxValidationView(transformInfoInterface TransformInfoInterface) *UpdateTaxValidationView {
    this := &UpdateTaxValidationView{}
        this.ShippingAddressView = *NewShippingAddressView(transformInfoInterface)
    return this
}


// throws Exception
func (this *UpdateTaxValidationView) IsValid() bool {
        this.StreetAddress = NewStreetAddress(this.GetRequest())
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetTransformInfoInterface().GetStoreName()); _ = storeFrontInterface
        if TaxFactory_GetInstance().GetInstance(this.AbeClientInformation, storeFrontInterface).IsValid(this.StreetAddress, storeFrontInterface) == false {
            return false
        }
        var billingAddressesEntity *ShippingAddressesEntity = NewShippingAddressesEntity(this.GetWeblisketSession().GetUserName()); _ = billingAddressesEntity
        var streetAddressList *BasicArrayList = billingAddressesEntity.Get(); _ = streetAddressList
        if streetAddressList == nil {
            return false
        }
        var count int = 0; _ = count
        var size int = streetAddressList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var aStreetAddress *StreetAddress = any(streetAddressList.Get(index)).(*StreetAddress); _ = aStreetAddress
            if strings.Compare(aStreetAddress.GetName(), ShippingAddressData_TAX) == 0 {
                count++
            }
        }
        if count != 1 {
            return false
        }
        return true
}


func (this *UpdateTaxValidationView) ValidationInfo() string {
        return "Unable to update tax location."
}


func (this *UpdateTaxValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *UpdateTaxValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

