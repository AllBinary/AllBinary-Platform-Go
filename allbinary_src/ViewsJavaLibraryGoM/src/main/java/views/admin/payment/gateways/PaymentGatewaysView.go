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

type PaymentGatewaysView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    PaymentGatewayVector IBasicArrayList
}

// Implements DomNodeInterface


func (this *PaymentGatewaysView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewPaymentGatewaysView(transformInfoInterface TransformInfoInterface, gatewayVector IBasicArrayList) *PaymentGatewaysView {
    this := &PaymentGatewaysView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.PaymentGatewayVector = gatewayVector
    return this
}


// throws Exception
func (this *PaymentGatewaysView) ToXmlNode(document *Document) *Node {
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
            var storeName string = this.GetTransformInfoInterface().GetStoreName(); _ = storeName
            var paymentGatewaysNode *Node = document.CreateElement(PaymentGatewaysData_NAME); _ = paymentGatewaysNode
            var size int = this.PaymentGatewayVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var paymentType *BasicPaymentType = any(this.PaymentGatewayVector.Get(index)).(*BasicPaymentType); _ = paymentType
                var hashMap *HashMap = paymentType.ToHashMap(); _ = hashMap
                hashMap.Put(EntryData_GetInstance().DEFAULT, PaymentTypeUtil_GetInstance().GetDefault(storeName).GetBasicPaymentType().GetName())
                var paymentGatewayNode *Node = ModDomHelper_CreateNodeWithValueNodes(document, fmt.Sprint(PaymentGatewayData_NAME), hashMap); _ = paymentGatewayNode
                paymentGatewaysNode.AppendChild(paymentGatewayNode)
            }
            tryControl1, tryResult1 = 1, paymentGatewaysNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewaysView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


func (this *PaymentGatewaysView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view payment gateways"; _ = error_
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

