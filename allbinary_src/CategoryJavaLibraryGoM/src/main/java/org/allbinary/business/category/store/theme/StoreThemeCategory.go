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

type StoreThemeCategory struct {
    Category
    TransformInfoInterface TransformInfoInterface
    ThemeVector IBasicArrayList
}

// Implements StoreThemeCategoryInterface


func (this *StoreThemeCategory) initFields() {
    
    
}




// Java overloads of new StoreThemeCategory
func NewStoreThemeCategory(args ...any) *StoreThemeCategory {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[CategoryPropertiesFactoryInterface](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return NewStoreThemeCategory_1(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[CategoryPropertiesFactoryInterface](args[1]), Java_Arg[int](args[2]))
        }
        if len(args) == 4 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[CategoryInterface](args[2], pass) && Java_IsArg[CategoryPropertiesFactoryInterface](args[3], pass) {
            return NewStoreThemeCategory_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[CategoryInterface](args[2]), Java_Arg[CategoryPropertiesFactoryInterface](args[3]))
        }
        if len(args) == 4 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) && Java_IsArg[CategoryPropertiesFactoryInterface](args[2], pass) && Java_IsArg[int](args[3], pass) {
            return NewStoreThemeCategory_3(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Node](args[1]), Java_Arg[CategoryPropertiesFactoryInterface](args[2]), Java_Arg[int](args[3]))
        }
        if len(args) == 5 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) && Java_IsArg[CategoryInterface](args[2], pass) && Java_IsArg[CategoryInterface](args[3], pass) && Java_IsArg[CategoryPropertiesFactoryInterface](args[4], pass) {
            return NewStoreThemeCategory_4(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Node](args[1]), Java_Arg[CategoryInterface](args[2]), Java_Arg[CategoryInterface](args[3]), Java_Arg[CategoryPropertiesFactoryInterface](args[4]))
        }
    }
    panic("No Java overload of new StoreThemeCategory for the arguments")
}
func NewStoreThemeCategory_1(transformInfoInterface TransformInfoInterface, categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface, level int) *StoreThemeCategory {
    this := &StoreThemeCategory{}
        this.Category = *NewCategory(categoryPropertiesFactoryInterface, level)
        this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.ThemeVector = NewBasicArrayListD()
    return this
}


func NewStoreThemeCategory_2(transformInfoInterface TransformInfoInterface, rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface) *StoreThemeCategory {
    this := &StoreThemeCategory{}
        this.Category = *NewCategory(rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)
        this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.ThemeVector = NewBasicArrayListD()
    return this
}


func NewStoreThemeCategory_3(transformInfoInterface TransformInfoInterface, node *Node, categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface, level int) *StoreThemeCategory {
    this := &StoreThemeCategory{}
        this.Category = *NewCategory(categoryPropertiesFactoryInterface, level)
        this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.ThemeVector = NewBasicArrayListD()
        this.AddThemes(node)
    return this
}


func NewStoreThemeCategory_4(transformInfoInterface TransformInfoInterface, node *Node, rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesFactoryInterface CategoryPropertiesFactoryInterface) *StoreThemeCategory {
    this := &StoreThemeCategory{}
        this.Category = *NewCategory(rootCategoryInterface, parentCategoryInterface, categoryPropertiesFactoryInterface)
        this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        this.ThemeVector = NewBasicArrayListD()
        this.AddThemes(node)
    return this
}


func (this *StoreThemeCategory) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


// throws Exception
func (this *StoreThemeCategory) AddThemes(node *Node) {
        var themesNode *Node = DomSearchHelper_GetNode(ThemesData_GetInstance().NAME, node.GetChildNodes()); _ = themesNode
        var nodeList *NodeList = themesNode.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var themeNode *Node = nodeList.Item(index); _ = themeNode
            if themeNode.GetNodeName().CompareTo(ThemeData_GetInstance().NAME) == 0 {
                this.ThemeVector.Add(NewThemeValidation(this, themeNode))
            }
        }
}


func (this *StoreThemeCategory) AddTheme(themeValidation *ThemeValidation) {
        this.ThemeVector.Add(themeValidation)
}


func (this *StoreThemeCategory) GetThemes() IBasicArrayList {
        return this.ThemeVector
}


// throws Exception
func (this *StoreThemeCategory) IsValid() bool {
        if !this.Category.IsValid() {
            return false
        }
        return true
}


// throws Exception
func (this *StoreThemeCategory) ToHashMap() *HashMap {
        var categoryHashMap *HashMap = this.Category.ToHashMap(); _ = categoryHashMap
        categoryHashMap.Put(ThemesData_GetInstance().NAME, this.ThemeVector)
        return categoryHashMap
}


// throws Exception
func (this *StoreThemeCategory) ToVector() IBasicArrayList {
        var categoryVector IBasicArrayList = this.Category.ToVector(); _ = categoryVector
        return categoryVector
}

