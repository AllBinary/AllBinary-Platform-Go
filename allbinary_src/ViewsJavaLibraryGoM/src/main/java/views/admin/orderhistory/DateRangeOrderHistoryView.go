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

type DateRangeOrderHistoryView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    VIEWALLORDERS string
    VIEWORDERSINLASTHOUR string
    VIEWORDERSINLASTDAY string
    VIEWORDERSINLASTWEEK string
    VIEWORDERSINLAST30DAYS string
    Request *HttpServletRequest
    Shipped string
    PartiallyShipped string
    Processing string
    Preprocessing string
    Cancelled string
    DateType string
    ToDate string
    FromDate string
    FromYear string
    FromMonth string
    FromDay string
    FromHour string
    ToYear string
    ToMonth string
    ToDay string
    ToHour string
    Command string
    ON string
    MAXLEN int
}

// Implements ValidationComponentInterface, DomNodeInterface


func (this *DateRangeOrderHistoryView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.VIEWALLORDERS = "View All"
    this.VIEWORDERSINLASTHOUR = "Last Hour"
    this.VIEWORDERSINLASTDAY = "Last Day"
    this.VIEWORDERSINLASTWEEK = "Last Week"
    this.VIEWORDERSINLAST30DAYS = "30 Days"
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    this.ON = "on"
    this.MAXLEN = 15
}


func NewDateRangeOrderHistoryView(transformInfoInterface TransformInfoInterface) *DateRangeOrderHistoryView {
    this := &DateRangeOrderHistoryView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func (this *DateRangeOrderHistoryView) GetFormData() {
        this.DateType = this.Request.GetParameter(OrderHistoryData_DATETYPE)
        this.Preprocessing = this.Request.GetParameter(OrderHistoryData_PREPROCESSINGNAME)
        this.Shipped = this.Request.GetParameter(OrderHistoryData_SHIPPEDNAME)
        this.PartiallyShipped = this.Request.GetParameter(OrderHistoryData_PARTIALLYSHIPPEDNAME)
        this.Processing = this.Request.GetParameter(OrderHistoryData_PROCESSINGNAME)
        this.Cancelled = this.Request.GetParameter(OrderHistoryData_CANCELLEDNAME)
        this.FromYear = this.Request.GetParameter(OrderHistoryData_FROMYEAR)
        this.FromMonth = this.Request.GetParameter(OrderHistoryData_FROMMONTH)
        this.FromDay = this.Request.GetParameter(OrderHistoryData_FROMDAY)
        this.FromHour = this.Request.GetParameter(OrderHistoryData_FROMHOUR)
        this.ToYear = this.Request.GetParameter(OrderHistoryData_TOYEAR)
        this.ToMonth = this.Request.GetParameter(OrderHistoryData_TOMONTH)
        this.ToDay = this.Request.GetParameter(OrderHistoryData_TODAY)
        this.ToHour = this.Request.GetParameter(OrderHistoryData_TOHOUR)
        this.Command = this.Request.GetParameter(HttpStoreComponentView_GLOBALS2.ADMINCOMMAND)
}


func (this *DateRangeOrderHistoryView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *DateRangeOrderHistoryView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            this.AddDomNodeInterfaces()
            tryControl1, tryResult1 = 1, this.HttpStoreComponentView.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DateRangeOrderHistoryView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Attempt to View orders in (fromDate=")
                stringBuffer.Append(this.FromDate)
                stringBuffer.Append(",toDate=")
                stringBuffer.Append(this.ToDate)
                stringBuffer.Append(") and status")
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "view")
            }
            var node *Node = document.CreateElement(OrderData_ORDERS); _ = node
            if this.Preprocessing != "" && strings.Compare(this.Preprocessing, this.ON) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Adding Preprocessing Orders", this, "toXmlNode")
                }
                var orderHistoryVector *BasicArrayList = OrderHistoryEntityFactory_GetInstance().GetOrders(OrderHistoryData_PREPROCESSING, this.FromDate, this.ToDate); _ = orderHistoryVector
                var size int = orderHistoryVector.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var orderHistory *OrderHistory = any(orderHistoryVector.Get(index)).(*OrderHistory); _ = orderHistory
                    node.AppendChild(orderHistory.ToXmlNode(document))
                }
            }
            if this.Shipped != "" && strings.Compare(this.Shipped, this.ON) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Adding Shipped Orders", this, "toXmlNode")
                }
                var orderHistoryVector *BasicArrayList = OrderHistoryEntityFactory_GetInstance().GetOrders(OrderHistoryData_SHIPPED, this.FromDate, this.ToDate); _ = orderHistoryVector
                var size int = orderHistoryVector.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var orderHistory *OrderHistory = any(orderHistoryVector.Get(index)).(*OrderHistory); _ = orderHistory
                    node.AppendChild(orderHistory.ToXmlNode(document))
                }
            }
            if this.PartiallyShipped != "" && strings.Compare(this.PartiallyShipped, this.ON) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Adding Partially Shipped Orders", this, "toDomNode")
                }
                var orderHistoryVector *BasicArrayList = OrderHistoryEntityFactory_GetInstance().GetOrders(OrderHistoryData_PARTIALLYSHIPPED, this.FromDate, this.ToDate); _ = orderHistoryVector
                var size int = orderHistoryVector.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var orderHistory *OrderHistory = any(orderHistoryVector.Get(index)).(*OrderHistory); _ = orderHistory
                    node.AppendChild(orderHistory.ToXmlNode(document))
                }
            }
            if this.Processing != "" && strings.Compare(this.Processing, this.ON) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Adding Processing Orders", this, "toXmlNode")
                }
                var orderHistoryVector *BasicArrayList = OrderHistoryEntityFactory_GetInstance().GetOrders(OrderHistoryData_PROCESSING, this.FromDate, this.ToDate); _ = orderHistoryVector
                var size int = orderHistoryVector.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var orderHistory *OrderHistory = any(orderHistoryVector.Get(index)).(*OrderHistory); _ = orderHistory
                    node.AppendChild(orderHistory.ToXmlNode(document))
                }
            }
            if this.Cancelled != "" && strings.Compare(this.Cancelled, this.ON) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Adding Cancelled Orders", this, "view")
                }
                var orderHistoryVector *BasicArrayList = OrderHistoryEntityFactory_GetInstance().GetOrders(OrderHistoryData_CANCELLED, this.FromDate, this.ToDate); _ = orderHistoryVector
                var size int = orderHistoryVector.Size(); _ = size
                for index := int(0); index < size; index++ {
                    var orderHistory *OrderHistory = any(orderHistoryVector.Get(index)).(*OrderHistory); _ = orderHistory
                    node.AppendChild(orderHistory.ToXmlNode(document))
                }
            }
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PREPROCESSINGNAME, OrderHistoryData_PREPROCESSING))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PROCESSINGNAME, OrderHistoryData_PROCESSING))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_CANCELLEDNAME, OrderHistoryData_CANCELLED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_PARTIALLYSHIPPEDNAME, OrderHistoryData_PARTIALLYSHIPPED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, OrderHistoryData_SHIPPEDNAME, OrderHistoryData_SHIPPED))
            node.AppendChild(ModDomHelper_CreateNameValueNodes(document, HttpStoreComponentView_GLOBALS2.VIEWNAME, HttpStoreComponentView_GLOBALS2.VIEW))
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DateRangeOrderHistoryView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.Put("Exception in validation", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var isValid bool = true; _ = isValid
            if this.DateType == "" || (strings.Compare(this.DateType, OrderHistoryData_TYPELONG) != 0 && strings.Compare(this.DateType, OrderHistoryData_TYPECAESAR) != 0) {
                isValid = false
            }
            if this.DateType != "" {
                if strings.Compare(this.DateType, OrderHistoryData_TYPELONG) == 0 {
                    var calendar *Calendar = Calendar_GetInstance(); _ = calendar
                    var currentTime int64 = calendar.GetTimeInMillis(); _ = currentTime
                    var rangeDate string = StringUtil_GetInstance().EMPTY_STRING; _ = rangeDate
                    if strings.Compare(this.Command, this.VIEWORDERSINLASTHOUR) == 0 {
                        var oneHour int64 = 60 * 60 * 1000; _ = oneHour
                        rangeDate = fmt.Sprint(int64(oneHour))
                    } else if strings.Compare(this.Command, this.VIEWORDERSINLASTDAY) == 0 {
                        var oneDay int64 = 24 * 60 * 60 * 1000; _ = oneDay
                        rangeDate = fmt.Sprint(int64(oneDay))
                    } else if strings.Compare(this.Command, this.VIEWORDERSINLASTWEEK) == 0 {
                        var oneWeek int64 = 7 * 24 * 60 * 60 * 1000; _ = oneWeek
                        rangeDate = fmt.Sprint(int64(oneWeek))
                    } else if strings.Compare(this.Command, this.VIEWORDERSINLAST30DAYS) == 0 {
                        var thirtyDays int64 = 30 * 24 * 60 * 60 * 1000; _ = thirtyDays
                        rangeDate = fmt.Sprint(int64(thirtyDays))
                    } else if strings.Compare(this.Command, this.VIEWALLORDERS) == 0 {
                        rangeDate = fmt.Sprint(int64(currentTime))
                    }
                    var rangeDateLong int64 = int64(int64(rangeDate)); _ = rangeDateLong
                    this.FromDate = fmt.Sprint(int64(currentTime - rangeDateLong))
                    var time string = fmt.Sprint(int64(currentTime)); _ = time
                    this.ToDate = time
                } else if strings.Compare(this.DateType, OrderHistoryData_TYPECAESAR) == 0 {
                    var calendar *Calendar = Calendar_GetInstance(); _ = calendar
                    TimeUtil_GetInstance().SetCalendar(calendar, this.FromYear, this.FromMonth, this.FromDay, this.FromHour)
                    this.FromDate = fmt.Sprint(int64(calendar.GetTimeInMillis()))
                    TimeUtil_GetInstance().SetCalendar(calendar, this.ToYear, this.ToMonth, this.ToDay, this.ToHour)
                    this.ToDate = fmt.Sprint(int64(calendar.GetTimeInMillis()))
                }
                if StringValidationUtil_GetInstance().IsEmpty(this.ToDate) || len(this.ToDate) > this.MAXLEN {
                    isValid = false
                }
                if StringValidationUtil_GetInstance().IsEmpty(this.FromDate) || len(this.FromDate) > this.MAXLEN {
                    isValid = false
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Attempt to View orders in (fromDate=")
                stringBuffer.Append(this.FromDate)
                stringBuffer.Append(",toDate=")
                stringBuffer.Append(this.ToDate)
                stringBuffer.Append(")")
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DateRangeOrderHistoryView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *DateRangeOrderHistoryView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *DateRangeOrderHistoryView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Getting Validation Info"; return
                }
            }()
            var result *StringMaker = NewStringMaker(); _ = result
            if this.DateType == "" || (strings.Compare(this.DateType, OrderHistoryData_TYPELONG) != 0 && strings.Compare(this.DateType, OrderHistoryData_TYPECAESAR) != 0) {
                result.Append("DATETYPE not recognized")
            }
            if StringValidationUtil_GetInstance().IsEmpty(this.ToDate) || len(this.ToDate) > this.MAXLEN {
                result.Append("Invalid To Date")
            }
            if StringValidationUtil_GetInstance().IsEmpty(this.FromDate) || len(this.FromDate) > this.MAXLEN {
                result.Append("Invalid From Date")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(result); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

