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

type InventoryEntityFactory struct {
    LogUtil *LogUtil
}

var InventoryEntityFactory_Instance *InventoryEntityFactory = NewInventoryEntityFactory()

func (this *InventoryEntityFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func InventoryEntityFactory_GetInstance() *InventoryEntityFactory {
        return InventoryEntityFactory_Instance
}


func NewInventoryEntityFactory() *InventoryEntityFactory {
    this := &InventoryEntityFactory{}
    this.initFields()
    return this
}


func (this *InventoryEntityFactory) GetInventoryEntityInstance() *InventoryEntity {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *InventoryEntity) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "getInventoryEntityInstance()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, NewInventoryEntity(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

