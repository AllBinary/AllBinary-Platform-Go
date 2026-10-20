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

type StoreFrontInventoryStatistics struct {
    TotalNumberOfItems int64
    TotalInventorySaleValueMoney *Money
}

// Implements StoreFrontInventoryStatisticsInterface


func (this *StoreFrontInventoryStatistics) initFields() {
    
    
}


func NewStoreFrontInventoryStatistics(storeFrontInterface StoreFrontInterface) *StoreFrontInventoryStatistics {
    this := &StoreFrontInventoryStatistics{}
    this.initFields()
        this.TotalInventorySaleValueMoney = NewMoney()
        var inventoryEntityInterface InventoryEntityInterface = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntityInterface
        var itemVector IBasicArrayList = inventoryEntityInterface.GetItems(storeFrontInterface); _ = itemVector
        var size int = itemVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var itemInterface ItemInterface = any(itemVector.Get(index)).(ItemInterface); _ = itemInterface
            var numberInStock int = int(fmt.Sprint(itemInterface.GetNumber())); _ = numberInStock
            this.TotalNumberOfItems = this.TotalNumberOfItems + int64(numberInStock)
            var itemPriceMoney *Money = itemInterface.GetPrice(); _ = itemPriceMoney
            itemPriceMoney.Multiply(numberInStock)
            this.TotalInventorySaleValueMoney.Add(fmt.Sprint(itemPriceMoney))
        }
    return this
}


func (this *StoreFrontInventoryStatistics) GetNumber() int64 {
        return int64(this.TotalNumberOfItems)
}


func (this *StoreFrontInventoryStatistics) GetTotal() *Money {
        return this.TotalInventorySaleValueMoney
}


func (this *StoreFrontInventoryStatistics) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(StoreFrontInventoryStatisticsData_GetInstance().NUMBEROFITEMS, fmt.Sprint(this.GetNumber()))
        hashMap.Put(StoreFrontInventoryStatisticsData_GetInstance().TOTALVALUE, fmt.Sprint(this.GetTotal()))
        return hashMap
}


func (this *StoreFrontInventoryStatistics) ToVector() IBasicArrayList {
        return nil
}


func (this *StoreFrontInventoryStatistics) GetKey() any {
        return nil
}

