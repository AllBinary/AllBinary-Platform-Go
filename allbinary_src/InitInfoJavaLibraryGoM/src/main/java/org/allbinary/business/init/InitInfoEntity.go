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

type InitInfoEntity struct {
    InitSql
    NAME string
    NOTHING string
    TableName string
    TableData string
    NOT_IN_DB string
    METHOD_GET string
    METHOD_IS string
    METHOD_ADD string
    CREATED_SUCCESS string
    FAILED_TO_CREATE string
}


func (this *InitInfoEntity) initFields() {
    this.NAME = "InitInfoEntity"
    this.NOTHING = "NOTHING"
    this.TableName = "initdata"
    
    this.NOT_IN_DB = "Not In DB"
    this.METHOD_GET = "get()"
    this.METHOD_IS = "is()"
    this.METHOD_ADD = "add()"
    this.CREATED_SUCCESS = " Created Successfully"
    this.FAILED_TO_CREATE = "Failed to create "
}


func NewInitInfoEntity() *InitInfoEntity {
    this := &InitInfoEntity{}
        this.InitSql = *NewInitSql(NewUserDbInitInfo())
        this.initFields()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(this.NOTHING)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(InitInfo_GetInstance().TESTING)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(InitInfo_GetInstance().TESTHTMLPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(InitInfo_GetInstance().MAINPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(this.NOTHING)
        stringBuffer.Append(this.SqlStrings.END)
        this.TableData = fmt.Sprint(stringBuffer)
        this.SetTable(this.TableName)
    return this
}


func (this *InitInfoEntity) Get() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.NAME, this.METHOD_GET, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var keyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyAndValue
            keyAndValue.Put(this.NOTHING, this.NOTHING)
            var hashMap *HashMap = this.GetRow(keyAndValue); _ = hashMap
            if hashMap != nil {
                InitInfo_GetInstance().Set(hashMap)
                tryControl1, tryResult1 = 1, true; return
            } else {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                    PreLogUtil_Put(this.NOT_IN_DB, this.NAME, this.METHOD_GET)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitInfoEntity) Is() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.NAME, this.METHOD_IS, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var keyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyAndValue
            keyAndValue.Put(this.NOTHING, this.NOTHING)
            var hashMap *HashMap = this.GetRow(keyAndValue); _ = hashMap
            if hashMap != nil {
                tryControl1, tryResult1 = 1, true; return
            } else {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                    PreLogUtil_Put(this.NOT_IN_DB, this.NAME, this.METHOD_IS)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitInfoEntity) Update() {
        this.UpdateWhere(this.NOTHING, this.NOTHING, InitInfo_GetInstance().ToHashMap())
}


func (this *InitInfoEntity) Add() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this.NAME, this.METHOD_ADD, e)
            }
                }
            }()
            var values *BasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(this.NOTHING)
            values.Add(InitInfo_GetInstance().GetTesting())
            values.Add(InitInfo_GetInstance().GetTestHtmlPath())
            values.Add(InitInfo_GetInstance().GetMainPath())
            this.Insert(values)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InitInfoEntity) CreateTable() string {
        if this.InitSql.CreateTable(this.TableData) {
            return this.TableName + this.CREATED_SUCCESS
        } else {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.FAILED_TO_CREATE)
            stringBuffer.Append(this.TableData)
            stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
            return fmt.Sprint(stringBuffer)
        }
}

