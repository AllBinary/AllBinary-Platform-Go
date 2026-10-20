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

type EmailServerConfigurationView struct {
    EmailServerConfigurationInterface EmailServerConfigurationInterface
}

// Implements DomNodeInterface


func (this *EmailServerConfigurationView) initFields() {
    
}


func NewEmailServerConfigurationView(node *Node) *EmailServerConfigurationView {
    this := &EmailServerConfigurationView{}
    this.initFields()
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        var server string = DomNodeHelper_GetTextNodeValue(EmailServerConfigurationData_SERVER, childNodeList); _ = server
        var accountName string = DomNodeHelper_GetTextNodeValue(EmailServerConfigurationData_ACCOUNT, childNodeList); _ = accountName
        var password string = DomNodeHelper_GetTextNodeValue(EmailServerConfigurationData_PASSWORD, childNodeList); _ = password
        this.EmailServerConfigurationInterface = any(NewEmailServerConfiguration(accountName, password, server)).(EmailServerConfigurationInterface)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewEmailServerConfigurationView(emailServerConfigurationInterface EmailServerConfigurationInterface) *EmailServerConfigurationView {
    this := &EmailServerConfigurationView{}
    this.initFields()
        this.EmailServerConfigurationInterface = emailServerConfigurationInterface
    return this
}


func (this *EmailServerConfigurationView) GetEmailConfigurationInterface() EmailServerConfigurationInterface {
        return this.EmailServerConfigurationInterface
}


// throws Exception
func (this *EmailServerConfigurationView) ToXmlNode(document *Document) *Node {
        var emailConfigurationMapping *EmailServerConfigurationMapping = NewEmailServerConfigurationMapping(this.EmailServerConfigurationInterface); _ = emailConfigurationMapping
        var hashMap *HashMap = emailConfigurationMapping.ToHashMap(); _ = hashMap
        var node *Node = ModDomHelper_CreateNodeWithValueNodes(document, EmailServerConfigurationData_NAME, hashMap); _ = node
        return node
}

