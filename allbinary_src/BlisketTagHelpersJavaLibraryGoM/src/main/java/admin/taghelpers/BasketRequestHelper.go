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

type BasketRequestHelper struct {
    TagHelper
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    WeblisketSession *WeblisketSession
    StoreFrontInterface StoreFrontInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
    Request *HttpServletRequest
    Id string
    Num string
    MAX int
}


func (this *BasketRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
    
    this.MAX = 200
}


func NewBasketRequestHelper(propertiesHashMap *HashMap, pageContext *PageContext) *BasketRequestHelper {
    this := &BasketRequestHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        this.PropertiesHashMap = propertiesHashMap
        this.PageContext = pageContext
        this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
        var storeName string = any(propertiesHashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
        } else {
            this.StoreFrontInterface = NewStoreFront()
        }
        this.WeblisketSession = NewWeblisketSession(propertiesHashMap, pageContext)
        this.GetFormData()
    return this
}


func (this *BasketRequestHelper) GetFormData() {
        this.Id = this.Request.GetParameter(this.BasicItemData.ID)
        this.Num = this.Request.GetParameter(BasketData_ITEMTOTALINBASKET)
}


func (this *BasketRequestHelper) AddItemToBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "addItemToBasket()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.Id != "" && this.Num != "" {
                var basket BasketInterface = this.WeblisketSession.GetOrder().GetBasket(); _ = basket
                if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.Id) == nil {
                    tryControl1, tryResult1 = 1, false; return
                }
                basket.AddItem(this.Id, this.Num)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successfully added item to Basket", this, "addItemToBasket()")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasketRequestHelper) RemoveItemFromBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "removeItemFromBasket()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.Id != "" {
                var basket BasketInterface = this.WeblisketSession.GetOrder().GetBasket(); _ = basket
                basket.RemoveItem(this.Id)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successfully removed item from Basket", this, "removeItemFromBasket()")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasketRequestHelper) AdjustBasket() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "adjustBasket()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var basket BasketInterface = this.WeblisketSession.GetOrder().GetBasket(); _ = basket
            if this.Id != "" && this.Num != "" {
                basket.AdjustItem(this.Id, this.Num)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Successfull Adjusted Basket", this, "adjustBasket()")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

