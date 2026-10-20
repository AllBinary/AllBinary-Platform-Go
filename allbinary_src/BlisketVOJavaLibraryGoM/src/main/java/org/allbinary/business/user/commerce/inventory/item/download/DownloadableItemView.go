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

type DownloadableItemView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    DownloadableItem *DownloadableItem
}

// Implements DomNodeInterface


func (this *DownloadableItemView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewDownloadableItemView(downloadableItem *DownloadableItem) *DownloadableItemView {
    this := &DownloadableItemView{}
    this.initFields()
        this.DownloadableItem = downloadableItem
    return this
}


// throws Exception
func (this *DownloadableItemView) ToXmlNode(document *Document) *Node {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "toXmlNode")
        }
        var hashMap *HashMap = this.DownloadableItem.ToHashMap(); _ = hashMap
        var totalTime int64 = int64(this.DownloadableItem.GetValidTime()); _ = totalTime
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var year int = calendar.Get(Calendar_YEAR); _ = year
        var month int = calendar.Get(Calendar_MONTH); _ = month
        var day int = calendar.Get(Calendar_DAY_OF_MONTH); _ = day
        var hour int = calendar.Get(Calendar_HOUR); _ = hour
        var minute int = calendar.Get(Calendar_MINUTE); _ = minute
        var second int = calendar.Get(Calendar_SECOND); _ = second
        calendar.SetTimeInMillis(calendar.GetTimeInMillis() + totalTime)
        var yearDelta int = calendar.Get(Calendar_YEAR); _ = yearDelta
        var monthDelta int = calendar.Get(Calendar_MONTH); _ = monthDelta
        var dayDelta int = calendar.Get(Calendar_DAY_OF_MONTH); _ = dayDelta
        var hourDelta int = calendar.Get(Calendar_HOUR); _ = hourDelta
        var minuteDelta int = calendar.Get(Calendar_MINUTE); _ = minuteDelta
        var secondDelta int = calendar.Get(Calendar_SECOND); _ = secondDelta
        hashMap.Put(DownloadItemData_VALID_TIME_YEARS, fmt.Sprint(yearDelta - year))
        hashMap.Put(DownloadItemData_VALID_TIME_MONTHS, fmt.Sprint(monthDelta - month))
        hashMap.Put(DownloadItemData_VALID_TIME_DAYS, fmt.Sprint(dayDelta - day))
        hashMap.Put(DownloadItemData_VALID_TIME_HOURS, fmt.Sprint(hourDelta - hour))
        hashMap.Put(DownloadItemData_VALID_TIME_MINUTES, fmt.Sprint(minuteDelta - minute))
        hashMap.Put(DownloadItemData_VALID_TIME_SECONDS, fmt.Sprint(secondDelta - second))
        var keySet Set = hashMap.KeySet(); _ = keySet
        var node *Node = document.CreateElement(DownloadItemData_NAME); _ = node
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var nameArray []any = keySet.ToArray(); _ = nameArray
        var size int = len(nameArray); _ = size
        for index := int(0); index < size; index++ {
            var name string = any(nameArray[index]).(string); _ = name
            var value string = any(hashMap.Get(name)).(string); _ = value
            value = stringUtil.GetNonNull(value)
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, name, value))
        }
        return node
}


func (this *DownloadableItemView) ToXmlDoc() *Document {
        return nil
}


func (this *DownloadableItemView) View() string {
        return nil
}

