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

type GenericProfileActionScriptJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    DefaultMutableTreeNode *DefaultMutableTreeNode
    GenericProfileActionScript *GenericProfileActionScript
    ActionScriptJTree *JTree
    JScrollPane1 *JScrollPane
}

// Implements JTreeInterface


func (this *GenericProfileActionScriptJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.DefaultMutableTreeNode = NewDefaultMutableTreeNode("Root")
    
    
    
}


func NewGenericProfileActionScriptJPanel() *GenericProfileActionScriptJPanel {
    this := &GenericProfileActionScriptJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
    return this
}


func (this *GenericProfileActionScriptJPanel) UpdateJTree() {
        var defaultTreeModel *DefaultTreeModel = NewDefaultTreeModel(this.GetGenericProfileActionScript()); _ = defaultTreeModel
        this.GetActionScriptJTree().SetModel(defaultTreeModel)
        this.Repaint()
        this.LogUtil.PutF("Updated UI", this, "updateJTree")
}


func (this *GenericProfileActionScriptJPanel) GetGenericProfileActionScript() *GenericProfileActionScript {
        return this.GenericProfileActionScript
}


func (this *GenericProfileActionScriptJPanel) SetGenericProfileActionScript(genericProfileActionScript *GenericProfileActionScript) {
        this.GenericProfileActionScript = genericProfileActionScript
}


func (this *GenericProfileActionScriptJPanel) GetActionScriptJTree() *JTree {
        return this.ActionScriptJTree
}


func (this *GenericProfileActionScriptJPanel) SetActionScriptJTree(actionScriptJTree *JTree) {
        this.ActionScriptJTree = actionScriptJTree
}


func (this *GenericProfileActionScriptJPanel) InitComponents() {
        this.JScrollPane1 = NewJScrollPane()
        this.ActionScriptJTree = NewJTree()
        this.SetLayout(NewGridLayout(1, 0))
        this.ActionScriptJTree.SetPreferredSize(NewDimension(500, 500))
        this.ActionScriptJTree.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *GenericProfileActionScriptJPanel) MouseClicked(evt IMouseEvent) {
                this.ActionScriptJTreeMouseClicked(evt)
        }
            

        func (this *GenericProfileActionScriptJPanel) MouseReleased(evt IMouseEvent) {
                this.ActionScriptJTreeMouseReleased(evt)
        }
*/ NewMouseAdapter())
        this.JScrollPane1.SetViewportView(this.ActionScriptJTree)
        this.Add(this.JScrollPane1)
}


func (this *GenericProfileActionScriptJPanel) ActionScriptJTreeMouseReleased(evt IMouseEvent) {
        var button int = evt.GetButton(); _ = button
        if button == MouseEvent_BUTTON3 {
            var treePath *TreePath = this.GetActionScriptJTree().GetSelectionPath(); _ = treePath
            if treePath != nil {
                var anyType any = treePath.GetLastPathComponent(); _ = anyType
                if func() bool { _, ok := any(anyType).(CustomTreeNodeInterface); return ok }() {
                    var customTreeNodeInterface CustomTreeNodeInterface = any(anyType).(CustomTreeNodeInterface); _ = customTreeNodeInterface
                    var jPopupMenu *JPopupMenu = customTreeNodeInterface.GetJPopupMenu(); _ = jPopupMenu
                    jPopupMenu.Show(any(evt.GetSource()).(*JComponent), evt.GetX(), evt.GetY())
                }
            }
        }
}


func (this *GenericProfileActionScriptJPanel) ActionScriptJTreeMouseClicked(evt IMouseEvent) {
}

