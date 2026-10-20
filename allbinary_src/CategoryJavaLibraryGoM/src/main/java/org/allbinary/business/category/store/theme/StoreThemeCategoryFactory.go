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

type StoreThemeCategoryFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TransformInfoInterface TransformInfoInterface
}

// Implements CategoryFactoryInterface


func (this *StoreThemeCategoryFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewStoreThemeCategoryFactory(transformInfoInterface TransformInfoInterface) *StoreThemeCategoryFactory {
    this := &StoreThemeCategoryFactory{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "StoreCategoryFactory(TransformInfoInterface transformInfoInterface)")
        }
        this.TransformInfoInterface = transformInfoInterface
    return this
}


func (this *StoreThemeCategoryFactory) GetRootInstance_1() CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getRootInstance()", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreThemeCategoryPropertiesFactory(this.TransformInfoInterface)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, categoryPropertiesFactoryInterface, 0)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of StoreThemeCategoryFactory.getRootInstance
func (this *StoreThemeCategoryFactory) GetRootInstance(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GetRootInstance_1()
        }
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            return this.GetRootInstance_2(Java_Arg[IAbPath](args[0]))
        }
    }
    panic("No Java overload of StoreThemeCategoryFactory.getRootInstance for the arguments")
}


func (this *StoreThemeCategoryFactory) GetRootInstance_2(categoryAbPath IAbPath) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getRootInstance(String)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var level int = CategoryUtil_GetPathLevel(categoryAbPath); _ = level
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreThemeCategoryPropertiesFactory(this.TransformInfoInterface, categoryAbPath)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetRootInstanceFromNode(node *Node) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getRootInstance(node)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewRootStoreThemeCategoryPropertiesFactory(this.TransformInfoInterface, node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            var categoryPath IAbPath = NewAbPath(CategoryUtil_GetNameFromNode(node), StringUtil_GetInstance().EMPTY_STRING); _ = categoryPath
            var level int = CategoryUtil_GetPathLevel(categoryPath); _ = level
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, node, categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetInstance_1(categoryPath string, level int) CategoryInterface {
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
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryPath)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of StoreThemeCategoryFactory.getInstance
func (this *StoreThemeCategoryFactory) GetInstance(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.GetInstance_1(Java_Arg[string](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetInstance_2(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 2 && Java_IsArg[*Node](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.GetInstance_3(Java_Arg[*Node](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[*Node](args[2], pass) {
            return this.GetInstance_4(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[*Node](args[2]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return this.GetInstance_5(Java_Arg[*HashMap](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return this.GetInstance_6(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[*HashMap](args[2]))
        }
    }
    panic("No Java overload of StoreThemeCategoryFactory.getInstance for the arguments")
}


func (this *StoreThemeCategoryFactory) GetInstance_2(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPath string) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(CategoryInterface, CategoryInterface , String, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryPath)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetInstance_3(node *Node, level int) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(Node, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetInstance_4(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, node *Node) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(CategoryInterface, CategoryInterface, Node, int)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetInstance_5(categoryPropertiesHashMap *HashMap, level int) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(HashMap)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryPropertiesHashMap)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, categoryPropertiesFactoryInterface, level)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *StoreThemeCategoryFactory) GetInstance_6(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesHashMap *HashMap) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().ENTITYFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(CategoryInterface, CategoryInterface, HashMap)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryPropertiesHashMap)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewStoreThemeCategory(this.TransformInfoInterface, rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

