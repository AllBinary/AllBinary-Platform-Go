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

type SqlConnectionPool struct {
    LogUtil *LogUtil
    StdUtil *StdUtil
    ConnectionHashMap *HashMap
    FIRST_NEW_CONNECTION_CREATED string
    FIRST string
    NEW_CONNECTION_FOR string
    CREATED string
    NUMBER_OF_SQL_CONNECTIONS_FOR string
    NUMBER_OF_SQL_CONNECTION_VECTORS string
    IS string
    CONNECTION_ALLREADY_CLOSED string
    METHOD_GET string
    METHOD_ADD string
}

var SqlConnectionPool_Instance *SqlConnectionPool = NewSqlConnectionPool()

func (this *SqlConnectionPool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StdUtil = StdUtil_GetInstance()
    this.ConnectionHashMap = nil
    this.FIRST_NEW_CONNECTION_CREATED = "First New Connection Created: "
    this.FIRST = "First "
    this.NEW_CONNECTION_FOR = "New Connection For "
    this.CREATED = " Created"
    this.NUMBER_OF_SQL_CONNECTIONS_FOR = "Number Of Sql Connections for: "
    this.NUMBER_OF_SQL_CONNECTION_VECTORS = "Number Of Sql Connection Vectors: "
    this.IS = " is "
    this.CONNECTION_ALLREADY_CLOSED = "Connection AllReady Closed"
    this.METHOD_GET = "get()"
    this.METHOD_ADD = "add()"
}


func SqlConnectionPool_GetInstance() *SqlConnectionPool {
        return SqlConnectionPool_Instance
}


func NewSqlConnectionPool() *SqlConnectionPool {
    this := &SqlConnectionPool{}
    this.initFields()
    return this
}


// throws SQLException
// synchronized
func (this *SqlConnectionPool) Get(url string) IConnection {
        if this.ConnectionHashMap == nil {
            this.ConnectionHashMap = this.StdUtil.CreateHashMap()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
                this.LogUtil.PutF(this.FIRST_NEW_CONNECTION_CREATED + url, this, this.METHOD_GET)
            }
            return DriverManager_GetConnection(url)
        } else {
            var connectionVector IBasicArrayList = any(this.ConnectionHashMap.Get(url)).(IBasicArrayList); _ = connectionVector
            if connectionVector == nil {
                connectionVector = NewBasicArrayListD()
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(this.FIRST)
                    stringBuffer.Append(this.NEW_CONNECTION_FOR)
                    stringBuffer.Append(url)
                    stringBuffer.Append(this.CREATED)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET)
                }
                return DriverManager_GetConnection(url)
            } else if connectionVector.Size() == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.NEW_CONNECTION_FOR).Append(url).Append(this.CREATED)), this, this.METHOD_GET)
                }
                return DriverManager_GetConnection(url)
            } else {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                var size int = connectionVector.Size(); _ = size
                for i := int(0); i < size; i++ {
                    var sqlConnection IConnection = any(connectionVector.Get(i)).(IConnection); _ = sqlConnection
                    if !sqlConnection.IsClosed() {
                        connectionVector.Remove(sqlConnection)
                        this.ConnectionHashMap.Put(url, connectionVector)
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
                            stringBuffer.Delete(0, stringBuffer.Length())
                            stringBuffer.Append(this.NUMBER_OF_SQL_CONNECTIONS_FOR)
                            stringBuffer.Append(url)
                            stringBuffer.Append(this.IS)
                            stringBuffer.Appendint(connectionVector.Size())
                            this.LogUtil.PutF(this.NUMBER_OF_SQL_CONNECTION_VECTORS + fmt.Sprint(this.ConnectionHashMap.Size()), this, this.METHOD_GET)
                            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET)
                        }
                        return any(sqlConnection).(IConnection)
                    }
                }
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.NEW_CONNECTION_FOR)
            stringBuffer.Append(url)
            stringBuffer.Append(this.CREATED)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET)
        }
        return DriverManager_GetConnection(url)
}


// throws SQLException
// synchronized
func (this *SqlConnectionPool) Add(url string, sqlConnection IConnection) {
        if !sqlConnection.IsClosed() {
            var connectionVector IBasicArrayList; _ = connectionVector
            if this.ConnectionHashMap == nil {
                this.ConnectionHashMap = this.StdUtil.CreateHashMap()
                connectionVector = NewBasicArrayListD()
                connectionVector.Add(sqlConnection)
            } else {
                connectionVector = any(this.ConnectionHashMap.Get(url)).(IBasicArrayList)
                if connectionVector == nil {
                    connectionVector = NewBasicArrayListD()
                }
                connectionVector.Add(sqlConnection)
                this.ConnectionHashMap.Put(url, connectionVector)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append(this.NUMBER_OF_SQL_CONNECTIONS_FOR)
                stringBuffer.Append(url)
                stringBuffer.Append(this.IS)
                stringBuffer.Appendint(connectionVector.Size())
                this.LogUtil.PutF(this.NUMBER_OF_SQL_CONNECTION_VECTORS + fmt.Sprint(this.ConnectionHashMap.Size()), this, this.METHOD_ADD)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_ADD)
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGPOOL) {
            this.LogUtil.PutF(this.CONNECTION_ALLREADY_CLOSED, this, this.METHOD_ADD)
        }
}

