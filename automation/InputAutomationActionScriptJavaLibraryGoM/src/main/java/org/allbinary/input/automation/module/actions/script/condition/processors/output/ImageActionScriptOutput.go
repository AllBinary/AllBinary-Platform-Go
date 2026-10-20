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

type ImageActionScriptOutput struct {
    BasicProfileActionScriptOutput
    LogUtil *LogUtil
    Saved bool
    Display bool
    ImageTypes *ImageTypes
    ActionScriptOutputJPanel *ImageActionScriptOutputJPanel
}

// Implements ImageActionScriptOutputInterface

var ImageActionScriptOutput_NAME string = "Captures"

func (this *ImageActionScriptOutput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Saved = false
    this.Display = false
    
    
}




// Java overloads of new ImageActionScriptOutput
func NewImageActionScriptOutput(args ...any) *ImageActionScriptOutput {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewImageActionScriptOutput_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewImageActionScriptOutput_2()
        }
    }
    panic("No Java overload of new ImageActionScriptOutput for the arguments")
}
func NewImageActionScriptOutput_1(node *Node) *ImageActionScriptOutput {
    this := &ImageActionScriptOutput{}
        this.BasicProfileActionScriptOutput = *NewBasicProfileActionScriptOutput(ImageActionScriptOutput_NAME, node)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        var actionNode *Node = DomSearchHelper_GetNode(ImageActionScriptOutputData_NAME, node.GetChildNodes()); _ = actionNode
        if actionNode != nil {
            var nodeList *NodeList = actionNode.GetChildNodes(); _ = nodeList
            for index := int(0); index < nodeList.GetLength(); index++ {
                var childNode *Node = nodeList.Item(index); _ = childNode
                if childNode.GetNodeName().CompareTo(ImageActionScriptOutputData_SAVE) == 0 {
                    var booleanString string = DomNodeHelper_GetTextNodeValue(childNode); _ = booleanString
                    this.SetSaved(booleanString)
                } else if childNode.GetNodeName().CompareTo(ImageActionScriptOutputData_DISPLAY) == 0 {
                    var booleanString string = DomNodeHelper_GetTextNodeValue(childNode); _ = booleanString
                    this.SetSaved(booleanString)
                } else if childNode.GetNodeName().CompareTo(ImageActionScriptOutputData_TYPES) == 0 {
                    this.SetImageTypes(NewImageTypes(childNode))
                } else {
                    panic(errors.New(fmt.Sprint("Action Script Output Unknown Node")))
                }
            }
        } else {
            panic(errors.New(fmt.Sprint("Action Script Output Node Null")))
        }
        this.SetAllowsChildren(false)
        this.ActionScriptOutputJPanel = NewImageActionScriptOutputJPanel(this)
    return this
}


func NewImageActionScriptOutput_2() *ImageActionScriptOutput {
    this := &ImageActionScriptOutput{}
        this.BasicProfileActionScriptOutput = *NewBasicProfileActionScriptOutput(ImageActionScriptOutput_NAME)
        this.initFields()
        this.SetImageTypes(NewImageTypes())
        this.SetAllowsChildren(false)
        this.ActionScriptOutputJPanel = NewImageActionScriptOutputJPanel(this)
    return this
}


func (this *ImageActionScriptOutput) GetActionScriptOutputJPanel() *ImageActionScriptOutputJPanel {
        return this.ActionScriptOutputJPanel
}


func (this *ImageActionScriptOutput) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(ImageActionScriptOutputData_DISPLAY, fmt.Sprint(this.IsDisplay()))
        hashMap.Put(ImageActionScriptOutputData_SAVE, fmt.Sprint(this.IsSaved()))
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *ImageActionScriptOutput) ToXmlNode(document *Document) *Node {
        var node *Node = this.BasicProfileActionScriptOutput.ToXmlNode(document); _ = node
        var newNode *Node = ModDomHelper_CreateNodeWithValueNodes(document, ImageActionScriptOutputData_NAME, this.ToHashMap()); _ = newNode
        newNode.AppendChild(this.GetImageTypes().ToXmlNode(document))
        node.AppendChild(newNode)
        return node
}


// throws Exception
func (this *ImageActionScriptOutput) Process(frame int64) {
        ImageActionScriptOutputProcessor_Process(this, frame)
}


func (this *ImageActionScriptOutput) ShowDialog() {
        this.ActionScriptOutputJPanel.GetCapturedImageActionJDialog().SetVisible(true)
}


func (this *ImageActionScriptOutput) IsSaved() bool {
        return this.Saved
}


func (this *ImageActionScriptOutput) SetSaved(saved bool) {
        this.Saved = saved
}


func (this *ImageActionScriptOutput) IsDisplay() bool {
        return this.Display
}


func (this *ImageActionScriptOutput) SetDisplay(display bool) {
        this.Display = display
}


func (this *ImageActionScriptOutput) Future_toString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(" Is Save: ")
        stringBuffer.Appendboolean(this.IsSaved())
        stringBuffer.Append(" Is Display: ")
        stringBuffer.Appendboolean(this.IsDisplay())
        stringBuffer.Append(" ")
        stringBuffer.Append(fmt.Sprint(this.GetImageTypes()))
        return fmt.Sprint(stringBuffer)
}


func (this *ImageActionScriptOutput) Log() {
        this.LogUtil.PutF(this.Future_toString(), this, "log")
}


func (this *ImageActionScriptOutput) GetImageTypes() *ImageTypes {
        return this.ImageTypes
}


func (this *ImageActionScriptOutput) SetImageTypes(imageTypes *ImageTypes) {
        this.ImageTypes = imageTypes
}

