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

type AffiliateSalesEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements AffiliateSalesEntityInterface


func (this *AffiliateSalesEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "affiliate"
}


func NewAffiliateSalesEntity() *AffiliateSalesEntity {
    this := &AffiliateSalesEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *AffiliateSalesEntity) Delete(value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.Delete, e)
            }
                }
            }()
            this.AbSqlBean.DeleteWhere(StoreFrontData_GetInstance().NAME, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AffiliateSalesEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(entryData.ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(StoreFrontData_GetInstance().NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(GroupData_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DynamicObjectData_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(entryData.TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(entryData.LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(entryData.ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *AffiliateSalesEntity) CreateTable() string {
        var returnStr string = this.AbSqlBean.CreateTable(this.CreateTableStatement()); _ = returnStr
        return returnStr
}


func (this *AffiliateSalesEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(EntryData_GetInstance().ID, any(updatedValues.Get(EntryData_GetInstance().ID)).(string), updatedValues)
}


func (this *AffiliateSalesEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

