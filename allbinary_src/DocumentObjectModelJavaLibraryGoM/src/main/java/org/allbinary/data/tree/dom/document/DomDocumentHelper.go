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


func DomDocumentHelper_Create_1() *Document {
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

// Java overloads of DomDocumentHelper.create
func DomDocumentHelper_Create(args ...any) *Document {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return DomDocumentHelper_Create_1()
        }
        if len(args) == 1 && Java_IsArg[IInputStream](args[0], pass) {
            return DomDocumentHelper_Create_2(Java_Arg[IInputStream](args[0]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return DomDocumentHelper_Create_3(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of DomDocumentHelper.create for the arguments")
}


// throws Exception
func DomDocumentHelper_Create_2(inputStream IInputStream) *Document {
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
func DomDocumentHelper_Create_3(xmlString string) *Document {
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

