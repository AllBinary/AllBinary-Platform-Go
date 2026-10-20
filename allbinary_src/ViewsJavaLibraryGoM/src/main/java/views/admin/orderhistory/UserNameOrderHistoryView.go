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

type UserNameOrderHistoryView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    Request *HttpServletRequest
    UserName string
    Shipped string
    PartiallyShipped string
    Processing string
    Preprocessing string
    Cancelled string
    ON string
}

// Implements ValidationComponentInterface, DomNodeInterface


func (this *UserNameOrderHistoryView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    this.ON = "on"
}


func NewUserNameOrderHistoryView(transformInfoInterface TransformInfoInterface) *UserNameOrderHistoryView {
    this := &UserNameOrderHistoryView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.UserName = this.Request.GetParameter(UserData_USERNAME)
        this.Preprocessing = this.Request.GetParameter(OrderHistoryData_PREPROCESSINGNAME)
        this.Shipped = this.Request.GetParameter(OrderHistoryData_SHIPPEDNAME)
        this.PartiallyShipped = this.Request.GetParameter(OrderHistoryData_PARTIALLYSHIPPEDNAME)
        this.Processing = this.Request.GetParameter(OrderHistoryData_PROCESSINGNAME)
        this.Cancelled = this.Request.GetParameter(OrderHistoryData_CANCELLEDNAME)
    return this
}


func (this *UserNameOrderHistoryView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *UserNameOrderHistoryView) View() string {
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


func (this *UserNameOrderHistoryView) ToXmlNode(document *Document) *Node {
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
            var node *Node = document.CreateElement(OrderData_ORDERS); _ = node
            var orderHistoryEntity *OrderHistoryEntity = NewOrderHistoryEntity(); _ = orderHistoryEntity
            var orderReviewVector *BasicArrayList = orderHistoryEntity.GetOrders(this.UserName); _ = orderReviewVector
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


func (this *UserNameOrderHistoryView) IsValid() bool {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Started", this, this.CommonStrings.IS_VALID)
        }
        if UserName_GetInstance().IsValid(this.UserName) == true {
            return true
        }
        return false
}


func (this *UserNameOrderHistoryView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *UserNameOrderHistoryView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *UserNameOrderHistoryView) ValidationInfo() string {
        if this.UserName == "" {
            return "No User Name Specified<br />"
        }
        if UserName_GetInstance().IsValid(this.UserName) == false {
            return "Invalid User Name<br />"
        }
        return "Unknown Error<br />"
}

