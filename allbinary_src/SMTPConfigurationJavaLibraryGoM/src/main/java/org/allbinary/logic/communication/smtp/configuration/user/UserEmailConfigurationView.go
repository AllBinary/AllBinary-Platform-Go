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

type UserEmailConfigurationView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    UserEmailConfigurationInterface UserEmailConfigurationInterface
}

// Implements DomNodeInterface


func (this *UserEmailConfigurationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new UserEmailConfigurationView
func NewUserEmailConfigurationView(args ...any) *UserEmailConfigurationView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewUserEmailConfigurationView_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[UserEmailConfigurationInterface](args[0], pass) {
            return NewUserEmailConfigurationView_2(Java_Arg[UserEmailConfigurationInterface](args[0]))
        }
    }
    panic("No Java overload of new UserEmailConfigurationView for the arguments")
}
func NewUserEmailConfigurationView_1(node *Node) *UserEmailConfigurationView {
    this := &UserEmailConfigurationView{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        }
        this.UserEmailConfigurationInterface = any(NewUserEmailConfiguration()).(UserEmailConfigurationInterface)
        var childNodeList *NodeList = node.GetChildNodes(); _ = childNodeList
        var userEmailEventsConfigurationNode *Node = DomSearchHelper_GetNodeNoThrow(UserEmailEventsConfigurationData_NAME, childNodeList); _ = userEmailEventsConfigurationNode
        var userEmailEventsConfigurationView *UserEmailEventsConfigurationView = NewUserEmailEventsConfigurationView(userEmailEventsConfigurationNode); _ = userEmailEventsConfigurationView
        this.UserEmailConfigurationInterface.SetUserEmailEventsConfigurationInterface(userEmailEventsConfigurationView.GetUserEmailEventsConfigurationInterface())
    return this
}


func NewUserEmailConfigurationView_2(userEmailConfigurationInterface UserEmailConfigurationInterface) *UserEmailConfigurationView {
    this := &UserEmailConfigurationView{}
    this.initFields()
        this.UserEmailConfigurationInterface = userEmailConfigurationInterface
    return this
}


func (this *UserEmailConfigurationView) GetEmailConfigurationInterface() UserEmailConfigurationInterface {
        return this.UserEmailConfigurationInterface
}


// throws Exception
func (this *UserEmailConfigurationView) ToXmlNode(document *Document) *Node {
        var emailConfigurationMapping *UserEmailConfigurationMapping = NewUserEmailConfigurationMapping(this.UserEmailConfigurationInterface); _ = emailConfigurationMapping
        var hashMap *HashMap = emailConfigurationMapping.ToHashMap(); _ = hashMap
        var node *Node = ModDomHelper_CreateNameValueNodes(document, UserEmailConfigurationData_NAME, hashMap); _ = node
        var userEmailConfigurationInterface UserEmailConfigurationInterface = this.GetEmailConfigurationInterface(); _ = userEmailConfigurationInterface
        var userEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface = userEmailConfigurationInterface.GetUserEmailEventsConfigurationInterface(); _ = userEmailEventsConfigurationInterface
        var userEmailEventsConfigurationView *UserEmailEventsConfigurationView = NewUserEmailEventsConfigurationView(userEmailEventsConfigurationInterface); _ = userEmailEventsConfigurationView
        node.AppendChild(userEmailEventsConfigurationView.ToXmlNode(document))
        return node
}

