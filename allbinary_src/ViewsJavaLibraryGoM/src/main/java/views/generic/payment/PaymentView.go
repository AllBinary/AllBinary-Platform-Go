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

type PaymentView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}

// Implements DomNodeInterface

// The Java class PaymentView has subclasses so its variables are this interface
type IPaymentView interface {
    IHttpStoreComponentView
    PaymentView_() *PaymentView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *PaymentView) PaymentView_() *PaymentView {
    return this
}


func (this *PaymentView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPaymentView(transformInfoInterface TransformInfoInterface) *PaymentView {
    this := &PaymentView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *PaymentView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            panic(e)
                }
            }()
            var paymentNode *Node = document.CreateElement(PaymentData_PAYMENT); _ = paymentNode
            document.AppendChild(paymentNode)
            var paymentVector IBasicArrayList = PaymentEntityFactory_GetInstance().GetPaymentEntityInstance().Get(this.GetWeblisketSession().GetUserName()); _ = paymentVector
            var size int = paymentVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var payment *Payment = any(paymentVector.Get(index)).(*Payment); _ = payment
                paymentNode.AppendChild(payment.ToXmlNode(document))
            }
            tryControl1, tryResult1 = 1, paymentNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


func (this *PaymentView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view Payment"; _ = error_
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

