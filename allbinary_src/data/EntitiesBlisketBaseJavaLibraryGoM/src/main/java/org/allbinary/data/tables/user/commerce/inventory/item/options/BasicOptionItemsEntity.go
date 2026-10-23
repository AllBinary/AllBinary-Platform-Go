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

type BasicOptionItemsEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    TableName string
}

// Implements BasicOptionItemsEntityInterface


func (this *BasicOptionItemsEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    this.TableName = "basicoptionitems"
}


func NewBasicOptionItemsEntity() *BasicOptionItemsEntity {
    this := &BasicOptionItemsEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewInventoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *BasicOptionItemsEntity) Insert(values IBasicArrayList) {
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
            this.AbSqlBean.Insert(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, AbSqlBean_INSERT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BasicOptionItemsEntity) Delete(value string) {
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
            this.AbSqlBean.DeleteWhere(this.BasicItemData.ID, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BasicOptionItemsEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(this.BasicItemData.ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(BasicOptionItemData_OPTION_ONE_TITLE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_DEFAULT_OPTION_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_DEFAULT_OPTION_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(BasicOptionItemData_OPTION_ONE_ONE_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_ONE_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_TWO_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_TWO_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_THREE_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_THREE_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_FOUR_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_FOUR_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_FIVE_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_FIVE_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_SIX_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_SIX_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_SEVEN_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_SEVEN_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_EIGHT_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_EIGHT_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_NINE_ITEM).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(BasicOptionItemData_OPTION_ONE_NINE_VALUE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(this.BasicItemData.ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *BasicOptionItemsEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *BasicOptionItemsEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(this.BasicItemData.ID, any(updatedValues.Get(this.BasicItemData.ID)).(string), updatedValues)
}


func (this *BasicOptionItemsEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

