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
func (this *XslHelper) Translate_1(xsltStreamSource *StreamSource, xmlStreamSource *StreamSource, streamResult *StreamResult) IOutputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IOutputStream) {
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

// Java overloads of XslHelper.translate
func (this *XslHelper) Translate(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[*StreamSource](args[0], pass) && Java_IsArg[*StreamSource](args[1], pass) && Java_IsArg[*StreamResult](args[2], pass) {
            return this.Translate_1(Java_Arg[*StreamSource](args[0]), Java_Arg[*StreamSource](args[1]), Java_Arg[*StreamResult](args[2]))
        }
        if len(args) == 4 && Java_IsArg[*URIResolver](args[0], pass) && Java_IsArg[*StreamSource](args[1], pass) && Java_IsArg[*StreamSource](args[2], pass) && Java_IsArg[*StreamResult](args[3], pass) {
            return this.Translate_2(Java_Arg[*URIResolver](args[0]), Java_Arg[*StreamSource](args[1]), Java_Arg[*StreamSource](args[2]), Java_Arg[*StreamResult](args[3]))
        }
        if len(args) == 2 && Java_IsArg[*StreamSource](args[0], pass) && Java_IsArg[*Document](args[1], pass) {
            return this.Translate_3(Java_Arg[*StreamSource](args[0]), Java_Arg[*Document](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*StreamSource](args[0], pass) && Java_IsArg[*StreamSource](args[1], pass) {
            return this.Translate_4(Java_Arg[*StreamSource](args[0]), Java_Arg[*StreamSource](args[1]))
        }
        if len(args) == 3 && Java_IsArg[*URIResolver](args[0], pass) && Java_IsArg[*StreamSource](args[1], pass) && Java_IsArg[*StreamSource](args[2], pass) {
            return this.Translate_5(Java_Arg[*URIResolver](args[0]), Java_Arg[*StreamSource](args[1]), Java_Arg[*StreamSource](args[2]))
        }
    }
    panic("No Java overload of XslHelper.translate for the arguments")
}


// throws Exception
func (this *XslHelper) Translate_2(resolver *URIResolver, xsltStreamSource *StreamSource, xmlStreamSource *StreamSource, streamResult *StreamResult) IOutputStream {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IOutputStream) {
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
func (this *XslHelper) Translate_3(xsltStreamSource *StreamSource, xmlDocument *Document) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(xsltStreamSource, NewStreamSource(NewStringBufferInputStream(XmlDocumentHelper_ToString(xmlDocument))), NewStreamResult(outputStream)).(IOutputStream)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *XslHelper) Translate_4(xsltStreamSource *StreamSource, xmlStreamSource *StreamSource) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(xsltStreamSource, xmlStreamSource, NewStreamResult(outputStream)).(IOutputStream)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *XslHelper) Translate_5(resolver *URIResolver, xsltStreamSource *StreamSource, xmlStreamSource *StreamSource) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(this.Translate(resolver, xsltStreamSource, xmlStreamSource, NewStreamResult(outputStream)).(IOutputStream)); return
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
            this.Translate(NewStreamSource(xsltFilePath), NewStreamSource(NewStringBufferInputStream(XmlDocumentHelper_ToString(xmlDocument))), NewStreamResult(outputFile)).(IOutputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
}

