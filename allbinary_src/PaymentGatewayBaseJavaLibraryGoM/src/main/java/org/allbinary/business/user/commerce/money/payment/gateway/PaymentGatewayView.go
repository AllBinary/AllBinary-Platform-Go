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

type PaymentGatewayView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PaymentGatewayInterface PaymentGatewayInterface
}

// Implements DomNodeInterface


func (this *PaymentGatewayView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewPaymentGatewayView(paymentGatewayInterface PaymentGatewayInterface) *PaymentGatewayView {
    this := &PaymentGatewayView{}
    this.initFields()
        this.PaymentGatewayInterface = paymentGatewayInterface
    return this
}


func (this *PaymentGatewayView) GetPaymentGatewayInterface() PaymentGatewayInterface {
        return this.PaymentGatewayInterface
}


// throws Exception
func (this *PaymentGatewayView) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        return hashMap
}


// throws Exception
func (this *PaymentGatewayView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "toXmlNode")
            }
            var hashMap *HashMap = NewPaymentGatewayMapping(this.PaymentGatewayInterface).ToHashMap(); _ = hashMap
            hashMap.PutAll(this.ToHashMap())
            var paymentGatewayNode *Node = ModDomHelper_CreateNodeWithValueNodes(document, fmt.Sprint(PaymentGatewayData_NAME), hashMap); _ = paymentGatewayNode
            tryControl1, tryResult1 = 1, paymentGatewayNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

