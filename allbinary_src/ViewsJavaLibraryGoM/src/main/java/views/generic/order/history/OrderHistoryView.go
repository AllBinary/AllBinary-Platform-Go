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

type OrderHistoryView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    Request *HttpServletRequest
    Shipped string
    PartiallyShipped string
    Processing string
    Preprocessing string
    Cancelled string
}

// Implements DomNodeInterface

// The Java class OrderHistoryView has subclasses so its variables are this interface
type IOrderHistoryView interface {
    IHttpStoreComponentView
    OrderHistoryView_() *OrderHistoryView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *OrderHistoryView) OrderHistoryView_() *OrderHistoryView {
    return this
}


func (this *OrderHistoryView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}


func NewOrderHistoryView(transformInfoInterface TransformInfoInterface) *OrderHistoryView {
    this := &OrderHistoryView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.Preprocessing = this.Request.GetParameter(OrderHistoryData_PREPROCESSINGNAME)
        this.Shipped = this.Request.GetParameter(OrderHistoryData_SHIPPEDNAME)
        this.PartiallyShipped = this.Request.GetParameter(OrderHistoryData_PARTIALLYSHIPPEDNAME)
        this.Processing = this.Request.GetParameter(OrderHistoryData_PROCESSINGNAME)
        this.Cancelled = this.Request.GetParameter(OrderHistoryData_CANCELLEDNAME)
    return this
}


// throws Exception
func (this *OrderHistoryView) ToXmlNode(document *Document) *Node {
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
            var node *Node = document.CreateElement(OrderData_ORDERS); _ = node
            var orderHistoryEntity *OrderHistoryEntity = NewOrderHistoryEntity(); _ = orderHistoryEntity
            var orderReviewVector IBasicArrayList = orderHistoryEntity.GetOrders(this.GetWeblisketSession().GetUserName()); _ = orderReviewVector
            var size int = orderReviewVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var orderHistory *OrderHistory = any(orderReviewVector.Get(index)).(*OrderHistory); _ = orderHistory
                var orderHistoryNode *Node = orderHistory.ToXmlNode(document); _ = orderHistoryNode
                var orderNode *Node = document.CreateElement(orderHistory.GetPaymentMethod()); _ = orderNode
                node.AppendChild(orderHistory.ToXmlNode(document))
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Attempt to View a users order history", this, "view")
            }
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PREPROCESSINGNAME, OrderHistoryData_PREPROCESSING))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PROCESSINGNAME, OrderHistoryData_PROCESSING))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_CANCELLEDNAME, OrderHistoryData_CANCELLED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PARTIALLYSHIPPEDNAME, OrderHistoryData_PARTIALLYSHIPPED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_SHIPPEDNAME, OrderHistoryData_SHIPPED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, HttpStoreComponentView_GLOBALS2.VIEWNAME, HttpStoreComponentView_GLOBALS2.VIEW))
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *OrderHistoryView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *OrderHistoryView) View() string {
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

