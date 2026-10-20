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

type StoreFrontsEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
}

// Implements StoreFrontsEntityInterface


func (this *StoreFrontsEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "storefronts"
}


func NewStoreFrontsEntity() *StoreFrontsEntity {
    this := &StoreFrontsEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *StoreFrontsEntity) Insert(values IBasicArrayList) {
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


func (this *StoreFrontsEntity) Delete(value string) {
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


// throws Exception
func (this *StoreFrontsEntity) GetStoreFrontInterface(name string) *StoreFront {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, name)
        var storeHashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = storeHashMap
        if storeHashMap != nil {
            return NewStoreFront(storeHashMap)
        } else {
            return NewStoreFront()
        }
}


func (this *StoreFrontsEntity) GetStoreFrontNames() IBasicArrayList {
        var storeFrontNames IBasicArrayList = this.AbSqlBean.GetColumn(StoreFrontData_GetInstance().NAME); _ = storeFrontNames
        return storeFrontNames
}


func (this *StoreFrontsEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(StoreFrontData_GetInstance().NAME, any(updatedValues.Get(StoreFrontData_GetInstance().NAME)).(string), updatedValues)
}


func (this *StoreFrontsEntity) CreateTableStatement() string {
        var storeFrontData *StoreFrontData = StoreFrontData_GetInstance(); _ = storeFrontData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(storeFrontData.NAME)
        stringBuffer.Append(this.SqlTypeStrings.SIXTY_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.HOMEHOSTNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.HOMEHOSTNAMEPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.HOSTNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.HOSTNAMEPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTHOMEHOSTNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTHOMEHOSTNAMEPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTHOSTNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTHOSTNAMEPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.IMAGEPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.STATICPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.CATEGORYPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.INVENTORYCONTROL)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.CONFIGURATION)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(storeFrontData.SUBSTORES)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TAGLOCATION)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.PACKAGELOCATION)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.FTP)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.FTPPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.FTPUSERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.FTPPASSWORD)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTFTP)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTFTPPATH)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTFTPUSERNAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(storeFrontData.TESTFTPPASSWORD)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().TIMECREATED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().LASTMODIFIED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(storeFrontData.NAME)
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *StoreFrontsEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *StoreFrontsEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

