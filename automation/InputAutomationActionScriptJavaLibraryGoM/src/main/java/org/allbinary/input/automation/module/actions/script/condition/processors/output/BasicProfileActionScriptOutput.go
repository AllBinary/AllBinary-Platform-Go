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

type BasicProfileActionScriptOutput struct {
    BasicProfileActionScriptProcessor
}

// Implements ProfileActionScriptOutputInterface

// The Java class BasicProfileActionScriptOutput has subclasses so its variables are this interface
type IBasicProfileActionScriptOutput interface {
    IBasicProfileActionScriptProcessor
    BasicProfileActionScriptOutput_() *BasicProfileActionScriptOutput
    ToXmlNode(document *Document) *Node
}

func (this *BasicProfileActionScriptOutput) BasicProfileActionScriptOutput_() *BasicProfileActionScriptOutput {
    return this
}





// Java overloads of new BasicProfileActionScriptOutput
func NewBasicProfileActionScriptOutput(args ...any) *BasicProfileActionScriptOutput {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewBasicProfileActionScriptOutput_1(Java_Arg[string](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicProfileActionScriptOutput_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new BasicProfileActionScriptOutput for the arguments")
}
func NewBasicProfileActionScriptOutput_1(label string, node *Node) *BasicProfileActionScriptOutput {
    this := &BasicProfileActionScriptOutput{}
        this.BasicProfileActionScriptProcessor = *NewBasicProfileActionScriptProcessor(label, node)
    return this
}


func NewBasicProfileActionScriptOutput_2(label string) *BasicProfileActionScriptOutput {
    this := &BasicProfileActionScriptOutput{}
        this.BasicProfileActionScriptProcessor = *NewBasicProfileActionScriptProcessor(label)
    return this
}


// throws Exception
func (this *BasicProfileActionScriptOutput) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionScriptOutputData_NAME); _ = node
        return node
}

