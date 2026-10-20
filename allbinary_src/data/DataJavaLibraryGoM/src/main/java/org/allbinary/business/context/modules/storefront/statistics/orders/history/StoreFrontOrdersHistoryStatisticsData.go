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

type StoreFrontOrdersHistoryStatisticsData struct {
    NAME string
    NUMBEROFORDERS string
    SUBTOTAL string
    SHIPPINGCOST string
    TAXES string
    TOTAL string
}

var StoreFrontOrdersHistoryStatisticsData_Instance *StoreFrontOrdersHistoryStatisticsData = NewStoreFrontOrdersHistoryStatisticsData()

func (this *StoreFrontOrdersHistoryStatisticsData) initFields() {
    this.NAME = "STOREFRONT_ORDERS_HISTORY_STATISTICS_NAME"
    this.NUMBEROFORDERS = "STOREFRONT_ORDERS_HISTORY_STATISTICS_NUMBER_OF_ORDERS"
    this.SUBTOTAL = "STOREFRONT_ORDERS_HISTORY_STATISTICS_SUBTOTAL"
    this.SHIPPINGCOST = "STOREFRONT_ORDERS_HISTORY_STATISTICS_SHIPPINGCOST"
    this.TAXES = "STOREFRONT_ORDERS_HISTORY_STATISTICS_TAXES"
    this.TOTAL = "STOREFRONT_ORDERS_HISTORY_STATISTICS_TOTAL"
}


func StoreFrontOrdersHistoryStatisticsData_GetInstance() *StoreFrontOrdersHistoryStatisticsData {
        return StoreFrontOrdersHistoryStatisticsData_Instance
}


func NewStoreFrontOrdersHistoryStatisticsData() *StoreFrontOrdersHistoryStatisticsData {
    this := &StoreFrontOrdersHistoryStatisticsData{}
    this.initFields()
    return this
}

