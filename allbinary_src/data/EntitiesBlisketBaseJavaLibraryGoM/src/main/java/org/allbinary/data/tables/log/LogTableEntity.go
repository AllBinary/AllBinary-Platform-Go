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

type LogTableEntity struct {
    AbSqlBean
    TableName string
}

// Implements LogTableEntityInterface


func (this *LogTableEntity) initFields() {
    this.TableName = "log"
}


func NewLogTableEntity() *LogTableEntity {
    this := &LogTableEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewLogDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *LogTableEntity) CreateTableStatement() string {
        var abeHttpRequestInfoData *AbeHttpRequestInfoData = AbeHttpRequestInfoData_GetInstance(); _ = abeHttpRequestInfoData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(this.SqlStrings.ID).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(abeHttpRequestInfoData.HTTP_USER_AGENT).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(abeHttpRequestInfoData.REMOTE_HOST).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(abeHttpRequestInfoData.REMOTE_ADDRESS).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(abeHttpRequestInfoData.REQUEST_FILE_PATH).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(abeHttpRequestInfoData.REMOTE_PORT).Append(this.SqlTypeStrings.MAX_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(this.SqlStrings.ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *LogTableEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}

