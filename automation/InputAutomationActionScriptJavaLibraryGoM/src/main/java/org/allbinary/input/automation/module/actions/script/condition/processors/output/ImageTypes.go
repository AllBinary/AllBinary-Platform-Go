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

type ImageTypes struct {
    Vector IBasicArrayList
}


func (this *ImageTypes) initFields() {
    this.Vector = NewBasicArrayListD()
}




// Java overloads of new ImageTypes
func NewImageTypes(args ...any) *ImageTypes {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewImageTypes_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewImageTypes_2(Java_Arg[*Node](args[0]))
        }
    }
    panic("No Java overload of new ImageTypes for the arguments")
}
func NewImageTypes_1() *ImageTypes {
    this := &ImageTypes{}
    this.initFields()
    return this
}


func NewImageTypes_2(node *Node) *ImageTypes {
    this := &ImageTypes{}
    this.initFields()
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var childNode *Node = nodeList.Item(index); _ = childNode
            if childNode.GetNodeName().CompareTo(ImageActionScriptOutputData_TYPE) == 0 {
                var nextImageTypeString string = DomNodeHelper_GetTextNodeValue(childNode); _ = nextImageTypeString
                this.GetVector().Add(InputImageType_GetInstance(nextImageTypeString))
            } else {
                panic(errors.New(fmt.Sprint("ImageTypes Unknown Node")))
            }
        }
    return this
}


func ImageTypes_ValueOf(anyType []any) *ImageTypes {
        var imageTypes *ImageTypes = NewImageTypes(); _ = imageTypes
        for index := int(0); index < len(anyType); index++ {
            var imageType *InputImageType = InputImageType_GetInstance(any(anyType[index]).(string)); _ = imageType
            imageTypes.GetVector().Add(imageType)
        }
        return imageTypes
}


func (this *ImageTypes) GetVector() IBasicArrayList {
        return this.Vector
}


// throws Exception
func (this *ImageTypes) ToXmlNode(document *Document) *Node {
        var newNode *Node = document.CreateElement(ImageActionScriptOutputData_TYPES); _ = newNode
        var size int = this.Vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var imageType *InputImageType = any(this.Vector.Get(index)).(*InputImageType); _ = imageType
            newNode.AppendChild(ModDomHelper_CreateTextNode(document, ImageActionScriptOutputData_TYPE, imageType.GetName()))
        }
        return newNode
}


func (this *ImageTypes) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("ImageTypes: ")
        var size int = this.Vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var imageType *InputImageType = any(this.Vector.Get(index)).(*InputImageType); _ = imageType
            stringBuffer.Append(imageType.GetName())
            if index < size - 1 {
                stringBuffer.Append(", ")
            }
        }
        return fmt.Sprint(stringBuffer)
}

func (this *ImageTypes) String() string {
    return this.ToString()
}

