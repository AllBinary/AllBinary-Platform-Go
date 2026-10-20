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


func NewLogConfig(aName string, aDescription string, aPath string, aFileName string) *LogConfig {
    this := &LogConfig{}
    this.initFields()
        this.IsEnabled = false
        this.Name = aName
        this.Description = aDescription
        this.Path = aPath
        this.FileName = aFileName
    return this
}


// Java overloaded constructor: Go needs a different name
func NewLogConfig(node *Node) *LogConfig {
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
func (this *LogConfig) GetTypeVector() *BasicArrayList {
        var logKeyVector *BasicArrayList = NewBasicArrayListD(); _ = logKeyVector
        return logKeyVector
}

