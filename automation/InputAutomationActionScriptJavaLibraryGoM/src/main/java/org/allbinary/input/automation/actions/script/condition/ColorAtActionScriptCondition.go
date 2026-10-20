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

type ColorAtActionScriptCondition struct {
    BasicProfileActionScriptCondition
    LogUtil *LogUtil
    ColorAtActionScriptConditionJPanel *ColorAtActionScriptConditionJPanel
    Point *Point
    ColorRangeInterface ColorRangeInterface
}

// Implements ColorAtActionScriptConditionInterface

var ColorAtActionScriptCondition_NAME string = "If Color At"

func (this *ColorAtActionScriptCondition) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    this.ColorRangeInterface = any(NewColorRange()).(ColorRangeInterface)
}




// Java overloads of new ColorAtActionScriptCondition
func NewColorAtActionScriptCondition(args ...any) *ColorAtActionScriptCondition {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewColorAtActionScriptCondition_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewColorAtActionScriptCondition_2()
        }
    }
    panic("No Java overload of new ColorAtActionScriptCondition for the arguments")
}
func NewColorAtActionScriptCondition_1(node *Node) *ColorAtActionScriptCondition {
    this := &ColorAtActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(ColorAtActionScriptCondition_NAME, node)
        this.initFields()
        this.Point = NewPoint(0, 0)
        var actionNode *Node = DomSearchHelper_GetNode(ColorAtActionScriptConditionData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_LOCATION_X) == 0 {
                    var locationX string = DomNodeHelper_GetTextNodeValue(childNode); _ = locationX
                    this.Point.X = fmt.Sprint(locationX)
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_LOCATION_Y) == 0 {
                    var locationY string = DomNodeHelper_GetTextNodeValue(childNode); _ = locationY
                    this.Point.Y = fmt.Sprint(locationY)
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MIN_RED) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMinRed(fmt.Sprint(string_))
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MAX_RED) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMaxRed(fmt.Sprint(string_))
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MIN_GREEN) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMinGreen(fmt.Sprint(string_))
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MAX_GREEN) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMaxGreen(fmt.Sprint(string_))
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MIN_BLUE) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMinBlue(fmt.Sprint(string_))
                } else if childNode.GetNodeName().CompareTo(ColorAtActionScriptConditionData_MAX_BLUE) == 0 {
                    var string_ string = DomNodeHelper_GetTextNodeValue(childNode); _ = string_
                    this.GetColorRangeInterface().SetMaxBlue(fmt.Sprint(string_))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Color At Action Script Condition Node Null")))
        }
        this.Init()
    return this
}


func NewColorAtActionScriptCondition_2() *ColorAtActionScriptCondition {
    this := &ColorAtActionScriptCondition{}
        this.BasicProfileActionScriptCondition = *NewBasicProfileActionScriptCondition(ColorAtActionScriptCondition_NAME)
        this.initFields()
        this.Point = NewPoint(0, 0)
        this.Init()
    return this
}


func (this *ColorAtActionScriptCondition) Init() {
        this.ColorAtActionScriptConditionJPanel = NewColorAtActionScriptConditionJPanel(this)
}


func (this *ColorAtActionScriptCondition) GetPoint() *Point {
        return this.Point
}


func (this *ColorAtActionScriptCondition) SetPoint(point *Point) {
        this.Point = point
}


func (this *ColorAtActionScriptCondition) ShowDialog() {
        this.ColorAtActionScriptConditionJPanel.GetColorAtActionJDialog().SetVisible(true)
}


func (this *ColorAtActionScriptCondition) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(ColorAtActionScriptConditionData_LOCATION_X, fmt.Sprint(this.GetPoint().X))
        hashMap.Put(ColorAtActionScriptConditionData_LOCATION_Y, fmt.Sprint(this.GetPoint().Y))
        hashMap.Put(ColorAtActionScriptConditionData_MIN_RED, fmt.Sprint(this.GetColorRangeInterface().GetMinRed()))
        hashMap.Put(ColorAtActionScriptConditionData_MAX_RED, fmt.Sprint(this.GetColorRangeInterface().GetMaxRed()))
        hashMap.Put(ColorAtActionScriptConditionData_MIN_GREEN, fmt.Sprint(this.GetColorRangeInterface().GetMinGreen()))
        hashMap.Put(ColorAtActionScriptConditionData_MAX_GREEN, fmt.Sprint(this.GetColorRangeInterface().GetMaxGreen()))
        hashMap.Put(ColorAtActionScriptConditionData_MIN_BLUE, fmt.Sprint(this.GetColorRangeInterface().GetMinBlue()))
        hashMap.Put(ColorAtActionScriptConditionData_MAX_BLUE, fmt.Sprint(this.GetColorRangeInterface().GetMaxBlue()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *ColorAtActionScriptCondition) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptCondition.ToXmlNode(document); _ = node
        node.AppendChild(ModDomHelper_CreateNodeWithValueNodes(document, ColorAtActionScriptConditionData_NAME, this.ToHashMap()))
        return node
}


// throws Exception
func (this *ColorAtActionScriptCondition) ShouldProcess(frame int64) bool {
        var capturedBufferedImageCacheable *BufferedImageFrameCacheable = any((any(CapturedBufferedImagesCacheSingleton_GetInstance()).(AutomaticCacheInterface)).Get(fmt.Sprint(frame))).(*BufferedImageFrameCacheable); _ = capturedBufferedImageCacheable
        var bufferedImage *BufferedImage = capturedBufferedImageCacheable.GetBufferedImage(); _ = bufferedImage
        var automaticCacheInterface AutomaticCacheInterface = ColorCacheFactory_GetInstance(); _ = automaticCacheInterface
        var colorInteger int = fmt.Sprint(bufferedImage.GetRGB(this.GetPoint().X, this.GetPoint().Y)); _ = colorInteger
        var colorCacheable *ColorCacheable = any(automaticCacheInterface.Get(colorInteger)).(*ColorCacheable); _ = colorCacheable
        var color *Color = colorCacheable.GetColor(); _ = color
        if this.ColorRangeInterface.IsInRange(color) {
            return this.BasicProfileActionScriptCondition.ShouldProcess(frame)
        }
        return false
}


func (this *ColorAtActionScriptCondition) Log() {
        var message string = fmt.Sprint(NewStringMaker().Append("getPoint(): ").Append(fmt.Sprint(this.GetPoint())).Append(CommonSeps_GetInstance().NEW_LINE).Append(fmt.Sprint(this.GetColorRangeInterface()))); _ = message
        this.LogUtil.PutF(message, this, "log")
}


func (this *ColorAtActionScriptCondition) GetColorRangeInterface() ColorRangeInterface {
        return this.ColorRangeInterface
}


func (this *ColorAtActionScriptCondition) SetColorRangeInterface(colorRangeInterface ColorRangeInterface) {
        this.ColorRangeInterface = colorRangeInterface
}

