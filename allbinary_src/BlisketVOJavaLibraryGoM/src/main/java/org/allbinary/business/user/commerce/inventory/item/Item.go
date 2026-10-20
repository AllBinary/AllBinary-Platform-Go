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

type Item struct {
    BasicItem
    Group string
    Tax string
    ShippingType string
    ShippingCost string
    ShippedDate string
    ShippingAddress *StreetAddress
    Special string
}

// Implements OrderItemInterface

// The Java class Item has subclasses so its variables are this interface
type IItem interface {
    IBasicItem
    Item_() *Item
    SetData(itemHashMap *HashMap)
    SetGroup(value string)
    SetTax(value string)
    SetShippingType(value string)
    SetShippingCost(value string)
    SetShippingAddress(value *StreetAddress)
    SetSpecial(value string)
    GetGroup() string
    GetTax() string
    GetShippingType() string
    GetShippingCost() string
    GetShippingAddress() *StreetAddress
    GetSpecial() string
    ToHashMap() *HashMap
}

func (this *Item) Item_() *Item {
    return this
}


func (this *Item) initFields() {
    
    
    
    
    
    
    
}




// Java overloads of new Item
func NewItem(args ...any) *Item {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewItem_1(Java_Arg[*HashMap](args[0]))
        }
        if len(args) == 0 {
            return NewItem_2()
        }
    }
    panic("No Java overload of new Item for the arguments")
}
func NewItem_1(itemHashMap *HashMap) *Item {
    this := &Item{}
        this.BasicItem = *NewBasicItem(itemHashMap)
        this.initFields()
        this.SetData(itemHashMap)
    return this
}


func NewItem_2() *Item {
    this := &Item{}
    this.BasicItem = *NewBasicItem()
    this.initFields()
        var EMPTY string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY
        this.Group = EMPTY
        this.Tax = EMPTY
        this.ShippingType = EMPTY
        this.ShippingCost = EMPTY
        this.ShippingAddress = NewStreetAddress()
    return this
}


func (this *Item) SetData(itemHashMap *HashMap) {
        this.Group = any(itemHashMap.Get(ShippingMethodData_GROUP)).(string)
        this.Tax = any(itemHashMap.Get(OrderHistoryData_TAX)).(string)
        this.ShippingType = any(itemHashMap.Get(ShippingMethodData_NAME)).(string)
        this.ShippingCost = any(itemHashMap.Get(ShippingMethodData_COST)).(string)
        this.ShippedDate = any(itemHashMap.Get(OrderHistoryData_SHIPPEDDATE)).(string)
        itemHashMap.Put(StreetAddressData_NAME, itemHashMap.Get(ShippingAddressData_NAME))
        itemHashMap.Put(StreetAddressData_STREET, itemHashMap.Get(ShippingAddressData_STREET))
        itemHashMap.Put(StreetAddressData_CITY, itemHashMap.Get(ShippingAddressData_CITY))
        itemHashMap.Put(StreetAddressData_STATE, itemHashMap.Get(ShippingAddressData_STATE))
        itemHashMap.Put(StreetAddressData_CODE, itemHashMap.Get(ShippingAddressData_CODE))
        itemHashMap.Put(StreetAddressData_COUNTRY, itemHashMap.Get(ShippingAddressData_COUNTRY))
        this.ShippingAddress = NewStreetAddress(itemHashMap)
}


func (this *Item) SetGroup(value string) {
        this.Group = value
}


func (this *Item) SetTax(value string) {
        this.Tax = value
}


func (this *Item) SetShippingType(value string) {
        this.ShippingType = value
}


func (this *Item) SetShippingCost(value string) {
        this.ShippingCost = value
}


func (this *Item) SetShippingAddress(value *StreetAddress) {
        this.ShippingAddress = value
}


func (this *Item) SetSpecial(value string) {
        this.Special = value
}


func (this *Item) GetGroup() string {
        return this.Group
}


func (this *Item) GetTax() string {
        return this.Tax
}


func (this *Item) GetShippingType() string {
        return this.ShippingType
}


func (this *Item) GetShippingCost() string {
        return this.ShippingCost
}


func (this *Item) GetShippingAddress() *StreetAddress {
        return this.ShippingAddress
}


func (this *Item) GetSpecial() string {
        return this.Special
}


func (this *Item) ToHashMap() *HashMap {
        var hashMap *HashMap = this.BasicItem.ToHashMap(); _ = hashMap
        hashMap.Put(ShippingMethodData_GROUP, this.Group)
        hashMap.Put(OrderHistoryData_TAX, this.Tax)
        hashMap.Put(ShippingMethodData_NAME, this.ShippingType)
        hashMap.Put(ShippingMethodData_COST, this.ShippingCost)
        hashMap.Put(OrderHistoryData_SHIPPEDDATE, this.ShippedDate)
        hashMap.Put(EntryData_GetInstance().SPECIAL, this.Special)
        return hashMap
}

