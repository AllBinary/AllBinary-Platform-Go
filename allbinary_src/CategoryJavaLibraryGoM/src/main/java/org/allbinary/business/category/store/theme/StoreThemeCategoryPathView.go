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

type StoreThemeCategoryPathView struct {
    CategoryView
}

// Implements DomNodeInterface

// The Java class StoreThemeCategoryPathView has subclasses so its variables are this interface
type IStoreThemeCategoryPathView interface {
    ICategoryView
    StoreThemeCategoryPathView_() *StoreThemeCategoryPathView
    ToXmlNode(document *Document) *Node
}

func (this *StoreThemeCategoryPathView) StoreThemeCategoryPathView_() *StoreThemeCategoryPathView {
    return this
}



func NewStoreThemeCategoryPathView(categoryInterface StoreThemeCategoryInterface) *StoreThemeCategoryPathView {
    this := &StoreThemeCategoryPathView{}
        this.CategoryView = *NewCategoryView(any(categoryInterface).(StoreThemeCategoryInterface))
    return this
}


// throws Exception
func (this *StoreThemeCategoryPathView) ToXmlNode(document *Document) *Node {
        var node *Node = this.CategoryView.ToXmlNode(document); _ = node
        var storeThemeCategoryInterface StoreThemeCategoryInterface = any(this.GetCategoryInterface()).(StoreThemeCategoryInterface); _ = storeThemeCategoryInterface
        var themeArray []any = storeThemeCategoryInterface.GetThemes().ToArray(); _ = themeArray
        var size int = len(themeArray); _ = size
        for i := int(0); i < size; i++ {
            var themeValidation *ThemeValidation = any(themeArray[i]).(*ThemeValidation); _ = themeValidation
            var domNodeInterface DomNodeInterface = any(NewThemePropertiesView(any(themeValidation).(ThemeInterface))).(DomNodeInterface); _ = domNodeInterface
            node.AppendChild(domNodeInterface.ToXmlNode(document))
        }
        return node
}

