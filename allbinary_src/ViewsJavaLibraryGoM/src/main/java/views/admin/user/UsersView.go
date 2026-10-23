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

type UsersView struct {
    HttpContextView
    LogUtil *LogUtil
    UserVector IBasicArrayList
}

// Implements DomNodeInterface

// The Java class UsersView has subclasses so its variables are this interface
type IUsersView interface {
    IHttpContextView
    UsersView_() *UsersView
    ToXmlNode(document *Document) *Node
    AddDomNodeInterfaces()
    View() string
}

func (this *UsersView) UsersView_() *UsersView {
    return this
}


func (this *UsersView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewUsersView(transformInfoInterface TransformInfoInterface) *UsersView {
    this := &UsersView{}
        this.HttpContextView = *NewHttpContextView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *UsersView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var usersNode *Node = document.CreateElement(UsersData_NAME); _ = usersNode
            var size int = this.UserVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var userInterface IUser = any(this.UserVector.Get(index)).(IUser); _ = userInterface
                if userInterface != nil {
                    var node *Node = NewUserDomNode(userInterface).ToXmlNode(document); _ = node
                    usersNode.AppendChild(node)
                } else {
                    this.LogUtil.PutF("UserHelper", this, "toXmlNode")
                }
            }
            tryControl1, tryResult1 = 1, usersNode; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *UsersView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *UsersView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            this.AddDomNodeInterfaces()
            tryControl1, tryResult1 = 1, this.HttpContextView.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

