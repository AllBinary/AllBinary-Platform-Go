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

type DownloadItemsEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    TableName string
}

// Implements DownloadItemsEntityInterface


func (this *DownloadItemsEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    this.TableName = "downloaditems"
}


func NewDownloadItemsEntity() *DownloadItemsEntity {
    this := &DownloadItemsEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewInventoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *DownloadItemsEntity) Insert(values *BasicArrayList) {
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


func (this *DownloadItemsEntity) Delete(value string) {
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
            this.AbSqlBean.DeleteWhere(DownloadItemData_ID, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DownloadItemsEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(DownloadItemData_ID, any(updatedValues.Get(DownloadItemData_ID)).(string), updatedValues)
}


func (this *DownloadItemsEntity) GetForItem(id string) *BasicArrayList {
        var returnVector *BasicArrayList = NewBasicArrayListD(); _ = returnVector
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(this.BasicItemData.ID, id)
        var vector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var hashMap *HashMap = any(vector.Get(index)).(*HashMap); _ = hashMap
            returnVector.Add(NewDownloadableItem(hashMap))
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("Found: " + fmt.Sprint(size), this, "getForItem")
        }
        return returnVector
}


// Java overloaded method: Go needs a different name
func (this *DownloadItemsEntity) GetForItem(id string, downloadItemId string) *BasicArrayList {
        var returnVector *BasicArrayList = NewBasicArrayListD(); _ = returnVector
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(this.BasicItemData.ID, id)
        keysAndValues.Put(DownloadItemData_ID, downloadItemId)
        var vector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var hashMap *HashMap = any(vector.Get(index)).(*HashMap); _ = hashMap
            returnVector.Add(NewDownloadableItem(hashMap))
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("Found: " + fmt.Sprint(size), this, "getForItem")
        }
        return returnVector
}


func (this *DownloadItemsEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(DownloadItemData_ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.BasicItemData.ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().ENABLE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_SPECIAL_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_VERSION).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_CHANGES).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(DownloadItemData_SYSTEM).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_PLATFORM).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_LICENSE_FILE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(DownloadItemData_FILE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DownloadItemData_SIZE).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(DownloadItemData_VALID_TIME).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(DownloadItemData_RETRIES).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(DownloadItemData_ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *DownloadItemsEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *DownloadItemsEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

