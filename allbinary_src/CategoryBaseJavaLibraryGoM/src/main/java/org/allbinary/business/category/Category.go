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

type Category struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CategoryHierarchyInterface CategoryHierarchyInterface
    CategoryPropertiesInterface CategoryPropertiesInterface
    ChildCategoryVector IBasicArrayList
    TypeVector IBasicArrayList
    PROPERTIES int
    CATEGORY int
}

// Implements CategoryInterface

// The Java class Category has subclasses so its variables are this interface
type ICategory interface {
    Category_() *Category
    GetKey() any
    GetRootFilePath() IAbPath
    GetFilePath() IAbPath
    GetWebAppPath() IAbPath
    GetPath() IAbPath
    AddChildProperty(categoryPropertiesInterface CategoryPropertiesInterface) bool
    AddChild(categoryInterface CategoryInterface) bool
    GetChildNodes() IBasicArrayList
    RemoveChild(categoryInterface CategoryInterface) bool
    Removal(removalVector IBasicArrayList)
    RemoveDuplicateChild(categoryInterface CategoryInterface) bool
    IsLeaf() bool
    GetProperties() CategoryPropertiesInterface
    SetProperties(categoryPropertiesInterface CategoryPropertiesInterface)
    GetHierarchy() CategoryHierarchyInterface
    SetHierarchy(categoryHierarchyInterface CategoryHierarchyInterface)
    IsValid() bool
    ToValidationInfoDoc() *Document
    ToValidationInfoNode(document *Document) *Node
    ValidationInfo() string
    ToHashMap() *HashMap
    ToVector() IBasicArrayList
    Log()
}

func (this *Category) Category_() *Category {
    return this
}

var Category_NOT_ROOT string = "This is not the root so it has no file path."

func (this *Category) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    this.ChildCategoryVector = NewBasicArrayListD()
    this.TypeVector = NewBasicArrayListD()
    this.PROPERTIES = int(1)
    this.CATEGORY = int(0)
}




// Java overloads of new Category
func NewCategory(args ...any) *Category {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[CategoryPropertiesFactoryInterface](args[0], pass) {
            return NewCategory_1(Java_Arg[CategoryPropertiesFactoryInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[CategoryPropertiesFactoryInterface](args[0], pass) && Java_IsArg[int](args[1], pass) {
            return NewCategory_2(Java_Arg[CategoryPropertiesFactoryInterface](args[0]), Java_Arg[int](args[1]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[CategoryPropertiesFactoryInterface](args[2], pass) {
            return NewCategory_3(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[CategoryPropertiesFactoryInterface](args[2]))
        }
    }
    panic("No Java overload of new Category for the arguments")
}
func NewCategory_1(categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface) *Category {
    this := &Category{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, "Category(CategoryPropertiesFactoryInterface)")
        }
        this.CategoryPropertiesInterface = categoryPropertiesFactoryInterface.GetInstance()
        this.CategoryHierarchyInterface = any(NewCategoryHierarchy(this, this)).(CategoryHierarchyInterface)
        this.Log()
    return this
}


func NewCategory_2(categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface, level int) *Category {
    this := &Category{}
    this.initFields()
        this.CategoryPropertiesInterface = categoryPropertiesFactoryInterface.GetInstance()
        this.CategoryHierarchyInterface = any(NewCategoryHierarchy(this, this, level)).(CategoryHierarchyInterface)
        this.Log()
    return this
}


func NewCategory_3(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface) *Category {
    this := &Category{}
    this.initFields()
        this.CategoryPropertiesInterface = any(categoryPropertiesFactoryInterface.GetInstance()).(CategoryPropertiesInterface)
        this.CategoryHierarchyInterface = any(NewCategoryHierarchy(rootCategoryInterface, parentCategoryInterface)).(CategoryHierarchyInterface)
        this.Log()
    return this
}


// throws Exception
func (this *Category) GetKey() any {
        return this.CategoryPropertiesInterface.GetKey()
}


// throws Exception
func (this *Category) GetRootFilePath() IAbPath {
        var rootAbPath IAbPath = NewAbPath(Category_NOT_ROOT, StringUtil_GetInstance().EMPTY_STRING); _ = rootAbPath
        var categoryInterface CategoryInterface = this.CategoryHierarchyInterface.GetRoot(); _ = categoryInterface
        if categoryInterface.GetProperties().IsRoot() {
            var rootCategoryPropertiesInterface RootCategoryPropertiesInterface = any(categoryInterface.GetProperties()).(RootCategoryPropertiesInterface); _ = rootCategoryPropertiesInterface
            rootAbPath = rootCategoryPropertiesInterface.GetRootFilePath()
        }
        return rootAbPath
}


// throws Exception
func (this *Category) GetFilePath() IAbPath {
        return NewAbPath(fmt.Sprint(this.GetRootFilePath()) + fmt.Sprint(this.GetProperties().GetPath(this.GetHierarchy())), this.GetProperties().GetFileName())
}


// throws Exception
func (this *Category) GetWebAppPath() IAbPath {
        var categoryInterface CategoryInterface = this.CategoryHierarchyInterface.GetRoot(); _ = categoryInterface
        return categoryInterface.GetProperties().GetWebAppPath()
}


// throws Exception
func (this *Category) GetPath() IAbPath {
        return this.CategoryPropertiesInterface.GetPath(this.CategoryHierarchyInterface)
}


// synchronized
func (this *Category) AddChildProperty(categoryPropertiesInterface CategoryPropertiesInterface) bool {
        this.TypeVector.Add(this.PROPERTIES)
        return this.ChildCategoryVector.Add(categoryPropertiesInterface)
}


// synchronized
func (this *Category) AddChild(categoryInterface CategoryInterface) bool {
        this.RemoveDuplicateChild(categoryInterface)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF("adding", this, "addChild")
        }
        var childCategoryHierarchyInterface CategoryHierarchyInterface = categoryInterface.GetHierarchy(); _ = childCategoryHierarchyInterface
        childCategoryHierarchyInterface.SetParent(this)
        childCategoryHierarchyInterface.SetRoot(this.GetHierarchy().GetRoot())
        categoryInterface.SetHierarchy(childCategoryHierarchyInterface)
        this.TypeVector.Add(this.CATEGORY)
        return this.ChildCategoryVector.Add(categoryInterface)
}


func (this *Category) GetChildNodes() IBasicArrayList {
        return this.ChildCategoryVector
}


// synchronized
func (this *Category) RemoveChild(categoryInterface CategoryInterface) bool {
        var removalVector IBasicArrayList = NewBasicArrayListD(); _ = removalVector
        var bool_return bool = false; _ = bool_return
        var size int = this.ChildCategoryVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            if this.CATEGORY == this.TypeVector.Get(index) {
                var anyType any = this.ChildCategoryVector.Get(index); _ = anyType
                var currentCategoryInterface CategoryInterface = any(anyType).(CategoryInterface); _ = currentCategoryInterface
                if strings.Compare(currentCategoryInterface.GetProperties().GetValue(), categoryInterface.GetProperties().GetValue()) == 0 {
                    removalVector.Add(currentCategoryInterface)
                    bool_return = true
                }
            }
        }
        this.Removal(removalVector)
        return bool_return
}


func (this *Category) Removal(removalVector IBasicArrayList) {
        var removalSize int = removalVector.Size(); _ = removalSize
        for index := int(0); index < removalSize; index++ {
            var anyType any = removalVector.Get(index); _ = anyType
            var objectIndex int = this.ChildCategoryVector.IndexOf(anyType); _ = objectIndex
            this.TypeVector.Remove(objectIndex)
            this.ChildCategoryVector.Remove(objectIndex)
        }
}


// synchronized
func (this *Category) RemoveDuplicateChild(categoryInterface CategoryInterface) bool {
        var removalVector IBasicArrayList = NewBasicArrayListD(); _ = removalVector
        var bool_return bool = false; _ = bool_return
        var size int = this.ChildCategoryVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = this.ChildCategoryVector.Get(index); _ = anyType
            if this.PROPERTIES == this.TypeVector.Get(index) {
                var categoryPropertiesInterface CategoryPropertiesInterface = any(anyType).(CategoryPropertiesInterface); _ = categoryPropertiesInterface
                if strings.Compare(categoryInterface.GetProperties().GetValue(), categoryPropertiesInterface.GetValue()) == 0 {
                    removalVector.Add(categoryPropertiesInterface)
                    bool_return = true
                    break
                }
            } else if this.CATEGORY == this.TypeVector.Get(index) {
                var existingCategoryInterface CategoryInterface = any(anyType).(CategoryInterface); _ = existingCategoryInterface
                var categoryPropertiesInterface CategoryPropertiesInterface = existingCategoryInterface.GetProperties(); _ = categoryPropertiesInterface
                if strings.Compare(categoryInterface.GetProperties().GetValue(), categoryPropertiesInterface.GetValue()) == 0 {
                    removalVector.Add(existingCategoryInterface)
                    bool_return = true
                    break
                }
            }
        }
        this.Removal(removalVector)
        return bool_return
}


// throws Exception
func (this *Category) IsLeaf() bool {
        if this.ChildCategoryVector == nil {
            panic(errors.New(fmt.Sprint("Category Error")))
        }
        if this.ChildCategoryVector.Size() == 0 {
            return true
        } else {
            return false
        }
}


func (this *Category) GetProperties() CategoryPropertiesInterface {
        return this.CategoryPropertiesInterface
}


func (this *Category) SetProperties(categoryPropertiesInterface CategoryPropertiesInterface) {
        this.CategoryPropertiesInterface = categoryPropertiesInterface
}


func (this *Category) GetHierarchy() CategoryHierarchyInterface {
        return this.CategoryHierarchyInterface
}


func (this *Category) SetHierarchy(categoryHierarchyInterface CategoryHierarchyInterface) {
        this.CategoryHierarchyInterface = categoryHierarchyInterface
}


// throws Exception
// synchronized
func (this *Category) IsValid() bool {
        if !this.CategoryPropertiesInterface.IsValid() {
            return false
        }
        var size int = this.ChildCategoryVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = this.ChildCategoryVector.Get(index); _ = anyType
            if this.PROPERTIES == this.TypeVector.Get(index) {
                var categoryPropertiesInterface CategoryPropertiesInterface = any(anyType).(CategoryPropertiesInterface); _ = categoryPropertiesInterface
                if !categoryPropertiesInterface.IsValid() {
                    return false
                }
            } else if this.CATEGORY == this.TypeVector.Get(index) {
                var categoryInterface CategoryInterface = any(anyType).(CategoryInterface); _ = categoryInterface
                if !categoryInterface.IsValid() {
                    return false
                }
            }
        }
        return true
}


// throws Exception
func (this *Category) ToValidationInfoDoc() *Document {
        return nil
}


// throws Exception
func (this *Category) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *Category) ValidationInfo() string {
        return nil
}


// throws Exception
func (this *Category) ToHashMap() *HashMap {
        var categoryHashMap *HashMap = this.CategoryPropertiesInterface.ToHashMap(); _ = categoryHashMap
        return categoryHashMap
}


// throws Exception
func (this *Category) ToVector() IBasicArrayList {
        var categoryVector IBasicArrayList = this.CategoryPropertiesInterface.ToVector(); _ = categoryVector
        categoryVector.Add(this.ChildCategoryVector)
        return categoryVector
}


// throws Exception
func (this *Category) Log() {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Category Name: ")
            stringBuffer.Append(this.CategoryPropertiesInterface.GetValue())
            stringBuffer.Append("\nPath = ")
            stringBuffer.Append(fmt.Sprint(this.GetPath()))
            stringBuffer.Append("\nFile Path: ")
            stringBuffer.Append(fmt.Sprint(this.GetFilePath()))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "log()")
        }
}

