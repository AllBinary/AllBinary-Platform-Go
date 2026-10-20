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

type OrderProcessorUtil struct {
    ORDER_PROCESSED string
}

var OrderProcessorUtil_Instance *OrderProcessorUtil = NewOrderProcessorUtil()

func (this *OrderProcessorUtil) initFields() {
    this.ORDER_PROCESSED = "Order Processed"
}

func NewOrderProcessorUtil() *OrderProcessorUtil {
    this := &OrderProcessorUtil{}
    this.initFields()
    return this
}


func OrderProcessorUtil_GetInstance() *OrderProcessorUtil {
        return OrderProcessorUtil_Instance
}


// throws Exception
func (this *OrderProcessorUtil) Process(userName string, order IOrder) string {
        var orderItems *OrderItemsEntity = NewOrderItemsEntity(); _ = orderItems
        var orderHistory *OrderHistoryEntity = NewOrderHistoryEntity(); _ = orderHistory
        order.SetId(NewOrderIdGenerator().GetNext())
        orderItems.Insert(userName, order)
        orderHistory.Insert(userName, order)
        return this.ORDER_PROCESSED
}

