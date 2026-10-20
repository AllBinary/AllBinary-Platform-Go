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

//import org.allbinary.data.tree.dom.document.DomDocumentHelper;
type XslHelper struct {
}

var XslHelper_Instance *XslHelper = NewXslHelper()


func XslHelper_GetInstance() *XslHelper {
        return XslHelper_Instance
}


func NewXslHelper() *XslHelper {
    this := &XslHelper{}
    return this
}


// throws Exception
func (this *XslHelper) Translate(xsltStreamSource *StreamSource, xmlStreamSource *StreamSource, streamResult *StreamResult) *OutputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *OutputStream) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var transformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = transformerFactory
            var transformer *Transformer = transformerFactory.NewTransformer(xsltStreamSource); _ = transformer
            transformer.Transform(xmlStreamSource, streamResult)
            tryControl1, tryResult1 = 1, streamResult.GetOutputStream(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *XslHelper) Translate(resolver *URIResolver, xsltStreamSource *StreamSource, xmlStreamSource *StreamSource, streamResult *StreamResult) *OutputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *OutputStream) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var tFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = tFactory
            tFactory.SetURIResolver(resolver)
            var transformer *Transformer = tFactory.NewTransformer(xsltStreamSource); _ = transformer
            transformer.Transform(xmlStreamSource, streamResult)
            tryControl1, tryResult1 = 1, streamResult.GetOutputStream(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *XslHelper) Translate(xsltStreamSource *StreamSource, xmlDocument *Document) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(xsltStreamSource, NewStreamSource(NewStringBufferInputStream(XmlDocumentHelper_ToString(xmlDocument))), NewStreamResult(outputStream))); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *XslHelper) Translate(xsltStreamSource *StreamSource, xmlStreamSource *StreamSource) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(xsltStreamSource, xmlStreamSource, NewStreamResult(outputStream))); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *XslHelper) Translate(resolver *URIResolver, xsltStreamSource *StreamSource, xmlStreamSource *StreamSource) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(resolver, xsltStreamSource, xmlStreamSource, NewStreamResult(outputStream))); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *XslHelper) Export(outputFile *File, xsltFilePath string, xmlDocument *Document) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            outputFile.CreateNewFile()
            this.Translate(NewStreamSource(xsltFilePath), NewStreamSource(NewStringBufferInputStream(XmlDocumentHelper_ToString(xmlDocument))), NewStreamResult(outputFile))
            return
        }(); tryControl1 == 1 {
            return
        }
}

