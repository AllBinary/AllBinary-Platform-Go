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

type CategoryLoader struct {
    CategoryModifierTree
    LogUtil *LogUtil
}

// Implements CategoryLoaderInterface


func (this *CategoryLoader) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewCategoryLoader(categoryFactoryInterface CategoryFactoryInterface) *CategoryLoader {
    this := &CategoryLoader{}
        this.CategoryModifierTree = *NewCategoryModifierTree(categoryFactoryInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "CategoryLoader(CategoryFactoryInterface categoryFactoryInterface)")
        }
    return this
}


// synchronized
func (this *CategoryLoader) GetDoc(categoryInterface CategoryInterface) *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 *Document) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e2 := Java_ToError(recovered)
                            _ = e2
                    this.LogUtil.Put("Could Not Log Category", this, "getDoc", e)
                        }
                    }()
                    categoryInterface.Log()
                    return
                }(); tryControl2 == 1 {
                    tryControl1, tryResult1 = 1, tryResult2; return
                }
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getDoc", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if categoryInterface.GetProperties().GetPath(categoryInterface.GetHierarchy()) != nil {
                if categoryInterface.GetHierarchy().GetLevel() == 1 {
                    var rootAbFile *AbFile = AbFile_CreateAbFileFromAbPath(categoryInterface.GetRootFilePath()); _ = rootAbFile
                    if !rootAbFile.IsDirectory() {
                        rootAbFile.Mkdirs()
                        if !rootAbFile.IsDirectory() {
                            panic(errors.New(fmt.Sprint("Could Not Create Directory")))
                        }
                        this.Save(categoryInterface)
                    }
                }
                var fileAbPath IAbPath = categoryInterface.GetFilePath(); _ = fileAbPath
                var cryptFileReader *CryptFileReader = NewCryptFileReader(CategoryData_GetInstance().UNCRYPTED_EXTENSION, CategoryData_GetInstance().ENCRYPTED_EXTENSION); _ = cryptFileReader
                var document *Document = DomDocumentHelper_Create(cryptFileReader.Get(fileAbPath)); _ = document
                tryControl1, tryResult1 = 1, document; return
            } else {
                panic(errors.New(fmt.Sprint("Null Category")))
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *CategoryLoader) Get_1(categoryInterface CategoryInterface) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var document *Document = this.GetDoc(categoryInterface); _ = document
            var categoryNode *Node = DomSearchHelper_GetNode(CategoryData_GetInstance().NAME, document.GetChildNodes()); _ = categoryNode
            tryControl1, tryResult1 = 1, this.AddProperties(categoryInterface, categoryNode.GetChildNodes()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of CategoryLoader.get
func (this *CategoryLoader) Get(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[CategoryInterface](args[0], pass) {
            return this.Get_1(Java_Arg[CategoryInterface](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return this.Get_2(Java_Arg[*Document](args[0]))
        }
    }
    panic("No Java overload of CategoryLoader.get for the arguments")
}


// synchronized
func (this *CategoryLoader) Get_2(document *Document) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryNode *Node = DomSearchHelper_GetNode(CategoryData_GetInstance().NAME, document.GetChildNodes()); _ = categoryNode
            var loadedCategoryInterface CategoryInterface = this.CategoryFactoryInterface.GetRootInstanceFromNode(categoryNode); _ = loadedCategoryInterface
            tryControl1, tryResult1 = 1, this.AddProperties(loadedCategoryInterface, categoryNode.GetChildNodes()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *CategoryLoader) GetAll_1(categoryInterface CategoryInterface) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var document *Document = this.GetDoc(categoryInterface); _ = document
            tryControl1, tryResult1 = 1, this.GetAll(document); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of CategoryLoader.getAll
func (this *CategoryLoader) GetAll(args ...any) CategoryInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[CategoryInterface](args[0], pass) {
            return this.GetAll_1(Java_Arg[CategoryInterface](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return this.GetAll_2(Java_Arg[*Document](args[0]))
        }
        if len(args) == 2 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[*NodeList](args[1], pass) {
            return this.GetAll_3(Java_Arg[CategoryInterface](args[0]), Java_Arg[*NodeList](args[1]))
        }
    }
    panic("No Java overload of CategoryLoader.getAll for the arguments")
}


// synchronized
func (this *CategoryLoader) GetAll_2(document *Document) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var categoryNode *Node = DomSearchHelper_GetNode(CategoryData_GetInstance().NAME, document.GetChildNodes()); _ = categoryNode
            var loadedCategoryInterface CategoryInterface = this.CategoryFactoryInterface.GetRootInstanceFromNode(categoryNode); _ = loadedCategoryInterface
            tryControl1, tryResult1 = 1, this.AddProperties(loadedCategoryInterface, categoryNode.GetChildNodes()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *CategoryLoader) GetAll_3(parentCategoryInterface CategoryInterface, categoryNodeList *NodeList) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if categoryNodeList != nil {
                for index := int(0); index < categoryNodeList.GetLength(); index++ {
                    var categoryNode *Node = categoryNodeList.Item(index); _ = categoryNode
                    if categoryNode != nil && categoryNode.GetNodeName().CompareTo(CategoryData_GetInstance().NAME) == 0 {
                        var categoryInterface CategoryInterface = any(this.CategoryFactoryInterface.GetInstance(parentCategoryInterface.GetHierarchy().GetRoot(), parentCategoryInterface, categoryNode)).(CategoryInterface); _ = categoryInterface
                        parentCategoryInterface.AddChild(categoryInterface)
                    }
                }
            }
            tryControl1, tryResult1 = 1, parentCategoryInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// synchronized
func (this *CategoryLoader) AddProperties(loadedCategoryInterface CategoryInterface, categoryNodeList *NodeList) CategoryInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 CategoryInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.GET, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            if categoryNodeList != nil {
                for index := int(0); index < categoryNodeList.GetLength(); index++ {
                    var categoryNode *Node = categoryNodeList.Item(index); _ = categoryNode
                    if categoryNode != nil && categoryNode.GetNodeName().CompareTo(CategoryData_GetInstance().NAME) == 0 {
                        var categoryPropertiesInterface CategoryPropertiesInterface = any(NewCategoryProperties(categoryNode)).(CategoryPropertiesInterface); _ = categoryPropertiesInterface
                        loadedCategoryInterface.AddChildProperty(categoryPropertiesInterface)
                    }
                }
            }
            tryControl1, tryResult1 = 1, loadedCategoryInterface; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

