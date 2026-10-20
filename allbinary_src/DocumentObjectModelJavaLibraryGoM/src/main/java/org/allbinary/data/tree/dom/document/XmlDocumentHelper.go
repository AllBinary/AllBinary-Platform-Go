/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type XmlDocumentHelper struct {
    LogUtil *LogUtil
    INDENT_NUMBER string
    Transformer *Transformer
}

var XmlDocumentHelper_Instance *XmlDocumentHelper = NewXmlDocumentHelper()

func (this *XmlDocumentHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.INDENT_NUMBER = "indent-number"
    
}


func XmlDocumentHelper_GetInstance() *XmlDocumentHelper {
        return XmlDocumentHelper_Instance
}


func NewXmlDocumentHelper() *XmlDocumentHelper {
    this := &XmlDocumentHelper{}
    this.initFields()
        this.Init(4, true)
    return this
}


func (this *XmlDocumentHelper) Init(indent int, ignoreDeclaration bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
                }
            }()
            var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
            var transformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = transformerFactory
            transformerFactory.SetAttribute(this.INDENT_NUMBER, indent)
            this.Transformer = transformerFactory.NewTransformer()
            this.Transformer.SetOutputProperty(OutputKeys_ENCODING, CharacterSetData_GetInstance().UTF_8)
            this.Transformer.SetOutputProperty(OutputKeys_OMIT_XML_DECLARATION, func() any { if ignoreDeclaration { return booleanFactory.YES }; return booleanFactory.NO }())
            this.Transformer.SetOutputProperty(OutputKeys_INDENT, booleanFactory.YES)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *XmlDocumentHelper) Format(xmlString string) string {
        var src *InputSource = NewInputSource(NewStringReader(xmlString)); _ = src
        var document *Document = DocumentBuilderFactory_NewInstance().NewDocumentBuilder().Parse(src); _ = document
        var out *Writer = NewStringWriter(); _ = out
        this.Transformer.Transform(NewDOMSource(document), NewStreamResult(out))
        return fmt.Sprint(out)
}


// throws Exception
func XmlDocumentHelper_ToString(document *Document) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var domSource *DOMSource = NewDOMSource(document); _ = domSource
            var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
            var streamResult *StreamResult = NewStreamResult(byteArrayOutputStream); _ = streamResult
            var copyTransformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = copyTransformerFactory
            var copyTransformer *Transformer = copyTransformerFactory.NewTransformer(); _ = copyTransformer
            copyTransformer.Transform(domSource, streamResult)
            tryControl1, tryResult1 = 1, byteArrayOutputStream.String(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

