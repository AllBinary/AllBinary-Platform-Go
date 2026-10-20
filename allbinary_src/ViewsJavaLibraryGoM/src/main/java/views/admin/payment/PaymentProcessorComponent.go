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

type PaymentProcessorComponent struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}


func (this *PaymentProcessorComponent) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPaymentProcessorComponent(transformInfoInterface TransformInfoInterface) *PaymentProcessorComponent {
    this := &PaymentProcessorComponent{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *PaymentProcessorComponent) AddDomNodeInterfaces() {
        var paymentProcessorInterface PaymentProcessorInterface = PaymentProcessorInterfaceFactory_GetInstance().GetInstance(this.GetTransformInfoInterface()); _ = paymentProcessorInterface
        paymentProcessorInterface.Process()
        var httpTransformInfoInterface TransformInfoHttpInterface = any(this.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        httpTransformInfoInterface.GetWeblisketSession().RemoveBasket()
        this.AddDomNodeInterface(any(paymentProcessorInterface).(DomNodeInterface))
}


// throws Exception
func (this *PaymentProcessorComponent) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
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

