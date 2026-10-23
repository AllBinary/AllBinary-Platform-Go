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

type InitSql struct {
    StringUtil *StringUtil
    CommonSeps *CommonSeps
    SqlTypeStrings *SqlTypeStrings
    SqlStrings *SqlStrings
    StdUtil *StdUtil
    INSERT string
    UPDATE string
    DatabaseConnectionInfoInterface IDbConnectionInfo
    Userid string
    Password string
    TableName string
    Conn IConnection
    UseridAndPassword bool
    SUCCESS_SQL_STATEMENT string
    FAILED_SQL_STATEMENT string
    INIT_SQL string
    CommonStrings *CommonStrings
    EQUAL_QUOTE string
    METHOD_GET_ROW string
    METHOD_UPDATE_WHERE string
    ROW_VALUE_LABEL string
    NO_RESULTS_IN_RESULT_SET string
    INSERT_END string
}

// The Java class InitSql has subclasses so its variables are this interface
type IInitSql interface {
    InitSql_() *InitSql
    SetTable(tableName string)
    SetDatabaseConnectionInfoInterface(databaseConnectionInfoInterface IDbConnectionInfo)
    CreateTable(tableData string) bool
    DropTable() bool
    GetRow(keysAndValues *HashMap) *HashMap
    UpdateWhere(key string, value string, updatedKeyValuePairs *HashMap)
    Insert(values IBasicArrayList)
    GetValue(value string) string
    ExecuteSQLStatement(statement string) *ResultSet
    CreateConnection()
    Initialize()
    GetDatabaseConnectionInfoInterface() DatabaseConnectionInfoInterface
}

func (this *InitSql) InitSql_() *InitSql {
    return this
}


func (this *InitSql) initFields() {
    this.StringUtil = StringUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.SqlTypeStrings = SqlTypeStrings_GetInstance()
    this.SqlStrings = SqlStrings_GetInstance()
    this.StdUtil = StdUtil_GetInstance()
    this.INSERT = "insert"
    this.UPDATE = "update"
    
    
    
    
    
    
    this.SUCCESS_SQL_STATEMENT = "Success\nSQL Statement: "
    this.FAILED_SQL_STATEMENT = "Failed\nSQL Statement: "
    this.INIT_SQL = "InitSql"
    this.CommonStrings = CommonStrings_GetInstance()
    this.EQUAL_QUOTE = "=\""
    this.METHOD_GET_ROW = "getRow()"
    this.METHOD_UPDATE_WHERE = "updateWhere()"
    this.ROW_VALUE_LABEL = "Row Value: "
    this.NO_RESULTS_IN_RESULT_SET = "No Results in ResultSet"
    this.INSERT_END = "')"
}


func NewInitSql(databaseConnectionInfoInterface IDbConnectionInfo) *InitSql {
    this := &InitSql{}
    this.initFields()
        this.SetDatabaseConnectionInfoInterface(databaseConnectionInfoInterface)
    return this
}


func (this *InitSql) SetTable(tableName string) {
        this.TableName = tableName
}


func (this *InitSql) SetDatabaseConnectionInfoInterface(databaseConnectionInfoInterface IDbConnectionInfo) {
        this.DatabaseConnectionInfoInterface = databaseConnectionInfoInterface
}


func (this *InitSql) CreateTable(tableData string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.INIT_SQL, "createTable()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Creating Table: ")
            stringBuffer.Append(tableData)
            PreLogUtil_Put(fmt.Sprint(stringBuffer), this.INIT_SQL, "createTable()")
            this.ExecuteSQLStatement(tableData)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitSql) DropTable() bool {
        var sqlStatement string = fmt.Sprint(this.SqlStrings.DROP_TABLE) + this.TableName; _ = sqlStatement
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.INIT_SQL, "dropTables()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.ExecuteSQLStatement(sqlStatement)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitSql) GetRow(keysAndValues *HashMap) *HashMap {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.WHERE)
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HashMap) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.FAILED_SQL_STATEMENT + fmt.Sprint(stringBuffer), this.INIT_SQL, this.METHOD_GET_ROW, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var result *HashMap = nil; _ = result
            var set Set = keysAndValues.KeySet(); _ = set
            var keyArray []any = set.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for i := int(0); i < size; i++ {
                var key string = any(keyArray[i]).(string); _ = key
                var value string = any(keysAndValues.Get(key)).(string); _ = value
                stringBuffer.Append(key)
                stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                stringBuffer.Append(this.GetValue(value))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if i < size - 1 {
                    stringBuffer.Append(this.SqlStrings.AND)
                }
            }
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                PreLogUtil_Put(fmt.Sprint(this.SqlStrings.SQL_STATEMENT_LABEL) + sqlStatement, this.INIT_SQL, this.METHOD_GET_ROW)
            }
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var resultSetMetaData *ResultSetMetaData = rset.GetMetaData(); _ = resultSetMetaData
            for rset.Next() {
                result = this.StdUtil.CreateHashMap()
                var columnCount int = resultSetMetaData.GetColumnCount(); _ = columnCount
                for index := int(1); index <= columnCount; index++ {
                    var columnName string = resultSetMetaData.GetColumnName(index); _ = columnName
                    var field string = rset.GetString(columnName); _ = field
                    result.Put(columnName, field)
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    PreLogUtil_Put(this.ROW_VALUE_LABEL + fmt.Sprint(result), this.INIT_SQL, this.METHOD_GET_ROW)
                }
                tryControl1, tryResult1 = 1, result; return
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_Put(this.NO_RESULTS_IN_RESULT_SET, this.INIT_SQL, this.METHOD_GET_ROW)
            }
            tryControl1, tryResult1 = 1, nil; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *InitSql) UpdateWhere(key string, value string, updatedKeyValuePairs *HashMap) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.UPDATE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.SET)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.FAILED_SQL_STATEMENT + fmt.Sprint(stringBuffer), this.INIT_SQL, this.METHOD_UPDATE_WHERE, e)
            }
                }
            }()
            var columnArray []any = updatedKeyValuePairs.KeySet().ToArray(); _ = columnArray
            var size int = len(columnArray); _ = size
            for i := int(0); i < size; i++ {
                var columnName string = fmt.Sprint(columnArray[i]); _ = columnName
                stringBuffer.Append(this.CommonSeps.SPACE)
                stringBuffer.Append(columnName)
                stringBuffer.Append(this.EQUAL_QUOTE)
                var columnValue string = any(updatedKeyValuePairs.Get(columnName)).(string); _ = columnValue
                if columnValue == "" {
                    columnValue = this.StringUtil.EMPTY_STRING
                } else {
                }
                stringBuffer.Append(this.GetValue(columnValue))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if i < size - 1 {
                    stringBuffer.Append(this.CommonSeps.COMMA)
                }
            }
            stringBuffer.Append(this.SqlStrings.WHERE)
            stringBuffer.Append(key)
            stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
            stringBuffer.Append(this.GetValue(value))
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                PreLogUtil_Put(this.SUCCESS_SQL_STATEMENT + sqlStatement, this.INIT_SQL, this.METHOD_UPDATE_WHERE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InitSql) Insert(values IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.INSERT_INTO)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.VALUES)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.FAILED_SQL_STATEMENT + fmt.Sprint(stringBuffer), this.INIT_SQL, this.INSERT, e)
            }
                }
            }()
            var size int = values.Size(); _ = size
            for i := int(0); i < size - 1; i++ {
                var value string = this.GetValue(any(values.Get(i)).(string)); _ = value
                stringBuffer.Append(value)
                stringBuffer.Append(this.SqlStrings.SINGLE_QUOTE_COMMA_SEP)
            }
            var value string = this.GetValue(any(values.Get(values.Size() - 1)).(string)); _ = value
            stringBuffer.Append(value)
            stringBuffer.Append(this.INSERT_END)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                PreLogUtil_Put(this.SUCCESS_SQL_STATEMENT + sqlStatement, this.INIT_SQL, this.INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *InitSql) GetValue(value string) string {
        if StringValidationUtil_GetInstance().IsEmpty(value) {
            return this.StringUtil.EMPTY_STRING
        } else {
            return value
        }
}


// throws Exception, SQLException
func (this *InitSql) ExecuteSQLStatement(statement string) *ResultSet {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *ResultSet) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.INIT_SQL, "executeSQLStatement()", e)
            }
            panic(e)
                }
            }()
            if this.Conn == nil {
                this.Initialize()
            }
            var stmt *Statement = this.Conn.CreateStatement(); _ = stmt
            stmt.Execute(statement)
            var rset *ResultSet = stmt.GetResultSet(); _ = rset
            stmt.Close()
            tryControl1, tryResult1 = 1, rset; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws SQLException
func (this *InitSql) CreateConnection() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    se := Java_ToError(recovered)
                    _ = se
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.INIT_SQL, "createConnection()", se)
            panic(se)
                }
            }()
            if this.UseridAndPassword == true {
                this.Conn = DriverManager_GetConnection(this.DatabaseConnectionInfoInterface.GetUrl(), this.Userid, this.Password)
            } else {
                this.Conn = DriverManager_GetConnection(this.DatabaseConnectionInfoInterface.GetUrl())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InitSql) Initialize() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    se := Java_ToError(recovered)
                    _ = se
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().SQLLOGGINGERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.INIT_SQL, "initialize()", se)
            }
            panic(se)
                }
            }()
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().SQLLOGGINGERROR) {
                    PreLogUtil_PutOE("LoadDriver Failed: " + fmt.Sprint(this.DatabaseConnectionInfoInterface.GetJdbcDriver()), this.INIT_SQL, "initialize()", e)
                }
                panic(e)
                    }
                }()
                var jdbcDriver string = this.DatabaseConnectionInfoInterface.GetJdbcDriver(); _ = jdbcDriver
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Loading DbConnnectionInfo: ").Append(reflect.TypeOf(this.DatabaseConnectionInfoInterface).String()).Append(" Driver: ").Append(jdbcDriver)), this.INIT_SQL, "initialize()")
                Class_ForName(jdbcDriver).NewInstance()
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            if this.Userid == "" && this.Password == "" {
                this.UseridAndPassword = true
            }
            this.CreateConnection()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InitSql) GetDatabaseConnectionInfoInterface() DatabaseConnectionInfoInterface {
        return this.DatabaseConnectionInfoInterface
}

