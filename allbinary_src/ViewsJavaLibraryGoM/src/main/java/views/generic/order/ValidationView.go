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

type ValidationView struct {
    OrderView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *ValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewValidationView(transformInfoInterface TransformInfoInterface) *ValidationView {
    this := &ValidationView{}
        this.OrderView = *NewOrderView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *ValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if this.Order.IsIdValid() == false {
                tryControl1, tryResult1 = 1, false; return
            } else {
                this.OrderHistory = OrderHistoryFactory_GetInstance(this.Order.GetId())
                if this.OrderHistory == nil || this.OrderHistory.IsValid() == false {
                    tryControl1, tryResult1 = 1, false; return
                }
                var orderItems *OrderItemsEntity = NewOrderItemsEntity(); _ = orderItems
                this.BasketReview = orderItems.GetBasketReview(this.Order.GetId())
                if this.BasketReview == nil || this.BasketReview.IsValid() == false {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.Order.IsIdValid() == false {
                stringBuffer.Append(this.Order.GetIdValidationInfo())
            } else {
                this.OrderHistory = OrderHistoryFactory_GetInstance(this.Order.GetId())
                if this.OrderHistory == nil || this.OrderHistory.IsValid() == false {
                    stringBuffer.Append("Order History data error for: ")
                    stringBuffer.Append(this.Order.GetId())
                    stringBuffer.Append("<br />")
                }
                var orderItems *OrderItemsEntity = NewOrderItemsEntity(); _ = orderItems
                this.BasketReview = orderItems.GetBasketReview(this.Order.GetId())
                if this.BasketReview == nil {
                    stringBuffer.Append("Order Items data error - does not exist for: ")
                    stringBuffer.Append(this.Order.GetId())
                    stringBuffer.Append("<br />")
                } else if this.BasketReview.IsValid() == false {
                    stringBuffer.Append("Order Items data error - not valid for: ")
                    stringBuffer.Append(this.Order.GetId())
                    stringBuffer.Append("<br />")
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *ValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

