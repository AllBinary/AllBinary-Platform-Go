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

type AbSqlBean struct {
    AbSqlRow
    LogUtil *LogUtil
    METHOD_GET_FIELD string
    METHOD_IS_SUBSET_OF_ENTRY string
}

// The Java class AbSqlBean has subclasses so its variables are this interface
type IAbSqlBean interface {
    IAbSqlRow
    AbSqlBean_() *AbSqlBean
    GetField(args ...any) string
    IsSubsetOfEntry(key string, value string, columnsAndValues *HashMap) string
}

func (this *AbSqlBean) AbSqlBean_() *AbSqlBean {
    return this
}


func (this *AbSqlBean) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.METHOD_GET_FIELD = "getField"
    this.METHOD_IS_SUBSET_OF_ENTRY = "isSubsetOfEntry"
}


func NewAbSqlBean(databaseConnectionInfoInterface IDbConnectionInfo) *AbSqlBean {
    this := &AbSqlBean{}
        this.AbSqlRow = *NewAbSqlRow(databaseConnectionInfoInterface)
        this.initFields()
    return this
}


func (this *AbSqlBean) GetField_1(key string, value string, requestedField string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        stringBuffer.Append(requestedField)
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
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_GET_FIELD, e)
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement; return
                }
            }()
            var field string = this.StringUtil.EMPTY_STRING; _ = field
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            for rset.Next() {
                field = rset.GetString(requestedField)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append(this.SqlStrings.SQL_STATEMENT_LABEL)
                    stringBuffer.Append(sqlStatement)
                    stringBuffer.Append(this.SqlStrings.FIELD_VALUE)
                    stringBuffer.Append(field)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET_FIELD)
                }
                tryControl1, tryResult1 = 1, field; return
            }
            tryControl1, tryResult1 = 1, nil; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AbSqlBean.getField
func (this *AbSqlBean) GetField(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetField_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetField_2(Java_Arg[*HashMap](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of AbSqlBean.getField for the arguments")
}


func (this *AbSqlBean) GetField_2(keysAndValues *HashMap, requestedField string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        stringBuffer.Append(requestedField)
        stringBuffer.Append(this.SqlStrings.FROM)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.WHERE)
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_GET_FIELD, e)
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement; return
                }
            }()
            var field string = this.StringUtil.EMPTY_STRING; _ = field
            var set Set = keysAndValues.KeySet(); _ = set
            var keyArray []any = set.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            var key string; _ = key
            var value string; _ = value
            for i := int(0); i < size; i++ {
                key = any(keyArray[i]).(string)
                value = any(keysAndValues.Get(key)).(string)
                stringBuffer.Append(key)
                stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
                stringBuffer.Append(value)
                stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
                if i < size - 1 {
                    stringBuffer.Append(this.SqlStrings.AND)
                }
            }
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            for rset.Next() {
                field = rset.GetString(requestedField)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append(this.SqlStrings.SQL_STATEMENT_LABEL)
                    stringBuffer.Append(sqlStatement)
                    stringBuffer.Append(this.SqlStrings.FIELD_VALUE)
                    stringBuffer.Append(field)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.METHOD_GET_FIELD)
                }
                tryControl1, tryResult1 = 1, field; return
            }
            tryControl1, tryResult1 = 1, nil; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbSqlBean) IsSubsetOfEntry(key string, value string, columnsAndValues *HashMap) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.SELECT)
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGINGERROR) {
                this.LogUtil.Put(fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement, this, this.METHOD_IS_SUBSET_OF_ENTRY, e)
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(this.FAILED_SQL_STATEMENT) + sqlStatement; return
                }
            }()
            var keyArray []any = columnsAndValues.KeySet().ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            stringBuffer.Append(this.CommonSeps.SPACE)
            stringBuffer.Append(key)
            for i := int(0); i < size; i++ {
                stringBuffer.Append(this.CommonSeps.SPACE)
                stringBuffer.Append(fmt.Sprint(keyArray[i]))
            }
            stringBuffer.Append(this.SqlStrings.FROM)
            stringBuffer.Append(this.GetTableName())
            stringBuffer.Append(this.SqlStrings.WHERE)
            stringBuffer.Append(key)
            stringBuffer.Append(this.SqlStrings.EQUAL_QUOTE)
            stringBuffer.Append(value)
            stringBuffer.Append(this.SqlStrings.CLOSE_QUOTE)
            var sqlStatement string = fmt.Sprint(stringBuffer); _ = sqlStatement
            var rset *ResultSet = this.ExecuteSQLStatement(sqlStatement); _ = rset
            var columnName string; _ = columnName
            var field string; _ = field
            for rset.Next() {
                var keyArray2 []any = columnsAndValues.KeySet().ToArray(); _ = keyArray2
                var size2 int = len(keyArray2); _ = size2
                for i := int(0); i < size2; i++ {
                    columnName = fmt.Sprint(keyArray2[i])
                    field = fmt.Sprint(rset.GetObject(columnName))
                    if strings.Compare(field, any(columnsAndValues.Get(columnName)).(string)) != 0 {
                        tryControl1, tryResult1 = 1, org_Allbinary.Globals.GLOBALS2.NOTASUBSET; return
                    }
                }
            }
            if rset.GetRow() != 1 {
                tryControl1, tryResult1 = 1, org_Allbinary.Globals.GLOBALS2.NOTASUBSET; return
            } else {
                tryControl1, tryResult1 = 1, org_Allbinary.Globals.GLOBALS2.ISASUBSET; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

