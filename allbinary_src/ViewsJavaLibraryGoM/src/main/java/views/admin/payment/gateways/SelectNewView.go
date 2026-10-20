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

type SelectNewView struct {
    PaymentGatewayViewAbstract
    LogUtil *LogUtil
}


func (this *SelectNewView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewSelectNewView(transformInfoInterface TransformInfoInterface) *SelectNewView {
    this := &SelectNewView{}
        this.PaymentGatewayViewAbstract = *NewPaymentGatewayViewAbstract(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *SelectNewView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view", e)
            }
            panic(e)
                }
            }()
            var paymentGatewayEntityInterface PaymentGatewayEntityInterface = PaymentGatewayEntityFactory_GetInstance(); _ = paymentGatewayEntityInterface
            var existingGateways *BasicArrayList = paymentGatewayEntityInterface.FindPaymentTypeVectorByStore(this.GetPaymentGatewayPrimaryKey().GetStoreName()); _ = existingGateways
            var allGatewaysLessExisting *BasicArrayList = BasicPaymentTypeUtil_GetInstance().Difference(existingGateways); _ = allGatewaysLessExisting
            tryControl1, tryResult1 = 1, NewPaymentGatewaysView(this.GetTransformInfoInterface(), allGatewaysLessExisting).View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

