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

type AbSqlRow struct {
    AbSqlColumn
    LogUtil *LogUtil
    EQUAL_QUOTE string
    ESCAPE_QUOTES string
    GET_ROW string
    METHOD_GET_ROW string
    METHOD_GET_ROWS string
    METHOD_GET_ALL_ROWS string
    METHOD_UPDATE_WHERE string
    METHOD_GET_ROWS_WHERE_BETWEEN string
    METHOD_DELETE_WHERE string
    ROW_VALUE_LABEL string
    ROW_VALUES_LABEL string
    NO_RESULTS_IN_RESULT_SET string
    COLUMN_NAME_LABEL string
    INSERT_END string
}

// The Java class AbSqlRow has subclasses so its variables are this interface
type IAbSqlRow interface {
    IAbSqlColumn
    AbSqlRow_() *AbSqlRow
    GetValue(value string) string
    UpdateWhere(args ...any)
    DeleteWhere(args ...any)
    Insert(values IBasicArrayList)
    GetRow(keysAndValues *HashMap) *HashMap
    GetRows(args ...any) IBasicArrayList
    GetAllRows() IBasicArrayList
    GetRowsWhereBetween(args ...any) IBasicArrayList
}

func (this *AbSqlRow) AbSqlRow_() *AbSqlRow {
    return this
}


func (this *AbSqlRow) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.EQUAL_QUOTE = "=\""
    this.ESCAPE_QUOTES = "\\\""
    this.GET_ROW = "Row Value: "
    this.METHOD_GET_ROW = "getRow()"
    this.METHOD_GET_ROWS = "getRows()"
    this.METHOD_GET_ALL_ROWS = "getAllRows()"
    this.METHOD_UPDATE_WHERE = "updateWhere()"
    this.METHOD_GET_ROWS_WHERE_BETWEEN = "getRowsWhereBetween()"
    this.METHOD_DELETE_WHERE = "deleteWhere()"
    this.ROW_VALUE_LABEL = "Row Value: "
    this.ROW_VALUES_LABEL = "Row Values: "
    this.NO_RESULTS_IN_RESULT_SET = "No Results in ResultSet"
    this.COLUMN_NAME_LABEL = "columnName: "
    this.INSERT_END = "')"
}


func NewAbSqlRow(databaseConnectionInfoInterface IDbConnectionInfo) *AbSqlRow {
    this := &AbSqlRow{}
        this.AbSqlColumn = *NewAbSqlColumn(databaseConnectionInfoInterface)
        this.initFields()
    return this
}


// synchronized
func (this *AbSqlRow) GetValue(value string) string {
        if StringValidationUtil_GetInstance().IsEmpty(value) {
            return this.StringUtil.EMPTY_STRING
        } else {
            return value
        }
}


// synchronized
func (this *AbSqlRow) UpdateWhere_1(key string, value string, updatedKeyValuePairs *HashMap) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.UPDATE)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.SET)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_UPDATE_WHERE, e)
            }
                }
            }()
            var keyArray []any = updatedKeyValuePairs.KeySet().ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for i := int(0); i < size; i++ {
                var columnName string = fmt.Sprint(keyArray[i]); _ = columnName
                stringBuffer.Append(this.CommonSeps.SPACE)
                stringBuffer.Append(columnName)
                stringBuffer.Append(this.EQUAL_QUOTE)
                var columnValue string = any(updatedKeyValuePairs.Get(columnName)).(string); _ = columnValue
                if columnValue == "" {
                    columnValue = this.StringUtil.EMPTY_STRING
                } else {
                    columnValue = NewReplace(this.SqlStrings.CLOSE_QUOTE, this.ESCAPE_QUOTES).All(columnValue)
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
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.METHOD_UPDATE_WHERE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of AbSqlRow.updateWhere
func (this *AbSqlRow) UpdateWhere(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            this.UpdateWhere_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[*HashMap](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            this.UpdateWhere_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HashMap](args[1]))
            return
        }
    }
    panic("No Java overload of AbSqlRow.updateWhere for the arguments")
}


// throws Exception
// synchronized
func (this *AbSqlRow) UpdateWhere_2(whereKeyValuePairs *HashMap, updatedKeyValuePairs *HashMap) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.UPDATE)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.SET)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_UPDATE_WHERE, e)
            }
            panic(e)
                }
            }()
            var keyArray []any = updatedKeyValuePairs.KeySet().ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for i := int(0); i < size; i++ {
                var columnName string = fmt.Sprint(keyArray[i]); _ = columnName
                stringBuffer.Append(this.CommonSeps.SPACE)
                stringBuffer.Append(columnName)
                stringBuffer.Append(this.EQUAL_QUOTE)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.COLUMN_NAME_LABEL + columnName, this, this.METHOD_UPDATE_WHERE)
                }
                var columnValue string = any(updatedKeyValuePairs.Get(columnName)).(string); _ = columnValue
                if columnValue == "" {
                    columnValue = this.StringUtil.EMPTY_STRING
                } else {
                    columnValue = NewReplace(this.SqlStrings.CLOSE_QUOTE, this.ESCAPE_QUOTES).All(columnValue)
                }
                stringBuffer.Append(this.GetValue(columnValue))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if i < size - 1 {
                    stringBuffer.Append(this.CommonSeps.COMMA)
                }
            }
            stringBuffer.Append(this.SqlStrings.WHERE)
            var set Set = whereKeyValuePairs.KeySet(); _ = set
            var keyArray2 []any = set.ToArray(); _ = keyArray2
            var size2 int = len(keyArray2); _ = size2
            for i := int(0); i < size2; i++ {
                var key string = any(keyArray2[i]).(string); _ = key
                var value string = any(whereKeyValuePairs.Get(key)).(string); _ = value
                {
                    stringBuffer.Append(key)
                    stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                    stringBuffer.Append(this.GetValue(value))
                    stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                    if i < size2 - 1 {
                        stringBuffer.Append(this.SqlStrings.AND)
                    }
                }
            }
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.METHOD_UPDATE_WHERE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *AbSqlRow) DeleteWhere_1(key string, value string) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.DELETE)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        stringBuffer.Append(key)
        stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
        stringBuffer.Append(value)
        stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + fmt.Sprint(stringBuffer), this, this.METHOD_DELETE_WHERE, e)
            }
                }
            }()
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.METHOD_DELETE_WHERE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of AbSqlRow.deleteWhere
func (this *AbSqlRow) DeleteWhere(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.DeleteWhere_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            this.DeleteWhere_2(Java_Arg[*HashMap](args[0]))
            return
        }
    }
    panic("No Java overload of AbSqlRow.deleteWhere for the arguments")
}


// synchronized
func (this *AbSqlRow) DeleteWhere_2(keysAndValues *HashMap) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.DELETE)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + fmt.Sprint(stringBuffer), this, this.METHOD_DELETE_WHERE, e)
            }
                }
            }()
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
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.METHOD_DELETE_WHERE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbSqlRow) Insert(values IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.INSERT_INTO)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.VALUES)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + fmt.Sprint(stringBuffer), this, AbSqlColumn_INSERT, e)
            }
                }
            }()
            for i := int(0); i < values.Size() - 1; i++ {
                var value string = this.GetValue(any(values.Get(i)).(string)); _ = value
                value = NewReplace(this.SqlStrings.ESCAPE, this.SqlStrings.DOUBLE_ESCAPE).All(value)
                stringBuffer.Append(value)
                stringBuffer.Append(this.SqlStrings.SINGLE_QUOTE_COMMA_SEP)
            }
            var value string = this.GetValue(any(values.Get(values.Size() - 1)).(string)); _ = value
            value = NewReplace(this.SqlStrings.ESCAPE, this.SqlStrings.DOUBLE_ESCAPE).All(value)
            stringBuffer.Append(value)
            stringBuffer.Append(this.INSERT_END)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            this.ExecuteSQLStatement(sqlStatement)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SUCCESS_SQL_STATEMENT) + sqlStatement, this, this.INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AbSqlRow) GetRow(keysAndValues *HashMap) *HashMap {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HashMap) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_GET_ROW, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var result *HashMap = nil; _ = result
            var set Set = keysAndValues.KeySet(); _ = set
            var keyArray []any = set.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for index := int(0); index < size; index++ {
                var key string = any(keyArray[index]).(string); _ = key
                var value string = any(keysAndValues.Get(key)).(string); _ = value
                stringBuffer.Append(key)
                stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                stringBuffer.Append(this.GetValue(value))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if index < size - 1 {
                    stringBuffer.Append(this.SqlStrings.AND)
                }
            }
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SqlStrings.SQL_STATEMENT_LABEL) + sqlStatement, this, this.METHOD_GET_ROW)
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
                    this.LogUtil.PutF(this.ROW_VALUE_LABEL + fmt.Sprint(result), this, this.METHOD_GET_ROW)
                }
                tryControl1, tryResult1 = 1, result; return
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.PutF(this.NO_RESULTS_IN_RESULT_SET, this, this.METHOD_GET_ROW)
            }
            tryControl1, tryResult1 = 1, nil; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlRow) GetRows_1(keysAndValues *HashMap) IBasicArrayList {
        return this.GetRows(keysAndValues, StringUtil_GetInstance().EMPTY_STRING)
}

// Java overloads of AbSqlRow.getRows
func (this *AbSqlRow) GetRows(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return this.GetRows_1(Java_Arg[*HashMap](args[0]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetRows_2(Java_Arg[*HashMap](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of AbSqlRow.getRows for the arguments")
}


func (this *AbSqlRow) GetRows_2(keysAndValues *HashMap, more string) IBasicArrayList {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_GET_ROWS, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var rows IBasicArrayList = NewBasicArrayListD(); _ = rows
            var set Set = keysAndValues.KeySet(); _ = set
            var keyArray []any = set.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for index := int(0); index < size; index++ {
                var key string = any(keyArray[index]).(string); _ = key
                var value string = any(keysAndValues.Get(key)).(string); _ = value
                stringBuffer.Append(key)
                stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                stringBuffer.Append(this.GetValue(value))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if index < size - 1 {
                    stringBuffer.Append(this.SqlStrings.AND)
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.SqlStrings.SQL_STATEMENT_LABEL + stringBuffer, this, this.METHOD_GET_ROWS)
            }
            stringBuffer.Append(more)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var resultSetMetaData *ResultSetMetaData = rset.GetMetaData(); _ = resultSetMetaData
            for rset.Next() {
                var result *HashMap = this.StdUtil.CreateHashMap(); _ = result
                var columnCount int = resultSetMetaData.GetColumnCount(); _ = columnCount
                for index := int(1); index <= columnCount; index++ {
                    var columnName string = resultSetMetaData.GetColumnName(index); _ = columnName
                    var field string = rset.GetString(columnName); _ = field
                    result.Put(columnName, field)
                }
                var resultAsString string = fmt.Sprint(result); _ = resultAsString
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.GET_ROW + resultAsString, this, this.METHOD_GET_ROWS)
                }
                rows.Add(result)
            }
            tryControl1, tryResult1 = 1, rows; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlRow) GetAllRows() IBasicArrayList {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.GetTableName())
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_GET_ALL_ROWS, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.SqlStrings.SQL_STATEMENT_LABEL + stringBuffer, this, this.METHOD_GET_ALL_ROWS)
            }
            var rows IBasicArrayList = NewBasicArrayListD(); _ = rows
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var resultSetMetaData *ResultSetMetaData = rset.GetMetaData(); _ = resultSetMetaData
            for rset.Next() {
                var result *HashMap = this.StdUtil.CreateHashMap(); _ = result
                var columnCount int = resultSetMetaData.GetColumnCount(); _ = columnCount
                for index := int(1); index <= columnCount; index++ {
                    var columnName string = resultSetMetaData.GetColumnName(index); _ = columnName
                    var field string = rset.GetString(columnName); _ = field
                    result.Put(columnName, field)
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.ROW_VALUES_LABEL + fmt.Sprint(result), this, this.METHOD_GET_ALL_ROWS)
                }
                rows.Add(result)
            }
            tryControl1, tryResult1 = 1, rows; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlRow) GetRowsWhereBetween_1(whereKeyValuePairs *HashMap, betweenColumn string, smallest string, largest string) IBasicArrayList {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.GetTableName())
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_GET_ROWS_WHERE_BETWEEN, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var rows IBasicArrayList = NewBasicArrayListD(); _ = rows
            stringBuffer.Append(this.SqlStrings.WHERE)
            var set Set = whereKeyValuePairs.KeySet(); _ = set
            var key string; _ = key
            var value string; _ = value
            var keyArray []any = set.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            for index := int(0); index < size; index++ {
                key = any(keyArray[index]).(string)
                value = any(whereKeyValuePairs.Get(key)).(string)
                stringBuffer.Append(key)
                stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                stringBuffer.Append(this.GetValue(value))
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                stringBuffer.Append(this.SqlStrings.AND)
            }
            stringBuffer.Append(betweenColumn)
            stringBuffer.Append(this.SqlStrings.MORE_THAN_QUOTE)
            stringBuffer.Append(smallest)
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            stringBuffer.Append(this.SqlStrings.AND)
            stringBuffer.Append(betweenColumn)
            stringBuffer.Append(this.SqlStrings.LESS_THAN_QUOTE)
            stringBuffer.Append(largest)
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SqlStrings.SQL_STATEMENT_LABEL) + sqlStatement, this, this.METHOD_GET_ROWS_WHERE_BETWEEN)
            }
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var resultSetMetaData *ResultSetMetaData = rset.GetMetaData(); _ = resultSetMetaData
            var result *HashMap; _ = result
            var columnNames IBasicArrayList; _ = columnNames
            var columnCount int; _ = columnCount
            var columnName string; _ = columnName
            var field string; _ = field
            for rset.Next() {
                result = this.StdUtil.CreateHashMap()
                columnNames = NewBasicArrayListD()
                columnCount = resultSetMetaData.GetColumnCount()
                for index := int(1); index <= columnCount; index++ {
                    columnName = resultSetMetaData.GetColumnName(index)
                    field = rset.GetString(columnName)
                    result.Put(columnName, field)
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.ROW_VALUE_LABEL + fmt.Sprint(result), this, this.METHOD_GET_ROWS_WHERE_BETWEEN)
                }
                rows.Add(result)
            }
            tryControl1, tryResult1 = 1, rows; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AbSqlRow.getRowsWhereBetween
func (this *AbSqlRow) GetRowsWhereBetween(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return this.GetRowsWhereBetween_1(Java_Arg[*HashMap](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetRowsWhereBetween_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of AbSqlRow.getRowsWhereBetween for the arguments")
}


func (this *AbSqlRow) GetRowsWhereBetween_2(betweenColumn string, smallest string, largest string) IBasicArrayList {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT_ALL_FROM)
        stringBuffer.Append(this.GetTableName())
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(this.FAILED_SQL_STATEMENT + stringBuffer, this, this.METHOD_GET_ROWS_WHERE_BETWEEN, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var rows IBasicArrayList = NewBasicArrayListD(); _ = rows
            stringBuffer.Append(this.SqlStrings.WHERE)
            stringBuffer.Append(betweenColumn)
            stringBuffer.Append(this.SqlStrings.MORE_THAN_QUOTE)
            stringBuffer.Append(smallest)
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            stringBuffer.Append(this.SqlStrings.AND)
            stringBuffer.Append(betweenColumn)
            stringBuffer.Append(this.SqlStrings.LESS_THAN_QUOTE)
            stringBuffer.Append(largest)
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(fmt.Sprint(this.SqlStrings.SQL_STATEMENT_LABEL) + sqlStatement, this, this.METHOD_GET_ROWS_WHERE_BETWEEN)
            }
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var resultSetMetaData *ResultSetMetaData = rset.GetMetaData(); _ = resultSetMetaData
            var columnName string; _ = columnName
            var field string; _ = field
            var result *HashMap; _ = result
            var columnNames IBasicArrayList; _ = columnNames
            var columnCount int; _ = columnCount
            for rset.Next() {
                result = this.StdUtil.CreateHashMap()
                columnNames = NewBasicArrayListD()
                columnCount = resultSetMetaData.GetColumnCount()
                for index := int(1); index <= columnCount; index++ {
                    columnName = resultSetMetaData.GetColumnName(index)
                    field = rset.GetString(columnName)
                    result.Put(columnName, field)
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.ROW_VALUE_LABEL + fmt.Sprint(result), this, this.METHOD_GET_ROWS_WHERE_BETWEEN)
                }
                rows.Add(result)
            }
            tryControl1, tryResult1 = 1, rows; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

