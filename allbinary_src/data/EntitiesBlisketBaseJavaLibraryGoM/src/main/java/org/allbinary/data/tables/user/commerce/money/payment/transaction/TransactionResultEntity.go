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

type TransactionResultEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements TransactionResultEntityInterface


func (this *TransactionResultEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "vresults"
}


func NewTransactionResultEntity() *TransactionResultEntity {
    this := &TransactionResultEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *TransactionResultEntity) Remove(userName string, orderNumber string) {
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
            whereHashMap.Put(OrderData_ID, any(orderNumber).(string))
            whereHashMap.Put(UserData_USERNAME, userName)
            this.AbSqlBean.DeleteWhere(whereHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "remove")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TransactionResultEntity) Add(userName string, orderNumber string, transactionResultInterface TransactionResultInterface) {
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
            var values IBasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(orderNumber)
            values.Add(userName)
            values.AddAllList(transactionResultInterface.GetValues())
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            values.Add(time)
            values.Add(time)
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "add")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TransactionResultEntity) GetTransactionResultInterface(orderNumber string) TransactionResultInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 TransactionResultInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getTransactionResultInterface", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var resultHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = resultHashMap
            var updateKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateKeyAndValue
            updateKeyAndValue.Put(OrderData_ID, orderNumber)
            resultHashMap = this.AbSqlBean.GetRow(updateKeyAndValue)
            if resultHashMap != nil {
                var transactionResult *TransactionResult = NewTransactionResult(resultHashMap); _ = transactionResult
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "getTransactionResultInterface")
                }
                tryControl1, tryResult1 = 1, any(transactionResult).(TransactionResultInterface); return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *TransactionResultEntity) CreateTableStatement() string {
        var paymentTransactionKeysFactory *PaymentTransactionKeysFactory = PaymentTransactionKeysFactory_GetInstance(); _ = paymentTransactionKeysFactory
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(OrderData_ID).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(UserData_USERNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.RESULT)).Append(this.SqlTypeStrings.MAX_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.PNREF)).Append(this.SqlTypeStrings.TWELVE_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.RESPMSG)).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.AUTHCODE)).Append(this.SqlTypeStrings.SIX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.AVSADDR)).Append(this.SqlTypeStrings.ONE_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.AVSZIP)).Append(this.SqlTypeStrings.ONE_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ORIGRESULT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.STATUS)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.FRAUDCODE)).Append(this.SqlTypeStrings.TWO_INT_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.FRAUDMSG)).Append(this.SqlTypeStrings.MAX_INT_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.ERRCODE)).Append(this.SqlTypeStrings.MAX_INT_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.SCORE)).Append(this.SqlTypeStrings.THREE_INT_NOT_NULL)
        stringBuffer.Append(fmt.Sprint(paymentTransactionKeysFactory.REASON1)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.REASON2)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.REASON3))
        stringBuffer.Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION1)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION2)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION3)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION4)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION5)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION6)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXCEPTION7)).Append(this.SqlTypeStrings.FOUR_INT_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(OrderData_ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *TransactionResultEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *TransactionResultEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

