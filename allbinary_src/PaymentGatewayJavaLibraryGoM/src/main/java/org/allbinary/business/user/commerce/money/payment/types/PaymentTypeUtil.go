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

type PaymentTypeUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PAYMENTGATEWAYFILEPATHSTRING string
    SHIPPINGMETHODSFILENAME string
    DefaultName string
    PaymentTypeVector *BasicArrayList
}

var PaymentTypeUtil_Instance *PaymentTypeUtil = NewPaymentTypeUtil()

func (this *PaymentTypeUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    this.SHIPPINGMETHODSFILENAME = "paymentGatewayTypes.xml"
    this.DefaultName = nil
    this.PaymentTypeVector = NewBasicArrayListD()
}


func init() {
}


func PaymentTypeUtil_GetInstance() *PaymentTypeUtil {
        return PaymentTypeUtil_Instance
}


func NewPaymentTypeUtil() *PaymentTypeUtil {
    this := &PaymentTypeUtil{}
    this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var sep string = AbPathData_GetInstance().SEPARATOR; _ = sep
        stringBuffer.Append(sep)
        stringBuffer.Append("generic")
        stringBuffer.Append(sep)
        stringBuffer.Append("user")
        stringBuffer.Append(sep)
        stringBuffer.Append("commerce")
        stringBuffer.Append(sep)
        stringBuffer.Append("money")
        stringBuffer.Append(sep)
        stringBuffer.Append("payment")
        stringBuffer.Append(sep)
        stringBuffer.Append("gateway")
        stringBuffer.Append(sep)
        this.PAYMENTGATEWAYFILEPATHSTRING = fmt.Sprint(stringBuffer)
    return this
}


func (this *PaymentTypeUtil) Add(paymentType *PaymentType) {
        this.PaymentTypeVector.Add(paymentType)
}


// throws Exception
func (this *PaymentTypeUtil) Get(paymentTypeString string) *PaymentType {
        var size int = this.PaymentTypeVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var paymentType *PaymentType = any(this.PaymentTypeVector.Get(i)).(*PaymentType); _ = paymentType
            if strings.Compare(paymentType.GetBasicPaymentType().GetName(), paymentTypeString) == 0 {
                return paymentType
            }
        }
        panic(errors.New(fmt.Sprint("Unknown PaymentType: " + paymentTypeString)))
}


// throws Exception
func (this *PaymentTypeUtil) InitDefault(storeName string) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
            this.LogUtil.PutF("initDefault Payment: " + this.DefaultName, this, "initDefault")
        }
        if StringValidationUtil_GetInstance().IsEmpty(this.DefaultName) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetMainPath())
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
            stringBuffer.Append(storeName)
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            stringBuffer.Append(this.PAYMENTGATEWAYFILEPATHSTRING)
            var abPath *AbPath = any(NewAbPath(fmt.Sprint(stringBuffer), this.SHIPPINGMETHODSFILENAME)).(*AbPath); _ = abPath
            var data string = NewCryptFileReader(TransformInfoObjectConfigData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectConfigData_GetInstance().ENCRYPTED_EXTENSION).Get(abPath); _ = data
            var document *Document = DomDocumentHelper_Create(data); _ = document
            var nodeList *NodeList = document.GetElementsByTagName(PaymentGatewaysData_NAME); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var node *Node = nodeList.Item(index); _ = node
                var defaultPaymentGatewayNode *Node = DomSearchHelper_GetNode(EntryData_GetInstance().DEFAULT, node.GetChildNodes()); _ = defaultPaymentGatewayNode
                var defaultPaymentGatewayNodeValue string = DomNodeHelper_GetTextNodeValue(defaultPaymentGatewayNode); _ = defaultPaymentGatewayNodeValue
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                    this.LogUtil.PutF("DefaultPaymentGatewayNodeValue: " + defaultPaymentGatewayNodeValue, this, "initDefault")
                }
                if !StringValidationUtil_GetInstance().IsEmpty(defaultPaymentGatewayNodeValue) {
                    this.DefaultName = defaultPaymentGatewayNodeValue
                }
            }
        }
}


// throws Exception
func (this *PaymentTypeUtil) GetDefault(storeName string) *PaymentType {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "getDefault")
        }
        this.InitDefault(storeName)
        if !StringValidationUtil_GetInstance().IsEmpty(this.DefaultName) {
            return this.Get(this.DefaultName)
        }
        panic(errors.New(fmt.Sprint("No Default: " + this.DefaultName)))
}


func (this *PaymentTypeUtil) IsContain(paymentType *PaymentType) bool {
        return this.PaymentTypeVector.Contains(paymentType)
}


func (this *PaymentTypeUtil) Difference(a_PaymentTypeVector *BasicArrayList) *BasicArrayList {
        var diff *BasicArrayList = NewBasicArrayListD(); _ = diff
        var size int = this.PaymentTypeVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var paymentType *PaymentType = any(this.PaymentTypeVector.Get(i)).(*PaymentType); _ = paymentType
            if !a_PaymentTypeVector.Contains(paymentType) {
                diff.Add(paymentType)
            }
        }
        return diff
}

