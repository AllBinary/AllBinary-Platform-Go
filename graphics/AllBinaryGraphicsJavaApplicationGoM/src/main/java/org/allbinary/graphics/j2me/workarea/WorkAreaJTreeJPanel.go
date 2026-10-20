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

type WorkAreaJTreeJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RootTreeNode *DefaultMutableTreeNode
    WorkAreaJTree *JTree
    HighlightedBasicArrayList IBasicArrayList
    CopyJMenuItem *JMenuItem
    DeleteJMenuItem *JMenuItem
    ItemJPopupMenu *JPopupMenu
    ItemJScrollPane1 *JScrollPane
    RotateJMenuItem *JMenuItem
    WorkAreaPropertiesJPanel *JPanel
}


func (this *WorkAreaJTreeJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
}


func NewWorkAreaJTreeJPanel(workAreaName string) *WorkAreaJTreeJPanel {
    this := &WorkAreaJTreeJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
        this.RootTreeNode = NewDefaultMutableTreeNode(workAreaName)
        this.WorkAreaPropertiesJPanel.RemoveAll()
        this.UpdateTree()
        this.HighlightedBasicArrayList = NewBasicArrayListD()
    return this
}


func (this *WorkAreaJTreeJPanel) Add(treeNode *MutableTreeNode) {
        this.RootTreeNode.Add(treeNode)
}


func (this *WorkAreaJTreeJPanel) UpdateTree() {
        this.WorkAreaJTree = NewJTree(this.RootTreeNode)
        this.WorkAreaJTree.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) MousePressed(evt IMouseEvent) {
                this.WorkAreaJTreeMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.WorkAreaPropertiesJPanel.RemoveAll()
        this.WorkAreaPropertiesJPanel.Add(this.WorkAreaJTree)
}


func (this *WorkAreaJTreeJPanel) GetNode_1(treePathArray []*TreePath) IBasicArrayList {
        var basicArrayList IBasicArrayList = NewBasicArrayListD(); _ = basicArrayList
        if treePathArray != nil {
            for index := int(0); index < len(treePathArray); index++ {
                basicArrayList.Add(this.GetNode(treePathArray[index]).(IBasicArrayList))
            }
        }
        return basicArrayList
}

// Java overloads of WorkAreaJTreeJPanel.getNode
func (this *WorkAreaJTreeJPanel) GetNode(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[[]*TreePath](args[0], pass) {
            return this.GetNode_1(Java_Arg[[]*TreePath](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*TreePath](args[0], pass) {
            return this.GetNode_2(Java_Arg[*TreePath](args[0]))
        }
    }
    panic("No Java overload of WorkAreaJTreeJPanel.getNode for the arguments")
}


func (this *WorkAreaJTreeJPanel) GetNode_2(treePath *TreePath) *DefaultMutableTreeNode {
        if treePath != nil {
            var obj []any = treePath.GetPath(); _ = obj
            if obj != nil {
                var node *DefaultMutableTreeNode = any(obj[len(obj) - 1]).(*DefaultMutableTreeNode); _ = node
                return node
            }
        }
        return nil
}


// throws Exception
func (this *WorkAreaJTreeJPanel) DeselectAll() {
        var size int = this.HighlightedBasicArrayList.Size(); _ = size
        this.LogUtil.PutF("size: " + fmt.Sprint(size), this, "deselectAll")
        for index := int(0); index < size; index++ {
            var node *DefaultMutableTreeNode = any(this.HighlightedBasicArrayList.Get(index)).(*DefaultMutableTreeNode); _ = node
            MyGraphicItemEventService_Fire(NewMyGraphicItemEvent(any(NewMyGraphicItemEventSource(MyGraphicItemEventService_DESELECT, node)).(any)))
        }
        this.HighlightedBasicArrayList.Clear()
}


// throws Exception
func (this *WorkAreaJTreeJPanel) SelectGraphicItem(node *DefaultMutableTreeNode) {
        this.HighlightedBasicArrayList.Add(node)
        MyGraphicItemEventService_Fire(NewMyGraphicItemEvent(any(NewMyGraphicItemEventSource(MyGraphicItemEventService_SELECT, node)).(any)))
}


// throws Exception
func (this *WorkAreaJTreeJPanel) RemoveGraphicItem(node *DefaultMutableTreeNode) {
        MyGraphicItemEventService_Fire(NewMyGraphicItemEvent(any(NewMyGraphicItemEventSource(MyGraphicItemEventService_DELETE, node)).(any)))
}


// throws Exception
func (this *WorkAreaJTreeJPanel) RemoveCanvas(node *DefaultMutableTreeNode) {
        MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_DELETE, node)).(any)))
}


// throws Exception
func (this *WorkAreaJTreeJPanel) Remove(treePath *TreePath) {
        var node *DefaultMutableTreeNode = this.GetNode(treePath).(IBasicArrayList); _ = node
        if node != nil {
            if func() bool { _, ok := any(node.GetUserObject()).(*CanvasTreeLabel); return ok }() {
                this.RemoveCanvas(node)
            } else {
                this.RemoveGraphicItem(node)
            }
        }
        var parent *DefaultMutableTreeNode = any(node.GetParent()).(*DefaultMutableTreeNode); _ = parent
        parent.Remove(node)
}


// throws Exception
func (this *WorkAreaJTreeJPanel) Copy_1() {
        var treePath *TreePath = this.WorkAreaJTree.GetSelectionPath(); _ = treePath
        this.Copy(treePath)
        this.UpdateTree()
        this.Expand()
        this.Repaint()
}

// Java overloads of WorkAreaJTreeJPanel.copy
func (this *WorkAreaJTreeJPanel) Copy(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Copy_1()
            return
        }
        if len(args) == 1 && Java_IsArg[*TreePath](args[0], pass) {
            this.Copy_2(Java_Arg[*TreePath](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJTreeJPanel.copy for the arguments")
}


// throws Exception
func (this *WorkAreaJTreeJPanel) Copy_2(treePath *TreePath) {
        var node *DefaultMutableTreeNode = this.GetNode(treePath).(IBasicArrayList); _ = node
        if node != nil {
            if func() bool { _, ok := any(node.GetUserObject()).(*CanvasTreeLabel); return ok }() {
                MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_DUPLICATE, node)).(any)))
            } else {
                MyGraphicItemEventService_Fire(NewMyGraphicItemEvent(any(NewMyGraphicItemEventSource(MyGraphicItemEventService_DUPLICATE, node)).(any)))
            }
        }
}


// throws Exception
func (this *WorkAreaJTreeJPanel) Rotate(treePath *TreePath, angle float64) {
        var node *DefaultMutableTreeNode = this.GetNode(treePath).(IBasicArrayList); _ = node
        if node != nil {
            if func() bool { _, ok := any(node.GetUserObject()).(*CanvasTreeLabel); return ok }() {
                MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_ROTATE, angle, node)).(any)))
            } else {
                MyGraphicItemEventService_Fire(NewMyGraphicItemEvent(any(NewMyGraphicItemEventSource(MyGraphicItemEventService_ROTATE, angle, node)).(any)))
            }
        }
}


func (this *WorkAreaJTreeJPanel) GetRootTreeNode() *MutableTreeNode {
        return this.RootTreeNode
}


func (this *WorkAreaJTreeJPanel) Expand_1() {
        this.WorkAreaJTree.ExpandRow(0)
        this.WorkAreaJTree.ExpandRow(1)
}

// Java overloads of WorkAreaJTreeJPanel.expand
func (this *WorkAreaJTreeJPanel) Expand(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Expand_1()
            return
        }
        if len(args) == 1 && Java_IsArg[*TreePath](args[0], pass) {
            this.Expand_2(Java_Arg[*TreePath](args[0]))
            return
        }
    }
    panic("No Java overload of WorkAreaJTreeJPanel.expand for the arguments")
}


func (this *WorkAreaJTreeJPanel) Expand_2(treePath *TreePath) {
        this.WorkAreaJTree.ExpandPath(treePath)
}


func (this *WorkAreaJTreeJPanel) InitComponents() {
        this.ItemJPopupMenu = NewJPopupMenu()
        this.DeleteJMenuItem = NewJMenuItem()
        this.CopyJMenuItem = NewJMenuItem()
        this.RotateJMenuItem = NewJMenuItem()
        this.ItemJScrollPane1 = NewJScrollPane()
        this.WorkAreaPropertiesJPanel = NewJPanel()
        this.ItemJPopupMenu.AddPopupMenuListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) PopupMenuCanceled(evt *PopupMenuEvent) {
        }
            

        func (this *WorkAreaJTreeJPanel) PopupMenuWillBecomeInvisible(evt *PopupMenuEvent) {
                this.ItemJPopupMenuPopupMenuWillBecomeInvisible(evt)
        }
            

        func (this *WorkAreaJTreeJPanel) PopupMenuWillBecomeVisible(evt *PopupMenuEvent) {
                this.ItemJPopupMenuPopupMenuWillBecomeVisible(evt)
        }
*/ NewPopupMenuListener())
        this.DeleteJMenuItem.SetText("Delete")
        this.DeleteJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) MousePressed(evt IMouseEvent) {
                this.DeleteJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ItemJPopupMenu.Add(this.DeleteJMenuItem)
        this.CopyJMenuItem.SetText("Copy")
        this.CopyJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) MousePressed(evt IMouseEvent) {
                this.CopyJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ItemJPopupMenu.Add(this.CopyJMenuItem)
        this.RotateJMenuItem.SetText("Rotate 45*")
        this.RotateJMenuItem.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) MousePressed(evt IMouseEvent) {
                this.RotateJMenuItemMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.ItemJPopupMenu.Add(this.RotateJMenuItem)
        this.SetLayout(NewGridLayout(1, 1))
        this.WorkAreaPropertiesJPanel.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *WorkAreaJTreeJPanel) MousePressed(evt IMouseEvent) {
                this.WorkAreaPropertiesJPanelMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.WorkAreaPropertiesJPanel.SetLayout(NewGridLayout(1, 0))
        this.ItemJScrollPane1.SetViewportView(this.WorkAreaPropertiesJPanel)
        this.Add(this.ItemJScrollPane1)
}


func (this *WorkAreaJTreeJPanel) RotateJMenuItemMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "rotateJMenuItemMousePressed", e)
                }
            }()
            var treePath *TreePath = this.WorkAreaJTree.GetSelectionPath(); _ = treePath
            this.Rotate(treePath, 45)
            this.UpdateTree()
            this.Expand()
            this.Repaint()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WorkAreaJTreeJPanel) WorkAreaJTreeMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "workAreaJTreeMousePressed", e)
                }
            }()
            var selectedTreePathArray []*TreePath = this.WorkAreaJTree.GetSelectionPaths(); _ = selectedTreePathArray
            var list IBasicArrayList = this.GetNode(selectedTreePathArray).(IBasicArrayList); _ = list
            if list.Size() > 0 {
                if (evt.GetModifiers() & evt.MouseEvent_().BUTTON3_MASK) == evt.MouseEvent_().BUTTON3_MASK {
                    for index := int(0); index < len(selectedTreePathArray); index++ {
                        var defaultMutableTreeNode *DefaultMutableTreeNode = any(list.Get(index)).(*DefaultMutableTreeNode); _ = defaultMutableTreeNode
                        if func() bool { _, ok := any(defaultMutableTreeNode.GetUserObject()).(*CanvasTreeLabel); return ok }() {
                            this.ItemJPopupMenu.Show(this, evt.GetX(), evt.GetY())
                        } else {
                            this.ItemJPopupMenu.Show(this, evt.GetX(), evt.GetY())
                        }
                    }
                }
                if (evt.GetModifiers() & evt.MouseEvent_().BUTTON1_MASK) == evt.MouseEvent_().BUTTON1_MASK {
                    this.DeselectAll()
                    for index := int(0); index < len(selectedTreePathArray); index++ {
                        var defaultMutableTreeNode *DefaultMutableTreeNode = any(list.Get(index)).(*DefaultMutableTreeNode); _ = defaultMutableTreeNode
                        if func() bool { _, ok := any(defaultMutableTreeNode.GetUserObject()).(*CanvasTreeLabel); return ok }() {
                            MyCanvasEventService_Fire(NewMyCanvasEvent(any(NewMyCanvasEventSource(MyCanvasEventService_SELECT, defaultMutableTreeNode)).(any)))
                        } else {
                            this.SelectGraphicItem(defaultMutableTreeNode)
                        }
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WorkAreaJTreeJPanel) ItemJPopupMenuPopupMenuWillBecomeInvisible(evt *PopupMenuEvent) {
}


func (this *WorkAreaJTreeJPanel) ItemJPopupMenuPopupMenuWillBecomeVisible(evt *PopupMenuEvent) {
}


func (this *WorkAreaJTreeJPanel) CopyJMenuItemMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "copyJMenuItemMousePressed", e)
                }
            }()
            this.Copy()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WorkAreaJTreeJPanel) DeleteJMenuItemMousePressed(evt IMouseEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "deleteJMenuItemMousePressed", e)
                }
            }()
            var selectedTreePathArray []*TreePath = this.WorkAreaJTree.GetSelectionPaths(); _ = selectedTreePathArray
            for index := int(0); index < len(selectedTreePathArray); index++ {
                var treePath *TreePath = selectedTreePathArray[index]; _ = treePath
                this.Remove(treePath)
                this.UpdateTree()
                this.Expand()
                this.Repaint()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WorkAreaJTreeJPanel) WorkAreaPropertiesJPanelMousePressed(evt IMouseEvent) {
}

