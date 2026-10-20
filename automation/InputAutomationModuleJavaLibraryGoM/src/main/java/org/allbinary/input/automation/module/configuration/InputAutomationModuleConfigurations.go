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

type InputAutomationModuleConfigurations struct {
    LogUtil *LogUtil
    HashMap *HashMap
}


func (this *InputAutomationModuleConfigurations) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}




// Java overloads of new InputAutomationModuleConfigurations
func NewInputAutomationModuleConfigurations(args ...any) *InputAutomationModuleConfigurations {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*File](args[1], pass) {
            return NewInputAutomationModuleConfigurations_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*File](args[1]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[*Document](args[1], pass) {
            return NewInputAutomationModuleConfigurations_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[*Document](args[1]))
        }
        if len(args) == 1 && Java_IsArg[IList[*InputAutomationModuleConfiguration]](args[0], pass) {
            return NewInputAutomationModuleConfigurations_3(Java_Arg[IList[*InputAutomationModuleConfiguration]](args[0]))
        }
        if len(args) == 0 {
            return NewInputAutomationModuleConfigurations_4()
        }
    }
    panic("No Java overload of new InputAutomationModuleConfigurations for the arguments")
}
func NewInputAutomationModuleConfigurations_1(abeClientInformation AbeClientInformationInterface, file *File) *InputAutomationModuleConfigurations {
    this := &InputAutomationModuleConfigurations{}
    this.initFields()
        this.SetHashMap(NewHashMap())
        var bytes []int8 = make([]int8, 100000); _ = bytes
        var idFile *FileInputStream = NewFileInputStream(file); _ = idFile
        var length int = idFile.Read(bytes); _ = length
        var data string = string(bytes); _ = data
        var endIndex int = strings.LastIndex(data, string('>')); _ = endIndex
        var document *Document = DomDocumentHelper_Create(data[0:endIndex + 1]); _ = document
        var nodeList *NodeList = document.GetElementsByTagName(NewInputAutomationModulesData_NAME); _ = nodeList
        if nodeList.GetLength() > 0 {
            var nameNodeList *NodeList = document.GetElementsByTagName(InputAutomationModuleData_NAME); _ = nameNodeList
            this.LogUtil.PutF("Number Of Module(s) Specified: " + fmt.Sprint(nameNodeList.GetLength()), this, "Contructor")
            for index := int(0); index < nameNodeList.GetLength(); index++ {
                var node *Node = nameNodeList.Item(index); _ = node
                this.Add(NewInputAutomationModuleConfiguration(abeClientInformation, node))
            }
        }
    return this
}


func NewInputAutomationModuleConfigurations_2(abeClientInformation AbeClientInformationInterface, document *Document) *InputAutomationModuleConfigurations {
    this := &InputAutomationModuleConfigurations{}
    this.initFields()
        this.SetHashMap(NewHashMap())
        var nameNodeList *NodeList = document.GetElementsByTagName(InputAutomationModuleData_NAME); _ = nameNodeList
        this.LogUtil.PutF("Number Of Module(s) Specified: " + fmt.Sprint(nameNodeList.GetLength()), this, "Contructor")
        for index := int(0); index < nameNodeList.GetLength(); index++ {
            var node *Node = nameNodeList.Item(index); _ = node
            this.Add(NewInputAutomationModuleConfiguration(abeClientInformation, node))
        }
    return this
}


func NewInputAutomationModuleConfigurations_3(inputAutomationModuleConfigurationList IList[*InputAutomationModuleConfiguration]) *InputAutomationModuleConfigurations {
    this := &InputAutomationModuleConfigurations{}
    this.initFields()
        this.SetHashMap(NewHashMap())
        var size int = inputAutomationModuleConfigurationList.Size(); _ = size
        var inputAutomationModuleConfiguration *InputAutomationModuleConfiguration; _ = inputAutomationModuleConfiguration
        for index := int(0); index < size; index++ {
            inputAutomationModuleConfiguration = inputAutomationModuleConfigurationList.Get(index)
            this.Add(inputAutomationModuleConfiguration)
        }
    return this
}


func NewInputAutomationModuleConfigurations_4() *InputAutomationModuleConfigurations {
    this := &InputAutomationModuleConfigurations{}
    this.initFields()
        this.SetHashMap(NewHashMap())
    return this
}


func (this *InputAutomationModuleConfigurations) Add(inputAutomationModuleConfiguration *InputAutomationModuleConfiguration) {
        this.GetHashMap().Put(inputAutomationModuleConfiguration.GetClassName(), inputAutomationModuleConfiguration)
}


func (this *InputAutomationModuleConfigurations) Remove(inputAutomationModuleConfiguration *InputAutomationModuleConfiguration) {
        this.GetHashMap().Remove(inputAutomationModuleConfiguration.GetClassName())
}


func (this *InputAutomationModuleConfigurations) GetHashMap() *HashMap {
        return this.HashMap
}


func (this *InputAutomationModuleConfigurations) SetHashMap(hashMap *HashMap) {
        this.HashMap = hashMap
}

