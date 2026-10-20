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

type CategoryPropertiesFactory struct {
    LogUtil *LogUtil
    CategoryPropertiesInterface CategoryPropertiesInterface
}

// Implements CategoryPropertiesFactoryInterface


func (this *CategoryPropertiesFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new CategoryPropertiesFactory
func NewCategoryPropertiesFactory(args ...any) *CategoryPropertiesFactory {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewCategoryPropertiesFactory_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCategoryPropertiesFactory_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewCategoryPropertiesFactory_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new CategoryPropertiesFactory for the arguments")
}
func NewCategoryPropertiesFactory_1(name string) *CategoryPropertiesFactory {
    this := &CategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewCategoryProperties(name)).(CategoryPropertiesInterface)
    return this
}


func NewCategoryPropertiesFactory_2(node *Node) *CategoryPropertiesFactory {
    this := &CategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewCategoryProperties(node)).(CategoryPropertiesInterface)
    return this
}


func NewCategoryPropertiesFactory_3(categoryPropertiesHashMap *HashMap) *CategoryPropertiesFactory {
    this := &CategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewCategoryProperties(categoryPropertiesHashMap)).(CategoryPropertiesInterface)
    return this
}


func (this *CategoryPropertiesFactory) GetInstance() CategoryPropertiesInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryPropertiesInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, any(this.CategoryPropertiesInterface).(CategoryPropertiesInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

