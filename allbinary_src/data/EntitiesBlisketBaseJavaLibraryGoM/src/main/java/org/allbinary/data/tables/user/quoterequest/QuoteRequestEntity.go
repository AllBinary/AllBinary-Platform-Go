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

type QuoteRequestEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements QuoteRequestEntityInterface


func (this *QuoteRequestEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "quoterequest"
}


func NewQuoteRequestEntity() *QuoteRequestEntity {
    this := &QuoteRequestEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *QuoteRequestEntity) Insert(values IBasicArrayList) {
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


// throws Exception
func (this *QuoteRequestEntity) Get(userName string, id int) *QuoteRequest {
        var row *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = row
        row.Put(UserData_USERNAME, userName)
        row.Put(QuoteRequestData_GetInstance().ID, fmt.Sprint(id))
        var quoteRequestHashMap *HashMap = this.AbSqlBean.GetRow(row); _ = quoteRequestHashMap
        if quoteRequestHashMap != nil {
            return NewQuoteRequest(quoteRequestHashMap)
        } else {
            return nil
        }
}


func (this *QuoteRequestEntity) GetIds(userName string) IBasicArrayList {
        return this.AbSqlBean.GetColumnWhere(QuoteRequestData_GetInstance().ID, UserData_USERNAME, userName)
}


func (this *QuoteRequestEntity) DeleteWhere(key string, value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "deleteWhere", e)
            }
                }
            }()
            this.AbSqlBean.DeleteWhere(key, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "deleteWhere")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *QuoteRequestEntity) CreateTableStatement() string {
        var quoteRequestData *QuoteRequestData = QuoteRequestData_GetInstance(); _ = quoteRequestData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START)
        stringBuffer.Append(quoteRequestData.ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(UserData_USERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(quoteRequestData.PROJECT_INFO)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(quoteRequestData.CUSTOMER_COMMENTS)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(quoteRequestData.BUDGET)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(quoteRequestData.TIMEFRAME)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(quoteRequestData.COMMENTS)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().TIMECREATED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().LASTMODIFIED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(quoteRequestData.ID)
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *QuoteRequestEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *QuoteRequestEntity) DropTable() string {
        var result string = this.AbSqlBean.DropTable(); _ = result
        return result
}


func (this *QuoteRequestEntity) Update(userName string, updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(UserData_USERNAME, userName, updatedValues)
}

