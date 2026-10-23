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

type ShippingMethodsView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    ShippingMethods *ShippingMethods
    StoreFrontInterface StoreFrontInterface
}

// Implements DomNodeInterface

// The Java class ShippingMethodsView has subclasses so its variables are this interface
type IShippingMethodsView interface {
    IHttpStoreComponentView
    ShippingMethodsView_() *ShippingMethodsView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *ShippingMethodsView) ShippingMethodsView_() *ShippingMethodsView {
    return this
}


func (this *ShippingMethodsView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewShippingMethodsView(transformInfoInterface TransformInfoInterface) *ShippingMethodsView {
    this := &ShippingMethodsView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetTransformInfoInterface().GetStoreName())
        this.ShippingMethods = NewShippingMethods(this.AbeClientInformation, this.StoreFrontInterface)
    return this
}


// throws Exception
func (this *ShippingMethodsView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            panic(e)
                }
            }()
            var shippingMethodsNode *Node = document.CreateElement(ShippingMethodsData_SHIPPINGMETHODSORDERSUMMARIES); _ = shippingMethodsNode
            var order OrderInterface = this.GetWeblisketSession().GetOrder(); _ = order
            var basket BasketInterface = this.GetWeblisketSession().GetOrder().GetBasket(); _ = basket
            var shippingAddressesEntityInterface StreetAddressesEntityInterface = ShippingAddressesEntityFactory_GetInstance().GetInstance(this.GetWeblisketSession().GetUserName()); _ = shippingAddressesEntityInterface
            var streetAddress *StreetAddress = shippingAddressesEntityInterface.GetDefault(); _ = streetAddress
            var shippingVector IBasicArrayList = this.ShippingMethods.Get(); _ = shippingVector
            var size int = shippingVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var shippingMethodNode *Node = document.CreateElement(ShippingMethodsData_SHIPPINGMETHODORDERSUMMARY); _ = shippingMethodNode
                var shipping ShippingInterface = any(shippingVector.Get(index)).(ShippingInterface); _ = shipping
                var shippingView DomNodeInterface = any(shipping).(DomNodeInterface); _ = shippingView
                shippingMethodNode.AppendChild(shippingView.ToXmlNode(document))
                var shippingCost *Money = shipping.GetCost(order); _ = shippingCost
                var subTotal *Money = basket.GetSubTotal(); _ = subTotal
                var taxRate float32 = TaxFactory_GetInstance().GetInstance(this.AbeClientInformation, this.StoreFrontInterface).GetTaxRate(streetAddress, this.StoreFrontInterface); _ = taxRate
                var tax *Money = NewMoney(); _ = tax
                var total *Money = NewMoney(); _ = total
                total.Add(fmt.Sprint(shippingCost))
                total.Add(fmt.Sprint(subTotal))
                tax.Add(fmt.Sprint(total))
                tax.Multiply(taxRate)
                total.Add(fmt.Sprint(tax))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodsData_SHIPPINGCOST, fmt.Sprint(shippingCost)))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodsData_TAXRATE, fmt.Sprint(taxRate)))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodsData_SUBTOTAL, fmt.Sprint(subTotal)))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodsData_TAX, fmt.Sprint(tax)))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, ShippingMethodsData_TOTAL, fmt.Sprint(total)))
                shippingMethodNode.AppendChild(ModDomHelper_CreateNameValueNodes(document, EntryData_GetInstance().DEFAULT, this.ShippingMethods.GetDefault().GetName()))
                shippingMethodsNode.AppendChild(shippingMethodNode)
            }
            tryControl1, tryResult1 = 1, shippingMethodsNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ShippingMethodsView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


func (this *ShippingMethodsView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view ShippingMethods"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            this.AddDomNodeInterfaces()
            tryControl1, tryResult1 = 1, this.HttpStoreComponentView.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

