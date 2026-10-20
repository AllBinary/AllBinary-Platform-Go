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

type StoreThemeCategoriesView struct {
    CategoryInterface StoreThemeCategoryInterface
}

// Implements DomNodeInterface


func (this *StoreThemeCategoriesView) initFields() {
    
}


func NewStoreThemeCategoriesView(categoryInterface StoreThemeCategoryInterface) *StoreThemeCategoriesView {
    this := &StoreThemeCategoriesView{}
    this.initFields()
        this.CategoryInterface = categoryInterface
    return this
}


// throws Exception
func (this *StoreThemeCategoriesView) ToXmlNode(document *Document) *Node {
        var node *Node = NewStoreThemeCategoryView(this.CategoryInterface).ToXmlNode(document); _ = node
        var childArray []any = this.CategoryInterface.GetChildNodes().ToArray(); _ = childArray
        var size int = len(childArray); _ = size
        for i := int(0); i < size; i++ {
            var childCategoryInterface CategoryInterface = any(childArray[i]).(CategoryInterface); _ = childCategoryInterface
            var childCategoryNode *Node = NewStoreThemeCategoriesView(any(childCategoryInterface).(StoreThemeCategoryInterface)).ToXmlNode(document); _ = childCategoryNode
            node.AppendChild(childCategoryNode)
        }
        return node
}

