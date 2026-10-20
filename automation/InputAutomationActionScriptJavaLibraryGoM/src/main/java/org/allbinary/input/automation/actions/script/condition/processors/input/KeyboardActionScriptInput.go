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

type KeyboardActionScriptInput struct {
    BasicProfileActionScriptInput
    LogUtil *LogUtil
    KeyArray []int
    KeyboardActionScriptInputJPanel *KeyboardActionScriptInputJPanel
    Press bool
    Release bool
    DelayBetweenKeys int
}

// Implements KeyboardActionScriptInputInterface

// The Java class KeyboardActionScriptInput has subclasses so its variables are this interface
type IKeyboardActionScriptInput interface {
    IBasicProfileActionScriptInput
    KeyboardActionScriptInput_() *KeyboardActionScriptInput
    GetDelayBetweenKeys() int
    SetDelayBetweenKeys(delayBetweenKeys int)
    IsPress() bool
    SetPress(press bool)
    IsRelease() bool
    SetRelease(release bool)
    IsNormal() bool
    SetNormal()
    GetKeyArray() []int
    SetKeyArray(keyArray []int)
    ShowDialog()
    GetText() string
    SetText(text string)
    IntegerArrayValue(text string) []int
    ToHashMap() *HashMap
    ToXmlNode(document *Document) *Node
    Process(frame int64)
    Log()
    ToString() string
}

func (this *KeyboardActionScriptInput) KeyboardActionScriptInput_() *KeyboardActionScriptInput {
    return this
}

var KeyboardActionScriptInput_NAME string = "Keyboard"

func (this *KeyboardActionScriptInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.KeyArray = []int{KeySingletonFactory_GetInstance(KeyEvent_VK_0)}
    
    
    
    
}




// Java overloads of new KeyboardActionScriptInput
func NewKeyboardActionScriptInput(args ...any) *KeyboardActionScriptInput {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewKeyboardActionScriptInput_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewKeyboardActionScriptInput_2()
        }
    }
    panic("No Java overload of new KeyboardActionScriptInput for the arguments")
}
func NewKeyboardActionScriptInput_1(node *Node) *KeyboardActionScriptInput {
    this := &KeyboardActionScriptInput{}
        this.BasicProfileActionScriptInput = *NewBasicProfileActionScriptInput(KeyboardActionScriptInput_NAME, node)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        var actionNode *Node = DomSearchHelper_GetNode(KeyboardActionScriptInputData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().StartsWith(KeyboardActionScriptInputData_KEY) {
                    var keyString string = DomNodeHelper_GetTextNodeValue(childNode); _ = keyString
                    vector.Add(fmt.Sprint(keyString))
                } else if childNode.GetNodeName().CompareTo(KeyboardActionScriptInputData_DELAY_BETWEEN_KEYS) == 0 {
                    var time string = DomNodeHelper_GetTextNodeValue(childNode); _ = time
                    this.SetDelayBetweenKeys(int(fmt.Sprint(time)))
                } else if childNode.GetNodeName().CompareTo(KeyboardActionScriptInputData_PRESS) == 0 {
                    var value string = DomNodeHelper_GetTextNodeValue(childNode); _ = value
                    this.SetPress(fmt.Sprint(value))
                } else if childNode.GetNodeName().CompareTo(KeyboardActionScriptInputData_RELEASE) == 0 {
                    var value string = DomNodeHelper_GetTextNodeValue(childNode); _ = value
                    this.SetRelease(fmt.Sprint(value))
                } else {
                    panic(errors.New(fmt.Sprint("Action Script Input Unknown Node")))
                }
            }
            this.SetKeyArray(any(vector.ToArrayType(make([]int, vector.Size()))).([]int))
        } else {
            panic(errors.New(fmt.Sprint("Action Script Input Node Null")))
        }
        this.SetAllowsChildren(false)
        if !this.IsPress() && !this.IsRelease() {
            this.SetNormal()
        }
        this.KeyboardActionScriptInputJPanel = NewKeyboardActionScriptInputJPanel(this)
    return this
}


func NewKeyboardActionScriptInput_2() *KeyboardActionScriptInput {
    this := &KeyboardActionScriptInput{}
        this.BasicProfileActionScriptInput = *NewBasicProfileActionScriptInput(KeyboardActionScriptInput_NAME)
        this.initFields()
        if !this.IsPress() && !this.IsRelease() {
            this.SetNormal()
        }
        this.SetAllowsChildren(false)
        this.SetTime(120)
        this.KeyboardActionScriptInputJPanel = NewKeyboardActionScriptInputJPanel(this)
    return this
}


func (this *KeyboardActionScriptInput) GetDelayBetweenKeys() int {
        return this.DelayBetweenKeys
}


func (this *KeyboardActionScriptInput) SetDelayBetweenKeys(delayBetweenKeys int) {
        this.DelayBetweenKeys = delayBetweenKeys
}


func (this *KeyboardActionScriptInput) IsPress() bool {
        return this.Press
}


func (this *KeyboardActionScriptInput) SetPress(press bool) {
        this.Press = press
}


func (this *KeyboardActionScriptInput) IsRelease() bool {
        return this.Release
}


func (this *KeyboardActionScriptInput) SetRelease(release bool) {
        this.Release = release
}


func (this *KeyboardActionScriptInput) IsNormal() bool {
        return this.IsRelease() && this.IsPress()
}


func (this *KeyboardActionScriptInput) SetNormal() {
        this.SetPress(true)
        this.SetRelease(true)
}


func (this *KeyboardActionScriptInput) GetKeyArray() []int {
        return this.KeyArray
}


func (this *KeyboardActionScriptInput) SetKeyArray(keyArray []int) {
        if keyArray != nil {
            this.LogUtil.PutF(CommonLabels_GetInstance().START + len(keyArray), this, "setKeyArray")
        }
        this.KeyArray = keyArray
}


func (this *KeyboardActionScriptInput) ShowDialog() {
        this.KeyboardActionScriptInputJPanel.GetKeyActionJDialog().SetVisible(true)
}


func (this *KeyboardActionScriptInput) GetText() string {
        return this.GetText(this.GetKeyArray())
}


func (this *KeyboardActionScriptInput) SetText(text string) {
        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + text, this, "setText")
        var integerArray []int = this.IntegerArrayValue(text); _ = integerArray
        this.SetKeyArray(integerArray)
}


func KeyboardActionScriptInput_GetText(integerArray []int) string {
        var index int = 0; _ = index
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index < len(integerArray) {
            var nextInteger int = integerArray[index]; _ = nextInteger
            if KeyUtil_IsNormallyDisplayed(nextInteger) {
                stringBuffer.Append(KeyEvent_GetKeyText(nextInteger))
            } else {
                stringBuffer.Append("&#" + fmt.Sprint(nextInteger) + ";")
            }
            index++
        }
        return fmt.Sprint(stringBuffer)
}


func (this *KeyboardActionScriptInput) IntegerArrayValue(text string) []int {
        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + text, this, "integerArrayValue")
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var index int = 0; _ = index
        for index < len(text) {
            var aChar rune = rune(text[index]); _ = aChar
            if aChar == '&' {
                if rune(text[index + 1]) == '#' {
                    var endIndex int = text.IndexOf(';', index + 1); _ = endIndex
                    if endIndex != -1 {
                        var nextCharString string = text[index + 2:endIndex]; _ = nextCharString
                        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().NEXT) + " Char String: " + nextCharString, this, "integerArrayValue")
                        vector.Add(fmt.Sprint(nextCharString[index:len(nextCharString)]))
                        index = index + len(nextCharString) + 3
                        continue
                    }
                }
            }
            vector.Add(KeySingletonFactory_GetHashtable().Get(fmt.Sprint(aChar)))
            index++
        }
        var integerArray []int = any(vector.ToArrayType(make([]int, vector.Size()))).([]int); _ = integerArray
        return integerArray
}


func (this *KeyboardActionScriptInput) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        for index := int(0); index < len(this.GetKeyArray()); index++ {
            hashMap.Put(KeyboardActionScriptInputData_KEY + index, fmt.Sprint(this.GetKeyArray()[index]))
        }
        hashMap.Put(KeyboardActionScriptInputData_DELAY_BETWEEN_KEYS, fmt.Sprint(this.GetDelayBetweenKeys()))
        hashMap.Put(KeyboardActionScriptInputData_PRESS, fmt.Sprint(this.IsPress()))
        hashMap.Put(KeyboardActionScriptInputData_RELEASE, fmt.Sprint(this.IsRelease()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *KeyboardActionScriptInput) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptInput.ToXmlNode(document); _ = node
        node.AppendChild(ModDomHelper_CreateNodeWithValueNodes(document, KeyboardActionScriptInputData_NAME, this.ToHashMap()))
        return node
}


// throws Exception
func (this *KeyboardActionScriptInput) Process(frame int64) {
        KeyboardInputAutomationProcessor_Process(this)
}


func (this *KeyboardActionScriptInput) Log() {
        this.LogUtil.PutF(this.ToString(), this, "log")
}


func (this *KeyboardActionScriptInput) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicProfileActionScriptInput.ToString())
        stringBuffer.Append(" Text: ")
        stringBuffer.Append(this.GetText())
        stringBuffer.Append(" isPress: ")
        stringBuffer.Appendboolean(this.IsPress())
        stringBuffer.Append(" isRelease: ")
        stringBuffer.Appendboolean(this.IsRelease())
        stringBuffer.Append(" Between Key Delay: ")
        stringBuffer.Appendint(this.GetDelayBetweenKeys())
        return fmt.Sprint(stringBuffer)
}

func (this *KeyboardActionScriptInput) String() string {
    return this.ToString()
}

