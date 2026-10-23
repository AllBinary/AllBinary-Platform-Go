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

type TransformInfoEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    TableName string
    AbeClientInformation AbeClientInformationInterface
    TransformInfoFactoryInterface TransformInfoFactoryInterface
    TransformInfoObjectConfigAndManipulatorFactoryInterface ITransformInfoObjectConfigAndManipulatorFactoryBase
    TransformInfoObjectConfigGeneratorFactoryInterface TransformInfoObjectConfigGeneratorFactoryInterface
}

// Implements TransformInfoEntityInterface


func (this *TransformInfoEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "transforminfo"
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
}


func NewTransformInfoEntity(transformInfoObjectConfigGeneratorFactoryInterface TransformInfoObjectConfigGeneratorFactoryInterface, transformInfoObjectConfigAndManipulatorFactoryInterface ITransformInfoObjectConfigAndManipulatorFactoryBase, transformInfoFactoryInterface TransformInfoFactoryInterface) *TransformInfoEntity {
    this := &TransformInfoEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.SetTableName(this.TableName)
        this.TransformInfoObjectConfigGeneratorFactoryInterface = transformInfoObjectConfigGeneratorFactoryInterface
        this.TransformInfoObjectConfigAndManipulatorFactoryInterface = transformInfoObjectConfigAndManipulatorFactoryInterface
        this.TransformInfoFactoryInterface = transformInfoFactoryInterface
    return this
}


func (this *TransformInfoEntity) Insert(values IBasicArrayList) {
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


func (this *TransformInfoEntity) Delete(value string) {
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
            this.AbSqlBean.DeleteWhere(TransformInfoData_GetInstance().NAME, value)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, this.CommonStrings.Delete)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *TransformInfoEntity) Get(name string, propertiesHashMap *HashMap, pageContext *PageContext) TransformInfoInterface {
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(transformInfoData.NAME, name)
        var hashMap *HashMap = this.AbSqlBean.GetRow(keysAndValues); _ = hashMap
        if hashMap != nil {
            var anyType any = hashMap.Get(transformInfoData.OBJECTCONFIG); _ = anyType
            if anyType != nil {
                var string_ string = any(anyType).(string); _ = string_
                hashMap.Put(transformInfoData.OBJECTCONFIG, Encoder_Decode(string_))
            }
            var objectData any = hashMap.Get(transformInfoData.DATA); _ = objectData
            if objectData != nil {
                var string_ string = any(objectData).(string); _ = string_
                hashMap.Put(transformInfoData.DATA, Encoder_Decode(string_))
            }
            return this.TransformInfoFactoryInterface.GetInstance(hashMap, propertiesHashMap, pageContext)
        } else {
            return nil
        }
}


// throws Exception
func (this *TransformInfoEntity) GetObjectConfigs(storeName string) IBasicArrayList {
        var objectConfigVector IBasicArrayList = NewBasicArrayListD(); _ = objectConfigVector
        var objectConfigColumnVector IBasicArrayList = this.GetColumnWhere(TransformInfoData_GetInstance().OBJECTCONFIG, StoreFrontData_GetInstance().NAME, storeName); _ = objectConfigColumnVector
        var size int = objectConfigColumnVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var objectConfigString string = any(objectConfigColumnVector.Get(i)).(string); _ = objectConfigString
            objectConfigString = Encoder_Decode(objectConfigString)
            objectConfigVector.Add(this.TransformInfoObjectConfigAndManipulatorFactoryInterface.GetInstance(this.AbeClientInformation, any(this).(TransformInfoInterface), DomDocumentHelper_Create(objectConfigString)))
        }
        return objectConfigVector
}


// throws Exception
func (this *TransformInfoEntity) GetNames(storeName string) IBasicArrayList {
        var viewNameVector IBasicArrayList = NewBasicArrayListD(); _ = viewNameVector
        var columnVector IBasicArrayList = this.GetColumnWhere(TransformInfoData_GetInstance().NAME, StoreFrontData_GetInstance().NAME, storeName); _ = columnVector
        var size int = columnVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var viewNameString string = any(columnVector.Get(i)).(string); _ = viewNameString
            viewNameVector.Add(viewNameString)
        }
        return viewNameVector
}


func (this *TransformInfoEntity) CreateTableStatement() string {
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE)
        stringBuffer.Append(this.TableName)
        stringBuffer.Append(this.SqlStrings.START)
        stringBuffer.Append(transformInfoData.NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(StoreFrontData_GetInstance().NAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(transformInfoData.OBJECTFILENAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(transformInfoData.OBJECT)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(transformInfoData.OBJECTCONFIGFILENAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(transformInfoData.OBJECTCONFIG)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(transformInfoData.TEMPLATEFILENAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(transformInfoData.TEMPLATE)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(transformInfoData.DATAFILENAME)
        stringBuffer.Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL)
        stringBuffer.Append(transformInfoData.DATA)
        stringBuffer.Append(this.SqlTypeStrings.BLOB_NOT_NULL)
        stringBuffer.Append(EntryData_GetInstance().LASTMODIFIED)
        stringBuffer.Append(EntryData_GetInstance().TIMECREATED)
        stringBuffer.Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL)
        stringBuffer.Append(this.SqlStrings.PRIMARY_KEY)
        stringBuffer.Append(transformInfoData.NAME)
        stringBuffer.Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *TransformInfoEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *TransformInfoEntity) Update(updatedValues *HashMap) {
        this.AbSqlBean.UpdateWhere(TransformInfoData_GetInstance().NAME, any(updatedValues.Get(TransformInfoData_GetInstance().NAME)).(string), updatedValues)
}


func (this *TransformInfoEntity) DropTable() string {
        return this.AbSqlBean.DropTable()
}

