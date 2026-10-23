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

type OrderHistory struct {
    Order
    UserName string
    ShippingAddress *StreetAddress
    BillingAddress *StreetAddress
    Payment *Payment
    ShipDate string
    OrderDate string
    TransDate string
    CancelDate string
    Status string
    SubTotal *Money
    ShippingCost *Money
    Taxes *Money
    Total *Money
}

// Implements OrderHistoryInterface


func (this *OrderHistory) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
}




// Java overloads of new OrderHistory
func NewOrderHistory(args ...any) *OrderHistory {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[BasketInterface](args[0], pass) {
            return NewOrderHistory_1(Java_Arg[BasketInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[BasketInterface](args[0], pass) && Java_IsArg[*OrderHistory](args[1], pass) {
            return NewOrderHistory_2(Java_Arg[BasketInterface](args[0]), Java_Arg[*OrderHistory](args[1]))
        }
        if len(args) == 2 && Java_IsArg[BasketInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewOrderHistory_3(Java_Arg[BasketInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of new OrderHistory for the arguments")
}
func NewOrderHistory_1(basketInterface BasketInterface) *OrderHistory {
    this := &OrderHistory{}
        this.Order = *NewOrder(basketInterface)
        this.initFields()
    return this
}


func NewOrderHistory_2(basketInterface BasketInterface, orderHistory *OrderHistory) *OrderHistory {
    this := &OrderHistory{}
        this.Order = *NewOrder(orderHistory.GetStoreName(), basketInterface, orderHistory.GetId(), orderHistory.GetShippingMethod(), orderHistory.GetSpecial(), orderHistory.GetPaymentMethod())
        this.initFields()
        this.UserName = orderHistory.GetUserName()
        this.ShippingAddress = orderHistory.GetShippingAddress()
        this.BillingAddress = orderHistory.GetBillingAddress()
        this.Payment = orderHistory.GetPaymentInfo()
        this.ShipDate = orderHistory.GetShipDate()
        this.OrderDate = orderHistory.GetOrderDate()
        this.TransDate = orderHistory.GetTransDate()
        this.CancelDate = orderHistory.GetCancelDate()
        this.Status = orderHistory.GetStatus()
        this.SubTotal = orderHistory.GetSubTotal()
        this.ShippingCost = orderHistory.GetShippingCost()
        this.Taxes = orderHistory.GetTaxes()
        this.Total = orderHistory.GetTotal()
        this.Order.SetUserComments(any(orderHistory.GetUserComments()).(string))
        this.Order.SetUserCancelComments(any(orderHistory.GetUserCancelComments()).(string))
        this.Order.SetStoreComments(any(orderHistory.GetStoreComments()).(string))
        this.Order.SetStoreCancelComments(any(orderHistory.GetStoreCancelComments()).(string))
    return this
}


func NewOrderHistory_3(basketInterface BasketInterface, orderHistoryHashMap *HashMap) *OrderHistory {
    this := &OrderHistory{}
        this.Order = *NewOrder(any(orderHistoryHashMap.Get(StoreFrontData_GetInstance().NAME)).(string), basketInterface, any(orderHistoryHashMap.Get(OrderData_ID)).(string), any(orderHistoryHashMap.Get(ShippingMethodData_NAME)).(string), any(orderHistoryHashMap.Get(EntryData_GetInstance().SPECIAL)).(string), any(orderHistoryHashMap.Get(PaymentData_METHOD)).(string))
        this.initFields()
        this.ShippingAddress = NewStreetAddress(any(orderHistoryHashMap.Get(ShippingAddressData_NAME)).(string), any(orderHistoryHashMap.Get(ShippingAddressData_STREET)).(string), any(orderHistoryHashMap.Get(ShippingAddressData_CITY)).(string), any(orderHistoryHashMap.Get(ShippingAddressData_STATE)).(string), any(orderHistoryHashMap.Get(ShippingAddressData_CODE)).(string), any(orderHistoryHashMap.Get(ShippingAddressData_COUNTRY)).(string))
        this.BillingAddress = NewStreetAddress(any(orderHistoryHashMap.Get(BillingAddressData_NAME)).(string), any(orderHistoryHashMap.Get(BillingAddressData_STREET)).(string), any(orderHistoryHashMap.Get(BillingAddressData_CITY)).(string), any(orderHistoryHashMap.Get(BillingAddressData_STATE)).(string), any(orderHistoryHashMap.Get(BillingAddressData_CODE)).(string), any(orderHistoryHashMap.Get(BillingAddressData_COUNTRY)).(string))
        this.Payment = NewPayment(orderHistoryHashMap)
        this.UserName = any(orderHistoryHashMap.Get(UserData_USERNAME)).(string)
        this.OrderDate = any(orderHistoryHashMap.Get(OrderHistoryData_ORDERDATE)).(string)
        this.ShipDate = any(orderHistoryHashMap.Get(OrderHistoryData_SHIPPEDDATE)).(string)
        this.TransDate = any(orderHistoryHashMap.Get(OrderHistoryData_TRANSDATE)).(string)
        this.CancelDate = any(orderHistoryHashMap.Get(OrderHistoryData_CANCELDATE)).(string)
        this.Status = any(orderHistoryHashMap.Get(OrderHistoryData_STATUS)).(string)
        this.SubTotal = NewMoney(any(orderHistoryHashMap.Get(OrderHistoryData_SUBTOTAL)).(string))
        this.Total = NewMoney(any(orderHistoryHashMap.Get(OrderHistoryData_TOTAL)).(string))
        this.ShippingCost = NewMoney(any(orderHistoryHashMap.Get(OrderHistoryData_SHIPPINGCOST)).(string))
        this.Taxes = NewMoney(any(orderHistoryHashMap.Get(OrderHistoryData_TAX)).(string))
        this.Order.SetUserComments(any(orderHistoryHashMap.Get(OrderData_CUSTOMERCOMMENT)).(string))
        this.Order.SetUserCancelComments(any(orderHistoryHashMap.Get(OrderData_CUSTOMERCANCELCOMMENT)).(string))
        this.Order.SetStoreComments(any(orderHistoryHashMap.Get(OrderData_STORECOMMENT)).(string))
        this.Order.SetStoreCancelComments(any(orderHistoryHashMap.Get(OrderData_STORECANCELCOMMENT)).(string))
    return this
}


func (this *OrderHistory) SetUserName(value string) {
        this.UserName = value
}


func (this *OrderHistory) SetDefaultShippingAddress(shippingAddress *StreetAddress) {
        this.ShippingAddress = shippingAddress
}


func (this *OrderHistory) SetBillingAddress(billingAddress *StreetAddress) {
        this.BillingAddress = billingAddress
}


func (this *OrderHistory) SetPaymentInfo(payment *Payment) {
        this.Payment = payment
}


func (this *OrderHistory) SetShipDate(value string) {
        this.ShipDate = value
}


func (this *OrderHistory) SetOrderDate(value string) {
        this.OrderDate = value
}


func (this *OrderHistory) SetTransDate(value string) {
        this.TransDate = value
}


func (this *OrderHistory) SetCancelDate(value string) {
        this.CancelDate = value
}


func (this *OrderHistory) SetStatus(value string) {
        this.Status = value
}


func (this *OrderHistory) SetSubTotal(value *Money) {
        this.SubTotal = value
}


func (this *OrderHistory) SetShippingCost(value *Money) {
        this.ShippingCost = value
}


func (this *OrderHistory) SetTaxes(value *Money) {
        this.Taxes = value
}


func (this *OrderHistory) SetTotal(value *Money) {
        this.Total = value
}


func (this *OrderHistory) GetUserName() string {
        return this.UserName
}


func (this *OrderHistory) GetShippingAddress() *StreetAddress {
        return this.ShippingAddress
}


func (this *OrderHistory) GetBillingAddress() *StreetAddress {
        return this.BillingAddress
}


func (this *OrderHistory) GetPaymentInfo() *Payment {
        return this.Payment
}


func (this *OrderHistory) GetShipDate() string {
        return this.ShipDate
}


func (this *OrderHistory) GetOrderDate() string {
        return this.OrderDate
}


func (this *OrderHistory) GetTransDate() string {
        return this.TransDate
}


func (this *OrderHistory) GetCancelDate() string {
        return this.CancelDate
}


func (this *OrderHistory) GetStatus() string {
        return this.Status
}


func (this *OrderHistory) GetSubTotal() *Money {
        return this.SubTotal
}


func (this *OrderHistory) GetShippingCost() *Money {
        return this.ShippingCost
}


func (this *OrderHistory) GetTaxes() *Money {
        return this.Taxes
}


func (this *OrderHistory) GetTotal() *Money {
        return this.Total
}


func (this *OrderHistory) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(UserData_USERNAME, this.UserName)
        hashMap.Put(OrderHistoryData_ORDERDATE, this.OrderDate)
        hashMap.Put(OrderHistoryData_SHIPPEDDATE, this.ShipDate)
        hashMap.Put(OrderHistoryData_TRANSDATE, this.TransDate)
        hashMap.Put(OrderHistoryData_CANCELDATE, this.CancelDate)
        hashMap.Put(OrderHistoryData_ORDERDATEFORMATTED, fmt.Sprint(NewDate(int64(int64(this.OrderDate)))))
        hashMap.Put(OrderHistoryData_SHIPPEDDATEFORMATTED, fmt.Sprint(NewDate(int64(int64(this.ShipDate)))))
        hashMap.Put(OrderHistoryData_TRANSDATEFORMATTED, fmt.Sprint(NewDate(int64(int64(this.TransDate)))))
        hashMap.Put(OrderHistoryData_CANCELDATEFORMATTED, fmt.Sprint(NewDate(int64(int64(this.CancelDate)))))
        hashMap.Put(OrderHistoryData_STATUS, this.Status)
        hashMap.Put(OrderHistoryData_SUBTOTAL, fmt.Sprint(this.SubTotal))
        hashMap.Put(OrderHistoryData_TOTAL, fmt.Sprint(this.Total))
        hashMap.Put(OrderHistoryData_SHIPPINGCOST, fmt.Sprint(this.ShippingCost))
        hashMap.Put(OrderHistoryData_TAX, fmt.Sprint(this.Taxes))
        return hashMap
}


// throws Exception
func (this *OrderHistory) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        var keySet Set = hashMap.KeySet(); _ = keySet
        var keyArray []any = keySet.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        var node *Node = document.CreateElement(OrderHistoryData_ORDERHISTORY); _ = node
        node.AppendChild(this.Order.ToXmlNode(document))
        var billingAddressNode *Node = document.CreateElement(BillingAddressData_BILLINGADDRESS); _ = billingAddressNode
        billingAddressNode.AppendChild(this.BillingAddress.ToXmlNode(document))
        var shippingAddressNode *Node = document.CreateElement(ShippingAddressData_SHIPPINGADDRESS); _ = shippingAddressNode
        shippingAddressNode.AppendChild(this.ShippingAddress.ToXmlNode(document))
        node.AppendChild(billingAddressNode)
        node.AppendChild(shippingAddressNode)
        node.AppendChild(this.Payment.ToXmlNode(document))
        for i := int(0); i < size; i++ {
            var name string = any(keyArray[i]).(string); _ = name
            var value string = any(hashMap.Get(name)).(string); _ = value
            value = StringUtil_GetInstance().GetNonNull(value)
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        return node
}

