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

type LogConfig struct {
    IsEnabled bool
    Name string
    Description string
    Path string
    FileName string
}


func (this *LogConfig) initFields() {
    
    
    
    
    
}




// Java overloads of new LogConfig
func NewLogConfig(args ...any) *LogConfig {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return NewLogConfig_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewLogConfig_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new LogConfig for the arguments")
}
func NewLogConfig_1(aName string, aDescription string, aPath string, aFileName string) *LogConfig {
    this := &LogConfig{}
    this.initFields()
        this.IsEnabled = false
        this.Name = aName
        this.Description = aDescription
        this.Path = aPath
        this.FileName = aFileName
    return this
}


func NewLogConfig_2(node *Node) *LogConfig {
    this := &LogConfig{}
    this.initFields()
        var enabledValueNode *Node = DomSearchHelper_GetNode(EntryData_GetInstance().ENABLE, node.GetChildNodes()); _ = enabledValueNode
        var isEnabledString string = DomNodeHelper_GetTextNodeValue(enabledValueNode); _ = isEnabledString
        this.IsEnabled = BooleanUtil_GetInstance().GetFromString(isEnabledString)
        var logConfigData *LogConfigData = LogConfigData_GetInstance(); _ = logConfigData
        var nameValueNode *Node = DomSearchHelper_GetNode(logConfigData.NAME, node.GetChildNodes()); _ = nameValueNode
        this.Name = DomNodeHelper_GetTextNodeValue(nameValueNode)
        var descriptionValueNode *Node = DomSearchHelper_GetNode(logConfigData.DESCRIPTION, node.GetChildNodes()); _ = descriptionValueNode
        this.Description = DomNodeHelper_GetTextNodeValue(descriptionValueNode)
        var fileValueNode *Node = DomSearchHelper_GetNode(logConfigData.FILE, node.GetChildNodes()); _ = fileValueNode
        this.FileName = DomNodeHelper_GetTextNodeValue(fileValueNode)
    return this
}


func (this *LogConfig) GetName() string {
        return this.Name
}


func (this *LogConfig) IsEnabled() bool {
        return this.IsEnabled
}


func (this *LogConfig) Enable() {
        this.IsEnabled = true
}


func (this *LogConfig) Disable() {
        this.IsEnabled = false
}


func (this *LogConfig) GetDescription() string {
        return this.Description
}


func (this *LogConfig) GetFileName() string {
        return this.FileName
}


func (this *LogConfig) SetName(value string) {
        this.Name = value
}


func (this *LogConfig) SetDescription(value string) {
        this.Description = value
}


func (this *LogConfig) SetPath(value string) {
        this.Path = value
}


func (this *LogConfig) SetFileName(value string) {
        this.FileName = value
}


// throws Exception
func (this *LogConfig) GetTypeVector() IBasicArrayList {
        var logKeyVector IBasicArrayList = NewBasicArrayListD(); _ = logKeyVector
        return logKeyVector
}

