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

type CssElementView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TITLE string
    DESCRIPTION string
    Value string
    Title string
    Description string
    CssElementVector IBasicArrayList
    CssPropertyVector IBasicArrayList
}

// Implements DomNodeInterface

// The Java class CssElementView has subclasses so its variables are this interface
type ICssElementView interface {
    CssElementView_() *CssElementView
    GetValue() string
    ToHashMap() *HashMap
    ToXmlNode(document *Document) *Node
}

func (this *CssElementView) CssElementView_() *CssElementView {
    return this
}


func (this *CssElementView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.TITLE = "title"
    this.DESCRIPTION = "description"
    
    
    
    
    
}




// Java overloads of new CssElementView
func NewCssElementView(args ...any) *CssElementView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewCssElementView_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewCssElementView_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new CssElementView for the arguments")
}
func NewCssElementView_1() *CssElementView {
    this := &CssElementView{}
    this.initFields()
        this.CssPropertyVector = NewBasicArrayListD()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.Title = stringUtil.EMPTY_STRING
        this.Description = stringUtil.EMPTY_STRING
        this.Value = stringUtil.EMPTY_STRING
    return this
}


func NewCssElementView_2(node *Node) *CssElementView {
    this := &CssElementView{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append("Value: ").Append(this.Value).Append("\nTitle: ").Append(this.Title).Append("\nDescription: ").Append(this.Description)), this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(errors.New(fmt.Sprint("CssElementView")))
                }
            }()
            var valueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, node.GetChildNodes()); _ = valueNode
            this.Value = DomNodeHelper_GetTextNodeValue(valueNode)
            var titleValueNode *Node = DomSearchHelper_GetNodeNoThrow(this.TITLE, node.GetChildNodes()); _ = titleValueNode
            if titleValueNode != nil {
                this.Title = DomNodeHelper_GetTextNodeValue(titleValueNode)
            } else {
                var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
                this.Title = stringUtil.EMPTY_STRING
            }
            var descriptionValueNode *Node = DomSearchHelper_GetNodeNoThrow(this.DESCRIPTION, node.GetChildNodes()); _ = descriptionValueNode
            if descriptionValueNode != nil {
                this.Description = DomNodeHelper_GetTextNodeValue(descriptionValueNode)
            } else {
                this.Description = StringUtil_GetInstance().EMPTY_STRING
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Value: ").Append(this.Value).Append("\nTitle: ").Append(this.Title).Append("\nDescription: ").Append(this.Description)), this, this.CommonStrings.CONSTRUCTOR)
            }
            var cssElementStyleNodeVector IBasicArrayList = DomSearchHelper_GetAllNodes(CssElementData_GetInstance().NAME, node.GetChildNodes()); _ = cssElementStyleNodeVector
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Child Css Elements: " + fmt.Sprint(cssElementStyleNodeVector.Size()), this, "CssElementView()")
            }
            this.CssElementVector = CssElementsValidationFactory_GetInstance(cssElementStyleNodeVector)
            var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
            this.CssPropertyVector = CssPropertiesValidationFactory_GetInstance(nodeList)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *CssElementView) GetValue() string {
        return this.Value
}


func (this *CssElementView) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(DomData_VALUE, this.Value)
        hashMap.Put(this.TITLE, this.Title)
        hashMap.Put(this.DESCRIPTION, this.Description)
        return hashMap
}


// throws Exception
func (this *CssElementView) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("CssView HashMap: " + fmt.Sprint(hashMap), this, "toXmlNode")
        }
        var node *Node = ModDomHelper_CreateNodeWithValueNodes(document, CssElementData_GetInstance().NAME, hashMap); _ = node
        var size int = this.CssPropertyVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var cssPropertyDomNodeInterface DomNodeInterface = any(this.CssPropertyVector.Get(index)).(DomNodeInterface); _ = cssPropertyDomNodeInterface
            node.AppendChild(cssPropertyDomNodeInterface.ToXmlNode(document))
        }
        var size2 int = this.CssElementVector.Size(); _ = size2
        for index := int(0); index < size2; index++ {
            var cssElementDomNodeInterface DomNodeInterface = any(this.CssElementVector.Get(index)).(DomNodeInterface); _ = cssElementDomNodeInterface
            node.AppendChild(cssElementDomNodeInterface.ToXmlNode(document))
        }
        return node
}

