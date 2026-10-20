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

type BasicGroupItemsRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    Request *HttpServletRequest
    Id string
    ItemOne string
    ItemTwo string
    ItemThree string
    ItemFour string
    ItemFive string
    ItemSix string
    ItemSeven string
    ItemEight string
    ItemNine string
    ItemTen string
    TimeEntered string
    LastModified string
}


func (this *BasicGroupItemsRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewBasicGroupItemsRequestHelper(hashMap *HashMap, pageContext *PageContext) *BasicGroupItemsRequestHelper {
    this := &BasicGroupItemsRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *BasicGroupItemsRequestHelper) GetFormData() {
        this.Id = this.Request.GetParameter(this.BasicItemData.ID)
        this.ItemOne = this.Request.GetParameter(BasicGroupItemData_ITEM_ONE)
        this.ItemTwo = this.Request.GetParameter(BasicGroupItemData_ITEM_TWO)
        this.ItemThree = this.Request.GetParameter(BasicGroupItemData_ITEM_THREE)
        this.ItemFour = this.Request.GetParameter(BasicGroupItemData_ITEM_FOUR)
        this.ItemFive = this.Request.GetParameter(BasicGroupItemData_ITEM_FIVE)
        this.ItemSix = this.Request.GetParameter(BasicGroupItemData_ITEM_SIX)
        this.ItemSeven = this.Request.GetParameter(BasicGroupItemData_ITEM_SEVEN)
        this.ItemEight = this.Request.GetParameter(BasicGroupItemData_ITEM_EIGHT)
        this.ItemNine = this.Request.GetParameter(BasicGroupItemData_ITEM_NINE)
        this.ItemTen = this.Request.GetParameter(BasicGroupItemData_ITEM_TEN)
        this.TimeEntered = this.Request.GetParameter(EntryData_GetInstance().TIMECREATED)
        this.LastModified = this.Request.GetParameter(EntryData_GetInstance().LASTMODIFIED)
}


func (this *BasicGroupItemsRequestHelper) GetHashMap() *HashMap {
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(this.BasicItemData.ID, this.Id)
        values.Put(BasicGroupItemData_ITEM_ONE, this.ItemOne)
        values.Put(BasicGroupItemData_ITEM_TWO, this.ItemTwo)
        values.Put(BasicGroupItemData_ITEM_THREE, this.ItemThree)
        values.Put(BasicGroupItemData_ITEM_FOUR, this.ItemFour)
        values.Put(BasicGroupItemData_ITEM_FIVE, this.ItemFive)
        values.Put(BasicGroupItemData_ITEM_SIX, this.ItemSix)
        values.Put(BasicGroupItemData_ITEM_SEVEN, this.ItemSeven)
        values.Put(BasicGroupItemData_ITEM_EIGHT, this.ItemEight)
        values.Put(BasicGroupItemData_ITEM_NINE, this.ItemNine)
        values.Put(BasicGroupItemData_ITEM_TEN, this.ItemTen)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *BasicGroupItemsRequestHelper) Insert() string {
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
            values.Add(this.ItemOne)
            values.Add(this.ItemTwo)
            values.Add(this.ItemThree)
            values.Add(this.ItemFour)
            values.Add(this.ItemFive)
            values.Add(this.ItemSix)
            values.Add(this.ItemSeven)
            values.Add(this.ItemEight)
            values.Add(this.ItemNine)
            values.Add(this.ItemTen)
            values.Add(time)
            values.Add(time)
            BasicGroupItemsEntityFactory_GetInstance().GetBasicGroupItemsEntityInstance().Insert(values)
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


func (this *BasicGroupItemsRequestHelper) Delete() string {
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
            BasicGroupItemsEntityFactory_GetInstance().GetBasicGroupItemsEntityInstance().Delete(this.Id)
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


func (this *BasicGroupItemsRequestHelper) Update() string {
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
            BasicGroupItemsEntityFactory_GetInstance().GetBasicGroupItemsEntityInstance().Update(values)
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

