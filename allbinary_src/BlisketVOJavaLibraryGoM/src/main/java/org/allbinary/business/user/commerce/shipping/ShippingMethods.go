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

type ShippingMethods struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ShippingVector IBasicArrayList
    DefaultShippingMethodInterface ShippingInterface
}

// Implements ShippingMethodsInterface


func (this *ShippingMethods) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewShippingMethods(abeClientInformation AbeClientInformationInterface, storeFrontInterface StoreFrontInterface) *ShippingMethods {
    this := &ShippingMethods{}
    this.initFields()
        this.DefaultShippingMethodInterface = NewShippingMethodsFactory(abeClientInformation, storeFrontInterface).GetDefaultInstance()
        this.ShippingVector = NewShippingMethodsFactory(abeClientInformation, storeFrontInterface).GetInstance()
    return this
}


func (this *ShippingMethods) Get() IBasicArrayList {
        return this.ShippingVector
}


// throws Exception
func (this *ShippingMethods) GetShippingInterface(name string) ShippingInterface {
        var shipping ShippingInterface; _ = shipping
        var size int = this.ShippingVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            shipping = any(this.ShippingVector.Get(index)).(ShippingInterface)
            if strings.Compare(name, shipping.GetName()) == 0 {
                return shipping
            }
        }
        var error_ string = "Error Finding Shipping: " + name; _ = error_
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SHIPPINGERROR) {
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, "getShippingInterface(String name)")
        }
        panic(errors.New(fmt.Sprint(error_)))
}


// throws Exception
func (this *ShippingMethods) GetDefault() ShippingInterface {
        return this.DefaultShippingMethodInterface
}

