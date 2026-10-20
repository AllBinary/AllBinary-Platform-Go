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

type CategoryProperties struct {
    LogUtil *LogUtil
    Category string
}

// Implements CategoryPropertiesInterface

var CategoryProperties_MAXDEPTH int = 100

func (this *CategoryProperties) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new CategoryProperties
func NewCategoryProperties(args ...any) *CategoryProperties {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewCategoryProperties_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCategoryProperties_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewCategoryProperties_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new CategoryProperties for the arguments")
}
func NewCategoryProperties_1(name string) *CategoryProperties {
    this := &CategoryProperties{}
    this.initFields()
        this.Category = name
    return this
}


func NewCategoryProperties_2(node *Node) *CategoryProperties {
    this := &CategoryProperties{}
    this.initFields()
        this.Category = CategoryUtil_GetNameFromNode(node)
    return this
}


func NewCategoryProperties_3(categoryPropertiesHashMap *HashMap) *CategoryProperties {
    this := &CategoryProperties{}
    this.initFields()
        this.Category = any(categoryPropertiesHashMap.Get(CategoryData_GetInstance().NAME)).(string)
    return this
}


func (this *CategoryProperties) IsRealRoot() bool {
        return false
}


func (this *CategoryProperties) IsRoot() bool {
        return false
}


func (this *CategoryProperties) GetKey() any {
        return any(this.GetValue()).(any)
}


func (this *CategoryProperties) GetValue() string {
        return this.Category
}


// throws Exception
func (this *CategoryProperties) GetPath(categoryHierarchyInterface CategoryHierarchyInterface) IAbPath {
        var abPathData *AbPathData = AbPathData_GetInstance(); _ = abPathData
        var pathStringBuffer *StringMaker = NewStringMaker(); _ = pathStringBuffer
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var nextParentCategoryInterface CategoryInterface = categoryHierarchyInterface.GetParent(); _ = nextParentCategoryInterface
        if this != nextParentCategoryInterface.GetProperties() {
            var depthIndex int = 0; _ = depthIndex
            for nextParentCategoryInterface != nil {
                if nextParentCategoryInterface.GetProperties().IsRealRoot() {
                    break
                }
                if nextParentCategoryInterface.GetProperties().IsRoot() {
                    list.AddAt(0, fmt.Sprint(nextParentCategoryInterface.GetPath()))
                } else {
                    pathStringBuffer.Delete(0, pathStringBuffer.Length())
                    list.AddAt(0, fmt.Sprint(pathStringBuffer.Append(abPathData.SEPARATOR).Append(nextParentCategoryInterface.GetProperties().GetValue())))
                }
                if nextParentCategoryInterface == nextParentCategoryInterface.GetHierarchy().GetParent() {
                    break
                }
                nextParentCategoryInterface = nextParentCategoryInterface.GetHierarchy().GetParent()
                depthIndex++
                if depthIndex > CategoryProperties_MAXDEPTH {
                    panic(errors.New(fmt.Sprint("Probably Major Error")))
                }
            }
        }
        pathStringBuffer.Delete(0, pathStringBuffer.Length())
        pathStringBuffer.Append(abPathData.SEPARATOR).Append(this.GetValue())
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            pathStringBuffer.Append(any(list.Get(index)).(string))
        }
        pathStringBuffer.Append(abPathData.SEPARATOR)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF("path = " + fmt.Sprint(pathStringBuffer), this, "getPath")
        }
        return NewAbPath(fmt.Sprint(pathStringBuffer), StringUtil_GetInstance().EMPTY_STRING)
}


func (this *CategoryProperties) GetFileName() string {
        return this.GetValue() + AbPathData_GetInstance().EXTENSION_SEP + CategoryData_GetInstance().UNCRYPTED_EXTENSION
}


// throws Exception
func (this *CategoryProperties) GetWebAppPath() IAbPath {
        panic(errors.New(fmt.Sprint("Not Root")))
}


func (this *CategoryProperties) SetValue(value string) {
        this.Category = value
}


func (this *CategoryProperties) IsValid() bool {
        var returnBoolean bool = false; _ = returnBoolean
        if StringValidationUtil_GetInstance().IsValidRequired(this.Category, 1, AbSqlData_MAXSTRING) {
            returnBoolean = true
        }
        return returnBoolean
}


func (this *CategoryProperties) ToHashMap() *HashMap {
        var categoryHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = categoryHashMap
        categoryHashMap.Put(CategoryData_GetInstance().NAME, this.Category)
        return categoryHashMap
}


func (this *CategoryProperties) ToVector() IBasicArrayList {
        var categoryVector IBasicArrayList = NewBasicArrayListD(); _ = categoryVector
        categoryVector.Add(this.Category)
        return categoryVector
}


// throws Exception
func (this *CategoryProperties) ToValidationInfoDoc() *Document {
        return nil
}


// throws Exception
func (this *CategoryProperties) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *CategoryProperties) ValidationInfo() string {
        return nil
}

