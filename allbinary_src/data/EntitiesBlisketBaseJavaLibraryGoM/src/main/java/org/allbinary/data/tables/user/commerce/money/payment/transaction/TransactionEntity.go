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

type TransactionEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements TransactionEntityInterface


func (this *TransactionEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "vtrans"
}


func NewTransactionEntity() *TransactionEntity {
    this := &TransactionEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *TransactionEntity) Remove(userName string, orderNumber string) {
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


func (this *TransactionEntity) Add(userName string, orderNumber string, paymentTransactionInterface PaymentTransactionInterface) {
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
            var values *BasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(orderNumber)
            values.Add(userName)
            values.AddAllList(paymentTransactionInterface.ToVector())
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


func (this *TransactionEntity) CreateTableStatement() string {
        var paymentTransactionKeysFactory *PaymentTransactionKeysFactory = PaymentTransactionKeysFactory_GetInstance(); _ = paymentTransactionKeysFactory
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(OrderData_ID).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(UserData_USERNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.TRXTYPE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.TENDER)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ACCT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EXPDATE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.AMT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.AUTHCODE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.MICR)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.CHECKNUM)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.NAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.STREET)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.CITY)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.STATE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ZIP)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DL)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.EMAIL)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.COMMENT1)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.COMMENT2)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ORIGID)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.PONUM)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DESC)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DESC1TO4)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.INVNUM)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOZIP)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SWIPE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.TAXAMT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.COMMCARD)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DUTYAMT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.FREIGHTAMT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ORDERDATE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.TAXEXEMPT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.COUNTRYCODE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.CUSTCODE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.CVV2)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ABA)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.ACCTTYPE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DISCOUNT)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.FIRSTNAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.LASTNAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPFROMZIP)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.PRENOTE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.CHKTYPE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.DOB)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.PHONENUM)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SS)).Append(" VARCHAR(30) NOT NULL,").Append(fmt.Sprint(paymentTransactionKeysFactory.COMPANYNAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.COUNTRY)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOCITY)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOFIRSTNAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOLASTNAME)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOSTATE)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SHIPTOSTREET)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SPECIAL1)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SPECIAL2)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(fmt.Sprint(paymentTransactionKeysFactory.SPECIAL3)).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(OrderData_ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *TransactionEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *TransactionEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

