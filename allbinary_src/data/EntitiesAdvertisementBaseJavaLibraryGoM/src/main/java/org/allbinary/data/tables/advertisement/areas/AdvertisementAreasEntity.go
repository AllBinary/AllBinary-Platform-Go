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

type AdvertisementAreasEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements AdvertisementAreasEntityInterface


func (this *AdvertisementAreasEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "advertisements"
}


func NewAdvertisementAreasEntity() *AdvertisementAreasEntity {
    this := &AdvertisementAreasEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *AdvertisementAreasEntity) Delete(value string) {
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
            this.AbSqlBean.DeleteWhere(EntryData_GetInstance().ID, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AdvertisementAreasEntity) Get_1(storeName string) IBasicArrayList {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
        var hashMapVector IBasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = hashMapVector
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var size int = hashMapVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var hashMap *HashMap = any(hashMapVector.Get(index)).(*HashMap); _ = hashMap
            if hashMap != nil {
                vector.Add(any(NewAdvertisementArea(hashMap)).(AdvertisementAreaInterface))
            }
        }
        return vector
}

// Java overloads of AdvertisementAreasEntity.get
func (this *AdvertisementAreasEntity) Get(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.Get_1(Java_Arg[string](args[0]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.Get_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of AdvertisementAreasEntity.get for the arguments")
}


// throws Exception
func (this *AdvertisementAreasEntity) Get_2(storeName string, advertisementAreaName string) AdvertisementAreaInterface {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
        keysAndValues.Put(AdvertisementAreaData_GetInstance().NAME, advertisementAreaName)
        var hashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = hashMap
        if hashMap != nil {
            return any(NewAdvertisementArea(hashMap)).(AdvertisementAreaInterface)
        } else {
            return nil
        }
}


func (this *AdvertisementAreasEntity) CreateTableStatement() string {
        var advertisementAreaData *AdvertisementAreaData = AdvertisementAreaData_GetInstance(); _ = advertisementAreaData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.GetTableName())
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(advertisementAreaData.NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(StoreFrontData_GetInstance().NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(advertisementAreaData.DESCRIPTION).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(advertisementAreaData.CONSTRAINTS).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(AdvertisementCampaignData_GetInstance().NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(advertisementAreaData.NAME).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *AdvertisementAreasEntity) CreateTable() string {
        var returnStr string = this.AbSqlBean.CreateTable(this.CreateTableStatement()); _ = returnStr
        return returnStr
}


func (this *AdvertisementAreasEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(AdvertisementAreaData_GetInstance().NAME, any(updatedValues.Get(AdvertisementAreaData_GetInstance().NAME)).(string), updatedValues)
}


func (this *AdvertisementAreasEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

