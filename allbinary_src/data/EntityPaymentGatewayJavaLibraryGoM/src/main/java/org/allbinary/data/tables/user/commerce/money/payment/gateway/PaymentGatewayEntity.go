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

type PaymentGatewayEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements PaymentGatewayEntityInterface


func (this *PaymentGatewayEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "paymentgateways"
}


func NewPaymentGatewayEntity() *PaymentGatewayEntity {
    this := &PaymentGatewayEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *PaymentGatewayEntity) Add(paymentGatewayInterface PaymentGatewayInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "add", e)
            }
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "add")
            }
            var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
            vector.Add(NewPaymentGatewayIdGenerator().GetNext())
            vector.AddAllList(NewPaymentGatewayEncryptedMapping(paymentGatewayInterface).ToVector())
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            vector.Add(time)
            vector.Add(time)
            this.AbSqlBean.Insert(vector)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.END, this, "add")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentGatewayEntity) Update(paymentGatewayInterface PaymentGatewayInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "update", e)
            }
                }
            }()
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            paymentGatewayInterface.SetLastModified(time)
            var whereKeyValuePairs *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereKeyValuePairs
            whereKeyValuePairs.Put(StoreFrontData_GetInstance().NAME, paymentGatewayInterface.GetStoreName())
            whereKeyValuePairs.Put(fmt.Sprint(PaymentGatewayData_NAME), paymentGatewayInterface.GetName())
            if whereKeyValuePairs.Get(EntryData_GetInstance().TIMECREATED) != nil {
                panic(errors.New(fmt.Sprint("Don't update TimeCreated for update")))
            }
            var updateHashMap *HashMap = NewPaymentGatewayEncryptedMapping(paymentGatewayInterface).ToHashMap(); _ = updateHashMap
            this.AbSqlBean.UpdateWhere(whereKeyValuePairs, updateHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "update")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentGatewayEntity) GetPaymentGatewayInterface(storeName string, paymentType *BasicPaymentType) PaymentGatewayInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentGatewayInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put("Command Failed for paymentType: " + paymentType.GetName(), this, "getPaymentGatewayInterface", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var paymentGatewayHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = paymentGatewayHashMap
            var whereKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereKeyAndValue
            whereKeyAndValue.Put(StoreFrontData_GetInstance().NAME, storeName)
            whereKeyAndValue.Put(fmt.Sprint(PaymentGatewayData_NAME), paymentType.GetName())
            paymentGatewayHashMap = this.AbSqlBean.GetRow(whereKeyAndValue)
            if paymentGatewayHashMap != nil {
                var userName string = any(paymentGatewayHashMap.Get(UserData_USERNAME)).(string); _ = userName
                var password string = any(paymentGatewayHashMap.Get(UserData_PASSWORD)).(string); _ = password
                var special1 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL1))).(string); _ = special1
                var special2 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL2))).(string); _ = special2
                var special3 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL3))).(string); _ = special3
                var special4 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL4))).(string); _ = special4
                var special5 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL5))).(string); _ = special5
                var special6 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL6))).(string); _ = special6
                var special7 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL7))).(string); _ = special7
                var special8 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL8))).(string); _ = special8
                var special9 string = any(paymentGatewayHashMap.Get(fmt.Sprint(PaymentGatewayData_SPECIAL9))).(string); _ = special9
                var superCrypt *SuperCrypt = NewSuperCrypt(PaymentGatewayData_CRYPTNUM); _ = superCrypt
                paymentGatewayHashMap.Put(UserData_USERNAME, superCrypt.Decrypt(userName))
                paymentGatewayHashMap.Put(UserData_PASSWORD, superCrypt.Decrypt(password))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL1), superCrypt.Decrypt(special1))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL2), superCrypt.Decrypt(special2))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL3), superCrypt.Decrypt(special3))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL4), superCrypt.Decrypt(special4))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL5), superCrypt.Decrypt(special5))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL6), superCrypt.Decrypt(special6))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL7), superCrypt.Decrypt(special7))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL8), superCrypt.Decrypt(special8))
                paymentGatewayHashMap.Put(fmt.Sprint(PaymentGatewayData_SPECIAL9), superCrypt.Decrypt(special9))
                var paymentGatewayInterface PaymentGatewayInterface = NewPaymentGatewayInterfaceFactory().GetInstance(paymentGatewayHashMap); _ = paymentGatewayInterface
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "getPaymentGatewayInterface()")
                }
                tryControl1, tryResult1 = 1, paymentGatewayInterface; return
            } else {
                tryControl1, tryResult1 = 1, NewPaymentGatewayInterfaceFactory().GetInstance(paymentType); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayEntity) FindPaymentTypeVectorByStore(storeName string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "findPaymentTypeVectorByStore", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var paymentGatewayNameVector IBasicArrayList = NewBasicArrayListD(); _ = paymentGatewayNameVector
            paymentGatewayNameVector = this.AbSqlBean.GetColumnWhere(fmt.Sprint(PaymentGatewayData_NAME), StoreFrontData_GetInstance().NAME, storeName)
            if paymentGatewayNameVector != nil {
                var paymentGatewayVector IBasicArrayList = NewBasicArrayListD(); _ = paymentGatewayVector
                var size int = paymentGatewayNameVector.Size(); _ = size
                for i := int(0); i < size; i++ {
                    var paymentTypeString string = any(paymentGatewayNameVector.Get(i)).(string); _ = paymentTypeString
                    var paymentType *BasicPaymentType = BasicPaymentTypeUtil_GetInstance().Get(paymentTypeString); _ = paymentType
                    paymentGatewayVector.Add(paymentType)
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "findPaymentTypeVectorByStore()")
                }
                tryControl1, tryResult1 = 1, paymentGatewayVector; return
            } else {
                tryControl1, tryResult1 = 1, NewBasicArrayListD(); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentGatewayEntity) Remove(storeName string, paymentType *BasicPaymentType) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "remove", e)
            }
                }
            }()
            var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
            whereHashMap.Put(StoreFrontData_GetInstance().NAME, storeName)
            whereHashMap.Put(fmt.Sprint(PaymentGatewayData_NAME), paymentType.GetName())
            this.AbSqlBean.DeleteWhere(whereHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "remove")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentGatewayEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_ID))
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_AUTO_INCREMENT_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().ENABLE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(StoreFrontData_GetInstance().NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_NAME))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(PaymentData_METHOD)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_MODE))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_TESTPROTOCOL))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_TESTSERVER))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_TESTPORT))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_TESTPATH))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SERVERPROTOCOL))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SERVER))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SERVERPORT))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SERVERPATH))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(UserData_USERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(UserData_PASSWORD)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_TIMEOUT))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYPROTOCOL))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYSERVER))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYPORT))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYPATH))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYUSERNAME))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYPASSWORD))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_PROXYTIMEOUT))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL1))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL2))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL3))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL4))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL5))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL6))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL7))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL8))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_SPECIAL9))
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().LASTMODIFIED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().TIMECREATED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(fmt.Sprint(PaymentGatewayData_ID))
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *PaymentGatewayEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *PaymentGatewayEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

