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

type CategoryHierarchy struct {
    RootCategoryInterface CategoryInterface
    ParentCategoryInterface CategoryInterface
    Level int
}

// Implements CategoryHierarchyInterface


func (this *CategoryHierarchy) initFields() {
    
    
    
}




// Java overloads of new CategoryHierarchy
func NewCategoryHierarchy(args ...any) *CategoryHierarchy {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) {
            return NewCategoryHierarchy_1(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return NewCategoryHierarchy_2(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[int](args[2]))
        }
        if len(args) == 3 && Java_IsArg[CategoryInterface](args[0], pass) && Java_IsArg[CategoryInterface](args[1], pass) && Java_IsArg[*HashMap](args[2], pass) {
            return NewCategoryHierarchy_3(Java_Arg[CategoryInterface](args[0]), Java_Arg[CategoryInterface](args[1]), Java_Arg[*HashMap](args[2]))
        }
    }
    panic("No Java overload of new CategoryHierarchy for the arguments")
}
func NewCategoryHierarchy_1(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface) *CategoryHierarchy {
    this := &CategoryHierarchy{}
    this.initFields()
        this.Level = 0
        this.SetRoot(rootCategoryInterface)
        this.SetParent(parentCategoryInterface)
    return this
}


func NewCategoryHierarchy_2(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, level int) *CategoryHierarchy {
    this := &CategoryHierarchy{}
    this.initFields()
        this.Level = level
        this.SetRoot(rootCategoryInterface)
        this.SetParent(parentCategoryInterface)
    return this
}


func NewCategoryHierarchy_3(rootCategoryInterface CategoryInterface, parentCategoryInterface CategoryInterface, categoryPropertiesHashMap *HashMap) *CategoryHierarchy {
    this := &CategoryHierarchy{}
    this.initFields()
        var levelString string = StringUtil_GetInstance().GetNonNull(any(categoryPropertiesHashMap.Get(CategoryData_GetInstance().LEVEL)).(string)); _ = levelString
        this.Level = int(fmt.Sprint(levelString))
        this.SetRoot(rootCategoryInterface)
        this.SetParent(parentCategoryInterface)
    return this
}


func (this *CategoryHierarchy) GetLevel() int {
        return this.Level
}


func (this *CategoryHierarchy) GetLevelString() string {
        return fmt.Sprint(this.Level)
}


func (this *CategoryHierarchy) SetLevel(level int) {
        this.Level = level
}


func (this *CategoryHierarchy) SetParent(categoryInterface CategoryInterface) {
        this.ParentCategoryInterface = categoryInterface
}


func (this *CategoryHierarchy) SetRoot(categoryInterface CategoryInterface) {
        this.RootCategoryInterface = categoryInterface
}


func (this *CategoryHierarchy) GetRoot() CategoryInterface {
        return this.RootCategoryInterface
}


func (this *CategoryHierarchy) GetParent() CategoryInterface {
        return this.ParentCategoryInterface
}

