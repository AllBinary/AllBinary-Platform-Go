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

type BasicStoreFrontsEntity struct {
    InitSql
    LogUtil *LogUtil
    TableName string
}


func (this *BasicStoreFrontsEntity) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TableName = "storefronts"
}


func NewBasicStoreFrontsEntity() *BasicStoreFrontsEntity {
    this := &BasicStoreFrontsEntity{}
        this.InitSql = *NewInitSql(NewUserDbInitInfo())
        this.initFields()
        this.SetTable(this.TableName)
    return this
}


// throws Exception
func (this *BasicStoreFrontsEntity) GetStoreFrontInterface(name string) BasicStoreFrontInterface {
        var keysAndValues *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = keysAndValues
        keysAndValues.Put(StoreFrontData_GetInstance().NAME, name)
        var storeHashMap *HashMap = this.InitSql.GetRow(keysAndValues); _ = storeHashMap
        if storeHashMap == nil {
            panic(errors.New(fmt.Sprint("No Such Store: " + name)))
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().FACTORYERROR) {
            this.LogUtil.PutF("StoreFront: " + fmt.Sprint(storeHashMap), this, "getStore()")
        }
        return any(NewBasicStoreFront(storeHashMap)).(BasicStoreFrontInterface)
}

