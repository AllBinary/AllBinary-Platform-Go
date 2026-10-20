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

type RootCategoryPropertiesAbstract struct {
    Category string
    FileName string
}

// Implements RootCategoryPropertiesInterface


func (this *RootCategoryPropertiesAbstract) initFields() {
    this.Category = CategoryData_GetInstance().ROOTCATEGORY
    this.FileName = this.Category + fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + fmt.Sprint(CategoryData_GetInstance().UNCRYPTED_EXTENSION)
}


func NewRootCategoryPropertiesAbstract() *RootCategoryPropertiesAbstract {
    this := &RootCategoryPropertiesAbstract{}
    this.initFields()
    return this
}


func (this *RootCategoryPropertiesAbstract) GetKey() any {
        return any(this.GetValue()).(any)
}


func (this *RootCategoryPropertiesAbstract) GetValue() string {
        return this.Category
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) SetValue(value string) {
        panic(errors.New(fmt.Sprint("No Value Allowed")))
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) SetRootFilePath_1(value IAbPath) {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of RootCategoryPropertiesAbstract.setRootFilePath
func (this *RootCategoryPropertiesAbstract) SetRootFilePath(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            this.SetRootFilePath_1(Java_Arg[IAbPath](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            this.SetRootFilePath_2(Java_Arg[string](args[0]))
            return
        }
    }
    panic("No Java overload of RootCategoryPropertiesAbstract.setRootFilePath for the arguments")
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) SetRootFilePath_2(value string) {
        panic(errors.New(fmt.Sprint("No Value Allowed")))
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) GetRootFilePath() IAbPath {
        panic(errors.New(fmt.Sprint("No Value Allowed")))
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) IsRealRoot() bool {
        panic(errors.New(fmt.Sprint("No A Real Root")))
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) IsRoot() bool {
        panic(errors.New(fmt.Sprint("Root but not implemented")))
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) GetWebAppPath(categoryHierarchyInterface CategoryHierarchyInterface) string {
        panic(errors.New(fmt.Sprint("Not Root")))
}


func (this *RootCategoryPropertiesAbstract) GetPath(categoryHierarchyInterface CategoryHierarchyInterface) IAbPath {
        return AbPath_CreateAbPath()
}


func (this *RootCategoryPropertiesAbstract) GetFileName() string {
        return this.FileName
}


func (this *RootCategoryPropertiesAbstract) IsValid() bool {
        return true
}


func (this *RootCategoryPropertiesAbstract) ToHashMap() *HashMap {
        var categoryHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = categoryHashMap
        categoryHashMap.Put(CategoryData_GetInstance().NAME, this.Category)
        return categoryHashMap
}


func (this *RootCategoryPropertiesAbstract) ToVector() IBasicArrayList {
        var categoryVector IBasicArrayList = NewBasicArrayListD(); _ = categoryVector
        categoryVector.Add(this.Category)
        return categoryVector
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) ToValidationInfoDoc() *Document {
        return nil
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *RootCategoryPropertiesAbstract) ValidationInfo() string {
        return nil
}

