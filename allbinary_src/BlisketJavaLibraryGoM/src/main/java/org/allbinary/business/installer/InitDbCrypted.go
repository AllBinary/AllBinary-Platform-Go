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

//Warning you must have sql root access
type InitDbCrypted struct {
    AbDatabaseManagement
    LogUtil *LogUtil
    UserDbInitInfo *UserDbInitInfo
    InventoryDbInitInfo *InventoryDbInitInfo
    HistoryDbInitInfo *HistoryDbInitInfo
    StaticpagesDbInitInfo *StaticPagesDbInitInfo
    LogDbInitInfo *LogDbInitInfo
    SQL_FILE string
}


func (this *InitDbCrypted) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    this.SQL_FILE = "initWeblisket.sql"
}


func NewInitDbCrypted(databaseConnectionInfoInterface DatabaseConnectionInfoInterface) *InitDbCrypted {
    this := &InitDbCrypted{}
        this.AbDatabaseManagement = *NewAbDatabaseManagement(any(databaseConnectionInfoInterface).(IDbConnectionInfo))
        this.initFields()
    return this
}


func (this *InitDbCrypted) AddUsers() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.METHOD_ADD_USERS, e)
            tryControl1, tryResult1 = 1, true; return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.METHOD_ADD_USERS)
            this.UserDbInitInfo = NewUserDbInitInfo()
            this.InventoryDbInitInfo = NewInventoryDbInitInfo()
            this.HistoryDbInitInfo = NewHistoryDbInitInfo()
            this.StaticpagesDbInitInfo = NewStaticPagesDbInitInfo()
            this.LogDbInitInfo = NewLogDbInitInfo()
            var CUSTOMERDB string = this.UserDbInitInfo.GetName(); _ = CUSTOMERDB
            var CUSTOMERDBUSER string = this.UserDbInitInfo.GetUserName(); _ = CUSTOMERDBUSER
            var CUSTOMERDBPASSWORD string = this.UserDbInitInfo.GetPassword(); _ = CUSTOMERDBPASSWORD
            if !this.AddDbUser(CUSTOMERDB, CUSTOMERDBUSER, CUSTOMERDBPASSWORD) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_USER) + CUSTOMERDBUSER, this, this.METHOD_ADD_USERS)
                tryControl1, tryResult1 = 1, false; return
            }
            var INVENTORYDB string = this.InventoryDbInitInfo.GetName(); _ = INVENTORYDB
            var INVENTORYDBUSER string = this.InventoryDbInitInfo.GetUserName(); _ = INVENTORYDBUSER
            var INVENTORYDBPASSWORD string = this.InventoryDbInitInfo.GetPassword(); _ = INVENTORYDBPASSWORD
            if !this.AddDbUser(INVENTORYDB, INVENTORYDBUSER, INVENTORYDBPASSWORD) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_USER) + INVENTORYDBUSER, this, this.METHOD_ADD_USERS)
                tryControl1, tryResult1 = 1, false; return
            }
            var HISTORYDB string = this.HistoryDbInitInfo.GetName(); _ = HISTORYDB
            var HISTORYDBUSER string = this.HistoryDbInitInfo.GetUserName(); _ = HISTORYDBUSER
            var HISTORYDBPASSWORD string = this.HistoryDbInitInfo.GetPassword(); _ = HISTORYDBPASSWORD
            if !this.AddDbUser(HISTORYDB, HISTORYDBUSER, HISTORYDBPASSWORD) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_USER) + HISTORYDBUSER, this, this.METHOD_ADD_USERS)
                tryControl1, tryResult1 = 1, false; return
            }
            var STATICPAGESDB string = this.StaticpagesDbInitInfo.GetName(); _ = STATICPAGESDB
            var STATICPAGESDBUSER string = this.StaticpagesDbInitInfo.GetUserName(); _ = STATICPAGESDBUSER
            var STATICPAGESDBPASSWORD string = this.StaticpagesDbInitInfo.GetPassword(); _ = STATICPAGESDBPASSWORD
            if !this.AddDbUser(STATICPAGESDB, STATICPAGESDBUSER, STATICPAGESDBPASSWORD) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_USER) + STATICPAGESDBUSER, this, this.METHOD_ADD_USERS)
                tryControl1, tryResult1 = 1, false; return
            }
            var LOGDB string = this.LogDbInitInfo.GetName(); _ = LOGDB
            var LOGDBUSER string = this.LogDbInitInfo.GetUserName(); _ = LOGDBUSER
            var LOGDBPASSWORD string = this.LogDbInitInfo.GetPassword(); _ = LOGDBPASSWORD
            if !this.AddDbUser(LOGDB, LOGDBUSER, LOGDBPASSWORD) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_USER) + LOGDBUSER, this, this.METHOD_ADD_USERS)
                tryControl1, tryResult1 = 1, false; return
            }
            this.LogUtil.PutF(this.CommonStrings.END, this, this.METHOD_ADD_USERS)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDbCrypted) AddDatabases() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.METHOD_ADD_DATABASES, e)
            tryControl1, tryResult1 = 1, true; return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.METHOD_ADD_DATABASES)
            this.UserDbInitInfo = NewUserDbInitInfo()
            this.InventoryDbInitInfo = NewInventoryDbInitInfo()
            this.HistoryDbInitInfo = NewHistoryDbInitInfo()
            this.StaticpagesDbInitInfo = NewStaticPagesDbInitInfo()
            this.LogDbInitInfo = NewLogDbInitInfo()
            var CUSTOMERDB string = this.UserDbInitInfo.GetName(); _ = CUSTOMERDB
            if !this.AddDb(CUSTOMERDB) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_DATABASE) + CUSTOMERDB, this, this.METHOD_ADD_DATABASES)
                tryControl1, tryResult1 = 1, false; return
            }
            var INVENTORYDB string = this.InventoryDbInitInfo.GetName(); _ = INVENTORYDB
            if !this.AddDb(INVENTORYDB) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_DATABASE) + INVENTORYDB, this, this.METHOD_ADD_DATABASES)
                tryControl1, tryResult1 = 1, false; return
            }
            var HISTORYDB string = this.HistoryDbInitInfo.GetName(); _ = HISTORYDB
            if !this.AddDb(HISTORYDB) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_DATABASE) + HISTORYDB, this, this.METHOD_ADD_DATABASES)
                tryControl1, tryResult1 = 1, false; return
            }
            var STATICPAGESDB string = this.StaticpagesDbInitInfo.GetName(); _ = STATICPAGESDB
            if !this.AddDb(STATICPAGESDB) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_DATABASE) + STATICPAGESDB, this, this.METHOD_ADD_DATABASES)
                tryControl1, tryResult1 = 1, false; return
            }
            var LOGDB string = this.LogDbInitInfo.GetName(); _ = LOGDB
            if !this.AddDb(LOGDB) {
                this.LogUtil.PutF(fmt.Sprint(this.UNABLE_TO_CREATE_DATABASE) + LOGDB, this, this.METHOD_ADD_DATABASES)
                tryControl1, tryResult1 = 1, false; return
            }
            var file *AbFile = FileFactory_GetInstance().GetInstance(this.SQL_FILE); _ = file
            file.CreateNewFile()
            var idOutData *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(file); _ = idOutData
            idOutData.WriteBytes(fmt.Sprint(this.SqlCommandLog))
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDbCrypted) AddTables() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.METHOD_ADD_TABLES, e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            this.LogUtil.PutF(this.CommonStrings.START, this, this.METHOD_ADD_TABLES)
            stringBuffer.Append(UserEntityFactory_GetInstance().CreateTable())
            stringBuffer.Append(StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance().CreateTable())
            stringBuffer.Append(BillingAddressesEntityFactory_GetInstance().GetInstance(this.StringUtil.EMPTY_STRING).CreateTable())
            stringBuffer.Append(ShippingAddressesEntityFactory_GetInstance().GetInstance(this.StringUtil.EMPTY_STRING).CreateTable())
            stringBuffer.Append(PaymentEntityFactory_GetInstance().GetPaymentEntityInstance().CreateTable())
            stringBuffer.Append(QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance().CreateTable())
            stringBuffer.Append(OrderHistoryEntityFactory_GetInstance().CreateTable())
            stringBuffer.Append(OrderItemsEntityFactory_GetInstance().CreateTable())
            stringBuffer.Append(TransactionEntityFactory_GetInstance().GetTransactionEntityInstance().CreateTable())
            stringBuffer.Append(TransactionResultEntityFactory_GetInstance().GetTransactionResultEntityInstance().CreateTable())
            stringBuffer.Append(PaymentGatewayEntityFactory_GetInstance().CreateTable())
            stringBuffer.Append(InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().CreateTable())
            stringBuffer.Append(NewInitInfoEntity().CreateTable())
            stringBuffer.Append(LogTableEntityFactory_GetInstance().GetLogTableEntityInstance().CreateTable())
            stringBuffer.Append(StaticPagesEntityFactory_GetInstance().GetStaticPagesEntityInstance().CreateTable())
            stringBuffer.Append(TransformInfoEntityBuilder_GetInstance().CreateTable())
            this.LogUtil.PutF(fmt.Sprint(AbDatabaseManagement_ADD_TABLES_RESULTS_LABEL) + fmt.Sprint(stringBuffer), this, this.METHOD_ADD_TABLES)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

