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

type InventoryEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    TableName string
}

// Implements InventoryEntityInterface


func (this *InventoryEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    this.TableName = "basicinventory"
}


func NewInventoryEntity() *InventoryEntity {
    this := &InventoryEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewInventoryDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *InventoryEntity) Insert(values *BasicArrayList) {
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


func (this *InventoryEntity) Delete(value string) {
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


// throws Exception
func (this *InventoryEntity) GetItems(storeFrontInterface StoreFrontInterface) *BasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("Getting Items For: " + storeFrontInterface.GetName(), this, "getItems")
        }
        var itemVector *BasicArrayList = NewBasicArrayListD(); _ = itemVector
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        var itemHashMapVector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = itemHashMapVector
        var size int = itemHashMapVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var itemHashMap *HashMap = any(itemHashMapVector.Get(i)).(*HashMap); _ = itemHashMap
            if itemHashMap != nil {
                var category string = any(itemHashMap.Get(this.BasicItemData.CATEGORY)).(string); _ = category
                if !StringValidationUtil_GetInstance().IsEmpty(category) && strings.HasPrefix(category, storeFrontInterface.GetCategoryPath()) {
                    itemVector.Add(NewBasicItem(itemHashMap))
                }
            }
        }
        return itemVector
}


// throws MoneyException
func (this *InventoryEntity) GetItem(id string) ItemInterface {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(this.BasicItemData.ID, id)
        var itemHashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = itemHashMap
        if itemHashMap != nil {
            return NewBasicItem(itemHashMap)
        } else {
            return nil
        }
}


func (this *InventoryEntity) GetWeight(id string) string {
        return this.AbSqlBean.GetField(this.BasicItemData.ID, id, this.BasicItemData.WEIGHT)
}


func (this *InventoryEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(this.BasicItemData.ID)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.NUMBER)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.INBASKETS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.WEIGHT)
        stringBuffer.Append(" VARCHAR(20) NOT NULL,")
        stringBuffer.Append(EntryData_GetInstance().ENABLE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.NEWORUSED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.SUMMARY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DISTRIBUTOR)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.IDUSEDBYDISTRIBUTOR)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRODUCEDBY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRODUCTIONDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.STARTPRODUCTIONDATE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DESCRIPTION)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.KEYWORDS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.CATEGORY)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.TYPE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.SMALLIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.MEDIUMIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.LARGEIMAGE)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().LASTMODIFIED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().TIMECREATED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PRICE)
        stringBuffer.Append(" VARCHAR(20) NOT NULL,")
        stringBuffer.Append(this.BasicItemData.COMMENT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.CUSTOMS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.DOWNLOADS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.GROUPS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.OPTIONS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.PERMISSIONS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.BasicItemData.SPECIALS)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(this.BasicItemData.ID)
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *InventoryEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *InventoryEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(this.BasicItemData.ID, any(updatedValues.Get(this.BasicItemData.ID)).(string), updatedValues)
}


func (this *InventoryEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

