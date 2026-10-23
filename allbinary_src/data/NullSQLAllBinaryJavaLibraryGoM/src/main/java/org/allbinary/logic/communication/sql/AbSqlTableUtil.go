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

type AbSqlTableUtil struct {
    TABLE_LABEL string
    StringUtil *StringUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    SqlStrings *SqlStrings
}

var AbSqlTableUtil_Instance *AbSqlTableUtil = NewAbSqlTableUtil()

func (this *AbSqlTableUtil) initFields() {
    this.TABLE_LABEL = "Table: "
    this.StringUtil = StringUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.SqlStrings = SqlStrings_GetInstance()
}

func NewAbSqlTableUtil() *AbSqlTableUtil {
    this := &AbSqlTableUtil{}
    this.initFields()
    return this
}


func AbSqlTableUtil_GetInstance() *AbSqlTableUtil {
        return AbSqlTableUtil_Instance
}


// synchronized
func (this *AbSqlTableUtil) BackupTable(abSqlTable IAbSqlTable) string {
        var tableName string = abSqlTable.GetTableName(); _ = tableName
        return this.TABLE_LABEL + tableName + " Backup Failed"
}


// synchronized
func (this *AbSqlTableUtil) RestoreTable(abSqlTable IAbSqlTable, portion *Portion) string {
        var tableName string = abSqlTable.GetTableName(); _ = tableName
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.TABLE_LABEL)
        stringBuffer.Append(tableName)
        stringBuffer.Append(" Restoration Failed")
        return fmt.Sprint(stringBuffer)
}

