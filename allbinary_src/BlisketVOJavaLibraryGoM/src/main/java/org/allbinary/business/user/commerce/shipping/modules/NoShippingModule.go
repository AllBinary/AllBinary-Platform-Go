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

type NoShippingModule struct {
    CommonStrings *CommonStrings
    Name string
    Description string
}

// Implements ShippingInterface

// The Java class NoShippingModule has subclasses so its variables are this interface
type INoShippingModule interface {
    NoShippingModule_() *NoShippingModule
    GetCost(order OrderInterface) *Money
    GetDescription() string
    GetName() string
}

func (this *NoShippingModule) NoShippingModule_() *NoShippingModule {
    return this
}


func (this *NoShippingModule) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Name = "No Shipping"
    this.Description = "This shipping method is for items that are not shipped."
}


func NewNoShippingModule() *NoShippingModule {
    this := &NoShippingModule{}
    this.initFields()
    return this
}


func (this *NoShippingModule) GetCost(order OrderInterface) *Money {
        return NewMoney()
}


func (this *NoShippingModule) GetDescription() string {
        return this.Description
}


func (this *NoShippingModule) GetName() string {
        return this.Name
}

