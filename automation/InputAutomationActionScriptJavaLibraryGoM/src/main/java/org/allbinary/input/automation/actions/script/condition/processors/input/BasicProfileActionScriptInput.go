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

type BasicProfileActionScriptInput struct {
    BasicProfileActionScriptProcessor
    LogUtil *LogUtil
    InputRobotInterface InputRobotInterface
    Time int
}

// Implements ProfileActionScriptInputInterface

// The Java class BasicProfileActionScriptInput has subclasses so its variables are this interface
type IBasicProfileActionScriptInput interface {
    IBasicProfileActionScriptProcessor
    BasicProfileActionScriptInput_() *BasicProfileActionScriptInput
    GetTime() int
    SetTime(time int)
    BasicProfileActionScriptInput_toHashMap() *HashMap
    ToXmlNode(document *Document) *Node
    SetInputRobotInterface(inputRobotInterface InputRobotInterface)
    GetInputRobotInterface() InputRobotInterface
    ToString() string
}

func (this *BasicProfileActionScriptInput) BasicProfileActionScriptInput_() *BasicProfileActionScriptInput {
    return this
}


func (this *BasicProfileActionScriptInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new BasicProfileActionScriptInput
func NewBasicProfileActionScriptInput(args ...any) *BasicProfileActionScriptInput {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewBasicProfileActionScriptInput_1(Java_Arg[string](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicProfileActionScriptInput_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new BasicProfileActionScriptInput for the arguments")
}
func NewBasicProfileActionScriptInput_1(label string, node *Node) *BasicProfileActionScriptInput {
    this := &BasicProfileActionScriptInput{}
        this.BasicProfileActionScriptProcessor = *NewBasicProfileActionScriptProcessor(label, node)
        this.initFields()
        this.LogUtil.PutF(fmt.Sprint(this.CommonStrings.START) + label, this, this.CommonStrings.CONSTRUCTOR)
        var actionNode *Node = DomSearchHelper_GetNode(GenericProfileActionScriptInputData_TYPE, node.GetChildNodes()); _ = actionNode
        var inputTypeString string = DomNodeHelper_GetTextNodeValue(actionNode); _ = inputTypeString
        var timeNode *Node = DomSearchHelper_GetNodeNoThrow(GenericProfileActionScriptInputData_DELAY, node.GetChildNodes()); _ = timeNode
        if timeNode != nil {
            var delayString string = DomNodeHelper_GetTextNodeValue(timeNode); _ = delayString
            this.SetTime(int(fmt.Sprint(delayString)))
        } else {
            this.SetTime(0)
        }
        this.SetInputRobotInterface(InputRobotFactory_GetInstance().Get(inputTypeString).(InputRobotInterface))
    return this
}


func NewBasicProfileActionScriptInput_2(label string) *BasicProfileActionScriptInput {
    this := &BasicProfileActionScriptInput{}
        this.BasicProfileActionScriptProcessor = *NewBasicProfileActionScriptProcessor(label)
        this.initFields()
        this.LogUtil.PutF(fmt.Sprint(this.CommonStrings.START) + label, this, this.CommonStrings.CONSTRUCTOR)
        var inputRobotFactory *InputRobotFactory = InputRobotFactory_GetInstance(); _ = inputRobotFactory
        var hashtable *ABHashtable[any, any] = any(inputRobotFactory.Get().(*ABHashtable)).(*ABHashtable[any, any]); _ = hashtable
        var set Set = hashtable.KeySet(); _ = set
        var iterator Iterator = set.Iterator(); _ = iterator
        this.SetInputRobotInterface(inputRobotFactory.Get(any(iterator.Next()).(string)).(InputRobotInterface))
        this.SetTime(0)
    return this
}


func (this *BasicProfileActionScriptInput) GetTime() int {
        return this.Time
}


func (this *BasicProfileActionScriptInput) SetTime(time int) {
        this.Time = time
}


func (this *BasicProfileActionScriptInput) BasicProfileActionScriptInput_toHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(GenericProfileActionScriptInputData_TYPE, this.GetInputRobotInterface().GetName())
        hashMap.Put(GenericProfileActionScriptInputData_DELAY, fmt.Sprint(this.GetTime()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *BasicProfileActionScriptInput) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNodeWithValueNodes(document, GenericProfileActionScriptInputData_NAME, this.BasicProfileActionScriptInput_toHashMap()); _ = node
        return node
}


func (this *BasicProfileActionScriptInput) SetInputRobotInterface(inputRobotInterface InputRobotInterface) {
        this.InputRobotInterface = inputRobotInterface
}


func (this *BasicProfileActionScriptInput) GetInputRobotInterface() InputRobotInterface {
        return this.InputRobotInterface
}


func (this *BasicProfileActionScriptInput) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicProfileActionScriptProcessor.ToString())
        stringBuffer.Append(" Input Type: ")
        if this.GetInputRobotInterface() != nil {
            stringBuffer.Append(this.GetInputRobotInterface().GetName())
        } else {
            stringBuffer.Append(StringUtil_GetInstance().NULL_STRING)
        }
        stringBuffer.Append(" Time: ")
        stringBuffer.Appendlong(int64(this.GetTime()))
        return fmt.Sprint(stringBuffer)
}

func (this *BasicProfileActionScriptInput) String() string {
    return this.ToString()
}

