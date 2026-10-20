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

type BasicProfileActionScriptInputNode struct {
}





// Java overloads of new BasicProfileActionScriptInputNode
func NewBasicProfileActionScriptInputNode(args ...any) *BasicProfileActionScriptInputNode {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewBasicProfileActionScriptInputNode_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 0 {
            return NewBasicProfileActionScriptInputNode_2()
        }
    }
    panic("No Java overload of new BasicProfileActionScriptInputNode for the arguments")
}
func NewBasicProfileActionScriptInputNode_1(node *Node) *BasicProfileActionScriptInputNode {
    this := &BasicProfileActionScriptInputNode{}
        var actionNode *Node = DomSearchHelper_GetNode(GenericProfileActionScriptInputData_NAME, node.GetChildNodes()); _ = actionNode
    return this
}


func NewBasicProfileActionScriptInputNode_2() *BasicProfileActionScriptInputNode {
    this := &BasicProfileActionScriptInputNode{}
    return this
}


// throws Exception
func (this *BasicProfileActionScriptInputNode) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionScriptInputData_NAME); _ = node
        return node
}

