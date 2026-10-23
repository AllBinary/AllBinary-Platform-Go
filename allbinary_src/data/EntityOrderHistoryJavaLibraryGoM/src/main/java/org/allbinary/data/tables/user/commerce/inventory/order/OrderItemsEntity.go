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

type OrderItemsEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    TableName string
}

// Implements OrderItemsEntityInterface


func (this *OrderItemsEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    this.TableName = "orderitems"
}


func NewOrderItemsEntity() *OrderItemsEntity {
    this := &OrderItemsEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewHistoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *OrderItemsEntity) Insert(userName string, order OrderInterface) {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put("Command Failed: " + fmt.Sprint(vector), this, AbSqlBean_INSERT, e)
            }
                }
            }()
            var basket BasketInterface = order.GetBasket(); _ = basket
            var items Set = basket.GetIds(); _ = items
            var ZERO_STRING string = TableDataFactory_GetInstance().ZERO_STRING; _ = ZERO_STRING
            var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
            var inventoryEntity *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntity
            var itemsArray []any = items.ToArray(); _ = itemsArray
            var itemsSize int = len(itemsArray); _ = itemsSize
            for i := int(0); i < itemsSize; i++ {
                var item string = any(itemsArray[i]).(string); _ = item
                var itemInterface ItemInterface = inventoryEntity.GetItem(item); _ = itemInterface
                vector = NewBasicArrayListD()
                vector.Add(NewOrderItemIdGenerator().GetNext())
                vector.Add(order.GetId())
                vector.Add(TableDataFactory_GetInstance().INTEGER_MAX_VALUE_STRING)
                vector.Add(userName)
                vector.Add(order.GetStoreName())
                vector.Add(itemInterface.GetId())
                vector.Add(fmt.Sprint(basket.GetNumberOf(item)))
                vector.Add(itemInterface.GetNewOrUsed())
                vector.Add(itemInterface.GetSummary())
                vector.Add(itemInterface.GetDistributor())
                vector.Add(itemInterface.GetIdUsedByDistributor())
                vector.Add(itemInterface.GetProducedBy())
                vector.Add(itemInterface.GetProductionDate())
                vector.Add(itemInterface.GetStartProductionDate())
                vector.Add(itemInterface.GetDescription())
                vector.Add(itemInterface.GetKeywords())
                vector.Add(itemInterface.GetCategory())
                vector.Add(itemInterface.GetType())
                vector.Add(itemInterface.GetSmallImage())
                vector.Add(itemInterface.GetMediumImage())
                vector.Add(itemInterface.GetLargeImage())
                vector.Add(itemInterface.GetWeight())
                vector.Add(fmt.Sprint(itemInterface.GetPrice()))
                vector.Add(itemInterface.GetComment())
                vector.Add(itemInterface.GetCustoms())
                vector.Add(itemInterface.GetDownloads())
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(EMPTY_STRING)
                vector.Add(ZERO_STRING)
                vector.Add(ZERO_STRING)
                vector.Add(ZERO_STRING)
                vector.Add(ZERO_STRING)
                this.Insert(vector)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF("Command Success Added: " + fmt.Sprint(items.Size()), this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OrderItemsEntity) SetStatus(orderId string, groupId string, status string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "setStatus", e)
            }
                }
            }()
            var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
            var updateHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateHashMap
            whereHashMap.Put(OrderData_ID, orderId)
            whereHashMap.Put(ShippingMethodData_GROUP, groupId)
            updateHashMap.Put(OrderHistoryData_STATUS, status)
            this.AbSqlBean.UpdateWhere(whereHashMap, updateHashMap)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OrderItemsEntity) IsEverythingShipped(orderId string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "isEverythingShipped", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var itemStatusVector IBasicArrayList = this.AbSqlBean.GetColumnWhere(OrderHistoryData_STATUS, OrderData_ID, orderId); _ = itemStatusVector
            var size int = itemStatusVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var status string = any(itemStatusVector.Get(index)).(string); _ = status
                if strings.Compare(status, OrderHistoryData_SHIPPED) != 0 {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *OrderItemsEntity) GetBasketReview(orderId string) *BasketReview {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasketReview) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getBasketReview", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var basketReview *BasketReview = NewBasketReview(); _ = basketReview
            var keyValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyValues
            keyValues.Put(OrderData_ID, orderId)
            var items IBasicArrayList = this.AbSqlBean.GetRows(keyValues); _ = items
            var entryData *EntryData = EntryData_GetInstance(); _ = entryData
            if items != nil && items.Size() > 0 {
                var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
                var itemsArray []any = items.ToArray(); _ = itemsArray
                var itemsSize int = len(itemsArray); _ = itemsSize
                for i := int(0); i < itemsSize; i++ {
                    var itemHashMap *HashMap = any(itemsArray[i]).(*HashMap); _ = itemHashMap
                    itemHashMap.Put(this.BasicItemData.INBASKETS, EMPTY_STRING)
                    itemHashMap.Put(entryData.ENABLE, EMPTY_STRING)
                    itemHashMap.Put(entryData.TIMECREATED, EMPTY_STRING)
                    itemHashMap.Put(entryData.LASTMODIFIED, EMPTY_STRING)
                    itemHashMap.Put(this.BasicItemData.GROUPS, EMPTY_STRING)
                    itemHashMap.Put(this.BasicItemData.OPTIONS, EMPTY_STRING)
                    itemHashMap.Put(this.BasicItemData.PERMISSIONS, EMPTY_STRING)
                    itemHashMap.Put(this.BasicItemData.SPECIALS, EMPTY_STRING)
                    var item IItem = NewItem(itemHashMap); _ = item
                    basketReview.AddItem(item)
                }
                tryControl1, tryResult1 = 1, basketReview; return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *OrderItemsEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("CREATE TABLE ")
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(" (")
        stringBuffer.Append(entryData.ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(OrderData_ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(ShippingMethodData_GROUP)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(UserData_USERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(StoreFrontData_GetInstance().NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.NUMBER)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.NEWORUSED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.SUMMARY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DISTRIBUTOR)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.IDUSEDBYDISTRIBUTOR)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRODUCEDBY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRODUCTIONDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.STARTPRODUCTIONDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DESCRIPTION)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.KEYWORDS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.CATEGORY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.TYPE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.SMALLIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.MEDIUMIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.LARGEIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.WEIGHT)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRICE)
        stringBuffer.Append(" VARCHAR(20) NOT NULL,")
        stringBuffer.Append(this.BasicItemData.COMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.CUSTOMS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DOWNLOADS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(ShippingMethodData_COST)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_TAX)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_STREET)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_CITY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_STATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_CODE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingAddressData_COUNTRY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(ShippingMethodData_NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(entryData.SPECIAL)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(OrderHistoryData_STATUS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_CANCELINFO)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_CANCELTYPE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_SHIPPEDDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(OrderHistoryData_ORDERDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_TRANSDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(OrderHistoryData_CANCELDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append("PRIMARY KEY(")
        stringBuffer.Append(entryData.ID)
        stringBuffer.Append(") )")
        return fmt.Sprint(stringBuffer)
}


func (this *OrderItemsEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *OrderItemsEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

