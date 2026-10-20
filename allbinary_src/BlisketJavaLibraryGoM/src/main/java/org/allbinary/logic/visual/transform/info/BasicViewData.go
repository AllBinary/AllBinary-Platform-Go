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

type BasicViewData struct {
    GENERICDIR string
    USERDIR string
    ADDRESSDIR string
    BILLINGADDRESSDIR string
    SHIPPINGADDRESSDIR string
    COMMERCEDIR string
    INVENTORYDIR string
    BASKETDIR string
    ITEMDIR string
    MONEYDIR string
    SHIPPINGDIR string
    ORDERDIR string
    ORDEREMAILDIR string
    ORDERHISTORYDIR string
    ORDERREVIEWDIR string
    PAYMENTDIR string
    GENERATIONDIR string
    PRODUCTSDIR string
    IGNOREXMLXSL string
    REVIEWORDERVIEWXSL string
    ORDERVIEWXSL string
    INVENTORYXSL string
    BASKETXSL string
    CHECKOUTBASKETXSL string
    MINIBASKETXSL string
    BILLINGADDRESSCHANGEFORMXSL string
    BILLINGADDRESSINPUTFORMXSL string
    BILLINGADDRESSVIEWXSL string
    SHIPPINGADDRESSCHANGEFORMXSL string
    SHIPPINGADDRESSINPUTFORMXSL string
    SHIPPINGADDRESSVIEWXSL string
    NEWTAXADDRESSXSL string
    EDITTAXADDRESSXSL string
    SHIPPINGMETHODSXSL string
    ORDERCUSTOMEREMAILVIEWXSL string
    ORDERMANAGEREMAILVIEWXSL string
    EDITCUSTOMERXSL string
    ORDERHISTORYXSL string
    PAYMENTVIEWXSL string
    SUMMARYXSL string
    PAYMENTGATEWAYOPTIONSXSL string
    MAKEPAYMENTVIEWXSL string
}

var BasicViewData_Instance *BasicViewData = NewBasicViewData()

func (this *BasicViewData) initFields() {
    this.GENERICDIR = "generic/"
    this.USERDIR = this.GENERICDIR + "user/"
    this.ADDRESSDIR = this.USERDIR + "address/"
    this.BILLINGADDRESSDIR = this.ADDRESSDIR + "billing/"
    this.SHIPPINGADDRESSDIR = this.ADDRESSDIR + "shipping/"
    this.COMMERCEDIR = this.USERDIR + "commerce/"
    this.INVENTORYDIR = this.COMMERCEDIR + "inventory/"
    this.BASKETDIR = this.INVENTORYDIR + "basket/"
    this.ITEMDIR = this.INVENTORYDIR + "item/"
    this.MONEYDIR = this.COMMERCEDIR + "money/"
    this.SHIPPINGDIR = this.COMMERCEDIR + "shipping/"
    this.ORDERDIR = this.INVENTORYDIR + "order/"
    this.ORDEREMAILDIR = this.ORDERDIR + "email/"
    this.ORDERHISTORYDIR = this.ORDERDIR + "history/"
    this.ORDERREVIEWDIR = this.ORDERDIR + "review/"
    this.PAYMENTDIR = this.MONEYDIR + "payment/"
    this.GENERATIONDIR = this.GENERICDIR + "generation/"
    this.PRODUCTSDIR = this.GENERICDIR + "products/"
    this.IGNOREXMLXSL = this.GENERICDIR + "ignoreXmlView.xsl"
    this.REVIEWORDERVIEWXSL = this.ORDERREVIEWDIR + "reviewOrderView.xsl"
    this.ORDERVIEWXSL = this.REVIEWORDERVIEWXSL
    this.INVENTORYXSL = this.INVENTORYDIR + "inventory.xsl"
    this.BASKETXSL = this.BASKETDIR + "basket.xsl"
    this.CHECKOUTBASKETXSL = this.BASKETDIR + "staticBasket.xsl"
    this.MINIBASKETXSL = this.BASKETDIR + "miniBasket.xsl"
    this.BILLINGADDRESSCHANGEFORMXSL = this.BILLINGADDRESSDIR + "billingAddressChangeForm.xsl"
    this.BILLINGADDRESSINPUTFORMXSL = this.BILLINGADDRESSDIR + "billingAddressInputForm.xsl"
    this.BILLINGADDRESSVIEWXSL = this.BILLINGADDRESSDIR + "billingAddressView.xsl"
    this.SHIPPINGADDRESSCHANGEFORMXSL = this.SHIPPINGADDRESSDIR + "shippingAddressChangeForm.xsl"
    this.SHIPPINGADDRESSINPUTFORMXSL = this.SHIPPINGADDRESSDIR + "shippingAddressInputForm.xsl"
    this.SHIPPINGADDRESSVIEWXSL = this.SHIPPINGADDRESSDIR + "shippingAddressView.xsl"
    this.NEWTAXADDRESSXSL = this.SHIPPINGADDRESSDIR + "newTaxAddressView.xsl"
    this.EDITTAXADDRESSXSL = this.SHIPPINGADDRESSDIR + "editTaxAddressView.xsl"
    this.SHIPPINGMETHODSXSL = this.SHIPPINGDIR + "shippingMethods.xsl"
    this.ORDERCUSTOMEREMAILVIEWXSL = this.ORDEREMAILDIR + "orderCustomerEmailView.xsl"
    this.ORDERMANAGEREMAILVIEWXSL = this.ORDEREMAILDIR + "orderManagerEmailView.xsl"
    this.EDITCUSTOMERXSL = this.USERDIR + "editCustomerInputForm.xsl"
    this.ORDERHISTORYXSL = this.ORDERHISTORYDIR + "orderHistory.xsl"
    this.PAYMENTVIEWXSL = this.PAYMENTDIR + "paymentView.xsl"
    this.SUMMARYXSL = this.ITEMDIR + "summary/summary.xsl"
    this.PAYMENTGATEWAYOPTIONSXSL = this.PAYMENTDIR + "paymentGatewayOptions.xsl"
    this.MAKEPAYMENTVIEWXSL = this.PAYMENTDIR + "invoice/" + "makePayment.xsl"
}


func BasicViewData_GetInstance() *BasicViewData {
        return BasicViewData_Instance
}


func NewBasicViewData() *BasicViewData {
    this := &BasicViewData{}
    this.initFields()
    return this
}

