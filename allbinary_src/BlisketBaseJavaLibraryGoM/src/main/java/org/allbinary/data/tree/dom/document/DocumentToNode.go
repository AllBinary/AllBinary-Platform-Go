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

type DocumentToNode struct {
}

var DocumentToNode_XMLROOT string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
var DocumentToNode_XMLROOTSTART string = "<?xml"
var DocumentToNode_XMLROOTEND string = "?>"
var DocumentToNode_XMLJSPROOTSTART string = "<jsp:root"
var DocumentToNode_XMLEND string = ">"
var DocumentToNode_XMLJSPROOTENDSTART string = "</jsp:root"


func NewDocumentToNode() *DocumentToNode {
    this := &DocumentToNode{}
    return this
}


// throws Exception
func DocumentToNode_ConvertDocumentToNodeString(documentString string) string {
        documentString = DocumentToNode_Remove(documentString, DocumentToNode_XMLJSPROOTSTART, DocumentToNode_XMLEND)
        documentString = DocumentToNode_Remove(documentString, DocumentToNode_XMLJSPROOTENDSTART, DocumentToNode_XMLEND)
        var replaceXmlRoot *Replace = NewReplace(DocumentToNode_XMLROOT, StringUtil_GetInstance().EMPTY_STRING); _ = replaceXmlRoot
        return replaceXmlRoot.All(documentString)
}


func DocumentToNode_RemoveXmlDefinitionFromDocumentString(documentString string) string {
        return DocumentToNode_Remove(documentString, DocumentToNode_XMLROOTSTART, DocumentToNode_XMLROOTEND)
}


func DocumentToNode_Remove(documentString string, start string, end string) string {
        var startIndex int; _ = startIndex
        for (func() int { startIndex = strings.Index(documentString, string(start)); return startIndex }()) != -1 {
            if startIndex >= 0 {
                var endIndex int = documentString.IndexOf(end, startIndex); _ = endIndex
                if endIndex >= 0 {
                    var startString string = documentString[0:startIndex]; _ = startString
                    var resultString string = documentString[endIndex + 1:]; _ = resultString
                    return startString + resultString
                }
            }
        }
        return documentString
}

