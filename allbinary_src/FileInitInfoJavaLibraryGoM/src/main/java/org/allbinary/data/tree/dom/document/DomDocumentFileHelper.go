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

type DomDocumentFileHelper struct {
}



func NewDomDocumentFileHelper() *DomDocumentFileHelper {
    this := &DomDocumentFileHelper{}
    return this
}


// throws Exception
func DomDocumentFileHelper_CreateDocument(xmlFile *AbFile) *Document {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Document) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            tryControl1, tryResult1 = 1, DomDocumentHelper_Create(NewAbFileInputStream(xmlFile)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func DomDocumentFileHelper_Save(file *AbFile, document *Document) {
        var dataOutputStream *AbDataOutputStream = nil; _ = dataOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            StreamUtil_GetInstance().Close(dataOutputStream)
            }()
            var copyTransformerFactory *TransformerFactory = TransformerFactory_NewInstance(); _ = copyTransformerFactory
            var copyTransformer *Transformer = copyTransformerFactory.NewTransformer(); _ = copyTransformer
            var domSource *DOMSource = NewDOMSource(document); _ = domSource
            if file.IsFile() {
                file.Delete()
            }
            file.CreateNewFile()
            dataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(file)
            var streamResult *StreamResult = NewStreamResult(dataOutputStream); _ = streamResult
            copyTransformer.Transform(domSource, streamResult)
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}

