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

type TransformHttpRequestDocument struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BaseNode *Node
    Document *Document
}

// Implements TransformDocumentInterface

// The Java class TransformHttpRequestDocument has subclasses so its variables are this interface
type ITransformHttpRequestDocument interface {
    TransformHttpRequestDocument_() *TransformHttpRequestDocument
    GetBaseNode() *Node
    GetDoc() *Document
    Log() string
}

func (this *TransformHttpRequestDocument) TransformHttpRequestDocument_() *TransformHttpRequestDocument {
    return this
}

var TransformHttpRequestDocument_Language string = "en"
var TransformHttpRequestDocument_Local string = "US"

func (this *TransformHttpRequestDocument) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}




// Java overloads of new TransformHttpRequestDocument
func NewTransformHttpRequestDocument(args ...any) *TransformHttpRequestDocument {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*PageContext](args[0], pass) && Java_IsArg[*WeblisketSession](args[1], pass) {
            return NewTransformHttpRequestDocument_1(Java_Arg[*PageContext](args[0]), Java_Arg[*WeblisketSession](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*SearchRequest](args[0], pass) {
            return NewTransformHttpRequestDocument_2(Java_Arg[*SearchRequest](args[0]))
        }
    }
    panic("No Java overload of new TransformHttpRequestDocument for the arguments")
}
func NewTransformHttpRequestDocument_1(pageContext *PageContext, weblisketSession *WeblisketSession) *TransformHttpRequestDocument {
    this := &TransformHttpRequestDocument{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            var contentType string = AcceptableResponseGenerator_GetInstance().Get(any(pageContext.GetRequest()).(*HttpServletRequest)); _ = contentType
            this.Document = DomDocumentHelper_Create()
            var contentTypeNode *Node = this.Document.CreateElement(contentType); _ = contentTypeNode
            var languageTypeNode *Node = this.Document.CreateElement(TransformHttpRequestDocument_Language); _ = languageTypeNode
            var localTypeNode *Node = this.Document.CreateElement(TransformHttpRequestDocument_Local); _ = localTypeNode
            this.BaseNode = localTypeNode
            languageTypeNode.AppendChild(localTypeNode)
            contentTypeNode.AppendChild(languageTypeNode)
            this.Document.AppendChild(contentTypeNode)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.PutF(this.Log(), this, this.CommonStrings.CONSTRUCTOR)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewTransformHttpRequestDocument_2(searchRequest *SearchRequest) *TransformHttpRequestDocument {
    this := &TransformHttpRequestDocument{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            var contentType string = searchRequest.GetContentType(); _ = contentType
            this.Document = DomDocumentHelper_Create()
            var contentTypeNode *Node = this.Document.CreateElement(contentType); _ = contentTypeNode
            var languageTypeNode *Node = this.Document.CreateElement(TransformHttpRequestDocument_Language); _ = languageTypeNode
            var localTypeNode *Node = this.Document.CreateElement(TransformHttpRequestDocument_Local); _ = localTypeNode
            this.BaseNode = localTypeNode
            languageTypeNode.AppendChild(localTypeNode)
            contentTypeNode.AppendChild(languageTypeNode)
            this.Document.AppendChild(contentTypeNode)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *TransformHttpRequestDocument) GetBaseNode() *Node {
        return this.BaseNode
}


func (this *TransformHttpRequestDocument) GetDoc() *Document {
        return this.Document
}


// throws Exception
func (this *TransformHttpRequestDocument) Log() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if this.BaseNode != nil {
            stringBuffer.Append("BaseNode: ")
            stringBuffer.Append(this.BaseNode.GetNodeName())
        } else {
            stringBuffer.Append("Log-Error: BaseNode is Null")
        }
        stringBuffer.Append("\nDocument: ")
        stringBuffer.Append(XmlDocumentHelper_ToString(this.Document))
        return fmt.Sprint(stringBuffer)
}

