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

type ShippingAddressesEntityFactory struct {
    LogUtil *LogUtil
}

var ShippingAddressesEntityFactory_Instance *ShippingAddressesEntityFactory = NewShippingAddressesEntityFactory()

func (this *ShippingAddressesEntityFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func ShippingAddressesEntityFactory_GetInstance() *ShippingAddressesEntityFactory {
        return ShippingAddressesEntityFactory_Instance
}


func NewShippingAddressesEntityFactory() *ShippingAddressesEntityFactory {
    this := &ShippingAddressesEntityFactory{}
    this.initFields()
    return this
}


func (this *ShippingAddressesEntityFactory) GetInstance(userName string) IStreetAddressesEntity {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IStreetAddressesEntity) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "getShippingAddressesEntityInstance()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, NewShippingAddressesEntity(userName); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

