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

type BasicProfileActionScriptCondition struct {
    ProfileActionScriptItem
    LogUtil *LogUtil
    ProfileActionProcessorInterfaceVector IBasicArrayList
    ProfileActionConditionInterfaceVector IBasicArrayList
}

// Implements ProfileActionScriptConditionInterface

// The Java class BasicProfileActionScriptCondition has subclasses so its variables are this interface
type IBasicProfileActionScriptCondition interface {
    IProfileActionScriptItem
    BasicProfileActionScriptCondition_() *BasicProfileActionScriptCondition
    AddProcessorNodes(node *Node)
    Init()
    GetConditionJPopupMenu() *JMenu
    GetInputJPopupMenu() *JMenu
    GetOutputJPopupMenu() *JMenu
    AddProcessor(profileActionProcessorInterface ProfileActionScriptProcessorInterface)
    RemoveProcessor(profileActionProcessorInterface ProfileActionScriptProcessorInterface)
    RemoveCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface)
    AddCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface)
    ToXmlNode(document *Document) *Node
    GetProfileActionConditionInterfaceVector() IBasicArrayList
    SetProfileActionConditionInterfaceVector(profileActionConditionInterfaceVector IBasicArrayList)
    ActionPerformed(actionEvent *ActionEvent)
    GetProfileActionProcessorInterfaceVector() IBasicArrayList
    SetProfileActionProcessorInterfaceVector(profileActionProcessorInterfaceVector IBasicArrayList)
    UpdateTree()
    ShouldProcess(frame int64) bool
    Process(frame int64)
}

func (this *BasicProfileActionScriptCondition) BasicProfileActionScriptCondition_() *BasicProfileActionScriptCondition {
    return this
}


func (this *BasicProfileActionScriptCondition) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new BasicProfileActionScriptCondition
func NewBasicProfileActionScriptCondition(args ...any) *BasicProfileActionScriptCondition {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewBasicProfileActionScriptCondition_1(Java_Arg[string](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicProfileActionScriptCondition_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new BasicProfileActionScriptCondition for the arguments")
}
func NewBasicProfileActionScriptCondition_1(label string, node *Node) *BasicProfileActionScriptCondition {
    this := &BasicProfileActionScriptCondition{}
        this.ProfileActionScriptItem = *NewProfileActionScriptItem(label, node)
        this.initFields()
        this.Init()
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var actionItemNode *Node = nodeList.Item(index); _ = actionItemNode
            if actionItemNode.GetNodeType() == Node_ELEMENT_NODE {
                if actionItemNode.GetNodeName().CompareTo(GenericProfileActionScriptConditionData_NAME) == 0 {
                    this.AddCondition(ProfileActionScriptConditionFactory_GetInstance(actionItemNode))
                } else if actionItemNode.GetNodeName().CompareTo(GenericProfileActionScriptInputData_NAME) == 0 {
                    this.AddProcessor(ProfileActionScriptInputFactory_GetInstance(actionItemNode))
                } else if actionItemNode.GetNodeName().CompareTo(GenericProfileActionScriptOutputData_NAME) == 0 {
                    this.AddProcessor(ProfileActionScriptOutputFactory_GetInstance(actionItemNode))
                }
            }
        }
    return this
}


// throws Exception
func (this *BasicProfileActionScriptCondition) AddProcessorNodes(node *Node) {
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var processorActionItemNode *Node = nodeList.Item(index); _ = processorActionItemNode
            if processorActionItemNode.GetNodeName().CompareTo(GenericProfileActionScriptInputData_NAME) == 0 {
                this.AddProcessor(ProfileActionScriptInputFactory_GetInstance(processorActionItemNode))
            } else if processorActionItemNode.GetNodeName().CompareTo(GenericProfileActionScriptOutputData_NAME) == 0 {
                this.AddProcessor(ProfileActionScriptOutputFactory_GetInstance(processorActionItemNode))
            }
        }
}


func NewBasicProfileActionScriptCondition_2(label string) *BasicProfileActionScriptCondition {
    this := &BasicProfileActionScriptCondition{}
        this.ProfileActionScriptItem = *NewProfileActionScriptItem(label)
        this.initFields()
        this.Init()
    return this
}


func (this *BasicProfileActionScriptCondition) Init() {
        this.SetProfileActionProcessorInterfaceVector(NewBasicArrayListD())
        this.SetProfileActionConditionInterfaceVector(NewBasicArrayListD())
        this.GetJPopupMenu().Add(this.GetConditionJPopupMenu())
        this.GetJPopupMenu().Add(this.GetInputJPopupMenu())
        this.GetJPopupMenu().Add(this.GetOutputJPopupMenu())
}


func (this *BasicProfileActionScriptCondition) GetConditionJPopupMenu() *JMenu {
        var jMenu *JMenu = NewJMenu(ProfileActionScriptItem_NEW_CONDITION); _ = jMenu
        var jMenuItemColorAt *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_COLOR_AT); _ = jMenuItemColorAt
        var jMenuItemAlways *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_ALWAYS); _ = jMenuItemAlways
        var jMenuItemTimeInterval *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_TIME_INTERVAL); _ = jMenuItemTimeInterval
        jMenuItemColorAt.AddActionListener(this)
        jMenuItemAlways.AddActionListener(this)
        jMenuItemTimeInterval.AddActionListener(this)
        jMenu.Add(jMenuItemColorAt)
        jMenu.Add(jMenuItemAlways)
        jMenu.Add(jMenuItemTimeInterval)
        return jMenu
}


func (this *BasicProfileActionScriptCondition) GetInputJPopupMenu() *JMenu {
        var jMenu *JMenu = NewJMenu(ProfileActionScriptItem_NEW_INPUT); _ = jMenu
        var jMenuItemKey *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_KEY_INPUT); _ = jMenuItemKey
        var jMenuItemMouse *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_MOUSE_INPUT); _ = jMenuItemMouse
        var jMenuItemJoystick *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_JOYSTICK_INPUT); _ = jMenuItemJoystick
        jMenuItemKey.AddActionListener(this)
        jMenuItemMouse.AddActionListener(this)
        jMenuItemJoystick.AddActionListener(this)
        jMenu.Add(jMenuItemKey)
        jMenu.Add(jMenuItemMouse)
        jMenu.Add(jMenuItemJoystick)
        return jMenu
}


func (this *BasicProfileActionScriptCondition) GetOutputJPopupMenu() *JMenu {
        var jMenu *JMenu = NewJMenu(ProfileActionScriptItem_NEW_OUTPUT); _ = jMenu
        var jMenuItemImage *JMenuItem = NewJMenuItem(ProfileActionScriptItem_NEW_IMAGE_OUTPUT); _ = jMenuItemImage
        jMenuItemImage.AddActionListener(this)
        jMenu.Add(jMenuItemImage)
        return jMenu
}


func (this *BasicProfileActionScriptCondition) AddProcessor(profileActionProcessorInterface ProfileActionScriptProcessorInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "addInput")
        this.GetProfileActionProcessorInterfaceVector().Add(profileActionProcessorInterface)
        this.Add(profileActionProcessorInterface)
}


func (this *BasicProfileActionScriptCondition) RemoveProcessor(profileActionProcessorInterface ProfileActionScriptProcessorInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "removeInput")
        this.GetProfileActionProcessorInterfaceVector().Remove(profileActionProcessorInterface)
        this.Remove(profileActionProcessorInterface)
}


func (this *BasicProfileActionScriptCondition) RemoveCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "removeCondition")
        this.GetProfileActionConditionInterfaceVector().Remove(profileActionScriptNodeInterface)
        this.Remove(profileActionScriptNodeInterface)
}


func (this *BasicProfileActionScriptCondition) AddCondition(profileActionScriptNodeInterface ProfileActionScriptNodeInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "addCondition")
        this.GetProfileActionConditionInterfaceVector().Add(profileActionScriptNodeInterface)
        this.Add(profileActionScriptNodeInterface)
}


// throws Exception
func (this *BasicProfileActionScriptCondition) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionScriptConditionData_NAME); _ = node
        var profileActionConditionInterfaceVector IBasicArrayList = this.GetProfileActionConditionInterfaceVector(); _ = profileActionConditionInterfaceVector
        var size int = profileActionConditionInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var profileActionConditionInterface ProfileActionScriptConditionInterface = any(profileActionConditionInterfaceVector.Get(index)).(ProfileActionScriptConditionInterface); _ = profileActionConditionInterface
            node.AppendChild(profileActionConditionInterface.ToXmlNode(document))
        }
        var profileActionProcessorInterfaceVector IBasicArrayList = this.GetProfileActionProcessorInterfaceVector(); _ = profileActionProcessorInterfaceVector
        var size2 int = profileActionProcessorInterfaceVector.Size(); _ = size2
        for index := int(0); index < size2; index++ {
            var profileActionProcessorInterface ProfileActionScriptProcessorInterface = any(profileActionProcessorInterfaceVector.Get(index)).(ProfileActionScriptProcessorInterface); _ = profileActionProcessorInterface
            node.AppendChild(profileActionProcessorInterface.ToXmlNode(document))
        }
        return node
}


func (this *BasicProfileActionScriptCondition) GetProfileActionConditionInterfaceVector() IBasicArrayList {
        return this.ProfileActionConditionInterfaceVector
}


func (this *BasicProfileActionScriptCondition) SetProfileActionConditionInterfaceVector(profileActionConditionInterfaceVector IBasicArrayList) {
        this.ProfileActionConditionInterfaceVector = profileActionConditionInterfaceVector
}


func (this *BasicProfileActionScriptCondition) ActionPerformed(actionEvent *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Error", this, "actionPerformed", e)
                }
            }()
            this.ProfileActionScriptItem.ActionPerformed(actionEvent)
            if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_DELETE) == 0 {
                if func() bool { _, ok := any(this.GetParent()).(ProfileActionScriptNodeInterface); return ok }() {
                    var profileActionScriptNodeInterface ProfileActionScriptNodeInterface = any(this.GetParent()).(ProfileActionScriptNodeInterface); _ = profileActionScriptNodeInterface
                    profileActionScriptNodeInterface.RemoveCondition(this)
                    JTreeInterfaceFactory_GetInstance().GetJTreeInterface().UpdateJTree()
                }
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_COLOR_AT) == 0 {
                this.AddCondition(NewColorAtActionScriptCondition())
                this.UpdateTree()
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_ALWAYS) == 0 {
                this.AddCondition(NewAlwaysActionScriptCondition())
                this.UpdateTree()
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_TIME_INTERVAL) == 0 {
                this.AddCondition(NewTimeIntervalActionScriptCondition())
                this.UpdateTree()
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_KEY_INPUT) == 0 {
                this.AddProcessor(NewKeyboardActionScriptInput())
                this.UpdateTree()
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_MOUSE_INPUT) == 0 {
                this.AddProcessor(NewMouseActionScriptInput())
                this.UpdateTree()
            } else if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_NEW_IMAGE_OUTPUT) == 0 {
                this.AddProcessor(NewImageActionScriptOutput())
                this.UpdateTree()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BasicProfileActionScriptCondition) GetProfileActionProcessorInterfaceVector() IBasicArrayList {
        return this.ProfileActionProcessorInterfaceVector
}


func (this *BasicProfileActionScriptCondition) SetProfileActionProcessorInterfaceVector(profileActionProcessorInterfaceVector IBasicArrayList) {
        this.ProfileActionProcessorInterfaceVector = profileActionProcessorInterfaceVector
}


func (this *BasicProfileActionScriptCondition) UpdateTree() {
        JTreeInterfaceFactory_GetInstance().GetJTreeInterface().UpdateJTree()
        var treePath *TreePath = NewTreePath(this); _ = treePath
        JTreeInterfaceFactory_GetInstance().GetJTreeInterface().GetActionScriptJTree().ExpandPath(treePath)
}


// throws Exception
func (this *BasicProfileActionScriptCondition) ShouldProcess(frame int64) bool {
        var profileActionConditionInterfaceVector IBasicArrayList = this.GetProfileActionConditionInterfaceVector(); _ = profileActionConditionInterfaceVector
        var size int = profileActionConditionInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var profileActionScriptConditionInterface ProfileActionScriptConditionInterface = any(profileActionConditionInterfaceVector.Get(index)).(ProfileActionScriptConditionInterface); _ = profileActionScriptConditionInterface
            if !profileActionScriptConditionInterface.ShouldProcess(frame) {
                return false
            }
        }
        return true
}


// throws Exception
func (this *BasicProfileActionScriptCondition) Process(frame int64) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Start - Processing ").Appendint(this.GetProfileActionProcessorInterfaceVector().Size()).Append(" inputs")), this, this.CommonStrings.PROCESS)
        var profileActionProcessorInterfaceVector IBasicArrayList = this.GetProfileActionProcessorInterfaceVector(); _ = profileActionProcessorInterfaceVector
        var size int = profileActionProcessorInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var profileActionScriptProcessorInterface ProfileActionScriptProcessorInterface = any(profileActionProcessorInterfaceVector.Get(index)).(ProfileActionScriptProcessorInterface); _ = profileActionScriptProcessorInterface
            profileActionScriptProcessorInterface.Process(frame)
        }
}

