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

type InputAutomationModuleConfiguration struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ClassName string
    Name string
    InputAutomationModuleInterface IInputAutomationModuleFactoryInterface
}


func (this *InputAutomationModuleConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}




// Java overloads of new InputAutomationModuleConfiguration
func NewInputAutomationModuleConfiguration(args ...any) *InputAutomationModuleConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewInputAutomationModuleConfiguration_1()
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewInputAutomationModuleConfiguration_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewInputAutomationModuleConfiguration_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 1 && Java_IsArg[IInputAutomationModuleFactoryInterface](args[0], pass) {
            return NewInputAutomationModuleConfiguration_4(Java_Arg[IInputAutomationModuleFactoryInterface](args[0]))
        }
    }
    panic("No Java overload of new InputAutomationModuleConfiguration for the arguments")
}
func NewInputAutomationModuleConfiguration_1() *InputAutomationModuleConfiguration {
    this := &InputAutomationModuleConfiguration{}
    this.initFields()
    return this
}


func NewInputAutomationModuleConfiguration_2(abeClientInformation AbeClientInformationInterface, node *Node) *InputAutomationModuleConfiguration {
    this := &InputAutomationModuleConfiguration{}
    this.initFields()
        this.Init(abeClientInformation, node)
    return this
}


func NewInputAutomationModuleConfiguration_3(abeClientInformation AbeClientInformationInterface, name string, className string) *InputAutomationModuleConfiguration {
    this := &InputAutomationModuleConfiguration{}
    this.initFields()
        this.SetName(name)
        this.SetClassName(className)
        this.Init(abeClientInformation)
    return this
}


func NewInputAutomationModuleConfiguration_4(inputAutomationModuleInterface IInputAutomationModuleFactoryInterface) *InputAutomationModuleConfiguration {
    this := &InputAutomationModuleConfiguration{}
    this.initFields()
        this.SetInputAutomationModuleInterface(inputAutomationModuleInterface)
        this.SetClassName(reflect.TypeOf(this.InputAutomationModuleInterface).String())
    return this
}


// throws Exception
func (this *InputAutomationModuleConfiguration) Init_1(abeClientInformation AbeClientInformationInterface, node *Node) {
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        if nodeList != nil {
            var classNameNode *Node = DomSearchHelper_GetNode(DynamicObjectData_NAME, nodeList); _ = classNameNode
            if classNameNode != nil {
                this.SetClassName(DomNodeHelper_GetTextNodeValue(classNameNode))
                this.Init(abeClientInformation)
            } else {
                this.LogUtil.PutF("Class Node Null", this, this.CommonStrings.INIT)
            }
        } else {
            this.LogUtil.PutF(fmt.Sprint(InputAutomationData_NAME) + " Node Has No Children", this, "Contructor")
        }
}

// Java overloads of InputAutomationModuleConfiguration.init
func (this *InputAutomationModuleConfiguration) Init(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            this.Init_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*Node](args[1]))
            return
        }
        if len(args) == 1 && Java_IsArg[AbeClientInformationInterface](args[0], pass) {
            this.Init_2(Java_Arg[AbeClientInformationInterface](args[0]))
            return
        }
    }
    panic("No Java overload of InputAutomationModuleConfiguration.init for the arguments")
}


func (this *InputAutomationModuleConfiguration) Init_2(abeClientInformation AbeClientInformationInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
            panic(errors.New(fmt.Sprint()))
                }
            }()
            this.LogUtil.PutF("Name: " + this.GetName(), this, this.CommonStrings.INIT)
            this.LogUtil.PutF("ClassName: " + this.ClassName, this, this.CommonStrings.INIT)
            this.SetInputAutomationModuleInterface(any(AbeFactory_GetInstance().GetInstance(abeClientInformation, this.GetClassName())).(IInputAutomationModuleFactoryInterface))
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputAutomationModuleConfiguration) ToDomNode(document *Document) *Node {
        var node *Node = document.CreateElement(InputAutomationModuleData_NAME); _ = node
        var classNameNode *Node = ModDomHelper_CreateTextNode(document, DynamicObjectData_NAME, this.GetClassName()); _ = classNameNode
        node.AppendChild(classNameNode)
        return node
}


func (this *InputAutomationModuleConfiguration) GetName() string {
        return this.Name
}


func (this *InputAutomationModuleConfiguration) SetName(name string) {
        this.LogUtil.PutF("Name: " + name, this, "setName")
        this.Name = name
}


func (this *InputAutomationModuleConfiguration) GetClassName() string {
        return this.ClassName
}


func (this *InputAutomationModuleConfiguration) SetClassName(className string) {
        this.LogUtil.PutF("ClassName : " + className, this, "setClassName")
        this.ClassName = className
}


func (this *InputAutomationModuleConfiguration) GetInputAutomationModuleInterface() IInputAutomationModuleFactoryInterface {
        return this.InputAutomationModuleInterface
}


func (this *InputAutomationModuleConfiguration) SetInputAutomationModuleInterface(inputAutomationModuleInterface IInputAutomationModuleFactoryInterface) {
        this.LogUtil.PutF("InputAutomationModuleFactoryInterface : " + fmt.Sprint(inputAutomationModuleInterface), this, "setInputAutomationModuleInterface")
        this.InputAutomationModuleInterface = inputAutomationModuleInterface
        this.SetName(this.InputAutomationModuleInterface.GetName())
}

