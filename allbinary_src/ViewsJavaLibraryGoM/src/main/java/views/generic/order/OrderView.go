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

type OrderView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    Request *HttpServletRequest
    Order IOrder
    OrderHistory *OrderHistory
    BasketReview *BasketReview
}

// Implements DomNodeInterface

// The Java class OrderView has subclasses so its variables are this interface
type IOrderView interface {
    IHttpStoreComponentView
    OrderView_() *OrderView
    GetFormData()
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *OrderView) OrderView_() *OrderView {
    return this
}


func (this *OrderView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
}


func NewOrderView(transformInfoInterface TransformInfoInterface) *OrderView {
    this := &OrderView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


// throws Exception
func (this *OrderView) GetFormData() {
        var id string = this.Request.GetParameter(OrderData_ID); _ = id
        if id == "" {
            id = this.GetWeblisketSession().GetOrder().GetId()
        }
        this.Order = NewOrder(NewBasket())
        this.Order.SetId(id)
}


func (this *OrderView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var orderHistoryNode *Node = this.OrderHistory.ToXmlNode(document); _ = orderHistoryNode
            var node *Node = document.CreateElement(this.OrderHistory.GetPaymentMethod()); _ = node
            orderHistoryNode.AppendChild(this.BasketReview.ToXmlNode(document))
            node.AppendChild(orderHistoryNode)
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *OrderView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *OrderView) View() string {
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

