/*
 * 
 * To change this license header, choose License Headers in Project Properties. 
 * To change this template file, choose Tools | Templates 
 * and open the template in the editor.  
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

type AbDatabaseManagement struct {
    AbSqlBean
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SqlCommandLog *StringMaker
    GRANT_ALL string
    DOT_STAR string
    TO string
    AT_LOCALHOST string
    IDENTIFIED_BY string
    WITH_GRANT_OPTION string
    FAILED string
    GET_HOST_NAME string
    METHOD_ADD_USER string
    METHOD_ADD_USERS string
    METHOD_ADD_DB string
    METHOD_ADD_DATABASES string
    METHOD_ADD_TABLES string
    UNABLE_TO_CREATE_USER string
    UNABLE_TO_CREATE_DATABASE string
    UNABLE_TO_CREATE_TABLES string
    ADD_TABLES_RESULTS_LABEL string
}


func (this *AbDatabaseManagement) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SqlCommandLog = NewStringMaker()
    this.GRANT_ALL = "GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX, ALTER ON "
    this.DOT_STAR = ".*"
    this.TO = " TO "
    this.AT_LOCALHOST = "@localhost"
    this.IDENTIFIED_BY = " IDENTIFIED BY '"
    this.WITH_GRANT_OPTION = "' WITH GRANT OPTION"
    this.FAILED = "Failed"
    this.GET_HOST_NAME = "getHostName()"
    this.METHOD_ADD_USER = "addUser()"
    this.METHOD_ADD_USERS = "addUsers()"
    this.METHOD_ADD_DB = "addDb()"
    this.METHOD_ADD_DATABASES = "addDatabases()"
    this.METHOD_ADD_TABLES = "addTables()"
    this.UNABLE_TO_CREATE_USER = "Unable to Create User: "
    this.UNABLE_TO_CREATE_DATABASE = "Unable to Create Database: "
    this.UNABLE_TO_CREATE_TABLES = "Unable to Create Tables"
    this.ADD_TABLES_RESULTS_LABEL = "Add Table Results: "
}


func NewAbDatabaseManagement(databaseConnectionInfoInterface DatabaseConnectionInfoInterface) *AbDatabaseManagement {
    this := &AbDatabaseManagement{}
        this.AbSqlBean = *NewAbSqlBean(any(databaseConnectionInfoInterface).(*DbConnectionInfo))
        this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewAbDatabaseManagement(connectionInfo *DbConnectionInfo) *AbDatabaseManagement {
    this := &AbDatabaseManagement{}
        this.AbSqlBean = *NewAbSqlBean(connectionInfo)
        this.initFields()
    return this
}


func (this *AbDatabaseManagement) GetHostName() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.FAILED, this, this.GET_HOST_NAME, e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var addr *InetAddress = InetAddress_GetLocalHost(); _ = addr
            tryControl1, tryResult1 = 1, addr.GetHostName(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AbDatabaseManagement) AddDbUser(hostName string, db string, userName string, password string) bool {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.GRANT_ALL)
        stringBuffer.Append(db)
        stringBuffer.Append(this.DOT_STAR)
        stringBuffer.Append(this.TO)
        stringBuffer.Append(userName)
        stringBuffer.Append(hostName)
        stringBuffer.Append(this.IDENTIFIED_BY)
        stringBuffer.Append(password)
        stringBuffer.Append(this.WITH_GRANT_OPTION)
        var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
        this.SqlCommandLog.Append(sqlStatement)
        this.SqlCommandLog.Append(this.CommonSeps.NEW_LINE)
        this.AbSqlBean.ExecuteSQLStatement(sqlStatement)
        return true
}


// Java overloaded method: Go needs a different name
func (this *AbDatabaseManagement) AddDbUser(db string, userName string, password string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.UNABLE_TO_CREATE_USER + userName, this, this.METHOD_ADD_USER, e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.AddDbUser(this.AT_LOCALHOST, db, userName, password)
            var hostName string = this.GetHostName(); _ = hostName
            if hostName != "" {
                this.AddDbUser(hostName, db, userName, password)
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbDatabaseManagement) AddDb(db string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.UNABLE_TO_CREATE_DATABASE + db, this, this.METHOD_ADD_DB, e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var sqlStatement string = fmt.Sprint(this.SqlStrings.CREATE_DATABASE) + db; _ = sqlStatement
            this.SqlCommandLog.Append(sqlStatement)
            this.SqlCommandLog.Append(this.CommonSeps.NEW_LINE)
            this.AbSqlBean.ExecuteSQLStatement(sqlStatement)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

