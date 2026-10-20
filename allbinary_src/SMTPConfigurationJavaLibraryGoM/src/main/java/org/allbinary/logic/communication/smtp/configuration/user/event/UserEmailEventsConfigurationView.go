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

type UserEmailEventsConfigurationView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    UserEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface
}

// Implements DomNodeInterface


func (this *UserEmailEventsConfigurationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewUserEmailEventsConfigurationView(node *Node) *UserEmailEventsConfigurationView {
    this := &UserEmailEventsConfigurationView{}
    this.initFields()
        this.UserEmailEventsConfigurationInterface = any(NewUserEmailEventsConfiguration()).(UserEmailEventsConfigurationInterface)
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        var emailEventConfigurationNodeVector *BasicArrayList = DomSearchHelper_GetAllNodesNoThrow(UserEmailEventConfigurationData_NAME, childNodeList); _ = emailEventConfigurationNodeVector
        var size int = emailEventConfigurationNodeVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userEmailConfigurationNode *Node = any(emailEventConfigurationNodeVector.Get(index)).(*Node); _ = userEmailConfigurationNode
            var userEmailEventConfigurationView *UserEmailEventConfigurationView = NewUserEmailEventConfigurationView(userEmailConfigurationNode); _ = userEmailEventConfigurationView
            this.UserEmailEventsConfigurationInterface.AddUserEmailEventConfiguration(userEmailEventConfigurationView.GetUserEmailEventConfigurationInterface())
        }
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUserEmailEventsConfigurationView(userEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface) *UserEmailEventsConfigurationView {
    this := &UserEmailEventsConfigurationView{}
    this.initFields()
        this.UserEmailEventsConfigurationInterface = userEmailEventsConfigurationInterface
    return this
}


func (this *UserEmailEventsConfigurationView) GetUserEmailEventsConfigurationInterface() UserEmailEventsConfigurationInterface {
        return this.UserEmailEventsConfigurationInterface
}


// throws Exception
func (this *UserEmailEventsConfigurationView) ToXmlNode(document *Document) *Node {
        var userEmailConfigurationMapping *UserEmailEventsConfigurationMapping = NewUserEmailEventsConfigurationMapping(this.UserEmailEventsConfigurationInterface); _ = userEmailConfigurationMapping
        var node *Node = document.CreateElement(UserEmailEventsConfigurationData_NAME); _ = node
        var emailEventHashMap *HashMap = this.GetUserEmailEventsConfigurationInterface().GetEventConfigurationHashMap(); _ = emailEventHashMap
        var set Set = emailEventHashMap.KeySet(); _ = set
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("Number Of Email Events Specified in file: " + fmt.Sprint(set.Size()), this, this.CommonStrings.CONSTRUCTOR)
        }
        var eventNameArray []any = set.ToArray(); _ = eventNameArray
        var size int = len(eventNameArray); _ = size
        for index := int(0); index < size; index++ {
            var eventName string = any(eventNameArray[index]).(string); _ = eventName
            var userEmailEventsConfigurationInterface UserEmailEventConfigurationInterface = any(emailEventHashMap.Get(eventName)).(UserEmailEventConfigurationInterface); _ = userEmailEventsConfigurationInterface
            var userEmailEventConfigurationView *UserEmailEventConfigurationView = NewUserEmailEventConfigurationView(userEmailEventsConfigurationInterface); _ = userEmailEventConfigurationView
            node.AppendChild(userEmailEventConfigurationView.ToXmlNode(document))
        }
        return node
}

