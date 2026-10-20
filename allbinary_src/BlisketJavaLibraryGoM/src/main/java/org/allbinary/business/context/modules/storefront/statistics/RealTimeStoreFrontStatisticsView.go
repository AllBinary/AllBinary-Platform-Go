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

type RealTimeStoreFrontStatisticsView struct {
    StoreFrontStatisticsInterface StoreFrontStatisticsInterface
    UserRole IUserRole
}

// Implements DomNodeInterface


func (this *RealTimeStoreFrontStatisticsView) initFields() {
    
    
}


func NewRealTimeStoreFrontStatisticsView(storeFrontStatisticsInterface StoreFrontStatisticsInterface, userRole IUserRole) *RealTimeStoreFrontStatisticsView {
    this := &RealTimeStoreFrontStatisticsView{}
    this.initFields()
        this.StoreFrontStatisticsInterface = storeFrontStatisticsInterface
        this.UserRole = userRole
    return this
}


// throws Exception
func (this *RealTimeStoreFrontStatisticsView) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(StoreFrontStatisticsData_GetInstance().NAME); _ = node
        var storeFrontAdvertisementsStatisticsInterface StoreFrontAdvertisementsStatisticsInterface = this.StoreFrontStatisticsInterface.GetAdvertisements(); _ = storeFrontAdvertisementsStatisticsInterface
        var storeFrontAdvertisementsStatisticsDomNodeInterface DomNodeInterface = NewStoreFrontAdvertisementsStatisticsView(storeFrontAdvertisementsStatisticsInterface); _ = storeFrontAdvertisementsStatisticsDomNodeInterface
        node.AppendChild(storeFrontAdvertisementsStatisticsDomNodeInterface.ToXmlNode(document))
        var storeFrontInventoryStatisticsInterface StoreFrontInventoryStatisticsInterface = this.StoreFrontStatisticsInterface.GetInventory(); _ = storeFrontInventoryStatisticsInterface
        var storeFrontInventoryStatisticsDomNodeInterface DomNodeInterface = NewStoreFrontInventoryStatisticsView(storeFrontInventoryStatisticsInterface); _ = storeFrontInventoryStatisticsDomNodeInterface
        node.AppendChild(storeFrontInventoryStatisticsDomNodeInterface.ToXmlNode(document))
        var storeFrontOrdersHistoryStatisticsInterface StoreFrontOrdersHistoryStatisticsInterface = this.StoreFrontStatisticsInterface.GetOrders(); _ = storeFrontOrdersHistoryStatisticsInterface
        var storeFrontOrdersHistoryStatisticsDomNodeInterface DomNodeInterface = NewStoreFrontOrdersHistoryStatisticsView(storeFrontOrdersHistoryStatisticsInterface); _ = storeFrontOrdersHistoryStatisticsDomNodeInterface
        node.AppendChild(storeFrontOrdersHistoryStatisticsDomNodeInterface.ToXmlNode(document))
        if this.UserRole == UserRoleFactory_GetInstance().ADMINISTRATOR {
            var storeFrontUsersStatisticsInterface StoreFrontUsersStatisticsInterface = this.StoreFrontStatisticsInterface.GetUsers(); _ = storeFrontUsersStatisticsInterface
            var storeFrontUsersStatisticsDomNodeInterface DomNodeInterface = NewStoreFrontUsersStatisticsView(storeFrontUsersStatisticsInterface); _ = storeFrontUsersStatisticsDomNodeInterface
            node.AppendChild(storeFrontUsersStatisticsDomNodeInterface.ToXmlNode(document))
        }
        return node
}

