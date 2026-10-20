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

type AlwaysActionScriptCondition struct {
    BasicProfileActionScriptCondition
    LogUtil *LogUtil
    AlwaysActionScriptConditionJPanel *AlwaysActionScriptConditionJPanel
    IsOn bool
}

// Implements AlwaysActionScriptConditionInterface

var AlwaysActionScriptCondition_NAME string = "Always On/Off"

func (this *AlwaysActionScriptCondition) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new AlwaysActionScriptCondition
func NewAlwaysActionScriptCondition(args ...any) *AlwaysActionScriptCondition {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewAlwaysActionScriptCondition_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewAlwaysActionScriptCondition_2()
        }
    }
    panic("No Java overload of new AlwaysActionScriptCondition for the arguments")
}
func NewAlwaysActionScriptCondition_1(node *Node) *AlwaysActionScriptCondition {
    this := &AlwaysActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(AlwaysActionScriptCondition_NAME, node)
        this.initFields()
        var actionNode *Node = DomSearchHelper_GetNode(AlwaysActionScriptConditionData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().CompareTo(AlwaysActionScriptConditionData_IS_ON) == 0 {
                    var booleanString string = DomNodeHelper_GetTextNodeValue(childNode); _ = booleanString
                    this.SetIsOn(booleanString)
                } else {
                    panic(errors.New(fmt.Sprint("Action Script Condition Unknown Node")))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Action Script Condition Node Null")))
        }
        this.Init()
    return this
}


func NewAlwaysActionScriptCondition_2() *AlwaysActionScriptCondition {
    this := &AlwaysActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(AlwaysActionScriptCondition_NAME)
        this.initFields()
        this.Init()
    return this
}


func (this *AlwaysActionScriptCondition) Init() {
        this.AlwaysActionScriptConditionJPanel = NewAlwaysActionScriptConditionJPanel(this)
}


func (this *AlwaysActionScriptCondition) IsIsOn() bool {
        return this.IsOn
}


func (this *AlwaysActionScriptCondition) SetIsOn(isOn bool) {
        this.IsOn = isOn
}


func (this *AlwaysActionScriptCondition) ShowDialog() {
        this.AlwaysActionScriptConditionJPanel.GetAlwaysActionJDialog().SetVisible(true)
}


func (this *AlwaysActionScriptCondition) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(AlwaysActionScriptConditionData_IS_ON, fmt.Sprint(this.IsIsOn()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *AlwaysActionScriptCondition) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptCondition.ToXmlNode(document); _ = node
        node.AppendChild(ModDomHelper_CreateNodeWithValueNodes(document, AlwaysActionScriptConditionData_NAME, this.ToHashMap()))
        return node
}


// throws Exception
func (this *AlwaysActionScriptCondition) ShouldProcess(frame int64) bool {
        if this.IsIsOn() {
            return this.BasicProfileActionScriptCondition.ShouldProcess(frame)
        }
        return false
}


func (this *AlwaysActionScriptCondition) Log() {
        this.LogUtil.PutF("Is On: " + fmt.Sprint(this.IsIsOn()), this, "log")
}

