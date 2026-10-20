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

type DirectoryDataModel struct {
    FileBasicArrayList IBasicArrayList
    TransformDocumentInterface TransformDocumentInterface
    NO_TYPE int
}

// Implements DomNodeInterface, TransformInterface


func (this *DirectoryDataModel) initFields() {
    
    
    this.NO_TYPE = 0
}


func NewDirectoryDataModel(file *AbFile) *DirectoryDataModel {
    this := &DirectoryDataModel{}
    this.initFields()
        this.FileBasicArrayList = Directory_GetInstance().Search(file)
        this.TransformDocumentInterface = any(TransformDocumentFactory_GetInstance()).(TransformDocumentInterface)
    return this
}


func (this *DirectoryDataModel) GetTypeId() int {
        return this.NO_TYPE
}


// throws Exception
func (this *DirectoryDataModel) GetTransformInfoInterface() TransformInfoInterface {
        return nil
}


func (this *DirectoryDataModel) SetTransformDocumentInterface(transformDocumentInterface TransformDocumentInterface) {
        this.TransformDocumentInterface = transformDocumentInterface
}


func (this *DirectoryDataModel) GetTransformDocumentInterface() TransformDocumentInterface {
        return this.TransformDocumentInterface
}


// throws Exception
func (this *DirectoryDataModel) ToXmlNode(document *Document) *Node {
        var dirNode *Node = document.CreateElement("dir"); _ = dirNode
        var size int = this.FileBasicArrayList.Size(); _ = size
        var nextFile *AbFile; _ = nextFile
        var fileNode *Node; _ = fileNode
        for index := int(0); index < size; index++ {
            nextFile = any(this.FileBasicArrayList.Get(index)).(*AbFile)
            fileNode = NewFileDomDataModel(nextFile).ToXmlNode(document)
            dirNode.AppendChild(fileNode)
        }
        return dirNode
}


// throws Exception
func (this *DirectoryDataModel) ToXmlDoc() *Document {
        return this.GetTransformDocumentInterface().GetDoc()
}


// throws Exception
func (this *DirectoryDataModel) View() string {
        var node *Node = this.ToXmlNode(this.GetTransformDocumentInterface().GetDoc()); _ = node
        this.GetTransformDocumentInterface().GetBaseNode().AppendChild(node)
        return XmlDocumentHelper_ToString(this.GetTransformDocumentInterface().GetDoc())
}

