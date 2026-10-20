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

type CategoryEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements CategoryEntityInterface


func (this *CategoryEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "categories"
}




// Java overloads of new CategoryEntity
func NewCategoryEntity(args ...any) *CategoryEntity {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[CategoryFactoryInterface](args[0], pass) {
            return NewCategoryEntity_1(Java_Arg[CategoryFactoryInterface](args[0]))
        }
        if len(args) == 0 {
            return NewCategoryEntity_2()
        }
    }
    panic("No Java overload of new CategoryEntity for the arguments")
}
func NewCategoryEntity_1(categoryFactoryInterface CategoryFactoryInterface) *CategoryEntity {
    this := &CategoryEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewInventoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func NewCategoryEntity_2() *CategoryEntity {
    this := &CategoryEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewInventoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *CategoryEntity) Insert(category ICategory) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, AbSqlBean_INSERT, e)
            }
                }
            }()
            var categoryVector IBasicArrayList = category.ToVector(); _ = categoryVector
            var values IBasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(categoryVector.Get(0))
            values.Add(categoryVector.Get(1))
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CategoryEntity) Delete(value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.Delete, e)
            }
                }
            }()
            this.AbSqlBean.DeleteWhere(CategoryData_GetInstance().NAME, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CategoryEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(CategoryData_GetInstance().NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(CategoryData_GetInstance().LEVEL)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(CategoryData_GetInstance().NAME)
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *CategoryEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *CategoryEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(CategoryData_GetInstance().NAME, any(updatedValues.Get(CategoryData_GetInstance().NAME)).(string), updatedValues)
}


func (this *CategoryEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

