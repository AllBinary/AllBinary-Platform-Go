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

type AbSqlColumn struct {
    AbSqlTable
    LogUtil *LogUtil
    LARGEST_INT_VALUE_IN_COLUMN string
    METHOD_GET_LARGETS_INTEGER_IN_COLUMN string
    METHOD_GET_COLUMN string
    METHOD_GET_COLUMN_WHERE string
    INVALID_RESULT string
}


func (this *AbSqlColumn) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.LARGEST_INT_VALUE_IN_COLUMN = "\nLargest Int Value in column: "
    this.METHOD_GET_LARGETS_INTEGER_IN_COLUMN = "getLargestIntegerInColumn"
    this.METHOD_GET_COLUMN = "getColumn"
    this.METHOD_GET_COLUMN_WHERE = "getColumnWhere"
    this.INVALID_RESULT = "-1"
}


func NewAbSqlColumn(databaseConnectionInfoInterface *DbConnectionInfo) *AbSqlColumn {
    this := &AbSqlColumn{}
        this.AbSqlTable = *NewAbSqlTable(databaseConnectionInfoInterface)
        this.initFields()
    return this
}


func (this *AbSqlColumn) GetLargestIntegerInColumnWhere(columnName string, key string, value string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        stringBuffer.Append(columnName)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        stringBuffer.Append(key)
        stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
        stringBuffer.Append(value)
        stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
        var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_GET_LARGETS_INTEGER_IN_COLUMN, e)
            }
            tryControl1, tryResult1 = 1, this.INVALID_RESULT; return
                }
            }()
            var largest int = 0; _ = largest
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            for rset.Next() {
                var intValue int = rset.GetInt(columnName); _ = intValue
                if intValue > largest {
                    largest = intValue
                }
            }
            var largestAsString string = fmt.Sprint(int(largest)); _ = largestAsString
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.SqlStrings.SQL_STATEMENT_LABEL)
                stringBuffer.Append(sqlStatement)
                stringBuffer.Append(this.LARGEST_INT_VALUE_IN_COLUMN)
                stringBuffer.Append(largestAsString)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET_LARGETS_INTEGER_IN_COLUMN)
            }
            tryControl1, tryResult1 = 1, largestAsString; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlColumn) GetColumn(columnName string) *BasicArrayList {
        var column *BasicArrayList = NewBasicArrayListD(); _ = column
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        stringBuffer.Append(columnName)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_GET_COLUMN, e)
            }
            tryControl1, tryResult1 = 1, column; return
                }
            }()
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            for rset.Next() {
                var field string = fmt.Sprint(rset.GetObject(columnName)); _ = field
                column.Add(field)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.SqlStrings.SQL_STATEMENT_LABEL)
                stringBuffer.Append(sqlStatement)
                stringBuffer.Append(this.SqlStrings.COLUMN_VALUE)
                stringBuffer.Append(fmt.Sprint(column))
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET_COLUMN)
            }
            tryControl1, tryResult1 = 1, column; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlColumn) GetColumnWhere(columnName string, key string, value string) *BasicArrayList {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        stringBuffer.Append(columnName)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        stringBuffer.Append(key)
        stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
        stringBuffer.Append(value)
        stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
        var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
        var column *BasicArrayList = NewBasicArrayListD(); _ = column
        var rset *ResultSet; _ = rset
        var field string; _ = field
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_GET_COLUMN_WHERE, e)
            }
            tryControl1, tryResult1 = 1, column; return
                }
            }()
            rset = this.ExecuteSQLStatement(sqlStatement)
            for rset.Next() {
                field = fmt.Sprint(rset.GetObject(columnName))
                column.Add(field)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.SqlStrings.SQL_STATEMENT_LABEL)
                stringBuffer.Append(sqlStatement)
                stringBuffer.Append(this.SqlStrings.COLUMN_VALUE)
                stringBuffer.Append(fmt.Sprint(column))
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET_COLUMN_WHERE)
            }
            tryControl1, tryResult1 = 1, column; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

