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

type BasicItem struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BasicItemData *BasicItemData
    ItemId string
    Number string
    InBaskets string
    Weight string
    Enabled string
    NewOrUsed string
    Summary string
    Distributor string
    IdUsedByDistributor string
    ProducedBy string
    ProductionDate string
    StartProductionDate string
    Description string
    Keywords string
    Category string
    Type string
    SmallImage string
    MediumImage string
    LargeImage string
    Price *Money
    Comment string
    Customs string
    Downloads string
    Groups string
    Options string
    Permissions string
    Specials string
    Downloadable bool
    TimeEntered string
    LastModified string
}

// Implements ItemInterface

// The Java class BasicItem has subclasses so its variables are this interface
type IBasicItem interface {
    BasicItem_() *BasicItem
    SetId(itemId string)
    SetNumber(number string)
    SetInBaskets(value string)
    SetWeight(value string)
    SetNewOrUsed(value string)
    SetSummary(value string)
    SetDistributor(value string)
    SetIdUsedByDistributor(value string)
    SetProducedBy(value string)
    SetProductionDate(value string)
    SetStartProductionDate(value string)
    SetDescription(value string)
    SetKeywords(value string)
    SetCategory(value string)
    SetType(value string)
    SetSmallImage(value string)
    SetMediumImage(value string)
    SetLargeImage(value string)
    SetTimeEntered(value string)
    SetLastModified(value string)
    SetPrice(value *Money)
    SetComment(value string)
    SetCustoms(value string)
    SetDownloads(value string)
    SetGroups(value string)
    SetOptions(value string)
    SetPermissions(value string)
    SetSpecials(value string)
    SetEnabled(value string)
    GetId() string
    GetNumber() string
    GetInBaskets() string
    GetWeight() string
    GetNewOrUsed() string
    GetSummary() string
    GetDistributor() string
    GetIdUsedByDistributor() string
    GetProducedBy() string
    GetProductionDate() string
    GetStartProductionDate() string
    GetDescription() string
    GetKeywords() string
    GetCategory() string
    GetType() string
    GetSmallImage() string
    GetMediumImage() string
    GetLargeImage() string
    GetTimeEntered() string
    GetLastModified() string
    GetPrice() *Money
    GetComment() string
    GetCustoms() string
    GetDownloads() string
    GetGroups() string
    GetOptions() string
    GetPermissions() string
    GetSpecials() string
    SetDownloadable(downloadable bool)
    IsDownloadable() bool
    GetEnabled() string
    GetTotal() *Money
    ToHashMap() *HashMap
    ToVector() IBasicArrayList
    GetKey() any
}

func (this *BasicItem) BasicItem_() *BasicItem {
    return this
}


func (this *BasicItem) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}




// Java overloads of new BasicItem
func NewBasicItem(args ...any) *BasicItem {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewBasicItem_1(Java_Arg[*HashMap](args[0]))
        }
        if len(args) == 0 {
            return NewBasicItem_2()
        }
    }
    panic("No Java overload of new BasicItem for the arguments")
}
func NewBasicItem_1(itemHashMap *HashMap) *BasicItem {
    this := &BasicItem{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "Constructor(HashMap)")
        }
        this.ItemId = any(itemHashMap.Get(this.BasicItemData.ID)).(string)
        this.Number = any(itemHashMap.Get(this.BasicItemData.NUMBER)).(string)
        this.InBaskets = any(itemHashMap.Get(this.BasicItemData.INBASKETS)).(string)
        this.Weight = any(itemHashMap.Get(this.BasicItemData.WEIGHT)).(string)
        this.Enabled = any(itemHashMap.Get(EntryData_GetInstance().ENABLE)).(string)
        this.NewOrUsed = any(itemHashMap.Get(this.BasicItemData.NEWORUSED)).(string)
        this.Summary = any(itemHashMap.Get(this.BasicItemData.SUMMARY)).(string)
        this.Distributor = any(itemHashMap.Get(this.BasicItemData.DISTRIBUTOR)).(string)
        this.IdUsedByDistributor = any(itemHashMap.Get(this.BasicItemData.IDUSEDBYDISTRIBUTOR)).(string)
        this.ProducedBy = any(itemHashMap.Get(this.BasicItemData.PRODUCEDBY)).(string)
        this.ProductionDate = any(itemHashMap.Get(this.BasicItemData.PRODUCTIONDATE)).(string)
        this.StartProductionDate = any(itemHashMap.Get(this.BasicItemData.STARTPRODUCTIONDATE)).(string)
        this.Description = any(itemHashMap.Get(this.BasicItemData.DESCRIPTION)).(string)
        this.Keywords = any(itemHashMap.Get(this.BasicItemData.KEYWORDS)).(string)
        this.Category = any(itemHashMap.Get(this.BasicItemData.CATEGORY)).(string)
        this.Type = any(itemHashMap.Get(this.BasicItemData.TYPE)).(string)
        this.SmallImage = any(itemHashMap.Get(this.BasicItemData.SMALLIMAGE)).(string)
        this.MediumImage = any(itemHashMap.Get(this.BasicItemData.MEDIUMIMAGE)).(string)
        this.LargeImage = any(itemHashMap.Get(this.BasicItemData.LARGEIMAGE)).(string)
        this.TimeEntered = any(itemHashMap.Get(EntryData_GetInstance().TIMECREATED)).(string)
        this.LastModified = any(itemHashMap.Get(EntryData_GetInstance().LASTMODIFIED)).(string)
        this.Price = NewMoney(any(itemHashMap.Get(this.BasicItemData.PRICE)).(string))
        this.Comment = any(itemHashMap.Get(this.BasicItemData.COMMENT)).(string)
        this.Customs = any(itemHashMap.Get(this.BasicItemData.CUSTOMS)).(string)
        this.SetDownloads(any(itemHashMap.Get(this.BasicItemData.DOWNLOADS)).(string))
        this.Groups = any(itemHashMap.Get(this.BasicItemData.GROUPS)).(string)
        this.Options = any(itemHashMap.Get(this.BasicItemData.OPTIONS)).(string)
        this.Permissions = any(itemHashMap.Get(this.BasicItemData.PERMISSIONS)).(string)
        this.Specials = any(itemHashMap.Get(this.BasicItemData.SPECIALS)).(string)
    return this
}


func NewBasicItem_2() *BasicItem {
    this := &BasicItem{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
        var EMPTY string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY
        this.ItemId = NewProductIdGenerator().GetNext()
        this.Number = CommonPhoneStrings_GetInstance().ZERO
        this.InBaskets = EMPTY
        this.Weight = EMPTY
        this.Enabled = EMPTY
        this.NewOrUsed = EMPTY
        this.Summary = EMPTY
        this.Distributor = EMPTY
        this.IdUsedByDistributor = EMPTY
        this.ProducedBy = EMPTY
        this.ProductionDate = EMPTY
        this.StartProductionDate = EMPTY
        this.Description = EMPTY
        this.Keywords = EMPTY
        this.Category = EMPTY
        this.Type = EMPTY
        this.SmallImage = EMPTY
        this.MediumImage = EMPTY
        this.LargeImage = EMPTY
        this.TimeEntered = EMPTY
        this.LastModified = EMPTY
        this.Price = NewMoney()
        this.Comment = EMPTY
        this.Customs = EMPTY
        this.SetDownloads(EMPTY)
        this.Groups = EMPTY
        this.Options = EMPTY
        this.Permissions = EMPTY
        this.Specials = EMPTY
    return this
}


func (this *BasicItem) SetId(itemId string) {
        this.ItemId = itemId
}


func (this *BasicItem) SetNumber(number string) {
        this.Number = number
}


func (this *BasicItem) SetInBaskets(value string) {
        this.InBaskets = value
}


func (this *BasicItem) SetWeight(value string) {
        this.Weight = value
}


func (this *BasicItem) SetNewOrUsed(value string) {
        this.NewOrUsed = value
}


func (this *BasicItem) SetSummary(value string) {
        this.Summary = value
}


func (this *BasicItem) SetDistributor(value string) {
        this.Distributor = value
}


func (this *BasicItem) SetIdUsedByDistributor(value string) {
        this.IdUsedByDistributor = value
}


func (this *BasicItem) SetProducedBy(value string) {
        this.ProducedBy = value
}


func (this *BasicItem) SetProductionDate(value string) {
        this.ProductionDate = value
}


func (this *BasicItem) SetStartProductionDate(value string) {
        this.StartProductionDate = value
}


func (this *BasicItem) SetDescription(value string) {
        this.Description = value
}


func (this *BasicItem) SetKeywords(value string) {
        this.Keywords = value
}


func (this *BasicItem) SetCategory(value string) {
        this.Category = value
}


func (this *BasicItem) SetType(value string) {
        this.Type = value
}


func (this *BasicItem) SetSmallImage(value string) {
        this.SmallImage = value
}


func (this *BasicItem) SetMediumImage(value string) {
        this.MediumImage = value
}


func (this *BasicItem) SetLargeImage(value string) {
        this.LargeImage = value
}


func (this *BasicItem) SetTimeEntered(value string) {
        this.TimeEntered = value
}


func (this *BasicItem) SetLastModified(value string) {
        this.LastModified = value
}


func (this *BasicItem) SetPrice(value *Money) {
        this.Price = value
}


func (this *BasicItem) SetComment(value string) {
        this.Comment = value
}


func (this *BasicItem) SetCustoms(value string) {
        this.Customs = value
}


func (this *BasicItem) SetDownloads(value string) {
        this.Downloads = value
        if !StringValidationUtil_GetInstance().IsEmpty(this.Downloads) {
            var downloadInteger int = fmt.Sprint(this.Downloads); _ = downloadInteger
            if int(downloadInteger) != 0 {
                this.SetDownloadable(true)
            }
        }
}


func (this *BasicItem) SetGroups(value string) {
        this.Groups = value
}


func (this *BasicItem) SetOptions(value string) {
        this.Options = value
}


func (this *BasicItem) SetPermissions(value string) {
        this.Permissions = value
}


func (this *BasicItem) SetSpecials(value string) {
        this.Specials = value
}


func (this *BasicItem) SetEnabled(value string) {
        this.Enabled = value
}


func (this *BasicItem) GetId() string {
        return this.ItemId
}


func (this *BasicItem) GetNumber() string {
        return this.Number
}


func (this *BasicItem) GetInBaskets() string {
        return this.InBaskets
}


func (this *BasicItem) GetWeight() string {
        return this.Weight
}


func (this *BasicItem) GetNewOrUsed() string {
        return this.NewOrUsed
}


func (this *BasicItem) GetSummary() string {
        return this.Summary
}


func (this *BasicItem) GetDistributor() string {
        return this.Distributor
}


func (this *BasicItem) GetIdUsedByDistributor() string {
        return this.IdUsedByDistributor
}


func (this *BasicItem) GetProducedBy() string {
        return this.ProducedBy
}


func (this *BasicItem) GetProductionDate() string {
        return this.ProductionDate
}


func (this *BasicItem) GetStartProductionDate() string {
        return this.StartProductionDate
}


func (this *BasicItem) GetDescription() string {
        return this.Description
}


func (this *BasicItem) GetKeywords() string {
        return this.Keywords
}


func (this *BasicItem) GetCategory() string {
        return this.Category
}


func (this *BasicItem) GetType() string {
        return this.Type
}


func (this *BasicItem) GetSmallImage() string {
        return this.SmallImage
}


func (this *BasicItem) GetMediumImage() string {
        return this.MediumImage
}


func (this *BasicItem) GetLargeImage() string {
        return this.LargeImage
}


func (this *BasicItem) GetTimeEntered() string {
        return this.TimeEntered
}


func (this *BasicItem) GetLastModified() string {
        return this.LastModified
}


func (this *BasicItem) GetPrice() *Money {
        return this.Price
}


func (this *BasicItem) GetComment() string {
        return this.Comment
}


func (this *BasicItem) GetCustoms() string {
        return this.Customs
}


func (this *BasicItem) GetDownloads() string {
        return this.Downloads
}


func (this *BasicItem) GetGroups() string {
        return this.Groups
}


func (this *BasicItem) GetOptions() string {
        return this.Options
}


func (this *BasicItem) GetPermissions() string {
        return this.Permissions
}


func (this *BasicItem) GetSpecials() string {
        return this.Specials
}


func (this *BasicItem) SetDownloadable(downloadable bool) {
        this.Downloadable = downloadable
}


func (this *BasicItem) IsDownloadable() bool {
        return this.Downloadable
}


func (this *BasicItem) GetEnabled() string {
        return this.Enabled
}


func (this *BasicItem) GetTotal() *Money {
        var itemTotal *Money = NewMoney(this.Price); _ = itemTotal
        itemTotal.Multiply(int(int(this.Number)))
        return itemTotal
}


func (this *BasicItem) ToHashMap() *HashMap {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "toHashMap")
        }
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(this.BasicItemData.ID, this.ItemId)
        hashMap.Put(this.BasicItemData.NUMBER, this.Number)
        hashMap.Put(this.BasicItemData.INBASKETS, this.InBaskets)
        hashMap.Put(this.BasicItemData.WEIGHT, this.Weight)
        hashMap.Put(EntryData_GetInstance().ENABLE, this.Enabled)
        hashMap.Put(this.BasicItemData.NEWORUSED, this.NewOrUsed)
        hashMap.Put(this.BasicItemData.SUMMARY, this.Summary)
        hashMap.Put(this.BasicItemData.DISTRIBUTOR, this.Distributor)
        hashMap.Put(this.BasicItemData.IDUSEDBYDISTRIBUTOR, this.IdUsedByDistributor)
        hashMap.Put(this.BasicItemData.PRODUCEDBY, this.ProducedBy)
        hashMap.Put(this.BasicItemData.PRODUCTIONDATE, this.ProductionDate)
        hashMap.Put(this.BasicItemData.STARTPRODUCTIONDATE, this.StartProductionDate)
        hashMap.Put(this.BasicItemData.DESCRIPTION, this.Description)
        hashMap.Put(this.BasicItemData.KEYWORDS, this.Keywords)
        hashMap.Put(this.BasicItemData.CATEGORY, this.Category)
        hashMap.Put(this.BasicItemData.TYPE, this.Type)
        hashMap.Put(this.BasicItemData.SMALLIMAGE, this.SmallImage)
        hashMap.Put(this.BasicItemData.MEDIUMIMAGE, this.MediumImage)
        hashMap.Put(this.BasicItemData.LARGEIMAGE, this.LargeImage)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        hashMap.Put(EntryData_GetInstance().LASTMODIFIED, time)
        hashMap.Put(this.BasicItemData.PRICE, fmt.Sprint(this.Price))
        hashMap.Put(this.BasicItemData.COMMENT, this.Comment)
        hashMap.Put(this.BasicItemData.CUSTOMS, this.Customs)
        hashMap.Put(this.BasicItemData.DOWNLOADS, this.GetDownloads())
        hashMap.Put(this.BasicItemData.GROUPS, this.Groups)
        hashMap.Put(this.BasicItemData.OPTIONS, this.Options)
        hashMap.Put(this.BasicItemData.PERMISSIONS, this.Permissions)
        hashMap.Put(this.BasicItemData.SPECIALS, this.Specials)
        return hashMap
}


func (this *BasicItem) ToVector() IBasicArrayList {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "toVector")
        }
        var values IBasicArrayList = NewBasicArrayListD(); _ = values
        values.Add(this.ItemId)
        values.Add(this.Number)
        values.Add(this.InBaskets)
        values.Add(this.Weight)
        values.Add(this.Enabled)
        values.Add(this.NewOrUsed)
        values.Add(this.Summary)
        values.Add(this.Distributor)
        values.Add(this.IdUsedByDistributor)
        values.Add(this.ProducedBy)
        values.Add(this.ProductionDate)
        values.Add(this.StartProductionDate)
        values.Add(this.Description)
        values.Add(this.Keywords)
        values.Add(this.Category)
        values.Add(this.Type)
        values.Add(this.GetSmallImage())
        values.Add(this.GetMediumImage())
        values.Add(this.GetLargeImage())
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Add(time)
        values.Add(time)
        values.Add(fmt.Sprint(this.Price))
        values.Add(this.Comment)
        values.Add(this.Customs)
        values.Add(this.GetDownloads())
        values.Add(this.Groups)
        values.Add(this.Options)
        values.Add(this.Permissions)
        values.Add(this.Specials)
        return values
}


// throws Exception
func (this *BasicItem) GetKey() any {
        return any(this.GetId()).(any)
}

