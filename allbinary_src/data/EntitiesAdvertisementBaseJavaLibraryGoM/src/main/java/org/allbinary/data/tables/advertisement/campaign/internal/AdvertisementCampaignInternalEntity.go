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

type AdvertisementCampaignInternalEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements AdvertisementCampaignInternalEntityInterface


func (this *AdvertisementCampaignInternalEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "adCampaignInternal"
}


func NewAdvertisementCampaignInternalEntity() *AdvertisementCampaignInternalEntity {
    this := &AdvertisementCampaignInternalEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *AdvertisementCampaignInternalEntity) Delete(value string) {
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
            this.AbSqlBean.DeleteWhere(StoreFrontData_GetInstance().NAME, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AdvertisementCampaignInternalEntity) GetCampaignsInStore(storeName string) AdvertisementCampaignsInterface {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
        var hashMapVector IBasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = hashMapVector
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var size int = hashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var hashMap *HashMap = any(hashMapVector.Get(index)).(*HashMap); _ = hashMap
            if hashMap != nil {
                vector.Add(any(NewAdvertisementCampaign(hashMap)).(AdvertisementCampaignInterface))
            }
        }
        return any(NewAdvertisementCampaigns(vector)).(AdvertisementCampaignsInterface)
}


func (this *AdvertisementCampaignInternalEntity) Get(storeName string, name string) AdvertisementCampaignInterface {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
        keysAndValues.Put(AdvertisementCampaignData_GetInstance().NAME, name)
        var hashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = hashMap
        if hashMap != nil {
            return any(NewAdvertisementCampaign(hashMap)).(AdvertisementCampaignInterface)
        } else {
            return nil
        }
}


func (this *AdvertisementCampaignInternalEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(EntryData_GetInstance().ID, any(updatedValues.Get(EntryData_GetInstance().ID)).(string), updatedValues)
}


func (this *AdvertisementCampaignInternalEntity) CreateTableStatement() string {
        var entryData *EntryData = EntryData_GetInstance(); _ = entryData
        var advertisementCampaignData *AdvertisementCampaignData = AdvertisementCampaignData_GetInstance(); _ = advertisementCampaignData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(entryData.ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(StoreFrontData_GetInstance().NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(advertisementCampaignData.NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(advertisementCampaignData.DESCRIPTION).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(advertisementCampaignData.CONFIG).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(TimeData_GetInstance().START).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(TimeData_GetInstance().END).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(entryData.TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(entryData.LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(entryData.ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *AdvertisementCampaignInternalEntity) CreateTable() string {
        var returnStr string = this.AbSqlBean.CreateTable(this.CreateTableStatement()); _ = returnStr
        return returnStr
}


func (this *AdvertisementCampaignInternalEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

