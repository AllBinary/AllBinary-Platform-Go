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

type DomDocumentHelper struct {
}



func NewDomDocumentHelper() *DomDocumentHelper {
    this := &DomDocumentHelper{}
    return this
}


func DomDocumentHelper_Init() {
        System_SetProperty("jdk.xml.maxElementDepth", "0")
        System_SetProperty("jdk.xml.xpathExprOpLimit", "0")
        System_SetProperty("jdk.xml.xpathTotalOpLimit", "0")
        var LIMIT string = "500000"; _ = LIMIT
        System_SetProperty("jdk.xml.maxGeneralEntitySizeLimit", LIMIT)
        System_SetProperty("jdk.xml.totalEntitySizeLimit", LIMIT)
}


func DomDocumentHelper_Create() *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var factory *DocumentBuilderFactory = DocumentBuilderFactory_NewInstance(); _ = factory
            var builder *DocumentBuilder = factory.NewDocumentBuilder(); _ = builder
            var document *Document = builder.NewDocument(); _ = document
            tryControl1, tryResult1 = 1, document; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func DomDocumentHelper_Create(inputStream *InputStream) *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var factory *DocumentBuilderFactory = DocumentBuilderFactory_NewInstance(); _ = factory
            var builder *DocumentBuilder = factory.NewDocumentBuilder(); _ = builder
            var document *Document = builder.Parse(inputStream); _ = document
            tryControl1, tryResult1 = 1, document; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func DomDocumentHelper_Create(xmlString string) *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var factory *DocumentBuilderFactory = DocumentBuilderFactory_NewInstance(); _ = factory
            var builder *DocumentBuilder = factory.NewDocumentBuilder(); _ = builder
            var document *Document = builder.Parse(NewStringBufferInputStream(xmlString)); _ = document
            tryControl1, tryResult1 = 1, document; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

