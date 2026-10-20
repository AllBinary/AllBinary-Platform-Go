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

type MouseActionScriptInput struct {
    BasicProfileActionScriptInput
    LogUtil *LogUtil
    Point *Point
    Buttons int
    MouseActionScriptInputJPanel *MouseActionScriptInputJPanel
}

// Implements MouseActionScriptInputInterface

var MouseActionScriptInput_NAME string = "Mouse"

func (this *MouseActionScriptInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}




// Java overloads of new MouseActionScriptInput
func NewMouseActionScriptInput(args ...any) *MouseActionScriptInput {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewMouseActionScriptInput_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewMouseActionScriptInput_2()
        }
    }
    panic("No Java overload of new MouseActionScriptInput for the arguments")
}
func NewMouseActionScriptInput_1(node *Node) *MouseActionScriptInput {
    this := &MouseActionScriptInput{}
        this.BasicProfileActionScriptInput = *NewBasicProfileActionScriptInput(MouseActionScriptInput_NAME, node)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.Point = NewPoint()
        var actionNode *Node = DomSearchHelper_GetNode(MouseActionScriptInputData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().CompareTo(MouseActionScriptInputData_BUTTONS) == 0 {
                    var buttons string = DomNodeHelper_GetTextNodeValue(childNode); _ = buttons
                    this.SetButtonClicks(int(fmt.Sprint(buttons)))
                } else if childNode.GetNodeName().CompareTo(MouseActionScriptInputData_MOVE) == 0 {
                    var mouseXNode *Node = DomSearchHelper_GetNode(MouseActionScriptInputData_MOVE_X, childNode.GetChildNodes()); _ = mouseXNode
                    var mouseXString string = DomNodeHelper_GetTextNodeValue(mouseXNode); _ = mouseXString
                    var mouseXInteger int = fmt.Sprint(mouseXString); _ = mouseXInteger
                    var mouseYNode *Node = DomSearchHelper_GetNode(MouseActionScriptInputData_MOVE_Y, childNode.GetChildNodes()); _ = mouseYNode
                    var mouseYString string = DomNodeHelper_GetTextNodeValue(mouseYNode); _ = mouseYString
                    var mouseYInteger int = fmt.Sprint(mouseYString); _ = mouseYInteger
                    var newPoint *Point = NewPoint(int(mouseXInteger), int(mouseYInteger)); _ = newPoint
                    this.SetPoint(newPoint)
                } else {
                    panic(errors.New(fmt.Sprint("Action Script Input Unknown Node")))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Action Script Input Node Null")))
        }
        this.SetAllowsChildren(false)
        this.MouseActionScriptInputJPanel = NewMouseActionScriptInputJPanel(this)
    return this
}


func NewMouseActionScriptInput_2() *MouseActionScriptInput {
    this := &MouseActionScriptInput{}
        this.BasicProfileActionScriptInput = *NewBasicProfileActionScriptInput(MouseActionScriptInput_NAME)
        this.initFields()
        this.Point = NewPoint()
        this.SetAllowsChildren(false)
        this.MouseActionScriptInputJPanel = NewMouseActionScriptInputJPanel(this)
    return this
}


func (this *MouseActionScriptInput) SetPoint(point *Point) {
        this.Point = point
}


func (this *MouseActionScriptInput) GetPoint() *Point {
        return this.Point
}


func (this *MouseActionScriptInput) GetButtonClicks() int {
        return this.Buttons
}


func (this *MouseActionScriptInput) SetButtonClicks(buttons int) {
        this.Buttons = buttons
}


func (this *MouseActionScriptInput) ShowDialog() {
        this.MouseActionScriptInputJPanel.GetMouseActionJDialog().SetVisible(true)
}


func (this *MouseActionScriptInput) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(MouseActionScriptInputData_BUTTONS, fmt.Sprint(this.GetButtonClicks()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *MouseActionScriptInput) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptInput.ToXmlNode(document); _ = node
        var mouseNode *Node = ModDomHelper_CreateNodeWithValueNodes(document, MouseActionScriptInputData_NAME, this.ToHashMap()); _ = mouseNode
        node.AppendChild(mouseNode)
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(MouseActionScriptInputData_MOVE_X, fmt.Sprint(this.GetPoint().X))
        hashMap.Put(MouseActionScriptInputData_MOVE_Y, fmt.Sprint(this.GetPoint().Y))
        mouseNode.AppendChild(ModDomHelper_CreateNodeWithValueNodes(document, MouseActionScriptInputData_MOVE, hashMap))
        return node
}


// throws Exception
func (this *MouseActionScriptInput) Process(frame int64) {
        MouseInputAutomationProcessor_Process(this)
}


func (this *MouseActionScriptInput) Log() {
        var buttonStringBuffer *StringMaker = NewStringMaker(); _ = buttonStringBuffer
        buttonStringBuffer.Append("1: ")
        var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
        if (this.GetButtonClicks() & InputEvent_BUTTON1_MASK) != 0 {
            buttonStringBuffer.Append(booleanFactory.TRUE_STRING)
        } else {
            buttonStringBuffer.Append(booleanFactory.FALSE_STRING)
        }
        buttonStringBuffer.Append(" 2: ")
        if (this.GetButtonClicks() & InputEvent_BUTTON2_MASK) != 0 {
            buttonStringBuffer.Append(booleanFactory.TRUE_STRING)
        } else {
            buttonStringBuffer.Append(booleanFactory.FALSE_STRING)
        }
        buttonStringBuffer.Append(" 3: ")
        if (this.GetButtonClicks() & InputEvent_BUTTON3_MASK) != 0 {
            buttonStringBuffer.Append(booleanFactory.TRUE_STRING)
        } else {
            buttonStringBuffer.Append(booleanFactory.FALSE_STRING)
        }
        this.LogUtil.PutF("Input Type: " + this.GetInputRobotInterface().GetName() + " Point: " + fmt.Sprint(this.GetPoint()) + " Buttons Clicked: " + fmt.Sprint(buttonStringBuffer), this, "log")
}

