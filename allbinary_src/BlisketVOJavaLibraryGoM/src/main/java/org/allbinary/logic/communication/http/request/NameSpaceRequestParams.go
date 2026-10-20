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

//containing string values and other hashmaps of the like
type NameSpaceRequestParams struct {
    RequestParams
    LogUtil *LogUtil
}


func (this *NameSpaceRequestParams) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new NameSpaceRequestParams
func NewNameSpaceRequestParams(args ...any) *NameSpaceRequestParams {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewNameSpaceRequestParams_1()
        }
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return NewNameSpaceRequestParams_2(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*PageContext](args[0], pass) {
            return NewNameSpaceRequestParams_3(Java_Arg[*PageContext](args[0]))
        }
    }
    panic("No Java overload of new NameSpaceRequestParams for the arguments")
}
func NewNameSpaceRequestParams_1() *NameSpaceRequestParams {
    this := &NameSpaceRequestParams{}
        this.RequestParams = *NewRequestParams()
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("Constructing empty", this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func NewNameSpaceRequestParams_2(httpServletRequest *HttpServletRequest) *NameSpaceRequestParams {
    this := &NameSpaceRequestParams{}
        this.RequestParams = *NewRequestParams(httpServletRequest)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("Constructing from HttpServletRequest", this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func NewNameSpaceRequestParams_3(pageContext *PageContext) *NameSpaceRequestParams {
    this := &NameSpaceRequestParams{}
        this.RequestParams = *NewRequestParams(pageContext)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("Constructing from PageContext", this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *NameSpaceRequestParams) CreatePackageNode(document *Document, packageName string, packagePropertiesHashMap *HashMap) *Node {
        var node *Node = document.CreateElement(packageName); _ = node
        var propertyKeySet Set = packagePropertiesHashMap.KeySet(); _ = propertyKeySet
        var propertyKeySetArray []any = propertyKeySet.ToArray(); _ = propertyKeySetArray
        var size int = len(propertyKeySetArray); _ = size
        for index := int(0); index < size; index++ {
            var propertyName string = any(propertyKeySetArray[index]).(string); _ = propertyName
            var propertyValue string = any(packagePropertiesHashMap.Get(propertyName)).(string); _ = propertyValue
            var propertyNode *Node = ModDomHelper_CreateTextNode(document, propertyName, propertyValue); _ = propertyNode
            node.AppendChild(propertyNode)
        }
        return node
}


// throws Exception
func (this *NameSpaceRequestParams) CreatePackageMultiNode(document *Document, nameSpaceRequestParam *NameSpaceRequestParam, packageName string, packagePropertiesHashMap *HashMap) *Node {
        var beginIndex int = strings.Index(packageName, string('[')); _ = beginIndex
        var indexValue string = packageName[beginIndex + 1:len(packageName) - 1]; _ = indexValue
        var multiElementPackage string = packageName[0:beginIndex]; _ = multiElementPackage
        var selectedValue string = nameSpaceRequestParam.GetValue(); _ = selectedValue
        var propertyValue string = any(packagePropertiesHashMap.Get(DomData_VALUE)).(string); _ = propertyValue
        if strings.Compare(propertyValue, NameSpaceRequestParamData_VALUE) != 0 {
            selectedValue = propertyValue
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("\nAppending Multinode Element: " + multiElementPackage + "[" + indexValue + "] " + selectedValue, this, "createPackageMultiNode")
        }
        var nextNode *Node = ModDomHelper_CreateNameValueNodes(document, multiElementPackage, selectedValue); _ = nextNode
        var indexNode *Node = ModDomHelper_CreateTextNode(document, DomData_INDEX, indexValue); _ = indexNode
        nextNode.AppendChild(indexNode)
        return nextNode
}


// throws Exception
func (this *NameSpaceRequestParams) GetRootNode(rootElementPackage string, packagePropertiesHashMap *HashMap, document *Document) *Node {
        var rootNode *Node = DomSearchHelper_GetNodeNoThrow(rootElementPackage, document.GetChildNodes()); _ = rootNode
        if rootNode == nil {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("\nAppending Document Root: " + "\nElementName: " + rootElementPackage + "\nProperties: " + fmt.Sprint(packagePropertiesHashMap), this, "getRootNode")
            }
            var node *Node = this.CreatePackageNode(document, rootElementPackage, packagePropertiesHashMap); _ = node
            document.AppendChild(node)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("\nAppended Document Created: " + XmlDocumentHelper_ToString(document), this, "getRootNode")
            }
            return node
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("\nUsing Existing Root Node", this, "getRootNode")
            }
            return rootNode
        }
}


// throws Exception
func (this *NameSpaceRequestParams) AddNewProperties(document *Document, node *Node, nextPackagePropertiesHashMap *HashMap) *Node {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("\nAdding Any New Properties: ").Append(fmt.Sprint(nextPackagePropertiesHashMap)).Append("\nto node: ").Append(node.GetNodeName())), this, "addNewProperties")
        }
        var propertyNameSet Set = nextPackagePropertiesHashMap.KeySet(); _ = propertyNameSet
        var propertyNameSetArray []any = propertyNameSet.ToArray(); _ = propertyNameSetArray
        var size int = len(propertyNameSetArray); _ = size
        for index := int(0); index < size; index++ {
            var nextPropertyName string = any(propertyNameSetArray[index]).(string); _ = nextPropertyName
            var propertyValue string = StringUtil_GetInstance().GetNonNull(any(nextPackagePropertiesHashMap.Get(nextPropertyName)).(string)); _ = propertyValue
            var valueNode *Node = DomSearchHelper_GetNodeNoThrow(nextPropertyName, node.GetChildNodes()); _ = valueNode
            if valueNode != nil {
                var existingElementNodeTextNodeValue string = DomNodeHelper_GetTextNodeValue(valueNode); _ = existingElementNodeTextNodeValue
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("\nIf property: \"" + nextPropertyName + " is the same as:\n" + "Node: " + fmt.Sprint(node.GetNodeName()) + " then: " + propertyValue + "==" + existingElementNodeTextNodeValue, this, "addNewProperties")
                }
                if strings.Compare(propertyValue, existingElementNodeTextNodeValue) != 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                        this.LogUtil.PutF("At least one Property from NameSpaceRequest Package was different from Node", this, "addNewProperties")
                    }
                    panic(errors.New(fmt.Sprint("Property Modifications Should Not Occur")))
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("Adding a new Property: " + nextPropertyName + "=" + propertyValue, this, "addNewProperties")
                }
                var newPropertyNode *Node = ModDomHelper_CreateTextNode(document, nextPropertyName, propertyValue); _ = newPropertyNode
                node.AppendChild(newPropertyNode)
            }
        }
        return node
}


// throws Exception
func (this *NameSpaceRequestParams) IsElementValueTextNodeEqual(nextPackagePropertiesHashMap *HashMap, node *Node) bool {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("\nComparing the value property in: " + fmt.Sprint(nextPackagePropertiesHashMap) + "\nto node child leaves of node: " + fmt.Sprint(node.GetNodeName()), this, "isElementValueTextNodeEqual")
        }
        var propertyValue string = any(nextPackagePropertiesHashMap.Get(DomData_VALUE)).(string); _ = propertyValue
        var valueNode *Node = DomSearchHelper_GetNodeNoThrow(DomData_VALUE, node.GetChildNodes()); _ = valueNode
        if valueNode != nil {
            var existingElementNodeTextNodeValue string = DomNodeHelper_GetTextNodeValue(valueNode); _ = existingElementNodeTextNodeValue
            if strings.Compare(propertyValue, existingElementNodeTextNodeValue) == 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("Value Node is the same", this, "isElementValueTextNodeEqual")
                }
                return true
            }
        } else {
            return true
        }
        return false
}


// throws Exception
func (this *NameSpaceRequestParams) IsElementValueTextNodeUnique(nextPackagePropertiesHashMap *HashMap, elementNodeVector IBasicArrayList) int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("\nComparing Properties of: " + fmt.Sprint(elementNodeVector.Size()) + " Nodes", this, "isElementValueTextNodeUnique")
        }
        var size int = elementNodeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var existingElementNode *Node = any(elementNodeVector.Get(index)).(*Node); _ = existingElementNode
            if this.IsElementValueTextNodeEqual(nextPackagePropertiesHashMap, existingElementNode) {
                return index
            }
        }
        return -1
}


// throws Exception
func (this *NameSpaceRequestParams) AddChildren(document *Document, rootNode *Node, nameSpaceRequestParam *NameSpaceRequestParam) *Document {
        var node *Node = rootNode; _ = node
        var packageVector IBasicArrayList = nameSpaceRequestParam.GetPackages(); _ = packageVector
        var size int = packageVector.Size(); _ = size
        for index := int(1); index < size; index++ {
            var nextPackageName string = any(packageVector.Get(index)).(string); _ = nextPackageName
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("\nTrying to Add Child to Node: ").Append(node.GetNodeName()).Append("\nwith new PackageName: ").Append(nextPackageName)), this, "addChildren")
            }
            var nextPackagePropertiesHashMap *HashMap = nameSpaceRequestParam.GetPackageProperties(index); _ = nextPackagePropertiesHashMap
            var elementNodeVector IBasicArrayList = DomSearchHelper_GetAllNodesNoThrow(nextPackageName, node.GetChildNodes()); _ = elementNodeVector
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("\nDocument Contains " + fmt.Sprint(elementNodeVector.Size()) + " Node(s) With Same Name", this, "addChildren")
            }
            var isElementValueTextNodeUniqueIndex int = this.IsElementValueTextNodeUnique(nextPackagePropertiesHashMap, elementNodeVector); _ = isElementValueTextNodeUniqueIndex
            if (elementNodeVector.Size() == 0 || isElementValueTextNodeUniqueIndex == -1) && !strings.HasSuffix(nextPackageName, CommonSeps_GetInstance().BRACKET_CLOSE) {
                var nextNode *Node = this.CreatePackageNode(document, nextPackageName, nextPackagePropertiesHashMap); _ = nextNode
                node.AppendChild(nextNode)
                node = nextNode
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("\nAppended Package: " + fmt.Sprint(node.GetNodeName()), this, "addChildren")
                    this.LogUtil.PutF("\nAppended Document Created: " + XmlDocumentHelper_ToString(document), this, "addChildren")
                }
            } else if strings.HasSuffix(nextPackageName, CommonSeps_GetInstance().BRACKET_CLOSE) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("Package is MultiNode: just append", this, "addChildren")
                }
                var nextNode *Node = this.CreatePackageMultiNode(document, nameSpaceRequestParam, nextPackageName, nextPackagePropertiesHashMap); _ = nextNode
                node.AppendChild(nextNode)
                node = nextNode
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("\nAppended Multinode Element: " + XmlDocumentHelper_ToString(document), this, "addChildren")
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("Value Node already exists: Now compare (properties to leaf nodes)", this, "addChildren")
                }
                var nodeNameDuplicateNode *Node = any(elementNodeVector.Get(isElementValueTextNodeUniqueIndex)).(*Node); _ = nodeNameDuplicateNode
                node = this.AddNewProperties(document, nodeNameDuplicateNode, nextPackagePropertiesHashMap)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                    this.LogUtil.PutF("\nDuplicate Node Found: Moving to next child ^", this, "addChildre")
                }
            }
        }
        return document
}


// throws Exception
func (this *NameSpaceRequestParams) AddNameSpace(key string, value string, document *Document) *Document {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("NameSpace key: ").Append(key).Append(" Value: ").Append(value)), this, "addNameSpace")
        }
        var nameSpaceRequestParam *NameSpaceRequestParam = NewNameSpaceRequestParam(key, value); _ = nameSpaceRequestParam
        var packageVector IBasicArrayList = nameSpaceRequestParam.GetPackages(); _ = packageVector
        var size int = packageVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var packageName string = any(packageVector.Get(index)).(string); _ = packageName
            var packagePropertiesHashMap *HashMap = nameSpaceRequestParam.GetPackageProperties(packageVector.IndexOf(packageName)); _ = packagePropertiesHashMap
            var rootNode *Node = this.GetRootNode(packageName, packagePropertiesHashMap, document); _ = rootNode
            if index < size - 1 {
                return this.AddChildren(document, any(rootNode).(*Node), nameSpaceRequestParam)
            }
        }
        return document
}


// throws Exception
func (this *NameSpaceRequestParams) ToHashMap() *HashMap {
        var document *Document = DomDocumentHelper_Create(); _ = document
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var keys Set = this.GetMap().KeySet(); _ = keys
        var keyArray []any = keys.ToArray(); _ = keyArray
        var keySize int = len(keyArray); _ = keySize
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("NameSpace Request Size: " + fmt.Sprint(keys.Size()), this, "toHashMap")
        }
        for i := int(0); i < keySize; i++ {
            var key string = any(keyArray[i]).(string); _ = key
            var anyType any = this.GetMap().Get(key); _ = anyType
            var className string = any(reflect.TypeOf(anyType).String()).(string); _ = className
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                this.LogUtil.PutF("Request Param Class: " + className, this, "toHashMap")
            }
            if func() bool { _, ok := any(this.GetMap().Get(key)).(string); return ok }() {
                var value string = any(this.GetMap().Get(key)).(string); _ = value
                if strings.Index(key, string(NameSpaceRequestParamData_NAME)) < 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                        this.LogUtil.PutF("Not NameSpace key: " + key + " Value: " + value, this, "toHashMap()")
                    }
                    hashMap.Put(key, value)
                } else {
                    document = this.AddNameSpace(key, value, document)
                }
            } else if func() bool { _, ok := any(this.GetMap().Get(key)).([]string); return ok }() {
                var values []string = any(this.GetMap().Get(key)).([]string); _ = values
                if strings.Index(key, string(NameSpaceRequestParamData_NAME)) < 0 {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
                        this.LogUtil.PutF("Not NameSpace key: " + key + " Value: " + values[0], this, "toHashMap()")
                    }
                    hashMap.Put(key, values[0])
                } else {
                    document = this.AddNameSpace(key, values[0], document)
                }
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPREQUEST) {
            this.LogUtil.PutF("NameSpaceRequestParams: " + fmt.Sprint(this.GetMap()) + "\ntoHashMap(): " + fmt.Sprint(hashMap) + "\nDocument Created: " + XmlDocumentHelper_ToString(document), this, "toHashMap()")
        }
        hashMap.Put(NameSpaceRequestParamData_DOCUMENT, document)
        return hashMap
}

