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

type CategoryPropertiesView struct {
    CategoryPropertiesInterface CategoryPropertiesInterface
}

// Implements DomNodeInterface

// The Java class CategoryPropertiesView has subclasses so its variables are this interface
type ICategoryPropertiesView interface {
    CategoryPropertiesView_() *CategoryPropertiesView
    ToXmlNode(document *Document) *Node
}

func (this *CategoryPropertiesView) CategoryPropertiesView_() *CategoryPropertiesView {
    return this
}


func (this *CategoryPropertiesView) initFields() {
    
}


func NewCategoryPropertiesView(categoryPropertiesInterface CategoryPropertiesInterface) *CategoryPropertiesView {
    this := &CategoryPropertiesView{}
    this.initFields()
        this.CategoryPropertiesInterface = categoryPropertiesInterface
    return this
}


// throws Exception
func (this *CategoryPropertiesView) ToXmlNode(document *Document) *Node {
        var categoryNode *Node = document.CreateElement(CategoryData_GetInstance().NAME); _ = categoryNode
        var newCategoryNode *Element = any(categoryNode).(*Element); _ = newCategoryNode
        newCategoryNode.SetAttribute(CategoryData_GetInstance().LABEL, this.CategoryPropertiesInterface.GetValue())
        return categoryNode
}

