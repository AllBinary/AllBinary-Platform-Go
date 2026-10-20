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

type BasketHelper struct {
    TagHelper
    LogUtil *LogUtil
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
    MAX int
}


func (this *BasketHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    this.MAX = 200
}


func NewBasketHelper(hashMap *HashMap, pageContext *PageContext) *BasketHelper {
    this := &BasketHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.PropertiesHashMap = hashMap
        this.PageContext = pageContext
        var storeName string = any(this.PropertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        }
        this.WeblisketSession = NewWeblisketSession(hashMap, pageContext)
    return this
}


func (this *BasketHelper) IsBasketEmpty() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "isBasketEmpty()", e)
            }
            tryControl1, tryResult1 = 1, true; return
                }
            }()
            var basket BasketInterface = this.WeblisketSession.GetOrder().GetBasket(); _ = basket
            if int(basket.GetNumberOfItems()) <= 0 {
                tryControl1, tryResult1 = 1, true; return
            } else {
                tryControl1, tryResult1 = 1, false; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

