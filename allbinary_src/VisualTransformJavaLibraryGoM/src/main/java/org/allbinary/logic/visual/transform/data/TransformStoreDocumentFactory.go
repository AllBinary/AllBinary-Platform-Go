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

type TransformStoreDocumentFactory struct {
}



func NewTransformStoreDocumentFactory() *TransformStoreDocumentFactory {
    this := &TransformStoreDocumentFactory{}
    return this
}


// throws Exception
func TransformStoreDocumentFactory_GetInstance_1(pageContext *PageContext, weblisketSession *WeblisketSession) TransformDocumentInterface {
        return any(NewTransformStoreDocument(pageContext, weblisketSession)).(TransformDocumentInterface)
}

// Java overloads of TransformStoreDocumentFactory.getInstance
func TransformStoreDocumentFactory_GetInstance(args ...any) TransformDocumentInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*PageContext](args[0], pass) && Java_IsArg[*WeblisketSession](args[1], pass) {
            return TransformStoreDocumentFactory_GetInstance_1(Java_Arg[*PageContext](args[0]), Java_Arg[*WeblisketSession](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*SearchRequest](args[0], pass) {
            return TransformStoreDocumentFactory_GetInstance_2(Java_Arg[*SearchRequest](args[0]))
        }
    }
    panic("No Java overload of TransformStoreDocumentFactory.getInstance for the arguments")
}


// throws Exception
func TransformStoreDocumentFactory_GetInstance_2(searchRequest *SearchRequest) TransformDocumentInterface {
        return any(NewTransformStoreDocument(searchRequest)).(TransformDocumentInterface)
}

