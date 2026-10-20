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

type CategoryView struct {
    LogUtil *LogUtil
    CategoryInterface CategoryInterface
}

// Implements DomNodeInterface

// The Java class CategoryView has subclasses so its variables are this interface
type ICategoryView interface {
    CategoryView_() *CategoryView
    GetCategoryInterface() CategoryInterface
    ToXmlNode(document *Document) *Node
}

func (this *CategoryView) CategoryView_() *CategoryView {
    return this
}


func (this *CategoryView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewCategoryView(categoryInterface CategoryInterface) *CategoryView {
    this := &CategoryView{}
    this.initFields()
        this.CategoryInterface = categoryInterface
    return this
}


func (this *CategoryView) GetCategoryInterface() CategoryInterface {
        return this.CategoryInterface
}


// throws Exception
func (this *CategoryView) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(CategoryData_GetInstance().NAME); _ = node
        var childCategoryVector IBasicArrayList = this.CategoryInterface.GetChildNodes(); _ = childCategoryVector
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF("Number Of Children: " + fmt.Sprint(childCategoryVector.Size()), this, "toXmlNode")
        }
        var size int = childCategoryVector.Size(); _ = size
        for i := int(0); i < size; i++ {
            var anyType any = childCategoryVector.Get(i); _ = anyType
            var categoryPropertiesInterface CategoryPropertiesInterface = nil; _ = categoryPropertiesInterface
            if InterfaceUtil_IsImplemented("CategoryPropertiesInterface", anyType) {
                categoryPropertiesInterface = any(anyType).(CategoryPropertiesInterface)
            } else if InterfaceUtil_IsImplemented("CategoryInterface", anyType) {
                var categoryInterface CategoryInterface = any(anyType).(CategoryInterface); _ = categoryInterface
                categoryPropertiesInterface = categoryInterface.GetProperties()
            } else {
                var isImpl string = InterfaceUtil_ViewAll(reflect.TypeOf(anyType), CommonSeps_GetInstance().NEW_LINE); _ = isImpl
                panic(errors.New(fmt.Sprint("CategoryView found unknown object.\n" + "Show Testing:" + isImpl)))
            }
            var childCategoryNode *Node = NewCategoryPropertiesView(categoryPropertiesInterface).ToXmlNode(document); _ = childCategoryNode
            node.AppendChild(childCategoryNode)
        }
        return node
}

