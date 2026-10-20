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

type CategoryAbstractFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements CategoryFactoryInterface

// The Java class CategoryAbstractFactory has subclasses so its variables are this interface
type ICategoryAbstractFactory interface {
    CategoryAbstractFactory_() *CategoryAbstractFactory
    GetRootInstance(args ...any) CategoryInterface
    GetRootInstanceFromNode(node *Node) CategoryInterface
    GetInstance(args ...any) CategoryInterface
}

func (this *CategoryAbstractFactory) CategoryAbstractFactory_() *CategoryAbstractFactory {
    return this
}


func (this *CategoryAbstractFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewCategoryAbstractFactory() *CategoryAbstractFactory {
    this := &CategoryAbstractFactory{}
    this.initFields()
    return this
}


func (this *CategoryAbstractFactory) GetRootInstance_1() CategoryInterface {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of CategoryAbstractFactory.getRootInstance
func (this *CategoryAbstractFactory) GetRootInstance(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.GetRootInstance_1()
        }
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            return this.GetRootInstance_2(Java_Arg[IAbPath](args[0]))
        }
    }
    panic("No Java overload of CategoryAbstractFactory.getRootInstance for the arguments")
}


func (this *CategoryAbstractFactory) GetRootInstance_2(categoryAbPath IAbPath) CategoryInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *CategoryAbstractFactory) GetRootInstanceFromNode(node *Node) CategoryInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *CategoryAbstractFactory) GetInstance_1(categoryName string) CategoryInterface {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "getInstance(String)")
            }
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryName)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of CategoryAbstractFactory.getInstance
func (this *CategoryAbstractFactory) GetInstance(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetInstance_1(Java_Arg[string](args[0]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetInstance_2(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return this.GetInstance_3(Java_Arg[*Node](args[0]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[*Node](args[2], pass) {
            return this.GetInstance_4(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[*Node](args[2]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return this.GetInstance_5(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[*HashMap](args[2]))
        }
    }
    panic("No Java overload of CategoryAbstractFactory.getInstance for the arguments")
}


func (this *CategoryAbstractFactory) GetInstance_2(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryName string) CategoryInterface {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "getInstance(CategoryInterface, CategoryInterface, String)")
            }
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryName)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryAbstractFactory) GetInstance_3(node *Node) CategoryInterface {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "getInstance(Node)")
            }
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryAbstractFactory) GetInstance_4(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, node *Node) CategoryInterface {
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
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "getInstance(CategoryInterface, CategoryInterface, Node)")
            }
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(node)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryAbstractFactory) GetInstance_5(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesHashMap *HashMap) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstance(CategoryInterface, CategoryInterface, HashMap)", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "getInstance(CategoryInterface, CategoryInterface, HashMap)")
            }
            var categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface = any(NewCategoryPropertiesFactory(categoryPropertiesHashMap)).(CategoryPropertiesFactoryInterface); _ = categoryPropertiesFactoryInterface
            tryControl1, tryResult1 = 1, any(NewCategory(rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)).(CategoryInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

