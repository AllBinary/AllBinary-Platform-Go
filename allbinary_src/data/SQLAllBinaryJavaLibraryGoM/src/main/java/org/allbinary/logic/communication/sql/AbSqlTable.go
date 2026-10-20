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

type AbSqlTable struct {
    AbSqlBasic
    LogUtil *LogUtil
    TableName string
    METHOD_CREATED_TABLE string
    METHOD_DROP_TABLE string
    TABLE_CREATION_SUCCESS string
    DROPPED_SUCCESS string
}


func (this *AbSqlTable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.METHOD_CREATED_TABLE = "createTable()"
    this.METHOD_DROP_TABLE = "dropTable()"
    this.TABLE_CREATION_SUCCESS = "Table Creation Successful: "
    this.DROPPED_SUCCESS = " Dropped Successfully"
}


func NewAbSqlTable(databaseConnectionInfoInterface *DbConnectionInfo) *AbSqlTable {
    this := &AbSqlTable{}
        this.AbSqlBasic = *NewAbSqlBasic(databaseConnectionInfoInterface)
        this.initFields()
    return this
}


func (this *AbSqlTable) SetTableName(tableName string) {
        this.TableName = tableName
}


func (this *AbSqlTable) GetTableName() string {
        return this.TableName
}


// synchronized
func (this *AbSqlTable) CreateTable(data string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put("Table Creation Failed: " + this.TableName + " with statement: " + data, this, this.METHOD_CREATED_TABLE, e)
            }
            tryControl1, tryResult1 = 1, "Failed to Create " + this.TableName + " table."; return
                }
            }()
            this.ExecuteSQLStatement(data)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.TABLE_CREATION_SUCCESS + this.TableName + " with statement: " + data, this, this.METHOD_CREATED_TABLE)
            }
            tryControl1, tryResult1 = 1, this.TableName + fmt.Sprint(this.SqlStrings.CREATE_RETURN); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *AbSqlTable) DropTable() string {
        var sqlStatement string = fmt.Sprint(this.SqlStrings.DROP_TABLE) + this.TableName; _ = sqlStatement
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_DROP_TABLE, e)
            }
            tryControl1, tryResult1 = 1, "Failed to Drop " + this.TableName + " table."; return
                }
            }()
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.METHOD_DROP_TABLE)
            }
            tryControl1, tryResult1 = 1, this.TableName + this.DROPPED_SUCCESS; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

