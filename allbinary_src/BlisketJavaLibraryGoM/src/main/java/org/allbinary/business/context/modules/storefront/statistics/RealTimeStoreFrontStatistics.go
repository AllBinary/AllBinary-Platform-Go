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

type RealTimeStoreFrontStatistics struct {
    StoreFrontAdvertisementsStatisticsInterface StoreFrontAdvertisementsStatisticsInterface
    StoreFrontUsersStatisticsInterface StoreFrontUsersStatisticsInterface
    StoreFrontInventoryStatisticsInterface StoreFrontInventoryStatisticsInterface
    StoreFrontOrderHistoryStatisticsInterface StoreFrontOrdersHistoryStatisticsInterface
}

// Implements StoreFrontStatisticsInterface


func (this *RealTimeStoreFrontStatistics) initFields() {
    
    
    
    
}


func NewRealTimeStoreFrontStatistics(storeFrontInterface StoreFrontInterface) *RealTimeStoreFrontStatistics {
    this := &RealTimeStoreFrontStatistics{}
    this.initFields()
        this.StoreFrontAdvertisementsStatisticsInterface = any(NewStoreFrontAdvertisementsStatistics(storeFrontInterface)).(StoreFrontAdvertisementsStatisticsInterface)
        this.StoreFrontUsersStatisticsInterface = any(NewStoreFrontUsersStatistics(storeFrontInterface)).(StoreFrontUsersStatisticsInterface)
        this.StoreFrontInventoryStatisticsInterface = any(NewStoreFrontInventoryStatistics(storeFrontInterface)).(StoreFrontInventoryStatisticsInterface)
        this.StoreFrontOrderHistoryStatisticsInterface = any(NewStoreFrontOrdersHistoryStatistics(storeFrontInterface)).(StoreFrontOrdersHistoryStatisticsInterface)
    return this
}


func (this *RealTimeStoreFrontStatistics) GetAdvertisements() StoreFrontAdvertisementsStatisticsInterface {
        return this.StoreFrontAdvertisementsStatisticsInterface
}


func (this *RealTimeStoreFrontStatistics) GetUsers() StoreFrontUsersStatisticsInterface {
        return this.StoreFrontUsersStatisticsInterface
}


func (this *RealTimeStoreFrontStatistics) GetInventory() StoreFrontInventoryStatisticsInterface {
        return this.StoreFrontInventoryStatisticsInterface
}


func (this *RealTimeStoreFrontStatistics) GetOrders() StoreFrontOrdersHistoryStatisticsInterface {
        return this.StoreFrontOrderHistoryStatisticsInterface
}

