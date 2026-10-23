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

type BasketView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    Request *HttpServletRequest
}

// Implements DomNodeInterface

// The Java class BasketView has subclasses so its variables are this interface
type IBasketView interface {
    IHttpStoreComponentView
    BasketView_() *BasketView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *BasketView) BasketView_() *BasketView {
    return this
}


func (this *BasketView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewBasketView(transformInfoInterface TransformInfoInterface) *BasketView {
    this := &BasketView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
    return this
}


// throws Exception
func (this *BasketView) ToXmlNode(document *Document) *Node {
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
            var inventoryEntity *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntity
            var basketInterface BasketInterface = this.GetWeblisketSession().GetOrder().GetBasket(); _ = basketInterface
            var basketNode *Node = document.CreateElement(BasketData_BASKET); _ = basketNode
            var itemsAndNumberInBasket *HashMap = basketInterface.GetItems(); _ = itemsAndNumberInBasket
            var numberOfResults int = 1; _ = numberOfResults
            var items Set = itemsAndNumberInBasket.KeySet(); _ = items
            var productArray []any = items.ToArray(); _ = productArray
            var size int = len(productArray); _ = size
            for index := int(0); index < size; index++ {
                var product string = any(productArray[index]).(string); _ = product
                var itemInterface ItemInterface = inventoryEntity.GetItem(product); _ = itemInterface
                if itemInterface != nil {
                    var basicItemView *BasicItemView = NewBasicItemView(itemInterface, NewBasicArrayListD()); _ = basicItemView
                    var node *Node = basicItemView.ToXmlNode(document); _ = node
                    var numberInBasket string = fmt.Sprint(basketInterface.GetNumberOf(product)); _ = numberInBasket
                    node.AppendChild(ModDomHelper_CreateNameValueNodes(document, BasketData_ITEMTOTALINBASKET, numberInBasket))
                    var itemPrice *Money = itemInterface.GetPrice(); _ = itemPrice
                    var itemTotal *Money = NewMoney(itemPrice); _ = itemTotal
                    itemTotal.Multiply(int(basketInterface.GetNumberOf(product)))
                    node.AppendChild(ModDomHelper_CreateNameValueNodes(document, BasketData_ITEMTOTAL, fmt.Sprint(itemTotal)))
                    basketNode.AppendChild(node)
                } else {
                    this.LogUtil.PutF("Product Failed: " + product, this, "toXmlNode")
                }
                numberOfResults++
            }
            var totalNumberNode *Node = document.CreateElement(SearchData_TOTAL_NUMBER_ITEMS); _ = totalNumberNode
            var totalNumberTextNode *Node = document.CreateTextNode(fmt.Sprint(basketInterface.GetNumberOfItems())); _ = totalNumberTextNode
            totalNumberNode.AppendChild(totalNumberTextNode)
            var totalWeightNode *Node = document.CreateElement(BasketData_TOTALWEIGHT); _ = totalWeightNode
            var totalWeightTextNode *Node = document.CreateTextNode(basketInterface.GetTotalWeight()); _ = totalWeightTextNode
            totalWeightNode.AppendChild(totalWeightTextNode)
            var subTotalNode *Node = document.CreateElement(BasketData_SUBTOTAL); _ = subTotalNode
            var subTotalTextNode *Node = document.CreateTextNode(fmt.Sprint(basketInterface.GetSubTotal())); _ = subTotalTextNode
            subTotalNode.AppendChild(subTotalTextNode)
            basketNode.AppendChild(totalNumberNode)
            basketNode.AppendChild(totalWeightNode)
            basketNode.AppendChild(subTotalNode)
            tryControl1, tryResult1 = 1, basketNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasketView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *BasketView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
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

