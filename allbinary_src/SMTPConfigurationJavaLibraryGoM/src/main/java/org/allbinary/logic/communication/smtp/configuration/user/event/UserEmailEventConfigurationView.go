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

type UserEmailEventConfigurationView struct {
    UserEmailEventConfigurationInterface UserEmailEventConfigurationInterface
}

// Implements DomNodeInterface


func (this *UserEmailEventConfigurationView) initFields() {
    
}


func NewUserEmailEventConfigurationView(node *Node) *UserEmailEventConfigurationView {
    this := &UserEmailEventConfigurationView{}
    this.initFields()
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var name string = DomNodeHelper_GetTextNodeValue(UserEmailEventConfigurationData_NAME, childNodeList); _ = name
        hashMap.Put(UserEmailEventConfigurationData_NAME, name)
        var eventListenerClassPath string = DomNodeHelper_GetTextNodeValue(UserEmailEventConfigurationData_LISTENER_CLASSPATH, childNodeList); _ = eventListenerClassPath
        hashMap.Put(UserEmailEventConfigurationData_LISTENER_CLASSPATH, eventListenerClassPath)
        this.SetUserEmailEventConfigurationInterface(any(NewUserEmailEventConfiguration(hashMap)).(UserEmailEventConfigurationInterface))
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUserEmailEventConfigurationView(userEmailEventConfigurationInterface UserEmailEventConfigurationInterface) *UserEmailEventConfigurationView {
    this := &UserEmailEventConfigurationView{}
    this.initFields()
        this.SetUserEmailEventConfigurationInterface(userEmailEventConfigurationInterface)
    return this
}


// throws Exception
func (this *UserEmailEventConfigurationView) ToXmlNode(document *Document) *Node {
        var userEmailEventConfigurationMapping *UserEmailEventConfigurationMapping = NewUserEmailEventConfigurationMapping(this.GetUserEmailEventConfigurationInterface()); _ = userEmailEventConfigurationMapping
        var hashMap *HashMap = userEmailEventConfigurationMapping.ToHashMap(); _ = hashMap
        var node *Node = ModDomHelper_CreateNameValueNodes(document, UserEmailEventConfigurationData_NAME, hashMap); _ = node
        return node
}


func (this *UserEmailEventConfigurationView) GetUserEmailEventConfigurationInterface() UserEmailEventConfigurationInterface {
        return this.UserEmailEventConfigurationInterface
}


func (this *UserEmailEventConfigurationView) SetUserEmailEventConfigurationInterface(userEmailEventConfigurationInterface UserEmailEventConfigurationInterface) {
        this.UserEmailEventConfigurationInterface = userEmailEventConfigurationInterface
}

