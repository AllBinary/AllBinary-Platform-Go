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

type GenericProfile struct {
    LogUtil *LogUtil
    Name string
    Vector IBasicArrayList
    GenericProfileActions *GenericProfileActions
}


func (this *GenericProfile) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}




// Java overloads of new GenericProfile
func NewGenericProfile(args ...any) *GenericProfile {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewGenericProfile_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewGenericProfile_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new GenericProfile for the arguments")
}
func NewGenericProfile_1(node *Node) *GenericProfile {
    this := &GenericProfile{}
    this.initFields()
        this.SetName(DomNodeHelper_GetTextNodeValue(node))
        this.Vector = NewBasicArrayListD()
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        this.LogUtil.PutF("Name: " + this.GetName() + " Child Nodes: " + fmt.Sprint(nodeList.GetLength()), this, "Contructor")
        for index := int(0); index < nodeList.GetLength(); index++ {
            var profileNode *Node = nodeList.Item(index); _ = profileNode
            if profileNode.GetNodeName().CompareTo(GenericProfileDataWorkerData_NAME) == 0 {
                var genericProfileDataWorkerType IGenericProfileDataWorkerType = GenericProfileDataWorkerType_GetInstance(profileNode); _ = genericProfileDataWorkerType
                this.LogUtil.PutF("Adding GenericProfileDataWorkerType: " + fmt.Sprint(genericProfileDataWorkerType), this, "Contructor")
                this.Vector.Add(genericProfileDataWorkerType)
            }
        }
    return this
}


func NewGenericProfile_2(name string) *GenericProfile {
    this := &GenericProfile{}
    this.initFields()
        this.SetName(name)
        this.Vector = NewBasicArrayListD()
    return this
}


func (this *GenericProfile) GetGenericProfileDataWorkerTypeVector() IBasicArrayList {
        return this.Vector
}


func (this *GenericProfile) Add(genericProfileDataWorkerType IGenericProfileDataWorkerType) {
        if !this.Vector.Contains(genericProfileDataWorkerType) {
            this.Vector.Add(genericProfileDataWorkerType)
        }
}


func (this *GenericProfile) Remove(genericProfileDataWorkerType IGenericProfileDataWorkerType) {
        this.Vector.Remove(genericProfileDataWorkerType)
}


// throws Exception
func (this *GenericProfile) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateTextNode(document, GenericProfileData_NAME, this.Name); _ = node
        var size int = this.Vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var genericProfileDataWorkerType IGenericProfileDataWorkerType = any(this.Vector.Get(index)).(IGenericProfileDataWorkerType); _ = genericProfileDataWorkerType
            node.AppendChild(genericProfileDataWorkerType.ToXmlNode(document))
        }
        return node
}


func (this *GenericProfile) GetName() string {
        return this.Name
}


func (this *GenericProfile) SetName(name string) {
        this.Name = name
}


func (this *GenericProfile) GetGenericProfileActions() *GenericProfileActions {
        return this.GenericProfileActions
}


func (this *GenericProfile) SetGenericProfileActions(genericProfileActions *GenericProfileActions) {
        this.GenericProfileActions = genericProfileActions
}

