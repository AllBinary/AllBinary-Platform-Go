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

type QuoteRequest struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Id string
    UserName string
    ProjectInfo string
    UserComments string
    Budget string
    TimeFrame string
    Comments string
}


func (this *QuoteRequest) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
}




// Java overloads of new QuoteRequest
func NewQuoteRequest(args ...any) *QuoteRequest {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*HttpServletRequest](args[1], pass) {
            return NewQuoteRequest_1(Java_Arg[string](args[0]), Java_Arg[*HttpServletRequest](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewQuoteRequest_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new QuoteRequest for the arguments")
}
func NewQuoteRequest_1(userName string, request *HttpServletRequest) *QuoteRequest {
    this := &QuoteRequest{}
    this.initFields()
        var quoteRequestData *QuoteRequestData = QuoteRequestData_GetInstance(); _ = quoteRequestData
        this.Id = NewQuoteRequestIdGenerator().GetNext()
        this.SetUserName(userName)
        this.SetProjectInfo(any(request.GetParameter(quoteRequestData.PROJECT_INFO)).(string))
        this.SetUserComments(any(request.GetParameter(quoteRequestData.CUSTOMER_COMMENTS)).(string))
        this.SetBudget(any(request.GetParameter(quoteRequestData.BUDGET)).(string))
        this.SetTimeFrame(any(request.GetParameter(quoteRequestData.TIMEFRAME)).(string))
        this.SetComments(any(request.GetParameter(quoteRequestData.COMMENTS)).(string))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
            this.LogUtil.PutF("Created with: " + fmt.Sprint(this.ToHashMap()), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func NewQuoteRequest_2(hashMap *HashMap) *QuoteRequest {
    this := &QuoteRequest{}
    this.initFields()
        var quoteRequestData *QuoteRequestData = QuoteRequestData_GetInstance(); _ = quoteRequestData
        this.Id = any(hashMap.Get(quoteRequestData.ID)).(string)
        this.SetUserName(any(hashMap.Get(UserData_USERNAME)).(string))
        this.SetProjectInfo(any(hashMap.Get(quoteRequestData.PROJECT_INFO)).(string))
        this.SetUserComments(any(hashMap.Get(quoteRequestData.CUSTOMER_COMMENTS)).(string))
        this.SetBudget(any(hashMap.Get(quoteRequestData.BUDGET)).(string))
        this.SetTimeFrame(any(hashMap.Get(quoteRequestData.TIMEFRAME)).(string))
        this.SetComments(any(hashMap.Get(quoteRequestData.COMMENTS)).(string))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
            this.LogUtil.PutF("Created with: " + fmt.Sprint(this.ToHashMap()), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *QuoteRequest) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if this.UserName == "" || len(this.UserName) < 4 || len(this.UserName) > 250 {
                valid = false
            }
            if this.ProjectInfo == "" || len(this.ProjectInfo) < 0 || len(this.ProjectInfo) > AbSqlData_MAXBLOB {
                valid = false
            }
            if this.UserComments == "" || len(this.UserComments) < 0 || len(this.UserComments) > AbSqlData_MAXBLOB {
                valid = false
            }
            if this.Budget == "" || len(this.Budget) < 0 || len(this.Budget) > AbSqlData_MAXBLOB {
                valid = false
            }
            if this.TimeFrame == "" || len(this.TimeFrame) < 0 || len(this.TimeFrame) > AbSqlData_MAXBLOB {
                valid = false
            }
            if this.Comments == "" || len(this.Comments) < 0 || len(this.Comments) > AbSqlData_MAXBLOB {
                valid = false
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteRequest) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.UserName == "" || len(this.UserName) < 4 || len(this.UserName) > 250 {
                stringBuffer.Append("Please enter a User Name with more than 4 characters.<br>")
            }
            if this.ProjectInfo == "" || len(this.ProjectInfo) < 0 || len(this.ProjectInfo) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Please enter Project Info with less than " + fmt.Sprint(AbSqlData_MAXBLOB) + " characters.<br>")
            }
            if this.UserComments == "" || len(this.UserComments) < 0 || len(this.UserComments) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Please enter User Comments with less than " + fmt.Sprint(AbSqlData_MAXBLOB) + " characters.<br>")
            }
            if this.Budget == "" || len(this.Budget) < 0 || len(this.Budget) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Please enter Budget with less than " + fmt.Sprint(AbSqlData_MAXBLOB) + " characters.<br>")
            }
            if this.TimeFrame == "" || len(this.TimeFrame) < 0 || len(this.TimeFrame) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Please enter time frame with less than " + fmt.Sprint(AbSqlData_MAXBLOB) + " characters.<br>")
            }
            if this.Comments == "" || len(this.Comments) < 0 || len(this.Comments) > AbSqlData_MAXBLOB {
                stringBuffer.Append("Please enter comments with less than " + fmt.Sprint(AbSqlData_MAXBLOB) + " characters.<br>")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteRequest) ToVector() IBasicArrayList {
        var values IBasicArrayList = NewBasicArrayListD(); _ = values
        values.Add(this.Id)
        values.Add(this.UserName)
        values.Add(this.ProjectInfo)
        values.Add(this.UserComments)
        values.Add(this.Budget)
        values.Add(this.TimeFrame)
        values.Add(this.Comments)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Add(time)
        values.Add(time)
        return values
}


func (this *QuoteRequest) ToHashMap() *HashMap {
        var quoteRequestData *QuoteRequestData = QuoteRequestData_GetInstance(); _ = quoteRequestData
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        values.Put(UserData_USERNAME, this.UserName)
        values.Put(quoteRequestData.PROJECT_INFO, this.ProjectInfo)
        values.Put(quoteRequestData.CUSTOMER_COMMENTS, this.UserComments)
        values.Put(quoteRequestData.BUDGET, this.Budget)
        values.Put(quoteRequestData.TIMEFRAME, this.TimeFrame)
        values.Put(quoteRequestData.COMMENTS, this.Comments)
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        values.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return values
}


func (this *QuoteRequest) SetUserName(value string) {
        this.UserName = value
}


func (this *QuoteRequest) SetProjectInfo(value string) {
        this.ProjectInfo = value
}


func (this *QuoteRequest) SetUserComments(value string) {
        this.UserComments = value
}


func (this *QuoteRequest) SetBudget(value string) {
        this.Budget = value
}


func (this *QuoteRequest) SetTimeFrame(value string) {
        this.TimeFrame = value
}


func (this *QuoteRequest) SetComments(value string) {
        this.Comments = value
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        if this.Comments == "" {
            this.Comments = stringUtil.EMPTY_STRING
        }
}


func (this *QuoteRequest) GetUserName() string {
        return this.UserName
}


func (this *QuoteRequest) GetProjectInfo() string {
        return this.ProjectInfo
}


func (this *QuoteRequest) GetUserComments() string {
        return this.UserComments
}


func (this *QuoteRequest) GetBudget() string {
        return this.Budget
}


func (this *QuoteRequest) GetTimeFrame() string {
        return this.TimeFrame
}


func (this *QuoteRequest) GetComments() string {
        return this.Comments
}

