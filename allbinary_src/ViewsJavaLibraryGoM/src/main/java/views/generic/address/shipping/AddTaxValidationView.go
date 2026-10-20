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

type AddTaxValidationView struct {
    ShippingAddressView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *AddTaxValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewAddTaxValidationView(transformInfoInterface TransformInfoInterface) *AddTaxValidationView {
    this := &AddTaxValidationView{}
        this.ShippingAddressView = *NewShippingAddressView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *AddTaxValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.StreetAddress = NewStreetAddress(this.GetRequest())
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetTransformInfoInterface().GetStoreName()); _ = storeFrontInterface
            var taxInterface TaxModuleInterface = TaxFactory_GetInstance().GetInstance(this.AbeClientInformation, storeFrontInterface); _ = taxInterface
            if taxInterface == nil {
                tryControl1, tryResult1 = 1, false; return
            }
            if taxInterface.IsValid(this.StreetAddress, storeFrontInterface) == false {
                tryControl1, tryResult1 = 1, false; return
            }
            var billingAddressesEntity *ShippingAddressesEntity = NewShippingAddressesEntity(this.GetWeblisketSession().GetUserName()); _ = billingAddressesEntity
            var streetAddressList *BasicArrayList = billingAddressesEntity.Get(); _ = streetAddressList
            if streetAddressList == nil {
                tryControl1, tryResult1 = 1, false; return
            }
            var size int = streetAddressList.Size(); _ = size
            for index := int(0); index < size; index++ {
                var aStreetAddress *StreetAddress = any(streetAddressList.Get(index)).(*StreetAddress); _ = aStreetAddress
                if strings.Compare(aStreetAddress.GetName(), ShippingAddressData_TAX) == 0 {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AddTaxValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to create validateInfo", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Creating ValidationInfo"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetTransformInfoInterface().GetStoreName()); _ = storeFrontInterface
            var taxInterface TaxModuleInterface = TaxFactory_GetInstance().GetInstance(this.AbeClientInformation, storeFrontInterface); _ = taxInterface
            if taxInterface == nil {
                stringBuffer.Append("Unable to Load Tax Component<br/>")
            } else {
                if taxInterface.IsValid(this.StreetAddress, storeFrontInterface) == false {
                    stringBuffer.Append("Unable to validate address with Tax Component<br/>")
                }
            }
            var billingAddressesEntity *ShippingAddressesEntity = NewShippingAddressesEntity(this.GetWeblisketSession().GetUserName()); _ = billingAddressesEntity
            var streetAddressList *BasicArrayList = billingAddressesEntity.Get(); _ = streetAddressList
            var size int = streetAddressList.Size(); _ = size
            for index := int(0); index < size; index++ {
                var aStreetAddress *StreetAddress = any(streetAddressList.Get(index)).(*StreetAddress); _ = aStreetAddress
                if strings.Compare(aStreetAddress.GetName(), ShippingAddressData_TAX) == 0 {
                    stringBuffer.Append("Not a valid tax location<br/>")
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AddTaxValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *AddTaxValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

