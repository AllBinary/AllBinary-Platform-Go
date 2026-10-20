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

type PaymentGatewayData struct {
    Key string
}

var PaymentGatewayData_CRYPTNUM int = 5
var PaymentGatewayData_ID *PaymentGatewayData = NewPaymentGatewayData("PAYMENTGATEWAY_ID")
var PaymentGatewayData_NAME *PaymentGatewayData = NewPaymentGatewayData("PAYMENTGATEWAY_NAME")
var PaymentGatewayData_VALUE *PaymentGatewayData = NewPaymentGatewayData("PAYMENTGATEWAY_VALUE")
var PaymentGatewayData_PAYMENTTRANSACTIONINTERFACEFACTORYINTERFACE *PaymentGatewayData = NewPaymentGatewayData("PAYMENTGATEWAY_PAYMENTTRANSACTIONINTERFACEFACTORYINTERFACE")
var PaymentGatewayData_PAYMENTPROCESSORINTERFACEFACTORYINTERFACE *PaymentGatewayData = NewPaymentGatewayData("PAYMENTGATEWAY_PAYMENTPROCESSORINTERFACEFACTORYINTERFACE")
var PaymentGatewayData_MODE *PaymentGatewayData = NewPaymentGatewayData("MODE")
var PaymentGatewayData_SERVERPROTOCOL *PaymentGatewayData = NewPaymentGatewayData("SERVERPROTOCOL")
var PaymentGatewayData_SERVER *PaymentGatewayData = NewPaymentGatewayData("SERVER")
var PaymentGatewayData_SERVERPORT *PaymentGatewayData = NewPaymentGatewayData("SERVERPORT")
var PaymentGatewayData_SERVERPATH *PaymentGatewayData = NewPaymentGatewayData("SERVERPATH")
var PaymentGatewayData_TESTPROTOCOL *PaymentGatewayData = NewPaymentGatewayData("TESTSERVERPROTOCOL")
var PaymentGatewayData_TESTSERVER *PaymentGatewayData = NewPaymentGatewayData("TESTSERVER")
var PaymentGatewayData_TESTPORT *PaymentGatewayData = NewPaymentGatewayData("TESTPORT")
var PaymentGatewayData_TESTPATH *PaymentGatewayData = NewPaymentGatewayData("TESTSERVERPATH")
var PaymentGatewayData_TIMEOUT *PaymentGatewayData = NewPaymentGatewayData("TIMEOUT")
var PaymentGatewayData_PROXYPROTOCOL *PaymentGatewayData = NewPaymentGatewayData("PROXYPROTOCOL")
var PaymentGatewayData_PROXYSERVER *PaymentGatewayData = NewPaymentGatewayData("PROXYSERVER")
var PaymentGatewayData_PROXYPATH *PaymentGatewayData = NewPaymentGatewayData("PROXYPATH")
var PaymentGatewayData_PROXYUSERNAME *PaymentGatewayData = NewPaymentGatewayData("PROXYUSERNAME")
var PaymentGatewayData_PROXYPASSWORD *PaymentGatewayData = NewPaymentGatewayData("PROXYPASSWORD")
var PaymentGatewayData_PROXYPORT *PaymentGatewayData = NewPaymentGatewayData("PROXYPORT")
var PaymentGatewayData_PROXYTIMEOUT *PaymentGatewayData = NewPaymentGatewayData("PROXYTIMEOUT")
var PaymentGatewayData_SPECIAL1 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL1")
var PaymentGatewayData_SPECIAL2 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL2")
var PaymentGatewayData_SPECIAL3 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL3")
var PaymentGatewayData_SPECIAL4 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL4")
var PaymentGatewayData_SPECIAL5 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL5")
var PaymentGatewayData_SPECIAL6 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL6")
var PaymentGatewayData_SPECIAL7 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL7")
var PaymentGatewayData_SPECIAL8 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL8")
var PaymentGatewayData_SPECIAL9 *PaymentGatewayData = NewPaymentGatewayData("SPECIAL9")

func (this *PaymentGatewayData) initFields() {
    
}


func NewPaymentGatewayData(value string) *PaymentGatewayData {
    this := &PaymentGatewayData{}
    this.initFields()
        this.Key = value
    return this
}


func (this *PaymentGatewayData) ToString() string {
        return this.Key
}

func (this *PaymentGatewayData) String() string {
    return this.ToString()
}

