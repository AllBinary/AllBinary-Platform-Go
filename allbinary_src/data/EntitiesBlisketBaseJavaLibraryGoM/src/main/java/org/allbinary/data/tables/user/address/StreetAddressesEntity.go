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

type StreetAddressesEntity struct {
    AbSqlBean
    LogUtil *LogUtil
    UserName string
}

// Implements StreetAddressesEntityInterface


func (this *StreetAddressesEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewStreetAddressesEntity() *StreetAddressesEntity {
    this := &StreetAddressesEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewStreetAddressesEntity(userName string) *StreetAddressesEntity {
    this := &StreetAddressesEntity{}
        this.AbSqlBean = *NewAbSqlBean(NewUserDbInitInfo())
        this.initFields()
        this.UserName = userName
    return this
}


func (this *StreetAddressesEntity) Remove(index int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "remove", e)
            }
                }
            }()
            var whereHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereHashMap
            whereHashMap.Put(UserData_USERNAME, this.UserName)
            whereHashMap.Put(StreetAddressData_ID, any(fmt.Sprint(index)).(string))
            this.AbSqlBean.DeleteWhere(whereHashMap)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "remove")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StreetAddressesEntity) Add(address *StreetAddress) {
        this.Add(address, TableDataFactory_GetInstance().INTEGER_MAX_VALUE_STRING)
}


// Java overloaded method: Go needs a different name
func (this *StreetAddressesEntity) Add(address *StreetAddress, index string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "add", e)
            }
                }
            }()
            var values *BasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(index)
            values.Add(this.UserName)
            values.Add(StringUtil_GetInstance().EMPTY_STRING)
            values.Add(address.GetName())
            values.Add(address.GetStreet())
            values.Add(address.GetCity())
            values.Add(address.GetState())
            values.Add(address.GetCode())
            values.Add(address.GetCountry())
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            values.Add(time)
            this.AbSqlBean.Insert(values)
            this.SetDefault(this.GetLastId())
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "add")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StreetAddressesEntity) Update(address *StreetAddress) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "add", e)
            }
                }
            }()
            var addressHashMap *HashMap = address.ToHashMap(); _ = addressHashMap
            var whereKeyValuePairs *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereKeyValuePairs
            whereKeyValuePairs.Put(StreetAddressData_ID, address.GetId())
            whereKeyValuePairs.Put(UserData_USERNAME, this.UserName)
            this.AbSqlBean.UpdateWhere(whereKeyValuePairs, addressHashMap)
            this.SetDefault(address.GetId())
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "update")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StreetAddressesEntity) GetLastId() string {
        return this.AbSqlBean.GetLargestIntegerInColumnWhere(StreetAddressData_ID, UserData_USERNAME, this.UserName)
}


func (this *StreetAddressesEntity) Get() *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var streetAddressVector *BasicArrayList = NewBasicArrayListD(); _ = streetAddressVector
            var keyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyAndValue
            keyAndValue.Put(UserData_USERNAME, this.UserName)
            var addressList *BasicArrayList = this.AbSqlBean.GetRows(keyAndValue); _ = addressList
            var size int = addressList.Size(); _ = size
            for index := int(0); index < size; index++ {
                var addressHashMap *HashMap = any(addressList.Get(index)).(*HashMap); _ = addressHashMap
                var streetAddress *StreetAddress = NewStreetAddress(addressHashMap); _ = streetAddress
                if streetAddress != nil {
                    streetAddressVector.Add(streetAddress)
                }
            }
            tryControl1, tryResult1 = 1, streetAddressVector; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// Java overloaded method: Go needs a different name
func (this *StreetAddressesEntity) Get(index int) *StreetAddress {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *StreetAddress) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var keyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keyAndValue
            keyAndValue.Put(UserData_USERNAME, this.UserName)
            keyAndValue.Put(StreetAddressData_ID, fmt.Sprint(index))
            var addressHashMap *HashMap = this.AbSqlBean.GetRow(keyAndValue); _ = addressHashMap
            if addressHashMap != nil {
                tryControl1, tryResult1 = 1, NewStreetAddress(addressHashMap); return
            } else {
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StreetAddressesEntity) GetDefault() *StreetAddress {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *StreetAddress) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getDefault", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var addressHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = addressHashMap
            var updateKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateKeyAndValue
            updateKeyAndValue.Put(StreetAddressData_DEFAULT, StreetAddressData_DEFAULT)
            updateKeyAndValue.Put(UserData_USERNAME, this.UserName)
            addressHashMap = this.AbSqlBean.GetRow(updateKeyAndValue)
            if addressHashMap != nil {
                var streetAddress *StreetAddress = NewStreetAddress(addressHashMap); _ = streetAddress
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "getDefault")
                }
                tryControl1, tryResult1 = 1, streetAddress; return
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF("No Default Address Found", this, "getDefault")
                }
                tryControl1, tryResult1 = 1, nil; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StreetAddressesEntity) SetDefault(value string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "setShippingAddress", e)
            }
                }
            }()
            var updateKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = updateKeyAndValue
            var whereKeyAndValue *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = whereKeyAndValue
            whereKeyAndValue.Put(UserData_USERNAME, this.UserName)
            var streetAddress *StreetAddress = this.GetDefault(); _ = streetAddress
            if streetAddress != nil {
                updateKeyAndValue.Put(StreetAddressData_DEFAULT, StringUtil_GetInstance().EMPTY_STRING)
                whereKeyAndValue.Put(StreetAddressData_ID, streetAddress.GetId())
                this.AbSqlBean.UpdateWhere(whereKeyAndValue, updateKeyAndValue)
            }
            whereKeyAndValue.Put(StreetAddressData_ID, value)
            updateKeyAndValue.Put(StreetAddressData_DEFAULT, StreetAddressData_DEFAULT)
            this.AbSqlBean.UpdateWhere(whereKeyAndValue, updateKeyAndValue)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF(this.CommonStrings.SUCCESS, this, "setDefault")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StreetAddressesEntity) GetUserName() string {
        return this.UserName
}


func (this *StreetAddressesEntity) CreateTableStatement() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.SqlStrings.CREATE_TABLE).Append(this.GetTableName()).Append(this.SqlStrings.START).Append(StreetAddressData_ID).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_AUTO_INCREMENT_NOT_NULL).Append(UserData_USERNAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(StreetAddressData_DEFAULT).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(StreetAddressData_NAME).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN_NOT_NULL).Append(StreetAddressData_STREET).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(StreetAddressData_CITY).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(StreetAddressData_STATE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(StreetAddressData_CODE).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(StreetAddressData_COUNTRY).Append(this.SqlTypeStrings.MAX_CHAR_COLUMN).Append(EntryData_GetInstance().TIMECREATED).Append(this.SqlTypeStrings.MAX_BIG_INT_UNSIGNED_NOT_NULL).Append(this.SqlStrings.PRIMARY_KEY).Append(StreetAddressData_ID).Append(this.SqlStrings.END)
        return fmt.Sprint(stringBuffer)
}


func (this *StreetAddressesEntity) CreateTable() string {
        return this.AbSqlBean.CreateTable(this.CreateTableStatement())
}


func (this *StreetAddressesEntity) Drop() string {
        return this.AbSqlBean.DropTable()
}

