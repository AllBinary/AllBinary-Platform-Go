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

type ContextConfigurationView struct {
    ContextConfigurationInterface ContextConfigurationInterface
}

// Implements DomNodeInterface


func (this *ContextConfigurationView) initFields() {
    
}




// Java overloads of new ContextConfigurationView
func NewContextConfigurationView(args ...any) *ContextConfigurationView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewContextConfigurationView_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[ContextConfigurationInterface](args[0], pass) {
            return NewContextConfigurationView_2(Java_Arg[ContextConfigurationInterface](args[0]))
        }
    }
    panic("No Java overload of new ContextConfigurationView for the arguments")
}
func NewContextConfigurationView_1(node *Node) *ContextConfigurationView {
    this := &ContextConfigurationView{}
    this.initFields()
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        this.ContextConfigurationInterface = NewContextConfiguration()
        var emailConfigurationNode *Node = DomSearchHelper_GetNode(EmailServerConfigurationData_NAME, childNodeList); _ = emailConfigurationNode
        var emailServerConfigurationView *EmailServerConfigurationView = NewEmailServerConfigurationView(emailConfigurationNode); _ = emailServerConfigurationView
        this.ContextConfigurationInterface.SetEmailServerConfigurationInterface(any(emailServerConfigurationView.GetEmailConfigurationInterface()).(EmailServerConfigurationInterface))
    return this
}


func NewContextConfigurationView_2(contextConfigurationInterface ContextConfigurationInterface) *ContextConfigurationView {
    this := &ContextConfigurationView{}
    this.initFields()
        this.ContextConfigurationInterface = contextConfigurationInterface
    return this
}


func (this *ContextConfigurationView) GetContextConfigurationInterface() ContextConfigurationInterface {
        return this.ContextConfigurationInterface
}


// throws Exception
func (this *ContextConfigurationView) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(ContextConfigurationData_GetInstance().NAME); _ = node
        var emailServerConfigurationView *EmailServerConfigurationView = NewEmailServerConfigurationView(this.GetContextConfigurationInterface().GetEmailServerConfigurationInterface()); _ = emailServerConfigurationView
        node.AppendChild(emailServerConfigurationView.ToXmlNode(document))
        return node
}

