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

type RootStoreCategoryPropertiesFactory struct {
    LogUtil *LogUtil
    CategoryPropertiesInterface CategoryPropertiesInterface
}

// Implements CategoryPropertiesFactoryInterface


func (this *RootStoreCategoryPropertiesFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new RootStoreCategoryPropertiesFactory
func NewRootStoreCategoryPropertiesFactory(args ...any) *RootStoreCategoryPropertiesFactory {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewRootStoreCategoryPropertiesFactory_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) {
            return NewRootStoreCategoryPropertiesFactory_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[IAbPath](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewRootStoreCategoryPropertiesFactory_3(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewRootStoreCategoryPropertiesFactory_4(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of new RootStoreCategoryPropertiesFactory for the arguments")
}
func NewRootStoreCategoryPropertiesFactory_1(transformInfoInterface TransformInfoInterface) *RootStoreCategoryPropertiesFactory {
    this := &RootStoreCategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewRootStoreCategoryProperties(transformInfoInterface)).(CategoryPropertiesInterface)
    return this
}


func NewRootStoreCategoryPropertiesFactory_2(transformInfoInterface TransformInfoInterface, abPath IAbPath) *RootStoreCategoryPropertiesFactory {
    this := &RootStoreCategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewRootStoreCategoryProperties(transformInfoInterface, abPath)).(CategoryPropertiesInterface)
    return this
}


func NewRootStoreCategoryPropertiesFactory_3(transformInfoInterface TransformInfoInterface, node *Node) *RootStoreCategoryPropertiesFactory {
    this := &RootStoreCategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewRootStoreCategoryProperties(transformInfoInterface, node)).(CategoryPropertiesInterface)
    return this
}


func NewRootStoreCategoryPropertiesFactory_4(transformInfoInterface TransformInfoInterface, categoryPropertiesHashMap *HashMap) *RootStoreCategoryPropertiesFactory {
    this := &RootStoreCategoryPropertiesFactory{}
    this.initFields()
        this.CategoryPropertiesInterface = any(NewRootStoreCategoryProperties(transformInfoInterface, categoryPropertiesHashMap)).(CategoryPropertiesInterface)
    return this
}


func (this *RootStoreCategoryPropertiesFactory) GetInstance() CategoryPropertiesInterface {
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

