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

type BasicProfileActionScriptProcessorNode struct {
}





// Java overloads of new BasicProfileActionScriptProcessorNode
func NewBasicProfileActionScriptProcessorNode(args ...any) *BasicProfileActionScriptProcessorNode {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewBasicProfileActionScriptProcessorNode_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewBasicProfileActionScriptProcessorNode_2()
        }
    }
    panic("No Java overload of new BasicProfileActionScriptProcessorNode for the arguments")
}
func NewBasicProfileActionScriptProcessorNode_1(node *Node) *BasicProfileActionScriptProcessorNode {
    this := &BasicProfileActionScriptProcessorNode{}
        var actionNode *Node = DomSearchHelper_GetNode(GenericProfileActionScriptProcessorData_NAME, node.GetChildNodes()); _ = actionNode
    return this
}


func NewBasicProfileActionScriptProcessorNode_2() *BasicProfileActionScriptProcessorNode {
    this := &BasicProfileActionScriptProcessorNode{}
    return this
}


// throws Exception
func (this *BasicProfileActionScriptProcessorNode) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionScriptProcessorData_NAME); _ = node
        return node
}

