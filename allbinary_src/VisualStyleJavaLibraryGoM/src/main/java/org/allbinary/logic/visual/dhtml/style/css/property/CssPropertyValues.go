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

type CssPropertyValues struct {
    LogUtil *LogUtil
    PropertyValueVector IBasicArrayList
}

// Implements DomNodeInterface


func (this *CssPropertyValues) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new CssPropertyValues
func NewCssPropertyValues(args ...any) *CssPropertyValues {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewCssPropertyValues_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCssPropertyValues_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new CssPropertyValues for the arguments")
}
func NewCssPropertyValues_1() *CssPropertyValues {
    this := &CssPropertyValues{}
    this.initFields()
        this.PropertyValueVector = NewBasicArrayListD()
    return this
}


func NewCssPropertyValues_2(node *Node) *CssPropertyValues {
    this := &CssPropertyValues{}
    this.initFields()
        var indexPropertyValueHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = indexPropertyValueHashMap
        this.PropertyValueVector = NewBasicArrayListD()
        var cssPropertyNodeVector IBasicArrayList = DomSearchHelper_GetAllNodes(CssPropertyValueData_GetInstance().NAME, node.GetChildNodes()); _ = cssPropertyNodeVector
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STYLE) {
            this.LogUtil.PutF("Number Of Properties: " + fmt.Sprint(cssPropertyNodeVector.Size()), this, "CssProperties()")
        }
        var size int = cssPropertyNodeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var cssPropertyValueNode *Node = any(cssPropertyNodeVector.Get(index)).(*Node); _ = cssPropertyValueNode
            var valueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, cssPropertyValueNode.GetChildNodes()); _ = valueNode
            var propertyValue string = DomNodeHelper_GetTextNodeValue(valueNode); _ = propertyValue
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().STYLE) {
                this.LogUtil.PutF("Property Value: " + propertyValue, this, "CssProperties()")
            }
            var indexNode *Node = DomSearchHelper_GetNode(DomData_INDEX, cssPropertyValueNode.GetChildNodes()); _ = indexNode
            var indexValue string = DomNodeHelper_GetTextNodeValue(indexNode); _ = indexValue
            indexPropertyValueHashMap.Put(indexValue, propertyValue)
        }
        var size2 int = indexPropertyValueHashMap.KeySet().Size(); _ = size2
        for index := int(0); index < size2; index++ {
            var indexInt int = int(index); _ = indexInt
            var propertyValue string = any(indexPropertyValueHashMap.Get(fmt.Sprint(indexInt))).(string); _ = propertyValue
            this.PropertyValueVector.Add(propertyValue)
        }
    return this
}


func (this *CssPropertyValues) ToHashMap() *HashMap {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var cssPropertyValueData *CssPropertyValueData = CssPropertyValueData_GetInstance(); _ = cssPropertyValueData
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var size int = this.PropertyValueVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var value string = any(this.PropertyValueVector.Get(index)).(string); _ = value
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(cssPropertyValueData.NAME)
            stringBuffer.Append(commonSeps.BRACKET_OPEN)
            stringBuffer.Appendint(index)
            stringBuffer.Append(commonSeps.BRACKET_CLOSE)
            hashMap.Put(fmt.Sprint(stringBuffer), value)
            index++
        }
        return hashMap
}


// throws Exception
func (this *CssPropertyValues) ToXmlNode(document *Document) *Node {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var hashMap *HashMap = this.ToHashMap(); _ = hashMap
            this.LogUtil.PutF("CssPropertyValues HashMap: " + fmt.Sprint(hashMap), this, "toXmlNode()")
        }
        return ModDomHelper_CreateNameValueIndexNodes(document, CssPropertyData_GetInstance().VALUES, CssPropertyValueData_GetInstance().NAME, this.PropertyValueVector)
}

