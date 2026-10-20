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

type BasicOptionItemsRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    Request *HttpServletRequest
    Id string
    OptionOneTitle string
    DefaultOptionItem string
    DefaultOptionValue string
    OptionItem IBasicArrayList
    OptionValue IBasicArrayList
    TimeEntered string
    LastModified string
}


func (this *BasicOptionItemsRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
    
    
    
}


func NewBasicOptionItemsRequestHelper(hashMap *HashMap, pageContext *PageContext) *BasicOptionItemsRequestHelper {
    this := &BasicOptionItemsRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *BasicOptionItemsRequestHelper) GetFormData() {
        this.OptionItem = NewBasicArrayListD()
        this.OptionValue = NewBasicArrayListD()
        this.Id = this.Request.GetParameter(this.BasicItemData.ID)
        this.OptionOneTitle = this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_TITLE)
        this.DefaultOptionItem = this.Request.GetParameter(BasicOptionItemData_DEFAULT_OPTION_ITEM)
        this.DefaultOptionValue = this.Request.GetParameter(BasicOptionItemData_DEFAULT_OPTION_VALUE)
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_ONE_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_ONE_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_TWO_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_TWO_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_THREE_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_THREE_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_FOUR_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_FOUR_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_FIVE_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_FIVE_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_SIX_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_SIX_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_SEVEN_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_SEVEN_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_EIGHT_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_EIGHT_VALUE))
        this.OptionItem.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_NINE_ITEM))
        this.OptionValue.Add(this.Request.GetParameter(BasicOptionItemData_OPTION_ONE_NINE_VALUE))
        this.TimeEntered = this.Request.GetParameter(EntryData_GetInstance().TIMECREATED)
        this.LastModified = this.Request.GetParameter(EntryData_GetInstance().LASTMODIFIED)
}


func (this *BasicOptionItemsRequestHelper) GetHashMap() *HashMap {
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(this.BasicItemData.ID, this.Id)
        values.Put(BasicOptionItemData_OPTION_ONE_ONE_ITEM, this.OptionItem.Get(0))
        values.Put(BasicOptionItemData_OPTION_ONE_ONE_VALUE, this.OptionValue.Get(0))
        values.Put(BasicOptionItemData_OPTION_ONE_TWO_ITEM, this.OptionItem.Get(1))
        values.Put(BasicOptionItemData_OPTION_ONE_TWO_VALUE, this.OptionValue.Get(1))
        values.Put(BasicOptionItemData_OPTION_ONE_THREE_ITEM, this.OptionItem.Get(2))
        values.Put(BasicOptionItemData_OPTION_ONE_THREE_VALUE, this.OptionValue.Get(2))
        values.Put(BasicOptionItemData_OPTION_ONE_FOUR_ITEM, this.OptionItem.Get(3))
        values.Put(BasicOptionItemData_OPTION_ONE_FOUR_VALUE, this.OptionValue.Get(3))
        values.Put(BasicOptionItemData_OPTION_ONE_FIVE_ITEM, this.OptionItem.Get(4))
        values.Put(BasicOptionItemData_OPTION_ONE_FIVE_VALUE, this.OptionValue.Get(4))
        values.Put(BasicOptionItemData_OPTION_ONE_SIX_ITEM, this.OptionItem.Get(5))
        values.Put(BasicOptionItemData_OPTION_ONE_SIX_VALUE, this.OptionValue.Get(5))
        values.Put(BasicOptionItemData_OPTION_ONE_SEVEN_ITEM, this.OptionItem.Get(6))
        values.Put(BasicOptionItemData_OPTION_ONE_SEVEN_VALUE, this.OptionValue.Get(6))
        values.Put(BasicOptionItemData_OPTION_ONE_EIGHT_ITEM, this.OptionItem.Get(7))
        values.Put(BasicOptionItemData_OPTION_ONE_EIGHT_VALUE, this.OptionValue.Get(7))
        values.Put(BasicOptionItemData_OPTION_ONE_NINE_ITEM, this.OptionItem.Get(8))
        values.Put(BasicOptionItemData_OPTION_ONE_NINE_VALUE, this.OptionValue.Get(8))
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *BasicOptionItemsRequestHelper) Insert() string {
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
            var size int = this.OptionValue.Size(); _ = size
            for index := int(0); index < size; index++ {
                values.Add(this.OptionValue.Get(index))
                values.Add(this.OptionItem.Get(index))
            }
            values.Add(time)
            values.Add(time)
            BasicOptionItemsEntityFactory_GetInstance().GetBasicOptionItemsEntityInstance().Insert(values)
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


func (this *BasicOptionItemsRequestHelper) Delete() string {
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
            BasicOptionItemsEntityFactory_GetInstance().GetBasicOptionItemsEntityInstance().Delete(this.Id)
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


func (this *BasicOptionItemsRequestHelper) Update() string {
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
            BasicOptionItemsEntityFactory_GetInstance().GetBasicOptionItemsEntityInstance().Update(values)
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

