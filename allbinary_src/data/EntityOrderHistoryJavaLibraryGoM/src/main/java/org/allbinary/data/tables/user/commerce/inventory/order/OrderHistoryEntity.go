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

type OrderHistoryEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements OrderHistoryEntityInterface


func (this *OrderHistoryEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "orderhistory"
}


func NewOrderHistoryEntity() *OrderHistoryEntity {
    this := &OrderHistoryEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewHistoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *OrderHistoryEntity) Insert_1(userName string, order IOrder) {
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
            var paymentInterface PaymentInterface = NewPaymentEntity().GetDefault(userName); _ = paymentInterface
            var billingAddress *StreetAddress = NewBillingAddressesEntity(userName).GetDefault(); _ = billingAddress
            var shippingAddress *StreetAddress = NewShippingAddressesEntity(userName).GetDefault(); _ = shippingAddress
            var abeClientInformation AbeClientInformationInterface = ServiceClientInformationInterfaceFactory_GetInstance(); _ = abeClientInformation
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(order.GetStoreName()); _ = storeFrontInterface
            var shippingInterface ShippingInterface = NewShippingMethods(abeClientInformation, storeFrontInterface).GetShippingInterface(order.GetShippingMethod()); _ = shippingInterface
            var shippingCost *Money = shippingInterface.GetCost(order); _ = shippingCost
            var basketInterface BasketInterface = order.GetBasket(); _ = basketInterface
            var subTotal *Money = basketInterface.GetSubTotal(); _ = subTotal
            var taxRate float32 = TaxFactory_GetInstance().GetInstance(abeClientInformation, storeFrontInterface).GetTaxRate(shippingAddress, storeFrontInterface); _ = taxRate
            var tax *Money = NewMoney(); _ = tax
            var total *Money = NewMoney(); _ = total
            total.Add(fmt.Sprint(shippingCost))
            total.Add(fmt.Sprint(subTotal))
            tax.Add(fmt.Sprint(total))
            tax.Multiply(taxRate)
            total.Add(fmt.Sprint(tax))
            var empty string = StringUtil_GetInstance().EMPTY_STRING; _ = empty
            vector.Add(NewOrderHistoryIdGenerator().GetNext())
            var ZERO string = TableDataFactory_GetInstance().ZERO_STRING; _ = ZERO
            vector.Add(order.GetId())
            vector.Add(userName)
            vector.Add(order.GetStoreName())
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            vector.Add(ZERO)
            vector.Add(time)
            vector.Add(ZERO)
            vector.Add(ZERO)
            vector.Add(OrderHistoryData_PREPROCESSING)
            vector.Add(order.GetPaymentMethod())
            if paymentInterface != nil {
                vector.Add(paymentInterface.GetName())
                vector.Add(paymentInterface.GetType())
                vector.Add(paymentInterface.GetExpiration())
                var random int = NewRandom().NextInt(SuperCrypt_KEYMAX); _ = random
                vector.Add(NewSuperCrypt(random).Encrypt(paymentInterface.GetNumber()))
                vector.Add(fmt.Sprint(int(random)))
            } else {
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(ZERO)
            }
            if billingAddress != nil {
                vector.Add(billingAddress.GetName())
                vector.Add(billingAddress.GetStreet())
                vector.Add(billingAddress.GetCity())
                vector.Add(billingAddress.GetState())
                vector.Add(billingAddress.GetCode())
                vector.Add(billingAddress.GetCountry())
            } else {
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
            }
            if shippingAddress != nil {
                vector.Add(shippingAddress.GetName())
                vector.Add(shippingAddress.GetStreet())
                vector.Add(shippingAddress.GetCity())
                vector.Add(shippingAddress.GetState())
                vector.Add(shippingAddress.GetCode())
                vector.Add(shippingAddress.GetCountry())
            } else {
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
                vector.Add(empty)
            }
            vector.Add(order.GetShippingMethod())
            vector.Add(fmt.Sprint(subTotal))
            vector.Add(fmt.Sprint(shippingCost))
            vector.Add(fmt.Sprint(tax))
            vector.Add(fmt.Sprint(total))
            vector.Add(order.GetSpecial())
            vector.Add(order.GetUserComments())
            vector.Add(order.GetUserCancelComments())
            vector.Add(order.GetStoreComments())
            vector.Add(order.GetStoreCancelComments())
            this.Insert(vector)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of OrderHistoryEntity.insert
func (this *OrderHistoryEntity) Insert(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[IOrder](args[1], pass) {
            this.Insert_1(Java_Arg[string](args[0]), Java_Arg[IOrder](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[IBasicArrayList](args[0], pass) {
            this.Insert_2(Java_Arg[IBasicArrayList](args[0]))
            return
        }
    }
    panic("No Java overload of OrderHistoryEntity.insert for the arguments")
}


func (this *OrderHistoryEntity) Insert_2(values IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, AbSqlBean_INSERT, e)
            }
                }
            }()
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OrderHistoryEntity) SetStatus(orderId string, status string) {
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
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            var updateHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateHashMap
            updateHashMap.Put(OrderHistoryData_STATUS, status)
            if strings.Compare(status, OrderHistoryData_CANCELLED) == 0 {
                updateHashMap.Put(OrderHistoryData_CANCELDATE, time)
            } else if strings.Compare(status, OrderHistoryData_SHIPPED) == 0 {
                updateHashMap.Put(OrderHistoryData_SHIPPEDDATE, time)
            }
            this.AbSqlBean.UpdateWhere(OrderData_ID, orderId, updateHashMap)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *OrderHistoryEntity) SetPaymentMethod(orderId string, paymentMethod string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "setPaymentMethod", e)
            }
                }
            }()
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            var updateHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateHashMap
            updateHashMap.Put(PaymentData_METHOD, paymentMethod)
            updateHashMap.Put(OrderHistoryData_TRANSDATE, time)
            this.AbSqlBean.UpdateWhere(OrderData_ID, orderId, updateHashMap)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *OrderHistoryEntity) GetStoreOrders(storeFrontInterface StoreFrontInterface) IBasicArrayList {
        var orderReviewVector IBasicArrayList = NewBasicArrayListD(); _ = orderReviewVector
        var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
        whereHashMap.Put(StoreFrontData_GetInstance().NAME, storeFrontInterface.GetName())
        var orderHashMapVector IBasicArrayList = this.AbSqlBean.GetRows(whereHashMap); _ = orderHashMapVector
        var size int = orderHashMapVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var orderReviewHashMap *HashMap = any(orderHashMapVector.Get(i)).(*HashMap); _ = orderReviewHashMap
            var orderReview *OrderHistory = NewOrderHistory(NewBasket(), orderReviewHashMap); _ = orderReview
            orderReviewVector.Add(orderReview)
        }
        return orderReviewVector
}


// throws Exception
func (this *OrderHistoryEntity) GetOrders_1(userName string) IBasicArrayList {
        var orderReviewVector IBasicArrayList = NewBasicArrayListD(); _ = orderReviewVector
        var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
        whereHashMap.Put(UserData_USERNAME, userName)
        var orderHashMapVector IBasicArrayList = this.AbSqlBean.GetRows(whereHashMap); _ = orderHashMapVector
        var size int = orderHashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var orderReviewHashMap *HashMap = any(orderHashMapVector.Get(index)).(*HashMap); _ = orderReviewHashMap
            var orderReview *OrderHistory = NewOrderHistory(NewBasket(), orderReviewHashMap); _ = orderReview
            orderReviewVector.Add(orderReview)
        }
        return orderReviewVector
}

// Java overloads of OrderHistoryEntity.getOrders
func (this *OrderHistoryEntity) GetOrders(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetOrders_1(Java_Arg[string](args[0]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetOrders_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetOrders_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of OrderHistoryEntity.getOrders for the arguments")
}


// throws Exception
func (this *OrderHistoryEntity) GetOrders_2(status string, fromDate string, toDate string) IBasicArrayList {
        var orderReviewVector IBasicArrayList = NewBasicArrayListD(); _ = orderReviewVector
        var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
        whereHashMap.Put(OrderHistoryData_STATUS, status)
        var orderHashMapVector IBasicArrayList = this.AbSqlBean.GetRowsWhereBetween(whereHashMap, OrderHistoryData_ORDERDATE, fromDate, toDate); _ = orderHashMapVector
        var size int = orderHashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var orderReviewHashMap *HashMap = any(orderHashMapVector.Get(index)).(*HashMap); _ = orderReviewHashMap
            var orderReview *OrderHistory = NewOrderHistory(NewBasket(), orderReviewHashMap); _ = orderReview
            orderReviewVector.Add(orderReview)
        }
        return orderReviewVector
}


// throws Exception
func (this *OrderHistoryEntity) GetOrders_3(fromDate string, toDate string) IBasicArrayList {
        var orderReviewVector IBasicArrayList = NewBasicArrayListD(); _ = orderReviewVector
        var orderHashMapVector IBasicArrayList = this.AbSqlBean.GetRowsWhereBetween(OrderHistoryData_ORDERDATE, fromDate, toDate); _ = orderHashMapVector
        var size int = orderHashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var orderReviewHashMap *HashMap = any(orderHashMapVector.Get(index)).(*HashMap); _ = orderReviewHashMap
            var orderReview *OrderHistory = NewOrderHistory(NewBasket(), orderReviewHashMap); _ = orderReview
            orderReviewVector.Add(orderReview)
        }
        return orderReviewVector
}


// throws Exception
func (this *OrderHistoryEntity) GetOrder(id string) *OrderHistory {
        var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
        whereHashMap.Put(OrderData_ID, id)
        var orderReviewHashMap *HashMap = this.AbSqlBean.GetRow(whereHashMap); _ = orderReviewHashMap
        if orderReviewHashMap != nil {
            var orderHistory *OrderHistory = NewOrderHistory(NewBasket(), orderReviewHashMap); _ = orderHistory
            return orderHistory
        } else {
            return nil
        }
}


func (this *OrderHistoryEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("CREATE TABLE ")
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(" (")
        stringBuffer.Append(entryData.ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(OrderData_ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(UserData_USERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(StoreFrontData_GetInstance().NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_SHIPPEDDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(OrderHistoryData_ORDERDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_TRANSDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(OrderHistoryData_CANCELDATE)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(OrderHistoryData_STATUS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(PaymentData_METHOD)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(PaymentData_NAME)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(PaymentData_TYPE)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(PaymentData_EXPIRATION)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(PaymentData_NUMBER)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(entryData.ENCRYPTION)
        stringBuffer.Append(" BIGINT(19) UNSIGNED ,")
        stringBuffer.Append(BillingAddressData_NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BillingAddressData_STREET)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BillingAddressData_CITY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BillingAddressData_STATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BillingAddressData_CODE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BillingAddressData_COUNTRY)
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
        stringBuffer.Append(OrderHistoryData_SUBTOTAL)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_SHIPPINGCOST)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_TAX)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(OrderHistoryData_TOTAL)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(entryData.SPECIAL)
        stringBuffer.Append(" VARCHAR(255) ,")
        stringBuffer.Append(OrderData_CUSTOMERCOMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(OrderData_CUSTOMERCANCELCOMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(OrderData_STORECOMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(OrderData_STORECANCELCOMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append("PRIMARY KEY(")
        stringBuffer.Append(entryData.ID)
        stringBuffer.Append(") )")
        return fmt.Sprint(stringBuffer)
}


func (this *OrderHistoryEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *OrderHistoryEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}


// throws Exception
func (this *OrderHistoryEntity) Update(whereHashMap *HashMap, orderHashMap *HashMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.UpdateWhere(whereHashMap, orderHashMap)
            return
        }(); tryControl1 == 1 {
            return
        }
}

