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

type TimeIntervalActionScriptCondition struct {
    BasicProfileActionScriptCondition
    LogUtil *LogUtil
    TimeIntervalActionScriptConditionJPanel *TimeIntervalActionScriptConditionJPanel
    TimeHelper ITimeDelayHelper
}

// Implements TimeIntervalActionScriptConditionInterface

var TimeIntervalActionScriptCondition_NAME string = "Time Interval"

func (this *TimeIntervalActionScriptCondition) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new TimeIntervalActionScriptCondition
func NewTimeIntervalActionScriptCondition(args ...any) *TimeIntervalActionScriptCondition {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewTimeIntervalActionScriptCondition_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewTimeIntervalActionScriptCondition_2()
        }
    }
    panic("No Java overload of new TimeIntervalActionScriptCondition for the arguments")
}
func NewTimeIntervalActionScriptCondition_1(node *Node) *TimeIntervalActionScriptCondition {
    this := &TimeIntervalActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(TimeIntervalActionScriptCondition_NAME, node)
        this.initFields()
        var actionNode *Node = DomSearchHelper_GetNode(TimeIntervalActionScriptConditionData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().CompareTo(TimeIntervalActionScriptConditionData_TIME) == 0 {
                    var interval string = DomNodeHelper_GetTextNodeValue(childNode); _ = interval
                    this.SetTimeDelayHelper(NewTimeDelayHelper(fmt.Sprint(interval)))
                } else {
                    panic(errors.New(fmt.Sprint("Time Interval Action Script Condition Unknown Node")))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Time Interval Action Script Condition Node Null")))
        }
        this.Init()
    return this
}


func NewTimeIntervalActionScriptCondition_2() *TimeIntervalActionScriptCondition {
    this := &TimeIntervalActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(TimeIntervalActionScriptCondition_NAME)
        this.initFields()
        this.SetTimeDelayHelper(NewTimeDelayHelper(0))
        this.Init()
    return this
}


func (this *TimeIntervalActionScriptCondition) Init() {
        this.TimeIntervalActionScriptConditionJPanel = NewTimeIntervalActionScriptConditionJPanel(this)
}


func (this *TimeIntervalActionScriptCondition) ShowDialog() {
        this.TimeIntervalActionScriptConditionJPanel.GetTimeIntervalActionJDialog().SetVisible(true)
}


func (this *TimeIntervalActionScriptCondition) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(TimeIntervalActionScriptConditionData_TIME, fmt.Sprint(this.TimeHelper.TimeDelayHelper_().Delay))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *TimeIntervalActionScriptCondition) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptCondition.ToXmlNode(document); _ = node
        node.AppendChild(ModDomHelper_CreateNodeWithValueNodes(document, TimeIntervalActionScriptConditionData_NAME, this.ToHashMap()))
        return node
}


func (this *TimeIntervalActionScriptCondition) Log() {
        this.LogUtil.PutF("Time Interval: " + fmt.Sprint(this.TimeHelper.TimeDelayHelper_().Delay), this, "log")
}


func (this *TimeIntervalActionScriptCondition) GetTimeDelayHelper() ITimeDelayHelper {
        return this.TimeHelper
}


func (this *TimeIntervalActionScriptCondition) SetTimeDelayHelper(timeHelper ITimeDelayHelper) {
        this.TimeHelper = timeHelper
}


// throws Exception
func (this *TimeIntervalActionScriptCondition) ShouldProcess(frame int64) bool {
        var timeHelper ITimeDelayHelper = this.GetTimeDelayHelper(); _ = timeHelper
        if timeHelper.IsTimeTNT() {
            timeHelper.SetStartTimeTNT()
            return this.BasicProfileActionScriptCondition.ShouldProcess(frame)
        }
        return false
}

