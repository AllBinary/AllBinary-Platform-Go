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

type GenericProfileAction struct {
    LogUtil *LogUtil
    Name string
    GenericProfileActionScript *GenericProfileActionScript
    GenericProfileActionJPanel *GenericProfileActionJPanel
}

// Implements DomNodeInterface


func (this *GenericProfileAction) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}




// Java overloads of new GenericProfileAction
func NewGenericProfileAction(args ...any) *GenericProfileAction {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*GenericProfileActionJPanel](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewGenericProfileAction_1(Java_Arg[*GenericProfileActionJPanel](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewGenericProfileAction_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new GenericProfileAction for the arguments")
}
func NewGenericProfileAction_1(genericProfileActionJPanel *GenericProfileActionJPanel, node *Node) *GenericProfileAction {
    this := &GenericProfileAction{}
    this.initFields()
        var actionNameNode *Node = DomSearchHelper_GetNode(GenericProfileActionData_NAME, node.GetChildNodes()); _ = actionNameNode
        if actionNameNode != nil {
            this.Name = DomNodeHelper_GetTextNodeValue(actionNameNode)
        } else {
            panic(errors.New(fmt.Sprint("Profile Action Name Node Null")))
        }
        this.SetGenericProfileActionScript(NewGenericProfileActionScript(this.GenericProfileActionJPanel, node))
    return this
}


func NewGenericProfileAction_2(name string) *GenericProfileAction {
    this := &GenericProfileAction{}
    this.initFields()
        this.Name = name
        this.SetGenericProfileActionScript(NewGenericProfileActionScript(this.GenericProfileActionJPanel))
    return this
}


func (this *GenericProfileAction) GetName() string {
        return this.Name
}


func (this *GenericProfileAction) SetName(name string) {
        this.Name = name
}


func (this *GenericProfileAction) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(GenericProfileActionData_NAME, this.Name)
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *GenericProfileAction) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNodeWithValueNodes(document, GenericProfileActionData_NAME, this.ToHashMap()); _ = node
        node.AppendChild(this.GetGenericProfileActionScript().ToXmlNode(document))
        return node
}


func (this *GenericProfileAction) GetGenericProfileActionScript() *GenericProfileActionScript {
        return this.GenericProfileActionScript
}


func (this *GenericProfileAction) SetGenericProfileActionScript(genericProfileActionScript *GenericProfileActionScript) {
        this.GenericProfileActionScript = genericProfileActionScript
}

