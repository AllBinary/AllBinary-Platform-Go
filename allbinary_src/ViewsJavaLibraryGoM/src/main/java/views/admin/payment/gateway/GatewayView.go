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

type GatewayView struct {
    PaymentGatewayViewAbstract
    LogUtil *LogUtil
}


func (this *GatewayView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewGatewayView(transformInfoInterface TransformInfoInterface) *GatewayView {
    this := &GatewayView{}
        this.PaymentGatewayViewAbstract = *NewPaymentGatewayViewAbstract(transformInfoInterface)
        this.initFields()
    return this
}


func (this *GatewayView) GetFormData(request *HttpServletRequest) {
        this.PaymentGatewayViewAbstract.GetFormData(request)
        if StringValidationUtil_GetInstance().IsEmpty(this.GetPaymentGatewayPrimaryKey().GetName()) {
            this.GetPaymentGatewayPrimaryKey().SetName(this.GetWeblisketSession().GetPaymentMethod())
        }
}


// throws Exception
func (this *GatewayView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "edit()", e)
            }
            panic(e)
                }
            }()
            var paymentGatewayInterface PaymentGatewayInterface = PaymentGatewayEntityFactory_GetInstance().GetPaymentGatewayInterface(this.GetPaymentGatewayPrimaryKey().GetStoreName(), BasicPaymentTypeUtil_GetInstance().Get(this.GetPaymentGatewayPrimaryKey().GetName())); _ = paymentGatewayInterface
            tryControl1, tryResult1 = 1, NewPaymentGatewayComponent(this.GetTransformInfoInterface(), paymentGatewayInterface).View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

