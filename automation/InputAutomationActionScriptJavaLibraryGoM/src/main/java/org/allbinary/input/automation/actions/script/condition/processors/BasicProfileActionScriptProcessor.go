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

type BasicProfileActionScriptProcessor struct {
    ProfileActionScriptItem
}

// Implements ProfileActionScriptProcessorInterface

// The Java class BasicProfileActionScriptProcessor has subclasses so its variables are this interface
type IBasicProfileActionScriptProcessor interface {
    IProfileActionScriptItem
    BasicProfileActionScriptProcessor_() *BasicProfileActionScriptProcessor
    ActionPerformed(actionEvent *ActionEvent)
    ToXmlNode(document *Document) *Node
}

func (this *BasicProfileActionScriptProcessor) BasicProfileActionScriptProcessor_() *BasicProfileActionScriptProcessor {
    return this
}





// Java overloads of new BasicProfileActionScriptProcessor
func NewBasicProfileActionScriptProcessor(args ...any) *BasicProfileActionScriptProcessor {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewBasicProfileActionScriptProcessor_1(Java_Arg[string](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicProfileActionScriptProcessor_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new BasicProfileActionScriptProcessor for the arguments")
}
func NewBasicProfileActionScriptProcessor_1(label string, node *Node) *BasicProfileActionScriptProcessor {
    this := &BasicProfileActionScriptProcessor{}
        this.ProfileActionScriptItem = *NewProfileActionScriptItem(label, node)
    return this
}


func NewBasicProfileActionScriptProcessor_2(label string) *BasicProfileActionScriptProcessor {
    this := &BasicProfileActionScriptProcessor{}
        this.ProfileActionScriptItem = *NewProfileActionScriptItem(label)
    return this
}


func (this *BasicProfileActionScriptProcessor) ActionPerformed(actionEvent *ActionEvent) {
        this.ProfileActionScriptItem.ActionPerformed(actionEvent)
        if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_DELETE) == 0 {
            if func() bool { _, ok := any(this.GetParent()).(ProfileActionScriptConditionInterface); return ok }() {
                var profileActionScriptConditionInterface ProfileActionScriptConditionInterface = any(this.GetParent()).(ProfileActionScriptConditionInterface); _ = profileActionScriptConditionInterface
                profileActionScriptConditionInterface.RemoveProcessor(any(this).(ProfileActionScriptProcessorInterface))
                JTreeInterfaceFactory_GetInstance().GetJTreeInterface().UpdateJTree()
            }
        }
}


// throws Exception
func (this *BasicProfileActionScriptProcessor) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionScriptProcessorData_NAME); _ = node
        return node
}

