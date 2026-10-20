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

type PermissionItemsRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    Request *HttpServletRequest
    Id string
    Enabled string
    Number string
    Who string
    What string
    Type string
    RemoteAddr string
    StartTime string
    EndTime string
    Price string
    TimeEntered string
    LastModified string
}


func (this *PermissionItemsRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewPermissionItemsRequestHelper(hashMap *HashMap, pageContext *PageContext) *PermissionItemsRequestHelper {
    this := &PermissionItemsRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *PermissionItemsRequestHelper) GetFormData() {
        this.Id = this.Request.GetParameter(this.BasicItemData.ID)
        this.Number = this.Request.GetParameter(this.BasicItemData.NUMBER)
        this.Enabled = this.Request.GetParameter(EntryData_GetInstance().ENABLE)
        this.Who = this.Request.GetParameter(PermissionItemData_WHO)
        this.What = this.Request.GetParameter(PermissionItemData_WHAT)
        this.Type = this.Request.GetParameter(PermissionItemData_TYPE)
        this.RemoteAddr = this.Request.GetParameter(PermissionItemData_REMOTE_ADDR)
        this.StartTime = this.Request.GetParameter(PermissionItemData_START_TIME)
        this.EndTime = this.Request.GetParameter(PermissionItemData_END_TIME)
        this.Price = this.Request.GetParameter(this.BasicItemData.PRICE)
        this.TimeEntered = this.Request.GetParameter(EntryData_GetInstance().TIMECREATED)
        this.LastModified = this.Request.GetParameter(EntryData_GetInstance().LASTMODIFIED)
}


func (this *PermissionItemsRequestHelper) GetHashMap() *HashMap {
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(this.BasicItemData.ID, this.Id)
        values.Put(this.BasicItemData.NUMBER, this.Number)
        values.Put(EntryData_GetInstance().ENABLE, this.Enabled)
        values.Put(PermissionItemData_WHO, this.Who)
        values.Put(PermissionItemData_WHAT, this.What)
        values.Put(PermissionItemData_TYPE, this.Type)
        values.Put(PermissionItemData_REMOTE_ADDR, this.RemoteAddr)
        values.Put(PermissionItemData_START_TIME, this.StartTime)
        values.Put(PermissionItemData_END_TIME, this.EndTime)
        values.Put(this.BasicItemData.PRICE, this.Price)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *PermissionItemsRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to insert " + this.Id + " into items table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "inserts()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var calendar *Calendar = Calendar_GetInstance(); _ = calendar
            var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
            var values IBasicArrayList = NewBasicArrayListD(); _ = values
            values.Add(this.Id)
            values.Add(this.Number)
            values.Add(this.Enabled)
            values.Add(this.Who)
            values.Add(this.What)
            values.Add(this.Type)
            values.Add(this.RemoteAddr)
            values.Add(this.StartTime)
            values.Add(this.EndTime)
            values.Add(this.Price)
            values.Add(time)
            values.Add(time)
            PermissionItemsEntityFactory_GetPermissionItemsEntityInstance().Insert(values)
            var success string = "Successfully inserted " + this.Id + " into items table"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "insert()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PermissionItemsRequestHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to delete"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            PermissionItemsEntityFactory_GetPermissionItemsEntityInstance().Delete(this.Id)
            var success string = "Successfully deleted"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "delete()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PermissionItemsRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update: " + this.Id; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Update Pricing Successful"; _ = success
            var values *HashMap = this.GetHashMap(); _ = values
            PermissionItemsEntityFactory_GetPermissionItemsEntityInstance().Update(values)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.Id).Append(CommonSeps_GetInstance().SPACE).Append(success)), this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

