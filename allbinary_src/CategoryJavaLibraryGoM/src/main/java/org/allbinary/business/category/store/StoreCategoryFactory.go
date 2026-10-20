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

type StoreCategoryFactory struct {
    CategoryAbstractFactory
    LogUtil *LogUtil
    TransformInfoInterface TransformInfoInterface
}

// Implements CategoryFactoryInterface


func (this *StoreCategoryFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewStoreCategoryFactory(transformInfoInterface TransformInfoInterface) *StoreCategoryFactory {
    this := &StoreCategoryFactory{}
    this.CategoryAbstractFactory = *NewCategoryAbstractFactory()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "StoreCategoryFactory(TransformInfoInterface transformInfoInterface)")
        }
        this.TransformInfoInterface = transformInfoInterface
    return this
}


func (this *StoreCategoryFactory) GetRootInstance_1() CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(String, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreCategoryPropertiesFactory(this.TransformInfoInterface)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of StoreCategoryFactory.getRootInstance
func (this *StoreCategoryFactory) GetRootInstance(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GetRootInstance_1()
        }
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            return this.GetRootInstance_2(Java_Arg[IAbPath](args[0]))
        }
    }
    panic("No Java overload of StoreCategoryFactory.getRootInstance for the arguments")
}


func (this *StoreCategoryFactory) GetRootInstance_2(categoryPath IAbPath) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(String, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var level int = CategoryUtil_GetPathLevel(categoryPath); _ = level
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreCategoryPropertiesFactory(this.TransformInfoInterface, categoryPath)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreCategoryFactory) GetRootInstanceFromNode(node *Node) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(String, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreCategoryPropertiesFactory(this.TransformInfoInterface, node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            var categoryPath IAbPath = NewAbPath(CategoryUtil_GetNameFromNode(node), StringUtil_GetInstance().EMPTY_STRING); _ = categoryPath
            var level int = CategoryUtil_GetPathLevel(categoryPath); _ = level
            tryControl1, tryResult1 = 1, any(NewCategory(categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

