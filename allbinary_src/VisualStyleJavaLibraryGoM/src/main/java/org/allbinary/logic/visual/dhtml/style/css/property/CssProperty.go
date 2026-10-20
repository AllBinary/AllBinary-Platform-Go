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

type CssProperty struct {
    LogUtil *LogUtil
    PropertyName string
    CssPropertyValues *CssPropertyValues
}

// Implements DomNodeInterface


func (this *CssProperty) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewCssProperty(name string) *CssProperty {
    this := &CssProperty{}
    this.initFields()
        this.PropertyName = name
        this.CssPropertyValues = NewCssPropertyValues()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewCssProperty(node *Node) *CssProperty {
    this := &CssProperty{}
    this.initFields()
        var propertyNode *Node = DomSearchHelper_GetNode(DomData_VALUE, node.GetChildNodes()); _ = propertyNode
        this.PropertyName = DomNodeHelper_GetTextNodeValue(propertyNode)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STYLE) {
            this.LogUtil.PutF("Name: " + this.GetName(), this, "CssProperty()")
        }
        var propertyValuesNode *Node = DomSearchHelper_GetNode(CssPropertyData_GetInstance().VALUES, node.GetChildNodes()); _ = propertyValuesNode
        this.CssPropertyValues = NewCssPropertyValues(propertyValuesNode)
    return this
}


func (this *CssProperty) GetName() string {
        return this.PropertyName
}


func (this *CssProperty) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(CssPropertyData_GetInstance().NAME, this.PropertyName)
        return hashMap
}


// throws Exception
func (this *CssProperty) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("CssView HashMap: " + fmt.Sprint(hashMap), this, "toXmlNode()")
        }
        var node *Node = ModDomHelper_CreateNameValueNodes(document, CssPropertyData_GetInstance().NAME, this.PropertyName); _ = node
        node.AppendChild(this.CssPropertyValues.ToXmlNode(document))
        return node
}

