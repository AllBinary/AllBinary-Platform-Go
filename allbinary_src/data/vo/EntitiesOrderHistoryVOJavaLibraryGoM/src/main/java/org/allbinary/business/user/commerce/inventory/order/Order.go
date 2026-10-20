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

type Order struct {
    Basket BasketInterface
    PaymentMethod string
    ShippingType string
    StoreName string
    Special string
    Id string
    UserComment string
    UserCancelComment string
    StoreComment string
    StoreCancelComment string
}

// Implements OrderInterface, *Serializable


func (this *Order) initFields() {
    
    
    
    
    
    
    
    
    
    
}


func NewOrder(basket BasketInterface) *Order {
    this := &Order{}
    this.initFields()
        this.Basket = basket
        this.Id = nil
        this.ShippingType = StringUtil_GetInstance().EMPTY_STRING
    return this
}


// Java overloaded constructor: Go needs a different name
func NewOrder(storeName string, basket BasketInterface, id string, shippingType string, special string, paymentMethod string) *Order {
    this := &Order{}
    this.initFields()
        this.StoreName = storeName
        this.Basket = basket
        this.Id = id
        this.ShippingType = shippingType
        this.Special = special
        this.PaymentMethod = paymentMethod
    return this
}


func (this *Order) IsIdValid() bool {
        if this.GetId() == nil || this.GetId().CompareTo(StringUtil_GetInstance().EMPTY_STRING) == 0 || this.GetId().Length() > OrderHistoryData_MAXIDLEN || !StringValidationUtil_GetInstance().IsNumber(this.GetId()) {
            return false
        }
        return true
}


func (this *Order) IsValid() bool {
        return true
}


func (this *Order) GetIdValidationInfo() string {
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        if this.GetId() == nil || this.GetId().CompareTo(EMPTY_STRING) == 0 || this.GetId().Length() > OrderHistoryData_MAXIDLEN || !StringValidationUtil_GetInstance().IsNumber(this.GetId()) {
            return "Please enter a valid order Id < " + fmt.Sprint(OrderHistoryData_MAXIDLEN) + " digits in length.<br>"
        }
        return EMPTY_STRING
}


func (this *Order) SetShippingMethod(shippingType string) {
        this.ShippingType = shippingType
}


func (this *Order) SetPaymentMethod(value string) {
        this.PaymentMethod = value
}


func (this *Order) SetSpecial(special string) {
        this.Special = special
}


func (this *Order) SetId(value string) {
        this.Id = value
}


func (this *Order) SetStoreName(value string) {
        this.StoreName = value
}


func (this *Order) SetUserComments(value string) {
        this.UserComment = value
}


func (this *Order) SetUserCancelComments(value string) {
        this.UserCancelComment = value
}


func (this *Order) SetStoreComments(value string) {
        this.StoreComment = value
}


func (this *Order) SetStoreCancelComments(value string) {
        this.StoreCancelComment = value
}


func (this *Order) GetBasket() BasketInterface {
        return this.Basket
}


func (this *Order) GetShippingMethod() string {
        return this.ShippingType
}


func (this *Order) GetPaymentMethod() string {
        return this.PaymentMethod
}


func (this *Order) GetSpecial() string {
        return this.Special
}


func (this *Order) GetId() string {
        return this.Id
}


func (this *Order) GetStoreName() string {
        return this.StoreName
}


func (this *Order) GetUserComments() string {
        return this.UserComment
}


func (this *Order) GetUserCancelComments() string {
        return this.UserCancelComment
}


func (this *Order) GetStoreComments() string {
        return this.StoreComment
}


func (this *Order) GetStoreCancelComments() string {
        return this.StoreCancelComment
}


func (this *Order) ReviewOrder(paymentMethod string) string {
        return "Create your own Order.reviewOrder"
}


func (this *Order) ViewOrder(paymentMethod string) string {
        return "Create your own Order.viewOrder"
}


func (this *Order) EmailOrder(userName string, paymentMethod string, storeFrontInterface StoreFrontInterface) string {
        return "Create your own Order.emailOrder"
}


func (this *Order) ToHashMapOrder() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(PaymentData_METHOD, this.PaymentMethod)
        hashMap.Put(ShippingMethodData_NAME, this.ShippingType)
        hashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
        hashMap.Put(EntryData_GetInstance().SPECIAL, this.Special)
        hashMap.Put(OrderData_ID, this.Id)
        hashMap.Put(OrderData_CUSTOMERCOMMENT, this.UserComment)
        hashMap.Put(OrderData_CUSTOMERCANCELCOMMENT, this.UserCancelComment)
        hashMap.Put(OrderData_STORECOMMENT, this.StoreComment)
        hashMap.Put(OrderData_STORECANCELCOMMENT, this.StoreCancelComment)
        return hashMap
}


// throws Exception
func (this *Order) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMapOrder(); _ = hashMap
        var keySet Set = hashMap.KeySet(); _ = keySet
        var keyArray []any = keySet.ToArray(); _ = keyArray
        var size int = len(keyArray); _ = size
        var node *Node = document.CreateElement(OrderData_ORDER); _ = node
        for i := int(0); i < size; i++ {
            var name string = any(keyArray[i]).(string); _ = name
            var value string = any(hashMap.Get(name)).(string); _ = value
            value = StringUtil_GetInstance().GetNonNull(value)
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        return node
}

