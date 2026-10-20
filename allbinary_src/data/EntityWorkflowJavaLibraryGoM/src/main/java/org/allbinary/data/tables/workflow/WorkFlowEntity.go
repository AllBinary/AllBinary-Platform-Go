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

type WorkFlowEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    TableName string
    METHOD_GET string
    METHOD_UPDATE string
}

// Implements WorkFlowEntityInterface


func (this *WorkFlowEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    this.TableName = "workflows"
    this.METHOD_GET = "get()"
    this.METHOD_UPDATE = "update"
}


func NewWorkFlowEntity() *WorkFlowEntity {
    this := &WorkFlowEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
    return this
}


func (this *WorkFlowEntity) Insert(values *BasicArrayList) {
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


func (this *WorkFlowEntity) Delete(name string, storeName string) {
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
            var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
            keysAndValues.Put(WorkFlowData_GetInstance().NAME, name)
            keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
            this.AbSqlBean.DeleteWhere(keysAndValues)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception, LicensingException
func (this *WorkFlowEntity) Get(name string, storeName string) WorkFlowInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 WorkFlowInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.METHOD_GET, e)
            }
            panic(e)
                }
            }()
            var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
            keysAndValues.Put(WorkFlowData_GetInstance().NAME, name)
            keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
            var hashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = hashMap
            tryControl1, tryResult1 = 1, any(DbWorkFlowFactory_GetInstance().GetInstance(this.AbeClientInformation, hashMap)).(WorkFlowInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// Java overloaded method: Go needs a different name
func (this *WorkFlowEntity) Get(storeName string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.METHOD_GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var workFlowsVector *BasicArrayList = NewBasicArrayListD(); _ = workFlowsVector
            var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
            keysAndValues.Put(StoreFrontData_GetInstance().NAME, storeName)
            var hashMapVector *BasicArrayList = this.AbSqlBean.GetRows(keysAndValues); _ = hashMapVector
            var size int = hashMapVector.Size(); _ = size
            for i := int(0); i < size; i++ {
                var workFlowHashMap *HashMap = any(hashMapVector.Get(i)).(*HashMap); _ = workFlowHashMap
                if workFlowHashMap != nil {
                    workFlowsVector.Add(DbWorkFlowFactory_GetInstance().GetInstance(this.AbeClientInformation, workFlowHashMap))
                }
            }
            tryControl1, tryResult1 = 1, workFlowsVector; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *WorkFlowEntity) Update(updatedValues *HashMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.METHOD_UPDATE, e)
            }
                }
            }()
            var wherekeysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = wherekeysAndValues
            wherekeysAndValues.Put(WorkFlowData_GetInstance().NAME, any(updatedValues.Get(WorkFlowData_GetInstance().NAME)).(string))
            wherekeysAndValues.Put(StoreFrontData_GetInstance().NAME, any(updatedValues.Get(StoreFrontData_GetInstance().NAME)).(string))
            this.AbSqlBean.UpdateWhere(wherekeysAndValues, updatedValues)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WorkFlowEntity) CreateTableStatement() string {
        var workFlowData *WorkFlowData = WorkFlowData_GetInstance(); _ = workFlowData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.TableName).Append(this.SqlStrings.START).Append(workFlowData.NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(StoreFrontData_GetInstance().NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(DynamicObjectData_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(workFlowData.DATA).Append(this.SqlTypeStrings.BLOB_NOT_NULL).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(EntryData_GetInstance().LASTMODIFIED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(workFlowData.NAME).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *WorkFlowEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *WorkFlowEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

