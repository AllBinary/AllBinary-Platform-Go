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

type PaymentEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements PaymentEntityInterface


func (this *PaymentEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "payment"
}


func NewPaymentEntity() *PaymentEntity {
    this := &PaymentEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *PaymentEntity) GetLastId(userName string) string {
        return this.AbSqlBean.GetLargestIntegerInColumnWhere(PaymentData_ID, UserData_USERNAME, userName)
}


func (this *PaymentEntity) SetDefault(userName string, index int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "setDefault", e)
            }
                }
            }()
            var updateKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateKeyAndValue
            var whereKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereKeyAndValue
            whereKeyAndValue.Put(UserData_USERNAME, userName)
            var paymentInterface PaymentInterface = this.GetDefault(userName); _ = paymentInterface
            if paymentInterface != nil {
                updateKeyAndValue.Put(EntryData_GetInstance().DEFAULT, StringUtil_GetInstance().EMPTY_STRING)
                whereKeyAndValue.Put(PaymentData_ID, paymentInterface.GetId())
                this.AbSqlBean.UpdateWhere(whereKeyAndValue, updateKeyAndValue)
            }
            updateKeyAndValue.Put(EntryData_GetInstance().DEFAULT, EntryData_GetInstance().DEFAULT)
            whereKeyAndValue.Put(PaymentData_ID, fmt.Sprint(index))
            this.AbSqlBean.UpdateWhere(whereKeyAndValue, updateKeyAndValue)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "setDefault")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentEntity) Get(userName string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var paymentVector IBasicArrayList = NewBasicArrayListD(); _ = paymentVector
            var keyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyAndValue
            keyAndValue.Put(UserData_USERNAME, userName)
            var paymentList IBasicArrayList = this.AbSqlBean.GetRows(keyAndValue); _ = paymentList
            var size int = paymentList.Size(); _ = size
            for index := int(0); index < size; index++ {
                var paymentHashMap *HashMap = any(paymentList.Get(index)).(*HashMap); _ = paymentHashMap
                var payment *Payment = NewPayment(paymentHashMap); _ = payment
                if payment != nil {
                    paymentVector.Add(payment)
                }
            }
            tryControl1, tryResult1 = 1, paymentVector; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentEntity) GetDefault(userName string) PaymentInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 PaymentInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getDefault", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var paymentHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = paymentHashMap
            var updateKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateKeyAndValue
            updateKeyAndValue.Put(EntryData_GetInstance().DEFAULT, EntryData_GetInstance().DEFAULT)
            updateKeyAndValue.Put(UserData_USERNAME, userName)
            paymentHashMap = this.AbSqlBean.GetRow(updateKeyAndValue)
            if paymentHashMap != nil {
                var payment *Payment = NewPayment(paymentHashMap); _ = payment
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "getDefault")
                }
                tryControl1, tryResult1 = 1, any(payment).(PaymentInterface); return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PaymentEntity) Remove(userName string, index int) {
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
            whereHashMap.Put(UserData_USERNAME, userName)
            whereHashMap.Put(PaymentData_ID, any(fmt.Sprint(index)).(string))
            this.AbSqlBean.DeleteWhere(whereHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "remove")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentEntity) Add(userName string, paymentInterface PaymentInterface) {
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
            var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
            vector.Add(NewPaymentIdGenerator().GetNext())
            vector.Add(userName)
            vector.Add(StringUtil_GetInstance().EMPTY_STRING)
            vector.Add(paymentInterface.GetName())
            vector.Add(paymentInterface.GetType())
            vector.Add(paymentInterface.GetExpiration())
            var random int = NewRandom().NextInt(SuperCrypt_KEYMAX); _ = random
            vector.Add(NewSuperCrypt(random).Encrypt(paymentInterface.GetNumber()))
            vector.Add(fmt.Sprint(int(random)))
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            vector.Add(time)
            this.AbSqlBean.Insert(vector)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "add")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PaymentEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(PaymentData_ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_AUTO_INCREMENT_NOT_NULL).Append(UserData_USERNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(entryData.DEFAULT).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(PaymentData_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(PaymentData_TYPE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(PaymentData_EXPIRATION).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(PaymentData_NUMBER).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(entryData.ENCRYPTION).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(entryData.TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(PaymentData_ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *PaymentEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *PaymentEntity) Drop() string {
        return this.AbSqlBean.DropTable()
}

