/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type IdGeneratorEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
    TableData string
    NAME string
    VALUE string
}

// Implements IdGeneratorEntityInterface


func (this *IdGeneratorEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "idgenerator"
    
    this.NAME = "NAME"
    this.VALUE = "VALUE"
}


func NewIdGeneratorEntity() *IdGeneratorEntity {
    this := &IdGeneratorEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.TableData = fmt.Sprint(NewStringMaker().Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(this.NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(this.VALUE).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(this.NAME).Append(this.SqlStrings.END))
        this.AbSqlBean.SetTableName(this.TableName)
    return this
}


// throws Exception
func (this *IdGeneratorEntity) Get(name string) int64 {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(this.NAME, name)
        var hashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = hashMap
        if strings.Compare((any(hashMap.Get(this.NAME)).(string)), name) != 0 {
            panic(errors.New(fmt.Sprint("results do not match")))
        }
        var value string = any(hashMap.Get(this.VALUE)).(string); _ = value
        return func() int64 { value, _ := strconv.ParseInt(value, 10, 64); return value }()
}


func (this *IdGeneratorEntity) Insert(values IBasicArrayList) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.FAILURE, this, AbSqlBean_INSERT, e)
                }
            }()
            this.AbSqlBean.Insert(values)
            this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *IdGeneratorEntity) Delete(value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.Delete, e)
                }
            }()
            this.AbSqlBean.DeleteWhere(this.NAME, value)
            this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *IdGeneratorEntity) Update_1(name string, value int64) {
        var map_ *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = map_
        map_.Put(this.NAME, name)
        map_.Put(this.VALUE, value)
        this.Update(map_)
}

// Java overloads of IdGeneratorEntity.update
func (this *IdGeneratorEntity) Update(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            this.Update_1(Java_Arg[string](args[0]), Java_Arg[int64](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            this.Update_2(Java_Arg[*HashMap](args[0]))
            return
        }
    }
    panic("No Java overload of IdGeneratorEntity.update for the arguments")
}


func (this *IdGeneratorEntity) Update_2(hashMap *HashMap) {
        this.AbSqlBean.UpdateWhere(this.NAME, any(hashMap.Get(this.NAME)).(string), hashMap)
}


func (this *IdGeneratorEntity) CreateTable() string {
        var returnStr string = this.AbSqlBean.CreateTable(this.TableData); _ = returnStr
        return returnStr
}


func (this *IdGeneratorEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

