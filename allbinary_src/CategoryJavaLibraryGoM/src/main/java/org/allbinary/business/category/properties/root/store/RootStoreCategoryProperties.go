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

type RootStoreCategoryProperties struct {
    LogUtil *LogUtil
    PathData *AbPathData
    FileAbPath IAbPath
    AbPath IAbPath
    Category string
    IsRealRoot bool
    WebAppAbPath IAbPath
    TransformInfoInterface TransformInfoInterface
}

// Implements RootCategoryPropertiesInterface, CategoryPropertiesInterface

// The Java class RootStoreCategoryProperties has subclasses so its variables are this interface
type IRootStoreCategoryProperties interface {
    RootStoreCategoryProperties_() *RootStoreCategoryProperties
    InitPath()
    IsRealRoot() bool
    IsRoot() bool
    GetKey() any
    SetPath(path IAbPath)
    SetRootFilePath(value IAbPath)
    GetRootFilePath() IAbPath
    GetValue() string
    GetWebAppPath() IAbPath
    GetPath(categoryHierarchyInterface CategoryHierarchyInterface) IAbPath
    GetFileName() string
    SetValue(value string)
    IsValid() bool
    ToHashMap() *HashMap
    ToVector() IBasicArrayList
    ToValidationInfoDoc() *Document
    ToValidationInfoNode(document *Document) *Node
    ValidationInfo() string
    Log()
}

func (this *RootStoreCategoryProperties) RootStoreCategoryProperties_() *RootStoreCategoryProperties {
    return this
}


func (this *RootStoreCategoryProperties) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PathData = AbPathData_GetInstance()
    
    
    
    
    
    
}




// Java overloads of new RootStoreCategoryProperties
func NewRootStoreCategoryProperties(args ...any) *RootStoreCategoryProperties {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewRootStoreCategoryProperties_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) {
            return NewRootStoreCategoryProperties_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[IAbPath](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewRootStoreCategoryProperties_3(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewRootStoreCategoryProperties_4(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of new RootStoreCategoryProperties for the arguments")
}
func NewRootStoreCategoryProperties_1(transformInfoInterface TransformInfoInterface) *RootStoreCategoryProperties {
    this := &RootStoreCategoryProperties{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.AbPath = AbPath_CreateAbPath()
        this.Category = CategoryData_GetInstance().ROOTCATEGORY
        this.IsRealRoot = true
        this.InitPath()
        this.Log()
    return this
}


func NewRootStoreCategoryProperties_2(transformInfoInterface TransformInfoInterface, categoryAbPath IAbPath) *RootStoreCategoryProperties {
    this := &RootStoreCategoryProperties{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.AbPath = categoryAbPath
        this.Category = this.PathData.GetNameFromPath(fmt.Sprint(categoryAbPath))
        if StringValidationUtil_GetInstance().IsEmpty(this.Category) {
            this.IsRealRoot = true
            this.Category = CategoryData_GetInstance().ROOTCATEGORY
        }
        this.InitPath()
        this.Log()
    return this
}


func NewRootStoreCategoryProperties_3(transformInfoInterface TransformInfoInterface, node *Node) *RootStoreCategoryProperties {
    this := &RootStoreCategoryProperties{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        var categoryPath string = CategoryUtil_GetNameFromNode(node); _ = categoryPath
        this.AbPath = NewAbPath(categoryPath, StringUtil_GetInstance().EMPTY_STRING)
        this.Category = this.PathData.GetNameFromPath(categoryPath)
        if StringValidationUtil_GetInstance().IsEmpty(this.Category) {
            this.IsRealRoot = true
            this.Category = CategoryData_GetInstance().ROOTCATEGORY
        }
        this.InitPath()
        this.Log()
    return this
}


func NewRootStoreCategoryProperties_4(transformInfoInterface TransformInfoInterface, categoryPropertiesHashMap *HashMap) *RootStoreCategoryProperties {
    this := &RootStoreCategoryProperties{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        var categoryPath string = any(categoryPropertiesHashMap.Get(CategoryData_GetInstance().NAME)).(string); _ = categoryPath
        this.AbPath = NewAbPath(categoryPath, StringUtil_GetInstance().EMPTY_STRING)
        this.Category = this.PathData.GetNameFromPath(categoryPath)
        if StringValidationUtil_GetInstance().IsEmpty(this.Category) {
            this.IsRealRoot = true
            this.Category = CategoryData_GetInstance().ROOTCATEGORY
        }
        this.InitPath()
        this.Log()
    return this
}


// throws Exception
func (this *RootStoreCategoryProperties) InitPath() {
        var transformInfoHttpStoreInterface TransformInfoHttpInterface = any(this.TransformInfoInterface).(TransformInfoHttpInterface); _ = transformInfoHttpStoreInterface
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(transformInfoHttpStoreInterface.GetStoreName()); _ = storeFrontInterface
        var postPath string = storeFrontInterface.GetCurrentHostNamePath() + storeFrontInterface.GetCategoryPath(); _ = postPath
        var httpServletRequest *HttpServletRequest = any(transformInfoHttpStoreInterface.GetPageContext().GetRequest()).(*HttpServletRequest); _ = httpServletRequest
        this.WebAppAbPath = NewAbPath(fmt.Sprint(httpServletRequest.GetContextPath()) + postPath, StringUtil_GetInstance().EMPTY_STRING)
        this.SetRootFilePath(NewAbPath(fmt.Sprint(URLGLOBALS.GetMainPath()) + postPath, StringUtil_GetInstance().EMPTY_STRING))
}


func (this *RootStoreCategoryProperties) IsRealRoot() bool {
        return this.IsRealRoot
}


func (this *RootStoreCategoryProperties) IsRoot() bool {
        return true
}


func (this *RootStoreCategoryProperties) GetKey() any {
        return any(this.GetValue()).(any)
}


func (this *RootStoreCategoryProperties) SetPath(path IAbPath) {
        this.AbPath = path
}


func (this *RootStoreCategoryProperties) SetRootFilePath(value IAbPath) {
        this.FileAbPath = value
}


func (this *RootStoreCategoryProperties) GetRootFilePath() IAbPath {
        return this.FileAbPath
}


func (this *RootStoreCategoryProperties) GetValue() string {
        return this.Category
}


// throws Exception
func (this *RootStoreCategoryProperties) GetWebAppPath() IAbPath {
        return this.WebAppAbPath
}


func (this *RootStoreCategoryProperties) GetPath(categoryHierarchyInterface CategoryHierarchyInterface) IAbPath {
        return this.AbPath
}


func (this *RootStoreCategoryProperties) GetFileName() string {
        return this.GetValue() + AbPathData_GetInstance().EXTENSION_SEP + CategoryData_GetInstance().UNCRYPTED_EXTENSION
}


func (this *RootStoreCategoryProperties) SetValue(value string) {
        this.Category = value
}


func (this *RootStoreCategoryProperties) IsValid() bool {
        return true
}


func (this *RootStoreCategoryProperties) ToHashMap() *HashMap {
        var categoryHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = categoryHashMap
        categoryHashMap.Put(CategoryData_GetInstance().NAME, this.GetValue())
        return categoryHashMap
}


func (this *RootStoreCategoryProperties) ToVector() IBasicArrayList {
        var categoryVector IBasicArrayList = NewBasicArrayListD(); _ = categoryVector
        categoryVector.Add(this.GetValue())
        return categoryVector
}


// throws Exception
func (this *RootStoreCategoryProperties) ToValidationInfoDoc() *Document {
        return nil
}


// throws Exception
func (this *RootStoreCategoryProperties) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *RootStoreCategoryProperties) ValidationInfo() string {
        return nil
}


func (this *RootStoreCategoryProperties) Log() {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("filePath = ").Append(fmt.Sprint(this.FileAbPath)).Append("\npath = ").Append(fmt.Sprint(this.AbPath)).Append("\ncategory = ").Append(this.Category)), this, "log()")
        }
}

