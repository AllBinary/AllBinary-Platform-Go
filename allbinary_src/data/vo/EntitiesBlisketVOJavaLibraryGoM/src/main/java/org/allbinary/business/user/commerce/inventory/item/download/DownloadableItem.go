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

type DownloadableItem struct {
    Id string
    BasicItemId string
    Enabled string
    System string
    Platform string
    SpecialName string
    Version string
    Changes string
    LicenseFile string
    File string
    Size int64
    ValidTime int64
    Retries string
}


func (this *DownloadableItem) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
}




// Java overloads of new DownloadableItem
func NewDownloadableItem(args ...any) *DownloadableItem {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewDownloadableItem_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewDownloadableItem_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new DownloadableItem for the arguments")
}
func NewDownloadableItem_1(id string) *DownloadableItem {
    this := &DownloadableItem{}
    this.initFields()
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        this.Id = NewProductIdGenerator().GetNext()
        this.BasicItemId = id
        this.Enabled = EntryData_GetInstance().YES
        this.System = "Java"
        this.Platform = "Android"
        this.SpecialName = EMPTY_STRING
        this.Version = EMPTY_STRING
        this.Changes = EMPTY_STRING
        this.LicenseFile = EMPTY_STRING
        this.File = EMPTY_STRING
        this.Size = fmt.Sprint(0)
        this.ValidTime = fmt.Sprint(0)
        this.Retries = EMPTY_STRING
    return this
}


func NewDownloadableItem_2(hashMap *HashMap) *DownloadableItem {
    this := &DownloadableItem{}
    this.initFields()
        this.Id = any(hashMap.Get(DownloadItemData_ID)).(string)
        this.BasicItemId = any(hashMap.Get(BasicItemData_GetInstance().ID)).(string)
        this.Enabled = any(hashMap.Get(EntryData_GetInstance().ENABLE)).(string)
        this.System = any(hashMap.Get(DownloadItemData_SYSTEM)).(string)
        this.Platform = any(hashMap.Get(DownloadItemData_PLATFORM)).(string)
        this.SpecialName = any(hashMap.Get(DownloadItemData_SPECIAL_NAME)).(string)
        this.Version = any(hashMap.Get(DownloadItemData_VERSION)).(string)
        this.Changes = any(hashMap.Get(DownloadItemData_CHANGES)).(string)
        this.LicenseFile = any(hashMap.Get(DownloadItemData_LICENSE_FILE)).(string)
        this.File = any(hashMap.Get(DownloadItemData_FILE)).(string)
        var value string = any(hashMap.Get(DownloadItemData_SIZE)).(string); _ = value
        if !StringValidationUtil_GetInstance().IsEmpty(value) {
            this.Size = fmt.Sprint(value)
        } else {
            this.Size = int64(0)
        }
        var validTime string = any(hashMap.Get(DownloadItemData_VALID_TIME)).(string); _ = validTime
        if !StringValidationUtil_GetInstance().IsEmpty(validTime) {
            this.ValidTime = fmt.Sprint(validTime)
        } else {
            var totalYears string = any(hashMap.Get(DownloadItemData_VALID_TIME_YEARS)).(string); _ = totalYears
            var totalMonths string = any(hashMap.Get(DownloadItemData_VALID_TIME_MONTHS)).(string); _ = totalMonths
            var totalDays string = any(hashMap.Get(DownloadItemData_VALID_TIME_DAYS)).(string); _ = totalDays
            var totalHours string = any(hashMap.Get(DownloadItemData_VALID_TIME_HOURS)).(string); _ = totalHours
            var totalMinutes string = any(hashMap.Get(DownloadItemData_VALID_TIME_MINUTES)).(string); _ = totalMinutes
            var totalSeconds string = any(hashMap.Get(DownloadItemData_VALID_TIME_SECONDS)).(string); _ = totalSeconds
            var totalTime int64 = TimeUtil_GetInstance().GetTotalTime(totalYears, totalMonths, totalDays, totalHours, totalMinutes, totalSeconds); _ = totalTime
            this.ValidTime = int64(totalTime)
        }
        this.Retries = any(hashMap.Get(DownloadItemData_RETRIES)).(string)
    return this
}


func (this *DownloadableItem) ToVector() IBasicArrayList {
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        var values IBasicArrayList = NewBasicArrayListD(); _ = values
        values.Add(this.GetId())
        values.Add(this.BasicItemId)
        values.Add(this.Enabled)
        values.Add(this.GetSpecialName())
        values.Add(this.GetVersion())
        values.Add(this.GetChanges())
        values.Add(this.GetSystem())
        values.Add(this.GetPlatform())
        values.Add(this.GetLicenseFile())
        values.Add(this.GetFile())
        values.Add(fmt.Sprint(this.GetSize()))
        values.Add(fmt.Sprint(this.GetValidTime()))
        values.Add(this.GetRetries())
        values.Add(time)
        values.Add(time)
        return values
}


func (this *DownloadableItem) ToHashMap() *HashMap {
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(DownloadItemData_ID, this.GetId())
        values.Put(BasicItemData_GetInstance().ID, this.BasicItemId)
        values.Put(EntryData_GetInstance().ENABLE, this.Enabled)
        values.Put(DownloadItemData_SYSTEM, this.GetSystem())
        values.Put(DownloadItemData_PLATFORM, this.GetPlatform())
        values.Put(DownloadItemData_SPECIAL_NAME, this.GetSpecialName())
        values.Put(DownloadItemData_VERSION, this.GetVersion())
        values.Put(DownloadItemData_CHANGES, this.GetChanges())
        values.Put(DownloadItemData_LICENSE_FILE, this.GetLicenseFile())
        values.Put(DownloadItemData_FILE, this.GetFile())
        values.Put(DownloadItemData_SIZE, fmt.Sprint(this.GetSize()))
        values.Put(DownloadItemData_VALID_TIME, fmt.Sprint(this.GetValidTime()))
        values.Put(DownloadItemData_RETRIES, this.GetRetries())
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *DownloadableItem) GetId() string {
        return this.Id
}


func (this *DownloadableItem) GetSystem() string {
        return this.System
}


func (this *DownloadableItem) GetPlatform() string {
        return this.Platform
}


func (this *DownloadableItem) GetSpecialName() string {
        return this.SpecialName
}


func (this *DownloadableItem) GetVersion() string {
        return this.Version
}


func (this *DownloadableItem) GetChanges() string {
        return this.Changes
}


func (this *DownloadableItem) GetLicenseFile() string {
        return this.LicenseFile
}


func (this *DownloadableItem) GetFile() string {
        return this.File
}


func (this *DownloadableItem) GetSize() int64 {
        return this.Size
}


func (this *DownloadableItem) GetValidTime() int64 {
        return this.ValidTime
}


func (this *DownloadableItem) GetRetries() string {
        return this.Retries
}

