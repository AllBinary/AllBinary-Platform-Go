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

type BasicWeightShippingModule struct {
    CommonStrings *CommonStrings
    Name string
    Description string
}

// Implements ShippingInterface

// The Java class BasicWeightShippingModule has subclasses so its variables are this interface
type IBasicWeightShippingModule interface {
    BasicWeightShippingModule_() *BasicWeightShippingModule
    GetCost(order OrderInterface) *Money
    GetDescription() string
    GetName() string
}

func (this *BasicWeightShippingModule) BasicWeightShippingModule_() *BasicWeightShippingModule {
    return this
}


func (this *BasicWeightShippingModule) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Name = "Basic Shipping (5 to 7 working days)"
    this.Description = "We will soon provide a variety of shipping methods."
}


func NewBasicWeightShippingModule() *BasicWeightShippingModule {
    this := &BasicWeightShippingModule{}
    this.initFields()
    return this
}


func (this *BasicWeightShippingModule) GetCost(order OrderInterface) *Money {
        var basket BasketInterface = order.GetBasket(); _ = basket
        var money *Money = NewMoney(); _ = money
        money.Add(basket.GetTotalWeight())
        money.Multiply(1)
        money.Add("3.50")
        return money
}


func (this *BasicWeightShippingModule) GetDescription() string {
        return this.Description
}


func (this *BasicWeightShippingModule) GetName() string {
        return this.Name
}

