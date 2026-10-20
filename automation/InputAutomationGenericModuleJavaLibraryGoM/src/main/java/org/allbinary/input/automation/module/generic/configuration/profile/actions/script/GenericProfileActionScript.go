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

type GenericProfileActionScript struct {
    DefaultMutableTreeNode
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GenericProfileActionJPanel *GenericProfileActionJPanel
    ProfileActionConditionInterfaceVector IBasicArrayList
}

// Implements ProfileActionScriptNodeInterface

var GenericProfileActionScript_NAME string = "Root"

func (this *GenericProfileActionScript) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}




// Java overloads of new GenericProfileActionScript
func NewGenericProfileActionScript(args ...any) *GenericProfileActionScript {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*GenericProfileActionJPanel](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewGenericProfileActionScript_1(Java_Arg[*GenericProfileActionJPanel](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*GenericProfileActionJPanel](args[0], pass) {
            return NewGenericProfileActionScript_2(Java_Arg[*GenericProfileActionJPanel](args[0]))
        }
    }
    panic("No Java overload of new GenericProfileActionScript for the arguments")
}
func NewGenericProfileActionScript_1(genericProfileActionJPanel *GenericProfileActionJPanel, node *Node) *GenericProfileActionScript {
    this := &GenericProfileActionScript{}
        this.DefaultMutableTreeNode = *NewDefaultMutableTreeNode(GenericProfileActionScript_NAME)
        this.initFields()
        this.GenericProfileActionJPanel = genericProfileActionJPanel
        this.SetProfileActionConditionInterfaceVector(NewBasicArrayListD())
        var actionScriptNode *Node = DomSearchHelper_GetNode(GenericProfileActionData_SCRIPT, node.GetChildNodes()); _ = actionScriptNode
        if actionScriptNode != nil {
            var nodeList *NodeList = actionScriptNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var actionConditionNode *Node = nodeList.Item(index); _ = actionConditionNode
                if actionConditionNode.GetNodeType() == Node_ELEMENT_NODE {
                    this.AddCondition(ProfileActionScriptConditionFactory_GetInstance(actionConditionNode))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Profile Action Script Node Null")))
        }
    return this
}


func NewGenericProfileActionScript_2(genericProfileActionJPanel *GenericProfileActionJPanel) *GenericProfileActionScript {
    this := &GenericProfileActionScript{}
    this.DefaultMutableTreeNode = *NewDefaultMutableTreeNode()
    this.initFields()
        this.GenericProfileActionJPanel = genericProfileActionJPanel
        this.SetProfileActionConditionInterfaceVector(NewBasicArrayListD())
    return this
}


func (this *GenericProfileActionScript) AddCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "addCondition")
        this.GetProfileActionConditionInterfaceVector().Add(profileActionScriptNodeInterface)
        this.Add(profileActionScriptNodeInterface)
}


func (this *GenericProfileActionScript) RemoveCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "removeCondition")
        this.GetProfileActionConditionInterfaceVector().Remove(profileActionScriptNodeInterface)
        this.Remove(profileActionScriptNodeInterface)
}


func (this *GenericProfileActionScript) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


func (this *GenericProfileActionScript) GetProfileActionConditionInterfaceVector() IBasicArrayList {
        return this.ProfileActionConditionInterfaceVector
}


func (this *GenericProfileActionScript) SetProfileActionConditionInterfaceVector(profileActionConditionInterfaceVector IBasicArrayList) {
        this.ProfileActionConditionInterfaceVector = profileActionConditionInterfaceVector
}


// throws Exception
func (this *GenericProfileActionScript) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionData_SCRIPT); _ = node
        var vector IBasicArrayList = this.GetProfileActionConditionInterfaceVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var profileActionConditionInterface ProfileActionScriptConditionInterface = any(vector.Get(index)).(ProfileActionScriptConditionInterface); _ = profileActionConditionInterface
            node.AppendChild(profileActionConditionInterface.ToXmlNode(document))
        }
        return node
}

