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

type UserConfigurationView struct {
    UserConfigurationInterface UserConfigurationInterface
}

// Implements DomNodeInterface


func (this *UserConfigurationView) initFields() {
    
}


func NewUserConfigurationView(node *Node) *UserConfigurationView {
    this := &UserConfigurationView{}
    this.initFields()
        this.UserConfigurationInterface = any(NewUserConfiguration()).(UserConfigurationInterface)
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        var emailConfigurationNode *Node = DomSearchHelper_GetNode(UserEmailConfigurationData_NAME, childNodeList); _ = emailConfigurationNode
        var userEmailConfigurationView *UserEmailConfigurationView = NewUserEmailConfigurationView(emailConfigurationNode); _ = userEmailConfigurationView
        var userEmailConfigurationInterface UserEmailConfigurationInterface = any(userEmailConfigurationView.GetEmailConfigurationInterface()).(UserEmailConfigurationInterface); _ = userEmailConfigurationInterface
        this.UserConfigurationInterface.SetUserEmailConfigurationInterface(userEmailConfigurationInterface)
        var timeCreatedNode *Node = DomSearchHelper_GetNode(EntryData_GetInstance().TIMECREATED, childNodeList); _ = timeCreatedNode
        var timeCreatedView *TimeCreatedView = NewTimeCreatedView(timeCreatedNode); _ = timeCreatedView
        this.UserConfigurationInterface.SetTimeCreated(timeCreatedView.GetTimeCreated())
        var timeLastModifiedNode *Node = DomSearchHelper_GetNode(EntryData_GetInstance().LASTMODIFIED, childNodeList); _ = timeLastModifiedNode
        var timeLastModifiedView *TimeLastModifiedView = NewTimeLastModifiedView(timeLastModifiedNode); _ = timeLastModifiedView
        this.UserConfigurationInterface.SetTimeLastModified(timeLastModifiedView.GetTimeLastModified())
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUserConfigurationView(userConfigurationInterface UserConfigurationInterface) *UserConfigurationView {
    this := &UserConfigurationView{}
    this.initFields()
        this.UserConfigurationInterface = userConfigurationInterface
    return this
}


func (this *UserConfigurationView) GetUserConfigurationInterface() UserConfigurationInterface {
        return this.UserConfigurationInterface
}


// throws Exception
func (this *UserConfigurationView) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(UserConfigurationData_NAME); _ = node
        var userEmailConfigurationView *UserEmailConfigurationView = NewUserEmailConfigurationView(this.UserConfigurationInterface.GetUserEmailConfigurationInterface()); _ = userEmailConfigurationView
        node.AppendChild(userEmailConfigurationView.ToXmlNode(document))
        var timeCreatedView *TimeCreatedView = NewTimeCreatedView(this.UserConfigurationInterface.GetTimeCreated()); _ = timeCreatedView
        node.AppendChild(timeCreatedView.ToXmlNode(document))
        var timeLastModifiedView *TimeLastModifiedView = NewTimeLastModifiedView(this.UserConfigurationInterface.GetTimeLastModified()); _ = timeLastModifiedView
        node.AppendChild(timeLastModifiedView.ToXmlNode(document))
        return node
}

