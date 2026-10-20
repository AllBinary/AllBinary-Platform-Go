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

//TWB - Am I missing the correct Quote Request? Probably not
type QuoteRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    QuoteRequest *QuoteRequest
}


func (this *QuoteRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewQuoteRequestHelper(hashMap *HashMap, pageContext *PageContext) *QuoteRequestHelper {
    this := &QuoteRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        var request *HttpServletRequest = any(pageContext.GetRequest()).(*HttpServletRequest); _ = request
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
        this.QuoteRequest = NewQuoteRequest(this.WeblisketSession.GetUserName(), request)
        var storeName string = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        }
    return this
}


func (this *QuoteRequestHelper) IsValid() bool {
        return this.QuoteRequest.IsValid()
}


func (this *QuoteRequestHelper) ValidationInfo() string {
        return this.QuoteRequest.ValidationInfo()
}


func (this *QuoteRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add QuoteRequest"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "add()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var values IBasicArrayList = this.QuoteRequest.ToVector(); _ = values
            QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance().Insert(values)
            var success string = "New QuoteRequest Successfully added to the QuoteRequest Table"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "add()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update QuoteRequest"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var values *HashMap = this.QuoteRequest.ToHashMap(); _ = values
            QuoteRequestEntityFactory_GetInstance().GetQuoteRequestEntityInstance().Update(this.QuoteRequest.GetUserName(), values)
            var success string = "New QuoteRequest Successfully updated"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *QuoteRequestHelper) Delete() string {
        return this.CommonStrings.NOT_IMPLEMENTED
}

