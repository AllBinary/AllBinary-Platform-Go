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

type TransformDocument struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    BaseNode *Node
    Document *Document
}

// Implements TransformDocumentInterface


func (this *TransformDocument) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewTransformDocument() *TransformDocument {
    this := &TransformDocument{}
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
            this.Document = DomDocumentHelper_Create()
            var allbinaryNode *Node = this.Document.CreateElement("allbinary"); _ = allbinaryNode
            this.BaseNode = allbinaryNode
            this.Document.AppendChild(this.BaseNode)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.PutF(this.Log(), this, this.CommonStrings.CONSTRUCTOR)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *TransformDocument) GetBaseNode() *Node {
        return this.BaseNode
}


func (this *TransformDocument) GetDoc() *Document {
        return this.Document
}


// throws Exception
func (this *TransformDocument) Log() string {
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

