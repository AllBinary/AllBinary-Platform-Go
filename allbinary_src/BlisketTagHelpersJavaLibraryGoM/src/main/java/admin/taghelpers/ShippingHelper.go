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

type ShippingHelper struct {
    TagHelper
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
    Request *HttpServletRequest
    ShippingType string
}


func (this *ShippingHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
}


func NewShippingHelper(hashMap *HashMap, pageContext *PageContext) *ShippingHelper {
    this := &ShippingHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.PageContext = pageContext
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        var storeName string = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        }
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
        this.GetFormData()
    return this
}


func (this *ShippingHelper) GetFormData() {
        this.ShippingType = this.Request.GetParameter(ShippingMethodData_NAME)
}


func (this *ShippingHelper) SetShippingType() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view Shipping Type"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setShippingType()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Successfully Set Shipping Type"; _ = success
            var order OrderInterface = this.WeblisketSession.GetOrder(); _ = order
            order.SetShippingMethod(this.ShippingType)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "setShippingType()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

