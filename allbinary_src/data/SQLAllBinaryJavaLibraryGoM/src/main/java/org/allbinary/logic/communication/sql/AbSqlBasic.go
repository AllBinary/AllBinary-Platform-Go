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

type AbSqlBasic struct {
    LogUtil *LogUtil
    SqlConnectionPool *SqlConnectionPool
    DatabaseConnectionInfoInterface *DbConnectionInfo
    ConnectAttemptCounter int
    Conn *Connection
    StringUtil *StringUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    StdUtil *StdUtil
    SqlTypeStrings *SqlTypeStrings
    SqlStrings *SqlStrings
    INSERT string
    UPDATE string
    FAILED_SQL_STATEMENT string
    SUCCESS_SQL_STATEMENT string
    METHOD_EXECUTED_SQL_STATEMENT string
    METHOD_CREATE string
    DATABASE_CREATED_LABEL string
    DATABASE_CREATION_FAILED_LABEL string
    SQL_CONNECTION_FAILED string
    SQL_CONNECTION_RETRYING string
    SQL_CONNECTION_RETRY string
    METHOD_CREATE_CONNECTION string
    METHOD_INITIALIZE string
    LOAD_JDBC_DRIVER_FAILED_LABEL string
}


func (this *AbSqlBasic) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SqlConnectionPool = SqlConnectionPool_GetInstance()
    
    
    
    this.StringUtil = StringUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.StdUtil = StdUtil_GetInstance()
    this.SqlTypeStrings = SqlTypeStrings_GetInstance()
    this.SqlStrings = SqlStrings_GetInstance()
    this.INSERT = "insert"
    this.UPDATE = "update"
    this.FAILED_SQL_STATEMENT = "Failed\nSQL Statement: "
    this.SUCCESS_SQL_STATEMENT = "Success\nSQL Statement: "
    this.METHOD_EXECUTED_SQL_STATEMENT = "executedSQLStatement"
    this.METHOD_CREATE = "create"
    this.DATABASE_CREATED_LABEL = "Database Created: "
    this.DATABASE_CREATION_FAILED_LABEL = "Database Creation Failed: "
    this.SQL_CONNECTION_FAILED = "SQL Connection Failed"
    this.SQL_CONNECTION_RETRYING = "SQL Connection Retrying"
    this.SQL_CONNECTION_RETRY = "SQL Connection Retry"
    this.METHOD_CREATE_CONNECTION = "createConnection()"
    this.METHOD_INITIALIZE = "initialize"
    this.LOAD_JDBC_DRIVER_FAILED_LABEL = "Load JDBC Driver Failed: "
}


func NewAbSqlBasic(databaseConnectionInfoInterface *DbConnectionInfo) *AbSqlBasic {
    this := &AbSqlBasic{}
    this.initFields()
        this.DatabaseConnectionInfoInterface = databaseConnectionInfoInterface
        this.ConnectAttemptCounter = 0
        this.Conn = nil
    return this
}


// throws SQLException
func (this *AbSqlBasic) ExecuteSQLStatement(statement string) *ResultSet {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *ResultSet) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + statement, this, this.METHOD_EXECUTED_SQL_STATEMENT, e)
            }
            this.Conn.Close()
            this.Conn = nil
            panic(e)
                }
            }()
            this.Initialize()
            var stmt *Statement = this.Conn.CreateStatement(); _ = stmt
            stmt.Execute(statement)
            var rset *ResultSet = stmt.GetResultSet(); _ = rset
            stmt.Close()
            var tempConnection *Connection = this.Conn; _ = tempConnection
            this.Conn = nil
            this.SqlConnectionPool.Add(this.GetDatabaseConnectionInfoInterface().GetUrl(), tempConnection)
            tryControl1, tryResult1 = 1, rset; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlBasic) Create(databaseName string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.DATABASE_CREATION_FAILED_LABEL + databaseName, this, this.METHOD_CREATE, e)
            }
            tryControl1, tryResult1 = 1, this.DATABASE_CREATION_FAILED_LABEL + databaseName; return
                }
            }()
            this.ExecuteSQLStatement(fmt.Sprint(this.SqlStrings.CREATE_DATABASE) + databaseName)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.DATABASE_CREATED_LABEL + databaseName, this, this.METHOD_CREATE)
            }
            tryControl1, tryResult1 = 1, databaseName + fmt.Sprint(this.SqlStrings.CREATE_RETURN); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlBasic) CreateConnection() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    se := Java_ToError(recovered)
                    _ = se
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.SQL_CONNECTION_RETRYING, this, this.METHOD_CREATE_CONNECTION, se)
            }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                    this.LogUtil.Put(this.SQL_CONNECTION_RETRY + se.Error(), this, this.METHOD_CREATE_CONNECTION, e)
                }
                    }
                }()
                Thread_CurrentThread().Sleep(2000)
                if this.ConnectAttemptCounter < 10 {
                    this.ConnectAttemptCounter++
                    this.Initialize()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
                }
            }()
            this.Conn = this.SqlConnectionPool.Get(this.GetDatabaseConnectionInfoInterface().GetUrl())
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbSqlBasic) Initialize() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    se := Java_ToError(recovered)
                    _ = se
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.SQL_CONNECTION_FAILED, this, this.METHOD_INITIALIZE, se)
            }
                }
            }()
            {
                if tryControl2 := func() (tryControl2 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                        this.LogUtil.Put(this.LOAD_JDBC_DRIVER_FAILED_LABEL + fmt.Sprint(this.GetDatabaseConnectionInfoInterface().GetJdbcDriver()), this, this.METHOD_INITIALIZE, e)
                    }
                        }
                    }()
                    var loader *ClassLoader = Thread_CurrentThread().GetContextClassLoader(); _ = loader
                    var jdbcDriverClass reflect.Type = loader.LoadClass(this.GetDatabaseConnectionInfoInterface().GetJdbcDriver()); _ = jdbcDriverClass
                    jdbcDriverClass.NewInstance()
                    return
                }(); tryControl2 == 1 {
                    tryControl1 = 1; return
                }
            }
            this.CreateConnection()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbSqlBasic) GetDatabaseConnectionInfoInterface() DatabaseConnectionInfoInterface {
        return this.DatabaseConnectionInfoInterface
}


func (this *AbSqlBasic) SetDatabaseConnectionInfoInterface(databaseConnectionInfoInterface *DbConnectionInfo) {
        this.DatabaseConnectionInfoInterface = databaseConnectionInfoInterface
}

