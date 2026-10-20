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

type StoreFrontOrdersHistoryStatistics struct {
    NumberOfOrdersLong int64
    SubTotalMoney *Money
    ShippingCostMoney *Money
    TaxesMoney *Money
    TotalMoney *Money
}

// Implements StoreFrontOrdersHistoryStatisticsInterface


func (this *StoreFrontOrdersHistoryStatistics) initFields() {
    
    
    
    
    
}


func NewStoreFrontOrdersHistoryStatistics(storeFrontInterface StoreFrontInterface) *StoreFrontOrdersHistoryStatistics {
    this := &StoreFrontOrdersHistoryStatistics{}
    this.initFields()
        this.SubTotalMoney = NewMoney()
        this.ShippingCostMoney = NewMoney()
        this.TaxesMoney = NewMoney()
        this.TotalMoney = NewMoney()
        var orderHistoryEntityInterface OrderHistoryEntityInterface = OrderHistoryEntityFactory_GetInstance(); _ = orderHistoryEntityInterface
        var orderHistoryInterfaceVector IBasicArrayList = orderHistoryEntityInterface.GetStoreOrders(storeFrontInterface); _ = orderHistoryInterfaceVector
        var numberOfOrders int64 = 0; _ = numberOfOrders
        var orderHistoryInterface OrderHistoryInterface; _ = orderHistoryInterface
        var size int = orderHistoryInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            orderHistoryInterface = any(orderHistoryInterfaceVector.Get(index)).(OrderHistoryInterface)
            numberOfOrders = numberOfOrders + 1
            this.SubTotalMoney.Add(orderHistoryInterface.GetSubTotal())
            this.ShippingCostMoney.Add(orderHistoryInterface.GetShippingCost())
            this.TaxesMoney.Add(orderHistoryInterface.GetTaxes())
            this.TotalMoney.Add(orderHistoryInterface.GetTotal())
        }
        this.NumberOfOrdersLong = int64(numberOfOrders)
    return this
}


func (this *StoreFrontOrdersHistoryStatistics) GetNumberOfOrders() int64 {
        return this.NumberOfOrdersLong
}


func (this *StoreFrontOrdersHistoryStatistics) GetSubTotal() *Money {
        return this.SubTotalMoney
}


func (this *StoreFrontOrdersHistoryStatistics) GetShippingCost() *Money {
        return this.ShippingCostMoney
}


func (this *StoreFrontOrdersHistoryStatistics) GetTaxes() *Money {
        return this.TaxesMoney
}


func (this *StoreFrontOrdersHistoryStatistics) GetTotal() *Money {
        return this.TotalMoney
}


func (this *StoreFrontOrdersHistoryStatistics) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var storeFrontOrdersHistoryStatisticsData *StoreFrontOrdersHistoryStatisticsData = StoreFrontOrdersHistoryStatisticsData_GetInstance(); _ = storeFrontOrdersHistoryStatisticsData
        hashMap.Put(storeFrontOrdersHistoryStatisticsData.NUMBEROFORDERS, fmt.Sprint(this.GetNumberOfOrders()))
        hashMap.Put(storeFrontOrdersHistoryStatisticsData.SUBTOTAL, fmt.Sprint(this.GetSubTotal()))
        hashMap.Put(storeFrontOrdersHistoryStatisticsData.SHIPPINGCOST, fmt.Sprint(this.GetShippingCost()))
        hashMap.Put(storeFrontOrdersHistoryStatisticsData.TAXES, fmt.Sprint(this.GetTaxes()))
        hashMap.Put(storeFrontOrdersHistoryStatisticsData.TOTAL, fmt.Sprint(this.GetTotal()))
        return hashMap
}


func (this *StoreFrontOrdersHistoryStatistics) ToVector() IBasicArrayList {
        return nil
}


func (this *StoreFrontOrdersHistoryStatistics) GetKey() any {
        return nil
}

