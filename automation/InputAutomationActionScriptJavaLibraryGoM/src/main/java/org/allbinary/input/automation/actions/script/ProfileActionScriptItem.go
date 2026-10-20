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

type ProfileActionScriptItem struct {
    DefaultMutableTreeNode
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    JPopupMenu *JPopupMenu
}

// Implements ProfileActionScriptItemInterface, *ActionListener

// The Java class ProfileActionScriptItem has subclasses so its variables are this interface
type IProfileActionScriptItem interface {
    ProfileActionScriptItem_() *ProfileActionScriptItem
    Init()
    Process(frame int64)
    GetJPopupMenu() *JPopupMenu
    ShowDialog()
    ActionPerformed(actionEvent *ActionEvent)
    SetJPopupMenu(jPopupMenu *JPopupMenu)
}

func (this *ProfileActionScriptItem) ProfileActionScriptItem_() *ProfileActionScriptItem {
    return this
}

var ProfileActionScriptItem_EDIT string = "Edit"
var ProfileActionScriptItem_DELETE string = "Delete"
var ProfileActionScriptItem_NEW_CONDITION string = "New Condition"
var ProfileActionScriptItem_NEW_INPUT string = "New Input"
var ProfileActionScriptItem_NEW_OUTPUT string = "New Output"
var ProfileActionScriptItem_NEW_COLOR_AT string = "Color At"
var ProfileActionScriptItem_NEW_TIME_INTERVAL string = "Time Interval"
var ProfileActionScriptItem_NEW_ALWAYS string = "Always On/Off"
var ProfileActionScriptItem_NEW_KEY_INPUT string = "Key"
var ProfileActionScriptItem_NEW_MOUSE_INPUT string = "Mouse"
var ProfileActionScriptItem_NEW_JOYSTICK_INPUT string = "Joystick"
var ProfileActionScriptItem_NEW_IMAGE_OUTPUT string = "Image"

func (this *ProfileActionScriptItem) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new ProfileActionScriptItem
func NewProfileActionScriptItem(args ...any) *ProfileActionScriptItem {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewProfileActionScriptItem_1(Java_Arg[string](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewProfileActionScriptItem_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new ProfileActionScriptItem for the arguments")
}
func NewProfileActionScriptItem_1(label string, node *Node) *ProfileActionScriptItem {
    this := &ProfileActionScriptItem{}
        this.DefaultMutableTreeNode = *NewDefaultMutableTreeNode(label)
        this.initFields()
        this.Init()
        this.LogUtil.PutF("Label: " + label, this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func NewProfileActionScriptItem_2(label string) *ProfileActionScriptItem {
    this := &ProfileActionScriptItem{}
        this.DefaultMutableTreeNode = *NewDefaultMutableTreeNode(label)
        this.initFields()
        this.Init()
        this.LogUtil.PutF("Label: " + label, this, this.CommonStrings.CONSTRUCTOR)
    return this
}


func (this *ProfileActionScriptItem) Init() {
        this.SetJPopupMenu(NewJPopupMenu())
        var jMenuItemEdit *JMenuItem = NewJMenuItem(ProfileActionScriptItem_EDIT); _ = jMenuItemEdit
        var jMenuItemDelete *JMenuItem = NewJMenuItem(ProfileActionScriptItem_DELETE); _ = jMenuItemDelete
        jMenuItemEdit.AddActionListener(this)
        jMenuItemDelete.AddActionListener(this)
        this.GetJPopupMenu().Add(jMenuItemEdit)
        this.GetJPopupMenu().Add(jMenuItemDelete)
}


// throws Exception
func (this *ProfileActionScriptItem) Process(frame int64) {
        panic(errors.New(fmt.Sprint()))
}


func (this *ProfileActionScriptItem) GetJPopupMenu() *JPopupMenu {
        return this.JPopupMenu
}


func (this *ProfileActionScriptItem) ShowDialog() {
        panic(errors.New(fmt.Sprint()))
}


func (this *ProfileActionScriptItem) ActionPerformed(actionEvent *ActionEvent) {
        if actionEvent.GetActionCommand().CompareTo(ProfileActionScriptItem_EDIT) == 0 {
            this.ShowDialog()
        }
}


func (this *ProfileActionScriptItem) SetJPopupMenu(jPopupMenu *JPopupMenu) {
        this.JPopupMenu = jPopupMenu
}

